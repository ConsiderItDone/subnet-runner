// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ics20bank

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

// ICS20BankMetaData contains all meta data concerning the ICS20Bank contract.
var ICS20BankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5060387fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177533603d565b5060e6565b6000828152602081815260408083206001600160a01b038516845290915281205460ff1660dc576000838152602081815260408083206001600160a01b03861684529091529020805460ff1916600117905560953390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600160e0565b5060005b92915050565b611366806100f56000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063b3ab15fb11610097578063d547741f11610066578063d547741f14610227578063f24dc1da1461023a578063f45346dc1461024d578063f5b541a61461026057600080fd5b8063b3ab15fb146101db578063b9b092c8146101ee578063ba7aef4314610201578063c45b71de1461021457600080fd5b806369328dec116100d357806369328dec1461018657806375b238fc1461019957806391d14854146101c0578063a217fddf146101d357600080fd5b806301ffc9a714610105578063248a9ca31461012d5780632f2ff15d1461015e57806336568abe14610173575b600080fd5b610118610113366004610eba565b610275565b60405190151581526020015b60405180910390f35b61015061013b366004610eeb565b60009081526020819052604090206001015490565b604051908152602001610124565b61017161016c366004610f20565b6102ac565b005b610171610181366004610f20565b6102d7565b610171610194366004610f4c565b61030f565b6101507fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b6101186101ce366004610f20565b610408565b610150600081565b6101716101e9366004610f88565b610431565b6101506101fc366004610fec565b6104d4565b61017161020f36600461103f565b610584565b61017161022236600461103f565b6105fa565b610171610235366004610f20565b610670565b610171610248366004611099565b610695565b61017161025b366004610f4c565b610919565b61015060008051602061131183398151915281565b60006001600160e01b03198216637965db0b60e01b14806102a657506301ffc9a760e01b6001600160e01b03198316145b92915050565b6000828152602081905260409020600101546102c781610a1d565b6102d18383610a2a565b50505050565b6001600160a01b03811633146103005760405163334bd91960e11b815260040160405180910390fd5b61030a8282610abc565b505050565b823b6103365760405162461bcd60e51b815260040161032d90611102565b60405180910390fd5b6103493361034385610b27565b84610b32565b60405163a9059cbb60e01b81526001600160a01b0382811660048301526024820184905284169063a9059cbb906044016020604051808303816000875af1158015610398573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103bc919061114c565b61030a5760405162461bcd60e51b815260206004820152601a60248201527f494353323042616e6b3a207472616e73666572206661696c6564000000000000604482015260640161032d565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b61045b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177533610408565b6104b85760405162461bcd60e51b815260206004820152602860248201527f6d75737420686176652061646d696e20726f6c6520746f20736574206e65772060448201526737b832b930ba37b960c11b606482015260840161032d565b6104d060008051602061131183398151915282610a2a565b5050565b60006001600160a01b0384166105425760405162461bcd60e51b815260206004820152602d60248201527f494353323042616e6b3a2062616c616e636520717565727920666f722074686560448201526c207a65726f206164647265737360981b606482015260840161032d565b6001838360405161055492919061116e565b908152604080519182900360209081019092206001600160a01b0387166000908152925290205490509392505050565b61059c60008051602061131183398151915233610408565b6105b85760405162461bcd60e51b815260040161032d9061117e565b6102d18484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250869250610c4d915050565b61061260008051602061131183398151915233610408565b61062e5760405162461bcd60e51b815260040161032d9061117e565b6102d18484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250869250610b32915050565b60008281526020819052604090206001015461068b81610a1d565b6102d18383610abc565b6001600160a01b0384166106fb5760405162461bcd60e51b815260206004820152602760248201527f494353323042616e6b3a207472616e7366657220746f20746865207a65726f206044820152666164647265737360c81b606482015260840161032d565b6001600160a01b038516331480610725575061072560008051602061131183398151915233610408565b6107855760405162461bcd60e51b815260206004820152602b60248201527f494353323042616e6b3a2063616c6c6572206973206e6f74206f776e6572206e60448201526a1bdc88185c1c1c9bdd995960aa1b606482015260840161032d565b60006001848460405161079992919061116e565b90815260408051602092819003830190206001600160a01b038916600090815292529020549050818110156108255760405162461bcd60e51b815260206004820152602c60248201527f494353323042616e6b3a20696e73756666696369656e742062616c616e63652060448201526b3337b9103a3930b739b332b960a11b606482015260840161032d565b61082f82826111dc565b6001858560405161084192919061116e565b9081526040805191829003602090810183206001600160a01b038b16600090815291522091909155829060019061087b908790879061116e565b90815260200160405180910390206000876001600160a01b03166001600160a01b0316815260200190815260200160002060008282546108bb91906111ef565b92505081905550846001600160a01b0316866001600160a01b03167f1d30d3db8e01fa0d5626c471596f822f597e720c26a2930ef20d3387313c3d7886868660405161090993929190611202565b60405180910390a3505050505050565b823b6109375760405162461bcd60e51b815260040161032d90611102565b6001600160a01b0383166323b872dd336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152306024820152604481018590526064016020604051808303816000875af115801561099a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109be919061114c565b610a0a5760405162461bcd60e51b815260206004820152601a60248201527f494353323042616e6b3a207472616e73666572206661696c6564000000000000604482015260640161032d565b61030a81610a1785610b27565b84610c4d565b610a278133610cf0565b50565b6000610a368383610408565b610ab4576000838152602081815260408083206001600160a01b03861684529091529020805460ff19166001179055610a6c3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016102a6565b5060006102a6565b6000610ac88383610408565b15610ab4576000838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016102a6565b60606102a682610d29565b6000600183604051610b44919061125f565b90815260408051602092819003830190206001600160a01b03871660009081529252902054905081811015610bca5760405162461bcd60e51b815260206004820152602660248201527f494353323042616e6b3a206275726e20616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161032d565b610bd482826111dc565b600184604051610be4919061125f565b90815260408051602092819003830181206001600160a01b03891660008181529190945291822093909355917f1d30d3db8e01fa0d5626c471596f822f597e720c26a2930ef20d3387313c3d7890610c3f908790879061127b565b60405180910390a350505050565b80600183604051610c5e919061125f565b90815260200160405180910390206000856001600160a01b03166001600160a01b031681526020019081526020016000206000828254610c9e91906111ef565b90915550506040516001600160a01b038416906000907f1d30d3db8e01fa0d5626c471596f822f597e720c26a2930ef20d3387313c3d7890610ce3908690869061127b565b60405180910390a3505050565b610cfa8282610408565b6104d05760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440161032d565b60606102a66001600160a01b03831660146060826000610d4a8460026112b6565b610d559060026111ef565b67ffffffffffffffff811115610d6d57610d6d6112cd565b6040519080825280601f01601f191660200182016040528015610d97576020820181803683370190505b509050600360fc1b81600081518110610db257610db26112e3565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110610de157610de16112e3565b60200101906001600160f81b031916908160001a9053506000610e058560026112b6565b610e109060016111ef565b90505b6001811115610e88576f181899199a1a9b1b9c1cb0b131b232b360811b83600f1660108110610e4457610e446112e3565b1a60f81b828281518110610e5a57610e5a6112e3565b60200101906001600160f81b031916908160001a90535060049290921c91610e81816112f9565b9050610e13565b508115610eb25760405163e22e27eb60e01b8152600481018690526024810185905260440161032d565b949350505050565b600060208284031215610ecc57600080fd5b81356001600160e01b031981168114610ee457600080fd5b9392505050565b600060208284031215610efd57600080fd5b5035919050565b80356001600160a01b0381168114610f1b57600080fd5b919050565b60008060408385031215610f3357600080fd5b82359150610f4360208401610f04565b90509250929050565b600080600060608486031215610f6157600080fd5b610f6a84610f04565b925060208401359150610f7f60408501610f04565b90509250925092565b600060208284031215610f9a57600080fd5b610ee482610f04565b60008083601f840112610fb557600080fd5b50813567ffffffffffffffff811115610fcd57600080fd5b602083019150836020828501011115610fe557600080fd5b9250929050565b60008060006040848603121561100157600080fd5b61100a84610f04565b9250602084013567ffffffffffffffff81111561102657600080fd5b61103286828701610fa3565b9497909650939450505050565b6000806000806060858703121561105557600080fd5b61105e85610f04565b9350602085013567ffffffffffffffff81111561107a57600080fd5b61108687828801610fa3565b9598909750949560400135949350505050565b6000806000806000608086880312156110b157600080fd5b6110ba86610f04565b94506110c860208701610f04565b9350604086013567ffffffffffffffff8111156110e457600080fd5b6110f088828901610fa3565b96999598509660600135949350505050565b6020808252602a908201527f494353323042616e6b3a20746f6b656e436f6e7472616374206973206e6f7420604082015269184818dbdb9d1c9858dd60b21b606082015260800190565b60006020828403121561115e57600080fd5b81518015158114610ee457600080fd5b8183823760009101908152919050565b60208082526028908201527f494353323042616e6b3a206d7573742068617665206d696e74657220726f6c65604082015267081d1bc81b5a5b9d60c21b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b818103818111156102a6576102a66111c6565b808201808211156102a6576102a66111c6565b604081528260408201528284606083013760006060848301015260006060601f19601f8601168301019050826020830152949350505050565b60005b8381101561125657818101518382015260200161123e565b50506000910152565b6000825161127181846020870161123b565b9190910192915050565b604081526000835180604084015261129a81606085016020880161123b565b602083019390935250601f91909101601f191601606001919050565b80820281158282048414176102a6576102a66111c6565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600081611308576113086111c6565b50600019019056fe97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929a26469706673582212203e9e7b0200939cd4e47a97084619691e11194bd223546d80379dc98caddf3e8e64736f6c63430008190033",
}

