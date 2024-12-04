// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20mintburntoken

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

// ERC20MintBurnTokenMetaData contains all meta data concerning the ERC20MintBurnToken contract.
var ERC20MintBurnTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"homeAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHomeAddress\",\"inputs\":[{\"name\":\"newHomeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HomeAddressSet\",\"inputs\":[{\"name\":\"previousHome\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newHome\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610e3a380380610e3a83398101604081905261002e91610195565b808484600361003d83826102ad565b50600461004a82826102ad565b5050506001600160a01b03811661007a57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b6100838161008d565b505050505061036c565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610101575f80fd5b81516001600160401b038082111561011b5761011b6100de565b604051601f8301601f19908116603f01168101908282118183101715610143576101436100de565b816040528381526020925086602085880101111561015f575f80fd5b5f91505b838210156101805785820183015181830184015290820190610163565b5f602085830101528094505050505092915050565b5f805f80608085870312156101a8575f80fd5b84516001600160401b03808211156101be575f80fd5b6101ca888389016100f2565b955060208701519150808211156101df575f80fd5b506101ec878288016100f2565b935050604085015160ff81168114610202575f80fd5b60608601519092506001600160a01b038116811461021e575f80fd5b939692955090935050565b600181811c9082168061023d57607f821691505b60208210810361025b57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156102a857805f5260205f20601f840160051c810160208510156102865750805b601f840160051c820191505b818110156102a5575f8155600101610292565b50505b505050565b81516001600160401b038111156102c6576102c66100de565b6102da816102d48454610229565b84610261565b602080601f83116001811461030d575f84156102f65750858301515b5f19600386901b1c1916600185901b178555610364565b5f85815260208120601f198616915b8281101561033b5788860151825594840194600190910190840161031c565b508582101561035857878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b610ac1806103795f395ff3fe608060405234801561000f575f80fd5b5060043610610106575f3560e01c8063715018a61161009e578063a9059cbb1161006e578063a9059cbb14610217578063c9072c381461022a578063dd62ed3e1461023d578063f2fde38b14610275578063f629ad3014610288575f80fd5b8063715018a6146101cd57806379cc6790146101d75780638da5cb5b146101ea57806395d89b411461020f575f80fd5b8063313ce567116100d9578063313ce5671461017057806340c10f191461017f57806342966c681461019257806370a08231146101a5575f80fd5b806306fdde031461010a578063095ea7b31461012857806318160ddd1461014b57806323b872dd1461015d575b5f80fd5b61011261029b565b60405161011f9190610907565b60405180910390f35b61013b610136366004610969565b61032b565b604051901515815260200161011f565b6002545b60405190815260200161011f565b61013b61016b366004610991565b610344565b6040516012815260200161011f565b61013b61018d366004610969565b610367565b61013b6101a03660046109ca565b6103af565b61014f6101b33660046109e1565b6001600160a01b03165f9081526020819052604090205490565b6101d56103ee565b005b6101d56101e5366004610969565b610401565b6005546001600160a01b03165b6040516001600160a01b03909116815260200161011f565b610112610439565b61013b610225366004610969565b610448565b6101d56102383660046109e1565b610455565b61014f61024b366004610a01565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6101d56102833660046109e1565b610502565b6006546101f7906001600160a01b031681565b6060600380546102aa90610a32565b80601f01602080910402602001604051908101604052809291908181526020018280546102d690610a32565b80156103215780601f106102f857610100808354040283529160200191610321565b820191905f5260205f20905b81548152906001019060200180831161030457829003601f168201915b5050505050905090565b5f3361033881858561053f565b60019150505b92915050565b5f33610351858285610551565b61035c8585856105cc565b506001949350505050565b6006545f906001600160a01b0316331461039c5760405162461bcd60e51b815260040161039390610a6a565b60405180910390fd5b6103a68383610629565b50600192915050565b6006545f906001600160a01b031633146103db5760405162461bcd60e51b815260040161039390610a6a565b6103e5338361065d565b5060015b919050565b6103f6610691565b6103ff5f6106be565b565b6006546001600160a01b0316331461042b5760405162461bcd60e51b815260040161039390610a6a565b610435828261065d565b5050565b6060600480546102aa90610a32565b5f336103388185856105cc565b61045d610691565b6001600160a01b0381166104a75760405162461bcd60e51b81526020600482015260116024820152705a65726f20686f6d65206164647265737360781b6044820152606401610393565b6006546040516001600160a01b038084169216907ff0f605335e9213619cc50a0fec453435f6eb65e0e181cf11eaf362b938e01a6e905f90a3600680546001600160a01b0319166001600160a01b0392909216919091179055565b61050a610691565b6001600160a01b03811661053357604051631e4fbdf760e01b81525f6004820152602401610393565b61053c816106be565b50565b61054c838383600161070f565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f1981146105c657818110156105b857604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610393565b6105c684848484035f61070f565b50505050565b6001600160a01b0383166105f557604051634b637e8f60e11b81525f6004820152602401610393565b6001600160a01b03821661061e5760405163ec442f0560e01b81525f6004820152602401610393565b61054c8383836107e1565b6001600160a01b0382166106525760405163ec442f0560e01b81525f6004820152602401610393565b6104355f83836107e1565b6001600160a01b03821661068657604051634b637e8f60e11b81525f6004820152602401610393565b610435825f836107e1565b6005546001600160a01b031633146103ff5760405163118cdaa760e01b8152336004820152602401610393565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6001600160a01b0384166107385760405163e602df0560e01b81525f6004820152602401610393565b6001600160a01b03831661076157604051634a1406b160e11b81525f6004820152602401610393565b6001600160a01b038085165f90815260016020908152604080832093871683529290522082905580156105c657826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516107d391815260200190565b60405180910390a350505050565b6001600160a01b03831661080b578060025f8282546108009190610a95565b9091555061087b9050565b6001600160a01b0383165f908152602081905260409020548181101561085d5760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610393565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b038216610897576002805482900390556108b5565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516108fa91815260200190565b60405180910390a3505050565b5f602080835283518060208501525f5b8181101561093357858101830151858201604001528201610917565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146103e9575f80fd5b5f806040838503121561097a575f80fd5b61098383610953565b946020939093013593505050565b5f805f606084860312156109a3575f80fd5b6109ac84610953565b92506109ba60208501610953565b9150604084013590509250925092565b5f602082840312156109da575f80fd5b5035919050565b5f602082840312156109f1575f80fd5b6109fa82610953565b9392505050565b5f8060408385031215610a12575f80fd5b610a1b83610953565b9150610a2960208401610953565b90509250929050565b600181811c90821680610a4657607f821691505b602082108103610a6457634e487b7160e01b5f52602260045260245ffd5b50919050565b602080825260119082015270139bdd081a1bdb594818dbdb9d1c9858dd607a1b604082015260600190565b8082018082111561033e57634e487b7160e01b5f52601160045260245ffdfea164736f6c6343000819000a",
}

