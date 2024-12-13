package internal

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ava-labs/subnet-evm/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"go.uber.org/zap"

	receiveronsubnet "subnet-runner/abi-bindings/go/ictt/mocks/ReceiverOnSubnet"
	senderonsubnet "subnet-runner/abi-bindings/go/ictt/mocks/SenderOnSubnet"
	teleportermessenger "subnet-runner/abi-bindings/go/teleporter/TeleporterMessenger"
	teleporterregistry "subnet-runner/abi-bindings/go/teleporter/registry/TeleporterRegistry"
)

// deployTeleporter deploys the teleporterMessenger and teleporterRegistry contracts.
func deployTeleporter(
	ctx context.Context,
	log logging.Logger,
	client ethclient.Client,
	rpcClient *rpc.Client,
	opts *bind.TransactOpts,
	pkey *ecdsa.PrivateKey,
	chainID *big.Int,
) (common.Address, common.Address, error) {
	// Roughly 3,010,000 gas needed to deploy contract. Padded to account for possible additions
	const defaultContractCreationGasLimit = uint64(4000000)

	// R and S values to use in a keyless transaction signature.
	// The values do not technically need to be the same when using Nick's method, but the AvalancheGo
	// APIs by default only allow legacy transactions to be broadcast if they have the same R and S values,
	// which is used as a heuristic to identify (and allow) Nick's method transactions.
	const rsValueHex = "3333333333333333333333333333333333333333333333333333333333333333"

	var (
		vValue = big.NewInt(
			27,
		) // Must be less than 35 to be considered non-EIP155
		defaultContractCreationGasPrice = big.NewInt(2500e9) // 2500 nAVAX/gas
	)

	// Convert the R and S values (which must be the same) from hex.
	rsValue, ok := new(big.Int).SetString(rsValueHex, 16)
	if !ok {
		return common.Address{}, common.Address{}, errors.New(
			"failed to convert R and S value to big.Int",
		)
	}

	// Construct the legacy transaction with pre-determined signature values.
	contractCreationTx := types.NewTx(&types.LegacyTx{
		Nonce:    0, // keyless transaction
		Gas:      defaultContractCreationGasLimit,
		GasPrice: defaultContractCreationGasPrice,
		To:       nil, // Contract creation transaction
		Value:    big.NewInt(0),
		Data:     common.FromHex(teleportermessenger.TeleporterMessengerMetaData.Bin),
		V:        vValue,
		R:        rsValue,
		S:        rsValue,
	})

	// Recover the "sender" address of the transaction.
	senderAddress, err := types.HomesteadSigner{}.Sender(contractCreationTx)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf(
			"failed to recover the sender address of transaction %s", err,
		)
	}

	if err := fundAddress(ctx, client, opts, senderAddress, pkey, chainID, log); err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to fund address: %w", err)
	}

	// Serialize the raw transaction and sender address.
	contractCreationTxBytes, err := contractCreationTx.MarshalBinary()
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf(
			"failed to serialize raw transaction %s", err,
		)
	}

	// Derive the resulting contract address given that it will be deployed from the sender address using the nonce of 0.
	tpContractAddress := crypto.CreateAddress(senderAddress, 0)
	log.Info("Check if the teleporterMessenger contract is already deployed", zap.String("address", tpContractAddress.Hex()))
	deployedCode, err := client.CodeAt(ctx, tpContractAddress, nil)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to check if contract is deployed: %w", err)
	}
	if len(deployedCode) > 0 {
		log.Info("Contract is already deployed", zap.String("address", tpContractAddress.Hex()))
		return common.Address{}, common.Address{}, nil
	}

	txHash := common.Hash{}
	err = rpcClient.CallContext(ctx, &txHash, "eth_sendRawTransaction", hexutil.Encode(contractCreationTxBytes))
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to send transaction: %v", err)
	}

	receipt, err := WaitMinedTxHash(ctx, client, txHash)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to wait for teleporterMessenger transaction mining: %w", err)
	}

	teleporterCode, err := client.CodeAt(ctx, tpContractAddress, nil)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get code at contract address: %w", err)
	}
	if len(teleporterCode) == 0 {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to deploy TeleporterMessenger contract")
	}
	log.Info(
		"TeleporterMessenger contract deployed",
		zap.String("address", tpContractAddress.Hex()),
		zap.String("chainID", chainID.String()),
	)

	tpRegistryAddress := common.Address{}

	// Deploy the TeleporterRegistry contract
	initialEntries := []teleporterregistry.ProtocolRegistryEntry{
		{
			Version:         big.NewInt(1),
			ProtocolAddress: tpContractAddress,
		},
	}
	tpRegistryAddress, tx, _, err := teleporterregistry.DeployTeleporterRegistry(opts, client, initialEntries)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to deploy TeleporterRegistry contract: %v", err)
	}

	// Wait for the deployment transaction to be mined
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to wait for deployment transaction mining: %v", err)
	}

	log.Info(
		"Deployed TeleporterRegistry contract",
		zap.String("address", tpRegistryAddress.Hex()),
		zap.String("txHash", receipt.TxHash.Hex()),
	)

	return tpContractAddress, tpRegistryAddress, nil
}