// ICS20BankABI is the input ABI used to generate the binding from.
// Deprecated: Use ICS20BankMetaData.ABI instead.
var ICS20BankABI = ICS20BankMetaData.ABI

// ICS20BankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ICS20BankMetaData.Bin instead.
var ICS20BankBin = ICS20BankMetaData.Bin

// DeployICS20Bank deploys a new Ethereum contract, binding an instance of ICS20Bank to it.
func DeployICS20Bank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ICS20Bank, error) {
	parsed, err := ICS20BankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ICS20BankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ICS20Bank{ICS20BankCaller: ICS20BankCaller{contract: contract}, ICS20BankTransactor: ICS20BankTransactor{contract: contract}, ICS20BankFilterer: ICS20BankFilterer{contract: contract}}, nil
}

// ICS20Bank is an auto generated Go binding around an Ethereum contract.
type ICS20Bank struct {
	ICS20BankCaller     // Read-only binding to the contract
	ICS20BankTransactor // Write-only binding to the contract
	ICS20BankFilterer   // Log filterer for contract events
}

// ICS20BankCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICS20BankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS20BankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICS20BankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS20BankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICS20BankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS20BankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICS20BankSession struct {
	Contract     *ICS20Bank        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICS20BankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICS20BankCallerSession struct {
	Contract *ICS20BankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ICS20BankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICS20BankTransactorSession struct {
	Contract     *ICS20BankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ICS20BankRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICS20BankRaw struct {
	Contract *ICS20Bank // Generic contract binding to access the raw methods on
}

// ICS20BankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICS20BankCallerRaw struct {
	Contract *ICS20BankCaller // Generic read-only contract binding to access the raw methods on
}

// ICS20BankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICS20BankTransactorRaw struct {
	Contract *ICS20BankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICS20Bank creates a new instance of ICS20Bank, bound to a specific deployed contract.
func NewICS20Bank(address common.Address, backend bind.ContractBackend) (*ICS20Bank, error) {
	contract, err := bindICS20Bank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICS20Bank{ICS20BankCaller: ICS20BankCaller{contract: contract}, ICS20BankTransactor: ICS20BankTransactor{contract: contract}, ICS20BankFilterer: ICS20BankFilterer{contract: contract}}, nil
}

// NewICS20BankCaller creates a new read-only instance of ICS20Bank, bound to a specific deployed contract.
func NewICS20BankCaller(address common.Address, caller bind.ContractCaller) (*ICS20BankCaller, error) {
	contract, err := bindICS20Bank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICS20BankCaller{contract: contract}, nil
}

// NewICS20BankTransactor creates a new write-only instance of ICS20Bank, bound to a specific deployed contract.
func NewICS20BankTransactor(address common.Address, transactor bind.ContractTransactor) (*ICS20BankTransactor, error) {
	contract, err := bindICS20Bank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransactor{contract: contract}, nil
}

// NewICS20BankFilterer creates a new log filterer instance of ICS20Bank, bound to a specific deployed contract.
func NewICS20BankFilterer(address common.Address, filterer bind.ContractFilterer) (*ICS20BankFilterer, error) {
	contract, err := bindICS20Bank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICS20BankFilterer{contract: contract}, nil
}

// bindICS20Bank binds a generic wrapper to an already deployed contract.
func bindICS20Bank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICS20BankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICS20Bank *ICS20BankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICS20Bank.Contract.ICS20BankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICS20Bank *ICS20BankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICS20Bank.Contract.ICS20BankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICS20Bank *ICS20BankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICS20Bank.Contract.ICS20BankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICS20Bank *ICS20BankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICS20Bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICS20Bank *ICS20BankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICS20Bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICS20Bank *ICS20BankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICS20Bank.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ICS20Bank *ICS20BankCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ICS20Bank.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ICS20Bank *ICS20BankSession) ADMINROLE() ([32]byte, error) {
	return _ICS20Bank.Contract.ADMINROLE(&_ICS20Bank.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ICS20Bank *ICS20BankCallerSession) ADMINROLE() ([32]byte, error) {
	return _ICS20Bank.Contract.ADMINROLE(&_ICS20Bank.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ICS20Bank *ICS20BankCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ICS20Bank.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ICS20Bank *ICS20BankSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ICS20Bank.Contract.DEFAULTADMINROLE(&_ICS20Bank.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ICS20Bank *ICS20BankCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ICS20Bank.Contract.DEFAULTADMINROLE(&_ICS20Bank.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_ICS20Bank *ICS20BankCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ICS20Bank.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_ICS20Bank *ICS20BankSession) OPERATORROLE() ([32]byte, error) {
	return _ICS20Bank.Contract.OPERATORROLE(&_ICS20Bank.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_ICS20Bank *ICS20BankCallerSession) OPERATORROLE() ([32]byte, error) {
	return _ICS20Bank.Contract.OPERATORROLE(&_ICS20Bank.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0xb9b092c8.
//
// Solidity: function balanceOf(address account, string id) view returns(uint256)
func (_ICS20Bank *ICS20BankCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id string) (*big.Int, error) {
	var out []interface{}
	err := _ICS20Bank.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xb9b092c8.
//
// Solidity: function balanceOf(address account, string id) view returns(uint256)
func (_ICS20Bank *ICS20BankSession) BalanceOf(account common.Address, id string) (*big.Int, error) {
	return _ICS20Bank.Contract.BalanceOf(&_ICS20Bank.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0xb9b092c8.
//
// Solidity: function balanceOf(address account, string id) view returns(uint256)
func (_ICS20Bank *ICS20BankCallerSession) BalanceOf(account common.Address, id string) (*big.Int, error) {
	return _ICS20Bank.Contract.BalanceOf(&_ICS20Bank.CallOpts, account, id)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ICS20Bank *ICS20BankCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ICS20Bank.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ICS20Bank *ICS20BankSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ICS20Bank.Contract.GetRoleAdmin(&_ICS20Bank.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ICS20Bank *ICS20BankCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ICS20Bank.Contract.GetRoleAdmin(&_ICS20Bank.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ICS20Bank *ICS20BankCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ICS20Bank.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ICS20Bank *ICS20BankSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ICS20Bank.Contract.HasRole(&_ICS20Bank.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ICS20Bank *ICS20BankCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ICS20Bank.Contract.HasRole(&_ICS20Bank.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ICS20Bank *ICS20BankCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ICS20Bank.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ICS20Bank *ICS20BankSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ICS20Bank.Contract.SupportsInterface(&_ICS20Bank.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ICS20Bank *ICS20BankCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ICS20Bank.Contract.SupportsInterface(&_ICS20Bank.CallOpts, interfaceId)
}

// Burn is a paid mutator transaction binding the contract method 0xc45b71de.
//
// Solidity: function burn(address account, string id, uint256 amount) returns()
func (_ICS20Bank *ICS20BankTransactor) Burn(opts *bind.TransactOpts, account common.Address, id string, amount *big.Int) (*types.Transaction, error) {
	return _ICS20Bank.contract.Transact(opts, "burn", account, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xc45b71de.
//
// Solidity: function burn(address account, string id, uint256 amount) returns()
func (_ICS20Bank *ICS20BankSession) Burn(account common.Address, id string, amount *big.Int) (*types.Transaction, error) {
	return _ICS20Bank.Contract.Burn(&_ICS20Bank.TransactOpts, account, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xc45b71de.
//
// Solidity: function burn(address account, string id, uint256 amount) returns()
func (_ICS20Bank *ICS20BankTransactorSession) Burn(account common.Address, id string, amount *big.Int) (*types.Transaction, error) {
	return _ICS20Bank.Contract.Burn(&_ICS20Bank.TransactOpts, account, id, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address tokenContract, uint256 amount, address receiver) returns()
func (_ICS20Bank *ICS20BankTransactor) Deposit(opts *bind.TransactOpts, tokenContract common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ICS20Bank.contract.Transact(opts, "deposit", tokenContract, amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address tokenContract, uint256 amount, address receiver) returns()
func (_ICS20Bank *ICS20BankSession) Deposit(tokenContract common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ICS20Bank.Contract.Deposit(&_ICS20Bank.TransactOpts, tokenContract, amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address tokenContract, uint256 amount, address receiver) returns()
func (_ICS20Bank *ICS20BankTransactorSession) Deposit(tokenContract common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ICS20Bank.Contract.Deposit(&_ICS20Bank.TransactOpts, tokenContract, amount, receiver)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ICS20Bank *ICS20BankTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ICS20Bank.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ICS20Bank *ICS20BankSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ICS20Bank.Contract.GrantRole(&_ICS20Bank.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ICS20Bank *ICS20BankTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ICS20Bank.Contract.GrantRole(&_ICS20Bank.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address account, string id, uint256 amount) returns()
func (_ICS20Bank *ICS20BankTransactor) Mint(opts *bind.TransactOpts, account common.Address, id string, amount *big.Int) (*types.Transaction, error) {
	return _ICS20Bank.contract.Transact(opts, "mint", account, id, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address account, string id, uint256 amount) returns()
func (_ICS20Bank *ICS20BankSession) Mint(account common.Address, id string, amount *big.Int) (*types.Transaction, error) {
	return _ICS20Bank.Contract.Mint(&_ICS20Bank.TransactOpts, account, id, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address account, string id, uint256 amount) returns()
func (_ICS20Bank *ICS20BankTransactorSession) Mint(account common.Address, id string, amount *big.Int) (*types.Transaction, error) {
	return _ICS20Bank.Contract.Mint(&_ICS20Bank.TransactOpts, account, id, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ICS20Bank *ICS20BankTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ICS20Bank.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ICS20Bank *ICS20BankSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ICS20Bank.Contract.RenounceRole(&_ICS20Bank.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ICS20Bank *ICS20BankTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ICS20Bank.Contract.RenounceRole(&_ICS20Bank.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ICS20Bank *ICS20BankTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ICS20Bank.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ICS20Bank *ICS20BankSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ICS20Bank.Contract.RevokeRole(&_ICS20Bank.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ICS20Bank *ICS20BankTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ICS20Bank.Contract.RevokeRole(&_ICS20Bank.TransactOpts, role, account)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_ICS20Bank *ICS20BankTransactor) SetOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ICS20Bank.contract.Transact(opts, "setOperator", operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_ICS20Bank *ICS20BankSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _ICS20Bank.Contract.SetOperator(&_ICS20Bank.TransactOpts, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_ICS20Bank *ICS20BankTransactorSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _ICS20Bank.Contract.SetOperator(&_ICS20Bank.TransactOpts, operator)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xf24dc1da.
//
// Solidity: function transferFrom(address from, address to, string id, uint256 amount) returns()
func (_ICS20Bank *ICS20BankTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id string, amount *big.Int) (*types.Transaction, error) {
	return _ICS20Bank.contract.Transact(opts, "transferFrom", from, to, id, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xf24dc1da.
//
// Solidity: function transferFrom(address from, address to, string id, uint256 amount) returns()
func (_ICS20Bank *ICS20BankSession) TransferFrom(from common.Address, to common.Address, id string, amount *big.Int) (*types.Transaction, error) {
	return _ICS20Bank.Contract.TransferFrom(&_ICS20Bank.TransactOpts, from, to, id, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xf24dc1da.
//
// Solidity: function transferFrom(address from, address to, string id, uint256 amount) returns()
func (_ICS20Bank *ICS20BankTransactorSession) TransferFrom(from common.Address, to common.Address, id string, amount *big.Int) (*types.Transaction, error) {
	return _ICS20Bank.Contract.TransferFrom(&_ICS20Bank.TransactOpts, from, to, id, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address tokenContract, uint256 amount, address receiver) returns()
func (_ICS20Bank *ICS20BankTransactor) Withdraw(opts *bind.TransactOpts, tokenContract common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ICS20Bank.contract.Transact(opts, "withdraw", tokenContract, amount, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address tokenContract, uint256 amount, address receiver) returns()
func (_ICS20Bank *ICS20BankSession) Withdraw(tokenContract common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ICS20Bank.Contract.Withdraw(&_ICS20Bank.TransactOpts, tokenContract, amount, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address tokenContract, uint256 amount, address receiver) returns()
func (_ICS20Bank *ICS20BankTransactorSession) Withdraw(tokenContract common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ICS20Bank.Contract.Withdraw(&_ICS20Bank.TransactOpts, tokenContract, amount, receiver)
}

// ICS20BankRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ICS20Bank contract.
type ICS20BankRoleAdminChangedIterator struct {
	Event *ICS20BankRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ICS20BankRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20BankRoleAdminChanged)
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
		it.Event = new(ICS20BankRoleAdminChanged)
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
func (it *ICS20BankRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20BankRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20BankRoleAdminChanged represents a RoleAdminChanged event raised by the ICS20Bank contract.
type ICS20BankRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ICS20Bank *ICS20BankFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ICS20BankRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ICS20Bank.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ICS20BankRoleAdminChangedIterator{contract: _ICS20Bank.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ICS20Bank *ICS20BankFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ICS20BankRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ICS20Bank.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20BankRoleAdminChanged)
				if err := _ICS20Bank.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ICS20Bank *ICS20BankFilterer) ParseRoleAdminChanged(log types.Log) (*ICS20BankRoleAdminChanged, error) {
	event := new(ICS20BankRoleAdminChanged)
	if err := _ICS20Bank.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICS20BankRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ICS20Bank contract.
type ICS20BankRoleGrantedIterator struct {
	Event *ICS20BankRoleGranted // Event containing the contract specifics and raw log

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
func (it *ICS20BankRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20BankRoleGranted)
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
		it.Event = new(ICS20BankRoleGranted)
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
func (it *ICS20BankRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20BankRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20BankRoleGranted represents a RoleGranted event raised by the ICS20Bank contract.
type ICS20BankRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ICS20Bank *ICS20BankFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ICS20BankRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ICS20Bank.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ICS20BankRoleGrantedIterator{contract: _ICS20Bank.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ICS20Bank *ICS20BankFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ICS20BankRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ICS20Bank.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20BankRoleGranted)
				if err := _ICS20Bank.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ICS20Bank *ICS20BankFilterer) ParseRoleGranted(log types.Log) (*ICS20BankRoleGranted, error) {
	event := new(ICS20BankRoleGranted)
	if err := _ICS20Bank.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICS20BankRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ICS20Bank contract.
type ICS20BankRoleRevokedIterator struct {
	Event *ICS20BankRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ICS20BankRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20BankRoleRevoked)
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
		it.Event = new(ICS20BankRoleRevoked)
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
func (it *ICS20BankRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20BankRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20BankRoleRevoked represents a RoleRevoked event raised by the ICS20Bank contract.
type ICS20BankRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ICS20Bank *ICS20BankFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ICS20BankRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ICS20Bank.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ICS20BankRoleRevokedIterator{contract: _ICS20Bank.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ICS20Bank *ICS20BankFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ICS20BankRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ICS20Bank.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20BankRoleRevoked)
				if err := _ICS20Bank.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ICS20Bank *ICS20BankFilterer) ParseRoleRevoked(log types.Log) (*ICS20BankRoleRevoked, error) {
	event := new(ICS20BankRoleRevoked)
	if err := _ICS20Bank.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICS20BankTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ICS20Bank contract.
type ICS20BankTransferIterator struct {
	Event *ICS20BankTransfer // Event containing the contract specifics and raw log

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
func (it *ICS20BankTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20BankTransfer)
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
		it.Event = new(ICS20BankTransfer)
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
func (it *ICS20BankTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20BankTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20BankTransfer represents a Transfer event raised by the ICS20Bank contract.
type ICS20BankTransfer struct {
	From  common.Address
	To    common.Address
	Path  string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x1d30d3db8e01fa0d5626c471596f822f597e720c26a2930ef20d3387313c3d78.
//
// Solidity: event Transfer(address indexed from, address indexed to, string path, uint256 value)
func (_ICS20Bank *ICS20BankFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ICS20BankTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICS20Bank.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferIterator{contract: _ICS20Bank.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x1d30d3db8e01fa0d5626c471596f822f597e720c26a2930ef20d3387313c3d78.
//
// Solidity: event Transfer(address indexed from, address indexed to, string path, uint256 value)
func (_ICS20Bank *ICS20BankFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ICS20BankTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICS20Bank.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20BankTransfer)
				if err := _ICS20Bank.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x1d30d3db8e01fa0d5626c471596f822f597e720c26a2930ef20d3387313c3d78.
//
// Solidity: event Transfer(address indexed from, address indexed to, string path, uint256 value)
func (_ICS20Bank *ICS20BankFilterer) ParseTransfer(log types.Log) (*ICS20BankTransfer, error) {
	event := new(ICS20BankTransfer)
	if err := _ICS20Bank.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
