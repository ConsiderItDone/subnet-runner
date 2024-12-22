package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"

	erc20mintburntoken "subnet-runner/abi-bindings/go/ictt/Erc20MintBurnToken"
	ics20banktransferapp "subnet-runner/abi-bindings/go/ictt/ICS20/ICS20BankTransferApp"
	tokenrouter "subnet-runner/abi-bindings/go/ictt/ICS20/TokenRouter"
	erc20tokenhome "subnet-runner/abi-bindings/go/ictt/TokenHome/ERC20TokenHome"
	erc20tokenremoteupgradeable "subnet-runner/abi-bindings/go/ictt/TokenRemote/ERC20TokenRemoteUpgradeable"
	"subnet-runner/internal/ics20"
)

const (
	pk = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"

	FlagBlockchainIDName  = "bc-id"
	FlagTransferAppName   = "app" // bank transfer app address
	FlagShowLogsAddrName  = "addr"
	FlagErc20ContractName = "erc20"
	FlagHashName          = "hash"
	FlagRouterAddressName = "router"
	FlagHomeAddressName   = "home"
	FlagRemoteAddressName = "remote"

	defaultHashValue          = "0xSomeHash"
	defaultBlockchainID       = "2m11W6dgpvs789P9cYCLDLTrY7ent858A75tq9k7ki9oVwb4oL"
	defaultTransferAppAddress = "0x4Ac1d98D9cEF99EC6546dEd4Bd550b0b287aaD6D"
	defaultShowLogsAddr       = "0x4Ac1d98D9cEF99EC6546dEd4Bd550b0b287aaD6D"
	defaultErc20Contract      = "0x5aa01B3b5877255cE50cc55e8986a7a5fe29C70e" // mint burn token
	defaultHomeAddr           = "0x5DB9A7629912EBF95876228C24A848de0bfB43A9" // token home address
	defaultRemoteAddr         = "0xA4cD3b0Eb6E5Ab5d8CE4065BcCD70040ADAB1F00" // token home address
	defaultTokenRouterAddress = "0x5DB9A7629912EBF95876228C24A848de0bfB43A9"
)

// getClient returns an instance of the ethclient.Client
func getClient(blockchainID string) (ethclient.Client, error) {
	rpc := fmt.Sprintf("http://127.0.0.1:9650/ext/bc/%s/rpc", blockchainID)
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// appInstance returns an instance of the ICS20BankTransferApp contract
func appInstance(c *cli.Context) (*ics20banktransferapp.ICS20BankTransferApp, ethclient.Client, error) {
	client, err := getClient(c.String(FlagBlockchainIDName))
	if err != nil {
		return nil, nil, err
	}

	addr := c.String(FlagTransferAppName)
	if !common.IsHexAddress(addr) {
		return nil, nil, fmt.Errorf("bad app address '%s'", addr)
	}

	app, err := ics20banktransferapp.NewICS20BankTransferApp(common.HexToAddress(addr), client)
	if err != nil {
		return nil, nil, err
	}

	return app, client, nil
}

// sendDirectToTransferApp sends a direct transfer to the transfer app
func sendDirectToTransferApp(
	c *cli.Context,
) error {
	contract, client, err := appInstance(c)
	if err != nil {
		return err
	}

	pkey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %w", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pkey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create auth: %w", err)
	}

	// Get token configuration from TokenRouter
	tokenRouterAddress := common.HexToAddress(c.String(FlagRouterAddressName))
	tokenRouter, err := tokenrouter.NewTokenRouter(tokenRouterAddress, client)
	if err != nil {
		return fmt.Errorf("failed to create token router instance: %w", err)
	}
	// Get token configuration from TokenRouter
	tokenConfig, err := tokenRouter.GetTokenConfig(nil, "transfer/channel-0/stake")
	if err != nil {
		return fmt.Errorf("failed to get token config: %w", err)
	}
	fmt.Printf("Token Config: %+v\n", tokenConfig)

	addressBytes := common.HexToAddress(c.String(FlagShowLogsAddrName)).Bytes()
	// send test onReceivePacket
	packedData, err := ics20.FungibleTokenPacketDataPack(
		"stake",
		big.NewInt(44444),
		[]byte("testSender"),
		addressBytes,
		[]byte("memo-test"),
	)
	if err != nil {
		return err
	}

	packet := ics20banktransferapp.Packet{
		Sequence:           big.NewInt(1),
		SourcePort:         "transfer",
		SourceChannel:      "channel-0",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-0",
		Data:               packedData,
		TimeoutHeight: ics20banktransferapp.Height{
			RevisionNumber: big.NewInt(1),
			RevisionHeight: big.NewInt(10),
		},
		TimeoutTimestamp: big.NewInt(time.Now().UnixNano()),
	}

	tx, err := contract.OnRecvPacket(auth, packet, nil)
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}

	fmt.Printf("TransferApp OnRecvPacket called: tx=%s, receipt=%d\n", tx.Hash().Hex(), receipt.Status)

	return nil
}

