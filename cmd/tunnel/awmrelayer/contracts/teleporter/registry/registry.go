// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// ProtocolRegistryEntry is an auto generated low-level Go binding around an user-defined struct.
type ProtocolRegistryEntry struct {
	Version         *big.Int
	ProtocolAddress common.Address
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialEntries\",\"type\":\"tuple[]\",\"internalType\":\"structProtocolRegistryEntry[]\",\"components\":[{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"protocolAddress\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_VERSION_INCREMENT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATORS_SOURCE_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addProtocolVersion\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blockchainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddressFromVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestTeleporter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITeleporterMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTeleporterFromVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITeleporterMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVersionFromAddress\",\"inputs\":[{\"name\":\"protocolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddProtocolVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"protocolAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LatestVersionUpdated\",\"inputs\":[{\"name\":\"oldVersion\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newVersion\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610cfa380380610cfa83398101604081905261002f9161011c565b7302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015610081573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100a591906101ff565b60805250610218565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156100e6576100e66100ae565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610114576101146100ae565b604052919050565b6000602080838503121561012f57600080fd5b82516001600160401b038082111561014657600080fd5b818501915085601f83011261015a57600080fd5b81518181111561016c5761016c6100ae565b61017a848260051b016100ec565b818152848101925060069190911b83018401908782111561019a57600080fd5b928401925b818410156101f457604084890312156101b85760008081fd5b6101c06100c4565b84518152858501516001600160a01b03811681146101de5760008081fd5b818701528352604093909301929184019161019f565b979650505050505050565b60006020828403121561021157600080fd5b5051919050565b608051610ac061023a6000396000818161014901526102640152610ac06000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063ac473ac311610066578063ac473ac314610124578063b771b3bc1461012d578063c07f47d41461013b578063d127dc9b14610144578063d820e64f1461016b57600080fd5b80630731775d146100a3578063215abce9146100c857806341f34ed9146100db57806346f9ef49146100f05780634c1f08ce14610103575b600080fd5b6100ab600081565b6040516001600160a01b0390911681526020015b60405180910390f35b6100ab6100d63660046107c5565b610173565b6100ee6100e93660046107de565b610184565b005b6100ab6100fe3660046107c5565b6103f9565b610116610111366004610823565b6104be565b6040519081526020016100bf565b6101166101f481565b6100ab6005600160991b0181565b61011660005481565b6101167f000000000000000000000000000000000000000000000000000000000000000081565b6100ab610566565b600061017e826103f9565b92915050565b6040516306f8253560e41b815263ffffffff8216600482015260009081906005600160991b0190636f82535090602401600060405180830381865afa1580156101d1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526101f991908101906108c5565b91509150806102605760405162461bcd60e51b815260206004820152602860248201527f54656c65706f7274657252656769737472793a20696e76616c69642077617270604482015267206d65737361676560c01b60648201526084015b60405180910390fd5b81517f0000000000000000000000000000000000000000000000000000000000000000146102e45760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f7274657252656769737472793a20696e76616c696420736f757260448201526a18d94818da185a5b88125160aa1b6064820152608401610257565b60208201516001600160a01b0316156103595760405162461bcd60e51b815260206004820152603160248201527f54656c65706f7274657252656769737472793a20696e76616c6964206f726967604482015270696e2073656e646572206164647265737360781b6064820152608401610257565b600080836040015180602001905181019061037491906109cd565b90925090506001600160a01b03811630146103e95760405162461bcd60e51b815260206004820152602f60248201527f54656c65706f7274657252656769737472793a20696e76616c6964206465737460448201526e696e6174696f6e206164647265737360881b6064820152608401610257565b6103f282610578565b5050505050565b60008160000361044b5760405162461bcd60e51b815260206004820181905260248201527f54656c65706f7274657252656769737472793a207a65726f2076657273696f6e6044820152606401610257565b6000828152600160205260409020546001600160a01b03168061017e5760405162461bcd60e51b815260206004820152602560248201527f54656c65706f7274657252656769737472793a2076657273696f6e206e6f7420604482015264199bdd5b9960da1b6064820152608401610257565b60006001600160a01b0382166104e65760405162461bcd60e51b815260040161025790610a49565b6001600160a01b0382166000908152600260205260408120549081900361017e5760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f7274657252656769737472793a2070726f746f636f6c2061646460448201526d1c995cdcc81b9bdd08199bdd5b9960921b6064820152608401610257565b60006105736000546103f9565b905090565b80516000036105c95760405162461bcd60e51b815260206004820181905260248201527f54656c65706f7274657252656769737472793a207a65726f2076657273696f6e6044820152606401610257565b80516000908152600160205260409020546001600160a01b0316156106435760405162461bcd60e51b815260206004820152602a60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20616c72656044820152696164792065786973747360b01b6064820152608401610257565b60208101516001600160a01b031661066d5760405162461bcd60e51b815260040161025790610a49565b60005461067c6101f482610a92565b825111156106e35760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20696e637260448201526d0cadacadce840e8dede40d0d2ced60931b6064820152608401610257565b602082810180518451600090815260018452604080822080546001600160a01b0319166001600160a01b039485161790559251909116815260029092529020548251111561074c5781516020808401516001600160a01b03166000908152600290915260409020555b602082015182516040516001600160a01b03909216917fa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a90600090a381518110156107c1578151600081815560405183917f30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b091a35b5050565b6000602082840312156107d757600080fd5b5035919050565b6000602082840312156107f057600080fd5b813563ffffffff8116811461080457600080fd5b9392505050565b6001600160a01b038116811461082057600080fd5b50565b60006020828403121561083557600080fd5b81356108048161080b565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561087957610879610840565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156108a8576108a8610840565b604052919050565b805180151581146108c057600080fd5b919050565b600080604083850312156108d857600080fd5b825167ffffffffffffffff808211156108f057600080fd5b908401906060828703121561090457600080fd5b61090c610856565b8251815260208084015161091f8161080b565b8282015260408401518381111561093557600080fd5b80850194505087601f85011261094a57600080fd5b83518381111561095c5761095c610840565b61096e601f8201601f1916830161087f565b9350808452888282870101111561098457600080fd5b60005b818110156109a2578581018301518582018401528201610987565b506000828286010152508260408301528195506109c08188016108b0565b9450505050509250929050565b60008082840360608112156109e157600080fd5b60408112156109ef57600080fd5b506040516040810181811067ffffffffffffffff82111715610a1357610a13610840565b604052835181526020840151610a288161080b565b60208201526040840151909250610a3e8161080b565b809150509250929050565b60208082526029908201527f54656c65706f7274657252656769737472793a207a65726f2070726f746f636f6040820152686c206164647265737360b81b606082015260800190565b8082018082111561017e57634e487b7160e01b600052601160045260246000fdfea164736f6c6343000819000a",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// RegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegistryMetaData.Bin instead.
var RegistryBin = RegistryMetaData.Bin

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, initialEntries []ProtocolRegistryEntry) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegistryBin), backend, initialEntries)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// MAXVERSIONINCREMENT is a free data retrieval call binding the contract method 0xac473ac3.
//
// Solidity: function MAX_VERSION_INCREMENT() view returns(uint256)
func (_Registry *RegistryCaller) MAXVERSIONINCREMENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "MAX_VERSION_INCREMENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERSIONINCREMENT is a free data retrieval call binding the contract method 0xac473ac3.
//
// Solidity: function MAX_VERSION_INCREMENT() view returns(uint256)
func (_Registry *RegistrySession) MAXVERSIONINCREMENT() (*big.Int, error) {
	return _Registry.Contract.MAXVERSIONINCREMENT(&_Registry.CallOpts)
}

