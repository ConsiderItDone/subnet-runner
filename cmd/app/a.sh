#!/bin/bash

# Source the addresses from the external file
source ./.addresses

# Run the register-remote command
go run app.go register-remote

# Fund the address
go run app.go fund-address --bc-id C --addr $DEFAULT_EVM_RECEIVER_ADDRESS
go run app.go balance-native --bc-id C --addr  $DEFAULT_EVM_RECEIVER_ADDRESS

# Check balances
./check_balances $DEFAULT_SENDER_COSMOS_ADDRESS $DEFAULT_EVM_RECEIVER_ADDRESS