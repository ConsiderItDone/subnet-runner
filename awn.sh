export AVALANCHEGO_EXEC_PATH="/tmp/e2e-test/avalanchego/avalanchego"
export AVALANCHEGO_PLUGIN_PATH="/tmp/e2e-test/avalanchego/plugins"

avalanche-network-runner server

#avalanche-network-runner control start \
#--log-level debug \
#--endpoint="localhost:9650" \
#--blockchain-specs '[{"vm_name": "subnetevm", "genesis": "/tmp/e2e-test/genesis.json"}]'