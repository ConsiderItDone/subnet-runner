package internal

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"

	erc20mintburntoken "subnet-runner/abi-bindings/go/ictt/ERC20MintBurnToken"
	ics20bank2 "subnet-runner/abi-bindings/go/ictt/ICS20/ICS20Bank"
	ics20banktransferapp "subnet-runner/abi-bindings/go/ictt/ICS20/ICS20BankTransferApp"
	tokenrouter "subnet-runner/abi-bindings/go/ictt/ICS20/TokenRouter"
	erc20tokenhome "subnet-runner/abi-bindings/go/ictt/TokenHome/ERC20TokenHome"
	erc20tokenremoteupgradeable "subnet-runner/abi-bindings/go/ictt/TokenRemote/ERC20TokenRemoteUpgradeable"
)

const pk = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"

func DeploySubnetContracts(
	log logging.Logger,
	urls []string,
	ibcAddr common.Address,
) error {
	ctx := context.Background()

	// Connect to the Avalanche node
	client := platformvm.NewClient("http://localhost:9650")

	// Get the list of blockchains
	blockchains, err := client.GetBlockchains(context.Background())
	if err != nil {
		fmt.Errorf("failed to get blockchains: %w", err)
	}

	// Print the transaction IDs
	for _, blockchain := range blockchains {
		fmt.Printf("Blockchain ID: %s, SubnetID: %s, Name: %s\n", blockchain.ID.Hex(), blockchain.SubnetID, blockchain.Name)
	}

	pkey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %w", err)
	}
	clientLnd, chainIDLnd, authLnd, rpcClientLnd, err := initializeClientAndAuth(urls[0], ctx, pkey)
	if err != nil {
		return err
	}
	defer rpcClientLnd.Close()

	// deployer address
	log.Info("Deployer address", zap.String("address", authLnd.From.Hex()))

	// Deploy teleporterMessenger contract
	tpMessengerAddressLnd, tpRegistryAddressLnd, err := deployTeleporter(ctx, log, clientLnd, rpcClientLnd, authLnd, pkey, chainIDLnd)
	if err != nil {
		return err
	}
	log.Info("Teleporter contracts deployed", zap.String("messenger", tpMessengerAddressLnd.Hex()), zap.String("registry", tpRegistryAddressLnd.Hex()))

	// 1. Deploy ICS20Bank
	_, bankTx, ics20bank, err := ics20bank2.DeployICS20Bank(authLnd, clientLnd)
	if err != nil {
		return fmt.Errorf("failed to deploy ICS20Bank: %w", err)
	}
	bankAddr, err := bind.WaitDeployed(ctx, clientLnd, bankTx)
	if err != nil {
		return fmt.Errorf("failed waiting for ICS20Bank deployment: %w", err)
	}
	log.Info("ICS20Bank deployed", zap.String("address", bankAddr.Hex()))

	// 2. Deploy TokenRouter
	_, routerTx, tokenRouter, err := tokenrouter.DeployTokenRouter(authLnd, clientLnd, authLnd.From)
	if err != nil {
		return fmt.Errorf("failed to deploy TokenRouter: %w", err)
	}
	routerAddr, err := bind.WaitDeployed(ctx, clientLnd, routerTx)
	if err != nil {
		return fmt.Errorf("failed waiting for TokenRouter deployment: %w", err)
	}
	log.Info("TokenRouter deployed", zap.String("address", routerAddr.Hex()))

	// 3. Deploy ICS20BankTransferApp
	_, transferTx, bankTransfer, err := ics20banktransferapp.DeployICS20BankTransferApp(
		authLnd,
		clientLnd,
		ibcAddr,
		bankAddr,
		authLnd.From,
		routerAddr,
	)
	if err != nil {
		return fmt.Errorf("failed to deploy ICS20BankTransferApp: %w", err)
	}
	transferAddr, err := bind.WaitDeployed(ctx, clientLnd, transferTx)
	if err != nil {
		return fmt.Errorf("failed waiting for ICS20BankTransferApp deployment: %w", err)
	}
	log.Info("ICS20BankTransferApp deployed", zap.String("address", transferAddr.Hex()))

	// Deploy remoteToken
	_, remoteTokenTx, _, err := erc20tokenremoteupgradeable.DeployERC20TokenRemoteUpgradeable(
		authLnd,
		clientLnd,
		0,
	)
	if err != nil {
		return fmt.Errorf("failed to deploy ERC20TokenRemoteUpgradeable: %w", err)
	}
	remoteTokenAddr, err := bind.WaitDeployed(ctx, clientLnd, remoteTokenTx)
	if err != nil {
		return fmt.Errorf("failed waiting for ERC20TokenRemoteUpgradeable deployment: %w", err)
	}
	log.Info("ERC20TokenRemoteUpgradeable deployed", zap.String("address", remoteTokenAddr.Hex()))

	// Setup operators for ICS20Bank
	if err := setOperators(ctx, &clientLnd, authLnd, ics20bank, ibcAddr, transferAddr, log); err != nil {
		return fmt.Errorf("failed to set operators: %w", err)
	}

	// Setup escrow addresses for ICS20BankTransferApp
	if err := setupTransferApp(ctx, &clientLnd, authLnd, bankTransfer, ibcAddr, log); err != nil {
		return fmt.Errorf("failed to setup transfer app: %w", err)
	}

	// Deploy home token at c-chain
	baseURL, err := parseBaseURL(urls[0])
	if err != nil {
		return err
	}
	CChainURL := baseURL + "/ext/bc/C/rpc"
	cChainBlockchainID, err := info.NewClient(baseURL).GetBlockchainID(ctx, "C")
	if err != nil {
		return fmt.Errorf("failed to get C-chain blockchain ID: %w", err)
	}
	log.Info("C-chain blockchain ID", zap.String("id", cChainBlockchainID.String()))

	clientC, chainIDC, authC, rpcClientC, err := initializeClientAndAuth(CChainURL, ctx, pkey)
	if err != nil {
		return err
	}
	defer rpcClientC.Close()

	// Deploy teleporterMessenger contract
	tpMessengerAddressC, tpRegistryAddressC, err := deployTeleporter(ctx, log, clientC, rpcClientC, authC, pkey, chainIDC)
	if err != nil {
		return err
	}
	log.Info("Teleporter contracts deployed to C-chain", zap.String("messenger", tpMessengerAddressC.Hex()), zap.String("registry", tpRegistryAddressC.Hex()))

	// Deploy ERC20 token
	_, erc20Tx, erc20Token, err := erc20mintburntoken.DeployERC20MintBurnToken(
		authC,
		clientC,
		"Landslide", // name
		"LND",       // symbol
		18,          // decimals
		authC.From,  // initial owner
	)
	if err != nil {
		return fmt.Errorf("failed to deploy ERC20MintBurnToken: %w", err)
	}
	tokenAddr, err := bind.WaitDeployed(ctx, clientC, erc20Tx)
	if err != nil {
		return fmt.Errorf("failed waiting for ERC20MintBurnToken deployment: %w", err)
	}
	log.Info("ERC20MintBurnToken deployed to C-chain", zap.String("address", tokenAddr.Hex()))

	// Deploy ERC20TokenHome
	_, homeTokenTx, _, err := erc20tokenhome.DeployERC20TokenHome(
		authC,
		clientC,
		tpRegistryAddressC,  // teleporter registry address
		tpMessengerAddressC, // teleporter manager
		big.NewInt(1),       // min teleporter version
		tokenAddr,           // token address
		18,                  // token decimals
	)
	if err != nil {
		return fmt.Errorf("failed to deploy ERC20TokenHome: %w", err)
	}
	homeAddr, err := bind.WaitDeployed(ctx, clientC, homeTokenTx)
	if err != nil {
		return fmt.Errorf("failed waiting for ERC20TokenHome deployment: %w", err)
	}
	log.Info("ERC20TokenHome deployed to C-chain", zap.String("address", homeAddr.Hex()))

	// Set home address in ERC20 token
	tx, err := erc20Token.SetHomeAddress(authC, homeAddr)
	if err != nil {
		return fmt.Errorf("failed to set home address in ERC20 token: %w", err)
	}
	if _, err := bind.WaitMined(ctx, clientC, tx); err != nil {
		return fmt.Errorf("failed waiting for SetHomeAddress tx: %w", err)
	}
	log.Info("Set home address in ERC20 token")

	remoteToken, err := erc20tokenremoteupgradeable.NewERC20TokenRemoteUpgradeable(remoteTokenAddr, clientLnd)
	if err != nil {
		return fmt.Errorf("failed to get ERC20TokenRemoteUpgradeable instance: %w", err)
	}

	// Create settings struct for initialization
	tokenHomeBlockchainID := [32]byte{}
	chainIDBytes := chainIDC.Bytes()
	// Copy bytes from right to left to preserve big-endian order
	copy(tokenHomeBlockchainID[32-len(chainIDBytes):], chainIDBytes)

	settings := erc20tokenremoteupgradeable.TokenRemoteSettings{
		TeleporterRegistryAddress: tpRegistryAddressLnd,
		TeleporterManager:         tpMessengerAddressLnd,
		MinTeleporterVersion:      big.NewInt(1),
		TokenHomeBlockchainID:     tokenHomeBlockchainID, // C-chain's chainID as bytes32
		TokenHomeAddress:          homeAddr,
		TokenHomeDecimals:         18,
	}

	// Initialize the remote token
	tx, err = remoteToken.Initialize(
		authLnd,
		settings,
		"Landslide",  // name
		"LND",        // symbol
		uint8(18),    // decimals
		routerAddr,   // tokenRouterChannelReader address
		transferAddr, // ibcBaseFungibleApp address
		transferAddr, // transferrer address (ibcBaseFungibleApp)
	)
	if err != nil {
		return fmt.Errorf("failed to initialize ERC20TokenRemoteUpgradeable: %w", err)
	}

	if _, err := bind.WaitMined(ctx, clientLnd, tx); err != nil {
		return fmt.Errorf("failed waiting for ERC20TokenRemoteUpgradeable initialization: %w", err)
	}
	log.Info("Initialized ERC20TokenRemoteUpgradeable")

	// Verify initialization was successful
	tokenName, err := remoteToken.Name(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get token name after initialization: %w", err)
	}
	if tokenName != "Landslide" {
		return fmt.Errorf("unexpected token name after initialization: got %s, want Landslide", tokenName)
	}

	tokenSymbol, err := remoteToken.Symbol(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get token symbol after initialization: %w", err)
	}
	if tokenSymbol != "LND" {
		return fmt.Errorf("unexpected token symbol after initialization: got %s, want LND", tokenSymbol)
	}

	tokenDecimals, err := remoteToken.Decimals(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get token decimals after initialization: %w", err)
	}
	if tokenDecimals != 18 {
		return fmt.Errorf("unexpected token decimals after initialization: got %d, want 18", tokenDecimals)
	}

	homeBlockchainID, err := remoteToken.GetTokenHomeBlockchainID(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get home blockchain ID after initialization: %w", err)
	}
	if homeBlockchainID != settings.TokenHomeBlockchainID {
		return fmt.Errorf("unexpected home blockchain ID after initialization")
	}

	homeAddress, err := remoteToken.GetTokenHomeAddress(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get home address after initialization: %w", err)
	}
	if homeAddress != settings.TokenHomeAddress {
		return fmt.Errorf("unexpected home address after initialization: got %s, want %s",
			homeAddress.Hex(), settings.TokenHomeAddress.Hex())
	}

	log.Info("Successfully verified ERC20TokenRemoteUpgradeable initialization")

	// 7. Setup TokenRouter configuration
	if err := setupTokenRouter(ctx, &clientLnd, authLnd, tokenRouter, tokenAddr, homeAddr, remoteTokenAddr, log); err != nil {
		return fmt.Errorf("failed to setup token router: %w", err)
	}

	// Deploy ReceiverOnSubnet contract
	receiverAddr, err := deployTestReceiver(ctx, authC, clientC)
	if err != nil {
		return err
	}
	log.Info("ReceiverOnSubnet С deployed", zap.String("address", receiverAddr.Hex()))

	// Deploy ReceiverOnSubnet contract
	senderAddr, err := deployTestSender(ctx, authLnd, clientLnd)
	if err != nil {
		return err
	}
	log.Info("SenderOnSubnet LND deployed", zap.String("address", senderAddr.Hex()))

	return nil
}

