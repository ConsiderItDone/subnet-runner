// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20tokenhome

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FromCosmosSendTokensInput is an auto generated low-level Go binding around an user-defined struct.
type FromCosmosSendTokensInput struct {
	MessageID                          [32]byte
	DestinationBlockchainID            [32]byte
	DestinationTokenTransferrerAddress common.Address
	SourceCosmosBlockchainID           [32]byte
	SourceCosmosAddress                []byte
	Recipient                          common.Address
	PrimaryFeeTokenAddress             common.Address
	PrimaryFee                         *big.Int
	SecondaryFee                       *big.Int
	RequiredGasLimit                   *big.Int
	MultiHopFallback                   common.Address
}

// RemoteTokenTransferrerSettings is an auto generated low-level Go binding around an user-defined struct.
type RemoteTokenTransferrerSettings struct {
	Registered       bool
	CollateralNeeded *big.Int
	TokenMultiplier  *big.Int
	MultiplyOnRemote bool
}

// SendAndCallInput is an auto generated low-level Go binding around an user-defined struct.
type SendAndCallInput struct {
	DestinationBlockchainID            [32]byte
	DestinationTokenTransferrerAddress common.Address
	RecipientContract                  common.Address
	RecipientPayload                   []byte
	RequiredGasLimit                   *big.Int
	RecipientGasLimit                  *big.Int
	MultiHopFallback                   common.Address
	FallbackRecipient                  common.Address
	PrimaryFeeTokenAddress             common.Address
	PrimaryFee                         *big.Int
	SecondaryFee                       *big.Int
}

// SendTokensInput is an auto generated low-level Go binding around an user-defined struct.
type SendTokensInput struct {
	DestinationBlockchainID            [32]byte
	DestinationTokenTransferrerAddress common.Address
	Recipient                          common.Address
	PrimaryFeeTokenAddress             common.Address
	PrimaryFee                         *big.Int
	SecondaryFee                       *big.Int
	RequiredGasLimit                   *big.Int
	MultiHopFallback                   common.Address
}

// ToCosmosSendTokensInput is an auto generated low-level Go binding around an user-defined struct.
type ToCosmosSendTokensInput struct {
	DestinationBlockchainID            [32]byte
	DestinationTokenTransferrerAddress common.Address
	DestinationCosmosRecipient         []byte
	DestinationCosmosBlockchainID      [32]byte
	PrimaryFeeTokenAddress             common.Address
	PrimaryFee                         *big.Int
	SecondaryFee                       *big.Int
	RequiredGasLimit                   *big.Int
	MultiHopFallback                   common.Address
}

