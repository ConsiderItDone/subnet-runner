package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	pk1Flag = &cli.StringFlag{
		Name:  "pk1",
		Usage: "First private key",
	}
	pk2Flag = &cli.StringFlag{
		Name:  "pk2",
		Usage: "Second private key",
	}
	delayFlag = &cli.UintFlag{
		Name:  "delay",
		Usage: "Delay in seconds",
		Value: 5,
	}
	rpcFlag = &cli.StringFlag{
		Name:  "rpc",
		Usage: "RPC endpoint",
	}
)
var closedOnShutdownCh = make(chan struct{})
var log *zap.Logger

func main() {

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.CallerKey = "" // Disable caller key

	log, _ = config.Build()
	defer log.Sync()

	app := &cli.App{
		Name:      "txspam",
		Usage:     "A CLI tool for spamming transactions",
		UsageText: "txspam spam --pk1 a --pk2 b --delay 1 --rpc http://somerpc",
		Flags: []cli.Flag{
			pk1Flag,
			pk2Flag,
			delayFlag,
			rpcFlag,
		},
		Commands: []*cli.Command{
			{
				Name:  "spam",
				Usage: "Spam transactions with given parameters",
				Flags: []cli.Flag{
					pk1Flag,
					pk2Flag,
					delayFlag,
					rpcFlag,
				},
				Action: txSpam,
			},
		},
	}
	go handleShutdown()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("error running app", zap.Error(err))
	}
}

func handleShutdown() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGKILL)
	<-sigCh
	close(closedOnShutdownCh)
}

func txSpam(c *cli.Context) error {
	pk1 := c.String("pk1")
	pk2 := c.String("pk2")
	if pk1 == pk2 {
		return fmt.Errorf("private keys must be different")
	}
	// remove 0x prefix
	if pk1[:2] == "0x" {
		pk1 = pk1[2:]
	}

	if pk2[:2] == "0x" {
		pk2 = pk2[2:]
	}

	delay := c.Uint("delay")
	if delay < 1 {
		return fmt.Errorf("delay must be greater than 0")
	}

	rpc := c.String("rpc")

	log.Info("Starting spamming transactions")
	for {
		select {
		case <-closedOnShutdownCh:
			return nil
		case <-time.After(time.Duration(delay) * time.Second):
			if err := doTx(pk1, pk2, rpc); err != nil {
				log.Error("error doing tx", zap.Error(err))
				return err
			}
			// swap pk1 and pk2
			// pk1, pk2 = pk2, pk1
		}
	}
}

func doTx(pkFrom, pkTo, rpc string) error {
	pkey, err := crypto.HexToECDSA(pkFrom)
	if err != nil {
		return err
	}
	addr := crypto.PubkeyToAddress(pkey.PublicKey)

	client, err := ethclient.Dial(rpc)
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

	pkeyTo, err := crypto.HexToECDSA(pkTo)
	if err != nil {
		return err
	}
	toAddress := crypto.PubkeyToAddress(pkeyTo.PublicKey)

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    big.NewInt(1000000),
		Gas:      21000,
		GasPrice: gasPrice,
		Data:     nil,
	})
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pkey)
	if err != nil {
		return err
	}

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
