// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenhome

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

// NativeTokenHomeMetaData contains all meta data concerning the NativeTokenHome contract.
var NativeTokenHomeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTeleporterVersion\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceCosmosBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sourceCosmosAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structFromCosmosSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FromCosmosTokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialCollateralNeeded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"RemoteRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"destinationCosmosRecipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCosmosBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structToCosmosSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ToCosmosTokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensCosmosWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NATIVE_TOKEN_HOME_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TELEPORTER_REGISTRY_APP_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HOME_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"}],\"name\":\"getRemoteTokenTransferrerSettings\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralNeeded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"multiplyOnRemote\",\"type\":\"bool\"}],\"internalType\":\"structRemoteTokenTransferrerSettings\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"}],\"name\":\"getTransferredBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTeleporterVersion\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"sendAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"destinationCosmosRecipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCosmosBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"internalType\":\"structToCosmosSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"toCosmosSend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051615e3a380380615e3a83398101604081905261002f916108a6565b61003b84848484610047565b5050505061090c565b50565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff1615906001600160401b03166000811580156100915750825b90506000826001600160401b031660011480156100ad5750303b155b9050811580156100bb575080155b156100d95760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b0319166001178555831561010757845460ff60401b1916680100000000000000001785555b61011389898989610164565b831561015957845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b61016c610189565b61017a8484848460126101d9565b61018381610205565b50505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff166101d757604051631afcd79f60e31b815260040160405180910390fd5b565b6101e1610189565b6101ec85858561024e565b6101f461026e565b6101fe828261027e565b5050505050565b61020d610189565b7f3b5030f10c94fcbdaa3022348ff0b82dbd4c0c71339e41ff59d0bdc92179d60080546001600160a01b0319166001600160a01b0392909216919091179055565b610256610189565b610260838261040d565b61026982610433565b505050565b610276610189565b6101d7610444565b610286610189565b6001600160a01b0382166102e15760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f20746f6b656e206164647265737300000060448201526064015b60405180910390fd5b60128160ff1611156103405760405162461bcd60e51b815260206004820152602260248201527f546f6b656e486f6d653a20746f6b656e20646563696d616c7320746f6f2068696044820152610ced60f31b60648201526084016102d8565b60007f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e60090507302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103db91906108f3565b8155600101805460ff909216600160a01b026001600160a81b03199092166001600160a01b0390931692909217179055565b610415610189565b61041d61046f565b61042561047f565b61042f8282610487565b5050565b61043b610189565b61004481610611565b60007fd2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c75005b6001905550565b610477610189565b6101d761064c565b6101d7610189565b61048f610189565b6001600160a01b03821661050b5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084016102d8565b60007fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d00905060008390506000816001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610575573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061059991906108f3565b116105ef5760405162461bcd60e51b81526020600482015260326024820152600080516020615e1a833981519152604482015271656c65706f7274657220726567697374727960701b60648201526084016102d8565b81546001600160a01b0319166001600160a01b0382161782556101838361067c565b610619610189565b6001600160a01b03811661064357604051631e4fbdf760e01b8152600060048201526024016102d8565b61004481610819565b610654610189565b60007f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00610468565b7fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d0080546040805163301fd1f560e21b815290516000926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa1580156106e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070a91906108f3565b60028301549091508184111561076a5760405162461bcd60e51b81526020600482015260316024820152600080516020615e1a83398151915260448201527032b632b837b93a32b9103b32b939b4b7b760791b60648201526084016102d8565b8084116107df5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e0060648201526084016102d8565b60028301849055604051849082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a350505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b80516001600160a01b03811681146108a157600080fd5b919050565b600080600080608085870312156108bc57600080fd5b6108c58561088a565b93506108d36020860161088a565b9250604085015191506108e86060860161088a565b905092959194509250565b60006020828403121561090557600080fd5b5051919050565b6154ff8061091b6000396000f3fe60806040526004361061012e5760003560e01c80638bf2fa94116100ab578063be2030941161006f578063be203094146103e8578063c8511ada14610408578063c868efaa146104eb578063d2cc7a701461050b578063efb5b95e1461053f578063f2fde38b1461056157600080fd5b80638bf2fa94146103335780638da5cb5b14610346578063909a6ac01461038357806397314297146103a5578063b0b78b26146103d557600080fd5b80635eb99514116100f25780635eb99514146102b657806362e3901b146102d65780636e6eef8d146102f8578063715018a61461030b5780637675695e1461032057600080fd5b806310fe9ae8146101d0578063154d625a146102265780632b0d8f18146102545780634213cf78146102745780634511243e1461029657600080fd5b366101cb577f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e601546001600160a01b031633146101c95760405162461bcd60e51b815260206004820152602f60248201527f4e6174697665546f6b656e486f6d653a20696e76616c6964207265636569766560448201526e103830bcb0b136329039b2b73232b960891b60648201526084015b60405180910390fd5b005b600080fd5b3480156101dc57600080fd5b507f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e601546001600160a01b03165b6040516001600160a01b0390911681526020015b60405180910390f35b34801561023257600080fd5b5061024661024136600461416e565b610581565b60405190815260200161021d565b34801561026057600080fd5b506101c961026f36600461419e565b6105ca565b34801561028057600080fd5b5060008051602061542a83398151915254610246565b3480156102a257600080fd5b506101c96102b136600461419e565b6106c5565b3480156102c257600080fd5b506101c96102d13660046141bb565b6107b6565b3480156102e257600080fd5b5061024660008051602061542a83398151915281565b6101c96103063660046141d4565b6107ca565b34801561031757600080fd5b506101c96107f4565b6101c961032e36600461420f565b610808565b6101c961034136600461424a565b61081a565b34801561035257600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610209565b34801561038f57600080fd5b506102466000805160206154aa83398151915281565b3480156103b157600080fd5b506103c56103c036600461419e565b610832565b604051901515815260200161021d565b6101c96103e336600461416e565b610854565b3480156103f457600080fd5b506101c9610403366004614263565b610863565b34801561041457600080fd5b506104b461042336600461416e565b604080516080808201835260008083526020808401829052838501829052606093840182905295815260008051602061546a83398151915286528381206001600160a01b039590951681529385529282902082519384018352805460ff9081161515855260018201549585019590955260028101549284019290925260039091015490921615159181019190915290565b60405161021d9190815115158152602080830151908201526040808301519082015260609182015115159181019190915260800190565b3480156104f757600080fd5b506101c96105063660046142b6565b610977565b34801561051757600080fd5b507fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d0254610246565b34801561054b57600080fd5b5061024660008051602061544a83398151915281565b34801561056d57600080fd5b506101c961057c36600461419e565b610b39565b60008281527f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e603602090815260408083206001600160a01b03851684529091529020545b92915050565b6000805160206154aa8339815191526105e1610b74565b6001600160a01b0382166106075760405162461bcd60e51b81526004016101c09061433e565b6106118183610b7c565b156106745760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b60648201526084016101c0565b6001600160a01b038216600081815260018381016020526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a25050565b6000805160206154aa8339815191526106dc610b74565b6001600160a01b0382166107025760405162461bcd60e51b81526004016101c09061433e565b61070c8183610b7c565b61076a5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472794170703a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b60648201526084016101c0565b6001600160a01b0382166000818152600183016020526040808220805460ff19169055517f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c39190a25050565b6107be610b74565b6107c781610b9e565b50565b6107c76107e360008051602061542a8339815191525490565b30336107ee85614502565b34610d3b565b6107fc610fa5565b6108066000611000565b565b6107c7610814826145d3565b34611071565b6107c761082c3683900383018361467f565b346112b4565b60006000805160206154aa83398151915261084d8184610b7c565b9392505050565b61085f828234611460565b5050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156108a85750825b90506000826001600160401b031660011480156108c45750303b155b9050811580156108d2575080155b156108f05760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561091a57845460ff60401b1916600160401b1785555b61092689898989611658565b831561096c57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b61097f611677565b60006000805160206154aa83398151915260028101548154919250906001600160a01b0316634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa1580156109ee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a12919061470b565b1015610a795760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b60648201526084016101c0565b610a838133610b7c565b15610ae95760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b60648201526084016101c0565b610b2a858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506116c192505050565b50610b33611b27565b50505050565b610b41610fa5565b6001600160a01b038116610b6b57604051631e4fbdf760e01b8152600060048201526024016101c0565b6107c781611000565b610806610fa5565b6001600160a01b03166000908152600191909101602052604090205460ff1690565b6000805160206154aa83398151915280546040805163301fd1f560e21b815290516000926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa158015610bf6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1a919061470b565b600283015490915081841115610c8c5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b60648201526084016101c0565b808411610d015760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e0060648201526084016101c0565b60028301849055604051849082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a350505050565b60008051602061548a8339815191528054600114610d6b5760405162461bcd60e51b81526004016101c090614724565b60028155610d7883611b52565b60c08301516001600160a01b031615610da35760405162461bcd60e51b81526004016101c090614768565b600080610dc58560000151866020015186886101000151896101200151611d39565b915091506000604051806040016040528060026007811115610de957610de96147ae565b81526020016040518061010001604052808c81526020018b6001600160a01b031681526020018a6001600160a01b0316815260200189604001516001600160a01b03168152602001868152602001896060015181526020018960a0015181526020018960e001516001600160a01b0316815250604051602001610e6c9190614814565b60405160208183030381529060405281525090506000610f4f6040518060c001604052808960000151815260200189602001516001600160a01b0316815260200160405180604001604052808b61010001516001600160a01b031681526020018781525081526020018960800151815260200160006001600160401b03811115610ef857610ef861438c565b604051908082528060200260200182016040528015610f21578160200160208202803683370190505b50815260200184604051602001610f3891906148b3565b604051602081830303815290604052815250611ec5565b9050876001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168987604051610f8d9291906148f8565b60405180910390a35050600190925550505050505050565b33610fd77f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146108065760405163118cdaa760e01b81523360048201526024016101c0565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b60008051602061548a83398151915280546001146110a15760405162461bcd60e51b81526004016101c090614724565b600281556110ae83611fe4565b6101008301516001600160a01b0316156110da5760405162461bcd60e51b81526004016101c090614768565b6000806110fa856000015186602001518688608001518960a0015161204f565b91509150600060405180604001604052806005600781111561111e5761111e6147ae565b8152602001604051806080016040528061113661219f565b8a516040516337c83a2f60e21b81526001600160a01b03929092169163df20e8bc916111689160040190815260200190565b602060405180830381865afa158015611185573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111a9919061470b565b81526020018960400151815260200189600001518152602001868152506040516020016111d691906149db565b604051602081830303815290604052815250905060006112616040518060c001604052808960000151815260200189602001516001600160a01b0316815260200160405180604001604052808b608001516001600160a01b031681526020018781525081526020018960e00151815260200160006001600160401b03811115610ef857610ef861438c565b9050336001600160a01b0316817f79bd37563e6b913ff0a0e0c938210c007ebef4c942d38333576d03cb04620e9b898760405161129f929190614a21565b60405180910390a35050600190925550505050565b60008051602061548a83398151915280546001146112e45760405162461bcd60e51b81526004016101c090614724565b600281556112f183612293565b60e08301516001600160a01b03161561131c5760405162461bcd60e51b81526004016101c090614768565b60008061133c856000015186602001518688606001518960800151611d39565b915091506000604051806040016040528060016007811115611360576113606147ae565b8152602001604051806040016040528089604001516001600160a01b03168152602001868152506040516020016113979190614adb565b604051602081830303815290604052815250905060006114226040518060c001604052808960000151815260200189602001516001600160a01b0316815260200160405180604001604052808b606001516001600160a01b031681526020018781525081526020018960c00151815260200160006001600160401b03811115610ef857610ef861438c565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52898760405161129f929190614afb565b60008051602061548a83398151915280546001146114905760405162461bcd60e51b81526004016101c090614724565b6002808255600085815260008051602061546a833981519152602090815260408083206001600160a01b03881684528252918290208251608081018452815460ff90811615158083526001840154948301949094529482015493810193909352600301549092161515606082015260008051602061542a8339815191529161152a5760405162461bcd60e51b81526004016101c090614b7d565b60008160200151116115885760405162461bcd60e51b815260206004820152602160248201527f546f6b656e486f6d653a207a65726f20636f6c6c61746572616c206e656564656044820152601960fa1b60648201526084016101c0565b61159184612303565b9350600080826020015186106115c2576020830151600092506115b49087614bc8565b9050826020015195506115d5565b8583602001516115d29190614bc8565b91505b600088815260028501602090815260408083206001600160a01b038b168085529083529281902060010185905580518981529182018590528a917f6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6910160405180910390a3801561164a5761164a338261232b565b505060019092555050505050565b6116606123ef565b61166e848484846012612438565b610b338161245d565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f008054600119016116bb57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b600060008051602061542a83398151915290506000828060200190518101906116ea9190614c20565b9050600181516007811115611701576117016147ae565b0361174c57600081602001518060200190518101906117209190614cae565b9050600061173387878460200151612494565b905061174382600001518261232b565b50505050505050565b600681516007811115611761576117616147ae565b036117aa57600081602001518060200190518101906117809190614ce8565b9050600061179387878460800151612522565b905061174387878460600151848660000151612620565b6002815160078111156117bf576117bf6147ae565b036118da57600081602001518060200190518101906117de9190614d8b565b905060006117f187878460800151612494565b825190915087146118575760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e486f6d653a206d69736d61746368656420736f7572636520626c6f60448201526918dad8da185a5b88125160b21b60648201526084016101c0565b856001600160a01b031682602001516001600160a01b0316146118d05760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e486f6d653a206d69736d617463686564206f726967696e2073656e60448201526a646572206164647265737360a81b60648201526084016101c0565b61174382826126f7565b6003815160078111156118ef576118ef6147ae565b036119c8576000816020015180602001905181019061190e9190614e5a565b9050600080611927888885606001518660800151612886565b915091506119be6040518061010001604052808560000151815260200185602001516001600160a01b0316815260200185604001516001600160a01b031681526020018760010160009054906101000a90046001600160a01b03166001600160a01b03168152602001838152602001600081526020018560a0015181526020018560c001516001600160a01b031681525083612935565b5050505050505050565b6004815160078111156119dd576119dd6147ae565b03611add57600081602001518060200190518101906119fc9190614ef5565b9050600080611a1688888560800151866101400151612886565b915091506119be888885600001516040518061016001604052808860200151815260200188604001516001600160a01b0316815260200188606001516001600160a01b031681526020018860a00151815260200188610100015181526020018860c0015181526020018861012001516001600160a01b031681526020018860e001516001600160a01b031681526020018a60010160009054906101000a90046001600160a01b03166001600160a01b03168152602001868152602001600081525086612ac8565b600081516007811115611af257611af26147ae565b03611b205760008160200151806020019051810190611b119190615003565b9050611b1e868683612cad565b505b5050505050565b60007f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005b6001905550565b60408101516001600160a01b0316611bbf5760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e486f6d653a207a65726f20726563697069656e7420636f6e7472616044820152696374206164647265737360b01b60648201526084016101c0565b6000816080015111611be35760405162461bcd60e51b81526004016101c090615067565b60008160a0015111611c435760405162461bcd60e51b815260206004820152602360248201527f546f6b656e486f6d653a207a65726f20726563697069656e7420676173206c696044820152621b5a5d60ea1b60648201526084016101c0565b80608001518160a0015110611ca95760405162461bcd60e51b815260206004820152602660248201527f546f6b656e486f6d653a20696e76616c696420726563697069656e7420676173604482015265081b1a5b5a5d60d21b60648201526084016101c0565b60e08101516001600160a01b0316611d165760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e486f6d653a207a65726f2066616c6c6261636b20726563697069656044820152696e74206164647265737360b01b60648201526084016101c0565b610140810151156107c75760405162461bcd60e51b81526004016101c0906150a9565b600085815260008051602061546a833981519152602090815260408083206001600160a01b038816845282528083208151608081018352815460ff908116151580835260018401549583019590955260028301549382019390935260039091015490911615156060820152829160008051602061542a8339815191529190611dd35760405162461bcd60e51b81526004016101c090614b7d565b602081015115611df55760405162461bcd60e51b81526004016101c0906150ea565b611dfe87612303565b96508415611e1557611e1286335b876130c1565b94505b6000611e2a826040015183606001518a613220565b905060008111611e7c5760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f207363616c656420616d6f756e7400000060448201526064016101c0565b60008a815260038401602090815260408083206001600160a01b038d16845290915281208054839290611eb0908490615131565b90915550909a95995094975050505050505050565b600080611ed061219f565b60408401516020015190915015611f75576040830151516001600160a01b0316611f525760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a207a65726f206665652060448201526c746f6b656e206164647265737360981b60648201526084016101c0565b604083015160208101519051611f75916001600160a01b03909116908390613237565b604051630624488560e41b81526001600160a01b03821690636244885090611fa1908690600401615144565b6020604051808303816000875af1158015611fc0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061084d919061470b565b6000816040015151116120095760405162461bcd60e51b81526004016101c0906151fd565b60008160e001511161202d5760405162461bcd60e51b81526004016101c090615067565b60c0810151156107c75760405162461bcd60e51b81526004016101c0906150a9565b600085815260008051602061546a833981519152602090815260408083206001600160a01b038816845282528083208151608081018352815460ff908116151580835260018401549583019590955260028301549382019390935260039091015490911615156060820152829160008051602061542a83398151915291906120e95760405162461bcd60e51b81526004016101c090614b7d565b60208101511561210b5760405162461bcd60e51b81526004016101c0906150ea565b61211487612303565b96508415612129576121268633611e0c565b94505b600061213e826040015183606001518a613220565b9050600081116121905760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f207363616c656420616d6f756e7400000060448201526064016101c0565b99949850939650505050505050565b6000805160206154aa83398151915280546040805163d820e64f60e01b815290516000939284926001600160a01b039091169163d820e64f916004808201926020929091908290030181865afa1580156121fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612221919061523e565b905061222d8282610b7c565b156105c45760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b60648201526084016101c0565b60408101516001600160a01b03166122bd5760405162461bcd60e51b81526004016101c0906151fd565b60008160c00151116122e15760405162461bcd60e51b81526004016101c090615067565b60a0810151156107c75760405162461bcd60e51b81526004016101c0906150a9565b60008051602061544a83398151915280546000919061084d906001600160a01b0316846132c1565b60405181815260008051602061544a833981519152906001600160a01b038416907f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b9060200160405180910390a28054604051632e1a7d4d60e01b8152600481018490526001600160a01b0390911690632e1a7d4d90602401600060405180830381600087803b1580156123be57600080fd5b505af11580156123d2573d6000803e3d6000fd5b506123ea925050506001600160a01b03841683613470565b505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661080657604051631afcd79f60e31b815260040160405180910390fd5b6124406123ef565b61244b858585613507565b612453613522565b611b208282613532565b6124656123ef565b60008051602061544a83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b600083815260008051602061546a833981519152602090815260408083206001600160a01b038616845282528083208151608081018352815460ff90811615158252600183015494820194909452600282015492810192909252600301549091161515606082015260008051602061542a833981519152906125188187878761369d565b9695505050505050565b600083815260008051602061546a833981519152602090815260408083206001600160a01b038616845282528083208151608081018352815460ff90811615158083526001840154958301959095526002830154938201939093526003909101549091161515606082015260008051602061542a833981519152916125b95760405162461bcd60e51b81526004016101c090614b7d565b60006125ce826040015183606001518761378e565b9050600081116125185760405162461bcd60e51b815260206004820152601c60248201527f546f6b656e486f6d653a207a65726f20746f6b656e20616d6f756e740000000060448201526064016101c0565b600060008051602061544a83398151915260408051848152602081018690529192506001600160a01b038616917ffc043975d4effd7520860981e63ea2ba7197dbe72a39dd32a3e227e820ca7ffe910160405180910390a28054604051632e1a7d4d60e01b8152600481018590526001600160a01b0390911690632e1a7d4d90602401600060405180830381600087803b1580156126bd57600080fd5b505af11580156126d1573d6000803e3d6000fd5b506126e9925050506001600160a01b03851684613470565b611b1e86868686600161379d565b600060008051602061544a8339815191528054604051632e1a7d4d60e01b8152600481018590529192506001600160a01b031690632e1a7d4d90602401600060405180830381600087803b15801561274e57600080fd5b505af1158015612762573d6000803e3d6000fd5b50508451602086015160408088015160a089015191516000965061278c955090919060240161525b565b60408051601f198184030181529190526020810180516001600160e01b031663161b12ff60e11b17905260c085015160608601519192506000916127d39190869085613880565b905080156128275784606001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff48560405161281a91815260200190565b60405180910390a2611b20565b84606001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb08560405161286691815260200190565b60405180910390a260e0850151611b20906001600160a01b031685613470565b600084815260008051602061546a833981519152602090815260408083206001600160a01b038716845282528083208151608081018352815460ff908116151582526001830154948201949094526002820154928101929092526003015490911615156060820152819060008051602061542a833981519152908261290d828a8a8a61369d565b90506000612924836040015184606001518961378e565b919a91995090975050505050505050565b60008051602061548a83398151915280546001146129655760405162461bcd60e51b81526004016101c090614724565b6002815561297283612293565b600061298c84600001518560200151858760800151613955565b9050806000036129aa576129a48460e001518461232b565b50612ac0565b604080518082019091526000908060018152602001604051806040016040528088604001516001600160a01b03168152602001858152506040516020016129f19190614adb565b60405160208183030381529060405281525090506000612a806040518060c001604052808860000151815260200188602001516001600160a01b0316815260200160405180604001604052808a606001516001600160a01b031681526020018a6080015181525081526020018860c00151815260200160006001600160401b03811115610ef857610ef861438c565b9050807f825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf08785604051612ab4929190614afb565b60405180910390a25050505b600190555050565b60008051602061548a8339815191528054600114612af85760405162461bcd60e51b81526004016101c090614724565b60028155612b0583611b52565b6000612b208460000151856020015185876101200151613955565b905080600003612b3e57612b388460c001518461232b565b50612ca2565b6040805180820190915260009080600281526020016040518061010001604052808b81526020018a6001600160a01b03168152602001896001600160a01b0316815260200188604001516001600160a01b03168152602001858152602001886060015181526020018860a0015181526020018860e001516001600160a01b0316815250604051602001612bd19190614814565b60405160208183030381529060405281525090506000612c626040518060c001604052808860000151815260200188602001516001600160a01b0316815260200160405180604001604052808a61010001516001600160a01b031681526020018a610120015181525081526020018860800151815260200160006001600160401b03811115610ef857610ef861438c565b9050807f42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb308785604051612c969291906148f8565b60405180910390a25050505b600190555050505050565b60008051602061542a83398151915283612d155760405162461bcd60e51b8152602060048201526024808201527f546f6b656e486f6d653a207a65726f2072656d6f746520626c6f636b636861696044820152631b88125160e21b60648201526084016101c0565b80548403612d7d5760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e486f6d653a2063616e6e6f742072656769737465722072656d6f7460448201526e329037b71039b0b6b29031b430b4b760891b60648201526084016101c0565b6001600160a01b038316612dec5760405162461bcd60e51b815260206004820152603060248201527f546f6b656e486f6d653a207a65726f2072656d6f746520746f6b656e2074726160448201526f6e73666572726572206164647265737360801b60648201526084016101c0565b600084815260028201602090815260408083206001600160a01b038716845290915290205460ff1615612e6d5760405162461bcd60e51b8152602060048201526024808201527f546f6b656e486f6d653a2072656d6f746520616c726561647920726567697374604482015263195c995960e21b60648201526084016101c0565b6012826040015160ff161115612ed75760405162461bcd60e51b815260206004820152602960248201527f546f6b656e486f6d653a2072656d6f746520746f6b656e20646563696d616c73604482015268040e8dede40d0d2ced60bb1b60648201526084016101c0565b6001810154602083015160ff908116600160a01b9092041614612f4b5760405162461bcd60e51b815260206004820152602660248201527f546f6b656e486f6d653a20696e76616c696420686f6d6520746f6b656e20646560448201526563696d616c7360d01b60648201526084016101c0565b600080612f6d8360010160149054906101000a900460ff168560400151613acf565b915091506000612f828383876000015161378e565b9050818015612f9c57508451612f999084906152a3565b15155b15612faf57612fac600182615131565b90505b60405180608001604052806001151581526020018281526020018481526020018315158152508460020160008981526020019081526020016000206000886001600160a01b03166001600160a01b0316815260200190815260200160002060008201518160000160006101000a81548160ff021916908315150217905550602082015181600101556040820151816002015560608201518160030160006101000a81548160ff021916908315150217905550905050856001600160a01b0316877ff229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b908388604001516040516130b092919091825260ff16602082015260400190565b60405180910390a350505050505050565b6040516370a0823160e01b815230600482015260009081906001600160a01b038616906370a0823190602401602060405180830381865afa15801561310a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061312e919061470b565b90506131456001600160a01b038616853086613b18565b6040516370a0823160e01b81523060048201526000906001600160a01b038716906370a0823190602401602060405180830381865afa15801561318c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131b0919061470b565b90508181116132165760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016101c0565b6125188282614bc8565b600061322f8484846001613b7f565b949350505050565b604051636eb1769f60e11b81523060048201526001600160a01b0383811660248301526000919085169063dd62ed3e90604401602060405180830381865afa158015613287573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132ab919061470b565b9050610b3384846132bc8585615131565b613ba7565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa15801561330a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061332e919061470b565b9050836001600160a01b031663d0e30db0846040518263ffffffff1660e01b81526004016000604051808303818588803b15801561336b57600080fd5b505af115801561337f573d6000803e3d6000fd5b50506040516370a0823160e01b8152306004820152600093506001600160a01b03881692506370a082319150602401602060405180830381865afa1580156133cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133ef919061470b565b905081811161345d5760405162461bcd60e51b815260206004820152603460248201527f53616665577261707065644e6174697665546f6b656e4465706f7369743a2062604482015273185b185b98d9481b9bdd081a5b98dc99585cd95960621b60648201526084016101c0565b6134678282614bc8565b95945050505050565b804710156134935760405163cd78605960e01b81523060048201526024016101c0565b6000826001600160a01b03168260405160006040518083038185875af1925050503d80600081146134e0576040519150601f19603f3d011682016040523d82523d6000602084013e6134e5565b606091505b50509050806123ea57604051630a12f52160e11b815260040160405180910390fd5b61350f6123ef565b6135198382613c37565b6123ea82613c59565b61352a6123ef565b610806613c6a565b61353a6123ef565b6001600160a01b0382166135905760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f20746f6b656e206164647265737300000060448201526064016101c0565b60128160ff1611156135ef5760405162461bcd60e51b815260206004820152602260248201527f546f6b656e486f6d653a20746f6b656e20646563696d616c7320746f6f2068696044820152610ced60f31b60648201526084016101c0565b600060008051602061542a83398151915290506005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015613647573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061366b919061470b565b8155600101805460ff909216600160a01b026001600160a81b03199092166001600160a01b0390931692909217179055565b83516000906136be5760405162461bcd60e51b81526004016101c090614b7d565b60208501511561371c5760405162461bcd60e51b8152602060048201526024808201527f546f6b656e486f6d653a2072656d6f7465206e6f7420636f6c6c61746572616c6044820152631a5e995960e21b60648201526084016101c0565b613727848484613c80565b600061373c866040015187606001518561378e565b9050600081116134675760405162461bcd60e51b815260206004820152601c60248201527f546f6b656e486f6d653a207a65726f20746f6b656e20616d6f756e740000000060448201526064016101c0565b600061322f8484846000613b7f565b6137e86040805160c081018252600080825260208083018290528351808501855282815290810191909152909182019081526020016000815260200160608152602001606081525090565b8581526001600160a01b03858116602080840191909152604080518082018252600781528151606080820184529489168082528185018981528815159285019283528451808701929092525181850152905115158186015282518082039095018552608001825280830193909352516138629291016148b3565b60408051601f1981840301815291905260a082015261174381611ec5565b6000845a10156138d25760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e7420676173000000000060448201526064016101c0565b834710156139225760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c756500000060448201526064016101c0565b826001600160a01b03163b60000361393c5750600061322f565b600080600084516020860188888bf19695505050505050565b600084815260008051602061546a833981519152602090815260408083206001600160a01b038716845282528083208151608081018352815460ff90811615801583526001840154958301959095526002830154938201939093526003909101549091161515606082015260008051602061542a83398151915291806139df575060008160200151115b156139ef5760009250505061322f565b838511613a535760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e486f6d653a20696e73756666696369656e7420616d6f756e74207460448201526b6f20636f766572206665657360a01b60648201526084016101c0565b613a5d8486614bc8565b94506000613a748260400151836060015188613220565b905080600003613a8a576000935050505061322f565b600088815260038401602090815260408083206001600160a01b038b16845290915281208054839290613abe908490615131565b909155509098975050505050505050565b60008060ff808516908416118181613af357613aeb85876152b7565b60ff16613b01565b613afd86866152b7565b60ff165b613b0c90600a6153b4565b96919550909350505050565b6040516001600160a01b038481166024830152838116604483015260648201839052610b339186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613d70565b600081151584151503613b9d57613b9685846153c0565b905061322f565b61346785846153d7565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052613bf88482613dd3565b610b33576040516001600160a01b03848116602483015260006044830152613c2d91869182169063095ea7b390606401613b4d565b610b338482613d70565b613c3f6123ef565b613c47613e76565b613c4f613e86565b61085f8282613e8e565b613c616123ef565b6107c781614018565b600060008051602061548a833981519152611b4b565b60008381527f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e603602090815260408083206001600160a01b038616845290915290205460008051602061542a8339815191529082811015613d395760405162461bcd60e51b815260206004820152602e60248201527f546f6b656e486f6d653a20696e73756666696369656e7420746f6b656e20747260448201526d616e736665722062616c616e636560901b60648201526084016101c0565b613d438382614bc8565b60009586526003909201602090815260408087206001600160a01b03909616875294905250919092205550565b6000613d856001600160a01b03841683614020565b90508051600014158015613daa575080806020019051810190613da891906153eb565b155b156123ea57604051635274afe760e01b81526001600160a01b03841660048201526024016101c0565b6000806000846001600160a01b031684604051613df0919061540d565b6000604051808303816000865af19150503d8060008114613e2d576040519150601f19603f3d011682016040523d82523d6000602084013e613e32565b606091505b5091509150818015613e5c575080511580613e5c575080806020019051810190613e5c91906153eb565b80156134675750505050506001600160a01b03163b151590565b613e7e6123ef565b61080661402e565b6108066123ef565b613e966123ef565b6001600160a01b038216613f125760405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084016101c0565b60006000805160206154aa833981519152905060008390506000816001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015613f6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613f8e919061470b565b11613ff65760405162461bcd60e51b815260206004820152603260248201527f54656c65706f7274657252656769737472794170703a20696e76616c69642054604482015271656c65706f7274657220726567697374727960701b60648201526084016101c0565b81546001600160a01b0319166001600160a01b038216178255610b3383610b9e565b610b416123ef565b606061084d83836000614036565b611b276123ef565b60608147101561405b5760405163cd78605960e01b81523060048201526024016101c0565b600080856001600160a01b03168486604051614077919061540d565b60006040518083038185875af1925050503d80600081146140b4576040519150601f19603f3d011682016040523d82523d6000602084013e6140b9565b606091505b50915091506125188683836060826140d9576140d482614120565b61084d565b81511580156140f057506001600160a01b0384163b155b1561411957604051639996b31560e01b81526001600160a01b03851660048201526024016101c0565b508061084d565b8051156141305780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b6001600160a01b03811681146107c757600080fd5b803561416981614149565b919050565b6000806040838503121561418157600080fd5b82359150602083013561419381614149565b809150509250929050565b6000602082840312156141b057600080fd5b813561084d81614149565b6000602082840312156141cd57600080fd5b5035919050565b6000602082840312156141e657600080fd5b81356001600160401b038111156141fc57600080fd5b8201610160818503121561084d57600080fd5b60006020828403121561422157600080fd5b81356001600160401b0381111561423757600080fd5b8201610120818503121561084d57600080fd5b6000610100828403121561425d57600080fd5b50919050565b6000806000806080858703121561427957600080fd5b843561428481614149565b9350602085013561429481614149565b92506040850135915060608501356142ab81614149565b939692955090935050565b600080600080606085870312156142cc57600080fd5b8435935060208501356142de81614149565b925060408501356001600160401b03808211156142fa57600080fd5b818701915087601f83011261430e57600080fd5b81358181111561431d57600080fd5b88602082850101111561432f57600080fd5b95989497505060200194505050565b6020808252602e908201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b60405161016081016001600160401b03811182821017156143c5576143c561438c565b60405290565b60405161012081016001600160401b03811182821017156143c5576143c561438c565b60405161010081016001600160401b03811182821017156143c5576143c561438c565b604080519081016001600160401b03811182821017156143c5576143c561438c565b60405160a081016001600160401b03811182821017156143c5576143c561438c565b604051601f8201601f191681016001600160401b038111828210171561447d5761447d61438c565b604052919050565b60006001600160401b0382111561449e5761449e61438c565b50601f01601f191660200190565b600082601f8301126144bd57600080fd5b81356144d06144cb82614485565b614455565b8181528460208386010111156144e557600080fd5b816020850160208301376000918101602001919091529392505050565b6000610160823603121561451557600080fd5b61451d6143a2565b8235815261452d6020840161415e565b602082015261453e6040840161415e565b604082015260608301356001600160401b0381111561455c57600080fd5b614568368286016144ac565b6060830152506080830135608082015260a083013560a082015261458e60c0840161415e565b60c082015261459f60e0840161415e565b60e08201526101006145b281850161415e565b90820152610120838101359082015261014092830135928101929092525090565b600061012082360312156145e657600080fd5b6145ee6143cb565b823581526145fe6020840161415e565b602082015260408301356001600160401b0381111561461c57600080fd5b614628368286016144ac565b604083015250606083013560608201526146446080840161415e565b608082015260a083013560a082015260c083013560c082015260e083013560e082015261010061467581850161415e565b9082015292915050565b6000610100828403121561469257600080fd5b61469a6143ee565b8235815260208301356146ac81614149565b602082015260408301356146bf81614149565b60408201526146d06060840161415e565b60608201526080830135608082015260a083013560a082015260c083013560c08201526146ff60e0840161415e565b60e08201529392505050565b60006020828403121561471d57600080fd5b5051919050565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b60208082526026908201527f546f6b656e486f6d653a206e6f6e2d7a65726f206d756c74692d686f702066616040820152656c6c6261636b60d01b606082015260800190565b634e487b7160e01b600052602160045260246000fd5b60005b838110156147df5781810151838201526020016147c7565b50506000910152565b600081518084526148008160208601602086016147c4565b601f01601f19169290920160200192915050565b60208152815160208201526000602083015160018060a01b0380821660408501528060408601511660608501525050606083015161485d60808401826001600160a01b03169052565b50608083015160a083015260a08301516101008060c08501526148846101208501836147e8565b915060c085015160e085015260e08501516148a9828601826001600160a01b03169052565b5090949350505050565b6020815260008251600881106148d957634e487b7160e01b600052602160045260246000fd5b80602084015250602083015160408084015261322f60608401826147e8565b60408152825160408201526000602084015161491f60608401826001600160a01b03169052565b5060408401516001600160a01b03166080830152606084015161016060a084018190526149506101a08501836147e8565b9150608086015160c085015260a086015160e085015260c0860151610100614982818701836001600160a01b03169052565b60e088015191506101206149a0818801846001600160a01b03169052565b908801519150610140906149be878301846001600160a01b03169052565b880151928601929092525090940151610180830152506020015290565b60208152815160208201526000602083015160806040840152614a0160a08401826147e8565b905060408401516060840152606084015160808401528091505092915050565b604081528251604082015260006020840151614a4860608401826001600160a01b03169052565b506040840151610120806080850152614a656101608501836147e8565b9150606086015160a08501526080860151614a8b60c08601826001600160a01b03169052565b5060a086015160e085015260c0860151610100818187015260e0880151838701528088015192505050614aca6101408501826001600160a01b03169052565b506020929092019290925292915050565b81516001600160a01b0316815260208083015190820152604081016105c4565b60006101208201905083518252602084015160018060a01b03808216602085015280604087015116604085015280606087015116606085015250506080840151608083015260a084015160a083015260c084015160c083015260e0840151614b6e60e08401826001600160a01b03169052565b50826101008301529392505050565b6020808252818101527f546f6b656e486f6d653a2072656d6f7465206e6f742072656769737465726564604082015260600190565b634e487b7160e01b600052601160045260246000fd5b818103818111156105c4576105c4614bb2565b600082601f830112614bec57600080fd5b8151614bfa6144cb82614485565b818152846020838601011115614c0f57600080fd5b61322f8260208301602087016147c4565b600060208284031215614c3257600080fd5b81516001600160401b0380821115614c4957600080fd5b9083019060408286031215614c5d57600080fd5b614c65614411565b825160088110614c7457600080fd5b8152602083015182811115614c8857600080fd5b614c9487828601614bdb565b60208301525095945050505050565b805161416981614149565b600060408284031215614cc057600080fd5b614cc8614411565b8251614cd381614149565b81526020928301519281019290925250919050565b600060208284031215614cfa57600080fd5b81516001600160401b0380821115614d1157600080fd5b9083019060a08286031215614d2557600080fd5b614d2d614433565b82518152602083015182811115614d4357600080fd5b614d4f87828601614bdb565b6020830152506040830151604082015260608301519150614d6f82614149565b8160608201526080830151608082015280935050505092915050565b600060208284031215614d9d57600080fd5b81516001600160401b0380821115614db457600080fd5b908301906101008286031215614dc957600080fd5b614dd16143ee565b82518152614de160208401614ca3565b6020820152614df260408401614ca3565b6040820152614e0360608401614ca3565b60608201526080830151608082015260a083015182811115614e2457600080fd5b614e3087828601614bdb565b60a08301525060c083015160c0820152614e4c60e08401614ca3565b60e082015295945050505050565b600060e08284031215614e6c57600080fd5b60405160e081018181106001600160401b0382111715614e8e57614e8e61438c565b604052825181526020830151614ea381614149565b60208201526040830151614eb681614149565b80604083015250606083015160608201526080830151608082015260a083015160a082015260c0830151614ee981614149565b60c08201529392505050565b600060208284031215614f0757600080fd5b81516001600160401b0380821115614f1e57600080fd5b908301906101608286031215614f3357600080fd5b614f3b6143a2565b614f4483614ca3565b815260208301516020820152614f5c60408401614ca3565b6040820152614f6d60608401614ca3565b60608201526080830151608082015260a083015182811115614f8e57600080fd5b614f9a87828601614bdb565b60a08301525060c083015160c0820152614fb660e08401614ca3565b60e082015261010083810151908201526101209150614fd6828401614ca3565b9181019190915261014091820151918101919091529392505050565b805160ff8116811461416957600080fd5b60006060828403121561501557600080fd5b604051606081018181106001600160401b03821117156150375761503761438c565b6040528251815261504a60208401614ff2565b602082015261505b60408401614ff2565b60408201529392505050565b60208082526022908201527f546f6b656e486f6d653a207a65726f20726571756972656420676173206c696d6040820152611a5d60f21b606082015260800190565b60208082526021908201527f546f6b656e486f6d653a206e6f6e2d7a65726f207365636f6e646172792066656040820152606560f81b606082015260800190565b60208082526027908201527f546f6b656e486f6d653a20636f6c6c61746572616c206e656564656420666f726040820152662072656d6f746560c81b606082015260800190565b808201808211156105c4576105c4614bb2565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501526000929161010085019190606087015160a0870152608087015160e060c0880152805193849052840192600092506101208701905b808410156151d3578451831682529385019360019390930192908501906151b1565b5060a0880151878203601f190160e089015294506151f181866147e8565b98975050505050505050565b60208082526021908201527f546f6b656e486f6d653a207a65726f20726563697069656e74206164647265736040820152607360f81b606082015260800190565b60006020828403121561525057600080fd5b815161084d81614149565b8481526001600160a01b03848116602083015283166040820152608060608201819052600090612518908301846147e8565b634e487b7160e01b600052601260045260246000fd5b6000826152b2576152b261528d565b500690565b60ff82811682821603908111156105c4576105c4614bb2565b600181815b8085111561530b5781600019048211156152f1576152f1614bb2565b808516156152fe57918102915b93841c93908002906152d5565b509250929050565b600082615322575060016105c4565b8161532f575060006105c4565b8160018114615345576002811461534f5761536b565b60019150506105c4565b60ff84111561536057615360614bb2565b50506001821b6105c4565b5060208310610133831016604e8410600b841016171561538e575081810a6105c4565b61539883836152d0565b80600019048211156153ac576153ac614bb2565b029392505050565b600061084d8383615313565b80820281158282048414176105c4576105c4614bb2565b6000826153e6576153e661528d565b500490565b6000602082840312156153fd57600080fd5b8151801515811461084d57600080fd5b6000825161541f8184602087016147c4565b919091019291505056fe9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e6003b5030f10c94fcbdaa3022348ff0b82dbd4c0c71339e41ff59d0bdc92179d6009316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e602d2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c7500de77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d00a26469706673582212207e49b1abfec7321c4f3e07df00d49d2e3753c48f2d354f20c586468cbd90963764736f6c6343000819003354656c65706f7274657252656769737472794170703a20696e76616c69642054",
}

// NativeTokenHomeABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenHomeMetaData.ABI instead.
var NativeTokenHomeABI = NativeTokenHomeMetaData.ABI

// NativeTokenHomeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenHomeMetaData.Bin instead.
var NativeTokenHomeBin = NativeTokenHomeMetaData.Bin

// DeployNativeTokenHome deploys a new Ethereum contract, binding an instance of NativeTokenHome to it.
func DeployNativeTokenHome(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterRegistryAddress common.Address, teleporterManager common.Address, minTeleporterVersion *big.Int, wrappedTokenAddress common.Address) (common.Address, *types.Transaction, *NativeTokenHome, error) {
	parsed, err := NativeTokenHomeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenHomeBin), backend, teleporterRegistryAddress, teleporterManager, minTeleporterVersion, wrappedTokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenHome{NativeTokenHomeCaller: NativeTokenHomeCaller{contract: contract}, NativeTokenHomeTransactor: NativeTokenHomeTransactor{contract: contract}, NativeTokenHomeFilterer: NativeTokenHomeFilterer{contract: contract}}, nil
}

// NativeTokenHome is an auto generated Go binding around an Ethereum contract.
type NativeTokenHome struct {
	NativeTokenHomeCaller     // Read-only binding to the contract
	NativeTokenHomeTransactor // Write-only binding to the contract
	NativeTokenHomeFilterer   // Log filterer for contract events
}

// NativeTokenHomeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenHomeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenHomeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenHomeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenHomeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenHomeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenHomeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenHomeSession struct {
	Contract     *NativeTokenHome  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeTokenHomeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenHomeCallerSession struct {
	Contract *NativeTokenHomeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NativeTokenHomeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenHomeTransactorSession struct {
	Contract     *NativeTokenHomeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NativeTokenHomeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenHomeRaw struct {
	Contract *NativeTokenHome // Generic contract binding to access the raw methods on
}

// NativeTokenHomeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenHomeCallerRaw struct {
	Contract *NativeTokenHomeCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenHomeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenHomeTransactorRaw struct {
	Contract *NativeTokenHomeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenHome creates a new instance of NativeTokenHome, bound to a specific deployed contract.
func NewNativeTokenHome(address common.Address, backend bind.ContractBackend) (*NativeTokenHome, error) {
	contract, err := bindNativeTokenHome(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHome{NativeTokenHomeCaller: NativeTokenHomeCaller{contract: contract}, NativeTokenHomeTransactor: NativeTokenHomeTransactor{contract: contract}, NativeTokenHomeFilterer: NativeTokenHomeFilterer{contract: contract}}, nil
}

// NewNativeTokenHomeCaller creates a new read-only instance of NativeTokenHome, bound to a specific deployed contract.
func NewNativeTokenHomeCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenHomeCaller, error) {
	contract, err := bindNativeTokenHome(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeCaller{contract: contract}, nil
}

// NewNativeTokenHomeTransactor creates a new write-only instance of NativeTokenHome, bound to a specific deployed contract.
func NewNativeTokenHomeTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenHomeTransactor, error) {
	contract, err := bindNativeTokenHome(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTransactor{contract: contract}, nil
}

// NewNativeTokenHomeFilterer creates a new log filterer instance of NativeTokenHome, bound to a specific deployed contract.
func NewNativeTokenHomeFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenHomeFilterer, error) {
	contract, err := bindNativeTokenHome(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeFilterer{contract: contract}, nil
}

// bindNativeTokenHome binds a generic wrapper to an already deployed contract.
func bindNativeTokenHome(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenHomeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenHome *NativeTokenHomeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenHome.Contract.NativeTokenHomeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenHome *NativeTokenHomeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.NativeTokenHomeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenHome *NativeTokenHomeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.NativeTokenHomeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenHome *NativeTokenHomeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenHome.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenHome *NativeTokenHomeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenHome *NativeTokenHomeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.contract.Transact(opts, method, params...)
}

// NATIVETOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0xefb5b95e.
//
// Solidity: function NATIVE_TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCaller) NATIVETOKENHOMESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "NATIVE_TOKEN_HOME_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NATIVETOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0xefb5b95e.
//
// Solidity: function NATIVE_TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeSession) NATIVETOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHome.Contract.NATIVETOKENHOMESTORAGELOCATION(&_NativeTokenHome.CallOpts)
}

// NATIVETOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0xefb5b95e.
//
// Solidity: function NATIVE_TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCallerSession) NATIVETOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHome.Contract.NATIVETOKENHOMESTORAGELOCATION(&_NativeTokenHome.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCaller) TELEPORTERREGISTRYAPPSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "TELEPORTER_REGISTRY_APP_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHome.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_NativeTokenHome.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCallerSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHome.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_NativeTokenHome.CallOpts)
}

// TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x62e3901b.
//
// Solidity: function TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCaller) TOKENHOMESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "TOKEN_HOME_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x62e3901b.
//
// Solidity: function TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeSession) TOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHome.Contract.TOKENHOMESTORAGELOCATION(&_NativeTokenHome.CallOpts)
}

// TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x62e3901b.
//
// Solidity: function TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCallerSession) TOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHome.Contract.TOKENHOMESTORAGELOCATION(&_NativeTokenHome.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCaller) GetBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "getBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeSession) GetBlockchainID() ([32]byte, error) {
	return _NativeTokenHome.Contract.GetBlockchainID(&_NativeTokenHome.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCallerSession) GetBlockchainID() ([32]byte, error) {
	return _NativeTokenHome.Contract.GetBlockchainID(&_NativeTokenHome.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenHome *NativeTokenHomeCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenHome *NativeTokenHomeSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenHome.Contract.GetMinTeleporterVersion(&_NativeTokenHome.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenHome *NativeTokenHomeCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenHome.Contract.GetMinTeleporterVersion(&_NativeTokenHome.CallOpts)
}

// GetRemoteTokenTransferrerSettings is a free data retrieval call binding the contract method 0xc8511ada.
//
// Solidity: function getRemoteTokenTransferrerSettings(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns((bool,uint256,uint256,bool))
func (_NativeTokenHome *NativeTokenHomeCaller) GetRemoteTokenTransferrerSettings(opts *bind.CallOpts, remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (RemoteTokenTransferrerSettings, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "getRemoteTokenTransferrerSettings", remoteBlockchainID, remoteTokenTransferrerAddress)

	if err != nil {
		return *new(RemoteTokenTransferrerSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(RemoteTokenTransferrerSettings)).(*RemoteTokenTransferrerSettings)

	return out0, err

}

// GetRemoteTokenTransferrerSettings is a free data retrieval call binding the contract method 0xc8511ada.
//
// Solidity: function getRemoteTokenTransferrerSettings(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns((bool,uint256,uint256,bool))
func (_NativeTokenHome *NativeTokenHomeSession) GetRemoteTokenTransferrerSettings(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (RemoteTokenTransferrerSettings, error) {
	return _NativeTokenHome.Contract.GetRemoteTokenTransferrerSettings(&_NativeTokenHome.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// GetRemoteTokenTransferrerSettings is a free data retrieval call binding the contract method 0xc8511ada.
//
// Solidity: function getRemoteTokenTransferrerSettings(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns((bool,uint256,uint256,bool))
func (_NativeTokenHome *NativeTokenHomeCallerSession) GetRemoteTokenTransferrerSettings(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (RemoteTokenTransferrerSettings, error) {
	return _NativeTokenHome.Contract.GetRemoteTokenTransferrerSettings(&_NativeTokenHome.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_NativeTokenHome *NativeTokenHomeCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_NativeTokenHome *NativeTokenHomeSession) GetTokenAddress() (common.Address, error) {
	return _NativeTokenHome.Contract.GetTokenAddress(&_NativeTokenHome.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_NativeTokenHome *NativeTokenHomeCallerSession) GetTokenAddress() (common.Address, error) {
	return _NativeTokenHome.Contract.GetTokenAddress(&_NativeTokenHome.CallOpts)
}

// GetTransferredBalance is a free data retrieval call binding the contract method 0x154d625a.
//
// Solidity: function getTransferredBalance(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns(uint256)
func (_NativeTokenHome *NativeTokenHomeCaller) GetTransferredBalance(opts *bind.CallOpts, remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "getTransferredBalance", remoteBlockchainID, remoteTokenTransferrerAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransferredBalance is a free data retrieval call binding the contract method 0x154d625a.
//
// Solidity: function getTransferredBalance(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns(uint256)
func (_NativeTokenHome *NativeTokenHomeSession) GetTransferredBalance(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*big.Int, error) {
	return _NativeTokenHome.Contract.GetTransferredBalance(&_NativeTokenHome.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// GetTransferredBalance is a free data retrieval call binding the contract method 0x154d625a.
//
// Solidity: function getTransferredBalance(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns(uint256)
func (_NativeTokenHome *NativeTokenHomeCallerSession) GetTransferredBalance(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*big.Int, error) {
	return _NativeTokenHome.Contract.GetTransferredBalance(&_NativeTokenHome.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenHome *NativeTokenHomeCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenHome *NativeTokenHomeSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenHome.Contract.IsTeleporterAddressPaused(&_NativeTokenHome.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenHome *NativeTokenHomeCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenHome.Contract.IsTeleporterAddressPaused(&_NativeTokenHome.CallOpts, teleporterAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenHome *NativeTokenHomeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenHome *NativeTokenHomeSession) Owner() (common.Address, error) {
	return _NativeTokenHome.Contract.Owner(&_NativeTokenHome.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenHome *NativeTokenHomeCallerSession) Owner() (common.Address, error) {
	return _NativeTokenHome.Contract.Owner(&_NativeTokenHome.CallOpts)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xb0b78b26.
//
// Solidity: function addCollateral(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) AddCollateral(opts *bind.TransactOpts, remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "addCollateral", remoteBlockchainID, remoteTokenTransferrerAddress)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xb0b78b26.
//
// Solidity: function addCollateral(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) payable returns()
func (_NativeTokenHome *NativeTokenHomeSession) AddCollateral(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.AddCollateral(&_NativeTokenHome.TransactOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xb0b78b26.
//
// Solidity: function addCollateral(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) AddCollateral(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.AddCollateral(&_NativeTokenHome.TransactOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address teleporterRegistryAddress, address teleporterManager, uint256 minTeleporterVersion, address wrappedTokenAddress) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) Initialize(opts *bind.TransactOpts, teleporterRegistryAddress common.Address, teleporterManager common.Address, minTeleporterVersion *big.Int, wrappedTokenAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "initialize", teleporterRegistryAddress, teleporterManager, minTeleporterVersion, wrappedTokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address teleporterRegistryAddress, address teleporterManager, uint256 minTeleporterVersion, address wrappedTokenAddress) returns()
func (_NativeTokenHome *NativeTokenHomeSession) Initialize(teleporterRegistryAddress common.Address, teleporterManager common.Address, minTeleporterVersion *big.Int, wrappedTokenAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.Initialize(&_NativeTokenHome.TransactOpts, teleporterRegistryAddress, teleporterManager, minTeleporterVersion, wrappedTokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address teleporterRegistryAddress, address teleporterManager, uint256 minTeleporterVersion, address wrappedTokenAddress) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) Initialize(teleporterRegistryAddress common.Address, teleporterManager common.Address, minTeleporterVersion *big.Int, wrappedTokenAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.Initialize(&_NativeTokenHome.TransactOpts, teleporterRegistryAddress, teleporterManager, minTeleporterVersion, wrappedTokenAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenHome *NativeTokenHomeSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.PauseTeleporterAddress(&_NativeTokenHome.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.PauseTeleporterAddress(&_NativeTokenHome.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenHome *NativeTokenHomeSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.ReceiveTeleporterMessage(&_NativeTokenHome.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.ReceiveTeleporterMessage(&_NativeTokenHome.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenHome *NativeTokenHomeSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenHome.Contract.RenounceOwnership(&_NativeTokenHome.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenHome.Contract.RenounceOwnership(&_NativeTokenHome.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.Send(&_NativeTokenHome.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.Send(&_NativeTokenHome.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "sendAndCall", input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.SendAndCall(&_NativeTokenHome.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.SendAndCall(&_NativeTokenHome.TransactOpts, input)
}

// ToCosmosSend is a paid mutator transaction binding the contract method 0x7675695e.
//
// Solidity: function toCosmosSend((bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) ToCosmosSend(opts *bind.TransactOpts, input ToCosmosSendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "toCosmosSend", input)
}

// ToCosmosSend is a paid mutator transaction binding the contract method 0x7675695e.
//
// Solidity: function toCosmosSend((bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeSession) ToCosmosSend(input ToCosmosSendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.ToCosmosSend(&_NativeTokenHome.TransactOpts, input)
}

// ToCosmosSend is a paid mutator transaction binding the contract method 0x7675695e.
//
// Solidity: function toCosmosSend((bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) ToCosmosSend(input ToCosmosSendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.ToCosmosSend(&_NativeTokenHome.TransactOpts, input)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenHome *NativeTokenHomeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.TransferOwnership(&_NativeTokenHome.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.TransferOwnership(&_NativeTokenHome.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenHome *NativeTokenHomeSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.UnpauseTeleporterAddress(&_NativeTokenHome.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.UnpauseTeleporterAddress(&_NativeTokenHome.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenHome *NativeTokenHomeSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.UpdateMinTeleporterVersion(&_NativeTokenHome.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.UpdateMinTeleporterVersion(&_NativeTokenHome.TransactOpts, version)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenHome.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenHome *NativeTokenHomeSession) Receive() (*types.Transaction, error) {
	return _NativeTokenHome.Contract.Receive(&_NativeTokenHome.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) Receive() (*types.Transaction, error) {
	return _NativeTokenHome.Contract.Receive(&_NativeTokenHome.TransactOpts)
}

// NativeTokenHomeCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the NativeTokenHome contract.
type NativeTokenHomeCallFailedIterator struct {
	Event *NativeTokenHomeCallFailed // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeCallFailed)
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
		it.Event = new(NativeTokenHomeCallFailed)
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
func (it *NativeTokenHomeCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeCallFailed represents a CallFailed event raised by the NativeTokenHome contract.
type NativeTokenHomeCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenHomeCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeCallFailedIterator{contract: _NativeTokenHome.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeCallFailed)
				if err := _NativeTokenHome.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseCallFailed(log types.Log) (*NativeTokenHomeCallFailed, error) {
	event := new(NativeTokenHomeCallFailed)
	if err := _NativeTokenHome.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the NativeTokenHome contract.
type NativeTokenHomeCallSucceededIterator struct {
	Event *NativeTokenHomeCallSucceeded // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeCallSucceeded)
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
		it.Event = new(NativeTokenHomeCallSucceeded)
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
func (it *NativeTokenHomeCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeCallSucceeded represents a CallSucceeded event raised by the NativeTokenHome contract.
type NativeTokenHomeCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenHomeCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeCallSucceededIterator{contract: _NativeTokenHome.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeCallSucceeded)
				if err := _NativeTokenHome.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseCallSucceeded(log types.Log) (*NativeTokenHomeCallSucceeded, error) {
	event := new(NativeTokenHomeCallSucceeded)
	if err := _NativeTokenHome.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeCollateralAddedIterator is returned from FilterCollateralAdded and is used to iterate over the raw logs and unpacked data for CollateralAdded events raised by the NativeTokenHome contract.
type NativeTokenHomeCollateralAddedIterator struct {
	Event *NativeTokenHomeCollateralAdded // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeCollateralAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeCollateralAdded)
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
		it.Event = new(NativeTokenHomeCollateralAdded)
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
func (it *NativeTokenHomeCollateralAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeCollateralAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeCollateralAdded represents a CollateralAdded event raised by the NativeTokenHome contract.
type NativeTokenHomeCollateralAdded struct {
	RemoteBlockchainID            [32]byte
	RemoteTokenTransferrerAddress common.Address
	Amount                        *big.Int
	Remaining                     *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdded is a free log retrieval operation binding the contract event 0x6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6.
//
// Solidity: event CollateralAdded(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 amount, uint256 remaining)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterCollateralAdded(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (*NativeTokenHomeCollateralAddedIterator, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "CollateralAdded", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeCollateralAddedIterator{contract: _NativeTokenHome.contract, event: "CollateralAdded", logs: logs, sub: sub}, nil
}

// WatchCollateralAdded is a free log subscription operation binding the contract event 0x6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6.
//
// Solidity: event CollateralAdded(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 amount, uint256 remaining)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchCollateralAdded(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeCollateralAdded, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (event.Subscription, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "CollateralAdded", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeCollateralAdded)
				if err := _NativeTokenHome.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseCollateralAdded(log types.Log) (*NativeTokenHomeCollateralAdded, error) {
	event := new(NativeTokenHomeCollateralAdded)
	if err := _NativeTokenHome.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeFromCosmosTokensSentIterator is returned from FilterFromCosmosTokensSent and is used to iterate over the raw logs and unpacked data for FromCosmosTokensSent events raised by the NativeTokenHome contract.
type NativeTokenHomeFromCosmosTokensSentIterator struct {
	Event *NativeTokenHomeFromCosmosTokensSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeFromCosmosTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeFromCosmosTokensSent)
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
		it.Event = new(NativeTokenHomeFromCosmosTokensSent)
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
func (it *NativeTokenHomeFromCosmosTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeFromCosmosTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeFromCosmosTokensSent represents a FromCosmosTokensSent event raised by the NativeTokenHome contract.
type NativeTokenHomeFromCosmosTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               FromCosmosSendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFromCosmosTokensSent is a free log retrieval operation binding the contract event 0xba11b5e7af3c06a53f098e63c9de142e163fc470b2e8565ff0de13af6010de2d.
//
// Solidity: event FromCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,bytes32,address,bytes32,bytes,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterFromCosmosTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenHomeFromCosmosTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "FromCosmosTokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeFromCosmosTokensSentIterator{contract: _NativeTokenHome.contract, event: "FromCosmosTokensSent", logs: logs, sub: sub}, nil
}

// WatchFromCosmosTokensSent is a free log subscription operation binding the contract event 0xba11b5e7af3c06a53f098e63c9de142e163fc470b2e8565ff0de13af6010de2d.
//
// Solidity: event FromCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,bytes32,address,bytes32,bytes,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchFromCosmosTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeFromCosmosTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "FromCosmosTokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeFromCosmosTokensSent)
				if err := _NativeTokenHome.contract.UnpackLog(event, "FromCosmosTokensSent", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseFromCosmosTokensSent(log types.Log) (*NativeTokenHomeFromCosmosTokensSent, error) {
	event := new(NativeTokenHomeFromCosmosTokensSent)
	if err := _NativeTokenHome.contract.UnpackLog(event, "FromCosmosTokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeTokenHome contract.
type NativeTokenHomeInitializedIterator struct {
	Event *NativeTokenHomeInitialized // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeInitialized)
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
		it.Event = new(NativeTokenHomeInitialized)
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
func (it *NativeTokenHomeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeInitialized represents a Initialized event raised by the NativeTokenHome contract.
type NativeTokenHomeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeTokenHomeInitializedIterator, error) {

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeInitializedIterator{contract: _NativeTokenHome.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeInitialized)
				if err := _NativeTokenHome.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseInitialized(log types.Log) (*NativeTokenHomeInitialized, error) {
	event := new(NativeTokenHomeInitialized)
	if err := _NativeTokenHome.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the NativeTokenHome contract.
type NativeTokenHomeMinTeleporterVersionUpdatedIterator struct {
	Event *NativeTokenHomeMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeMinTeleporterVersionUpdated)
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
		it.Event = new(NativeTokenHomeMinTeleporterVersionUpdated)
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
func (it *NativeTokenHomeMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the NativeTokenHome contract.
type NativeTokenHomeMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*NativeTokenHomeMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeMinTeleporterVersionUpdatedIterator{contract: _NativeTokenHome.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeMinTeleporterVersionUpdated)
				if err := _NativeTokenHome.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*NativeTokenHomeMinTeleporterVersionUpdated, error) {
	event := new(NativeTokenHomeMinTeleporterVersionUpdated)
	if err := _NativeTokenHome.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeTokenHome contract.
type NativeTokenHomeOwnershipTransferredIterator struct {
	Event *NativeTokenHomeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeOwnershipTransferred)
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
		it.Event = new(NativeTokenHomeOwnershipTransferred)
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
func (it *NativeTokenHomeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeOwnershipTransferred represents a OwnershipTransferred event raised by the NativeTokenHome contract.
type NativeTokenHomeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeTokenHomeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeOwnershipTransferredIterator{contract: _NativeTokenHome.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeOwnershipTransferred)
				if err := _NativeTokenHome.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenHomeOwnershipTransferred, error) {
	event := new(NativeTokenHomeOwnershipTransferred)
	if err := _NativeTokenHome.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeRemoteRegisteredIterator is returned from FilterRemoteRegistered and is used to iterate over the raw logs and unpacked data for RemoteRegistered events raised by the NativeTokenHome contract.
type NativeTokenHomeRemoteRegisteredIterator struct {
	Event *NativeTokenHomeRemoteRegistered // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeRemoteRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeRemoteRegistered)
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
		it.Event = new(NativeTokenHomeRemoteRegistered)
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
func (it *NativeTokenHomeRemoteRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeRemoteRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeRemoteRegistered represents a RemoteRegistered event raised by the NativeTokenHome contract.
type NativeTokenHomeRemoteRegistered struct {
	RemoteBlockchainID            [32]byte
	RemoteTokenTransferrerAddress common.Address
	InitialCollateralNeeded       *big.Int
	TokenDecimals                 uint8
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterRemoteRegistered is a free log retrieval operation binding the contract event 0xf229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b90.
//
// Solidity: event RemoteRegistered(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 initialCollateralNeeded, uint8 tokenDecimals)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterRemoteRegistered(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (*NativeTokenHomeRemoteRegisteredIterator, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "RemoteRegistered", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeRemoteRegisteredIterator{contract: _NativeTokenHome.contract, event: "RemoteRegistered", logs: logs, sub: sub}, nil
}

// WatchRemoteRegistered is a free log subscription operation binding the contract event 0xf229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b90.
//
// Solidity: event RemoteRegistered(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 initialCollateralNeeded, uint8 tokenDecimals)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchRemoteRegistered(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeRemoteRegistered, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (event.Subscription, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "RemoteRegistered", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeRemoteRegistered)
				if err := _NativeTokenHome.contract.UnpackLog(event, "RemoteRegistered", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseRemoteRegistered(log types.Log) (*NativeTokenHomeRemoteRegistered, error) {
	event := new(NativeTokenHomeRemoteRegistered)
	if err := _NativeTokenHome.contract.UnpackLog(event, "RemoteRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the NativeTokenHome contract.
type NativeTokenHomeTeleporterAddressPausedIterator struct {
	Event *NativeTokenHomeTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeTeleporterAddressPaused)
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
		it.Event = new(NativeTokenHomeTeleporterAddressPaused)
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
func (it *NativeTokenHomeTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the NativeTokenHome contract.
type NativeTokenHomeTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenHomeTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTeleporterAddressPausedIterator{contract: _NativeTokenHome.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeTeleporterAddressPaused)
				if err := _NativeTokenHome.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseTeleporterAddressPaused(log types.Log) (*NativeTokenHomeTeleporterAddressPaused, error) {
	event := new(NativeTokenHomeTeleporterAddressPaused)
	if err := _NativeTokenHome.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the NativeTokenHome contract.
type NativeTokenHomeTeleporterAddressUnpausedIterator struct {
	Event *NativeTokenHomeTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeTeleporterAddressUnpaused)
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
		it.Event = new(NativeTokenHomeTeleporterAddressUnpaused)
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
func (it *NativeTokenHomeTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the NativeTokenHome contract.
type NativeTokenHomeTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenHomeTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTeleporterAddressUnpausedIterator{contract: _NativeTokenHome.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeTeleporterAddressUnpaused)
				if err := _NativeTokenHome.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*NativeTokenHomeTeleporterAddressUnpaused, error) {
	event := new(NativeTokenHomeTeleporterAddressUnpaused)
	if err := _NativeTokenHome.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeToCosmosTokensSentIterator is returned from FilterToCosmosTokensSent and is used to iterate over the raw logs and unpacked data for ToCosmosTokensSent events raised by the NativeTokenHome contract.
type NativeTokenHomeToCosmosTokensSentIterator struct {
	Event *NativeTokenHomeToCosmosTokensSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeToCosmosTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeToCosmosTokensSent)
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
		it.Event = new(NativeTokenHomeToCosmosTokensSent)
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
func (it *NativeTokenHomeToCosmosTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeToCosmosTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeToCosmosTokensSent represents a ToCosmosTokensSent event raised by the NativeTokenHome contract.
type NativeTokenHomeToCosmosTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               ToCosmosSendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterToCosmosTokensSent is a free log retrieval operation binding the contract event 0x79bd37563e6b913ff0a0e0c938210c007ebef4c942d38333576d03cb04620e9b.
//
// Solidity: event ToCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterToCosmosTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenHomeToCosmosTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "ToCosmosTokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeToCosmosTokensSentIterator{contract: _NativeTokenHome.contract, event: "ToCosmosTokensSent", logs: logs, sub: sub}, nil
}

// WatchToCosmosTokensSent is a free log subscription operation binding the contract event 0x79bd37563e6b913ff0a0e0c938210c007ebef4c942d38333576d03cb04620e9b.
//
// Solidity: event ToCosmosTokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,bytes,bytes32,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchToCosmosTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeToCosmosTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "ToCosmosTokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeToCosmosTokensSent)
				if err := _NativeTokenHome.contract.UnpackLog(event, "ToCosmosTokensSent", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseToCosmosTokensSent(log types.Log) (*NativeTokenHomeToCosmosTokensSent, error) {
	event := new(NativeTokenHomeToCosmosTokensSent)
	if err := _NativeTokenHome.contract.UnpackLog(event, "ToCosmosTokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeTokensAndCallRoutedIterator is returned from FilterTokensAndCallRouted and is used to iterate over the raw logs and unpacked data for TokensAndCallRouted events raised by the NativeTokenHome contract.
type NativeTokenHomeTokensAndCallRoutedIterator struct {
	Event *NativeTokenHomeTokensAndCallRouted // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeTokensAndCallRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeTokensAndCallRouted)
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
		it.Event = new(NativeTokenHomeTokensAndCallRouted)
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
func (it *NativeTokenHomeTokensAndCallRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeTokensAndCallRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeTokensAndCallRouted represents a TokensAndCallRouted event raised by the NativeTokenHome contract.
type NativeTokenHomeTokensAndCallRouted struct {
	TeleporterMessageID [32]byte
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallRouted is a free log retrieval operation binding the contract event 0x42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterTokensAndCallRouted(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenHomeTokensAndCallRoutedIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "TokensAndCallRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTokensAndCallRoutedIterator{contract: _NativeTokenHome.contract, event: "TokensAndCallRouted", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallRouted is a free log subscription operation binding the contract event 0x42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchTokensAndCallRouted(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeTokensAndCallRouted, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "TokensAndCallRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeTokensAndCallRouted)
				if err := _NativeTokenHome.contract.UnpackLog(event, "TokensAndCallRouted", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseTokensAndCallRouted(log types.Log) (*NativeTokenHomeTokensAndCallRouted, error) {
	event := new(NativeTokenHomeTokensAndCallRouted)
	if err := _NativeTokenHome.contract.UnpackLog(event, "TokensAndCallRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the NativeTokenHome contract.
type NativeTokenHomeTokensAndCallSentIterator struct {
	Event *NativeTokenHomeTokensAndCallSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeTokensAndCallSent)
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
		it.Event = new(NativeTokenHomeTokensAndCallSent)
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
func (it *NativeTokenHomeTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeTokensAndCallSent represents a TokensAndCallSent event raised by the NativeTokenHome contract.
type NativeTokenHomeTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenHomeTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTokensAndCallSentIterator{contract: _NativeTokenHome.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeTokensAndCallSent)
				if err := _NativeTokenHome.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseTokensAndCallSent(log types.Log) (*NativeTokenHomeTokensAndCallSent, error) {
	event := new(NativeTokenHomeTokensAndCallSent)
	if err := _NativeTokenHome.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeTokensCosmosWithdrawnIterator is returned from FilterTokensCosmosWithdrawn and is used to iterate over the raw logs and unpacked data for TokensCosmosWithdrawn events raised by the NativeTokenHome contract.
type NativeTokenHomeTokensCosmosWithdrawnIterator struct {
	Event *NativeTokenHomeTokensCosmosWithdrawn // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeTokensCosmosWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeTokensCosmosWithdrawn)
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
		it.Event = new(NativeTokenHomeTokensCosmosWithdrawn)
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
func (it *NativeTokenHomeTokensCosmosWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeTokensCosmosWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeTokensCosmosWithdrawn represents a TokensCosmosWithdrawn event raised by the NativeTokenHome contract.
type NativeTokenHomeTokensCosmosWithdrawn struct {
	MessageID [32]byte
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensCosmosWithdrawn is a free log retrieval operation binding the contract event 0xfc043975d4effd7520860981e63ea2ba7197dbe72a39dd32a3e227e820ca7ffe.
//
// Solidity: event TokensCosmosWithdrawn(bytes32 messageID, address indexed recipient, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterTokensCosmosWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenHomeTokensCosmosWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "TokensCosmosWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTokensCosmosWithdrawnIterator{contract: _NativeTokenHome.contract, event: "TokensCosmosWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensCosmosWithdrawn is a free log subscription operation binding the contract event 0xfc043975d4effd7520860981e63ea2ba7197dbe72a39dd32a3e227e820ca7ffe.
//
// Solidity: event TokensCosmosWithdrawn(bytes32 messageID, address indexed recipient, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchTokensCosmosWithdrawn(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeTokensCosmosWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "TokensCosmosWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeTokensCosmosWithdrawn)
				if err := _NativeTokenHome.contract.UnpackLog(event, "TokensCosmosWithdrawn", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseTokensCosmosWithdrawn(log types.Log) (*NativeTokenHomeTokensCosmosWithdrawn, error) {
	event := new(NativeTokenHomeTokensCosmosWithdrawn)
	if err := _NativeTokenHome.contract.UnpackLog(event, "TokensCosmosWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeTokensRoutedIterator is returned from FilterTokensRouted and is used to iterate over the raw logs and unpacked data for TokensRouted events raised by the NativeTokenHome contract.
type NativeTokenHomeTokensRoutedIterator struct {
	Event *NativeTokenHomeTokensRouted // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeTokensRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeTokensRouted)
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
		it.Event = new(NativeTokenHomeTokensRouted)
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
func (it *NativeTokenHomeTokensRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeTokensRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeTokensRouted represents a TokensRouted event raised by the NativeTokenHome contract.
type NativeTokenHomeTokensRouted struct {
	TeleporterMessageID [32]byte
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensRouted is a free log retrieval operation binding the contract event 0x825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf0.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterTokensRouted(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenHomeTokensRoutedIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "TokensRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTokensRoutedIterator{contract: _NativeTokenHome.contract, event: "TokensRouted", logs: logs, sub: sub}, nil
}

// WatchTokensRouted is a free log subscription operation binding the contract event 0x825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf0.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchTokensRouted(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeTokensRouted, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "TokensRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeTokensRouted)
				if err := _NativeTokenHome.contract.UnpackLog(event, "TokensRouted", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseTokensRouted(log types.Log) (*NativeTokenHomeTokensRouted, error) {
	event := new(NativeTokenHomeTokensRouted)
	if err := _NativeTokenHome.contract.UnpackLog(event, "TokensRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the NativeTokenHome contract.
type NativeTokenHomeTokensSentIterator struct {
	Event *NativeTokenHomeTokensSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeTokensSent)
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
		it.Event = new(NativeTokenHomeTokensSent)
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
func (it *NativeTokenHomeTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeTokensSent represents a TokensSent event raised by the NativeTokenHome contract.
type NativeTokenHomeTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenHomeTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTokensSentIterator{contract: _NativeTokenHome.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeTokensSent)
				if err := _NativeTokenHome.contract.UnpackLog(event, "TokensSent", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseTokensSent(log types.Log) (*NativeTokenHomeTokensSent, error) {
	event := new(NativeTokenHomeTokensSent)
	if err := _NativeTokenHome.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the NativeTokenHome contract.
type NativeTokenHomeTokensWithdrawnIterator struct {
	Event *NativeTokenHomeTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeTokensWithdrawn)
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
		it.Event = new(NativeTokenHomeTokensWithdrawn)
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
func (it *NativeTokenHomeTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeTokensWithdrawn represents a TokensWithdrawn event raised by the NativeTokenHome contract.
type NativeTokenHomeTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenHomeTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTokensWithdrawnIterator{contract: _NativeTokenHome.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeTokensWithdrawn)
				if err := _NativeTokenHome.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseTokensWithdrawn(log types.Log) (*NativeTokenHomeTokensWithdrawn, error) {
	event := new(NativeTokenHomeTokensWithdrawn)
	if err := _NativeTokenHome.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