// registerRemote registers the remote at the HomeToken
func registerRemote(
	c *cli.Context,
) error {
	client, err := getClient(c.String(FlagBlockchainIDName))
	if err != nil {
		return err
	}

	pkey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %w", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pkey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create auth: %w", err)
	}

	tokenRemoteAddress := common.HexToAddress(c.String(FlagRemoteAddressName))
	tokenRemote, err := erc20tokenremoteupgradeable.NewERC20TokenRemoteUpgradeable(tokenRemoteAddress, client)
	if err != nil {
		return fmt.Errorf("failed to create token router instance: %w", err)
	}

	feeInfo := erc20tokenremoteupgradeable.TeleporterFeeInfo{
		FeeTokenAddress: common.Address{},
		Amount:          big.NewInt(0),
	}

	tx, err := tokenRemote.RegisterWithHome(auth, feeInfo)
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}

	fmt.Printf("Remote Register called: tx=%s, receipt=%d\n", tx.Hash().Hex(), receipt.Status)

	return nil
}

// showAllLogs prints all the logs from the ICS20BankTransferApp contract
func showAllLogs(c *cli.Context) error {
	filterOpts := &bind.FilterOpts{
		Start:   0,   // From block 0
		End:     nil, // to the latest block
		Context: context.Background(),
	}
	client, err := getClient(c.String(FlagBlockchainIDName))
	if err != nil {
		return err
	}

	address := common.HexToAddress(c.String(FlagTransferAppName))
	contract, err := ics20banktransferapp.NewICS20BankTransferApp(address, client)
	if err != nil {
		return fmt.Errorf("failed to create contract instance: %v", err)
	}

	ownershipIterator, err := contract.FilterOwnershipTransferred(filterOpts, nil, nil)
	if err == nil {
		defer ownershipIterator.Close()
		for ownershipIterator.Next() {
			event := ownershipIterator.Event
			fmt.Printf("OwnershipTransferred: PreviousOwner=%s, NewOwner=%s\n", event.PreviousOwner.Hex(), event.NewOwner.Hex())
		}
	}

	packetIterator, err := contract.FilterPacketReceived(filterOpts)
	if err == nil {
		defer packetIterator.Close()
		for packetIterator.Next() {
			event := packetIterator.Event
			fmt.Printf("PacketReceived: Denom=%s, PrefixedDenom=%s, Amount=%s, Sender=%x, Receiver=%x, Memo=%x, MsgSender=%x, IbcAddress=%x\n",
				event.Denom, event.PrefixedDenom, event.Amount.String(), event.Sender, event.Receiver, event.Memo, event.MsgSender, event.IbcAddress)
		}
	}

	tokenRoutingIterator, err := contract.FilterTokenRoutingFailed(filterOpts)
	if err == nil {
		defer tokenRoutingIterator.Close()
		for tokenRoutingIterator.Next() {
			event := tokenRoutingIterator.Event
			fmt.Printf("TokenRoutingFailed: Denom=%s, Reason=%s\n", event.Denom, event.Reason)
		}
	}

	return nil
}

