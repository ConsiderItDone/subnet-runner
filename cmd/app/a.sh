#!/bin/bash

# Run the register-remote command
go run app.go register-remote

# Fund the address
go run app.go fund-address --bc-id C --addr 0x976EA74026E726554dB657fA54763abd0C3a0aa9
go run app.go balance-native --bc-id C --addr 0x976EA74026E726554dB657fA54763abd0C3a0aa9

# Check balances
./check_balances cosmos1t36cnszflpzq6kvthpegafpqy9tv05pr9n7nga 0x976EA74026E726554dB657fA54763abd0C3a0aa9