// MAXVERSIONINCREMENT is a free data retrieval call binding the contract method 0xac473ac3.
//
// Solidity: function MAX_VERSION_INCREMENT() view returns(uint256)
func (_Registry *RegistryCallerSession) MAXVERSIONINCREMENT() (*big.Int, error) {
	return _Registry.Contract.MAXVERSIONINCREMENT(&_Registry.CallOpts)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_Registry *RegistryCaller) VALIDATORSSOURCEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "VALIDATORS_SOURCE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_Registry *RegistrySession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _Registry.Contract.VALIDATORSSOURCEADDRESS(&_Registry.CallOpts)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_Registry *RegistryCallerSession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _Registry.Contract.VALIDATORSSOURCEADDRESS(&_Registry.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_Registry *RegistryCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_Registry *RegistrySession) WARPMESSENGER() (common.Address, error) {
	return _Registry.Contract.WARPMESSENGER(&_Registry.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_Registry *RegistryCallerSession) WARPMESSENGER() (common.Address, error) {
	return _Registry.Contract.WARPMESSENGER(&_Registry.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_Registry *RegistryCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_Registry *RegistrySession) BlockchainID() ([32]byte, error) {
	return _Registry.Contract.BlockchainID(&_Registry.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_Registry *RegistryCallerSession) BlockchainID() ([32]byte, error) {
	return _Registry.Contract.BlockchainID(&_Registry.CallOpts)
}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_Registry *RegistryCaller) GetAddressFromVersion(opts *bind.CallOpts, version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getAddressFromVersion", version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_Registry *RegistrySession) GetAddressFromVersion(version *big.Int) (common.Address, error) {
	return _Registry.Contract.GetAddressFromVersion(&_Registry.CallOpts, version)
}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_Registry *RegistryCallerSession) GetAddressFromVersion(version *big.Int) (common.Address, error) {
	return _Registry.Contract.GetAddressFromVersion(&_Registry.CallOpts, version)
}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_Registry *RegistryCaller) GetLatestTeleporter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getLatestTeleporter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_Registry *RegistrySession) GetLatestTeleporter() (common.Address, error) {
	return _Registry.Contract.GetLatestTeleporter(&_Registry.CallOpts)
}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_Registry *RegistryCallerSession) GetLatestTeleporter() (common.Address, error) {
	return _Registry.Contract.GetLatestTeleporter(&_Registry.CallOpts)
}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_Registry *RegistryCaller) GetTeleporterFromVersion(opts *bind.CallOpts, version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getTeleporterFromVersion", version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_Registry *RegistrySession) GetTeleporterFromVersion(version *big.Int) (common.Address, error) {
	return _Registry.Contract.GetTeleporterFromVersion(&_Registry.CallOpts, version)
}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_Registry *RegistryCallerSession) GetTeleporterFromVersion(version *big.Int) (common.Address, error) {
	return _Registry.Contract.GetTeleporterFromVersion(&_Registry.CallOpts, version)
}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_Registry *RegistryCaller) GetVersionFromAddress(opts *bind.CallOpts, protocolAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getVersionFromAddress", protocolAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_Registry *RegistrySession) GetVersionFromAddress(protocolAddress common.Address) (*big.Int, error) {
	return _Registry.Contract.GetVersionFromAddress(&_Registry.CallOpts, protocolAddress)
}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_Registry *RegistryCallerSession) GetVersionFromAddress(protocolAddress common.Address) (*big.Int, error) {
	return _Registry.Contract.GetVersionFromAddress(&_Registry.CallOpts, protocolAddress)
}

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_Registry *RegistryCaller) LatestVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "latestVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_Registry *RegistrySession) LatestVersion() (*big.Int, error) {
	return _Registry.Contract.LatestVersion(&_Registry.CallOpts)
}

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_Registry *RegistryCallerSession) LatestVersion() (*big.Int, error) {
	return _Registry.Contract.LatestVersion(&_Registry.CallOpts)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_Registry *RegistryTransactor) AddProtocolVersion(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "addProtocolVersion", messageIndex)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_Registry *RegistrySession) AddProtocolVersion(messageIndex uint32) (*types.Transaction, error) {
	return _Registry.Contract.AddProtocolVersion(&_Registry.TransactOpts, messageIndex)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_Registry *RegistryTransactorSession) AddProtocolVersion(messageIndex uint32) (*types.Transaction, error) {
	return _Registry.Contract.AddProtocolVersion(&_Registry.TransactOpts, messageIndex)
}