// ERC20TokenHomeMetaData contains all meta data concerning the ERC20TokenHome contract.
var ERC20TokenHomeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTeleporterVersion\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceCosmosBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sourceCosmosAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structFromCosmosSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FromCosmosTokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialCollateralNeeded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"RemoteRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"destinationCosmosRecipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCosmosBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structToCosmosSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ToCosmosTokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensCosmosWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC20_TOKEN_HOME_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TELEPORTER_REGISTRY_APP_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HOME_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"}],\"name\":\"getRemoteTokenTransferrerSettings\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralNeeded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"multiplyOnRemote\",\"type\":\"bool\"}],\"internalType\":\"structRemoteTokenTransferrerSettings\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"}],\"name\":\"getTransferredBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTeleporterVersion\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"destinationCosmosRecipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCosmosBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"internalType\":\"structToCosmosSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"toCosmosSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051615d04380380615d0483398101604081905261002f916108a9565b61003c8585858585610049565b505050505061092a565b50565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff1615906001600160401b03166000811580156100935750825b90506000826001600160401b031660011480156100af5750303b155b9050811580156100bd575080155b156100db5760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b0319166001178555831561010957845460ff60401b1916680100000000000000001785555b6101168a8a8a8a8a610168565b831561015c57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b61017061018d565b61017d85858585856101dd565b61018682610202565b5050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff166101db57604051631afcd79f60e31b815260040160405180910390fd5b565b6101e561018d565b6101f085858561024b565b6101f861026b565b610186828261027b565b61020a61018d565b7f914a9547f6c3ddce1d5efbd9e687708f0d1d408ce129e8e1a88bce4f40e2950080546001600160a01b0319166001600160a01b0392909216919091179055565b61025361018d565b61025d838261040a565b61026682610430565b505050565b61027361018d565b6101db610441565b61028361018d565b6001600160a01b0382166102de5760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f20746f6b656e206164647265737300000060448201526064015b60405180910390fd5b60128160ff16111561033d5760405162461bcd60e51b815260206004820152602260248201527f546f6b656e486f6d653a20746f6b656e20646563696d616c7320746f6f2068696044820152610ced60f31b60648201526084016102d5565b60007f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e60090507302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d89190610911565b8155600101805460ff909216600160a01b026001600160a81b03199092166001600160a01b0390931692909217179055565b61041261018d565b61041a61046c565b61042261047c565b61042c8282610484565b5050565b61043861018d565b61004681610614565b60007fd2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c75005b6001905550565b61047461018d565b6101db61064f565b6101db61018d565b61048c61018d565b6001600160a01b0382166105085760405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084016102d5565b60007fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d00905060008390506000816001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610572573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105969190610911565b116105ec5760405162461bcd60e51b81526020600482015260326024820152600080516020615ce4833981519152604482015271656c65706f7274657220726567697374727960701b60648201526084016102d5565b81546001600160a01b0319166001600160a01b03821617825561060e8361067f565b50505050565b61061c61018d565b6001600160a01b03811661064657604051631e4fbdf760e01b8152600060048201526024016102d5565b6100468161081c565b61065761018d565b60007f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00610465565b7fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d0080546040805163301fd1f560e21b815290516000926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa1580156106e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070d9190610911565b60028301549091508184111561076d5760405162461bcd60e51b81526020600482015260316024820152600080516020615ce483398151915260448201527032b632b837b93a32b9103b32b939b4b7b760791b60648201526084016102d5565b8084116107e25760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e0060648201526084016102d5565b60028301849055604051849082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a350505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b80516001600160a01b03811681146108a457600080fd5b919050565b600080600080600060a086880312156108c157600080fd5b6108ca8661088d565b94506108d86020870161088d565b9350604086015192506108ed6060870161088d565b9150608086015160ff8116811461090357600080fd5b809150509295509295909350565b60006020828403121561092357600080fd5b5051919050565b6153ab806109396000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c806362e3901b116100b8578063973142971161007c57806397314297146102b9578063c8511ada146102dc578063c868efaa146103b2578063d2cc7a70146103c5578063f2fde38b146103ec578063fd658268146103ff57600080fd5b806362e3901b146102445780636569003814610259578063715018a61461026c5780638da5cb5b14610274578063909a6ac0146102a457600080fd5b80634213cf78116100ff5780634213cf78146101e15780634511243e146101f65780634797735f146102095780635d16225d1461021e5780635eb995141461023157600080fd5b806310fe9ae81461013c578063154d625a146101855780632b0d8f18146101a65780633bb03890146101bb5780633ea51daa146101ce575b600080fd5b7f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e601546001600160a01b03165b6040516001600160a01b0390911681526020015b60405180910390f35b610198610193366004613f93565b610412565b60405190815260200161017c565b6101b96101b4366004613fc3565b61045b565b005b6101b96101c9366004613fef565b61055f565b6101b96101dc366004614057565b610675565b6000805160206152d683398151915254610198565b6101b9610204366004613fc3565b61068b565b61019860008051602061533683398151915281565b6101b961022c3660046140a1565b61077c565b6101b961023f3660046140d3565b610794565b6101986000805160206152d683398151915281565b6101b96102673660046140ec565b6107a8565b6101b96107d2565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610168565b61019860008051602061535683398151915281565b6102cc6102c7366004613fc3565b6107e6565b604051901515815260200161017c565b61037b6102ea366004613f93565b60408051608080820183526000808352602080840182905283850182905260609384018290529581526000805160206152f683398151915286528381206001600160a01b039590951681529385529282902082519384018352805460ff9081161515855260018201549585019590955260028101549284019290925260039091015490921615159181019190915290565b60405161017c9190815115158152602080830151908201526040808301519082015260609182015115159181019190915260800190565b6101b96103c0366004614128565b610808565b7fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d0254610198565b6101b96103fa366004613fc3565b6109ca565b6101b961040d3660046141b0565b610a05565b60008281527f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e603602090815260408083206001600160a01b03851684529091529020545b92915050565b600080516020615356833981519152610472610a15565b6001600160a01b0382166104a15760405162461bcd60e51b8152600401610498906141e8565b60405180910390fd5b6104ab8183610a1d565b1561050e5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610498565b6001600160a01b038216600081815260018381016020526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a25050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156105a45750825b90506000826001600160401b031660011480156105c05750303b155b9050811580156105ce575080155b156105ec5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561061657845460ff60401b1916600160401b1785555b6106238a8a8a8a8a610a3f565b831561066957845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b610687610681836143ac565b82610a64565b5050565b6000805160206153568339815191526106a2610a15565b6001600160a01b0382166106c85760405162461bcd60e51b8152600401610498906141e8565b6106d28183610a1d565b6107305760405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472794170703a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610498565b6001600160a01b0382166000818152600183016020526040808220805460ff19169055517f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c39190a25050565b61068761078e36849003840184614458565b82610cfe565b61079c610a15565b6107a581610eaa565b50565b6106876107c16000805160206152d68339815191525490565b30336107cc866144e4565b85611047565b6107da61125a565b6107e460006112b5565b565b60006000805160206153568339815191526108018184610a1d565b9392505050565b610810611326565b600060008051602061535683398151915260028101548154919250906001600160a01b0316634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa15801561087f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108a391906145b5565b101561090a5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610498565b6109148133610a1d565b1561097a5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610498565b6109bb858585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061137092505050565b506109c46117d5565b50505050565b6109d261125a565b6001600160a01b0381166109fc57604051631e4fbdf760e01b815260006004820152602401610498565b6107a5816112b5565b610a10838383611800565b505050565b6107e461125a565b6001600160a01b03166000908152600191909101602052604090205460ff1690565b610a476119f8565b610a548585858585611a41565b610a5d82611a66565b5050505050565b6000805160206153168339815191528054600114610a945760405162461bcd60e51b8152600401610498906145ce565b60028155610aa183611a9d565b6101008301516001600160a01b031615610acd5760405162461bcd60e51b815260040161049890614612565b600080610aed856000015186602001518688608001518960a00151611b08565b915091506000604051806040016040528060056007811115610b1157610b11614658565b81526020016040518060800160405280610b29611c5a565b8a516040516337c83a2f60e21b81526001600160a01b03929092169163df20e8bc91610b5b9160040190815260200190565b602060405180830381865afa158015610b78573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9c91906145b5565b8152602001896040015181526020018960000151815260200186815250604051602001610bc991906146be565b60405160208183030381529060405281525090506000610cab6040518060c001604052808960000151815260200189602001516001600160a01b0316815260200160405180604001604052808b608001516001600160a01b031681526020018781525081526020018960e00151815260200160006001600160401b03811115610c5457610c54614236565b604051908082528060200260200182016040528015610c7d578160200160208202803683370190505b50815260200184604051602001610c949190614704565b604051602081830303815290604052815250611d4e565b9050336001600160a01b0316817f79bd37563e6b913ff0a0e0c938210c007ebef4c942d38333576d03cb04620e9b8987604051610ce9929190614749565b60405180910390a35050600190925550505050565b6000805160206153168339815191528054600114610d2e5760405162461bcd60e51b8152600401610498906145ce565b60028155610d3b83611e6d565b60e08301516001600160a01b031615610d665760405162461bcd60e51b815260040161049890614612565b600080610d86856000015186602001518688606001518960800151611edd565b915091506000604051806040016040528060016007811115610daa57610daa614658565b8152602001604051806040016040528089604001516001600160a01b0316815260200186815250604051602001610de19190614803565b60405160208183030381529060405281525090506000610e6c6040518060c001604052808960000151815260200189602001516001600160a01b0316815260200160405180604001604052808b606001516001600160a01b031681526020018781525081526020018960c00151815260200160006001600160401b03811115610c5457610c54614236565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb528987604051610ce9929190614823565b60008051602061535683398151915280546040805163301fd1f560e21b815290516000926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa158015610f02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f2691906145b5565b600283015490915081841115610f985760405162461bcd60e51b815260206004820152603160248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610498565b80841161100d5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610498565b60028301849055604051849082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a350505050565b60008051602061531683398151915280546001146110775760405162461bcd60e51b8152600401610498906145ce565b6002815561108483612067565b60c08301516001600160a01b0316156110af5760405162461bcd60e51b815260040161049890614612565b6000806110d18560000151866020015186886101000151896101200151611edd565b9150915060006040518060400160405280600260078111156110f5576110f5614658565b81526020016040518061010001604052808c81526020018b6001600160a01b031681526020018a6001600160a01b0316815260200189604001516001600160a01b03168152602001868152602001896060015181526020018960a0015181526020018960e001516001600160a01b031681525060405160200161117891906148a5565b604051602081830303815290604052815250905060006112046040518060c001604052808960000151815260200189602001516001600160a01b0316815260200160405180604001604052808b61010001516001600160a01b031681526020018781525081526020018960800151815260200160006001600160401b03811115610c5457610c54614236565b9050876001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168987604051611242929190614944565b60405180910390a35050600190925550505050505050565b3361128c7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146107e45760405163118cdaa760e01b8152336004820152602401610498565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080546001190161136a57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60006000805160206152d683398151915290506000828060200190518101906113999190614a6c565b90506001815160078111156113b0576113b0614658565b036113fb57600081602001518060200190518101906113cf9190614afa565b905060006113e28787846020015161224e565b90506113f28260000151826122dc565b50505050505050565b60068151600781111561141057611410614658565b03611459576000816020015180602001905181019061142f9190614b34565b9050600061144287878460800151612340565b90506113f28787846060015184866000015161243e565b60028151600781111561146e5761146e614658565b03611589576000816020015180602001905181019061148d9190614bd7565b905060006114a08787846080015161224e565b825190915087146115065760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e486f6d653a206d69736d61746368656420736f7572636520626c6f60448201526918dad8da185a5b88125160b21b6064820152608401610498565b856001600160a01b031682602001516001600160a01b03161461157f5760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e486f6d653a206d69736d617463686564206f726967696e2073656e60448201526a646572206164647265737360a81b6064820152608401610498565b6113f28282612617565b60038151600781111561159e5761159e614658565b0361167757600081602001518060200190518101906115bd9190614ca6565b90506000806115d68888856060015186608001516127ff565b9150915061166d6040518061010001604052808560000151815260200185602001516001600160a01b0316815260200185604001516001600160a01b031681526020018760010160009054906101000a90046001600160a01b03166001600160a01b03168152602001838152602001600081526020018560a0015181526020018560c001516001600160a01b0316815250836128ae565b5050505050505050565b60048151600781111561168c5761168c614658565b0361178c57600081602001518060200190518101906116ab9190614d41565b90506000806116c5888885608001518661014001516127ff565b9150915061166d888885600001516040518061016001604052808860200151815260200188604001516001600160a01b0316815260200188606001516001600160a01b031681526020018860a00151815260200188610100015181526020018860c0015181526020018861012001516001600160a01b031681526020018860e001516001600160a01b031681526020018a60010160009054906101000a90046001600160a01b03166001600160a01b03168152602001868152602001600081525086612a41565b6000815160078111156117a1576117a1614658565b03610a5d57600081602001518060200190518101906117c09190614e3e565b90506117cd868683612c26565b505050505050565b60007f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005b6001905550565b60008051602061531683398151915280546001146118305760405162461bcd60e51b8152600401610498906145ce565b600280825560008581526000805160206152f6833981519152602090815260408083206001600160a01b03881684528252918290208251608081018452815460ff9081161515808352600184015494830194909452948201549381019390935260030154909216151560608201526000805160206152d6833981519152916118ca5760405162461bcd60e51b815260040161049890614ea6565b60008160200151116119285760405162461bcd60e51b815260206004820152602160248201527f546f6b656e486f6d653a207a65726f20636f6c6c61746572616c206e656564656044820152601960fa1b6064820152608401610498565b6119318461303a565b935060008082602001518610611962576020830151600092506119549087614ef1565b905082602001519550611975565b8583602001516119729190614ef1565b91505b600088815260028501602090815260408083206001600160a01b038b168085529083529281902060010185905580518981529182018590528a917f6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6910160405180910390a380156119ea576119ea33826122dc565b505060019092555050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166107e457604051631afcd79f60e31b815260040160405180910390fd5b611a496119f8565b611a54858585613063565b611a5c61307e565b610a5d828261308e565b611a6e6119f8565b60008051602061533683398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b600081604001515111611ac25760405162461bcd60e51b815260040161049890614f04565b60008160e0015111611ae65760405162461bcd60e51b815260040161049890614f45565b60c0810151156107a55760405162461bcd60e51b815260040161049890614f87565b60008581526000805160206152f6833981519152602090815260408083206001600160a01b038816845282528083208151608081018352815460ff90811615158083526001840154958301959095526002830154938201939093526003909101549091161515606082015282916000805160206152d68339815191529190611ba25760405162461bcd60e51b815260040161049890614ea6565b602081015115611bc45760405162461bcd60e51b815260040161049890614fc8565b611bcd876131f9565b96508415611be457611be186335b87613279565b94505b6000611bf9826040015183606001518a6133d8565b905060008111611c4b5760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f207363616c656420616d6f756e740000006044820152606401610498565b99949850939650505050505050565b60008051602061535683398151915280546040805163d820e64f60e01b815290516000939284926001600160a01b039091169163d820e64f916004808201926020929091908290030181865afa158015611cb8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cdc919061500f565b9050611ce88282610a1d565b156104555760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610498565b600080611d59611c5a565b60408401516020015190915015611dfe576040830151516001600160a01b0316611ddb5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610498565b604083015160208101519051611dfe916001600160a01b039091169083906133ef565b604051630624488560e41b81526001600160a01b03821690636244885090611e2a90869060040161502c565b6020604051808303816000875af1158015611e49573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080191906145b5565b60408101516001600160a01b0316611e975760405162461bcd60e51b815260040161049890614f04565b60008160c0015111611ebb5760405162461bcd60e51b815260040161049890614f45565b60a0810151156107a55760405162461bcd60e51b815260040161049890614f87565b60008581526000805160206152f6833981519152602090815260408083206001600160a01b038816845282528083208151608081018352815460ff90811615158083526001840154958301959095526002830154938201939093526003909101549091161515606082015282916000805160206152d68339815191529190611f775760405162461bcd60e51b815260040161049890614ea6565b602081015115611f995760405162461bcd60e51b815260040161049890614fc8565b611fa28761303a565b96508415611fb757611fb48633611bdb565b94505b6000611fcc826040015183606001518a6133d8565b90506000811161201e5760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f207363616c656420616d6f756e740000006044820152606401610498565b60008a815260038401602090815260408083206001600160a01b038d168452909152812080548392906120529084906150e5565b90915550909a95995094975050505050505050565b60408101516001600160a01b03166120d45760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e486f6d653a207a65726f20726563697069656e7420636f6e7472616044820152696374206164647265737360b01b6064820152608401610498565b60008160800151116120f85760405162461bcd60e51b815260040161049890614f45565b60008160a00151116121585760405162461bcd60e51b815260206004820152602360248201527f546f6b656e486f6d653a207a65726f20726563697069656e7420676173206c696044820152621b5a5d60ea1b6064820152608401610498565b80608001518160a00151106121be5760405162461bcd60e51b815260206004820152602660248201527f546f6b656e486f6d653a20696e76616c696420726563697069656e7420676173604482015265081b1a5b5a5d60d21b6064820152608401610498565b60e08101516001600160a01b031661222b5760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e486f6d653a207a65726f2066616c6c6261636b20726563697069656044820152696e74206164647265737360b01b6064820152608401610498565b610140810151156107a55760405162461bcd60e51b815260040161049890614f87565b60008381526000805160206152f6833981519152602090815260408083206001600160a01b038616845282528083208151608081018352815460ff9081161515825260018301549482019490945260028201549281019290925260030154909116151560608201526000805160206152d6833981519152906122d281878787613479565b9695505050505050565b604051818152600080516020615336833981519152906001600160a01b038416907f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b9060200160405180910390a28054610a10906001600160a01b03168484613573565b60008381526000805160206152f6833981519152602090815260408083206001600160a01b038616845282528083208151608081018352815460ff9081161515808352600184015495830195909552600283015493820193909352600390910154909116151560608201526000805160206152d6833981519152916123d75760405162461bcd60e51b815260040161049890614ea6565b60006123ec82604001518360600151876135d2565b9050600081116122d25760405162461bcd60e51b815260206004820152601c60248201527f546f6b656e486f6d653a207a65726f20746f6b656e20616d6f756e74000000006044820152606401610498565b600060008051602061533683398151915280546040516370a0823160e01b81526001600160a01b038781166004830152929350600092909116906370a0823190602401602060405180830381865afa15801561249e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124c291906145b5565b82546040516340c10f1960e01b81526001600160a01b038881166004830152602482018890529293509116906340c10f19906044016020604051808303816000875af1925050508015612532575060408051601f3d908101601f1916820190925261252f918101906150f8565b60015b612549576125448787878760006135e1565b6113f2565b5081546040516370a0823160e01b81526001600160a01b03878116600483015260009216906370a0823190602401602060405180830381865afa158015612594573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125b891906145b5565b60408051868152602081018890529192506001600160a01b038816917ffc043975d4effd7520860981e63ea2ba7197dbe72a39dd32a3e227e820ca7ffe910160405180910390a261166d888888886126108787614ef1565b8a146135e1565b600080516020615336833981519152805460608401516001600160a01b0390911690612645908290856133ef565b600084600001518560200151866040015184878960a001516040516024016126729695949392919061511a565b60408051601f198184030181529190526020810180516001600160e01b03166394395edd60e01b17905260c086015160608701519192506000916126b79190846136c4565b6060870151604051636eb1769f60e11b81523060048201526001600160a01b0391821660248201529192506000919085169063dd62ed3e90604401602060405180830381865afa15801561270f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061273391906145b5565b905061274584886060015160006136d3565b81156127975786606001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff48760405161278a91815260200190565b60405180910390a26127df565b86606001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0876040516127d691815260200190565b60405180910390a25b80156113f25760e08701516113f2906001600160a01b0386169083613573565b60008481526000805160206152f6833981519152602090815260408083206001600160a01b038716845282528083208151608081018352815460ff90811615158252600183015494820194909452600282015492810192909252600301549091161515606082015281906000805160206152d68339815191529082612886828a8a8a613479565b9050600061289d83604001518460600151896135d2565b919a91995090975050505050505050565b60008051602061531683398151915280546001146128de5760405162461bcd60e51b8152600401610498906145ce565b600281556128eb83611e6d565b600061290584600001518560200151858760800151613763565b9050806000036129235761291d8460e00151846122dc565b50612a39565b604080518082019091526000908060018152602001604051806040016040528088604001516001600160a01b031681526020018581525060405160200161296a9190614803565b604051602081830303815290604052815250905060006129f96040518060c001604052808860000151815260200188602001516001600160a01b0316815260200160405180604001604052808a606001516001600160a01b031681526020018a6080015181525081526020018860c00151815260200160006001600160401b03811115610c5457610c54614236565b9050807f825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf08785604051612a2d929190614823565b60405180910390a25050505b600190555050565b6000805160206153168339815191528054600114612a715760405162461bcd60e51b8152600401610498906145ce565b60028155612a7e83612067565b6000612a998460000151856020015185876101200151613763565b905080600003612ab757612ab18460c00151846122dc565b50612c1b565b6040805180820190915260009080600281526020016040518061010001604052808b81526020018a6001600160a01b03168152602001896001600160a01b0316815260200188604001516001600160a01b03168152602001858152602001886060015181526020018860a0015181526020018860e001516001600160a01b0316815250604051602001612b4a91906148a5565b60405160208183030381529060405281525090506000612bdb6040518060c001604052808860000151815260200188602001516001600160a01b0316815260200160405180604001604052808a61010001516001600160a01b031681526020018a610120015181525081526020018860800151815260200160006001600160401b03811115610c5457610c54614236565b9050807f42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb308785604051612c0f929190614944565b60405180910390a25050505b600190555050505050565b6000805160206152d683398151915283612c8e5760405162461bcd60e51b8152602060048201526024808201527f546f6b656e486f6d653a207a65726f2072656d6f746520626c6f636b636861696044820152631b88125160e21b6064820152608401610498565b80548403612cf65760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e486f6d653a2063616e6e6f742072656769737465722072656d6f7460448201526e329037b71039b0b6b29031b430b4b760891b6064820152608401610498565b6001600160a01b038316612d655760405162461bcd60e51b815260206004820152603060248201527f546f6b656e486f6d653a207a65726f2072656d6f746520746f6b656e2074726160448201526f6e73666572726572206164647265737360801b6064820152608401610498565b600084815260028201602090815260408083206001600160a01b038716845290915290205460ff1615612de65760405162461bcd60e51b8152602060048201526024808201527f546f6b656e486f6d653a2072656d6f746520616c726561647920726567697374604482015263195c995960e21b6064820152608401610498565b6012826040015160ff161115612e505760405162461bcd60e51b815260206004820152602960248201527f546f6b656e486f6d653a2072656d6f746520746f6b656e20646563696d616c73604482015268040e8dede40d0d2ced60bb1b6064820152608401610498565b6001810154602083015160ff908116600160a01b9092041614612ec45760405162461bcd60e51b815260206004820152602660248201527f546f6b656e486f6d653a20696e76616c696420686f6d6520746f6b656e20646560448201526563696d616c7360d01b6064820152608401610498565b600080612ee68360010160149054906101000a900460ff1685604001516138dd565b915091506000612efb838387600001516135d2565b9050818015612f1557508451612f12908490615171565b15155b15612f2857612f256001826150e5565b90505b60405180608001604052806001151581526020018281526020018481526020018315158152508460020160008981526020019081526020016000206000886001600160a01b03166001600160a01b0316815260200190815260200160002060008201518160000160006101000a81548160ff021916908315150217905550602082015181600101556040820151816002015560608201518160030160006101000a81548160ff021916908315150217905550905050856001600160a01b0316877ff229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b9083886040015160405161302992919091825260ff16602082015260400190565b60405180910390a350505050505050565b600080516020615336833981519152805460009190610801906001600160a01b03163385613279565b61306b6119f8565b6130758382613926565b610a1082613948565b6130866119f8565b6107e4613959565b6130966119f8565b6001600160a01b0382166130ec5760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f20746f6b656e20616464726573730000006044820152606401610498565b60128160ff16111561314b5760405162461bcd60e51b815260206004820152602260248201527f546f6b656e486f6d653a20746f6b656e20646563696d616c7320746f6f2068696044820152610ced60f31b6064820152608401610498565b60006000805160206152d683398151915290506005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa1580156131a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131c791906145b5565b8155600101805460ff909216600160a01b026001600160a81b03199092166001600160a01b0390931692909217179055565b60008051602061533683398151915280546040805163079cc67960e41b8152336004820152602481018590529051600093926001600160a01b0316916379cc6790916044808301928792919082900301818387803b15801561325a57600080fd5b505af115801561326e573d6000803e3d6000fd5b509495945050505050565b6040516370a0823160e01b815230600482015260009081906001600160a01b038616906370a0823190602401602060405180830381865afa1580156132c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132e691906145b5565b90506132fd6001600160a01b03861685308661396f565b6040516370a0823160e01b81523060048201526000906001600160a01b038716906370a0823190602401602060405180830381865afa158015613344573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061336891906145b5565b90508181116133ce5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610498565b6122d28282614ef1565b60006133e784848460016139a8565b949350505050565b604051636eb1769f60e11b81523060048201526001600160a01b0383811660248301526000919085169063dd62ed3e90604401602060405180830381865afa15801561343f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061346391906145b5565b90506109c4848461347485856150e5565b6136d3565b835160009061349a5760405162461bcd60e51b815260040161049890614ea6565b6020850151156134f85760405162461bcd60e51b8152602060048201526024808201527f546f6b656e486f6d653a2072656d6f7465206e6f7420636f6c6c61746572616c6044820152631a5e995960e21b6064820152608401610498565b6135038484846139d0565b600061351886604001518760600151856135d2565b90506000811161356a5760405162461bcd60e51b815260206004820152601c60248201527f546f6b656e486f6d653a207a65726f20746f6b656e20616d6f756e74000000006044820152606401610498565b95945050505050565b6040516001600160a01b03838116602483015260448201839052610a1091859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613ac0565b60006133e784848460006139a8565b61362c6040805160c081018252600080825260208083018290528351808501855282815290810191909152909182019081526020016000815260200160608152602001606081525090565b8581526001600160a01b03858116602080840191909152604080518082018252600781528151606080820184529489168082528185018981528815159285019283528451808701929092525181850152905115158186015282518082039095018552608001825280830193909352516136a6929101614704565b60408051601f1981840301815291905260a08201526113f281611d4e565b60006133e78460008585613b23565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b1790526137248482613bf8565b6109c4576040516001600160a01b0384811660248301526000604483015261375991869182169063095ea7b3906064016135a0565b6109c48482613ac0565b60008481526000805160206152f6833981519152602090815260408083206001600160a01b038716845282528083208151608081018352815460ff9081161580158352600184015495830195909552600283015493820193909352600390910154909116151560608201526000805160206152d683398151915291806137ed575060008160200151115b156137fd576000925050506133e7565b8385116138615760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e486f6d653a20696e73756666696369656e7420616d6f756e74207460448201526b6f20636f766572206665657360a01b6064820152608401610498565b61386b8486614ef1565b9450600061388282604001518360600151886133d8565b90508060000361389857600093505050506133e7565b600088815260038401602090815260408083206001600160a01b038b168452909152812080548392906138cc9084906150e5565b909155509098975050505050505050565b60008060ff808516908416118181613901576138f98587615185565b60ff1661390f565b61390b8686615185565b60ff165b61391a90600a615282565b96919550909350505050565b61392e6119f8565b613936613c9b565b61393e613cab565b6106878282613cb3565b6139506119f8565b6107a581613e3d565b60006000805160206153168339815191526117f9565b6040516001600160a01b0384811660248301528381166044830152606482018390526109c49186918216906323b872dd906084016135a0565b6000811515841515036139c6576139bf858461528e565b90506133e7565b61356a85846152a5565b60008381527f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e603602090815260408083206001600160a01b03861684529091529020546000805160206152d68339815191529082811015613a895760405162461bcd60e51b815260206004820152602e60248201527f546f6b656e486f6d653a20696e73756666696369656e7420746f6b656e20747260448201526d616e736665722062616c616e636560901b6064820152608401610498565b613a938382614ef1565b60009586526003909201602090815260408087206001600160a01b03909616875294905250919092205550565b6000613ad56001600160a01b03841683613e45565b90508051600014158015613afa575080806020019051810190613af891906150f8565b155b15610a1057604051635274afe760e01b81526001600160a01b0384166004820152602401610498565b6000845a1015613b755760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e742067617300000000006044820152606401610498565b83471015613bc55760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c75650000006044820152606401610498565b826001600160a01b03163b600003613bdf575060006133e7565b600080600084516020860188888bf19695505050505050565b6000806000846001600160a01b031684604051613c1591906152b9565b6000604051808303816000865af19150503d8060008114613c52576040519150601f19603f3d011682016040523d82523d6000602084013e613c57565b606091505b5091509150818015613c81575080511580613c81575080806020019051810190613c8191906150f8565b801561356a5750505050506001600160a01b03163b151590565b613ca36119f8565b6107e4613e53565b6107e46119f8565b613cbb6119f8565b6001600160a01b038216613d375760405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560448201527f706f7274657220726567697374727920616464726573730000000000000000006064820152608401610498565b6000600080516020615356833981519152905060008390506000816001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015613d8f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613db391906145b5565b11613e1b5760405162461bcd60e51b815260206004820152603260248201527f54656c65706f7274657252656769737472794170703a20696e76616c69642054604482015271656c65706f7274657220726567697374727960701b6064820152608401610498565b81546001600160a01b0319166001600160a01b0382161782556109c483610eaa565b6109d26119f8565b606061080183836000613e5b565b6117d56119f8565b606081471015613e805760405163cd78605960e01b8152306004820152602401610498565b600080856001600160a01b03168486604051613e9c91906152b9565b60006040518083038185875af1925050503d8060008114613ed9576040519150601f19603f3d011682016040523d82523d6000602084013e613ede565b606091505b50915091506122d2868383606082613efe57613ef982613f45565b610801565b8151158015613f1557506001600160a01b0384163b155b15613f3e57604051639996b31560e01b81526001600160a01b0385166004820152602401610498565b5080610801565b805115613f555780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b6001600160a01b03811681146107a557600080fd5b8035613f8e81613f6e565b919050565b60008060408385031215613fa657600080fd5b823591506020830135613fb881613f6e565b809150509250929050565b600060208284031215613fd557600080fd5b813561080181613f6e565b60ff811681146107a557600080fd5b600080600080600060a0868803121561400757600080fd5b853561401281613f6e565b9450602086013561402281613f6e565b935060408601359250606086013561403981613f6e565b9150608086013561404981613fe0565b809150509295509295909350565b6000806040838503121561406a57600080fd5b82356001600160401b0381111561408057600080fd5b8301610120818603121561409357600080fd5b946020939093013593505050565b6000808284036101208112156140b657600080fd5b610100808212156140c657600080fd5b9395938601359450505050565b6000602082840312156140e557600080fd5b5035919050565b600080604083850312156140ff57600080fd5b82356001600160401b0381111561411557600080fd5b8301610160818603121561409357600080fd5b6000806000806060858703121561413e57600080fd5b84359350602085013561415081613f6e565b925060408501356001600160401b038082111561416c57600080fd5b818701915087601f83011261418057600080fd5b81358181111561418f57600080fd5b8860208285010111156141a157600080fd5b95989497505060200194505050565b6000806000606084860312156141c557600080fd5b8335925060208401356141d781613f6e565b929592945050506040919091013590565b6020808252602e908201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b60405161012081016001600160401b038111828210171561426f5761426f614236565b60405290565b60405161010081016001600160401b038111828210171561426f5761426f614236565b60405161016081016001600160401b038111828210171561426f5761426f614236565b604080519081016001600160401b038111828210171561426f5761426f614236565b60405160a081016001600160401b038111828210171561426f5761426f614236565b604051601f8201601f191681016001600160401b038111828210171561432757614327614236565b604052919050565b60006001600160401b0382111561434857614348614236565b50601f01601f191660200190565b600082601f83011261436757600080fd5b813561437a6143758261432f565b6142ff565b81815284602083860101111561438f57600080fd5b816020850160208301376000918101602001919091529392505050565b600061012082360312156143bf57600080fd5b6143c761424c565b823581526143d760208401613f83565b602082015260408301356001600160401b038111156143f557600080fd5b61440136828601614356565b6040830152506060830135606082015261441d60808401613f83565b608082015260a083013560a082015260c083013560c082015260e083013560e082015261010061444e818501613f83565b9082015292915050565b6000610100828403121561446b57600080fd5b614473614275565b82358152602083013561448581613f6e565b6020820152604083013561449881613f6e565b60408201526144a960608401613f83565b60608201526080830135608082015260a083013560a082015260c083013560c08201526144d860e08401613f83565b60e08201529392505050565b600061016082360312156144f757600080fd5b6144ff614298565b8235815261450f60208401613f83565b602082015261452060408401613f83565b604082015260608301356001600160401b0381111561453e57600080fd5b61454a36828601614356565b6060830152506080830135608082015260a083013560a082015261457060c08401613f83565b60c082015261458160e08401613f83565b60e0820152610100614594818501613f83565b90820152610120838101359082015261014092830135928101929092525090565b6000602082840312156145c757600080fd5b5051919050565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b60208082526026908201527f546f6b656e486f6d653a206e6f6e2d7a65726f206d756c74692d686f702066616040820152656c6c6261636b60d01b606082015260800190565b634e487b7160e01b600052602160045260246000fd5b60005b83811015614689578181015183820152602001614671565b50506000910152565b600081518084526146aa81602086016020860161466e565b601f01601f19169290920160200192915050565b602081528151602082015260006020830151608060408401526146e460a0840182614692565b905060408401516060840152606084015160808401528091505092915050565b60208152600082516008811061472a57634e487b7160e01b600052602160045260246000fd5b8060208401525060208301516040808401526133e76060840182614692565b60408152825160408201526000602084015161477060608401826001600160a01b03169052565b50604084015161012080608085015261478d610160850183614692565b9150606086015160a085015260808601516147b360c08601826001600160a01b03169052565b5060a086015160e085015260c0860151610100818187015260e08801518387015280880151925050506147f26101408501826001600160a01b03169052565b506020929092019290925292915050565b81516001600160a01b031681526020808301519082015260408101610455565b60006101208201905083518252602084015160018060a01b03808216602085015280604087015116604085015280606087015116606085015250506080840151608083015260a084015160a083015260c084015160c083015260e084015161489660e08401826001600160a01b03169052565b50826101008301529392505050565b60208152815160208201526000602083015160018060a01b038082166040850152806040860151166060850152505060608301516148ee60808401826001600160a01b03169052565b50608083015160a083015260a08301516101008060c0850152614915610120850183614692565b915060c085015160e085015260e085015161493a828601826001600160a01b03169052565b5090949350505050565b60408152825160408201526000602084015161496b60608401826001600160a01b03169052565b5060408401516001600160a01b03166080830152606084015161016060a0840181905261499c6101a0850183614692565b9150608086015160c085015260a086015160e085015260c08601516101006149ce818701836001600160a01b03169052565b60e088015191506101206149ec818801846001600160a01b03169052565b90880151915061014090614a0a878301846001600160a01b03169052565b880151928601929092525090940151610180830152506020015290565b600082601f830112614a3857600080fd5b8151614a466143758261432f565b818152846020838601011115614a5b57600080fd5b6133e782602083016020870161466e565b600060208284031215614a7e57600080fd5b81516001600160401b0380821115614a9557600080fd5b9083019060408286031215614aa957600080fd5b614ab16142bb565b825160088110614ac057600080fd5b8152602083015182811115614ad457600080fd5b614ae087828601614a27565b60208301525095945050505050565b8051613f8e81613f6e565b600060408284031215614b0c57600080fd5b614b146142bb565b8251614b1f81613f6e565b81526020928301519281019290925250919050565b600060208284031215614b4657600080fd5b81516001600160401b0380821115614b5d57600080fd5b9083019060a08286031215614b7157600080fd5b614b796142dd565b82518152602083015182811115614b8f57600080fd5b614b9b87828601614a27565b6020830152506040830151604082015260608301519150614bbb82613f6e565b8160608201526080830151608082015280935050505092915050565b600060208284031215614be957600080fd5b81516001600160401b0380821115614c0057600080fd5b908301906101008286031215614c1557600080fd5b614c1d614275565b82518152614c2d60208401614aef565b6020820152614c3e60408401614aef565b6040820152614c4f60608401614aef565b60608201526080830151608082015260a083015182811115614c7057600080fd5b614c7c87828601614a27565b60a08301525060c083015160c0820152614c9860e08401614aef565b60e082015295945050505050565b600060e08284031215614cb857600080fd5b60405160e081018181106001600160401b0382111715614cda57614cda614236565b604052825181526020830151614cef81613f6e565b60208201526040830151614d0281613f6e565b80604083015250606083015160608201526080830151608082015260a083015160a082015260c0830151614d3581613f6e565b60c08201529392505050565b600060208284031215614d5357600080fd5b81516001600160401b0380821115614d6a57600080fd5b908301906101608286031215614d7f57600080fd5b614d87614298565b614d9083614aef565b815260208301516020820152614da860408401614aef565b6040820152614db960608401614aef565b60608201526080830151608082015260a083015182811115614dda57600080fd5b614de687828601614a27565b60a08301525060c083015160c0820152614e0260e08401614aef565b60e082015261010083810151908201526101209150614e22828401614aef565b9181019190915261014091820151918101919091529392505050565b600060608284031215614e5057600080fd5b604051606081018181106001600160401b0382111715614e7257614e72614236565b604052825181526020830151614e8781613fe0565b60208201526040830151614e9a81613fe0565b60408201529392505050565b6020808252818101527f546f6b656e486f6d653a2072656d6f7465206e6f742072656769737465726564604082015260600190565b634e487b7160e01b600052601160045260246000fd5b8181038181111561045557610455614edb565b60208082526021908201527f546f6b656e486f6d653a207a65726f20726563697069656e74206164647265736040820152607360f81b606082015260800190565b60208082526022908201527f546f6b656e486f6d653a207a65726f20726571756972656420676173206c696d6040820152611a5d60f21b606082015260800190565b60208082526021908201527f546f6b656e486f6d653a206e6f6e2d7a65726f207365636f6e646172792066656040820152606560f81b606082015260800190565b60208082526027908201527f546f6b656e486f6d653a20636f6c6c61746572616c206e656564656420666f726040820152662072656d6f746560c81b606082015260800190565b60006020828403121561502157600080fd5b815161080181613f6e565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501526000929161010085019190606087015160a0870152608087015160e060c0880152805193849052840192600092506101208701905b808410156150bb57845183168252938501936001939093019290850190615099565b5060a0880151878203601f190160e089015294506150d98186614692565b98975050505050505050565b8082018082111561045557610455614edb565b60006020828403121561510a57600080fd5b8151801515811461080157600080fd5b8681526001600160a01b0386811660208301528581166040830152841660608201526080810183905260c060a082018190526000906150d990830184614692565b634e487b7160e01b600052601260045260246000fd5b6000826151805761518061515b565b500690565b60ff828116828216039081111561045557610455614edb565b600181815b808511156151d95781600019048211156151bf576151bf614edb565b808516156151cc57918102915b93841c93908002906151a3565b509250929050565b6000826151f057506001610455565b816151fd57506000610455565b8160018114615213576002811461521d57615239565b6001915050610455565b60ff84111561522e5761522e614edb565b50506001821b610455565b5060208310610133831016604e8410600b841016171561525c575081810a610455565b615266838361519e565b806000190482111561527a5761527a614edb565b029392505050565b600061080183836151e1565b808202811582820484141761045557610455614edb565b6000826152b4576152b461515b565b500490565b600082516152cb81846020870161466e565b919091019291505056fe9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e6009316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e602d2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c7500914a9547f6c3ddce1d5efbd9e687708f0d1d408ce129e8e1a88bce4f40e29500de77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d00a2646970667358221220a00ce6e2fb2b647d87286a2300a72d3a39605d1bafb4a7e258089bf91998afd464736f6c6343000819003354656c65706f7274657252656769737472794170703a20696e76616c69642054",
}

