package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

const rpcUrl = "http://10.11.0.3:9650/ext/bc/2rNGJ7nn4scafjYR6wo5Vd15UhjvZg48hZ3DjdnFpqVH63WbVa/rpc"

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
	for {
		time.Sleep(5 * time.Second)
		select {
		case <-time.After(5 * time.Second):
			if err := doTx(log, rpcUrl); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

}

func doTx(log logging.Logger, url string) error {
	pkey, err := crypto.HexToECDSA("56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027")
	if err != nil {
		return err
	}
	addr := crypto.PubkeyToAddress(pkey.PublicKey)

	client, err := ethclient.Dial(url)
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