// ERC20MintBurnTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MintBurnTokenMetaData.ABI instead.
var ERC20MintBurnTokenABI = ERC20MintBurnTokenMetaData.ABI

// ERC20MintBurnTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20MintBurnTokenMetaData.Bin instead.
var ERC20MintBurnTokenBin = ERC20MintBurnTokenMetaData.Bin

// DeployERC20MintBurnToken deploys a new Ethereum contract, binding an instance of ERC20MintBurnToken to it.
func DeployERC20MintBurnToken(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals_ uint8, initialOwner common.Address) (common.Address, *types.Transaction, *ERC20MintBurnToken, error) {
	parsed, err := ERC20MintBurnTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20MintBurnTokenBin), backend, name, symbol, decimals_, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20MintBurnToken{ERC20MintBurnTokenCaller: ERC20MintBurnTokenCaller{contract: contract}, ERC20MintBurnTokenTransactor: ERC20MintBurnTokenTransactor{contract: contract}, ERC20MintBurnTokenFilterer: ERC20MintBurnTokenFilterer{contract: contract}}, nil
}

// ERC20MintBurnToken is an auto generated Go binding around an Ethereum contract.
type ERC20MintBurnToken struct {
	ERC20MintBurnTokenCaller     // Read-only binding to the contract
	ERC20MintBurnTokenTransactor // Write-only binding to the contract
	ERC20MintBurnTokenFilterer   // Log filterer for contract events
}

