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

	req := internal.DeployRequest{
		Url:                      "http://167.71.94.136:9650/ext/bc/2ZNTQEbhXfuvFqrZEas2YKVicANd3Q53aKnFzuVNTFkHtyxFzi/rpc",
		Pk:                       "deployment-pk-here",
		TpMessengerAddressSubnet: "0x253b2784c75e510dD0fF1da844684a1aC0aa5fcf",
		TpRegistryAddressSubnet:  "0x7c604e63F3Faa40be4A91fA984788753CD2A699B",
		TpMessengerAddressC:      "0x253b2784c75e510dD0fF1da844684a1aC0aa5fcf",
		TpRegistryAddressC:       "0xF86Cb19Ad8405AEFa7d09C778215D2Cb6eBfB228",
		TokenName:                "Landslide",
		TokenSymbol:              "transfer/channel-0/stake",
	}

	err = internal.DeploySubnetContracts(log, ibcAddr, req)
	if err != nil {
		log.Fatal("fatal error", zap.Error(err))
	}
}