func setOperators(
	ctx context.Context,
	client *ethclient.Client,
	auth *bind.TransactOpts,
	bank *ics20bank2.ICS20Bank,
	ibcAddr, transferAddr common.Address,
	log logging.Logger,
) error {
	// Set deployer as operator
	tx1, err := bank.SetOperator(auth, auth.From)
	if err != nil {
		return err
	}
	if _, err := bind.WaitMined(ctx, *client, tx1); err != nil {
		return err
	}
	log.Info("Set deployer as operator")

	// Set IBC contract as operator
	tx2, err := bank.SetOperator(auth, ibcAddr)
	if err != nil {
		return err
	}
	if _, err := bind.WaitMined(ctx, *client, tx2); err != nil {
		return err
	}
	log.Info("Set IBC contract as operator")

	// Set TransferApp as operator
	tx3, err := bank.SetOperator(auth, transferAddr)
	if err != nil {
		return err
	}
	if _, err := bind.WaitMined(ctx, *client, tx3); err != nil {
		return err
	}
	log.Info("Set TransferApp as operator")

	return nil
}

func setupTransferApp(
	ctx context.Context,
	client *ethclient.Client,
	auth *bind.TransactOpts,
	app *ics20banktransferapp.ICS20BankTransferApp,
	ibcAddr common.Address,
	log logging.Logger,
) error {
	// Set escrow address for channel-0
	tx1, err := app.SetChannelEscrowAddresses(auth, "channel-0", auth.From)
	if err != nil {
		return err
	}
	if _, err := bind.WaitMined(ctx, *client, tx1); err != nil {
		return err
	}
	log.Info("Set channel escrow address")

	// Bind transfer port
	tx2, err := app.BindPort(auth, ibcAddr, "transfer")
	if err != nil {
		return err
	}
	if _, err := bind.WaitMined(ctx, *client, tx2); err != nil {
		return err
	}
	log.Info("Bound transfer port")

	return nil
}

func setupTokenRouter(
	ctx context.Context,
	client *ethclient.Client,
	auth *bind.TransactOpts,
	router *tokenrouter.TokenRouter,
	tokenAddr common.Address,
	homeAddr common.Address,
	remoteAddr common.Address,
	log logging.Logger,
) error {
	// Configure wrapped native token
	tx, err := router.SetTokenConfig(
		auth,
		"lnd",       // denom
		tokenAddr,   // token address
		remoteAddr,  // remote
		homeAddr,    // home
		"channel-0", // IBC channel
		18,          // decimals
		false,       // isNative
		true,        // isExternal
	)
	if err != nil {
		return err
	}
	if _, err := bind.WaitMined(ctx, *client, tx); err != nil {
		return err
	}
	log.Info("Configured erc20 token in router")

	return nil
}
