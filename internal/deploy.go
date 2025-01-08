package internal

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"

	erc20mintburntoken "subnet-runner/abi-bindings/go/ictt/ERC20MintBurnToken"
	ics20banktransferapp "subnet-runner/abi-bindings/go/ictt/ICS20/ICS20BankTransferApp"
	tokenrouter "subnet-runner/abi-bindings/go/ictt/ICS20/TokenRouter"
	erc20tokenhomeupgradable "subnet-runner/abi-bindings/go/ictt/TokenHome/ERC20TokenHomeUpgradeable"
	erc20tokenremoteupgradeable "subnet-runner/abi-bindings/go/ictt/TokenRemote/ERC20TokenRemoteUpgradeable"
)

const (
	pk               = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
	pkAwm            = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	AwmAddress       = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	erc20tokenName   = "Landslide"
	erc20tokenSymbol = "transfer/channel-0/stake"
)

// DeploySubnetContracts deploys contracts to the subnet and C-chain and sets up the relayer config
func DeploySubnetContracts(
	log logging.Logger,
	urls []string,
	ibcAddr common.Address,
) error {
	ctx := context.Background()

	baseURL, err := parseBaseURL(urls[0])
	if err != nil {
		return err
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

	// Deploy teleporterMessenger contract
	tpMessengerAddressLnd, tpRegistryAddressLnd, err := deployTeleporter(ctx, log, clientLnd, rpcClientLnd, authLnd, pkey, chainIDLnd)
	if err != nil {
		return err
	}
	log.Info("Teleporter contracts deployed", zap.String("messenger", tpMessengerAddressLnd.Hex()), zap.String("registry", tpRegistryAddressLnd.Hex()))

	// Deploy TokenRouter
	_, routerTx, tokenRouter, err := tokenrouter.DeployTokenRouter(authLnd, clientLnd, authLnd.From)
	if err != nil {
		return fmt.Errorf("failed to deploy TokenRouter: %w", err)
	}
	routerAddr, err := bind.WaitDeployed(ctx, clientLnd, routerTx)
	if err != nil {
		return fmt.Errorf("failed waiting for TokenRouter deployment: %w", err)
	}
	log.Info("TokenRouter deployed", zap.String("address", routerAddr.Hex()))

	// Deploy ICS20BankTransferApp
	_, transferTx, bankTransfer, err := ics20banktransferapp.DeployICS20BankTransferApp(
		authLnd,
		clientLnd,
		ibcAddr,
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
	_, remoteTokenTx, remoteToken, err := erc20tokenremoteupgradeable.DeployERC20TokenRemoteUpgradeable(
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

	// Setup escrow addresses for ICS20BankTransferApp
	if err := setupTransferApp(ctx, &clientLnd, authLnd, bankTransfer, ibcAddr, log); err != nil {
		return fmt.Errorf("failed to setup transfer app: %w", err)
	}

	// Deploy home token at c-chain
	CChainURL := baseURL + "/ext/bc/C/rpc"
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
		erc20tokenName,   // name
		erc20tokenSymbol, // symbol
		18,               // decimals
		authC.From,       // initial owner
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
	_, homeTokenTx, tokenHome, err := erc20tokenhomeupgradable.DeployERC20TokenHomeUpgradeable(
		authC,
		clientC,
		0)
	if err != nil {
		return fmt.Errorf("failed to deploy ERC20TokenHomeUpgradeable: %w", err)
	}

	homeAddr, err := bind.WaitDeployed(ctx, clientC, homeTokenTx)
	if err != nil {
		return fmt.Errorf("failed waiting for ERC20TokenHomeUpgradeable deployment: %w", err)
	}
	log.Info("ERC20TokenHomeUpgradeable deployed to C-chain", zap.String("address", homeAddr.Hex()))

	// Initialize the remote token
	tx, err := tokenHome.Initialize(
		authC,
		tpRegistryAddressC,  // teleporter registry address
		tpMessengerAddressC, // teleporter manager
		big.NewInt(1),       // min teleporter version
		tokenAddr,           // token address
		18,                  // token decimals
	)
	if err != nil {
		return fmt.Errorf("failed to initialize ERC20TokenHomeUpgradeable: %w", err)
	}

	if _, err := bind.WaitMined(ctx, clientC, tx); err != nil {
		return fmt.Errorf("failed waiting for ERC20TokenHomeUpgradeable initialization: %w", err)
	}
	log.Info("Initialized ERC20TokenHomeUpgradeable")

	// Set home address in ERC20 token
	tx, err = erc20Token.SetHomeAddress(authC, homeAddr)
	if err != nil {
		return fmt.Errorf("failed to set home address in ERC20 token: %w", err)
	}
	if _, err := bind.WaitMined(ctx, clientC, tx); err != nil {
		return fmt.Errorf("failed waiting for SetHomeAddress tx: %w", err)
	}
	log.Info("Set home address in ERC20 token")

	// remoteToken, err := erc20tokenremoteupgradeable.NewERC20TokenRemoteUpgradeable(remoteTokenAddr, clientLnd)
	// if err != nil {
	// 	return fmt.Errorf("failed to get ERC20TokenRemoteUpgradeable instance: %w", err)
	// }

	// Init token remote
	tokenHomeBlockchainID, err := getBlockchainID(baseURL, "C-Chain")
	if err != nil {
		return err
	}

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
		erc20tokenName,   // name
		erc20tokenSymbol, // symbol
		uint8(18),        // decimals
		routerAddr,       // tokenRouterChannelReader address
		transferAddr,     // ibcBaseFungibleApp address
		transferAddr,     // transferrer address (ibcBaseFungibleApp)
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
	if tokenName != erc20tokenName {
		return fmt.Errorf("unexpected token name after initialization: got %s, want %s", tokenName, erc20tokenName)
	}

	tokenSymbol, err := remoteToken.Symbol(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get token symbol after initialization: %w", err)
	}
	if tokenSymbol != erc20tokenSymbol {
		return fmt.Errorf("unexpected token symbol after initialization: got %s, want %s", tokenSymbol, erc20tokenSymbol)
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

	// Setup TokenRouter configuration
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

	// fund Awm Address at C-Chain and LND subnet
	if err := fundAddress(ctx, clientC, authC, common.HexToAddress(AwmAddress), pkey, chainIDC, log); err != nil {
		return fmt.Errorf("failed to fund address: %w", err)
	}
	log.Info("Awm Address funded at C-Chain")
	if err := fundAddress(ctx, clientLnd, authLnd, common.HexToAddress(AwmAddress), pkey, chainIDLnd, log); err != nil {
		return fmt.Errorf("failed to fund address: %w", err)
	}
	log.Info("Awm Address funded at LND subnet")

	// Update relayer config
	lndSubnetID, lndChainID, lndChainIDHex, cChainID, err := getBlockchainsInfo(baseURL)
	if err != nil {
		return err
	}
	if err := UpdateRelayerConfig(
		cChainID,
		lndChainID,
		lndSubnetID,
		tpRegistryAddressLnd.Hex(),
		tpMessengerAddressLnd.Hex(),
		pkAwm,
	); err != nil {
		return err
	}
	log.Info("Relayer config updated")

	data := map[string]string{
		"BLOCKCHAIN_ID":          lndChainID,
		"TRANSFER_APP_ADDRESS":   transferAddr.Hex(),
		"SHOW_LOGS_ADDR":         transferAddr.Hex(),
		"ERC20_CONTRACT_ADDRESS": tokenAddr.Hex(),
		"TOKEN_ROUTER_ADDRESS":   routerAddr.Hex(),
		"TOKEN_HOME_ADDRESS":     homeAddr.Hex(),
		"TOKEN_REMOTE_ADDRESS":   remoteTokenAddr.Hex(),
		"BLOCKCHAIN_ID_HEX":      lndChainIDHex,
	}

	if err := SaveToEnvFile("./cmd/app/.env", data); err != nil {
		fmt.Printf("Error saving to .env file: %v\n", err)
	} else {
		fmt.Println("Data saved to .env file successfully")
	}
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
		"transfer/channel-0/stake", // denom
		tokenAddr,                  // token address
		remoteAddr,                 // remote
		homeAddr,                   // home
		"channel-0",                // IBC channel
		18,                         // decimals
		false,                      // isNative
		true,                       // isExternal
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
