// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package senderonsubnet

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

// SenderOnSubnetMetaData contains all meta data concerning the SenderOnSubnet contract.
var SenderOnSubnetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052731411cc804d13231ce66217eb296a09d3c10ce785608052348015602757600080fd5b5060805161043d61004860003960008181604001526095015261043d6000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633cb747bf1461003b578063de6f24bb1461007e575b600080fd5b6100627f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b61009161008c366004610206565b610093565b005b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663624488506040518060c001604052807fbac32c66fba9d83a8cd38951c49dbc53d2325dbc891071767f6dddfcccb4195360001b8152602001866001600160a01b03168152602001604051806040016040528060006001600160a01b0316815260200160008152508152602001620186a08152602001600067ffffffffffffffff81111561014d5761014d610297565b604051908082528060200260200182016040528015610176578160200160208202803683370190505b508152602001858560405160200161018f9291906102ad565b6040516020818303038152906040528152506040518263ffffffff1660e01b81526004016101bd9190610367565b6020604051808303816000875af11580156101dc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020091906103ee565b50505050565b60008060006040848603121561021b57600080fd5b83356001600160a01b038116811461023257600080fd5b9250602084013567ffffffffffffffff8082111561024f57600080fd5b818601915086601f83011261026357600080fd5b81358181111561027257600080fd5b87602082850101111561028457600080fd5b6020830194508093505050509250925092565b634e487b7160e01b600052604160045260246000fd5b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60008151808452602080850194506020840160005b838110156103165781516001600160a01b0316875295820195908201906001016102f1565b509495945050505050565b6000815180845260005b818110156103475760208185018101518683018201520161032b565b506000602082860101526020601f19601f83011685010191505092915050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c08401526103c86101008401826102dc565b905060a0840151601f198483030160e08501526103e58282610321565b95945050505050565b60006020828403121561040057600080fd5b505191905056fea2646970667358221220ce0d5ac4fb8474ff8bf8bdc915ad5f4e2f8ba759d12a5fefa12050f0e2cf779f64736f6c63430008190033",
}

// SenderOnSubnetABI is the input ABI used to generate the binding from.
// Deprecated: Use SenderOnSubnetMetaData.ABI instead.
var SenderOnSubnetABI = SenderOnSubnetMetaData.ABI

// SenderOnSubnetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SenderOnSubnetMetaData.Bin instead.
var SenderOnSubnetBin = SenderOnSubnetMetaData.Bin

// DeploySenderOnSubnet deploys a new Ethereum contract, binding an instance of SenderOnSubnet to it.
func DeploySenderOnSubnet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SenderOnSubnet, error) {
	parsed, err := SenderOnSubnetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SenderOnSubnetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SenderOnSubnet{SenderOnSubnetCaller: SenderOnSubnetCaller{contract: contract}, SenderOnSubnetTransactor: SenderOnSubnetTransactor{contract: contract}, SenderOnSubnetFilterer: SenderOnSubnetFilterer{contract: contract}}, nil
}

// SenderOnSubnet is an auto generated Go binding around an Ethereum contract.
type SenderOnSubnet struct {
	SenderOnSubnetCaller     // Read-only binding to the contract
	SenderOnSubnetTransactor // Write-only binding to the contract
	SenderOnSubnetFilterer   // Log filterer for contract events
}

// SenderOnSubnetCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenderOnSubnetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderOnSubnetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenderOnSubnetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderOnSubnetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenderOnSubnetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderOnSubnetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenderOnSubnetSession struct {
	Contract     *SenderOnSubnet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenderOnSubnetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenderOnSubnetCallerSession struct {
	Contract *SenderOnSubnetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SenderOnSubnetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenderOnSubnetTransactorSession struct {
	Contract     *SenderOnSubnetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SenderOnSubnetRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenderOnSubnetRaw struct {
	Contract *SenderOnSubnet // Generic contract binding to access the raw methods on
}

// SenderOnSubnetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenderOnSubnetCallerRaw struct {
	Contract *SenderOnSubnetCaller // Generic read-only contract binding to access the raw methods on
}

// SenderOnSubnetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenderOnSubnetTransactorRaw struct {
	Contract *SenderOnSubnetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSenderOnSubnet creates a new instance of SenderOnSubnet, bound to a specific deployed contract.
func NewSenderOnSubnet(address common.Address, backend bind.ContractBackend) (*SenderOnSubnet, error) {
	contract, err := bindSenderOnSubnet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SenderOnSubnet{SenderOnSubnetCaller: SenderOnSubnetCaller{contract: contract}, SenderOnSubnetTransactor: SenderOnSubnetTransactor{contract: contract}, SenderOnSubnetFilterer: SenderOnSubnetFilterer{contract: contract}}, nil
}

// NewSenderOnSubnetCaller creates a new read-only instance of SenderOnSubnet, bound to a specific deployed contract.
func NewSenderOnSubnetCaller(address common.Address, caller bind.ContractCaller) (*SenderOnSubnetCaller, error) {
	contract, err := bindSenderOnSubnet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenderOnSubnetCaller{contract: contract}, nil
}

// NewSenderOnSubnetTransactor creates a new write-only instance of SenderOnSubnet, bound to a specific deployed contract.
func NewSenderOnSubnetTransactor(address common.Address, transactor bind.ContractTransactor) (*SenderOnSubnetTransactor, error) {
	contract, err := bindSenderOnSubnet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenderOnSubnetTransactor{contract: contract}, nil
}

// NewSenderOnSubnetFilterer creates a new log filterer instance of SenderOnSubnet, bound to a specific deployed contract.
func NewSenderOnSubnetFilterer(address common.Address, filterer bind.ContractFilterer) (*SenderOnSubnetFilterer, error) {
	contract, err := bindSenderOnSubnet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenderOnSubnetFilterer{contract: contract}, nil
}

// bindSenderOnSubnet binds a generic wrapper to an already deployed contract.
func bindSenderOnSubnet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SenderOnSubnetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenderOnSubnet *SenderOnSubnetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenderOnSubnet.Contract.SenderOnSubnetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenderOnSubnet *SenderOnSubnetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenderOnSubnet.Contract.SenderOnSubnetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenderOnSubnet *SenderOnSubnetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenderOnSubnet.Contract.SenderOnSubnetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenderOnSubnet *SenderOnSubnetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenderOnSubnet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenderOnSubnet *SenderOnSubnetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenderOnSubnet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenderOnSubnet *SenderOnSubnetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenderOnSubnet.Contract.contract.Transact(opts, method, params...)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_SenderOnSubnet *SenderOnSubnetCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SenderOnSubnet.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_SenderOnSubnet *SenderOnSubnetSession) Messenger() (common.Address, error) {
	return _SenderOnSubnet.Contract.Messenger(&_SenderOnSubnet.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_SenderOnSubnet *SenderOnSubnetCallerSession) Messenger() (common.Address, error) {
	return _SenderOnSubnet.Contract.Messenger(&_SenderOnSubnet.CallOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0xde6f24bb.
//
// Solidity: function sendMessage(address destinationAddress, string message) returns()
func (_SenderOnSubnet *SenderOnSubnetTransactor) SendMessage(opts *bind.TransactOpts, destinationAddress common.Address, message string) (*types.Transaction, error) {
	return _SenderOnSubnet.contract.Transact(opts, "sendMessage", destinationAddress, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xde6f24bb.
//
// Solidity: function sendMessage(address destinationAddress, string message) returns()
func (_SenderOnSubnet *SenderOnSubnetSession) SendMessage(destinationAddress common.Address, message string) (*types.Transaction, error) {
	return _SenderOnSubnet.Contract.SendMessage(&_SenderOnSubnet.TransactOpts, destinationAddress, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xde6f24bb.
//
// Solidity: function sendMessage(address destinationAddress, string message) returns()
func (_SenderOnSubnet *SenderOnSubnetTransactorSession) SendMessage(destinationAddress common.Address, message string) (*types.Transaction, error) {
	return _SenderOnSubnet.Contract.SendMessage(&_SenderOnSubnet.TransactOpts, destinationAddress, message)
}
