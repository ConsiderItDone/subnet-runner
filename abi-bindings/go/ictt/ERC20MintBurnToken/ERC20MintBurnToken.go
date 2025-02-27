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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousHome\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newHome\",\"type\":\"address\"}],\"name\":\"HomeAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"homeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newHomeAddress\",\"type\":\"address\"}],\"name\":\"setHomeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610f3f380380610f3f83398101604081905261002e916101b1565b808484600361003d83826102c9565b50600461004a82826102c9565b5050506001600160a01b03811661007a57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b6100838161008d565b5050505050610388565b600680546001600160a01b03191690556100a6816100a9565b50565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261011d575f80fd5b81516001600160401b0380821115610137576101376100fa565b604051601f8301601f19908116603f0116810190828211818310171561015f5761015f6100fa565b816040528381526020925086602085880101111561017b575f80fd5b5f91505b8382101561019c578582018301518183018401529082019061017f565b5f602085830101528094505050505092915050565b5f805f80608085870312156101c4575f80fd5b84516001600160401b03808211156101da575f80fd5b6101e68883890161010e565b955060208701519150808211156101fb575f80fd5b506102088782880161010e565b935050604085015160ff8116811461021e575f80fd5b60608601519092506001600160a01b038116811461023a575f80fd5b939692955090935050565b600181811c9082168061025957607f821691505b60208210810361027757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156102c457805f5260205f20601f840160051c810160208510156102a25750805b601f840160051c820191505b818110156102c1575f81556001016102ae565b50505b505050565b81516001600160401b038111156102e2576102e26100fa565b6102f6816102f08454610245565b8461027d565b602080601f831160018114610329575f84156103125750858301515b5f19600386901b1c1916600185901b178555610380565b5f85815260208120601f198616915b8281101561035757888601518255948401946001909101908401610338565b508582101561037457878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b610baa806103955f395ff3fe608060405234801561000f575f80fd5b506004361061011c575f3560e01c806379ba5097116100a9578063c9072c381161006e578063c9072c3814610248578063dd62ed3e1461025b578063e30c397814610293578063f2fde38b146102a4578063f629ad30146102b7575f80fd5b806379ba5097146101ed57806379cc6790146101f55780638da5cb5b1461020857806395d89b411461022d578063a9059cbb14610235575f80fd5b8063313ce567116100ef578063313ce5671461018657806340c10f191461019557806342966c68146101a857806370a08231146101bb578063715018a6146101e3575f80fd5b806306fdde0314610120578063095ea7b31461013e57806318160ddd1461016157806323b872dd14610173575b5f80fd5b6101286102ca565b60405161013591906109c7565b60405180910390f35b61015161014c366004610a29565b61035a565b6040519015158152602001610135565b6002545b604051908152602001610135565b610151610181366004610a51565b610373565b60405160128152602001610135565b6101516101a3366004610a29565b610396565b6101516101b6366004610a8a565b6103de565b6101656101c9366004610aa1565b6001600160a01b03165f9081526020819052604090205490565b6101eb61041d565b005b6101eb610430565b6101eb610203366004610a29565b610474565b6005546001600160a01b03165b6040516001600160a01b039091168152602001610135565b6101286104ac565b610151610243366004610a29565b6104bb565b6101eb610256366004610aa1565b6104c8565b610165610269366004610ac1565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6006546001600160a01b0316610215565b6101eb6102b2366004610aa1565b610575565b600754610215906001600160a01b031681565b6060600380546102d990610af2565b80601f016020809104026020016040519081016040528092919081815260200182805461030590610af2565b80156103505780601f1061032757610100808354040283529160200191610350565b820191905f5260205f20905b81548152906001019060200180831161033357829003601f168201915b5050505050905090565b5f336103678185856105e6565b60019150505b92915050565b5f336103808582856105f8565b61038b858585610673565b506001949350505050565b6007545f906001600160a01b031633146103cb5760405162461bcd60e51b81526004016103c290610b2a565b60405180910390fd5b6103d583836106d0565b50600192915050565b6007545f906001600160a01b0316331461040a5760405162461bcd60e51b81526004016103c290610b2a565b6104143383610704565b5060015b919050565b610425610738565b61042e5f610765565b565b60065433906001600160a01b031681146104685760405163118cdaa760e01b81526001600160a01b03821660048201526024016103c2565b61047181610765565b50565b6007546001600160a01b0316331461049e5760405162461bcd60e51b81526004016103c290610b2a565b6104a88282610704565b5050565b6060600480546102d990610af2565b5f33610367818585610673565b6104d0610738565b6001600160a01b03811661051a5760405162461bcd60e51b81526020600482015260116024820152705a65726f20686f6d65206164647265737360781b60448201526064016103c2565b6007546040516001600160a01b038084169216907ff0f605335e9213619cc50a0fec453435f6eb65e0e181cf11eaf362b938e01a6e905f90a3600780546001600160a01b0319166001600160a01b0392909216919091179055565b61057d610738565b600680546001600160a01b0383166001600160a01b031990911681179091556105ae6005546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6105f3838383600161077e565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f19811461066d578181101561065f57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016103c2565b61066d84848484035f61077e565b50505050565b6001600160a01b03831661069c57604051634b637e8f60e11b81525f60048201526024016103c2565b6001600160a01b0382166106c55760405163ec442f0560e01b81525f60048201526024016103c2565b6105f3838383610850565b6001600160a01b0382166106f95760405163ec442f0560e01b81525f60048201526024016103c2565b6104a85f8383610850565b6001600160a01b03821661072d57604051634b637e8f60e11b81525f60048201526024016103c2565b6104a8825f83610850565b6005546001600160a01b0316331461042e5760405163118cdaa760e01b81523360048201526024016103c2565b600680546001600160a01b031916905561047181610976565b6001600160a01b0384166107a75760405163e602df0560e01b81525f60048201526024016103c2565b6001600160a01b0383166107d057604051634a1406b160e11b81525f60048201526024016103c2565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561066d57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161084291815260200190565b60405180910390a350505050565b6001600160a01b03831661087a578060025f82825461086f9190610b55565b909155506108ea9050565b6001600160a01b0383165f90815260208190526040902054818110156108cc5760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016103c2565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661090657600280548290039055610924565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161096991815260200190565b60405180910390a3505050565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f602080835283518060208501525f5b818110156109f3578581018301518582016040015282016109d7565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610418575f80fd5b5f8060408385031215610a3a575f80fd5b610a4383610a13565b946020939093013593505050565b5f805f60608486031215610a63575f80fd5b610a6c84610a13565b9250610a7a60208501610a13565b9150604084013590509250925092565b5f60208284031215610a9a575f80fd5b5035919050565b5f60208284031215610ab1575f80fd5b610aba82610a13565b9392505050565b5f8060408385031215610ad2575f80fd5b610adb83610a13565b9150610ae960208401610a13565b90509250929050565b600181811c90821680610b0657607f821691505b602082108103610b2457634e487b7160e01b5f52602260045260245ffd5b50919050565b602080825260119082015270139bdd081a1bdb594818dbdb9d1c9858dd607a1b604082015260600190565b8082018082111561036d57634e487b7160e01b5f52601160045260245ffdfea26469706673582212207808e55a7aea15eb1f8160a76e406a7af186a92df915f455bce724a0e332e74a64736f6c63430008190033",
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

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20MintBurnToken.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) PendingOwner() (common.Address, error) {
	return _ERC20MintBurnToken.Contract.PendingOwner(&_ERC20MintBurnToken.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ERC20MintBurnToken *ERC20MintBurnTokenCallerSession) PendingOwner() (common.Address, error) {
	return _ERC20MintBurnToken.Contract.PendingOwner(&_ERC20MintBurnToken.CallOpts)
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

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20MintBurnToken.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.AcceptOwnership(&_ERC20MintBurnToken.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ERC20MintBurnToken *ERC20MintBurnTokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ERC20MintBurnToken.Contract.AcceptOwnership(&_ERC20MintBurnToken.TransactOpts)
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

// ERC20MintBurnTokenOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the ERC20MintBurnToken contract.
type ERC20MintBurnTokenOwnershipTransferStartedIterator struct {
	Event *ERC20MintBurnTokenOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *ERC20MintBurnTokenOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MintBurnTokenOwnershipTransferStarted)
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
		it.Event = new(ERC20MintBurnTokenOwnershipTransferStarted)
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
func (it *ERC20MintBurnTokenOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MintBurnTokenOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MintBurnTokenOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the ERC20MintBurnToken contract.
type ERC20MintBurnTokenOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20MintBurnTokenOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20MintBurnToken.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MintBurnTokenOwnershipTransferStartedIterator{contract: _ERC20MintBurnToken.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *ERC20MintBurnTokenOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20MintBurnToken.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MintBurnTokenOwnershipTransferStarted)
				if err := _ERC20MintBurnToken.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ERC20MintBurnToken *ERC20MintBurnTokenFilterer) ParseOwnershipTransferStarted(log types.Log) (*ERC20MintBurnTokenOwnershipTransferStarted, error) {
	event := new(ERC20MintBurnTokenOwnershipTransferStarted)
	if err := _ERC20MintBurnToken.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