// ERC20TokenHomeABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenHomeMetaData.ABI instead.
var ERC20TokenHomeABI = ERC20TokenHomeMetaData.ABI

// ERC20TokenHomeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenHomeMetaData.Bin instead.
var ERC20TokenHomeBin = ERC20TokenHomeMetaData.Bin

// DeployERC20TokenHome deploys a new Ethereum contract, binding an instance of ERC20TokenHome to it.
func DeployERC20TokenHome(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterRegistryAddress common.Address, teleporterManager common.Address, minTeleporterVersion *big.Int, tokenAddress common.Address, tokenDecimals uint8) (common.Address, *types.Transaction, *ERC20TokenHome, error) {
	parsed, err := ERC20TokenHomeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenHomeBin), backend, teleporterRegistryAddress, teleporterManager, minTeleporterVersion, tokenAddress, tokenDecimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20TokenHome{ERC20TokenHomeCaller: ERC20TokenHomeCaller{contract: contract}, ERC20TokenHomeTransactor: ERC20TokenHomeTransactor{contract: contract}, ERC20TokenHomeFilterer: ERC20TokenHomeFilterer{contract: contract}}, nil
}

// ERC20TokenHome is an auto generated Go binding around an Ethereum contract.
type ERC20TokenHome struct {
	ERC20TokenHomeCaller     // Read-only binding to the contract
	ERC20TokenHomeTransactor // Write-only binding to the contract
	ERC20TokenHomeFilterer   // Log filterer for contract events
}