// balanceERC20 prints the balance of the ERC20 token
func balanceERC20(c *cli.Context) error {
	client, err := getClient(c.String(FlagBlockchainIDName))
	if err != nil {
		return err
	}

	address := common.HexToAddress(c.String(FlagShowLogsAddrName))
	contractAddress := common.HexToAddress(c.String(FlagErc20ContractName))

	// Create an instance of the ERC20MintBurnToken contract
	erc20, err := erc20mintburntoken.NewERC20MintBurnToken(contractAddress, client)
	if err != nil {
		return fmt.Errorf("failed to create ERC20MintBurnToken instance: %w", err)
	}

	// Get the balance
	balance, err := erc20.BalanceOf(nil, address)
	if err != nil {
		return fmt.Errorf("failed to get ERC20 balance: %w", err)
	}
	fmt.Printf("ERC20 balance of %s: %s tokens\n", address.Hex(), balance.String())

	return nil
}

func balanceNative(c *cli.Context) error {
	client, err := getClient(c.String(FlagBlockchainIDName))
	if err != nil {
		return err
	}

	address := common.HexToAddress(c.String(FlagShowLogsAddrName))
	// Get the balance
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}
	fmt.Printf("EVM balance of %s: %s wei\n", address.Hex(), balance.String())

	return nil
}

// showTransferAppLogsByHash prints the logs of the transfer app by hash
func showTransferAppLogsByHash(c *cli.Context) error {
	contract, client, err := appInstance(c)
	if err != nil {
		return err
	}

	txHash := common.HexToHash(c.String(FlagHashName))
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return fmt.Errorf("failed to get transaction receipt: %v", err)
	}

	for _, lg := range receipt.Logs {
		event, err := contract.ParsePacketReceived(*lg)
		if err == nil {
			fmt.Printf("PacketReceived: Denom=%s, PrefixedDenom=%s, Amount=%s, Sender=%x, Receiver=%x, Memo=%x, MsgSender=%x, IbcAddress=%x\n",
				event.Denom, event.PrefixedDenom, event.Amount.String(), event.Sender, event.Receiver, event.Memo, event.MsgSender, event.IbcAddress)
		}

		eventOwnership, err := contract.ParseOwnershipTransferred(*lg)
		if err == nil {
			fmt.Printf("OwnershipTransferred: PreviousOwner=%s, NewOwner=%s\n", eventOwnership.PreviousOwner.Hex(), eventOwnership.NewOwner.Hex())
		}

		eventTokenRouting, err := contract.ParseTokenRoutingFailed(*lg)
		if err == nil {
			fmt.Printf("TokenRoutingFailed: Denom=%s, Reason=%s\n", eventTokenRouting.Denom, eventTokenRouting.Reason)
		}
	}

	return nil
}