// ERC20MintBurnTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20MintBurnTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20MintBurnTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20MintBurnTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20MintBurnTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20MintBurnTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20MintBurnTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20MintBurnTokenSession struct {
	Contract     *ERC20MintBurnToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ERC20MintBurnTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20MintBurnTokenCallerSession struct {
	Contract *ERC20MintBurnTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ERC20MintBurnTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20MintBurnTokenTransactorSession struct {
	Contract     *ERC20MintBurnTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ERC20MintBurnTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20MintBurnTokenRaw struct {
	Contract *ERC20MintBurnToken // Generic contract binding to access the raw methods on
}

// ERC20MintBurnTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20MintBurnTokenCallerRaw struct {
	Contract *ERC20MintBurnTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20MintBurnTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20MintBurnTokenTransactorRaw struct {
	Contract *ERC20MintBurnTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20MintBurnToken creates a new instance of ERC20MintBurnToken, bound to a specific deployed contract.
func NewERC20MintBurnToken(address common.Address, backend bind.ContractBackend) (*ERC20MintBurnToken, error) {
	contract, err := bindERC20MintBurnToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20MintBurnToken{ERC20MintBurnTokenCaller: ERC20MintBurnTokenCaller{contract: contract}, ERC20MintBurnTokenTransactor: ERC20MintBurnTokenTransactor{contract: contract}, ERC20MintBurnTokenFilterer: ERC20MintBurnTokenFilterer{contract: contract}}, nil
}

// NewERC20MintBurnTokenCaller creates a new read-only instance of ERC20MintBurnToken, bound to a specific deployed contract.
func NewERC20MintBurnTokenCaller(address common.Address, caller bind.ContractCaller) (*ERC20MintBurnTokenCaller, error) {
	contract, err := bindERC20MintBurnToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20MintBurnTokenCaller{contract: contract}, nil
}

// NewERC20MintBurnTokenTransactor creates a new write-only instance of ERC20MintBurnToken, bound to a specific deployed contract.
func NewERC20MintBurnTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20MintBurnTokenTransactor, error) {
	contract, err := bindERC20MintBurnToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20MintBurnTokenTransactor{contract: contract}, nil
}

// NewERC20MintBurnTokenFilterer creates a new log filterer instance of ERC20MintBurnToken, bound to a specific deployed contract.
func NewERC20MintBurnTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20MintBurnTokenFilterer, error) {
	contract, err := bindERC20MintBurnToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20MintBurnTokenFilterer{contract: contract}, nil
}

// bindERC20MintBurnToken binds a generic wrapper to an already deployed contract.
func bindERC20MintBurnToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20MintBurnTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20MintBurnToken *ERC20MintBurnTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20MintBurnToken.Contract.ERC20MintBurnTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20MintBurnToken *ERC20MintBurnTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.ERC20MintBurnTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20MintBurnToken *ERC20MintBurnTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.ERC20MintBurnTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20MintBurnToken *ERC20MintBurnTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20MintBurnToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20MintBurnToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20MintBurnToken.Contract.Allowance(&_ERC20MintBurnToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20MintBurnToken.Contract.Allowance(&_ERC20MintBurnToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20MintBurnToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20MintBurnToken.Contract.BalanceOf(&_ERC20MintBurnToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20MintBurnToken.Contract.BalanceOf(&_ERC20MintBurnToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20MintBurnToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) Decimals() (uint8, error) {
	return _ERC20MintBurnToken.Contract.Decimals(&_ERC20MintBurnToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCallerSession) Decimals() (uint8, error) {
	return _ERC20MintBurnToken.Contract.Decimals(&_ERC20MintBurnToken.CallOpts)
}

// HomeAddress is a free data retrieval call binding the contract method 0xf629ad30.
//
// Solidity: function homeAddress() view returns(address)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCaller) HomeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20MintBurnToken.contract.Call(opts, &out, "homeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HomeAddress is a free data retrieval call binding the contract method 0xf629ad30.
//
// Solidity: function homeAddress() view returns(address)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) HomeAddress() (common.Address, error) {
	return _ERC20MintBurnToken.Contract.HomeAddress(&_ERC20MintBurnToken.CallOpts)
}

// HomeAddress is a free data retrieval call binding the contract method 0xf629ad30.
//
// Solidity: function homeAddress() view returns(address)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCallerSession) HomeAddress() (common.Address, error) {
	return _ERC20MintBurnToken.Contract.HomeAddress(&_ERC20MintBurnToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20MintBurnToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) Name() (string, error) {
	return _ERC20MintBurnToken.Contract.Name(&_ERC20MintBurnToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCallerSession) Name() (string, error) {
	return _ERC20MintBurnToken.Contract.Name(&_ERC20MintBurnToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20MintBurnToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) Owner() (common.Address, error) {
	return _ERC20MintBurnToken.Contract.Owner(&_ERC20MintBurnToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCallerSession) Owner() (common.Address, error) {
	return _ERC20MintBurnToken.Contract.Owner(&_ERC20MintBurnToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20MintBurnToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) Symbol() (string, error) {
	return _ERC20MintBurnToken.Contract.Symbol(&_ERC20MintBurnToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCallerSession) Symbol() (string, error) {
	return _ERC20MintBurnToken.Contract.Symbol(&_ERC20MintBurnToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20MintBurnToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) TotalSupply() (*big.Int, error) {
	return _ERC20MintBurnToken.Contract.TotalSupply(&_ERC20MintBurnToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20MintBurnToken.Contract.TotalSupply(&_ERC20MintBurnToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.Approve(&_ERC20MintBurnToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.Approve(&_ERC20MintBurnToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.Burn(&_ERC20MintBurnToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.Burn(&_ERC20MintBurnToken.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.BurnFrom(&_ERC20MintBurnToken.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.BurnFrom(&_ERC20MintBurnToken.TransactOpts, account, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.Mint(&_ERC20MintBurnToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.Mint(&_ERC20MintBurnToken.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20MintBurnToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.RenounceOwnership(&_ERC20MintBurnToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.RenounceOwnership(&_ERC20MintBurnToken.TransactOpts)
}

// SetHomeAddress is a paid mutator transaction binding the contract method 0xc9072c38.
//
// Solidity: function setHomeAddress(address newHomeAddress) returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactor) SetHomeAddress(opts *bind.TransactOpts, newHomeAddress common.Address) (*types.Transaction, error) {
	return _ERC20MintBurnToken.contract.Transact(opts, "setHomeAddress", newHomeAddress)
}

// SetHomeAddress is a paid mutator transaction binding the contract method 0xc9072c38.
//
// Solidity: function setHomeAddress(address newHomeAddress) returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) SetHomeAddress(newHomeAddress common.Address) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.SetHomeAddress(&_ERC20MintBurnToken.TransactOpts, newHomeAddress)
}

// SetHomeAddress is a paid mutator transaction binding the contract method 0xc9072c38.
//
// Solidity: function setHomeAddress(address newHomeAddress) returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactorSession) SetHomeAddress(newHomeAddress common.Address) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.SetHomeAddress(&_ERC20MintBurnToken.TransactOpts, newHomeAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.Transfer(&_ERC20MintBurnToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.Transfer(&_ERC20MintBurnToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.TransferFrom(&_ERC20MintBurnToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.TransferFrom(&_ERC20MintBurnToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20MintBurnToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.TransferOwnership(&_ERC20MintBurnToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.TransferOwnership(&_ERC20MintBurnToken.TransactOpts, newOwner)
}

// ERC20MintBurnTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20MintBurnToken contract.
type ERC20MintBurnTokenApprovalIterator struct {
	Event *ERC20MintBurnTokenApproval // Event containing the contract specifics and raw log

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
func (it *ERC20MintBurnTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MintBurnTokenApproval)
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
		it.Event = new(ERC20MintBurnTokenApproval)
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
func (it *ERC20MintBurnTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MintBurnTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MintBurnTokenApproval represents a Approval event raised by the ERC20MintBurnToken contract.
type ERC20MintBurnTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20MintBurnTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20MintBurnToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MintBurnTokenApprovalIterator{contract: _ERC20MintBurnToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20MintBurnTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20MintBurnToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MintBurnTokenApproval)
				if err := _ERC20MintBurnToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) ParseApproval(log types.Log) (*ERC20MintBurnTokenApproval, error) {
	event := new(ERC20MintBurnTokenApproval)
	if err := _ERC20MintBurnToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20MintBurnTokenHomeAddressSetIterator is returned from FilterHomeAddressSet and is used to iterate over the raw logs and unpacked data for HomeAddressSet events raised by the ERC20MintBurnToken contract.
type ERC20MintBurnTokenHomeAddressSetIterator struct {
	Event *ERC20MintBurnTokenHomeAddressSet // Event containing the contract specifics and raw log

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
func (it *ERC20MintBurnTokenHomeAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MintBurnTokenHomeAddressSet)
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
		it.Event = new(ERC20MintBurnTokenHomeAddressSet)
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
func (it *ERC20MintBurnTokenHomeAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MintBurnTokenHomeAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MintBurnTokenHomeAddressSet represents a HomeAddressSet event raised by the ERC20MintBurnToken contract.
type ERC20MintBurnTokenHomeAddressSet struct {
	PreviousHome common.Address
	NewHome      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHomeAddressSet is a free log retrieval operation binding the contract event 0xf0f605335e9213619cc50a0fec453435f6eb65e0e181cf11eaf362b938e01a6e.
//
// Solidity: event HomeAddressSet(address indexed previousHome, address indexed newHome)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) FilterHomeAddressSet(opts *bind.FilterOpts, previousHome []common.Address, newHome []common.Address) (*ERC20MintBurnTokenHomeAddressSetIterator, error) {

	var previousHomeRule []interface{}
	for _, previousHomeItem := range previousHome {
		previousHomeRule = append(previousHomeRule, previousHomeItem)
	}
	var newHomeRule []interface{}
	for _, newHomeItem := range newHome {
		newHomeRule = append(newHomeRule, newHomeItem)
	}

	logs, sub, err := _ERC20MintBurnToken.contract.FilterLogs(opts, "HomeAddressSet", previousHomeRule, newHomeRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MintBurnTokenHomeAddressSetIterator{contract: _ERC20MintBurnToken.contract, event: "HomeAddressSet", logs: logs, sub: sub}, nil
}

// WatchHomeAddressSet is a free log subscription operation binding the contract event 0xf0f605335e9213619cc50a0fec453435f6eb65e0e181cf11eaf362b938e01a6e.
//
// Solidity: event HomeAddressSet(address indexed previousHome, address indexed newHome)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) WatchHomeAddressSet(opts *bind.WatchOpts, sink chan<- *ERC20MintBurnTokenHomeAddressSet, previousHome []common.Address, newHome []common.Address) (event.Subscription, error) {

	var previousHomeRule []interface{}
	for _, previousHomeItem := range previousHome {
		previousHomeRule = append(previousHomeRule, previousHomeItem)
	}
	var newHomeRule []interface{}
	for _, newHomeItem := range newHome {
		newHomeRule = append(newHomeRule, newHomeItem)
	}

	logs, sub, err := _ERC20MintBurnToken.contract.WatchLogs(opts, "HomeAddressSet", previousHomeRule, newHomeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MintBurnTokenHomeAddressSet)
				if err := _ERC20MintBurnToken.contract.UnpackLog(event, "HomeAddressSet", log); err != nil {
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

// ParseHomeAddressSet is a log parse operation binding the contract event 0xf0f605335e9213619cc50a0fec453435f6eb65e0e181cf11eaf362b938e01a6e.
//
// Solidity: event HomeAddressSet(address indexed previousHome, address indexed newHome)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) ParseHomeAddressSet(log types.Log) (*ERC20MintBurnTokenHomeAddressSet, error) {
	event := new(ERC20MintBurnTokenHomeAddressSet)
	if err := _ERC20MintBurnToken.contract.UnpackLog(event, "HomeAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20MintBurnTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20MintBurnToken contract.
type ERC20MintBurnTokenOwnershipTransferredIterator struct {
	Event *ERC20MintBurnTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20MintBurnTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MintBurnTokenOwnershipTransferred)
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
		it.Event = new(ERC20MintBurnTokenOwnershipTransferred)
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
func (it *ERC20MintBurnTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MintBurnTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MintBurnTokenOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20MintBurnToken contract.
type ERC20MintBurnTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20MintBurnTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20MintBurnToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MintBurnTokenOwnershipTransferredIterator{contract: _ERC20MintBurnToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20MintBurnTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20MintBurnToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MintBurnTokenOwnershipTransferred)
				if err := _ERC20MintBurnToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20MintBurnTokenOwnershipTransferred, error) {
	event := new(ERC20MintBurnTokenOwnershipTransferred)
	if err := _ERC20MintBurnToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20MintBurnTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20MintBurnToken contract.
type ERC20MintBurnTokenTransferIterator struct {
	Event *ERC20MintBurnTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20MintBurnTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MintBurnTokenTransfer)
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
		it.Event = new(ERC20MintBurnTokenTransfer)
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
func (it *ERC20MintBurnTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MintBurnTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MintBurnTokenTransfer represents a Transfer event raised by the ERC20MintBurnToken contract.
type ERC20MintBurnTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20MintBurnTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20MintBurnToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MintBurnTokenTransferIterator{contract: _ERC20MintBurnToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20MintBurnTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20MintBurnToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MintBurnTokenTransfer)
				if err := _ERC20MintBurnToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) ParseTransfer(log types.Log) (*ERC20MintBurnTokenTransfer, error) {
	event := new(ERC20MintBurnTokenTransfer)
	if err := _ERC20MintBurnToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