// fundAddress funds the deployer address with AVAX to cover transaction costs.
func fundAddress(
	ctx context.Context,
	client ethclient.Client,
	auth *bind.TransactOpts,
	deployerAddress common.Address,
	pkey *ecdsa.PrivateKey,
	chainID *big.Int,
	log logging.Logger,
) error {
	nonce, err := client.NonceAt(ctx, auth.From, nil)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %w", err)
	}

	value := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(11)) // 11 AVAX
	gasLimit := uint64(21000)                                    // in units
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %w", err)
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &deployerAddress,
		Value:    value,
	})
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pkey)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	if err := client.SendTransaction(ctx, signedTx); err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	receipt, err := bind.WaitMined(ctx, client, signedTx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction mining: %w", err)
	}
	log.Info("successfully funded", zap.String(
		"address",
		deployerAddress.String(),
	), zap.String("txHash", receipt.TxHash.Hex()))

	return nil
}

// WaitMinedTxHash waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
func WaitMinedTxHash(ctx context.Context, b bind.DeployBackend, txHash common.Hash) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	logger := log.New("hash", txHash)
	for {
		receipt, err := b.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}

		if errors.Is(err, interfaces.NotFound) {
			logger.Trace("Transaction not yet mined")
		} else {
			logger.Trace("Receipt retrieval failed", "err", err)
		}

		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

// parseBaseURL parses the base URL from a raw URL.
func parseBaseURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %s", err)
	}
	return fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host), nil
}

// initializeClientAndAuth initializes an ethclient.Client, bind.TransactOpts, and rpc.Client.
func initializeClientAndAuth(url string, ctx context.Context, pkey *ecdsa.PrivateKey) (
	ethclient.Client,
	*big.Int,
	*bind.TransactOpts,
	*rpc.Client,
	error,
) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to connect to network: %w", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pkey, chainID)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to create auth: %w", err)
	}

	rpcClient, err := rpc.DialContext(ctx, url)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to dial RPC client: %v", err)
	}

	return client, chainID, auth, rpcClient, nil
}

// deployTestReceiver deploys a ReceiverOnSubnet contract.
func deployTestReceiver(ctx context.Context, auth *bind.TransactOpts, client ethclient.Client) (common.Address, error) {
	_, receiverTx, _, err := receiveronsubnet.DeployReceiverOnSubnet(auth, client)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to deploy ReceiverOnSubnet: %w", err)
	}
	receiverAddr, err := bind.WaitDeployed(ctx, client, receiverTx)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed waiting for ReceiverOnSubnet deployment: %w", err)
	}
	return receiverAddr, nil
}

// deployTestSender deploys a SenderOnSubnet contract.
func deployTestSender(ctx context.Context, auth *bind.TransactOpts, client ethclient.Client) (common.Address, error) {
	_, receiverTx, _, err := senderonsubnet.DeploySenderOnSubnet(auth, client)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to deploy SenderOnSubnet: %w", err)
	}
	addr, err := bind.WaitDeployed(ctx, client, receiverTx)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed waiting for SenderOnSubnet deployment: %w", err)
	}
	return addr, nil
}

// getBlockchainsInfo gets the IDs of the LND and C-Chain blockchains.
func getBlockchainsInfo(baseURL string) (string, string, string, error) {
	client := platformvm.NewClient(baseURL)
	blockchains, err := client.GetBlockchains(context.Background())
	if err != nil {
		return "", "", "", fmt.Errorf("failed to get blockchains: %w", err)
	}

	var lndSubnetID, lndChainID, cChainID string
	for _, blockchain := range blockchains {
		fmt.Printf("Blockchain ID: %s, SubnetID: %s, Name: %s\n", blockchain.ID.Hex(), blockchain.SubnetID, blockchain.Name)
		if blockchain.Name == "subnetevm" {
			lndSubnetID = blockchain.SubnetID.String()
			lndChainID = blockchain.ID.String()
		} else if blockchain.Name == "C-Chain" {
			cChainID = blockchain.ID.String()
		}
	}
	return lndSubnetID, lndChainID, cChainID, nil
}
