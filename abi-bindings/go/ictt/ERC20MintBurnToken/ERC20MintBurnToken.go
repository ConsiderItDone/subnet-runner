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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousHome\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newHome\",\"type\":\"address\"}],\"name\":\"HomeAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"homeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newHomeAddress\",\"type\":\"address\"}],\"name\":\"setHomeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610eb3380380610eb383398101604081905261002f9161019f565b808484600361003e83826102c5565b50600461004b82826102c5565b5050506001600160a01b03811661007c57604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b6100858161008f565b5050505050610384565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261010857600080fd5b81516001600160401b0380821115610122576101226100e1565b604051601f8301601f19908116603f0116810190828211818310171561014a5761014a6100e1565b816040528381526020925086602085880101111561016757600080fd5b600091505b83821015610189578582018301518183018401529082019061016c565b6000602085830101528094505050505092915050565b600080600080608085870312156101b557600080fd5b84516001600160401b03808211156101cc57600080fd5b6101d8888389016100f7565b955060208701519150808211156101ee57600080fd5b506101fb878288016100f7565b935050604085015160ff8116811461021257600080fd5b60608601519092506001600160a01b038116811461022f57600080fd5b939692955090935050565b600181811c9082168061024e57607f821691505b60208210810361026e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156102c0576000816000526020600020601f850160051c8101602086101561029d5750805b601f850160051c820191505b818110156102bc578281556001016102a9565b5050505b505050565b81516001600160401b038111156102de576102de6100e1565b6102f2816102ec845461023a565b84610274565b602080601f831160018114610327576000841561030f5750858301515b600019600386901b1c1916600185901b1785556102bc565b600085815260208120601f198616915b8281101561035657888601518255948401946001909101908401610337565b50858210156103745787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610b20806103936000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063715018a6116100a2578063a9059cbb11610071578063a9059cbb1461021e578063c9072c3814610231578063dd62ed3e14610244578063f2fde38b1461027d578063f629ad301461029057600080fd5b8063715018a6146101d457806379cc6790146101de5780638da5cb5b146101f157806395d89b411461021657600080fd5b8063313ce567116100de578063313ce5671461017657806340c10f191461018557806342966c681461019857806370a08231146101ab57600080fd5b806306fdde0314610110578063095ea7b31461012e57806318160ddd1461015157806323b872dd14610163575b600080fd5b6101186102a3565b604051610125919061092a565b60405180910390f35b61014161013c366004610990565b610335565b6040519015158152602001610125565b6002545b604051908152602001610125565b6101416101713660046109ba565b61034f565b60405160128152602001610125565b610141610193366004610990565b610373565b6101416101a63660046109f6565b6103bc565b6101556101b9366004610a0f565b6001600160a01b031660009081526020819052604090205490565b6101dc6103fc565b005b6101dc6101ec366004610990565b610410565b6005546001600160a01b03165b6040516001600160a01b039091168152602001610125565b610118610448565b61014161022c366004610990565b610457565b6101dc61023f366004610a0f565b610465565b610155610252366004610a31565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6101dc61028b366004610a0f565b610513565b6006546101fe906001600160a01b031681565b6060600380546102b290610a64565b80601f01602080910402602001604051908101604052809291908181526020018280546102de90610a64565b801561032b5780601f106103005761010080835404028352916020019161032b565b820191906000526020600020905b81548152906001019060200180831161030e57829003601f168201915b5050505050905090565b600033610343818585610551565b60019150505b92915050565b60003361035d858285610563565b6103688585856105e1565b506001949350505050565b6006546000906001600160a01b031633146103a95760405162461bcd60e51b81526004016103a090610a9e565b60405180910390fd5b6103b38383610640565b50600192915050565b6006546000906001600160a01b031633146103e95760405162461bcd60e51b81526004016103a090610a9e565b6103f33383610676565b5060015b919050565b6104046106ac565b61040e60006106d9565b565b6006546001600160a01b0316331461043a5760405162461bcd60e51b81526004016103a090610a9e565b6104448282610676565b5050565b6060600480546102b290610a64565b6000336103438185856105e1565b61046d6106ac565b6001600160a01b0381166104b75760405162461bcd60e51b81526020600482015260116024820152705a65726f20686f6d65206164647265737360781b60448201526064016103a0565b6006546040516001600160a01b038084169216907ff0f605335e9213619cc50a0fec453435f6eb65e0e181cf11eaf362b938e01a6e90600090a3600680546001600160a01b0319166001600160a01b0392909216919091179055565b61051b6106ac565b6001600160a01b03811661054557604051631e4fbdf760e01b8152600060048201526024016103a0565b61054e816106d9565b50565b61055e838383600161072b565b505050565b6001600160a01b0383811660009081526001602090815260408083209386168352929052205460001981146105db57818110156105cc57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016103a0565b6105db8484848403600061072b565b50505050565b6001600160a01b03831661060b57604051634b637e8f60e11b8152600060048201526024016103a0565b6001600160a01b0382166106355760405163ec442f0560e01b8152600060048201526024016103a0565b61055e838383610800565b6001600160a01b03821661066a5760405163ec442f0560e01b8152600060048201526024016103a0565b61044460008383610800565b6001600160a01b0382166106a057604051634b637e8f60e11b8152600060048201526024016103a0565b61044482600083610800565b6005546001600160a01b0316331461040e5760405163118cdaa760e01b81523360048201526024016103a0565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b0384166107555760405163e602df0560e01b8152600060048201526024016103a0565b6001600160a01b03831661077f57604051634a1406b160e11b8152600060048201526024016103a0565b6001600160a01b03808516600090815260016020908152604080832093871683529290522082905580156105db57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516107f291815260200190565b60405180910390a350505050565b6001600160a01b03831661082b5780600260008282546108209190610ac9565b9091555061089d9050565b6001600160a01b0383166000908152602081905260409020548181101561087e5760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016103a0565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b0382166108b9576002805482900390556108d8565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161091d91815260200190565b60405180910390a3505050565b60006020808352835180602085015260005b818110156109585785810183015185820160400152820161093c565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146103f757600080fd5b600080604083850312156109a357600080fd5b6109ac83610979565b946020939093013593505050565b6000806000606084860312156109cf57600080fd5b6109d884610979565b92506109e660208501610979565b9150604084013590509250925092565b600060208284031215610a0857600080fd5b5035919050565b600060208284031215610a2157600080fd5b610a2a82610979565b9392505050565b60008060408385031215610a4457600080fd5b610a4d83610979565b9150610a5b60208401610979565b90509250929050565b600181811c90821680610a7857607f821691505b602082108103610a9857634e487b7160e01b600052602260045260246000fd5b50919050565b602080825260119082015270139bdd081a1bdb594818dbdb9d1c9858dd607a1b604082015260600190565b8082018082111561034957634e487b7160e01b600052601160045260246000fdfea2646970667358221220206167bab93b85181a12f22c2e27e318fb0ef1ecdc2e22e4019a7af18ebefd1364736f6c63430008190033",
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
