#!/bin/bash

# Define the configuration file paths
AWM_CONFIG_FILE="./awm-relayer-config.json"
# use full path to avoid issues with sed
RELAYER_CONFIG_FILE="/Users/vn/.relayer/config/config.yaml"

# Read values from awm-relayer-config.json
SUBNET_ID="256Y6fkZnXSaXykwqaH3PzBpbL7FYbQQDirpH7hayeVkHXCM69"

BLOCKCHAIN_ID=$(jq -r --arg SUBNET_ID "$SUBNET_ID" '.["source-blockchains"][] | select(.["subnet-id"] == $SUBNET_ID) | .["blockchain-id"]' $AWM_CONFIG_FILE)
BASE_URL=$(jq -r --arg SUBNET_ID "$SUBNET_ID" '.["source-blockchains"][] | select(.["subnet-id"] == $SUBNET_ID) | .["rpc-endpoint"]["base-url"]' $AWM_CONFIG_FILE)

# Update values in the relayer configuration file
sed -i '' "1,/rpc-addr: .*/s|rpc-addr: .*|rpc-addr: $BASE_URL|" $RELAYER_CONFIG_FILE
sed -i '' "1,/blockchain-id: .*/s|blockchain-id: .*|blockchain-id: $BLOCKCHAIN_ID|" $RELAYER_CONFIG_FILE

# Remove the specified lines from the config file
sed -i '' '/client-id: 07-tendermint-0/d' "$RELAYER_CONFIG_FILE"
sed -i '' '/connection-id: connection-0/d' "$RELAYER_CONFIG_FILE"
sed -i '' '/client-id: 14-avalanche-0/d' "$RELAYER_CONFIG_FILE"
sed -i '' '/connection-id: connection-0/d' "$RELAYER_CONFIG_FILE"

# Run the awm-relayer command with the specified configuration file
awm-relayer --config-file $AWM_CONFIG_FILE