// RegistryAddProtocolVersionIterator is returned from FilterAddProtocolVersion and is used to iterate over the raw logs and unpacked data for AddProtocolVersion events raised by the Registry contract.
type RegistryAddProtocolVersionIterator struct {
	Event *RegistryAddProtocolVersion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryAddProtocolVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryAddProtocolVersion)
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
		it.Event = new(RegistryAddProtocolVersion)
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
func (it *RegistryAddProtocolVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryAddProtocolVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryAddProtocolVersion represents a AddProtocolVersion event raised by the Registry contract.
type RegistryAddProtocolVersion struct {
	Version         *big.Int
	ProtocolAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddProtocolVersion is a free log retrieval operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_Registry *RegistryFilterer) FilterAddProtocolVersion(opts *bind.FilterOpts, version []*big.Int, protocolAddress []common.Address) (*RegistryAddProtocolVersionIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "AddProtocolVersion", versionRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return &RegistryAddProtocolVersionIterator{contract: _Registry.contract, event: "AddProtocolVersion", logs: logs, sub: sub}, nil
}

// WatchAddProtocolVersion is a free log subscription operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_Registry *RegistryFilterer) WatchAddProtocolVersion(opts *bind.WatchOpts, sink chan<- *RegistryAddProtocolVersion, version []*big.Int, protocolAddress []common.Address) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "AddProtocolVersion", versionRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryAddProtocolVersion)
				if err := _Registry.contract.UnpackLog(event, "AddProtocolVersion", log); err != nil {
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

// ParseAddProtocolVersion is a log parse operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_Registry *RegistryFilterer) ParseAddProtocolVersion(log types.Log) (*RegistryAddProtocolVersion, error) {
	event := new(RegistryAddProtocolVersion)
	if err := _Registry.contract.UnpackLog(event, "AddProtocolVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryLatestVersionUpdatedIterator is returned from FilterLatestVersionUpdated and is used to iterate over the raw logs and unpacked data for LatestVersionUpdated events raised by the Registry contract.
type RegistryLatestVersionUpdatedIterator struct {
	Event *RegistryLatestVersionUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryLatestVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryLatestVersionUpdated)
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
		it.Event = new(RegistryLatestVersionUpdated)
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
func (it *RegistryLatestVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryLatestVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryLatestVersionUpdated represents a LatestVersionUpdated event raised by the Registry contract.
type RegistryLatestVersionUpdated struct {
	OldVersion *big.Int
	NewVersion *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLatestVersionUpdated is a free log retrieval operation binding the contract event 0x30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b0.
//
// Solidity: event LatestVersionUpdated(uint256 indexed oldVersion, uint256 indexed newVersion)
func (_Registry *RegistryFilterer) FilterLatestVersionUpdated(opts *bind.FilterOpts, oldVersion []*big.Int, newVersion []*big.Int) (*RegistryLatestVersionUpdatedIterator, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "LatestVersionUpdated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return &RegistryLatestVersionUpdatedIterator{contract: _Registry.contract, event: "LatestVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchLatestVersionUpdated is a free log subscription operation binding the contract event 0x30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b0.
//
// Solidity: event LatestVersionUpdated(uint256 indexed oldVersion, uint256 indexed newVersion)
func (_Registry *RegistryFilterer) WatchLatestVersionUpdated(opts *bind.WatchOpts, sink chan<- *RegistryLatestVersionUpdated, oldVersion []*big.Int, newVersion []*big.Int) (event.Subscription, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "LatestVersionUpdated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryLatestVersionUpdated)
				if err := _Registry.contract.UnpackLog(event, "LatestVersionUpdated", log); err != nil {
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

// ParseLatestVersionUpdated is a log parse operation binding the contract event 0x30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b0.
//
// Solidity: event LatestVersionUpdated(uint256 indexed oldVersion, uint256 indexed newVersion)
func (_Registry *RegistryFilterer) ParseLatestVersionUpdated(log types.Log) (*RegistryLatestVersionUpdated, error) {
	event := new(RegistryLatestVersionUpdated)
	if err := _Registry.contract.UnpackLog(event, "LatestVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
