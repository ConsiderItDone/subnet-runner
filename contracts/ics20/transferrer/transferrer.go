// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transferrer

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Height is an auto generated low-level Go binding around an user-defined struct.
type Height struct {
	RevisionNumber *big.Int
	RevisionHeight *big.Int
}

// TransferrerMetaData contains all meta data concerning the Transferrer contract.
var TransferrerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"channelCapability\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TransferrerABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferrerMetaData.ABI instead.
var TransferrerABI = TransferrerMetaData.ABI

// Transferrer is an auto generated Go binding around an Ethereum contract.
type Transferrer struct {
	TransferrerCaller     // Read-only binding to the contract
	TransferrerTransactor // Write-only binding to the contract
	TransferrerFilterer   // Log filterer for contract events
}

// TransferrerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferrerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferrerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferrerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferrerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferrerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferrerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferrerSession struct {
	Contract     *Transferrer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferrerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferrerCallerSession struct {
	Contract *TransferrerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TransferrerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferrerTransactorSession struct {
	Contract     *TransferrerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TransferrerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferrerRaw struct {
	Contract *Transferrer // Generic contract binding to access the raw methods on
}

// TransferrerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferrerCallerRaw struct {
	Contract *TransferrerCaller // Generic read-only contract binding to access the raw methods on
}

// TransferrerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferrerTransactorRaw struct {
	Contract *TransferrerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferrer creates a new instance of Transferrer, bound to a specific deployed contract.
func NewTransferrer(address common.Address, backend bind.ContractBackend) (*Transferrer, error) {
	contract, err := bindTransferrer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transferrer{TransferrerCaller: TransferrerCaller{contract: contract}, TransferrerTransactor: TransferrerTransactor{contract: contract}, TransferrerFilterer: TransferrerFilterer{contract: contract}}, nil
}

// NewTransferrerCaller creates a new read-only instance of Transferrer, bound to a specific deployed contract.
func NewTransferrerCaller(address common.Address, caller bind.ContractCaller) (*TransferrerCaller, error) {
	contract, err := bindTransferrer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferrerCaller{contract: contract}, nil
}

// NewTransferrerTransactor creates a new write-only instance of Transferrer, bound to a specific deployed contract.
func NewTransferrerTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferrerTransactor, error) {
	contract, err := bindTransferrer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferrerTransactor{contract: contract}, nil
}

// NewTransferrerFilterer creates a new log filterer instance of Transferrer, bound to a specific deployed contract.
func NewTransferrerFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferrerFilterer, error) {
	contract, err := bindTransferrer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferrerFilterer{contract: contract}, nil
}

// bindTransferrer binds a generic wrapper to an already deployed contract.
func bindTransferrer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransferrerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transferrer *TransferrerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transferrer.Contract.TransferrerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transferrer *TransferrerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transferrer.Contract.TransferrerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transferrer *TransferrerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transferrer.Contract.TransferrerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transferrer *TransferrerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transferrer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transferrer *TransferrerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transferrer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transferrer *TransferrerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transferrer.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0x2c4a1bee.
//
// Solidity: function transfer(uint256 channelCapability, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes data) returns()
func (_Transferrer *TransferrerTransactor) Transfer(opts *bind.TransactOpts, channelCapability *big.Int, sourcePort string, sourceChannel string, timeoutHeight Height, timeoutTimestamp *big.Int, data []byte) (*types.Transaction, error) {
	return _Transferrer.contract.Transact(opts, "transfer", channelCapability, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// Transfer is a paid mutator transaction binding the contract method 0x2c4a1bee.
//
// Solidity: function transfer(uint256 channelCapability, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes data) returns()
func (_Transferrer *TransferrerSession) Transfer(channelCapability *big.Int, sourcePort string, sourceChannel string, timeoutHeight Height, timeoutTimestamp *big.Int, data []byte) (*types.Transaction, error) {
	return _Transferrer.Contract.Transfer(&_Transferrer.TransactOpts, channelCapability, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// Transfer is a paid mutator transaction binding the contract method 0x2c4a1bee.
//
// Solidity: function transfer(uint256 channelCapability, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes data) returns()
func (_Transferrer *TransferrerTransactorSession) Transfer(channelCapability *big.Int, sourcePort string, sourceChannel string, timeoutHeight Height, timeoutTimestamp *big.Int, data []byte) (*types.Transaction, error) {
	return _Transferrer.Contract.Transfer(&_Transferrer.TransactOpts, channelCapability, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}