func showAllLogsTokenHome(c *cli.Context) error {
	client, err := getClient("C")
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(c.String(FlagHomeAddressName))
	contract, err := erc20tokenhome.NewERC20TokenHome(contractAddress, client)
	if err != nil {
		return fmt.Errorf("failed to create contract instance: %v", err)
	}

	filterOpts := &bind.FilterOpts{
		Start:   0,   // Начало блока
		End:     nil, // Конец блока (nil означает текущий блок)
		Context: context.Background(),
	}

	remoteRegisteredIterator, err := contract.FilterRemoteRegistered(filterOpts, nil, nil)
	if err != nil {
		return err
	}
	defer remoteRegisteredIterator.Close()

	for remoteRegisteredIterator.Next() {
		event := remoteRegisteredIterator.Event
		fmt.Printf("RemoteRegistered: %+v\n", event)
	}
	if err := remoteRegisteredIterator.Error(); err != nil {
		return err
	}

	teleporterPausedIterator, err := contract.FilterTeleporterAddressPaused(filterOpts, nil)
	if err != nil {
		return err
	}
	defer teleporterPausedIterator.Close()

	for teleporterPausedIterator.Next() {
		event := teleporterPausedIterator.Event
		fmt.Printf("TeleporterAddressPaused: %+v\n", event)
	}
	if err := teleporterPausedIterator.Error(); err != nil {
		return err
	}

	teleporterUnpausedIterator, err := contract.FilterTeleporterAddressUnpaused(filterOpts, nil)
	if err != nil {
		return err
	}
	defer teleporterUnpausedIterator.Close()

	for teleporterUnpausedIterator.Next() {
		event := teleporterUnpausedIterator.Event
		fmt.Printf("TeleporterAddressUnpaused: %+v\n", event)
	}
	if err := teleporterUnpausedIterator.Error(); err != nil {
		return err
	}

	toCosmosTokensSentIterator, err := contract.FilterToCosmosTokensSent(filterOpts, nil, nil)
	if err != nil {
		return err
	}
	defer toCosmosTokensSentIterator.Close()

	for toCosmosTokensSentIterator.Next() {
		event := toCosmosTokensSentIterator.Event
		fmt.Printf("ToCosmosTokensSent: %+v\n", event)
	}
	if err := toCosmosTokensSentIterator.Error(); err != nil {
		return err
	}

	tokensAndCallRoutedIterator, err := contract.FilterTokensAndCallRouted(filterOpts, nil)
	if err != nil {
		return err
	}
	defer tokensAndCallRoutedIterator.Close()

	for tokensAndCallRoutedIterator.Next() {
		event := tokensAndCallRoutedIterator.Event
		fmt.Printf("TokensAndCallRouted: %+v\n", event)
	}
	if err := tokensAndCallRoutedIterator.Error(); err != nil {
		return err
	}

	tokensAndCallSentIterator, err := contract.FilterTokensAndCallSent(filterOpts, nil, nil)
	if err != nil {
		return err
	}
	defer tokensAndCallSentIterator.Close()

	for tokensAndCallSentIterator.Next() {
		event := tokensAndCallSentIterator.Event
		fmt.Printf("TokensAndCallSent: %+v\n", event)
	}
	if err := tokensAndCallSentIterator.Error(); err != nil {
		return err
	}

	tokensCosmosWithdrawnIterator, err := contract.FilterTokensCosmosWithdrawn(filterOpts, nil)
	if err != nil {
		return err
	}
	defer tokensCosmosWithdrawnIterator.Close()

	for tokensCosmosWithdrawnIterator.Next() {
		event := tokensCosmosWithdrawnIterator.Event
		fmt.Printf("TokensCosmosWithdrawn: %+v\n", event)
	}
	if err := tokensCosmosWithdrawnIterator.Error(); err != nil {
		return err
	}

	tokensRoutedIterator, err := contract.FilterTokensRouted(filterOpts, nil)
	if err != nil {
		return err
	}
	defer tokensRoutedIterator.Close()

	for tokensRoutedIterator.Next() {
		event := tokensRoutedIterator.Event
		fmt.Printf("TokensRouted: %+v\n", event)
	}
	if err := tokensRoutedIterator.Error(); err != nil {
		return err
	}

	tokensSentIterator, err := contract.FilterTokensSent(filterOpts, nil, nil)
	if err != nil {
		return err
	}
	defer tokensSentIterator.Close()

	for tokensSentIterator.Next() {
		event := tokensSentIterator.Event
		fmt.Printf("TokensSent: %+v\n", event)
	}
	if err := tokensSentIterator.Error(); err != nil {
		return err
	}

	tokensWithdrawnIterator, err := contract.FilterTokensWithdrawn(filterOpts, nil)
	if err != nil {
		return err
	}
	defer tokensWithdrawnIterator.Close()

	for tokensWithdrawnIterator.Next() {
		event := tokensWithdrawnIterator.Event
		fmt.Printf("TokensWithdrawn: %+v\n", event)
	}
	if err := tokensWithdrawnIterator.Error(); err != nil {
		return err
	}

	return nil
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using default values")
	}

	// Get values from environment variables or use default values
	blockchainID := getEnv("BLOCKCHAIN_ID", defaultBlockchainID)
	transferAppAddr := getEnv("APP_ADDRESS", defaultTransferAppAddress)
	showLogsAddr := getEnv("SHOW_LOGS_ADDR", defaultShowLogsAddr)
	erc20Contract := getEnv("ERC20_CONTRACT_ADDRESS", defaultErc20Contract)
	routerAddress := getEnv("TOKEN_ROUTER_ADDRESS", defaultTokenRouterAddress)
	homeAddress := getEnv("TOKEN_HOME_ADDRESS", defaultHomeAddr)
	remoteAddress := getEnv("TOKEN_REMOTE_ADDRESS", defaultRemoteAddr)

	// Define flags
	FlagBlockchainID := cli.StringFlag{
		Name:  FlagBlockchainIDName,
		Value: blockchainID,
	}
	FlagTransferApp := cli.StringFlag{
		Name:  FlagTransferAppName,
		Value: transferAppAddr,
	}
	FlagAddr := cli.StringFlag{
		Name:  FlagShowLogsAddrName,
		Value: showLogsAddr,
	}
	FlagContr := cli.StringFlag{
		Name:  FlagErc20ContractName,
		Value: erc20Contract,
	}

	FlagHash := cli.StringFlag{
		Name:  FlagHashName,
		Value: defaultHashValue,
	}
	FlagRouterAddress := cli.StringFlag{
		Name:  FlagRouterAddressName,
		Value: routerAddress,
	}
	FlagHomeAddress := cli.StringFlag{
		Name:  FlagHomeAddressName,
		Value: homeAddress,
	}
	FlagRemoteAddress := cli.StringFlag{
		Name:  FlagRemoteAddressName,
		Value: remoteAddress,
	}

	app := cli.NewApp()
	app.Name = "app"
	app.Usage = "ics20 bank transfer app cli"
	app.Flags = []cli.Flag{
		FlagBlockchainID,
		FlagTransferApp,
		FlagAddr,
		FlagContr,
		FlagHash,
		FlagRouterAddress,
		FlagRemoteAddress,
		FlagHomeAddress,
	}

	app.Commands = []cli.Command{
		{
			Flags:     []cli.Flag{FlagBlockchainID, FlagTransferApp},
			Name:      "show-all-logs",
			Usage:     "Show all logs from the transfer app",
			ShortName: "s",
			Action:    showAllLogs,
		},
		{
			Flags:     []cli.Flag{FlagBlockchainID, FlagAddr},
			Name:      "balance-native",
			Usage:     "Show the balance of the native token",
			ShortName: "n",
			Action:    balanceNative,
		},
		{
			Flags:     []cli.Flag{FlagBlockchainID, FlagAddr, FlagContr},
			Name:      "balance-erc20",
			Usage:     "Show the balance of the ERC20 token",
			ShortName: "e",
			Action:    balanceERC20,
		},
		{
			Flags:     []cli.Flag{FlagBlockchainID, FlagTransferApp, FlagHash},
			Name:      "show-app-logs-by-hash",
			Usage:     "Show the logs of the transfer app by hash",
			ShortName: "h",
			Action:    showTransferAppLogsByHash,
		},
		{
			Flags:     []cli.Flag{FlagBlockchainID, FlagTransferApp, FlagRouterAddress, FlagAddr},
			Name:      "send-direct-transfer",
			ShortName: "t",
			Usage:     "Send a direct transfer from subnets to C-Chain",
			Action:    sendDirectToTransferApp,
		},
		{
			Flags:     []cli.Flag{FlagBlockchainID, FlagHomeAddress},
			Name:      "get-token-home-events",
			ShortName: "g",
			Usage:     "Get all events from HomeToken",
			Action:    showAllLogsTokenHome,
		},
		{
			Flags:     []cli.Flag{FlagBlockchainID, FlagRemoteAddress},
			Name:      "register-remote",
			ShortName: "r",
			Usage:     "Register remote at the HomeToken",
			Action:    registerRemote,
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// getEnv gets the value of the environment variable or returns the default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
