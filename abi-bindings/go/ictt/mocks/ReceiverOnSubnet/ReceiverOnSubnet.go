// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package receiveronsubnet

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

// ReceiverOnSubnetMetaData contains all meta data concerning the ReceiverOnSubnet contract.
var ReceiverOnSubnetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"lastMessage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052731411cc804d13231ce66217eb296a09d3c10ce785608052348015602757600080fd5b50608051610529610049600039600081816069015261015101526105296000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806332970710146100465780633cb747bf14610064578063c868efaa146100a3575b600080fd5b61004e6100b8565b60405161005b91906101fd565b60405180910390f35b61008b7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161005b565b6100b66100b136600461024c565b610146565b005b600080546100c5906102e1565b80601f01602080910402602001604051908101604052809291908181526020018280546100f1906102e1565b801561013e5780601f106101135761010080835404028352916020019161013e565b820191906000526020600020905b81548152906001019060200180831161012157829003601f168201915b505050505081565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101dd5760405162461bcd60e51b815260206004820152603260248201527f52656365697665724f6e5375626e65743a20756e617574686f72697a6564205460448201527132b632b837b93a32b926b2b9b9b2b733b2b960711b606482015260840160405180910390fd5b6101e981830183610331565b6000906101f69082610433565b5050505050565b60006020808352835180602085015260005b8181101561022b5785810183015185820160400152820161020f565b506000604082860101526040601f19601f8301168501019250505092915050565b6000806000806060858703121561026257600080fd5b8435935060208501356001600160a01b038116811461028057600080fd5b9250604085013567ffffffffffffffff8082111561029d57600080fd5b818701915087601f8301126102b157600080fd5b8135818111156102c057600080fd5b8860208285010111156102d257600080fd5b95989497505060200194505050565b600181811c908216806102f557607f821691505b60208210810361031557634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561034357600080fd5b813567ffffffffffffffff8082111561035b57600080fd5b818401915084601f83011261036f57600080fd5b8135818111156103815761038161031b565b604051601f8201601f19908116603f011681019083821181831017156103a9576103a961031b565b816040528281528760208487010111156103c257600080fd5b826020860160208301376000928101602001929092525095945050505050565b601f82111561042e576000816000526020600020601f850160051c8101602086101561040b5750805b601f850160051c820191505b8181101561042a57828155600101610417565b5050505b505050565b815167ffffffffffffffff81111561044d5761044d61031b565b6104618161045b84546102e1565b846103e2565b602080601f831160018114610496576000841561047e5750858301515b600019600386901b1c1916600185901b17855561042a565b600085815260208120601f198616915b828110156104c5578886015182559484019460019091019084016104a6565b50858210156104e35787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea26469706673582212209ea622f297aa1a17e2ee12956287820e99fea367494f4c1acfd6014546fb6eef64736f6c63430008190033",
}

// ReceiverOnSubnetABI is the input ABI used to generate the binding from.
// Deprecated: Use ReceiverOnSubnetMetaData.ABI instead.
var ReceiverOnSubnetABI = ReceiverOnSubnetMetaData.ABI

// ReceiverOnSubnetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReceiverOnSubnetMetaData.Bin instead.
var ReceiverOnSubnetBin = ReceiverOnSubnetMetaData.Bin

// DeployReceiverOnSubnet deploys a new Ethereum contract, binding an instance of ReceiverOnSubnet to it.
func DeployReceiverOnSubnet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReceiverOnSubnet, error) {
	parsed, err := ReceiverOnSubnetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiverOnSubnetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReceiverOnSubnet{ReceiverOnSubnetCaller: ReceiverOnSubnetCaller{contract: contract}, ReceiverOnSubnetTransactor: ReceiverOnSubnetTransactor{contract: contract}, ReceiverOnSubnetFilterer: ReceiverOnSubnetFilterer{contract: contract}}, nil
}

// ReceiverOnSubnet is an auto generated Go binding around an Ethereum contract.
type ReceiverOnSubnet struct {
	ReceiverOnSubnetCaller     // Read-only binding to the contract
	ReceiverOnSubnetTransactor // Write-only binding to the contract
	ReceiverOnSubnetFilterer   // Log filterer for contract events
}

// ReceiverOnSubnetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiverOnSubnetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverOnSubnetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiverOnSubnetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverOnSubnetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiverOnSubnetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverOnSubnetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiverOnSubnetSession struct {
	Contract     *ReceiverOnSubnet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiverOnSubnetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiverOnSubnetCallerSession struct {
	Contract *ReceiverOnSubnetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ReceiverOnSubnetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiverOnSubnetTransactorSession struct {
	Contract     *ReceiverOnSubnetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ReceiverOnSubnetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiverOnSubnetRaw struct {
	Contract *ReceiverOnSubnet // Generic contract binding to access the raw methods on
}

// ReceiverOnSubnetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiverOnSubnetCallerRaw struct {
	Contract *ReceiverOnSubnetCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiverOnSubnetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiverOnSubnetTransactorRaw struct {
	Contract *ReceiverOnSubnetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiverOnSubnet creates a new instance of ReceiverOnSubnet, bound to a specific deployed contract.
func NewReceiverOnSubnet(address common.Address, backend bind.ContractBackend) (*ReceiverOnSubnet, error) {
	contract, err := bindReceiverOnSubnet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiverOnSubnet{ReceiverOnSubnetCaller: ReceiverOnSubnetCaller{contract: contract}, ReceiverOnSubnetTransactor: ReceiverOnSubnetTransactor{contract: contract}, ReceiverOnSubnetFilterer: ReceiverOnSubnetFilterer{contract: contract}}, nil
}

// NewReceiverOnSubnetCaller creates a new read-only instance of ReceiverOnSubnet, bound to a specific deployed contract.
func NewReceiverOnSubnetCaller(address common.Address, caller bind.ContractCaller) (*ReceiverOnSubnetCaller, error) {
	contract, err := bindReceiverOnSubnet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverOnSubnetCaller{contract: contract}, nil
}

// NewReceiverOnSubnetTransactor creates a new write-only instance of ReceiverOnSubnet, bound to a specific deployed contract.
func NewReceiverOnSubnetTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiverOnSubnetTransactor, error) {
	contract, err := bindReceiverOnSubnet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverOnSubnetTransactor{contract: contract}, nil
}

// NewReceiverOnSubnetFilterer creates a new log filterer instance of ReceiverOnSubnet, bound to a specific deployed contract.
func NewReceiverOnSubnetFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiverOnSubnetFilterer, error) {
	contract, err := bindReceiverOnSubnet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiverOnSubnetFilterer{contract: contract}, nil
}

// bindReceiverOnSubnet binds a generic wrapper to an already deployed contract.
func bindReceiverOnSubnet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReceiverOnSubnetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiverOnSubnet *ReceiverOnSubnetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiverOnSubnet.Contract.ReceiverOnSubnetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiverOnSubnet *ReceiverOnSubnetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiverOnSubnet.Contract.ReceiverOnSubnetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiverOnSubnet *ReceiverOnSubnetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiverOnSubnet.Contract.ReceiverOnSubnetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiverOnSubnet *ReceiverOnSubnetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiverOnSubnet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiverOnSubnet *ReceiverOnSubnetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiverOnSubnet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiverOnSubnet *ReceiverOnSubnetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiverOnSubnet.Contract.contract.Transact(opts, method, params...)
}

// LastMessage is a free data retrieval call binding the contract method 0x32970710.
//
// Solidity: function lastMessage() view returns(string)
func (_ReceiverOnSubnet *ReceiverOnSubnetCaller) LastMessage(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ReceiverOnSubnet.contract.Call(opts, &out, "lastMessage")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LastMessage is a free data retrieval call binding the contract method 0x32970710.
//
// Solidity: function lastMessage() view returns(string)
func (_ReceiverOnSubnet *ReceiverOnSubnetSession) LastMessage() (string, error) {
	return _ReceiverOnSubnet.Contract.LastMessage(&_ReceiverOnSubnet.CallOpts)
}

// LastMessage is a free data retrieval call binding the contract method 0x32970710.
//
// Solidity: function lastMessage() view returns(string)
func (_ReceiverOnSubnet *ReceiverOnSubnetCallerSession) LastMessage() (string, error) {
	return _ReceiverOnSubnet.Contract.LastMessage(&_ReceiverOnSubnet.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_ReceiverOnSubnet *ReceiverOnSubnetCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReceiverOnSubnet.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_ReceiverOnSubnet *ReceiverOnSubnetSession) Messenger() (common.Address, error) {
	return _ReceiverOnSubnet.Contract.Messenger(&_ReceiverOnSubnet.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_ReceiverOnSubnet *ReceiverOnSubnetCallerSession) Messenger() (common.Address, error) {
	return _ReceiverOnSubnet.Contract.Messenger(&_ReceiverOnSubnet.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 , address , bytes message) returns()
func (_ReceiverOnSubnet *ReceiverOnSubnetTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, arg0 [32]byte, arg1 common.Address, message []byte) (*types.Transaction, error) {
	return _ReceiverOnSubnet.contract.Transact(opts, "receiveTeleporterMessage", arg0, arg1, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 , address , bytes message) returns()
func (_ReceiverOnSubnet *ReceiverOnSubnetSession) ReceiveTeleporterMessage(arg0 [32]byte, arg1 common.Address, message []byte) (*types.Transaction, error) {
	return _ReceiverOnSubnet.Contract.ReceiveTeleporterMessage(&_ReceiverOnSubnet.TransactOpts, arg0, arg1, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 , address , bytes message) returns()
func (_ReceiverOnSubnet *ReceiverOnSubnetTransactorSession) ReceiveTeleporterMessage(arg0 [32]byte, arg1 common.Address, message []byte) (*types.Transaction, error) {
	return _ReceiverOnSubnet.Contract.ReceiveTeleporterMessage(&_ReceiverOnSubnet.TransactOpts, arg0, arg1, message)
}
