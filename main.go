package main

import (
	"context"
	_ "embed"
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
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

func await(nw network.Network, log logging.Logger, timeout time.Duration) error {
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
	if err := await(nw, log, healthyTimeout); err != nil {
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
		if _, err := copy(
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
	if err := await(nw, log, healthyTimeout); err != nil {
		return err
	}

	rpcUrls := make([]string, len(nodeNames))
	for i := range nodeNames {
		node, err := nw.GetNode(nodeNames[i])
		if err != nil {
			return err
		}
		rpcUrls[i] = fmt.Sprintf("http://127.0.0.1:%d/ext/bc/%s/rpc", node.GetAPIPort(), chains[0])
		log.Info("subnet rpc url", zap.String("node", nodeNames[i]), zap.String("url", rpcUrls[i]))
	}

	log.Info("Network will run until you CTRL + C to exit...")

	if err := doICS20(log, rpcUrls); err != nil {
		log.Error("can't deploy ICS20", zap.Error(err))
		return err
	}

	for {
		time.Sleep(5 * time.Second)
		select {
		case <-closedOnShutdownCh:
			return nil
		case <-time.After(5 * time.Second):
			if err := doTx(log, rpcUrls); err != nil {
				return err
			}
		}
	}

	return nil
}
