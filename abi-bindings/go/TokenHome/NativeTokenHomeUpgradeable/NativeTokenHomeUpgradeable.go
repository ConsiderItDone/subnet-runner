// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenhomeupgradeable

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

// NativeTokenHomeUpgradeableMetaData contains all meta data concerning the NativeTokenHomeUpgradeable contract.
var NativeTokenHomeUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICTTInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"NATIVE_TOKEN_HOME_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TELEPORTER_REGISTRY_APP_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_HOME_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addCollateral\",\"inputs\":[{\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getBlockchainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinTeleporterVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteTokenTransferrerSettings\",\"inputs\":[{\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRemoteTokenTransferrerSettings\",\"components\":[{\"name\":\"registered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"collateralNeeded\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenMultiplier\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiplyOnRemote\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTransferredBalance\",\"inputs\":[{\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"teleporterRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"teleporterManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minTeleporterVersion\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"wrappedTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTeleporterAddressPaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseTeleporterAddress\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receiveTeleporterMessage\",\"inputs\":[{\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originSenderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structSendTokensInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sendAndCall\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structSendAndCallInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"toCosmosSend\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structToCosmosSendTokensInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationCosmosRecipient\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destinationCosmosBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseTeleporterAddress\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMinTeleporterVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"recipientContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"recipientContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralAdded\",\"inputs\":[{\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"remaining\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FromCosmosTokensSent\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFromCosmosSendTokensInput\",\"components\":[{\"name\":\"messageID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceCosmosBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceCosmosAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinTeleporterVersionUpdated\",\"inputs\":[{\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteRegistered\",\"inputs\":[{\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"initialCollateralNeeded\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TeleporterAddressPaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TeleporterAddressUnpaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ToCosmosTokensSent\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structToCosmosSendTokensInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationCosmosRecipient\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destinationCosmosBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensAndCallRouted\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensAndCallSent\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensCosmosWithdrawn\",\"inputs\":[{\"name\":\"messageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensRouted\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensSent\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516154a43803806154a483398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b6153578061014d5f395ff3fe608060405260043610610129575f3560e01c80638bf2fa94116100a8578063be2030941161006d578063be203094146103d4578063c8511ada146103f3578063c868efaa146104d3578063d2cc7a70146104f2578063efb5b95e14610525578063f2fde38b14610545575f80fd5b80638bf2fa94146103235780638da5cb5b14610336578063909a6ac0146103725780639731429714610392578063b0b78b26146103c1575f80fd5b80635eb99514116100ee5780635eb99514146102aa57806362e3901b146102c95780636e6eef8d146102e9578063715018a6146102fc5780637675695e14610310575f80fd5b806310fe9ae8146101ca578063154d625a1461021f5780632b0d8f181461024c5780634213cf781461026b5780634511243e1461028b575f80fd5b366101c6577f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e601546001600160a01b031633146101c45760405162461bcd60e51b815260206004820152602f60248201527f4e6174697665546f6b656e486f6d653a20696e76616c6964207265636569766560448201526e103830bcb0b136329039b2b73232b960891b60648201526084015b60405180910390fd5b005b5f80fd5b3480156101d5575f80fd5b507f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e601546001600160a01b03165b6040516001600160a01b0390911681526020015b60405180910390f35b34801561022a575f80fd5b5061023e61023936600461405a565b610564565b604051908152602001610216565b348015610257575f80fd5b506101c4610266366004614088565b6105ac565b348015610276575f80fd5b505f805160206152ab8339815191525461023e565b348015610296575f80fd5b506101c46102a5366004614088565b6106a5565b3480156102b5575f80fd5b506101c46102c43660046140a3565b610794565b3480156102d4575f80fd5b5061023e5f805160206152ab83398151915281565b6101c46102f73660046140ba565b6107a8565b348015610307575f80fd5b506101c46107d1565b6101c461031e3660046140f1565b6107e4565b6101c4610331366004614128565b6107f6565b348015610341575f80fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610202565b34801561037d575f80fd5b5061023e5f8051602061532b83398151915281565b34801561039d575f80fd5b506103b16103ac366004614088565b61080e565b6040519015158152602001610216565b6101c46103cf36600461405a565b61082e565b3480156103df575f80fd5b506101c46103ee36600461413f565b61083d565b3480156103fe575f80fd5b5061049c61040d36600461405a565b60408051608080820183525f808352602080840182905283850182905260609384018290529581525f805160206152eb83398151915286528381206001600160a01b039590951681529385529282902082519384018352805460ff9081161515855260018201549585019590955260028101549284019290925260039091015490921615159181019190915290565b6040516102169190815115158152602080830151908201526040808301519082015260609182015115159181019190915260800190565b3480156104de575f80fd5b506101c46104ed36600461418f565b61094f565b3480156104fd575f80fd5b507fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d025461023e565b348015610530575f80fd5b5061023e5f805160206152cb83398151915281565b348015610550575f80fd5b506101c461055f366004614088565b610b0c565b5f8281527f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e603602090815260408083206001600160a01b03851684529091529020545b92915050565b5f8051602061532b8339815191526105c2610b46565b6001600160a01b0382166105e85760405162461bcd60e51b81526004016101bb90614210565b6105f28183610b4e565b156106555760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b60648201526084016101bb565b6001600160a01b0382165f81815260018381016020526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a25050565b5f8051602061532b8339815191526106bb610b46565b6001600160a01b0382166106e15760405162461bcd60e51b81526004016101bb90614210565b6106eb8183610b4e565b6107495760405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472794170703a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b60648201526084016101bb565b6001600160a01b0382165f818152600183016020526040808220805460ff19169055517f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c39190a25050565b61079c610b46565b6107a581610b6f565b50565b6107a56107c05f805160206152ab8339815191525490565b30336107cb856143cd565b34610d07565b6107d9610f6a565b6107e25f610fc5565b565b6107a56107f08261449b565b34611035565b6107a561080836839003830183614544565b3461126e565b5f5f8051602061532b8339815191526108278184610b4e565b9392505050565b610839828234611413565b5050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156108815750825b90505f826001600160401b0316600114801561089c5750303b155b9050811580156108aa575080155b156108c85760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156108f257845460ff60401b1916600160401b1785555b6108fe89898989611603565b831561094457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b610957611622565b5f5f8051602061532b83398151915260028101548154919250906001600160a01b0316634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa1580156109c2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109e691906145ce565b1015610a4d5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b60648201526084016101bb565b610a578133610b4e565b15610abd5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b60648201526084016101bb565b610afd858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061166c92505050565b50610b06611abb565b50505050565b610b14610f6a565b6001600160a01b038116610b3d57604051631e4fbdf760e01b81525f60048201526024016101bb565b6107a581610fc5565b6107e2610f6a565b6001600160a01b03165f908152600191909101602052604090205460ff1690565b5f8051602061532b83398151915280546040805163301fd1f560e21b815290515f926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa158015610bc3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610be791906145ce565b600283015490915081841115610c595760405162461bcd60e51b815260206004820152603160248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b60648201526084016101bb565b808411610cce5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e0060648201526084016101bb565b60028301849055604051849082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d905f90a350505050565b5f8051602061530b8339815191528054600114610d365760405162461bcd60e51b81526004016101bb906145e5565b60028155610d4383611ae5565b60c08301516001600160a01b031615610d6e5760405162461bcd60e51b81526004016101bb90614629565b5f80610d8e855f0151866020015186886101000151896101200151611cca565b915091505f604051806040016040528060026007811115610db157610db161466f565b81526020016040518061010001604052808c81526020018b6001600160a01b031681526020018a6001600160a01b0316815260200189604001516001600160a01b03168152602001868152602001896060015181526020018960a0015181526020018960e001516001600160a01b0316815250604051602001610e3491906146d0565b60405160208183030381529060405281525090505f610f146040518060c00160405280895f0151815260200189602001516001600160a01b0316815260200160405180604001604052808b61010001516001600160a01b03168152602001878152508152602001896080015181526020015f6001600160401b03811115610ebd57610ebd61425e565b604051908082528060200260200182016040528015610ee6578160200160208202803683370190505b50815260200184604051602001610efd919061476e565b604051602081830303815290604052815250611e50565b9050876001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168987604051610f529291906147b0565b60405180910390a35050600190925550505050505050565b33610f9c7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146107e25760405163118cdaa760e01b81523360048201526024016101bb565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f8051602061530b83398151915280546001146110645760405162461bcd60e51b81526004016101bb906145e5565b6002815561107183611f6b565b6101008301516001600160a01b03161561109d5760405162461bcd60e51b81526004016101bb90614629565b5f806110bb855f015186602001518688608001518960a00151611fd4565b915091505f6040518060400160405280600560078111156110de576110de61466f565b815260200160405180608001604052806110f661211f565b8a516040516337c83a2f60e21b81526001600160a01b03929092169163df20e8bc916111289160040190815260200190565b602060405180830381865afa158015611143573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061116791906145ce565b815260200189604001518152602001895f01518152602001868152506040516020016111939190614892565b60405160208183030381529060405281525090505f61121b6040518060c00160405280895f0151815260200189602001516001600160a01b0316815260200160405180604001604052808b608001516001600160a01b031681526020018781525081526020018960e0015181526020015f6001600160401b03811115610ebd57610ebd61425e565b9050336001600160a01b0316817f79bd37563e6b913ff0a0e0c938210c007ebef4c942d38333576d03cb04620e9b89876040516112599291906148d7565b60405180910390a35050600190925550505050565b5f8051602061530b833981519152805460011461129d5760405162461bcd60e51b81526004016101bb906145e5565b600281556112aa8361220f565b60e08301516001600160a01b0316156112d55760405162461bcd60e51b81526004016101bb90614629565b5f806112f3855f015186602001518688606001518960800151611cca565b915091505f6040518060400160405280600160078111156113165761131661466f565b8152602001604051806040016040528089604001516001600160a01b031681526020018681525060405160200161134d9190614990565b60405160208183030381529060405281525090505f6113d56040518060c00160405280895f0151815260200189602001516001600160a01b0316815260200160405180604001604052808b606001516001600160a01b031681526020018781525081526020018960c0015181526020015f6001600160401b03811115610ebd57610ebd61425e565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb5289876040516112599291906149b0565b5f8051602061530b83398151915280546001146114425760405162461bcd60e51b81526004016101bb906145e5565b60028082555f8581525f805160206152eb833981519152602090815260408083206001600160a01b03881684528252918290208251608081018452815460ff9081161515808352600184015494830194909452948201549381019390935260030154909216151560608201525f805160206152ab833981519152916114d95760405162461bcd60e51b81526004016101bb90614a31565b5f8160200151116115365760405162461bcd60e51b815260206004820152602160248201527f546f6b656e486f6d653a207a65726f20636f6c6c61746572616c206e656564656044820152601960fa1b60648201526084016101bb565b61153f8461227e565b93505f808260200151861061156e5760208301515f92506115609087614a7a565b905082602001519550611581565b85836020015161157e9190614a7a565b91505b5f88815260028501602090815260408083206001600160a01b038b168085529083529281902060010185905580518981529182018590528a917f6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6910160405180910390a380156115f5576115f533826122a4565b505060019092555050505050565b61160b612362565b6116198484848460126123ab565b610b06816123d0565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080546001190161166657604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f5f805160206152ab83398151915290505f828060200190518101906116929190614acf565b90506001815160078111156116a9576116a961466f565b036116f1575f81602001518060200190518101906116c79190614b57565b90505f6116d987878460200151612406565b90506116e8825f0151826122a4565b50505050505050565b6006815160078111156117065761170661466f565b0361174c575f81602001518060200190518101906117249190614b8f565b90505f61173687878460800151612491565b90506116e88787846060015184865f015161258a565b6002815160078111156117615761176161466f565b0361187a575f816020015180602001905181019061177f9190614c2d565b90505f61179187878460800151612406565b825190915087146117f75760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e486f6d653a206d69736d61746368656420736f7572636520626c6f60448201526918dad8da185a5b88125160b21b60648201526084016101bb565b856001600160a01b031682602001516001600160a01b0316146118705760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e486f6d653a206d69736d617463686564206f726967696e2073656e60448201526a646572206164647265737360a81b60648201526084016101bb565b6116e8828261265a565b60038151600781111561188f5761188f61466f565b03611963575f81602001518060200190518101906118ad9190614cf7565b90505f806118c58888856060015186608001516127e0565b91509150611959604051806101000160405280855f0151815260200185602001516001600160a01b0316815260200185604001516001600160a01b03168152602001876001015f9054906101000a90046001600160a01b03166001600160a01b031681526020018381526020015f81526020018560a0015181526020018560c001516001600160a01b03168152508361288b565b5050505050505050565b6004815160078111156119785761197861466f565b03611a73575f81602001518060200190518101906119969190614d90565b90505f806119af888885608001518661014001516127e0565b915091506119598888855f01516040518061016001604052808860200151815260200188604001516001600160a01b0316815260200188606001516001600160a01b031681526020018860a00151815260200188610100015181526020018860c0015181526020018861012001516001600160a01b031681526020018860e001516001600160a01b031681526020018a6001015f9054906101000a90046001600160a01b03166001600160a01b031681526020018681526020015f81525086612a16565b5f81516007811115611a8757611a8761466f565b03611ab4575f8160200151806020019051810190611aa59190614e98565b9050611ab2868683612bf3565b505b5050505050565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005b6001905550565b60408101516001600160a01b0316611b525760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e486f6d653a207a65726f20726563697069656e7420636f6e7472616044820152696374206164647265737360b01b60648201526084016101bb565b5f816080015111611b755760405162461bcd60e51b81526004016101bb90614efa565b5f8160a0015111611bd45760405162461bcd60e51b815260206004820152602360248201527f546f6b656e486f6d653a207a65726f20726563697069656e7420676173206c696044820152621b5a5d60ea1b60648201526084016101bb565b80608001518160a0015110611c3a5760405162461bcd60e51b815260206004820152602660248201527f546f6b656e486f6d653a20696e76616c696420726563697069656e7420676173604482015265081b1a5b5a5d60d21b60648201526084016101bb565b60e08101516001600160a01b0316611ca75760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e486f6d653a207a65726f2066616c6c6261636b20726563697069656044820152696e74206164647265737360b01b60648201526084016101bb565b610140810151156107a55760405162461bcd60e51b81526004016101bb90614f3c565b5f8581525f805160206152eb833981519152602090815260408083206001600160a01b038816845282528083208151608081018352815460ff90811615158083526001840154958301959095526002830154938201939093526003909101549091161515606082015282915f805160206152ab8339815191529190611d615760405162461bcd60e51b81526004016101bb90614a31565b602081015115611d835760405162461bcd60e51b81526004016101bb90614f7d565b611d8c8761227e565b96508415611da357611da086335b87612ffa565b94505b5f611db7826040015183606001518a613153565b90505f8111611e085760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f207363616c656420616d6f756e7400000060448201526064016101bb565b5f8a815260038401602090815260408083206001600160a01b038d16845290915281208054839290611e3b908490614fc4565b90915550909a95995094975050505050505050565b5f80611e5a61211f565b60408401516020015190915015611eff576040830151516001600160a01b0316611edc5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a207a65726f206665652060448201526c746f6b656e206164647265737360981b60648201526084016101bb565b604083015160208101519051611eff916001600160a01b03909116908390613169565b604051630624488560e41b81526001600160a01b03821690636244885090611f2b908690600401614fd7565b6020604051808303815f875af1158015611f47573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061082791906145ce565b5f81604001515111611f8f5760405162461bcd60e51b81526004016101bb9061508e565b5f8160e0015111611fb25760405162461bcd60e51b81526004016101bb90614efa565b60c0810151156107a55760405162461bcd60e51b81526004016101bb90614f3c565b5f8581525f805160206152eb833981519152602090815260408083206001600160a01b038816845282528083208151608081018352815460ff90811615158083526001840154958301959095526002830154938201939093526003909101549091161515606082015282915f805160206152ab833981519152919061206b5760405162461bcd60e51b81526004016101bb90614a31565b60208101511561208d5760405162461bcd60e51b81526004016101bb90614f7d565b6120968761227e565b965084156120ab576120a88633611d9a565b94505b5f6120bf826040015183606001518a613153565b90505f81116121105760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f207363616c656420616d6f756e7400000060448201526064016101bb565b99949850939650505050505050565b5f8051602061532b83398151915280546040805163d820e64f60e01b815290515f939284926001600160a01b039091169163d820e64f916004808201926020929091908290030181865afa158015612179573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061219d91906150cf565b90506121a98282610b4e565b156105a65760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b60648201526084016101bb565b60408101516001600160a01b03166122395760405162461bcd60e51b81526004016101bb9061508e565b5f8160c001511161225c5760405162461bcd60e51b81526004016101bb90614efa565b60a0810151156107a55760405162461bcd60e51b81526004016101bb90614f3c565b5f805160206152cb83398151915280545f9190610827906001600160a01b0316846131f0565b6040518181525f805160206152cb833981519152906001600160a01b038416907f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b9060200160405180910390a28054604051632e1a7d4d60e01b8152600481018490526001600160a01b0390911690632e1a7d4d906024015f604051808303815f87803b158015612333575f80fd5b505af1158015612345573d5f803e3d5ffd5b5061235d925050506001600160a01b03841683613395565b505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166107e257604051631afcd79f60e31b815260040160405180910390fd5b6123b3612362565b6123be858585613428565b6123c6613443565b611ab48282613453565b6123d8612362565b5f805160206152cb83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b5f8381525f805160206152eb833981519152602090815260408083206001600160a01b038616845282528083208151608081018352815460ff9081161515825260018301549482019490945260028201549281019290925260030154909116151560608201525f805160206152ab83398151915290612487818787876135ba565b9695505050505050565b5f8381525f805160206152eb833981519152602090815260408083206001600160a01b038616845282528083208151608081018352815460ff9081161515808352600184015495830195909552600283015493820193909352600390910154909116151560608201525f805160206152ab833981519152916125255760405162461bcd60e51b81526004016101bb90614a31565b5f61253982604001518360600151876136a8565b90505f81116124875760405162461bcd60e51b815260206004820152601c60248201527f546f6b656e486f6d653a207a65726f20746f6b656e20616d6f756e740000000060448201526064016101bb565b5f5f805160206152cb83398151915260408051848152602081018690529192506001600160a01b038616917ffc043975d4effd7520860981e63ea2ba7197dbe72a39dd32a3e227e820ca7ffe910160405180910390a28054604051632e1a7d4d60e01b8152600481018590526001600160a01b0390911690632e1a7d4d906024015f604051808303815f87803b158015612622575f80fd5b505af1158015612634573d5f803e3d5ffd5b5061264c925050506001600160a01b03851684613395565b611ab28686868660016136b5565b5f5f805160206152cb8339815191528054604051632e1a7d4d60e01b8152600481018590529192506001600160a01b031690632e1a7d4d906024015f604051808303815f87803b1580156126ac575f80fd5b505af11580156126be573d5f803e3d5ffd5b50508451602086015160408088015160a089015191515f96506126e795509091906024016150ea565b60408051601f198184030181529190526020810180516001600160e01b031663161b12ff60e11b17905260c085015160608601519192505f9161272d9190869085613796565b905080156127815784606001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff48560405161277491815260200190565b60405180910390a2611ab4565b84606001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0856040516127c091815260200190565b60405180910390a260e0850151611ab4906001600160a01b031685613395565b5f8481525f805160206152eb833981519152602090815260408083206001600160a01b038716845282528083208151608081018352815460ff90811615158252600183015494820194909452600282015492810192909252600301549091161515606082015281905f805160206152ab8339815191529082612864828a8a8a6135ba565b90505f61287a83604001518460600151896136a8565b919a91995090975050505050505050565b5f8051602061530b83398151915280546001146128ba5760405162461bcd60e51b81526004016101bb906145e5565b600281556128c78361220f565b5f6128df845f01518560200151858760800151613866565b9050805f036128fc576128f68460e00151846122a4565b50612a0e565b604080518082019091525f908060018152602001604051806040016040528088604001516001600160a01b03168152602001858152506040516020016129429190614990565b60405160208183030381529060405281525090505f6129ce6040518060c00160405280885f0151815260200188602001516001600160a01b0316815260200160405180604001604052808a606001516001600160a01b031681526020018a6080015181525081526020018860c0015181526020015f6001600160401b03811115610ebd57610ebd61425e565b9050807f825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf08785604051612a029291906149b0565b60405180910390a25050505b600190555050565b5f8051602061530b8339815191528054600114612a455760405162461bcd60e51b81526004016101bb906145e5565b60028155612a5283611ae5565b5f612a6b845f0151856020015185876101200151613866565b9050805f03612a8857612a828460c00151846122a4565b50612be8565b604080518082019091525f9080600281526020016040518061010001604052808b81526020018a6001600160a01b03168152602001896001600160a01b0316815260200188604001516001600160a01b03168152602001858152602001886060015181526020018860a0015181526020018860e001516001600160a01b0316815250604051602001612b1a91906146d0565b60405160208183030381529060405281525090505f612ba86040518060c00160405280885f0151815260200188602001516001600160a01b0316815260200160405180604001604052808a61010001516001600160a01b031681526020018a61012001518152508152602001886080015181526020015f6001600160401b03811115610ebd57610ebd61425e565b9050807f42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb308785604051612bdc9291906147b0565b60405180910390a25050505b600190555050505050565b5f805160206152ab83398151915283612c5a5760405162461bcd60e51b8152602060048201526024808201527f546f6b656e486f6d653a207a65726f2072656d6f746520626c6f636b636861696044820152631b88125160e21b60648201526084016101bb565b80548403612cc25760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e486f6d653a2063616e6e6f742072656769737465722072656d6f7460448201526e329037b71039b0b6b29031b430b4b760891b60648201526084016101bb565b6001600160a01b038316612d315760405162461bcd60e51b815260206004820152603060248201527f546f6b656e486f6d653a207a65726f2072656d6f746520746f6b656e2074726160448201526f6e73666572726572206164647265737360801b60648201526084016101bb565b5f84815260028201602090815260408083206001600160a01b038716845290915290205460ff1615612db15760405162461bcd60e51b8152602060048201526024808201527f546f6b656e486f6d653a2072656d6f746520616c726561647920726567697374604482015263195c995960e21b60648201526084016101bb565b6012826040015160ff161115612e1b5760405162461bcd60e51b815260206004820152602960248201527f546f6b656e486f6d653a2072656d6f746520746f6b656e20646563696d616c73604482015268040e8dede40d0d2ced60bb1b60648201526084016101bb565b6001810154602083015160ff908116600160a01b9092041614612e8f5760405162461bcd60e51b815260206004820152602660248201527f546f6b656e486f6d653a20696e76616c696420686f6d6520746f6b656e20646560448201526563696d616c7360d01b60648201526084016101bb565b5f80612eb08360010160149054906101000a900460ff1685604001516139d7565b915091505f612ec38383875f01516136a8565b9050818015612edd57508451612eda90849061512f565b15155b15612ef057612eed600182614fc4565b90505b6040518060800160405280600115158152602001828152602001848152602001831515815250846002015f8981526020019081526020015f205f886001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a81548160ff02191690831515021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff021916908315150217905550905050856001600160a01b0316877ff229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b90838860400151604051612fe992919091825260ff16602082015260400190565b60405180910390a350505050505050565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa158015613040573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061306491906145ce565b905061307b6001600160a01b038616853086613a1f565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa1580156130bf573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130e391906145ce565b90508181116131495760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016101bb565b6124878282614a7a565b5f6131618484846001613a86565b949350505050565b604051636eb1769f60e11b81523060048201526001600160a01b0383811660248301525f919085169063dd62ed3e90604401602060405180830381865afa1580156131b6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131da91906145ce565b9050610b0684846131eb8585614fc4565b613aad565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038516906370a0823190602401602060405180830381865afa158015613236573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061325a91906145ce565b9050836001600160a01b031663d0e30db0846040518263ffffffff1660e01b81526004015f604051808303818588803b158015613295575f80fd5b505af11580156132a7573d5f803e3d5ffd5b50506040516370a0823160e01b81523060048201525f93506001600160a01b03881692506370a082319150602401602060405180830381865afa1580156132f0573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061331491906145ce565b90508181116133825760405162461bcd60e51b815260206004820152603460248201527f53616665577261707065644e6174697665546f6b656e4465706f7369743a2062604482015273185b185b98d9481b9bdd081a5b98dc99585cd95960621b60648201526084016101bb565b61338c8282614a7a565b95945050505050565b804710156133b85760405163cd78605960e01b81523060048201526024016101bb565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114613401576040519150601f19603f3d011682016040523d82523d5f602084013e613406565b606091505b505090508061235d57604051630a12f52160e11b815260040160405180910390fd5b613430612362565b61343a8382613b3c565b61235d82613b5e565b61344b612362565b6107e2613b6f565b61345b612362565b6001600160a01b0382166134b15760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f20746f6b656e206164647265737300000060448201526064016101bb565b60128160ff1611156135105760405162461bcd60e51b815260206004820152602260248201527f546f6b656e486f6d653a20746f6b656e20646563696d616c7320746f6f2068696044820152610ced60f31b60648201526084016101bb565b5f5f805160206152ab83398151915290506005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015613564573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061358891906145ce565b8155600101805460ff909216600160a01b026001600160a81b03199092166001600160a01b0390931692909217179055565b83515f906135da5760405162461bcd60e51b81526004016101bb90614a31565b6020850151156136385760405162461bcd60e51b8152602060048201526024808201527f546f6b656e486f6d653a2072656d6f7465206e6f7420636f6c6c61746572616c6044820152631a5e995960e21b60648201526084016101bb565b613643848484613b83565b5f61365786604001518760600151856136a8565b90505f811161338c5760405162461bcd60e51b815260206004820152601c60248201527f546f6b656e486f6d653a207a65726f20746f6b656e20616d6f756e740000000060448201526064016101bb565b5f6131618484845f613a86565b6136fe6040805160c0810182525f80825260208083018290528351808501855282815290810191909152909182019081526020015f815260200160608152602001606081525090565b8581526001600160a01b038581166020808401919091526040805180820182526007815281516060808201845294891680825281850189815288151592850192835284518087019290925251818501529051151581860152825180820390950185526080018252808301939093525161377892910161476e565b60408051601f1981840301815291905260a08201526116e881611e50565b5f845a10156137e75760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e7420676173000000000060448201526064016101bb565b834710156138375760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c756500000060448201526064016101bb565b826001600160a01b03163b5f0361384f57505f613161565b5f805f84516020860188888bf19695505050505050565b5f8481525f805160206152eb833981519152602090815260408083206001600160a01b038716845282528083208151608081018352815460ff9081161580158352600184015495830195909552600283015493820193909352600390910154909116151560608201525f805160206152ab83398151915291806138ec57505f8160200151115b156138fb575f92505050613161565b83851161395f5760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e486f6d653a20696e73756666696369656e7420616d6f756e74207460448201526b6f20636f766572206665657360a01b60648201526084016101bb565b6139698486614a7a565b94505f61397f8260400151836060015188613153565b9050805f03613993575f9350505050613161565b5f88815260038401602090815260408083206001600160a01b038b168452909152812080548392906139c6908490614fc4565b909155509098975050505050505050565b5f8060ff8085169084161181816139fa576139f28587615142565b60ff16613a08565b613a048686615142565b60ff165b613a1390600a61523b565b96919550909350505050565b6040516001600160a01b038481166024830152838116604483015260648201839052610b069186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613c70565b5f81151584151503613aa357613a9c8584615246565b9050613161565b61338c858461525d565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052613afe8482613cd1565b610b06576040516001600160a01b0384811660248301525f6044830152613b3291869182169063095ea7b390606401613a54565b610b068482613c70565b613b44612362565b613b4c613d6e565b613b54613d7e565b6108398282613d86565b613b66612362565b6107a581613f0a565b5f5f8051602061530b833981519152611ade565b5f8381527f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e603602090815260408083206001600160a01b03861684529091529020545f805160206152ab8339815191529082811015613c3a5760405162461bcd60e51b815260206004820152602e60248201527f546f6b656e486f6d653a20696e73756666696369656e7420746f6b656e20747260448201526d616e736665722062616c616e636560901b60648201526084016101bb565b613c448382614a7a565b5f9586526003909201602090815260408087206001600160a01b03909616875294905250919092205550565b5f613c846001600160a01b03841683613f12565b905080515f14158015613ca8575080806020019051810190613ca69190615270565b155b1561235d57604051635274afe760e01b81526001600160a01b03841660048201526024016101bb565b5f805f846001600160a01b031684604051613cec919061528f565b5f604051808303815f865af19150503d805f8114613d25576040519150601f19603f3d011682016040523d82523d5f602084013e613d2a565b606091505b5091509150818015613d54575080511580613d54575080806020019051810190613d549190615270565b801561338c5750505050506001600160a01b03163b151590565b613d76612362565b6107e2613f1f565b6107e2612362565b613d8e612362565b6001600160a01b038216613e0a5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084016101bb565b5f5f8051602061532b83398151915290505f8390505f816001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015613e5c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613e8091906145ce565b11613ee85760405162461bcd60e51b815260206004820152603260248201527f54656c65706f7274657252656769737472794170703a20696e76616c69642054604482015271656c65706f7274657220726567697374727960701b60648201526084016101bb565b81546001600160a01b0319166001600160a01b038216178255610b0683610b6f565b610b14612362565b606061082783835f613f27565b611abb612362565b606081471015613f4c5760405163cd78605960e01b81523060048201526024016101bb565b5f80856001600160a01b03168486604051613f67919061528f565b5f6040518083038185875af1925050503d805f8114613fa1576040519150601f19603f3d011682016040523d82523d5f602084013e613fa6565b606091505b5091509150612487868383606082613fc657613fc18261400d565b610827565b8151158015613fdd57506001600160a01b0384163b155b1561400657604051639996b31560e01b81526001600160a01b03851660048201526024016101bb565b5080610827565b80511561401d5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b6001600160a01b03811681146107a5575f80fd5b803561405581614036565b919050565b5f806040838503121561406b575f80fd5b82359150602083013561407d81614036565b809150509250929050565b5f60208284031215614098575f80fd5b813561082781614036565b5f602082840312156140b3575f80fd5b5035919050565b5f602082840312156140ca575f80fd5b81356001600160401b038111156140df575f80fd5b82016101608185031215610827575f80fd5b5f60208284031215614101575f80fd5b81356001600160401b03811115614116575f80fd5b82016101208185031215610827575f80fd5b5f6101008284031215614139575f80fd5b50919050565b5f805f8060808587031215614152575f80fd5b843561415d81614036565b9350602085013561416d81614036565b925060408501359150606085013561418481614036565b939692955090935050565b5f805f80606085870312156141a2575f80fd5b8435935060208501356141b481614036565b925060408501356001600160401b03808211156141cf575f80fd5b818701915087601f8301126141e2575f80fd5b8135818111156141f0575f80fd5b886020828501011115614201575f80fd5b95989497505060200194505050565b6020808252602e908201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b634e487b7160e01b5f52604160045260245ffd5b60405161016081016001600160401b03811182821017156142955761429561425e565b60405290565b60405161012081016001600160401b03811182821017156142955761429561425e565b60405161010081016001600160401b03811182821017156142955761429561425e565b604080519081016001600160401b03811182821017156142955761429561425e565b60405160a081016001600160401b03811182821017156142955761429561425e565b604051601f8201601f191681016001600160401b038111828210171561434d5761434d61425e565b604052919050565b5f6001600160401b0382111561436d5761436d61425e565b50601f01601f191660200190565b5f82601f83011261438a575f80fd5b813561439d61439882614355565b614325565b8181528460208386010111156143b1575f80fd5b816020850160208301375f918101602001919091529392505050565b5f61016082360312156143de575f80fd5b6143e6614272565b823581526143f66020840161404a565b60208201526144076040840161404a565b604082015260608301356001600160401b03811115614424575f80fd5b6144303682860161437b565b6060830152506080830135608082015260a083013560a082015261445660c0840161404a565b60c082015261446760e0840161404a565b60e082015261010061447a81850161404a565b90820152610120838101359082015261014092830135928101929092525090565b5f61012082360312156144ac575f80fd5b6144b461429b565b823581526144c46020840161404a565b602082015260408301356001600160401b038111156144e1575f80fd5b6144ed3682860161437b565b604083015250606083013560608201526145096080840161404a565b608082015260a083013560a082015260c083013560c082015260e083013560e082015261010061453a81850161404a565b9082015292915050565b5f6101008284031215614555575f80fd5b61455d6142be565b82358152602083013561456f81614036565b6020820152604083013561458281614036565b60408201526145936060840161404a565b60608201526080830135608082015260a083013560a082015260c083013560c08201526145c260e0840161404a565b60e08201529392505050565b5f602082840312156145de575f80fd5b5051919050565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b60208082526026908201527f546f6b656e486f6d653a206e6f6e2d7a65726f206d756c74692d686f702066616040820152656c6c6261636b60d01b606082015260800190565b634e487b7160e01b5f52602160045260245ffd5b5f5b8381101561469d578181015183820152602001614685565b50505f910152565b5f81518084526146bc816020860160208601614683565b601f01601f19169290920160200192915050565b60208152815160208201525f602083015160018060a01b0380821660408501528060408601511660608501525050606083015161471860808401826001600160a01b03169052565b50608083015160a083015260a08301516101008060c085015261473f6101208501836146a5565b915060c085015160e085015260e0850151614764828601826001600160a01b03169052565b5090949350505050565b602081525f82516008811061479157634e487b7160e01b5f52602160045260245ffd5b80602084015250602083015160408084015261316160608401826146a5565b60408152825160408201525f60208401516147d660608401826001600160a01b03169052565b5060408401516001600160a01b03166080830152606084015161016060a084018190526148076101a08501836146a5565b9150608086015160c085015260a086015160e085015260c0860151610100614839818701836001600160a01b03169052565b60e08801519150610120614857818801846001600160a01b03169052565b90880151915061014090614875878301846001600160a01b03169052565b880151928601929092525090940151610180830152506020015290565b60208152815160208201525f6020830151608060408401526148b760a08401826146a5565b905060408401516060840152606084015160808401528091505092915050565b60408152825160408201525f60208401516148fd60608401826001600160a01b03169052565b50604084015161012080608085015261491a6101608501836146a5565b9150606086015160a0850152608086015161494060c08601826001600160a01b03169052565b5060a086015160e085015260c0860151610100818187015260e088015183870152808801519250505061497f6101408501826001600160a01b03169052565b506020929092019290925292915050565b81516001600160a01b0316815260208083015190820152604081016105a6565b5f6101208201905083518252602084015160018060a01b03808216602085015280604087015116604085015280606087015116606085015250506080840151608083015260a084015160a083015260c084015160c083015260e0840151614a2260e08401826001600160a01b03169052565b50826101008301529392505050565b6020808252818101527f546f6b656e486f6d653a2072656d6f7465206e6f742072656769737465726564604082015260600190565b634e487b7160e01b5f52601160045260245ffd5b818103818111156105a6576105a6614a66565b5f82601f830112614a9c575f80fd5b8151614aaa61439882614355565b818152846020838601011115614abe575f80fd5b613161826020830160208701614683565b5f60208284031215614adf575f80fd5b81516001600160401b0380821115614af5575f80fd5b9083019060408286031215614b08575f80fd5b614b106142e1565b825160088110614b1e575f80fd5b8152602083015182811115614b31575f80fd5b614b3d87828601614a8d565b60208301525095945050505050565b805161405581614036565b5f60408284031215614b67575f80fd5b614b6f6142e1565b8251614b7a81614036565b81526020928301519281019290925250919050565b5f60208284031215614b9f575f80fd5b81516001600160401b0380821115614bb5575f80fd5b9083019060a08286031215614bc8575f80fd5b614bd0614303565b82518152602083015182811115614be5575f80fd5b614bf187828601614a8d565b6020830152506040830151604082015260608301519150614c1182614036565b8160608201526080830151608082015280935050505092915050565b5f60208284031215614c3d575f80fd5b81516001600160401b0380821115614c53575f80fd5b908301906101008286031215614c67575f80fd5b614c6f6142be565b82518152614c7f60208401614b4c565b6020820152614c9060408401614b4c565b6040820152614ca160608401614b4c565b60608201526080830151608082015260a083015182811115614cc1575f80fd5b614ccd87828601614a8d565b60a08301525060c083015160c0820152614ce960e08401614b4c565b60e082015295945050505050565b5f60e08284031215614d07575f80fd5b60405160e081018181106001600160401b0382111715614d2957614d2961425e565b604052825181526020830151614d3e81614036565b60208201526040830151614d5181614036565b80604083015250606083015160608201526080830151608082015260a083015160a082015260c0830151614d8481614036565b60c08201529392505050565b5f60208284031215614da0575f80fd5b81516001600160401b0380821115614db6575f80fd5b908301906101608286031215614dca575f80fd5b614dd2614272565b614ddb83614b4c565b815260208301516020820152614df360408401614b4c565b6040820152614e0460608401614b4c565b60608201526080830151608082015260a083015182811115614e24575f80fd5b614e3087828601614a8d565b60a08301525060c083015160c0820152614e4c60e08401614b4c565b60e082015261010083810151908201526101209150614e6c828401614b4c565b9181019190915261014091820151918101919091529392505050565b805160ff81168114614055575f80fd5b5f60608284031215614ea8575f80fd5b604051606081018181106001600160401b0382111715614eca57614eca61425e565b60405282518152614edd60208401614e88565b6020820152614eee60408401614e88565b60408201529392505050565b60208082526022908201527f546f6b656e486f6d653a207a65726f20726571756972656420676173206c696d6040820152611a5d60f21b606082015260800190565b60208082526021908201527f546f6b656e486f6d653a206e6f6e2d7a65726f207365636f6e646172792066656040820152606560f81b606082015260800190565b60208082526027908201527f546f6b656e486f6d653a20636f6c6c61746572616c206e656564656420666f726040820152662072656d6f746560c81b606082015260800190565b808201808211156105a6576105a6614a66565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501525f929161010085019190606087015160a0870152608087015160e060c08801528051938490528401925f92506101208701905b8084101561506457845183168252938501936001939093019290850190615042565b5060a0880151878203601f190160e0890152945061508281866146a5565b98975050505050505050565b60208082526021908201527f546f6b656e486f6d653a207a65726f20726563697069656e74206164647265736040820152607360f81b606082015260800190565b5f602082840312156150df575f80fd5b815161082781614036565b8481526001600160a01b038481166020830152831660408201526080606082018190525f90612487908301846146a5565b634e487b7160e01b5f52601260045260245ffd5b5f8261513d5761513d61511b565b500690565b60ff82811682821603908111156105a6576105a6614a66565b600181815b8085111561519557815f190482111561517b5761517b614a66565b8085161561518857918102915b93841c9390800290615160565b509250929050565b5f826151ab575060016105a6565b816151b757505f6105a6565b81600181146151cd57600281146151d7576151f3565b60019150506105a6565b60ff8411156151e8576151e8614a66565b50506001821b6105a6565b5060208310610133831016604e8410600b8410161715615216575081810a6105a6565b615220838361515b565b805f190482111561523357615233614a66565b029392505050565b5f610827838361519d565b80820281158282048414176105a6576105a6614a66565b5f8261526b5761526b61511b565b500490565b5f60208284031215615280575f80fd5b81518015158114610827575f80fd5b5f82516152a0818460208701614683565b919091019291505056fe9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e6003b5030f10c94fcbdaa3022348ff0b82dbd4c0c71339e41ff59d0bdc92179d6009316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e602d2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c7500de77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d00a164736f6c6343000819000a",
}

// NativeTokenHomeUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenHomeUpgradeableMetaData.ABI instead.
var NativeTokenHomeUpgradeableABI = NativeTokenHomeUpgradeableMetaData.ABI

// NativeTokenHomeUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenHomeUpgradeableMetaData.Bin instead.
var NativeTokenHomeUpgradeableBin = NativeTokenHomeUpgradeableMetaData.Bin

// DeployNativeTokenHomeUpgradeable deploys a new Ethereum contract, binding an instance of NativeTokenHomeUpgradeable to it.
func DeployNativeTokenHomeUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *NativeTokenHomeUpgradeable, error) {
	parsed, err := NativeTokenHomeUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenHomeUpgradeableBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenHomeUpgradeable{NativeTokenHomeUpgradeableCaller: NativeTokenHomeUpgradeableCaller{contract: contract}, NativeTokenHomeUpgradeableTransactor: NativeTokenHomeUpgradeableTransactor{contract: contract}, NativeTokenHomeUpgradeableFilterer: NativeTokenHomeUpgradeableFilterer{contract: contract}}, nil
}

// NativeTokenHomeUpgradeable is an auto generated Go binding around an Ethereum contract.
type NativeTokenHomeUpgradeable struct {
	NativeTokenHomeUpgradeableCaller     // Read-only binding to the contract
	NativeTokenHomeUpgradeableTransactor // Write-only binding to the contract
	NativeTokenHomeUpgradeableFilterer   // Log filterer for contract events
}

// NativeTokenHomeUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenHomeUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenHomeUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenHomeUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenHomeUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenHomeUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenHomeUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenHomeUpgradeableSession struct {
	Contract     *NativeTokenHomeUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// NativeTokenHomeUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenHomeUpgradeableCallerSession struct {
	Contract *NativeTokenHomeUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// NativeTokenHomeUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenHomeUpgradeableTransactorSession struct {
	Contract     *NativeTokenHomeUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// NativeTokenHomeUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenHomeUpgradeableRaw struct {
	Contract *NativeTokenHomeUpgradeable // Generic contract binding to access the raw methods on
}

// NativeTokenHomeUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenHomeUpgradeableCallerRaw struct {
	Contract *NativeTokenHomeUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenHomeUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenHomeUpgradeableTransactorRaw struct {
	Contract *NativeTokenHomeUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenHomeUpgradeable creates a new instance of NativeTokenHomeUpgradeable, bound to a specific deployed contract.
func NewNativeTokenHomeUpgradeable(address common.Address, backend bind.ContractBackend) (*NativeTokenHomeUpgradeable, error) {
	contract, err := bindNativeTokenHomeUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeable{NativeTokenHomeUpgradeableCaller: NativeTokenHomeUpgradeableCaller{contract: contract}, NativeTokenHomeUpgradeableTransactor: NativeTokenHomeUpgradeableTransactor{contract: contract}, NativeTokenHomeUpgradeableFilterer: NativeTokenHomeUpgradeableFilterer{contract: contract}}, nil
}

// NewNativeTokenHomeUpgradeableCaller creates a new read-only instance of NativeTokenHomeUpgradeable, bound to a specific deployed contract.
func NewNativeTokenHomeUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenHomeUpgradeableCaller, error) {
	contract, err := bindNativeTokenHomeUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableCaller{contract: contract}, nil
}

