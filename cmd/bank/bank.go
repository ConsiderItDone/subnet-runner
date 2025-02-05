package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"

	"subnet-runner/contracts/ics20/ics20bank"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"
)

var (
	FlagRPC = cli.StringFlag{
		Name:  "rpc",
		Value: "http://127.0.0.1:9650/ext/bc/GKvUzzyNDQHYTWLZT82tAmVyRbshcsvX87Ghv4AvHc2seVVSn/rpc",
	}
	FlagBank = cli.StringFlag{
		Name:  "bank",
		Value: "0x52C84043CD9c865236f11d9Fc9F56aa003c1f922",
	}

	FlagDenom = cli.StringFlag{
		Name:  "denom",
		Value: "transfer/channel-0/stake",
	}

	FlagHash = cli.StringFlag{
		Name: "hash",
	}

	FlagAddr = cli.StringFlag{
		Name: "addr",
	}
)

func bankApp(c *cli.Context) (*ics20bank.ICS20Bank, *ethclient.Client, error) {
	fmt.Printf("Chain: %s\n", c.String("rpc"))
	fmt.Printf("Bank: %s\n", c.String("bank"))
	client, err := ethclient.Dial(c.String("rpc"))
	if err != nil {
		return nil, nil, err
	}

	if !common.IsHexAddress(c.String("bank")) {
		return nil, nil, fmt.Errorf("bad bank address '%s'", c.String("bank"))
	}

	bank, err := ics20bank.NewICS20Bank(common.HexToAddress(c.String("bank")), client)
	if err != nil {
		return nil, nil, err
	}

	return bank, client, nil
}

func balanceOf(c *cli.Context) error {
	bank, eth, err := bankApp(c)
	if err != nil {
		return err
	}

	address := common.HexToAddress(c.String("addr"))

	balance, err := bank.BalanceOf(nil, address, c.String("denom"))
	if err != nil {
		return err
	}

	fmt.Printf("Bank balance: %s\n", balance)

	// Get the balance
	balance, err = eth.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}
	fmt.Printf("EVM balance of %s: %s wei\n", address.Hex(), balance.String())

	return nil
}

func showTransferLogs(c *cli.Context) error {
	bank, eth, err := bankApp(c)
	if err != nil {
		return err
	}

	receipt, err := eth.TransactionReceipt(context.Background(), common.HexToHash(c.String("hash")))
	if err != nil {
		return err
	}

	for i := range receipt.Logs {
		e, err := bank.ParseTransfer(*receipt.Logs[i])
		if err != nil {
			continue
		}

		fmt.Printf("Transfer [%d]\n", i)
		fmt.Printf("  From:  %s\n", e.From)
		fmt.Printf("  To:    %s\n", e.To)
		fmt.Printf("  Path:  %s\n", e.Path)
		fmt.Printf("  Value: %s\n", e.Value.String())
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "ics20bank"
	app.Usage = "ics20 banl cli app"
	app.Commands = []cli.Command{
		{
			Flags:  []cli.Flag{FlagRPC, FlagBank, FlagDenom, FlagAddr},
			Name:   "balance-of",
			Action: balanceOf,
		},
		{
			Flags:  []cli.Flag{FlagRPC, FlagBank, FlagHash},
			Name:   "show-transfer-logs",
			Action: showTransferLogs,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
