package main

import (
	_ "embed"
	"io"
	"math/big"
	"subnet-runner/cmd/tunnel/awmrelayer/contracts/teleporter/messenger"
	"subnet-runner/cmd/tunnel/awmrelayer/contracts/teleporter/registry"
	"subnet-runner/contracts/ics20/ics20bank"
	"subnet-runner/contracts/ics20/ics20transferer"

	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ava-labs/avalanche-network-runner/local"
	"github.com/ava-labs/avalanche-network-runner/network"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

const (
	healthyTimeout = 2 * time.Minute
)

var (
	FlagAvalancheBinPath = cli.StringFlag{
		Name:  "avalanche-binary-path",
		Value: "/tmp/e2e-test/avalanchego",
	}
	FlagAvalancheWorkDir = cli.StringFlag{
		Name:  "avalanche-working-dir",
		Value: "/tmp/e2e-test/nodes",
	}
	FlagAvalanchePKey = cli.StringFlag{
		Name:  "avalanche-private-key",
		Value: "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027",
	}

	ibcAddr = common.HexToAddress("0x0300000000000000000000000000000000000002")

	logFactory logging.Factory

	//go:embed data/genesis.json
	genesis []byte
)

type (
	AvalancheCallbackFn func(*cli.Context, logging.Logger, network.Network) error
	AvalancheCallback   struct {
		Name    string
		Fn      AvalancheCallbackFn
		IsAsync bool
	}
	AvalancheCallbacks []AvalancheCallback
)

func main() {
	logFactory = logging.NewFactory(logging.Config{
		DisplayLevel: logging.Debug,
		LogLevel:     logging.Debug,
	})

	log, err := logFactory.Make("main")
	if err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Name = "tunnel"
	app.Usage = "e2e test env for token tranfering from C-chain to Cosmos via SubnetEVM"
	app.Flags = []cli.Flag{
		FlagAvalancheBinPath,
		FlagAvalancheWorkDir,
		FlagAvalanchePKey,
	}
	app.Action = mainFn

	if err := app.Run(os.Args); err != nil {
		log.Fatal("can't start app", zap.Error(err))
	}
}

func mainFn(ctx *cli.Context) error {
	os.RemoveAll(ctx.String("avalanche-working-dir"))
	os.MkdirAll(ctx.String("avalanche-working-dir"), 0777)

	var (
		chainid    ids.ID
		bank       *ics20bank.ICS20Bank
		transferer *ics20transferer.ICS20Transferer
	)
	network, err := runAvalanche(ctx, logFactory, AvalancheCallbacks{
		{
			Name: "copy subnet-evm plugin",
			Fn:   runCopyPlugin,
		},
		{
			Name: "create chain",
			Fn: func(ctx *cli.Context, l logging.Logger, n network.Network) error {
				ch, err := runBlockchain(ctx, l, n)
				if err != nil {
					return err
				}
				chainid = ch
				return nil
			},
		},
		{
			Name: "deploy ICS20",
			Fn: func(ctx *cli.Context, l logging.Logger, n network.Network) error {
				b, t, err := runICS20Contracts(chainid, ctx, l, n)
				if err != nil {
					return err
				}
				bank = b
				transferer = t
				return nil
			},
		},
		{
			Name: "run awm-relayer",
			Fn: func(ctx *cli.Context, l logging.Logger, n network.Network) error {
				return runAwmRelayer(chainid, ctx, l, n)
			},
		},
	})
	if err != nil {
		return err
	}

	_ = bank
	_ = transferer

	log, err := logFactory.Make("closer")
	if err != nil {
		return err
	}
	closer(log, func() error {
		return network.Stop(context.Background())
	})

	return nil
}

func closer(log logging.Logger, fn func() error) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	signal.Notify(signalChan, syscall.SIGTERM)

	sig := <-signalChan
	log.Info("got OS signal", zap.Stringer("signal", sig))
	if err := fn(); err != nil {
		log.Info("closer error", zap.Error(err))
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	if err := os.Chmod(dst, 0777); err != nil {
		return 0, err
	}
	return nBytes, err
}

func awaitHealthy(log logging.Logger, nw network.Network, timeout time.Duration) error {
	// Wait until the nodes in the network are ready
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Info("waiting for all nodes to report healthy...")
	err := nw.Healthy(ctx)
	if err == nil {
		log.Info("all nodes healthy...")
	}
	return err
}

func startAwmRelayer(log logging.Logger, cchainRpcURL, subnetRpcURL string, hexPkey string) error {
	{
		pkey, err := crypto.HexToECDSA(hexPkey)
		if err != nil {
			return err
		}
		cclient, err := ethclient.Dial(cchainRpcURL)
		if err != nil {
			return err
		}
		cchainID, err := cclient.ChainID(context.Background())
		if err != nil {
			return err
		}
		cauth, err := bind.NewKeyedTransactorWithChainID(pkey, cchainID)
		if err != nil {
			return err
		}
		_, cchainMesengerTx, cchainMesenger, err := messenger.DeployMessenger(cauth, cclient)
		if err != nil {
			return err
		}
		cchainMesengerAddr, err := bind.WaitDeployed(context.Background(), cclient, cchainMesengerTx)
		if err != nil {
			return err
		}
		log.Info("TeleporterMessenger delpoyed to C chain", zap.String("address", cchainMesengerAddr.Hex()))
		_ = cchainMesenger

		_, cchainRegistryTx, cchainRegistry, err := registry.DeployRegistry(cauth, cclient, []registry.ProtocolRegistryEntry{
			{
				Version:         big.NewInt(1),
				ProtocolAddress: cchainMesengerAddr,
			},
		})
		if err != nil {
			return err
		}
		cchainRegistryAddr, err := bind.WaitDeployed(context.Background(), cclient, cchainRegistryTx)
		if err != nil {
			return err
		}
		log.Info("TeleporterRegistry delpoyed to C chain", zap.String("address", cchainRegistryAddr.Hex()))
		_ = cchainRegistry
	}

	{
		pkey, err := crypto.HexToECDSA(hexPkey)
		if err != nil {
			return err
		}
		subnetclient, err := ethclient.Dial(subnetRpcURL)
		if err != nil {
			return err
		}
		subnetChainID, err := subnetclient.ChainID(context.Background())
		if err != nil {
			return err
		}

		subnetauth, err := bind.NewKeyedTransactorWithChainID(pkey, subnetChainID)
		if err != nil {
			return err
		}
		_, subnetMesengerTx, subnetMesenger, err := messenger.DeployMessenger(subnetauth, subnetclient)
		if err != nil {
			return err
		}
		subnetMesengerAddr, err := bind.WaitDeployed(context.Background(), subnetclient, subnetMesengerTx)
		if err != nil {
			return err
		}
		log.Info("TeleporterMessenger delpoyed to subnet", zap.String("address", subnetMesengerAddr.Hex()))
		_ = subnetMesenger

		_, subnetRegistryTx, subnetRegistry, err := registry.DeployRegistry(subnetauth, subnetclient, []registry.ProtocolRegistryEntry{
			{
				Version:         big.NewInt(1),
				ProtocolAddress: subnetMesengerAddr,
			},
		})
		if err != nil {
			return err
		}
		subnetRegistryAddr, err := bind.WaitDeployed(context.Background(), subnetclient, subnetRegistryTx)
		if err != nil {
			return err
		}
		log.Info("TeleporterRegistry delpoyed to subnet", zap.String("address", subnetRegistryAddr.Hex()))
		_ = subnetRegistry
	}

	return nil
}

func runAwmRelayer(bchainid ids.ID, ctx *cli.Context, log logging.Logger, nw network.Network) error {
	nodeNames, err := nw.GetNodeNames()
	if err != nil {
		return err
	}
	if len(nodeNames) == 0 {
		return fmt.Errorf("node list length must be more than zero")
	}
	node, err := nw.GetNode(nodeNames[0])
	if err != nil {
		return err
	}
	log.Info(
		"node",
		zap.String("node name", nodeNames[0]),
		zap.String("cchain", fmt.Sprintf("http://127.0.0.1:%d/ext/bc/C/rpc", node.GetAPIPort())),
		zap.String("subnet", fmt.Sprintf("http://127.0.0.1:%d/ext/bc/%s/rpc", node.GetAPIPort(), bchainid)),
	)

	return startAwmRelayer(
		log,
		fmt.Sprintf("http://127.0.0.1:%d/ext/bc/C/rpc", node.GetAPIPort()),
		fmt.Sprintf("http://127.0.0.1:%d/ext/bc/%s/rpc", node.GetAPIPort(), bchainid),
		ctx.String("avalanche-private-key"),
	)
}

func runICS20Contracts(bchainid ids.ID, ctx *cli.Context, log logging.Logger, nw network.Network) (*ics20bank.ICS20Bank, *ics20transferer.ICS20Transferer, error) {
	nodeNames, err := nw.GetNodeNames()
	if err != nil {
		return nil, nil, err
	}
	if len(nodeNames) == 0 {
		return nil, nil, fmt.Errorf("node list length must be more than zero")
	}
	node, err := nw.GetNode(nodeNames[0])
	if err != nil {
		return nil, nil, err
	}
	client, err := ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%d/ext/bc/%s/rpc", node.GetAPIPort(), bchainid))
	if err != nil {
		return nil, nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, nil, err
	}

	pkey, err := crypto.HexToECDSA(ctx.String("avalanche-private-key"))
	if err != nil {
		return nil, nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pkey, chainID)
	if err != nil {
		return nil, nil, err
	}

	_, ics20bankTx, ics20bank, err := ics20bank.DeployICS20Bank(auth, client)
	if err != nil {
		return nil, nil, err
	}

	ics20bankAddr, err := bind.WaitDeployed(context.Background(), client, ics20bankTx)
	if err != nil {
		return nil, nil, err
	}
	log.Info("ICS20 Bank delpoyed", zap.String("address", ics20bankAddr.Hex()))

	_, ics20transfererTx, ics20transferer, err := ics20transferer.DeployICS20Transferer(auth, client, ibcAddr, ics20bankAddr)
	if err != nil {
		return nil, nil, err
	}

	ics20transfererAddr, err := bind.WaitDeployed(context.Background(), client, ics20transfererTx)
	if err != nil {
		return nil, nil, err
	}
	log.Info("ICS20 Transferer delpoyed", zap.String("address", ics20transfererAddr.Hex()))

	setOperTx1, err := ics20bank.SetOperator(auth, auth.From)
	if err != nil {
		return nil, nil, err
	}
	setOperRe1, err := bind.WaitMined(context.Background(), client, setOperTx1)
	if err != nil {
		return nil, nil, err
	}
	log.Info("SetOperator key address", zap.String("hash", setOperRe1.TxHash.Hex()), zap.String("block", setOperRe1.BlockNumber.String()))

	setOperTx2, err := ics20bank.SetOperator(auth, ibcAddr)
	if err != nil {
		return nil, nil, err
	}
	setOperRe2, err := bind.WaitMined(context.Background(), client, setOperTx2)
	if err != nil {
		return nil, nil, err
	}
	log.Info("SetOperator ibc address", zap.String("hash", setOperRe2.TxHash.Hex()), zap.String("block", setOperRe2.BlockNumber.String()))

	setOperTx3, err := ics20bank.SetOperator(auth, ics20transfererAddr)
	if err != nil {
		return nil, nil, err
	}
	setOperRe3, err := bind.WaitMined(context.Background(), client, setOperTx3)
	if err != nil {
		return nil, nil, err
	}
	log.Info("SetOperator ics20 transferer address", zap.String("hash", setOperRe3.TxHash.Hex()), zap.String("block", setOperRe3.BlockNumber.String()))

	setChannelEscrowAddressesTx, err := ics20transferer.SetChannelEscrowAddresses(auth, "transfer", auth.From)
	if err != nil {
		return nil, nil, err
	}
	setChannelEscrowAddressesRe, err := bind.WaitMined(context.Background(), client, setChannelEscrowAddressesTx)
	if err != nil {
		return nil, nil, err
	}
	log.Info("ics20transferer.SetChannelEscrowAddresses", zap.String("addr", auth.From.Hex()), zap.String("port", "transfer"), zap.String("block", setChannelEscrowAddressesRe.BlockNumber.String()))

	bintPortTx, err := ics20transferer.BindPort(auth, ibcAddr, "transfer")
	if err != nil {
		return nil, nil, err
	}
	bintPortRe, err := bind.WaitMined(context.Background(), client, bintPortTx)
	if err != nil {
		return nil, nil, err
	}
	log.Info("ics20transferer.BindPort", zap.String("addr", ibcAddr.Hex()), zap.String("port", "transfer"), zap.String("block", bintPortRe.BlockNumber.String()))

	return ics20bank, ics20transferer, nil
}

func runBlockchain(ctx *cli.Context, log logging.Logger, nw network.Network) (ids.ID, error) {
	nodeNames, err := nw.GetNodeNames()
	if err != nil {
		return ids.Empty, err
	}

	config := []byte(`{"warp-api-enabled": true,"log-level": "debug"}`)
	chains, err := nw.CreateBlockchains(context.Background(), []network.BlockchainSpec{
		{
			VMName:      "subnetevm",
			Genesis:     genesis,
			ChainConfig: config,
			SubnetSpec: &network.SubnetSpec{
				SubnetConfig: config,
				Participants: nodeNames,
			},
		},
	})
	if err != nil {
		return ids.Empty, err
	}

	if len(chains) == 0 {
		return ids.Empty, fmt.Errorf("chain didn't created")
	}

	// Wait until the nodes in the network are ready
	if err := awaitHealthy(log, nw, healthyTimeout); err != nil {
		return ids.Empty, err
	}
	return chains[0], nil
}

func runCopyPlugin(ctx *cli.Context, log logging.Logger, nw network.Network) error {
	nodeNames, err := nw.GetNodeNames()
	if err != nil {
		return err
	}

	for i := range nodeNames {
		node, err := nw.GetNode(nodeNames[i])
		if err != nil {
			return err
		}
		if _, err := copy(
			fmt.Sprintf("%s/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy", ctx.String("avalanche-binary-path")),
			fmt.Sprintf("%s/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy", node.GetDataDir()),
		); err != nil {
			return err
		}
	}

	return nil
}

func runAvalanche(ctx *cli.Context, logFactory logging.Factory, callbacks AvalancheCallbacks) (network.Network, error) {
	// Create the network
	nwConfig, err := local.NewDefaultConfig(fmt.Sprintf("%s/avalanchego", ctx.String("avalanche-binary-path")))
	if err != nil {
		return nil, err
	}
	nwConfig.Flags["log-level"] = "INFO"

	log, err := logFactory.Make("avalanche")
	if err != nil {
		return nil, err
	}

	nw, err := local.NewNetwork(log, nwConfig, ctx.String("avalanche-working-dir"), "", true, true, true)
	if err != nil {
		return nil, err
	}
	if err := awaitHealthy(log, nw, healthyTimeout); err != nil {
		return nil, err
	}

	for i := range callbacks {
		log, err := logFactory.Make(callbacks[i].Name)
		if err != nil {
			return nil, err
		}

		log.Info(fmt.Sprintf("%s started", callbacks[i].Name))
		if callbacks[i].IsAsync {
			go func(log logging.Logger, fn AvalancheCallbackFn) {
				if err := callbacks[i].Fn(ctx, log, nw); err != nil {
					log.Error("callback error", zap.Error(err))
				}
			}(log, callbacks[i].Fn)
		} else {
			if err := callbacks[i].Fn(ctx, log, nw); err != nil {
				log.Error("callback error", zap.Error(err))
			}
		}
		log.Info(fmt.Sprintf("%s ended", callbacks[i].Name))
	}

	return nw, nil
}

/*
import (
	_ "embed"

	"context"
	"fmt"
	"go/build"
	"io"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ava-labs/avalanche-network-runner/local"
	"github.com/ava-labs/avalanche-network-runner/network"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"go.uber.org/zap"

	"subnet-runner/contracts/ics20/ics20bank"
	"subnet-runner/contracts/ics20/ics20transferer"
)

const (
	healthyTimeout = 2 * time.Minute
)

var (
	goPath = os.ExpandEnv("$GOPATH")

	ibcAddr = common.HexToAddress("0x0300000000000000000000000000000000000002")

	//go:embed data/genesis.json
	genesis []byte
)

// Shows example usage of the Avalanche Network Runner.
// Creates a local five node Avalanche network
// and waits for all nodes to become healthy.
// The network runs until the user provides a SIGINT or SIGTERM.
func main() {
	// Create the logger
	logFactory := logging.NewFactory(logging.Config{
		DisplayLevel: logging.Info,
		LogLevel:     logging.Info,
	})
	log, err := logFactory.Make("main")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if goPath == "" {
		goPath = build.Default.GOPATH
	}

	binaryPath := "/tmp/e2e-test/avalanchego"
	workDir := "/tmp/e2e-test/nodes"

	os.RemoveAll(workDir)
	os.MkdirAll(workDir, 0777)

	if err := run(log, binaryPath, workDir); err != nil {
		log.Fatal("fatal error", zap.Error(err))
		os.Exit(1)
	}
}

func run(log logging.Logger, binaryPath string, workDir string) error {
	// Create the network
	nwConfig, err := local.NewDefaultConfig(fmt.Sprintf("%s/avalanchego", binaryPath))
	if err != nil {
		return err
	}

	nwConfig.Flags["log-level"] = "INFO"

	nw, err := local.NewNetwork(log, nwConfig, workDir, "", true, false, true)
	if err != nil {
		return err
	}
	defer func() { // Stop the network when this function returns
		if err := nw.Stop(context.Background()); err != nil {
			log.Info("error stopping network", zap.Error(err))
		}
	}()

	// When we get a SIGINT or SIGTERM, stop the network and close [closedOnShutdownCh]
	signalsChan := make(chan os.Signal, 1)
	signal.Notify(signalsChan, syscall.SIGINT)
	signal.Notify(signalsChan, syscall.SIGTERM)
	closedOnShutdownCh := make(chan struct{})
	go func() {
		shutdownOnSignal(log, nw, signalsChan, closedOnShutdownCh)
	}()

	// Wait until the nodes in the network are ready
	if err := awaitHealthy(nw, log, healthyTimeout); err != nil {
		return err
	}

	// Add some chain
	nodeNames, err := nw.GetNodeNames()
	if err != nil {
		return err
	}

	for i := range nodeNames {
		node, err := nw.GetNode(nodeNames[i])
		if err != nil {
			return err
		}
		if _, err := copyFile(
			fmt.Sprintf("%s/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy", binaryPath),
			fmt.Sprintf("%s/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy", node.GetDataDir()),
		); err != nil {
			return err
		}
	}

	chains, err := nw.CreateBlockchains(context.Background(), []network.BlockchainSpec{
		{
			VMName:      "subnetevm",
			Genesis:     genesis,
			ChainConfig: []byte(`{"warp-api-enabled": true}`),
			SubnetSpec: &network.SubnetSpec{
				SubnetConfig: nil,
				Participants: nodeNames,
			},
		},
	})
	if err != nil {
		return err
	}

	// Wait until the nodes in the network are ready
	if err := awaitHealthy(nw, log, healthyTimeout); err != nil {
		return err
	}

	//rpcUrls := make([]string, len(nodeNames))
	//for i := range nodeNames {
	//	node, err := nw.GetNode(nodeNames[i])
	//	if err != nil {
	//		return err
	//	}
	//	rpcUrls[i] = fmt.Sprintf("http://127.0.0.1:%d/ext/bc/%s/rpc", node.GetAPIPort(), chains[0])
	//	log.Info("subnet rpc url", zap.String("node", nodeNames[i]), zap.String("url", rpcUrls[i]))
	//}
	//
	//log.Info("Network will run until you CTRL + C to exit...")
	//
	//if err := doICS20(log, rpcUrls); err != nil {
	//	log.Error("can't deploy ICS20", zap.Error(err))
	//	return err
	//}
	//
	//if err := deployContracts(log, no); err != nil {
	//	log.Error("can't deploy teleporter contracts", zap.Error(err))
	//	return err
	//}
	//
	//for {
	//	time.Sleep(5 * time.Second)
	//	select {
	//	case <-closedOnShutdownCh:
	//		return nil
	//	case <-time.After(5 * time.Second):
	//		if err := doTx(log, rpcUrls); err != nil {
	//			return err
	//		}
	//	}
	//}
}

// Blocks until a signal is received on [signalChan], upon which
// [n.Stop()] is called. If [signalChan] is closed, does nothing.
// Closes [closedOnShutdownChan] amd [signalChan] when done shutting down network.
// This function should only be called once.
func shutdownOnSignal(
	log logging.Logger,
	n network.Network,
	signalChan chan os.Signal,
	closedOnShutdownChan chan struct{},
) {
	sig := <-signalChan
	log.Info("got OS signal", zap.Stringer("signal", sig))

	if err := n.Stop(context.Background()); err != nil {
		log.Info("error stopping network", zap.Error(err))
	}
	signal.Reset()
	close(signalChan)
	close(closedOnShutdownChan)
}

func awaitHealthy(nw network.Network, log logging.Logger, timeout time.Duration) error {
	// Wait until the nodes in the network are ready
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	log.Info("waiting for all nodes to report healthy...")
	err := nw.Healthy(ctx)
	if err == nil {
		log.Info("all nodes healthy...")
	}
	return err
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	if err := os.Chmod(dst, 0777); err != nil {
		return 0, err
	}
	return nBytes, err
}

func doICS20(log logging.Logger, urls []string) error {
	pkey, err := crypto.HexToECDSA("56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027")
	if err != nil {
		return err
	}

	client, err := ethclient.Dial(urls[0])
	if err != nil {
		return err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pkey, chainID)
	if err != nil {
		return err
	}

	_, ics20bankTx, ics20bank, err := ics20bank.DeployICS20Bank(auth, client)
	if err != nil {
		return err
	}

	ics20bankAddr, err := bind.WaitDeployed(context.Background(), client, ics20bankTx)
	if err != nil {
		return err
	}
	log.Info("ICS20 Bank delpoyed", zap.String("address", ics20bankAddr.Hex()))

	_, ics20transfererTx, ics20transferer, err := ics20transferer.DeployICS20Transferer(auth, client, ibcAddr, ics20bankAddr)
	if err != nil {
		return err
	}

	ics20transfererAddr, err := bind.WaitDeployed(context.Background(), client, ics20transfererTx)
	if err != nil {
		return err
	}
	log.Info("ICS20 Transferer delpoyed", zap.String("address", ics20transfererAddr.Hex()))

	setOperTx1, err := ics20bank.SetOperator(auth, auth.From)
	if err != nil {
		return err
	}
	setOperRe1, err := bind.WaitMined(context.Background(), client, setOperTx1)
	if err != nil {
		return err
	}
	log.Info("SetOperator key address", zap.String("hash", setOperRe1.TxHash.Hex()), zap.String("block", setOperRe1.BlockNumber.String()))

	setOperTx2, err := ics20bank.SetOperator(auth, ibcAddr)
	if err != nil {
		return err
	}
	setOperRe2, err := bind.WaitMined(context.Background(), client, setOperTx2)
	if err != nil {
		return err
	}
	log.Info("SetOperator ibc address", zap.String("hash", setOperRe2.TxHash.Hex()), zap.String("block", setOperRe2.BlockNumber.String()))

	setOperTx3, err := ics20bank.SetOperator(auth, ics20transfererAddr)
	if err != nil {
		return err
	}
	setOperRe3, err := bind.WaitMined(context.Background(), client, setOperTx3)
	if err != nil {
		return err
	}
	log.Info("SetOperator ics20 transferer address", zap.String("hash", setOperRe3.TxHash.Hex()), zap.String("block", setOperRe3.BlockNumber.String()))

	setChannelEscrowAddressesTx, err := ics20transferer.SetChannelEscrowAddresses(auth, "transfer", auth.From)
	if err != nil {
		return err
	}
	setChannelEscrowAddressesRe, err := bind.WaitMined(context.Background(), client, setChannelEscrowAddressesTx)
	if err != nil {
		return err
	}
	log.Info("ics20transferer.SetChannelEscrowAddresses", zap.String("addr", auth.From.Hex()), zap.String("port", "transfer"), zap.String("block", setChannelEscrowAddressesRe.BlockNumber.String()))

	bintPortTx, err := ics20transferer.BindPort(auth, ibcAddr, "transfer")
	if err != nil {
		return err
	}
	bintPortRe, err := bind.WaitMined(context.Background(), client, bintPortTx)
	if err != nil {
		return err
	}
	log.Info("ics20transferer.BindPort", zap.String("addr", ibcAddr.Hex()), zap.String("port", "transfer"), zap.String("block", bintPortRe.BlockNumber.String()))

	return nil
}

func deployContracts(log logging.Logger, nw network.Network, subnetChainId ids.ID) error {
	nodeNames, err := nw.GetNodeNames()
	if err != nil {
		return err
	}
	if len(nodeNames) == 0 {
		return fmt.Errorf("network doesn't have nodes")
	}
	node, err := nw.GetNode(nodeNames[0])
	if err != nil {
		return err
	}

	cchainHttpRPC, err := rpc.Dial(fmt.Sprintf("http://127.0.0.1:%d/ext/bc/C/rpc", node.GetAPIPort()))
	if err != nil {
		return err
	}

	cchainWsRPC, err := rpc.Dial(fmt.Sprintf("ws://127.0.0.1:%d/ext/bc/C/rpc", node.GetAPIPort()))
	if err != nil {
		return err
	}

	subnetHttpRPC, err := rpc.Dial(fmt.Sprintf("http://127.0.0.1:%d/ext/bc/%s/rpc", node.GetAPIPort(), subnetChainId))
	if err != nil {
		return err
	}

	cchainHttpClient := ethclient.NewClient(cchainHttpRPC)
	cchainWslient := ethclient.NewClient(cchainWsRPC)
	subnetHttpClient := ethclient.NewClient(subnetHttpRPC)

	_ = cchainHttpClient
	_ = cchainWslient
	_ = subnetHttpClient

	teleportermessenger.New

	return nil
}

func doTx(log logging.Logger, urls []string) error {
	pkey, err := crypto.HexToECDSA("56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027")
	if err != nil {
		return err
	}
	addr := crypto.PubkeyToAddress(pkey.PublicKey)

	client, err := ethclient.Dial(urls[0])
	if err != nil {
		return err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return err
	}

	nonce, err := client.NonceAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	toAddress := common.HexToAddress("0xAB259A4830E2C7AA6EF3831BAC1590F855AE4C32")
	value := big.NewInt(1000000000000000000)
	tx := types.NewTransaction(nonce, toAddress, value, 21000, gasPrice, nil)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pkey)
	if err != nil {
		return err
	}

	log.Info("transaction hash", zap.String("hash", signedTx.Hash().Hex()))
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(context.Background(), client, signedTx)
	if err != nil {
		return err
	}
	log.Info("transaction mined",
		zap.String("hash", receipt.TxHash.Hex()),
		zap.String("block", receipt.BlockNumber.String()),
	)

	return nil
}
*/
