package internal

import (
	"context"
	"fmt"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"

	ics20bank2 "subnet-runner/abi-bindings/go/ictt/ICS20/ICS20Bank"
	ics20banktransferapp "subnet-runner/abi-bindings/go/ictt/ICS20/ICS20BankTransferApp"
	tokenrouter "subnet-runner/abi-bindings/go/ictt/ICS20/TokenRouter"
)

const pk = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"

func DeploySubnetContracts(
	log logging.Logger,
	urls []string,
	ibcAddr common.Address,
) error {
	const (
		mockTeleporterRegistry   = true
		noMockTeleporterRegistry = false
	)

	pkey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %w", err)
	}

	clientLnd, err := ethclient.Dial(urls[0])
	if err != nil {
		return fmt.Errorf("failed to connect to network: %w", err)
	}

	chainIDLnd, err := clientLnd.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}

	authLnd, err := bind.NewKeyedTransactorWithChainID(pkey, chainIDLnd)
	if err != nil {
		return fmt.Errorf("failed to create authLnd: %w", err)
	}
	ctx := context.Background()

	// Deploy Teleporter contract
	rpcClientLnd, err := rpc.DialContext(ctx, urls[0])
	if err != nil {
		return fmt.Errorf("failed to dial RPC client: %v", err)
	}
	defer rpcClientLnd.Close()

	// Deploy teleporterMessenger contract
	tpMessengerAddressLnd, tpRegistryAddressLnd, err := deployTeleporter(
		ctx,
		log,
		clientLnd,
		rpcClientLnd,
		authLnd,
		pkey,
		chainIDLnd,
		noMockTeleporterRegistry,
	)
	if err != nil {
		return err
	}
	log.Info("Teleporter contracts deployed", zap.String("messenger", tpMessengerAddressLnd.Hex()), zap.String("registry", tpRegistryAddressLnd.Hex()))

	// 1. Deploy ICS20Bank
	_, bankTx, _, err := ics20bank2.DeployICS20Bank(authLnd, clientLnd)
	if err != nil {
		return fmt.Errorf("failed to deploy ICS20Bank: %w", err)
	}
	bankAddr, err := bind.WaitDeployed(ctx, clientLnd, bankTx)
	if err != nil {
		return fmt.Errorf("failed waiting for ICS20Bank deployment: %w", err)
	}
	log.Info("ICS20Bank deployed", zap.String("address", bankAddr.Hex()))

	// // 2. Deploy TokenRouter
	// _, routerTx /*tokenRouter*/, _, err := tokenrouter.DeployTokenRouter(authLnd, clientLnd, authLnd.From)
	// if err != nil {
	// 	return fmt.Errorf("failed to deploy TokenRouter: %w", err)
	// }
	// routerAddr, err := bind.WaitDeployed(ctx, clientLnd, routerTx)
	// if err != nil {
	// 	return fmt.Errorf("failed waiting for TokenRouter deployment: %w", err)
	// }
	// log.Info("TokenRouter deployed", zap.String("address", routerAddr.Hex()))
	//
	// // 3. Deploy ICS20BankTransferApp
	// _, transferTx, bankTransfer, err := ics20banktransferapp.DeployICS20BankTransferApp(
	// 	authLnd,
	// 	clientLnd,
	// 	ibcAddr,
	// 	bankAddr,
	// 	authLnd.From,
	// 	routerAddr,
	// )
	// if err != nil {
	// 	return fmt.Errorf("failed to deploy ICS20BankTransferApp: %w", err)
	// }
	// transferAddr, err := bind.WaitDeployed(ctx, clientLnd, transferTx)
	// if err != nil {
	// 	return fmt.Errorf("failed waiting for ICS20BankTransferApp deployment: %w", err)
	// }
	// log.Info("ICS20BankTransferApp deployed", zap.String("address", transferAddr.Hex()))
	//
	// // 4. Deploy remoteToken
	// _, remoteTokenTx, _, err := erc20tokenremoteupgradeable.DeployERC20TokenRemoteUpgradeable(
	// 	authLnd,
	// 	clientLnd,
	// 	0,
	// )
	// if err != nil {
	// 	return fmt.Errorf("failed to deploy ERC20TokenRemoteUpgradeable: %w", err)
	// }
	// remoteTokenAddr, err := bind.WaitDeployed(ctx, clientLnd, remoteTokenTx)
	// if err != nil {
	// 	return fmt.Errorf("failed waiting for ERC20TokenRemoteUpgradeable deployment: %w", err)
	// }
	// log.Info("ERC20TokenRemoteUpgradeable deployed", zap.String("address", remoteTokenAddr.Hex()))
	//
	// // 5. Setup operators for ICS20Bank
	// if err := setOperators(ctx, &clientLnd, authLnd, ics20bank, ibcAddr, transferAddr, log); err != nil {
	// 	return fmt.Errorf("failed to set operators: %w", err)
	// }
	//
	// // 6. Setup escrow addresses for ICS20BankTransferApp
	// if err := setupTransferApp(ctx, &clientLnd, authLnd, bankTransfer, ibcAddr, log); err != nil {
	// 	return fmt.Errorf("failed to setup transfer app: %w", err)
	// }
	//
	//	deploy home token at c-chain
	baseURL, err := parseBaseURL(urls[0])
	if err != nil {
		return err
	}

	CChainURL := baseURL + "/ext/bc/C/rpc"

	clientC, err := ethclient.Dial(CChainURL)
	if err != nil {
		return fmt.Errorf("failed to connect to network: %w", err)
	}

	chainIDC, err := clientC.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}
	fmt.Printf("C-Chain blockchain ID (chainIDC): %s\n", chainIDC)
	authC, err := bind.NewKeyedTransactorWithChainID(pkey, chainIDC)
	if err != nil {
		return fmt.Errorf("failed to create authLnd: %w", err)
	}

	// Deploy Teleporter contract
	rpcClientC, err := rpc.DialContext(ctx, CChainURL)
	if err != nil {
		return fmt.Errorf("failed to dial RPC client: %v", err)
	}
	defer rpcClientC.Close()

	// Deploy teleporterMessenger contract
	tpMessengerAddressC, tpRegistryAddressC, err := deployTeleporter(
		ctx,
		log,
		clientC,
		rpcClientC,
		authC,
		pkey,
		chainIDLnd,
		noMockTeleporterRegistry,
	)
	if err != nil {
		return err
	}
	log.Info(
		"Teleporter contracts deployed to C-chain",
		zap.String("messenger", tpMessengerAddressC.Hex()),
		zap.String("registry", tpRegistryAddressC.Hex()),
	)

	// // todo: deploy tokens at c-chain
	// homeAddr := common.HexToAddress("0x0")
	// tokenAddr := common.HexToAddress("0x0")
	//
	// // 7. Setup TokenRouter configuration
	// if err := setupTokenRouter(ctx, &clientLnd, authLnd, tokenRouter, tokenAddr, homeAddr, remoteTokenAddr, log); err != nil {
	// 	return fmt.Errorf("failed to setup token router: %w", err)
	// }

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