// NewNativeTokenHomeUpgradeableTransactor creates a new write-only instance of NativeTokenHomeUpgradeable, bound to a specific deployed contract.
func NewNativeTokenHomeUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenHomeUpgradeableTransactor, error) {
	contract, err := bindNativeTokenHomeUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableTransactor{contract: contract}, nil
}

// NewNativeTokenHomeUpgradeableFilterer creates a new log filterer instance of NativeTokenHomeUpgradeable, bound to a specific deployed contract.
func NewNativeTokenHomeUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenHomeUpgradeableFilterer, error) {
	contract, err := bindNativeTokenHomeUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableFilterer{contract: contract}, nil
}

// bindNativeTokenHomeUpgradeable binds a generic wrapper to an already deployed contract.
func bindNativeTokenHomeUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenHomeUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenHomeUpgradeable.Contract.NativeTokenHomeUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.NativeTokenHomeUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.NativeTokenHomeUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenHomeUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// NATIVETOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0xefb5b95e.
//
// Solidity: function NATIVE_TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCaller) NATIVETOKENHOMESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenHomeUpgradeable.contract.Call(opts, &out, "NATIVE_TOKEN_HOME_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NATIVETOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0xefb5b95e.
//
// Solidity: function NATIVE_TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) NATIVETOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHomeUpgradeable.Contract.NATIVETOKENHOMESTORAGELOCATION(&_NativeTokenHomeUpgradeable.CallOpts)
}

// NATIVETOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0xefb5b95e.
//
// Solidity: function NATIVE_TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCallerSession) NATIVETOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHomeUpgradeable.Contract.NATIVETOKENHOMESTORAGELOCATION(&_NativeTokenHomeUpgradeable.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCaller) TELEPORTERREGISTRYAPPSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenHomeUpgradeable.contract.Call(opts, &out, "TELEPORTER_REGISTRY_APP_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHomeUpgradeable.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_NativeTokenHomeUpgradeable.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCallerSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHomeUpgradeable.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_NativeTokenHomeUpgradeable.CallOpts)
}

// TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x62e3901b.
//
// Solidity: function TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCaller) TOKENHOMESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenHomeUpgradeable.contract.Call(opts, &out, "TOKEN_HOME_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x62e3901b.
//
// Solidity: function TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) TOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHomeUpgradeable.Contract.TOKENHOMESTORAGELOCATION(&_NativeTokenHomeUpgradeable.CallOpts)
}

// TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x62e3901b.
//
// Solidity: function TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCallerSession) TOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHomeUpgradeable.Contract.TOKENHOMESTORAGELOCATION(&_NativeTokenHomeUpgradeable.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCaller) GetBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenHomeUpgradeable.contract.Call(opts, &out, "getBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) GetBlockchainID() ([32]byte, error) {
	return _NativeTokenHomeUpgradeable.Contract.GetBlockchainID(&_NativeTokenHomeUpgradeable.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCallerSession) GetBlockchainID() ([32]byte, error) {
	return _NativeTokenHomeUpgradeable.Contract.GetBlockchainID(&_NativeTokenHomeUpgradeable.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenHomeUpgradeable.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenHomeUpgradeable.Contract.GetMinTeleporterVersion(&_NativeTokenHomeUpgradeable.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenHomeUpgradeable.Contract.GetMinTeleporterVersion(&_NativeTokenHomeUpgradeable.CallOpts)
}

// GetRemoteTokenTransferrerSettings is a free data retrieval call binding the contract method 0xc8511ada.
//
// Solidity: function getRemoteTokenTransferrerSettings(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns((bool,uint256,uint256,bool))
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCaller) GetRemoteTokenTransferrerSettings(opts *bind.CallOpts, remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (RemoteTokenTransferrerSettings, error) {
	var out []interface{}
	err := _NativeTokenHomeUpgradeable.contract.Call(opts, &out, "getRemoteTokenTransferrerSettings", remoteBlockchainID, remoteTokenTransferrerAddress)

	if err != nil {
		return *new(RemoteTokenTransferrerSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(RemoteTokenTransferrerSettings)).(*RemoteTokenTransferrerSettings)

	return out0, err

}

// GetRemoteTokenTransferrerSettings is a free data retrieval call binding the contract method 0xc8511ada.
//
// Solidity: function getRemoteTokenTransferrerSettings(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns((bool,uint256,uint256,bool))
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) GetRemoteTokenTransferrerSettings(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (RemoteTokenTransferrerSettings, error) {
	return _NativeTokenHomeUpgradeable.Contract.GetRemoteTokenTransferrerSettings(&_NativeTokenHomeUpgradeable.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// GetRemoteTokenTransferrerSettings is a free data retrieval call binding the contract method 0xc8511ada.
//
// Solidity: function getRemoteTokenTransferrerSettings(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns((bool,uint256,uint256,bool))
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCallerSession) GetRemoteTokenTransferrerSettings(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (RemoteTokenTransferrerSettings, error) {
	return _NativeTokenHomeUpgradeable.Contract.GetRemoteTokenTransferrerSettings(&_NativeTokenHomeUpgradeable.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenHomeUpgradeable.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) GetTokenAddress() (common.Address, error) {
	return _NativeTokenHomeUpgradeable.Contract.GetTokenAddress(&_NativeTokenHomeUpgradeable.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCallerSession) GetTokenAddress() (common.Address, error) {
	return _NativeTokenHomeUpgradeable.Contract.GetTokenAddress(&_NativeTokenHomeUpgradeable.CallOpts)
}

// GetTransferredBalance is a free data retrieval call binding the contract method 0x154d625a.
//
// Solidity: function getTransferredBalance(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns(uint256)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCaller) GetTransferredBalance(opts *bind.CallOpts, remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenHomeUpgradeable.contract.Call(opts, &out, "getTransferredBalance", remoteBlockchainID, remoteTokenTransferrerAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransferredBalance is a free data retrieval call binding the contract method 0x154d625a.
//
// Solidity: function getTransferredBalance(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns(uint256)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) GetTransferredBalance(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*big.Int, error) {
	return _NativeTokenHomeUpgradeable.Contract.GetTransferredBalance(&_NativeTokenHomeUpgradeable.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// GetTransferredBalance is a free data retrieval call binding the contract method 0x154d625a.
//
// Solidity: function getTransferredBalance(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns(uint256)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCallerSession) GetTransferredBalance(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*big.Int, error) {
	return _NativeTokenHomeUpgradeable.Contract.GetTransferredBalance(&_NativeTokenHomeUpgradeable.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenHomeUpgradeable.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenHomeUpgradeable.Contract.IsTeleporterAddressPaused(&_NativeTokenHomeUpgradeable.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenHomeUpgradeable.Contract.IsTeleporterAddressPaused(&_NativeTokenHomeUpgradeable.CallOpts, teleporterAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenHomeUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) Owner() (common.Address, error) {
	return _NativeTokenHomeUpgradeable.Contract.Owner(&_NativeTokenHomeUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableCallerSession) Owner() (common.Address, error) {
	return _NativeTokenHomeUpgradeable.Contract.Owner(&_NativeTokenHomeUpgradeable.CallOpts)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xb0b78b26.
//
// Solidity: function addCollateral(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactor) AddCollateral(opts *bind.TransactOpts, remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.contract.Transact(opts, "addCollateral", remoteBlockchainID, remoteTokenTransferrerAddress)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xb0b78b26.
//
// Solidity: function addCollateral(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) AddCollateral(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.AddCollateral(&_NativeTokenHomeUpgradeable.TransactOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xb0b78b26.
//
// Solidity: function addCollateral(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorSession) AddCollateral(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.AddCollateral(&_NativeTokenHomeUpgradeable.TransactOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address teleporterRegistryAddress, address teleporterManager, uint256 minTeleporterVersion, address wrappedTokenAddress) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactor) Initialize(opts *bind.TransactOpts, teleporterRegistryAddress common.Address, teleporterManager common.Address, minTeleporterVersion *big.Int, wrappedTokenAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.contract.Transact(opts, "initialize", teleporterRegistryAddress, teleporterManager, minTeleporterVersion, wrappedTokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address teleporterRegistryAddress, address teleporterManager, uint256 minTeleporterVersion, address wrappedTokenAddress) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) Initialize(teleporterRegistryAddress common.Address, teleporterManager common.Address, minTeleporterVersion *big.Int, wrappedTokenAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.Initialize(&_NativeTokenHomeUpgradeable.TransactOpts, teleporterRegistryAddress, teleporterManager, minTeleporterVersion, wrappedTokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address teleporterRegistryAddress, address teleporterManager, uint256 minTeleporterVersion, address wrappedTokenAddress) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorSession) Initialize(teleporterRegistryAddress common.Address, teleporterManager common.Address, minTeleporterVersion *big.Int, wrappedTokenAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.Initialize(&_NativeTokenHomeUpgradeable.TransactOpts, teleporterRegistryAddress, teleporterManager, minTeleporterVersion, wrappedTokenAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.PauseTeleporterAddress(&_NativeTokenHomeUpgradeable.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.PauseTeleporterAddress(&_NativeTokenHomeUpgradeable.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.ReceiveTeleporterMessage(&_NativeTokenHomeUpgradeable.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.ReceiveTeleporterMessage(&_NativeTokenHomeUpgradeable.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.RenounceOwnership(&_NativeTokenHomeUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.RenounceOwnership(&_NativeTokenHomeUpgradeable.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.Send(&_NativeTokenHomeUpgradeable.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.Send(&_NativeTokenHomeUpgradeable.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.contract.Transact(opts, "sendAndCall", input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.SendAndCall(&_NativeTokenHomeUpgradeable.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.SendAndCall(&_NativeTokenHomeUpgradeable.TransactOpts, input)
}

// ToCosmosSend is a paid mutator transaction binding the contract method 0x7675695e.
//
// Solidity: function toCosmosSend((bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactor) ToCosmosSend(opts *bind.TransactOpts, input ToCosmosSendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.contract.Transact(opts, "toCosmosSend", input)
}

// ToCosmosSend is a paid mutator transaction binding the contract method 0x7675695e.
//
// Solidity: function toCosmosSend((bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) ToCosmosSend(input ToCosmosSendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.ToCosmosSend(&_NativeTokenHomeUpgradeable.TransactOpts, input)
}

// ToCosmosSend is a paid mutator transaction binding the contract method 0x7675695e.
//
// Solidity: function toCosmosSend((bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorSession) ToCosmosSend(input ToCosmosSendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.ToCosmosSend(&_NativeTokenHomeUpgradeable.TransactOpts, input)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.TransferOwnership(&_NativeTokenHomeUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.TransferOwnership(&_NativeTokenHomeUpgradeable.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.UnpauseTeleporterAddress(&_NativeTokenHomeUpgradeable.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.UnpauseTeleporterAddress(&_NativeTokenHomeUpgradeable.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.UpdateMinTeleporterVersion(&_NativeTokenHomeUpgradeable.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.UpdateMinTeleporterVersion(&_NativeTokenHomeUpgradeable.TransactOpts, version)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableSession) Receive() (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.Receive(&_NativeTokenHomeUpgradeable.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableTransactorSession) Receive() (*types.Transaction, error) {
	return _NativeTokenHomeUpgradeable.Contract.Receive(&_NativeTokenHomeUpgradeable.TransactOpts)
}

// NativeTokenHomeUpgradeableCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableCallFailedIterator struct {
	Event *NativeTokenHomeUpgradeableCallFailed // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableCallFailed)
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
		it.Event = new(NativeTokenHomeUpgradeableCallFailed)
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
func (it *NativeTokenHomeUpgradeableCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableCallFailed represents a CallFailed event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenHomeUpgradeableCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableCallFailedIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableCallFailed)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseCallFailed(log types.Log) (*NativeTokenHomeUpgradeableCallFailed, error) {
	event := new(NativeTokenHomeUpgradeableCallFailed)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableCallSucceededIterator struct {
	Event *NativeTokenHomeUpgradeableCallSucceeded // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableCallSucceeded)
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
		it.Event = new(NativeTokenHomeUpgradeableCallSucceeded)
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
func (it *NativeTokenHomeUpgradeableCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableCallSucceeded represents a CallSucceeded event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenHomeUpgradeableCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableCallSucceededIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableCallSucceeded)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseCallSucceeded(log types.Log) (*NativeTokenHomeUpgradeableCallSucceeded, error) {
	event := new(NativeTokenHomeUpgradeableCallSucceeded)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableCollateralAddedIterator is returned from FilterCollateralAdded and is used to iterate over the raw logs and unpacked data for CollateralAdded events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableCollateralAddedIterator struct {
	Event *NativeTokenHomeUpgradeableCollateralAdded // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableCollateralAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableCollateralAdded)
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
		it.Event = new(NativeTokenHomeUpgradeableCollateralAdded)
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
func (it *NativeTokenHomeUpgradeableCollateralAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableCollateralAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableCollateralAdded represents a CollateralAdded event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableCollateralAdded struct {
	RemoteBlockchainID            [32]byte
	RemoteTokenTransferrerAddress common.Address
	Amount                        *big.Int
	Remaining                     *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdded is a free log retrieval operation binding the contract event 0x6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6.
//
// Solidity: event CollateralAdded(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 amount, uint256 remaining)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterCollateralAdded(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (*NativeTokenHomeUpgradeableCollateralAddedIterator, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "CollateralAdded", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableCollateralAddedIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "CollateralAdded", logs: logs, sub: sub}, nil
}

// WatchCollateralAdded is a free log subscription operation binding the contract event 0x6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6.
//
// Solidity: event CollateralAdded(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 amount, uint256 remaining)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchCollateralAdded(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableCollateralAdded, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (event.Subscription, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "CollateralAdded", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableCollateralAdded)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseCollateralAdded(log types.Log) (*NativeTokenHomeUpgradeableCollateralAdded, error) {
	event := new(NativeTokenHomeUpgradeableCollateralAdded)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableFromCosmosTokensSentIterator is returned from FilterFromCosmosTokensSent and is used to iterate over the raw logs and unpacked data for FromCosmosTokensSent events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableFromCosmosTokensSentIterator struct {
	Event *NativeTokenHomeUpgradeableFromCosmosTokensSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableFromCosmosTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableFromCosmosTokensSent)
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
		it.Event = new(NativeTokenHomeUpgradeableFromCosmosTokensSent)
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
func (it *NativeTokenHomeUpgradeableFromCosmosTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableFromCosmosTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableFromCosmosTokensSent represents a FromCosmosTokensSent event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableFromCosmosTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               FromCosmosSendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFromCosmosTokensSent is a free log retrieval operation binding the contract event 0xba11b5e7af3c06a53f098e63c9de142e163fc470b2e8565ff0de13af6010de2d.
//
// Solidity: event FromCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,bytes32,address,bytes32,bytes,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterFromCosmosTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenHomeUpgradeableFromCosmosTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "FromCosmosTokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableFromCosmosTokensSentIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "FromCosmosTokensSent", logs: logs, sub: sub}, nil
}

// WatchFromCosmosTokensSent is a free log subscription operation binding the contract event 0xba11b5e7af3c06a53f098e63c9de142e163fc470b2e8565ff0de13af6010de2d.
//
// Solidity: event FromCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,bytes32,address,bytes32,bytes,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchFromCosmosTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableFromCosmosTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "FromCosmosTokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableFromCosmosTokensSent)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "FromCosmosTokensSent", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseFromCosmosTokensSent(log types.Log) (*NativeTokenHomeUpgradeableFromCosmosTokensSent, error) {
	event := new(NativeTokenHomeUpgradeableFromCosmosTokensSent)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "FromCosmosTokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableInitializedIterator struct {
	Event *NativeTokenHomeUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableInitialized)
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
		it.Event = new(NativeTokenHomeUpgradeableInitialized)
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
func (it *NativeTokenHomeUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableInitialized represents a Initialized event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeTokenHomeUpgradeableInitializedIterator, error) {

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableInitializedIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableInitialized)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseInitialized(log types.Log) (*NativeTokenHomeUpgradeableInitialized, error) {
	event := new(NativeTokenHomeUpgradeableInitialized)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableMinTeleporterVersionUpdatedIterator struct {
	Event *NativeTokenHomeUpgradeableMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableMinTeleporterVersionUpdated)
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
		it.Event = new(NativeTokenHomeUpgradeableMinTeleporterVersionUpdated)
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
func (it *NativeTokenHomeUpgradeableMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*NativeTokenHomeUpgradeableMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableMinTeleporterVersionUpdatedIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableMinTeleporterVersionUpdated)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*NativeTokenHomeUpgradeableMinTeleporterVersionUpdated, error) {
	event := new(NativeTokenHomeUpgradeableMinTeleporterVersionUpdated)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableOwnershipTransferredIterator struct {
	Event *NativeTokenHomeUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableOwnershipTransferred)
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
		it.Event = new(NativeTokenHomeUpgradeableOwnershipTransferred)
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
func (it *NativeTokenHomeUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeTokenHomeUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableOwnershipTransferredIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableOwnershipTransferred)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenHomeUpgradeableOwnershipTransferred, error) {
	event := new(NativeTokenHomeUpgradeableOwnershipTransferred)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableRemoteRegisteredIterator is returned from FilterRemoteRegistered and is used to iterate over the raw logs and unpacked data for RemoteRegistered events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableRemoteRegisteredIterator struct {
	Event *NativeTokenHomeUpgradeableRemoteRegistered // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableRemoteRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableRemoteRegistered)
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
		it.Event = new(NativeTokenHomeUpgradeableRemoteRegistered)
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
func (it *NativeTokenHomeUpgradeableRemoteRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableRemoteRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableRemoteRegistered represents a RemoteRegistered event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableRemoteRegistered struct {
	RemoteBlockchainID            [32]byte
	RemoteTokenTransferrerAddress common.Address
	InitialCollateralNeeded       *big.Int
	TokenDecimals                 uint8
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterRemoteRegistered is a free log retrieval operation binding the contract event 0xf229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b90.
//
// Solidity: event RemoteRegistered(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 initialCollateralNeeded, uint8 tokenDecimals)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterRemoteRegistered(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (*NativeTokenHomeUpgradeableRemoteRegisteredIterator, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "RemoteRegistered", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableRemoteRegisteredIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "RemoteRegistered", logs: logs, sub: sub}, nil
}

// WatchRemoteRegistered is a free log subscription operation binding the contract event 0xf229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b90.
//
// Solidity: event RemoteRegistered(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 initialCollateralNeeded, uint8 tokenDecimals)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchRemoteRegistered(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableRemoteRegistered, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (event.Subscription, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "RemoteRegistered", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableRemoteRegistered)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "RemoteRegistered", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseRemoteRegistered(log types.Log) (*NativeTokenHomeUpgradeableRemoteRegistered, error) {
	event := new(NativeTokenHomeUpgradeableRemoteRegistered)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "RemoteRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTeleporterAddressPausedIterator struct {
	Event *NativeTokenHomeUpgradeableTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableTeleporterAddressPaused)
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
		it.Event = new(NativeTokenHomeUpgradeableTeleporterAddressPaused)
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
func (it *NativeTokenHomeUpgradeableTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenHomeUpgradeableTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableTeleporterAddressPausedIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableTeleporterAddressPaused)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseTeleporterAddressPaused(log types.Log) (*NativeTokenHomeUpgradeableTeleporterAddressPaused, error) {
	event := new(NativeTokenHomeUpgradeableTeleporterAddressPaused)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTeleporterAddressUnpausedIterator struct {
	Event *NativeTokenHomeUpgradeableTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableTeleporterAddressUnpaused)
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
		it.Event = new(NativeTokenHomeUpgradeableTeleporterAddressUnpaused)
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
func (it *NativeTokenHomeUpgradeableTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenHomeUpgradeableTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableTeleporterAddressUnpausedIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableTeleporterAddressUnpaused)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*NativeTokenHomeUpgradeableTeleporterAddressUnpaused, error) {
	event := new(NativeTokenHomeUpgradeableTeleporterAddressUnpaused)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableToCosmosTokensSentIterator is returned from FilterToCosmosTokensSent and is used to iterate over the raw logs and unpacked data for ToCosmosTokensSent events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableToCosmosTokensSentIterator struct {
	Event *NativeTokenHomeUpgradeableToCosmosTokensSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableToCosmosTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableToCosmosTokensSent)
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
		it.Event = new(NativeTokenHomeUpgradeableToCosmosTokensSent)
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
func (it *NativeTokenHomeUpgradeableToCosmosTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableToCosmosTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableToCosmosTokensSent represents a ToCosmosTokensSent event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableToCosmosTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               ToCosmosSendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterToCosmosTokensSent is a free log retrieval operation binding the contract event 0x79bd37563e6b913ff0a0e0c938210c007ebef4c942d38333576d03cb04620e9b.
//
// Solidity: event ToCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterToCosmosTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenHomeUpgradeableToCosmosTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "ToCosmosTokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableToCosmosTokensSentIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "ToCosmosTokensSent", logs: logs, sub: sub}, nil
}

// WatchToCosmosTokensSent is a free log subscription operation binding the contract event 0x79bd37563e6b913ff0a0e0c938210c007ebef4c942d38333576d03cb04620e9b.
//
// Solidity: event ToCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchToCosmosTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableToCosmosTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "ToCosmosTokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableToCosmosTokensSent)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "ToCosmosTokensSent", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseToCosmosTokensSent(log types.Log) (*NativeTokenHomeUpgradeableToCosmosTokensSent, error) {
	event := new(NativeTokenHomeUpgradeableToCosmosTokensSent)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "ToCosmosTokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableTokensAndCallRoutedIterator is returned from FilterTokensAndCallRouted and is used to iterate over the raw logs and unpacked data for TokensAndCallRouted events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTokensAndCallRoutedIterator struct {
	Event *NativeTokenHomeUpgradeableTokensAndCallRouted // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableTokensAndCallRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableTokensAndCallRouted)
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
		it.Event = new(NativeTokenHomeUpgradeableTokensAndCallRouted)
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
func (it *NativeTokenHomeUpgradeableTokensAndCallRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableTokensAndCallRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableTokensAndCallRouted represents a TokensAndCallRouted event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTokensAndCallRouted struct {
	TeleporterMessageID [32]byte
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallRouted is a free log retrieval operation binding the contract event 0x42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterTokensAndCallRouted(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenHomeUpgradeableTokensAndCallRoutedIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "TokensAndCallRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableTokensAndCallRoutedIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "TokensAndCallRouted", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallRouted is a free log subscription operation binding the contract event 0x42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchTokensAndCallRouted(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableTokensAndCallRouted, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "TokensAndCallRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableTokensAndCallRouted)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TokensAndCallRouted", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseTokensAndCallRouted(log types.Log) (*NativeTokenHomeUpgradeableTokensAndCallRouted, error) {
	event := new(NativeTokenHomeUpgradeableTokensAndCallRouted)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TokensAndCallRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTokensAndCallSentIterator struct {
	Event *NativeTokenHomeUpgradeableTokensAndCallSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableTokensAndCallSent)
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
		it.Event = new(NativeTokenHomeUpgradeableTokensAndCallSent)
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
func (it *NativeTokenHomeUpgradeableTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableTokensAndCallSent represents a TokensAndCallSent event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenHomeUpgradeableTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableTokensAndCallSentIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableTokensAndCallSent)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseTokensAndCallSent(log types.Log) (*NativeTokenHomeUpgradeableTokensAndCallSent, error) {
	event := new(NativeTokenHomeUpgradeableTokensAndCallSent)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableTokensCosmosWithdrawnIterator is returned from FilterTokensCosmosWithdrawn and is used to iterate over the raw logs and unpacked data for TokensCosmosWithdrawn events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTokensCosmosWithdrawnIterator struct {
	Event *NativeTokenHomeUpgradeableTokensCosmosWithdrawn // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableTokensCosmosWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableTokensCosmosWithdrawn)
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
		it.Event = new(NativeTokenHomeUpgradeableTokensCosmosWithdrawn)
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
func (it *NativeTokenHomeUpgradeableTokensCosmosWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableTokensCosmosWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableTokensCosmosWithdrawn represents a TokensCosmosWithdrawn event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTokensCosmosWithdrawn struct {
	MessageID [32]byte
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensCosmosWithdrawn is a free log retrieval operation binding the contract event 0xfc043975d4effd7520860981e63ea2ba7197dbe72a39dd32a3e227e820ca7ffe.
//
// Solidity: event TokensCosmosWithdrawn(bytes32 messageID, address indexed recipient, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterTokensCosmosWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenHomeUpgradeableTokensCosmosWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "TokensCosmosWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableTokensCosmosWithdrawnIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "TokensCosmosWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensCosmosWithdrawn is a free log subscription operation binding the contract event 0xfc043975d4effd7520860981e63ea2ba7197dbe72a39dd32a3e227e820ca7ffe.
//
// Solidity: event TokensCosmosWithdrawn(bytes32 messageID, address indexed recipient, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchTokensCosmosWithdrawn(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableTokensCosmosWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "TokensCosmosWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableTokensCosmosWithdrawn)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TokensCosmosWithdrawn", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseTokensCosmosWithdrawn(log types.Log) (*NativeTokenHomeUpgradeableTokensCosmosWithdrawn, error) {
	event := new(NativeTokenHomeUpgradeableTokensCosmosWithdrawn)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TokensCosmosWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableTokensRoutedIterator is returned from FilterTokensRouted and is used to iterate over the raw logs and unpacked data for TokensRouted events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTokensRoutedIterator struct {
	Event *NativeTokenHomeUpgradeableTokensRouted // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableTokensRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableTokensRouted)
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
		it.Event = new(NativeTokenHomeUpgradeableTokensRouted)
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
func (it *NativeTokenHomeUpgradeableTokensRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableTokensRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableTokensRouted represents a TokensRouted event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTokensRouted struct {
	TeleporterMessageID [32]byte
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensRouted is a free log retrieval operation binding the contract event 0x825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf0.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterTokensRouted(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenHomeUpgradeableTokensRoutedIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "TokensRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableTokensRoutedIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "TokensRouted", logs: logs, sub: sub}, nil
}

// WatchTokensRouted is a free log subscription operation binding the contract event 0x825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf0.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchTokensRouted(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableTokensRouted, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "TokensRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableTokensRouted)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TokensRouted", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseTokensRouted(log types.Log) (*NativeTokenHomeUpgradeableTokensRouted, error) {
	event := new(NativeTokenHomeUpgradeableTokensRouted)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TokensRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTokensSentIterator struct {
	Event *NativeTokenHomeUpgradeableTokensSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableTokensSent)
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
		it.Event = new(NativeTokenHomeUpgradeableTokensSent)
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
func (it *NativeTokenHomeUpgradeableTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableTokensSent represents a TokensSent event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenHomeUpgradeableTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableTokensSentIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableTokensSent)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TokensSent", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseTokensSent(log types.Log) (*NativeTokenHomeUpgradeableTokensSent, error) {
	event := new(NativeTokenHomeUpgradeableTokensSent)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeUpgradeableTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTokensWithdrawnIterator struct {
	Event *NativeTokenHomeUpgradeableTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeUpgradeableTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeUpgradeableTokensWithdrawn)
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
		it.Event = new(NativeTokenHomeUpgradeableTokensWithdrawn)
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
func (it *NativeTokenHomeUpgradeableTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeUpgradeableTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeUpgradeableTokensWithdrawn represents a TokensWithdrawn event raised by the NativeTokenHomeUpgradeable contract.
type NativeTokenHomeUpgradeableTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenHomeUpgradeableTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeUpgradeableTokensWithdrawnIterator{contract: _NativeTokenHomeUpgradeable.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeUpgradeableTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenHomeUpgradeable.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeUpgradeableTokensWithdrawn)
				if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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
func (_NativeTokenHomeUpgradeable *NativeTokenHomeUpgradeableFilterer) ParseTokensWithdrawn(log types.Log) (*NativeTokenHomeUpgradeableTokensWithdrawn, error) {
	event := new(NativeTokenHomeUpgradeableTokensWithdrawn)
	if err := _NativeTokenHomeUpgradeable.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
