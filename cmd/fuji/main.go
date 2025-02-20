package main

import (
	"fmt"
	"os"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"

	"subnet-runner/internal"
)

var (
	ibcAddr = common.HexToAddress("0x0300000000000000000000000000000000000002")
)

func main() {
	fmt.Println("fuji")

	// Create the logger
	logFactory := logging.NewFactory(logging.Config{
		DisplayLevel: logging.Info,
		LogLevel:     logging.Info,
	})
	log, err := logFactory.Make("fuji")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = internal.DeploySubnetContracts(log, []string{"http://147.28.163.24:9650/ext/bc/X95M7DGW2beNdu9MqUBc7VhSb5H93JXCCaqxuqAJBrSDsXfka/rpc"}, ibcAddr)
	if err != nil {
		log.Fatal("fatal error", zap.Error(err))
	}
}
