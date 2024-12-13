package internal

import (
	"fmt"
	"os"
	"strings"
)

const (
	awmRelayerConfig = "./awm-relayer-config.json"
	awmRelayerTmpl   = "./awm-relayer-config.tmpl"
)

// UpdateRelayerConfig updates the relayer config file with the provided values.
func UpdateRelayerConfig(
	cChainBlockchainID, lndBlockchainID, lndSubnetID,
	teleporterRegistryAddress, teleporterMessengerAddress,
	accountPrivateKey string,
) error {
	data, err := os.ReadFile(awmRelayerTmpl)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	template := string(data)
	template = strings.ReplaceAll(template, "{{.LndSubnetID}}", lndSubnetID)
	template = strings.ReplaceAll(template, "{{.LndBlockchainID}}", lndBlockchainID)
	template = strings.ReplaceAll(template, "{{.CChainBlockchainID}}", cChainBlockchainID)
	template = strings.ReplaceAll(template, "{{.TeleporterRegistryAddress}}", teleporterRegistryAddress)
	template = strings.ReplaceAll(template, "{{.TeleporterMessengerAddress}}", teleporterMessengerAddress)
	template = strings.ReplaceAll(template, "{{.AccountPrivateKey}}", accountPrivateKey)

	if err := os.WriteFile(awmRelayerConfig, []byte(template), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