// ERC20TokenHomeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenHomeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenHomeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenHomeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenHomeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenHomeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenHomeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenHomeSession struct {
	Contract     *ERC20TokenHome   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20TokenHomeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenHomeCallerSession struct {
	Contract *ERC20TokenHomeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC20TokenHomeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenHomeTransactorSession struct {
	Contract     *ERC20TokenHomeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20TokenHomeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenHomeRaw struct {
	Contract *ERC20TokenHome // Generic contract binding to access the raw methods on
}

// ERC20TokenHomeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenHomeCallerRaw struct {
	Contract *ERC20TokenHomeCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenHomeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenHomeTransactorRaw struct {
	Contract *ERC20TokenHomeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenHome creates a new instance of ERC20TokenHome, bound to a specific deployed contract.
func NewERC20TokenHome(address common.Address, backend bind.ContractBackend) (*ERC20TokenHome, error) {
	contract, err := bindERC20TokenHome(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHome{ERC20TokenHomeCaller: ERC20TokenHomeCaller{contract: contract}, ERC20TokenHomeTransactor: ERC20TokenHomeTransactor{contract: contract}, ERC20TokenHomeFilterer: ERC20TokenHomeFilterer{contract: contract}}, nil
}

// NewERC20TokenHomeCaller creates a new read-only instance of ERC20TokenHome, bound to a specific deployed contract.
func NewERC20TokenHomeCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenHomeCaller, error) {
	contract, err := bindERC20TokenHome(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeCaller{contract: contract}, nil
}

// NewERC20TokenHomeTransactor creates a new write-only instance of ERC20TokenHome, bound to a specific deployed contract.
func NewERC20TokenHomeTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenHomeTransactor, error) {
	contract, err := bindERC20TokenHome(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeTransactor{contract: contract}, nil
}

// NewERC20TokenHomeFilterer creates a new log filterer instance of ERC20TokenHome, bound to a specific deployed contract.
func NewERC20TokenHomeFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenHomeFilterer, error) {
	contract, err := bindERC20TokenHome(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeFilterer{contract: contract}, nil
}

// bindERC20TokenHome binds a generic wrapper to an already deployed contract.
func bindERC20TokenHome(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TokenHomeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenHome *ERC20TokenHomeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenHome.Contract.ERC20TokenHomeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenHome *ERC20TokenHomeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.ERC20TokenHomeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenHome *ERC20TokenHomeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.ERC20TokenHomeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenHome *ERC20TokenHomeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenHome.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenHome *ERC20TokenHomeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenHome *ERC20TokenHomeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.contract.Transact(opts, method, params...)
}

// ERC20TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x4797735f.
//
// Solidity: function ERC20_TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenHome *ERC20TokenHomeCaller) ERC20TOKENHOMESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenHome.contract.Call(opts, &out, "ERC20_TOKEN_HOME_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC20TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x4797735f.
//
// Solidity: function ERC20_TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenHome *ERC20TokenHomeSession) ERC20TOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenHome.Contract.ERC20TOKENHOMESTORAGELOCATION(&_ERC20TokenHome.CallOpts)
}

// ERC20TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x4797735f.
//
// Solidity: function ERC20_TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenHome *ERC20TokenHomeCallerSession) ERC20TOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenHome.Contract.ERC20TOKENHOMESTORAGELOCATION(&_ERC20TokenHome.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenHome *ERC20TokenHomeCaller) TELEPORTERREGISTRYAPPSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenHome.contract.Call(opts, &out, "TELEPORTER_REGISTRY_APP_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenHome *ERC20TokenHomeSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenHome.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_ERC20TokenHome.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenHome *ERC20TokenHomeCallerSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenHome.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_ERC20TokenHome.CallOpts)
}

// TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x62e3901b.
//
// Solidity: function TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenHome *ERC20TokenHomeCaller) TOKENHOMESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenHome.contract.Call(opts, &out, "TOKEN_HOME_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x62e3901b.
//
// Solidity: function TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenHome *ERC20TokenHomeSession) TOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenHome.Contract.TOKENHOMESTORAGELOCATION(&_ERC20TokenHome.CallOpts)
}

// TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x62e3901b.
//
// Solidity: function TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenHome *ERC20TokenHomeCallerSession) TOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenHome.Contract.TOKENHOMESTORAGELOCATION(&_ERC20TokenHome.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_ERC20TokenHome *ERC20TokenHomeCaller) GetBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenHome.contract.Call(opts, &out, "getBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_ERC20TokenHome *ERC20TokenHomeSession) GetBlockchainID() ([32]byte, error) {
	return _ERC20TokenHome.Contract.GetBlockchainID(&_ERC20TokenHome.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_ERC20TokenHome *ERC20TokenHomeCallerSession) GetBlockchainID() ([32]byte, error) {
	return _ERC20TokenHome.Contract.GetBlockchainID(&_ERC20TokenHome.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20TokenHome *ERC20TokenHomeCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenHome.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20TokenHome *ERC20TokenHomeSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ERC20TokenHome.Contract.GetMinTeleporterVersion(&_ERC20TokenHome.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20TokenHome *ERC20TokenHomeCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ERC20TokenHome.Contract.GetMinTeleporterVersion(&_ERC20TokenHome.CallOpts)
}

// GetRemoteTokenTransferrerSettings is a free data retrieval call binding the contract method 0xc8511ada.
//
// Solidity: function getRemoteTokenTransferrerSettings(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns((bool,uint256,uint256,bool))
func (_ERC20TokenHome *ERC20TokenHomeCaller) GetRemoteTokenTransferrerSettings(opts *bind.CallOpts, remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (RemoteTokenTransferrerSettings, error) {
	var out []interface{}
	err := _ERC20TokenHome.contract.Call(opts, &out, "getRemoteTokenTransferrerSettings", remoteBlockchainID, remoteTokenTransferrerAddress)

	if err != nil {
		return *new(RemoteTokenTransferrerSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(RemoteTokenTransferrerSettings)).(*RemoteTokenTransferrerSettings)

	return out0, err

}

// GetRemoteTokenTransferrerSettings is a free data retrieval call binding the contract method 0xc8511ada.
//
// Solidity: function getRemoteTokenTransferrerSettings(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns((bool,uint256,uint256,bool))
func (_ERC20TokenHome *ERC20TokenHomeSession) GetRemoteTokenTransferrerSettings(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (RemoteTokenTransferrerSettings, error) {
	return _ERC20TokenHome.Contract.GetRemoteTokenTransferrerSettings(&_ERC20TokenHome.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// GetRemoteTokenTransferrerSettings is a free data retrieval call binding the contract method 0xc8511ada.
//
// Solidity: function getRemoteTokenTransferrerSettings(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns((bool,uint256,uint256,bool))
func (_ERC20TokenHome *ERC20TokenHomeCallerSession) GetRemoteTokenTransferrerSettings(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (RemoteTokenTransferrerSettings, error) {
	return _ERC20TokenHome.Contract.GetRemoteTokenTransferrerSettings(&_ERC20TokenHome.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_ERC20TokenHome *ERC20TokenHomeCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenHome.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_ERC20TokenHome *ERC20TokenHomeSession) GetTokenAddress() (common.Address, error) {
	return _ERC20TokenHome.Contract.GetTokenAddress(&_ERC20TokenHome.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_ERC20TokenHome *ERC20TokenHomeCallerSession) GetTokenAddress() (common.Address, error) {
	return _ERC20TokenHome.Contract.GetTokenAddress(&_ERC20TokenHome.CallOpts)
}

// GetTransferredBalance is a free data retrieval call binding the contract method 0x154d625a.
//
// Solidity: function getTransferredBalance(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns(uint256)
func (_ERC20TokenHome *ERC20TokenHomeCaller) GetTransferredBalance(opts *bind.CallOpts, remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenHome.contract.Call(opts, &out, "getTransferredBalance", remoteBlockchainID, remoteTokenTransferrerAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransferredBalance is a free data retrieval call binding the contract method 0x154d625a.
//
// Solidity: function getTransferredBalance(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns(uint256)
func (_ERC20TokenHome *ERC20TokenHomeSession) GetTransferredBalance(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*big.Int, error) {
	return _ERC20TokenHome.Contract.GetTransferredBalance(&_ERC20TokenHome.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// GetTransferredBalance is a free data retrieval call binding the contract method 0x154d625a.
//
// Solidity: function getTransferredBalance(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns(uint256)
func (_ERC20TokenHome *ERC20TokenHomeCallerSession) GetTransferredBalance(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*big.Int, error) {
	return _ERC20TokenHome.Contract.GetTransferredBalance(&_ERC20TokenHome.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20TokenHome *ERC20TokenHomeCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20TokenHome.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20TokenHome *ERC20TokenHomeSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _ERC20TokenHome.Contract.IsTeleporterAddressPaused(&_ERC20TokenHome.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20TokenHome *ERC20TokenHomeCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _ERC20TokenHome.Contract.IsTeleporterAddressPaused(&_ERC20TokenHome.CallOpts, teleporterAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenHome *ERC20TokenHomeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenHome.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenHome *ERC20TokenHomeSession) Owner() (common.Address, error) {
	return _ERC20TokenHome.Contract.Owner(&_ERC20TokenHome.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenHome *ERC20TokenHomeCallerSession) Owner() (common.Address, error) {
	return _ERC20TokenHome.Contract.Owner(&_ERC20TokenHome.CallOpts)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xfd658268.
//
// Solidity: function addCollateral(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress, uint256 amount) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactor) AddCollateral(opts *bind.TransactOpts, remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.contract.Transact(opts, "addCollateral", remoteBlockchainID, remoteTokenTransferrerAddress, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xfd658268.
//
// Solidity: function addCollateral(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress, uint256 amount) returns()
func (_ERC20TokenHome *ERC20TokenHomeSession) AddCollateral(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.AddCollateral(&_ERC20TokenHome.TransactOpts, remoteBlockchainID, remoteTokenTransferrerAddress, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xfd658268.
//
// Solidity: function addCollateral(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress, uint256 amount) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactorSession) AddCollateral(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.AddCollateral(&_ERC20TokenHome.TransactOpts, remoteBlockchainID, remoteTokenTransferrerAddress, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x3bb03890.
//
// Solidity: function initialize(address teleporterRegistryAddress, address teleporterManager, uint256 minTeleporterVersion, address tokenAddress, uint8 tokenDecimals) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactor) Initialize(opts *bind.TransactOpts, teleporterRegistryAddress common.Address, teleporterManager common.Address, minTeleporterVersion *big.Int, tokenAddress common.Address, tokenDecimals uint8) (*types.Transaction, error) {
	return _ERC20TokenHome.contract.Transact(opts, "initialize", teleporterRegistryAddress, teleporterManager, minTeleporterVersion, tokenAddress, tokenDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x3bb03890.
//
// Solidity: function initialize(address teleporterRegistryAddress, address teleporterManager, uint256 minTeleporterVersion, address tokenAddress, uint8 tokenDecimals) returns()
func (_ERC20TokenHome *ERC20TokenHomeSession) Initialize(teleporterRegistryAddress common.Address, teleporterManager common.Address, minTeleporterVersion *big.Int, tokenAddress common.Address, tokenDecimals uint8) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.Initialize(&_ERC20TokenHome.TransactOpts, teleporterRegistryAddress, teleporterManager, minTeleporterVersion, tokenAddress, tokenDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x3bb03890.
//
// Solidity: function initialize(address teleporterRegistryAddress, address teleporterManager, uint256 minTeleporterVersion, address tokenAddress, uint8 tokenDecimals) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactorSession) Initialize(teleporterRegistryAddress common.Address, teleporterManager common.Address, minTeleporterVersion *big.Int, tokenAddress common.Address, tokenDecimals uint8) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.Initialize(&_ERC20TokenHome.TransactOpts, teleporterRegistryAddress, teleporterManager, minTeleporterVersion, tokenAddress, tokenDecimals)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenHome.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenHome *ERC20TokenHomeSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.PauseTeleporterAddress(&_ERC20TokenHome.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.PauseTeleporterAddress(&_ERC20TokenHome.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenHome.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenHome *ERC20TokenHomeSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.ReceiveTeleporterMessage(&_ERC20TokenHome.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.ReceiveTeleporterMessage(&_ERC20TokenHome.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenHome.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenHome *ERC20TokenHomeSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.RenounceOwnership(&_ERC20TokenHome.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.RenounceOwnership(&_ERC20TokenHome.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x5d16225d.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactor) Send(opts *bind.TransactOpts, input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.contract.Transact(opts, "send", input, amount)
}

// Send is a paid mutator transaction binding the contract method 0x5d16225d.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenHome *ERC20TokenHomeSession) Send(input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.Send(&_ERC20TokenHome.TransactOpts, input, amount)
}

// Send is a paid mutator transaction binding the contract method 0x5d16225d.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactorSession) Send(input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.Send(&_ERC20TokenHome.TransactOpts, input, amount)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x65690038.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.contract.Transact(opts, "sendAndCall", input, amount)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x65690038.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount) returns()
func (_ERC20TokenHome *ERC20TokenHomeSession) SendAndCall(input SendAndCallInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.SendAndCall(&_ERC20TokenHome.TransactOpts, input, amount)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x65690038.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactorSession) SendAndCall(input SendAndCallInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.SendAndCall(&_ERC20TokenHome.TransactOpts, input, amount)
}

// ToCosmosSend is a paid mutator transaction binding the contract method 0x3ea51daa.
//
// Solidity: function toCosmosSend((bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactor) ToCosmosSend(opts *bind.TransactOpts, input ToCosmosSendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.contract.Transact(opts, "toCosmosSend", input, amount)
}

// ToCosmosSend is a paid mutator transaction binding the contract method 0x3ea51daa.
//
// Solidity: function toCosmosSend((bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenHome *ERC20TokenHomeSession) ToCosmosSend(input ToCosmosSendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.ToCosmosSend(&_ERC20TokenHome.TransactOpts, input, amount)
}

// ToCosmosSend is a paid mutator transaction binding the contract method 0x3ea51daa.
//
// Solidity: function toCosmosSend((bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactorSession) ToCosmosSend(input ToCosmosSendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.ToCosmosSend(&_ERC20TokenHome.TransactOpts, input, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenHome.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenHome *ERC20TokenHomeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.TransferOwnership(&_ERC20TokenHome.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.TransferOwnership(&_ERC20TokenHome.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenHome.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenHome *ERC20TokenHomeSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.UnpauseTeleporterAddress(&_ERC20TokenHome.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.UnpauseTeleporterAddress(&_ERC20TokenHome.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20TokenHome *ERC20TokenHomeSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.UpdateMinTeleporterVersion(&_ERC20TokenHome.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20TokenHome *ERC20TokenHomeTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _ERC20TokenHome.Contract.UpdateMinTeleporterVersion(&_ERC20TokenHome.TransactOpts, version)
}

// ERC20TokenHomeCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the ERC20TokenHome contract.
type ERC20TokenHomeCallFailedIterator struct {
	Event *ERC20TokenHomeCallFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeCallFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeCallFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeCallFailed represents a CallFailed event raised by the ERC20TokenHome contract.
type ERC20TokenHomeCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*ERC20TokenHomeCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeCallFailedIterator{contract: _ERC20TokenHome.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeCallFailed)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "CallFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallFailed is a log parse operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseCallFailed(log types.Log) (*ERC20TokenHomeCallFailed, error) {
	event := new(ERC20TokenHomeCallFailed)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the ERC20TokenHome contract.
type ERC20TokenHomeCallSucceededIterator struct {
	Event *ERC20TokenHomeCallSucceeded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeCallSucceeded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeCallSucceeded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeCallSucceeded represents a CallSucceeded event raised by the ERC20TokenHome contract.
type ERC20TokenHomeCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*ERC20TokenHomeCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeCallSucceededIterator{contract: _ERC20TokenHome.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeCallSucceeded)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallSucceeded is a log parse operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseCallSucceeded(log types.Log) (*ERC20TokenHomeCallSucceeded, error) {
	event := new(ERC20TokenHomeCallSucceeded)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeCollateralAddedIterator is returned from FilterCollateralAdded and is used to iterate over the raw logs and unpacked data for CollateralAdded events raised by the ERC20TokenHome contract.
type ERC20TokenHomeCollateralAddedIterator struct {
	Event *ERC20TokenHomeCollateralAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeCollateralAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeCollateralAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeCollateralAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeCollateralAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeCollateralAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeCollateralAdded represents a CollateralAdded event raised by the ERC20TokenHome contract.
type ERC20TokenHomeCollateralAdded struct {
	RemoteBlockchainID            [32]byte
	RemoteTokenTransferrerAddress common.Address
	Amount                        *big.Int
	Remaining                     *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdded is a free log retrieval operation binding the contract event 0x6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6.
//
// Solidity: event CollateralAdded(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 amount, uint256 remaining)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterCollateralAdded(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (*ERC20TokenHomeCollateralAddedIterator, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "CollateralAdded", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeCollateralAddedIterator{contract: _ERC20TokenHome.contract, event: "CollateralAdded", logs: logs, sub: sub}, nil
}

// WatchCollateralAdded is a free log subscription operation binding the contract event 0x6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6.
//
// Solidity: event CollateralAdded(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 amount, uint256 remaining)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchCollateralAdded(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeCollateralAdded, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (event.Subscription, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "CollateralAdded", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeCollateralAdded)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollateralAdded is a log parse operation binding the contract event 0x6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6.
//
// Solidity: event CollateralAdded(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 amount, uint256 remaining)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseCollateralAdded(log types.Log) (*ERC20TokenHomeCollateralAdded, error) {
	event := new(ERC20TokenHomeCollateralAdded)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeFromCosmosTokensSentIterator is returned from FilterFromCosmosTokensSent and is used to iterate over the raw logs and unpacked data for FromCosmosTokensSent events raised by the ERC20TokenHome contract.
type ERC20TokenHomeFromCosmosTokensSentIterator struct {
	Event *ERC20TokenHomeFromCosmosTokensSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeFromCosmosTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeFromCosmosTokensSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeFromCosmosTokensSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeFromCosmosTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeFromCosmosTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeFromCosmosTokensSent represents a FromCosmosTokensSent event raised by the ERC20TokenHome contract.
type ERC20TokenHomeFromCosmosTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               FromCosmosSendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFromCosmosTokensSent is a free log retrieval operation binding the contract event 0xba11b5e7af3c06a53f098e63c9de142e163fc470b2e8565ff0de13af6010de2d.
//
// Solidity: event FromCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,bytes32,address,bytes32,bytes,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterFromCosmosTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*ERC20TokenHomeFromCosmosTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "FromCosmosTokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeFromCosmosTokensSentIterator{contract: _ERC20TokenHome.contract, event: "FromCosmosTokensSent", logs: logs, sub: sub}, nil
}

// WatchFromCosmosTokensSent is a free log subscription operation binding the contract event 0xba11b5e7af3c06a53f098e63c9de142e163fc470b2e8565ff0de13af6010de2d.
//
// Solidity: event FromCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,bytes32,address,bytes32,bytes,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchFromCosmosTokensSent(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeFromCosmosTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "FromCosmosTokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeFromCosmosTokensSent)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "FromCosmosTokensSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFromCosmosTokensSent is a log parse operation binding the contract event 0xba11b5e7af3c06a53f098e63c9de142e163fc470b2e8565ff0de13af6010de2d.
//
// Solidity: event FromCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,bytes32,address,bytes32,bytes,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseFromCosmosTokensSent(log types.Log) (*ERC20TokenHomeFromCosmosTokensSent, error) {
	event := new(ERC20TokenHomeFromCosmosTokensSent)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "FromCosmosTokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20TokenHome contract.
type ERC20TokenHomeInitializedIterator struct {
	Event *ERC20TokenHomeInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeInitialized represents a Initialized event raised by the ERC20TokenHome contract.
type ERC20TokenHomeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20TokenHomeInitializedIterator, error) {

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeInitializedIterator{contract: _ERC20TokenHome.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeInitialized)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseInitialized(log types.Log) (*ERC20TokenHomeInitialized, error) {
	event := new(ERC20TokenHomeInitialized)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the ERC20TokenHome contract.
type ERC20TokenHomeMinTeleporterVersionUpdatedIterator struct {
	Event *ERC20TokenHomeMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeMinTeleporterVersionUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeMinTeleporterVersionUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the ERC20TokenHome contract.
type ERC20TokenHomeMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*ERC20TokenHomeMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeMinTeleporterVersionUpdatedIterator{contract: _ERC20TokenHome.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeMinTeleporterVersionUpdated)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinTeleporterVersionUpdated is a log parse operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*ERC20TokenHomeMinTeleporterVersionUpdated, error) {
	event := new(ERC20TokenHomeMinTeleporterVersionUpdated)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20TokenHome contract.
type ERC20TokenHomeOwnershipTransferredIterator struct {
	Event *ERC20TokenHomeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20TokenHome contract.
type ERC20TokenHomeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20TokenHomeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeOwnershipTransferredIterator{contract: _ERC20TokenHome.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeOwnershipTransferred)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20TokenHomeOwnershipTransferred, error) {
	event := new(ERC20TokenHomeOwnershipTransferred)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeRemoteRegisteredIterator is returned from FilterRemoteRegistered and is used to iterate over the raw logs and unpacked data for RemoteRegistered events raised by the ERC20TokenHome contract.
type ERC20TokenHomeRemoteRegisteredIterator struct {
	Event *ERC20TokenHomeRemoteRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeRemoteRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeRemoteRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeRemoteRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeRemoteRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeRemoteRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeRemoteRegistered represents a RemoteRegistered event raised by the ERC20TokenHome contract.
type ERC20TokenHomeRemoteRegistered struct {
	RemoteBlockchainID            [32]byte
	RemoteTokenTransferrerAddress common.Address
	InitialCollateralNeeded       *big.Int
	TokenDecimals                 uint8
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterRemoteRegistered is a free log retrieval operation binding the contract event 0xf229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b90.
//
// Solidity: event RemoteRegistered(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 initialCollateralNeeded, uint8 tokenDecimals)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterRemoteRegistered(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (*ERC20TokenHomeRemoteRegisteredIterator, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "RemoteRegistered", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeRemoteRegisteredIterator{contract: _ERC20TokenHome.contract, event: "RemoteRegistered", logs: logs, sub: sub}, nil
}

// WatchRemoteRegistered is a free log subscription operation binding the contract event 0xf229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b90.
//
// Solidity: event RemoteRegistered(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 initialCollateralNeeded, uint8 tokenDecimals)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchRemoteRegistered(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeRemoteRegistered, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (event.Subscription, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "RemoteRegistered", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeRemoteRegistered)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "RemoteRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoteRegistered is a log parse operation binding the contract event 0xf229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b90.
//
// Solidity: event RemoteRegistered(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 initialCollateralNeeded, uint8 tokenDecimals)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseRemoteRegistered(log types.Log) (*ERC20TokenHomeRemoteRegistered, error) {
	event := new(ERC20TokenHomeRemoteRegistered)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "RemoteRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the ERC20TokenHome contract.
type ERC20TokenHomeTeleporterAddressPausedIterator struct {
	Event *ERC20TokenHomeTeleporterAddressPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeTeleporterAddressPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeTeleporterAddressPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the ERC20TokenHome contract.
type ERC20TokenHomeTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*ERC20TokenHomeTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeTeleporterAddressPausedIterator{contract: _ERC20TokenHome.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeTeleporterAddressPaused)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeleporterAddressPaused is a log parse operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseTeleporterAddressPaused(log types.Log) (*ERC20TokenHomeTeleporterAddressPaused, error) {
	event := new(ERC20TokenHomeTeleporterAddressPaused)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the ERC20TokenHome contract.
type ERC20TokenHomeTeleporterAddressUnpausedIterator struct {
	Event *ERC20TokenHomeTeleporterAddressUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeTeleporterAddressUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeTeleporterAddressUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the ERC20TokenHome contract.
type ERC20TokenHomeTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*ERC20TokenHomeTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeTeleporterAddressUnpausedIterator{contract: _ERC20TokenHome.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeTeleporterAddressUnpaused)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeleporterAddressUnpaused is a log parse operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*ERC20TokenHomeTeleporterAddressUnpaused, error) {
	event := new(ERC20TokenHomeTeleporterAddressUnpaused)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeToCosmosTokensSentIterator is returned from FilterToCosmosTokensSent and is used to iterate over the raw logs and unpacked data for ToCosmosTokensSent events raised by the ERC20TokenHome contract.
type ERC20TokenHomeToCosmosTokensSentIterator struct {
	Event *ERC20TokenHomeToCosmosTokensSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeToCosmosTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeToCosmosTokensSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeToCosmosTokensSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeToCosmosTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeToCosmosTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeToCosmosTokensSent represents a ToCosmosTokensSent event raised by the ERC20TokenHome contract.
type ERC20TokenHomeToCosmosTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               ToCosmosSendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterToCosmosTokensSent is a free log retrieval operation binding the contract event 0x79bd37563e6b913ff0a0e0c938210c007ebef4c942d38333576d03cb04620e9b.
//
// Solidity: event ToCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterToCosmosTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*ERC20TokenHomeToCosmosTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "ToCosmosTokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeToCosmosTokensSentIterator{contract: _ERC20TokenHome.contract, event: "ToCosmosTokensSent", logs: logs, sub: sub}, nil
}

// WatchToCosmosTokensSent is a free log subscription operation binding the contract event 0x79bd37563e6b913ff0a0e0c938210c007ebef4c942d38333576d03cb04620e9b.
//
// Solidity: event ToCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchToCosmosTokensSent(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeToCosmosTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "ToCosmosTokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeToCosmosTokensSent)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "ToCosmosTokensSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseToCosmosTokensSent is a log parse operation binding the contract event 0x79bd37563e6b913ff0a0e0c938210c007ebef4c942d38333576d03cb04620e9b.
//
// Solidity: event ToCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseToCosmosTokensSent(log types.Log) (*ERC20TokenHomeToCosmosTokensSent, error) {
	event := new(ERC20TokenHomeToCosmosTokensSent)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "ToCosmosTokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeTokensAndCallRoutedIterator is returned from FilterTokensAndCallRouted and is used to iterate over the raw logs and unpacked data for TokensAndCallRouted events raised by the ERC20TokenHome contract.
type ERC20TokenHomeTokensAndCallRoutedIterator struct {
	Event *ERC20TokenHomeTokensAndCallRouted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeTokensAndCallRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeTokensAndCallRouted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeTokensAndCallRouted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeTokensAndCallRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeTokensAndCallRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeTokensAndCallRouted represents a TokensAndCallRouted event raised by the ERC20TokenHome contract.
type ERC20TokenHomeTokensAndCallRouted struct {
	TeleporterMessageID [32]byte
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallRouted is a free log retrieval operation binding the contract event 0x42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterTokensAndCallRouted(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*ERC20TokenHomeTokensAndCallRoutedIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "TokensAndCallRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeTokensAndCallRoutedIterator{contract: _ERC20TokenHome.contract, event: "TokensAndCallRouted", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallRouted is a free log subscription operation binding the contract event 0x42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchTokensAndCallRouted(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeTokensAndCallRouted, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "TokensAndCallRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeTokensAndCallRouted)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "TokensAndCallRouted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensAndCallRouted is a log parse operation binding the contract event 0x42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseTokensAndCallRouted(log types.Log) (*ERC20TokenHomeTokensAndCallRouted, error) {
	event := new(ERC20TokenHomeTokensAndCallRouted)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "TokensAndCallRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the ERC20TokenHome contract.
type ERC20TokenHomeTokensAndCallSentIterator struct {
	Event *ERC20TokenHomeTokensAndCallSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeTokensAndCallSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeTokensAndCallSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeTokensAndCallSent represents a TokensAndCallSent event raised by the ERC20TokenHome contract.
type ERC20TokenHomeTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*ERC20TokenHomeTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeTokensAndCallSentIterator{contract: _ERC20TokenHome.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeTokensAndCallSent)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensAndCallSent is a log parse operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseTokensAndCallSent(log types.Log) (*ERC20TokenHomeTokensAndCallSent, error) {
	event := new(ERC20TokenHomeTokensAndCallSent)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeTokensCosmosWithdrawnIterator is returned from FilterTokensCosmosWithdrawn and is used to iterate over the raw logs and unpacked data for TokensCosmosWithdrawn events raised by the ERC20TokenHome contract.
type ERC20TokenHomeTokensCosmosWithdrawnIterator struct {
	Event *ERC20TokenHomeTokensCosmosWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeTokensCosmosWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeTokensCosmosWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeTokensCosmosWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeTokensCosmosWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeTokensCosmosWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeTokensCosmosWithdrawn represents a TokensCosmosWithdrawn event raised by the ERC20TokenHome contract.
type ERC20TokenHomeTokensCosmosWithdrawn struct {
	MessageID [32]byte
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensCosmosWithdrawn is a free log retrieval operation binding the contract event 0xfc043975d4effd7520860981e63ea2ba7197dbe72a39dd32a3e227e820ca7ffe.
//
// Solidity: event TokensCosmosWithdrawn(bytes32 messageID, address indexed recipient, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterTokensCosmosWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ERC20TokenHomeTokensCosmosWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "TokensCosmosWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeTokensCosmosWithdrawnIterator{contract: _ERC20TokenHome.contract, event: "TokensCosmosWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensCosmosWithdrawn is a free log subscription operation binding the contract event 0xfc043975d4effd7520860981e63ea2ba7197dbe72a39dd32a3e227e820ca7ffe.
//
// Solidity: event TokensCosmosWithdrawn(bytes32 messageID, address indexed recipient, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchTokensCosmosWithdrawn(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeTokensCosmosWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "TokensCosmosWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeTokensCosmosWithdrawn)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "TokensCosmosWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensCosmosWithdrawn is a log parse operation binding the contract event 0xfc043975d4effd7520860981e63ea2ba7197dbe72a39dd32a3e227e820ca7ffe.
//
// Solidity: event TokensCosmosWithdrawn(bytes32 messageID, address indexed recipient, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseTokensCosmosWithdrawn(log types.Log) (*ERC20TokenHomeTokensCosmosWithdrawn, error) {
	event := new(ERC20TokenHomeTokensCosmosWithdrawn)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "TokensCosmosWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeTokensRoutedIterator is returned from FilterTokensRouted and is used to iterate over the raw logs and unpacked data for TokensRouted events raised by the ERC20TokenHome contract.
type ERC20TokenHomeTokensRoutedIterator struct {
	Event *ERC20TokenHomeTokensRouted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeTokensRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeTokensRouted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeTokensRouted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeTokensRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeTokensRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeTokensRouted represents a TokensRouted event raised by the ERC20TokenHome contract.
type ERC20TokenHomeTokensRouted struct {
	TeleporterMessageID [32]byte
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensRouted is a free log retrieval operation binding the contract event 0x825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf0.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterTokensRouted(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*ERC20TokenHomeTokensRoutedIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "TokensRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeTokensRoutedIterator{contract: _ERC20TokenHome.contract, event: "TokensRouted", logs: logs, sub: sub}, nil
}

// WatchTokensRouted is a free log subscription operation binding the contract event 0x825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf0.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchTokensRouted(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeTokensRouted, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "TokensRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeTokensRouted)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "TokensRouted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensRouted is a log parse operation binding the contract event 0x825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf0.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseTokensRouted(log types.Log) (*ERC20TokenHomeTokensRouted, error) {
	event := new(ERC20TokenHomeTokensRouted)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "TokensRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the ERC20TokenHome contract.
type ERC20TokenHomeTokensSentIterator struct {
	Event *ERC20TokenHomeTokensSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeTokensSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeTokensSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeTokensSent represents a TokensSent event raised by the ERC20TokenHome contract.
type ERC20TokenHomeTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*ERC20TokenHomeTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeTokensSentIterator{contract: _ERC20TokenHome.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeTokensSent)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "TokensSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensSent is a log parse operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseTokensSent(log types.Log) (*ERC20TokenHomeTokensSent, error) {
	event := new(ERC20TokenHomeTokensSent)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenHomeTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the ERC20TokenHome contract.
type ERC20TokenHomeTokensWithdrawnIterator struct {
	Event *ERC20TokenHomeTokensWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenHomeTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenHomeTokensWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenHomeTokensWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenHomeTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenHomeTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenHomeTokensWithdrawn represents a TokensWithdrawn event raised by the ERC20TokenHome contract.
type ERC20TokenHomeTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ERC20TokenHomeTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenHomeTokensWithdrawnIterator{contract: _ERC20TokenHome.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *ERC20TokenHomeTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenHome.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenHomeTokensWithdrawn)
				if err := _ERC20TokenHome.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_ERC20TokenHome *ERC20TokenHomeFilterer) ParseTokensWithdrawn(log types.Log) (*ERC20TokenHomeTokensWithdrawn, error) {
	event := new(ERC20TokenHomeTokensWithdrawn)
	if err := _ERC20TokenHome.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
