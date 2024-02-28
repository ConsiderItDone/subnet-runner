package main

import (
	"context"
	_ "embed"
	"fmt"
	"go/build"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ava-labs/avalanche-network-runner/local"
	"github.com/ava-labs/avalanche-network-runner/network"
	"github.com/ava-labs/avalanchego/utils/logging"
	"go.uber.org/zap"
)

const (
	healthyTimeout = 2 * time.Minute
)

var (
	goPath = os.ExpandEnv("$GOPATH")

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

func run(log logging.Logger, binaryPath string, workDir string) error {
	// Create the network
	nwConfig, err := local.NewDefaultConfig(fmt.Sprintf("%s/avalanchego", binaryPath))
	if err != nil {
		return err
	}

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
		network.BlockchainSpec{
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
	// Wait until done shutting down network after SIGINT/SIGTERM
	<-closedOnShutdownCh
	return nil
}
