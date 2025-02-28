// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ics20banktransferapp

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

// Height is an auto generated low-level Go binding around an user-defined struct.
type Height struct {
	RevisionNumber *big.Int
	RevisionHeight *big.Int
}

// Packet is an auto generated low-level Go binding around an user-defined struct.
type Packet struct {
	Sequence           *big.Int
	SourcePort         string
	SourceChannel      string
	DestinationPort    string
	DestinationChannel string
	Data               []byte
	TimeoutHeight      Height
	TimeoutTimestamp   *big.Int
}

// ICS20BankTransferAppMetaData contains all meta data concerning the ICS20BankTransferApp contract.
var ICS20BankTransferAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ibcAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"contractITokenRouter\",\"name\":\"tokenRouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"}],\"name\":\"PacketAcknowledged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"PacketTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"binder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"}],\"name\":\"PortBound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"TokenRoutingFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"TokenTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"someAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"}],\"name\":\"bindPort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibcAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onAcknowledgementPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onRecvPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onTimeoutPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenRouter\",\"outputs\":[{\"internalType\":\"contractTokenRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051611ad4380380611ad483398101604081905261002e916101d4565b80826001600160a01b03811661005e57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61006781610155565b506001600160a01b0381166100d35760405162461bcd60e51b815260206004820152602c60248201527f4942434261736546756e6769626c65546f6b656e4170703a207a65726f20726f60448201526b75746572206164647265737360a01b6064820152608401610055565b6001600160a01b0390811660805283166101225760405162461bcd60e51b815260206004820152601060248201526f7a65726f20494243206164647265737360801b6044820152606401610055565b600280546001600160a01b039485166001600160a01b03199182161790915560038054929094169116179091555061021e565b600180546001600160a01b031916905561016e81610171565b50565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038116811461016e575f80fd5b5f805f606084860312156101e6575f80fd5b83516101f1816101c0565b6020850151909350610202816101c0565b6040850151909250610213816101c0565b809150509250925092565b60805161189761023d5f395f81816101750152610aad01526118975ff3fe608060405234801561000f575f80fd5b50600436106100b1575f3560e01c80638da5cb5b1161006e5780638da5cb5b1461013a578063977c6d331461014a5780639d1987651461015d578063a562bdba14610170578063e30c397814610197578063f2fde38b146101a8575f80fd5b80633df7d406146100b55780635aa1e0ff146100dd57806363c1bec5146100f2578063696a9bf414610105578063715018a61461012a57806379ba509714610132575b5f80fd5b6100c86100c3366004610fc1565b6101bb565b60405190151581526020015b60405180910390f35b6100f06100eb366004611020565b61024b565b005b6100c86101003660046110f3565b610350565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016100d4565b6100f06103d7565b6100f06103ea565b5f546001600160a01b0316610112565b6100c8610158366004610fc1565b61042e565b6100f061016b366004611188565b610685565b6101127f000000000000000000000000000000000000000000000000000000000000000081565b6001546001600160a01b0316610112565b6100f06101b63660046111ca565b610723565b6002545f906001600160a01b031633146101f05760405162461bcd60e51b81526004016101e7906111e5565b60405180910390fd5b825f01517f57ff39a74ccd9ca587e73592bed69151fdbb1b89c99269a71aa879fa67c5f53b84602001518560400151866060015187608001516040516102399493929190611286565b60405180910390a25060015b92915050565b60025f9054906101000a90046001600160a01b03166001600160a01b0316636052bf6f5f878787876102a08f8f8f8b60405160200161028c91815260200190565b604051602081830303815290604052610793565b6040518763ffffffff1660e01b81526004016102c196959493929190611306565b5f604051808303815f87803b1580156102d8575f80fd5b505af11580156102ea573d5f803e3d5ffd5b5050505080886040516102fd9190611368565b60405180910390207fe693ddbb0a238dd2b07003241439cebbaf4e1d246a519bd4e1950b23431bc37389898989898960405161033e96959493929190611383565b60405180910390a35050505050505050565b6002545f906001600160a01b0316331461037c5760405162461bcd60e51b81526004016101e7906111e5565b835f01517f585093c5d85b859e4f94900605b213a12e85859091fee98f5afc7b1d8395b09b85602001518660400151876060015188608001516040516103c594939291906113ec565b60405180910390a25060019392505050565b6103df610822565b6103e85f61084e565b565b60015433906001600160a01b031681146104225760405163118cdaa760e01b81526001600160a01b03821660048201526024016101e7565b61042b8161084e565b50565b6002545f906001600160a01b0316331461045a5760405162461bcd60e51b81526004016101e7906111e5565b5f8360a00151806020019051810190610473919061148d565b90505f610482825f0151610867565b90505f61049786602001518760400151610893565b90505f6104a383610920565b90506104af838361098b565b6104c0576104bd82846109cc565b90505b5f806104cb83610a60565b915091508161052d5760405162461bcd60e51b815260206004820152602860248201527f546f6b656e20726f757465206e6f7420666f756e6420666f722070726566697860448201526765642064656e6f6d60c01b60648201526084016101e7565b8060a001511561056f577f476f2c085b42da8c6a24daf2d3a6d5546fc2cafdd94b1ac39b635fda15cd63a6836040516105669190611567565b60405180910390a15b8060a00151156105c15760405162461bcd60e51b815260206004820152601a60248201527f6e617469766520746f6b656e206e6f7420737570706f7274656400000000000060448201526064016101e7565b5f81602001519050806001600160a01b031663c1affb1265636f736d6f7360d01b89604001516040516020016105f79190611368565b60408051601f1981840301815291905260608b0151610616905f610bc3565b8b602001516106288d60800151610c27565b6040518663ffffffff1660e01b81526004016106489594939291906115a7565b5f604051808303815f87803b15801561065f575f80fd5b505af1158015610671573d5f803e3d5ffd5b5060019d9c50505050505050505050505050565b60405163c13b184f60e01b81526001600160a01b0383169063c13b184f906106b19084906004016115e5565b5f604051808303815f87803b1580156106c8575f80fd5b505af11580156106da573d5f803e3d5ffd5b50505050336001600160a01b03167fa11e5bab29d27b01501f85603d08d3e3d62334a45dab10389ecf672daf9c5f568260405161071791906115e5565b60405180910390a25050565b61072b610822565b600180546001600160a01b0383166001600160a01b0319909116811790915561075b5f546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b60606040518060a001604052808681526020018581526020016107b33390565b6040516020016107db919060609190911b6bffffffffffffffffffffffff1916815260140190565b60405160208183030381529060405281526020018481526020018381525060405160200161080991906115f7565b6040516020818303038152906040529050949350505050565b5f546001600160a01b031633146103e85760405163118cdaa760e01b81523360048201526024016101e7565b600180546001600160a01b031916905561042b81610ca8565b6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082019091525f80825260208201526109196109146108ce604051806040016040528060018152602001602f60f81b815250610867565b61090e6109146108dd87610867565b61090e610914610905604051806040016040528060018152602001602f60f81b815250610867565b61090e8c610867565b906109cc565b610867565b9392505050565b60605f825f01516001600160401b0381111561093e5761093e610d70565b6040519080825280601f01601f191660200182016040528015610968576020820181803683370190505b5090505f602082019050610984818560200151865f0151610cf7565b5092915050565b805182515f91111561099e57505f610245565b81602001518360200151036109b557506001610245565b508051602092830151929091015181902091201490565b805182516060915f916109df9190611692565b6001600160401b038111156109f6576109f6610d70565b6040519080825280601f01601f191660200182016040528015610a20576020820181803683370190505b5090505f602082019050610a3c818660200151875f0151610cf7565b8451610a5890610a4c9083611692565b60208601518651610cf7565b509392505050565b6040805160e0810182525f808252602082018190529181018290526060808201526080810182905260a0810182905260c081018290526040516381f868f760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906381f868f790610ae29086906004016115e5565b5f60405180830381865afa925050508015610b1e57506040513d5f823e601f3d908101601f19168201604052610b1b91908101906116b9565b60015b610b6f5750506040805160e0810182525f8082526020808301829052828401829052835190810190935280835260608201929092526080810182905260a0810182905260c081018290529092909150565b6040805160e0810182526001600160a01b03988916815296881660208801529490961693850193909352606084019190915260ff166080830152151560a082015290151560c0820152600194909350915050565b5f610bcf826014611692565b83511015610c175760405162461bcd60e51b8152602060048201526015602482015274746f416464726573735f6f75744f66426f756e647360581b60448201526064016101e7565b500160200151600160601b900490565b5f81515f03610c3757505f919050565b602082511115610ca05760405162461bcd60e51b815260206004820152602e60248201527f6279746573546f427974657333323a20736f75726365206c656e67746820657860448201526d636565647320333220627974657360901b60648201526084016101e7565b506020015190565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60208110610d2f5781518352610d0e602084611692565b9250610d1b602083611692565b9150610d28602082611763565b9050610cf7565b5f198115610d5d576001610d44836020611763565b610d5090610100611856565b610d5a9190611763565b90505b9151835183169219169190911790915250565b634e487b7160e01b5f52604160045260245ffd5b60405161010081016001600160401b0381118282101715610da757610da7610d70565b60405290565b60405160a081016001600160401b0381118282101715610da757610da7610d70565b604051601f8201601f191681016001600160401b0381118282101715610df757610df7610d70565b604052919050565b5f6001600160401b03821115610e1757610e17610d70565b50601f01601f191660200190565b5f82601f830112610e34575f80fd5b8135610e47610e4282610dff565b610dcf565b818152846020838601011115610e5b575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60408284031215610e87575f80fd5b604051604081018181106001600160401b0382111715610ea957610ea9610d70565b604052823581526020928301359281019290925250919050565b5f6101208284031215610ed4575f80fd5b610edc610d84565b90508135815260208201356001600160401b0380821115610efb575f80fd5b610f0785838601610e25565b60208401526040840135915080821115610f1f575f80fd5b610f2b85838601610e25565b60408401526060840135915080821115610f43575f80fd5b610f4f85838601610e25565b60608401526080840135915080821115610f67575f80fd5b610f7385838601610e25565b608084015260a0840135915080821115610f8b575f80fd5b50610f9884828501610e25565b60a083015250610fab8360c08401610e77565b60c082015261010082013560e082015292915050565b5f8060408385031215610fd2575f80fd5b82356001600160401b0380821115610fe8575f80fd5b610ff486838701610ec3565b93506020850135915080821115611009575f80fd5b5061101685828601610e25565b9150509250929050565b5f805f805f805f80610120898b031215611038575f80fd5b88356001600160401b038082111561104e575f80fd5b61105a8c838d01610e25565b995060208b0135985060408b0135915080821115611076575f80fd5b6110828c838d01610e25565b975060608b0135915080821115611097575f80fd5b6110a38c838d01610e25565b965060808b01359150808211156110b8575f80fd5b506110c58b828c01610e25565b9450506110d58a60a08b01610e77565b925060e0890135915061010089013590509295985092959890939650565b5f805f60608486031215611105575f80fd5b83356001600160401b038082111561111b575f80fd5b61112787838801610ec3565b9450602086013591508082111561113c575f80fd5b61114887838801610e25565b9350604086013591508082111561115d575f80fd5b5061116a86828701610e25565b9150509250925092565b6001600160a01b038116811461042b575f80fd5b5f8060408385031215611199575f80fd5b82356111a481611174565b915060208301356001600160401b038111156111be575f80fd5b61101685828601610e25565b5f602082840312156111da575f80fd5b813561091981611174565b60208082526034908201527f4261736546756e6769626c65546f6b656e4170703a2063616c6c6572206973206040820152731b9bdd081d1a1948125090c818dbdb9d1c9858dd60621b606082015260800190565b5f5b8381101561125357818101518382015260200161123b565b50505f910152565b5f8151808452611272816020860160208601611239565b601f01601f19169290920160200192915050565b60a081525f61129860a083018761125b565b82810360208401526112aa818761125b565b905082810360408401526112be818661125b565b905082810360608401526112d2818561125b565b8381036080909401939093525050600e81526d1c1858dad95d081d1a5b595bdd5d60921b6020820152604001949350505050565b86815260e060208201525f61131e60e083018861125b565b8281036040840152611330818861125b565b865160608501526020870151608085015290508460a084015282810360c084015261135b818561125b565b9998505050505050505050565b5f8251611379818460208701611239565b9190910192915050565b86815260e060208201525f61139b60e083018861125b565b82810360408401526113ad818861125b565b905082810360608401526113c1818761125b565b85516080850152602086015160a085015291506113db9050565b8260c0830152979650505050505050565b608081525f6113fe608083018761125b565b8281036020840152611410818761125b565b90508281036040840152611424818661125b565b90508281036060840152611438818561125b565b979650505050505050565b5f82601f830112611452575f80fd5b8151611460610e4282610dff565b818152846020838601011115611474575f80fd5b611485826020830160208701611239565b949350505050565b5f6020828403121561149d575f80fd5b81516001600160401b03808211156114b3575f80fd5b9083019060a082860312156114c6575f80fd5b6114ce610dad565b8251828111156114dc575f80fd5b6114e887828601611443565b82525060208301516020820152604083015182811115611506575f80fd5b61151287828601611443565b604083015250606083015182811115611529575f80fd5b61153587828601611443565b60608301525060808301518281111561154c575f80fd5b61155887828601611443565b60808301525095945050505050565b604081525f611579604083018461125b565b8281036020840152600c81526b3730ba34bb32903a37b5b2b760a11b60208201526040810191505092915050565b85815260a060208201525f6115bf60a083018761125b565b6001600160a01b0395909516604083015250606081019290925260809091015292915050565b602081525f610919602083018461125b565b602081525f825160a0602084015261161260c084018261125b565b9050602084015160408401526040840151601f198085840301606086015261163a838361125b565b92506060860151915080858403016080860152611657838361125b565b925060808601519150808584030160a086015250611675828261125b565b95945050505050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156102455761024561167e565b805180151581146116b4575f80fd5b919050565b5f805f805f805f60e0888a0312156116cf575f80fd5b87516116da81611174565b60208901519097506116eb81611174565b60408901519096506116fc81611174565b60608901519095506001600160401b03811115611717575f80fd5b6117238a828b01611443565b945050608088015160ff81168114611739575f80fd5b925061174760a089016116a5565b915061175560c089016116a5565b905092959891949750929550565b818103818111156102455761024561167e565b600181815b808511156117b057815f19048211156117965761179661167e565b808516156117a357918102915b93841c939080029061177b565b509250929050565b5f826117c657506001610245565b816117d257505f610245565b81600181146117e857600281146117f25761180e565b6001915050610245565b60ff8411156118035761180361167e565b50506001821b610245565b5060208310610133831016604e8410600b8410161715611831575081810a610245565b61183b8383611776565b805f190482111561184e5761184e61167e565b029392505050565b5f61091983836117b856fea26469706673582212202bb784d781b7f39639a8d33735527bbf23d878d66b98e7e2091f1b1c1d38151564736f6c63430008190033",
}

// ICS20BankTransferAppABI is the input ABI used to generate the binding from.
// Deprecated: Use ICS20BankTransferAppMetaData.ABI instead.
var ICS20BankTransferAppABI = ICS20BankTransferAppMetaData.ABI

// ICS20BankTransferAppBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ICS20BankTransferAppMetaData.Bin instead.
var ICS20BankTransferAppBin = ICS20BankTransferAppMetaData.Bin

// DeployICS20BankTransferApp deploys a new Ethereum contract, binding an instance of ICS20BankTransferApp to it.
func DeployICS20BankTransferApp(auth *bind.TransactOpts, backend bind.ContractBackend, ibcAddr common.Address, initialOwner common.Address, tokenRouter common.Address) (common.Address, *types.Transaction, *ICS20BankTransferApp, error) {
	parsed, err := ICS20BankTransferAppMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ICS20BankTransferAppBin), backend, ibcAddr, initialOwner, tokenRouter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ICS20BankTransferApp{ICS20BankTransferAppCaller: ICS20BankTransferAppCaller{contract: contract}, ICS20BankTransferAppTransactor: ICS20BankTransferAppTransactor{contract: contract}, ICS20BankTransferAppFilterer: ICS20BankTransferAppFilterer{contract: contract}}, nil
}

// ICS20BankTransferApp is an auto generated Go binding around an Ethereum contract.
type ICS20BankTransferApp struct {
	ICS20BankTransferAppCaller     // Read-only binding to the contract
	ICS20BankTransferAppTransactor // Write-only binding to the contract
	ICS20BankTransferAppFilterer   // Log filterer for contract events
}

// ICS20BankTransferAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICS20BankTransferAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS20BankTransferAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICS20BankTransferAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS20BankTransferAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICS20BankTransferAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS20BankTransferAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICS20BankTransferAppSession struct {
	Contract     *ICS20BankTransferApp // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ICS20BankTransferAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICS20BankTransferAppCallerSession struct {
	Contract *ICS20BankTransferAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ICS20BankTransferAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICS20BankTransferAppTransactorSession struct {
	Contract     *ICS20BankTransferAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ICS20BankTransferAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICS20BankTransferAppRaw struct {
	Contract *ICS20BankTransferApp // Generic contract binding to access the raw methods on
}

// ICS20BankTransferAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICS20BankTransferAppCallerRaw struct {
	Contract *ICS20BankTransferAppCaller // Generic read-only contract binding to access the raw methods on
}

// ICS20BankTransferAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICS20BankTransferAppTransactorRaw struct {
	Contract *ICS20BankTransferAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICS20BankTransferApp creates a new instance of ICS20BankTransferApp, bound to a specific deployed contract.
func NewICS20BankTransferApp(address common.Address, backend bind.ContractBackend) (*ICS20BankTransferApp, error) {
	contract, err := bindICS20BankTransferApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferApp{ICS20BankTransferAppCaller: ICS20BankTransferAppCaller{contract: contract}, ICS20BankTransferAppTransactor: ICS20BankTransferAppTransactor{contract: contract}, ICS20BankTransferAppFilterer: ICS20BankTransferAppFilterer{contract: contract}}, nil
}

// NewICS20BankTransferAppCaller creates a new read-only instance of ICS20BankTransferApp, bound to a specific deployed contract.
func NewICS20BankTransferAppCaller(address common.Address, caller bind.ContractCaller) (*ICS20BankTransferAppCaller, error) {
	contract, err := bindICS20BankTransferApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferAppCaller{contract: contract}, nil
}

// NewICS20BankTransferAppTransactor creates a new write-only instance of ICS20BankTransferApp, bound to a specific deployed contract.
func NewICS20BankTransferAppTransactor(address common.Address, transactor bind.ContractTransactor) (*ICS20BankTransferAppTransactor, error) {
	contract, err := bindICS20BankTransferApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferAppTransactor{contract: contract}, nil
}

// NewICS20BankTransferAppFilterer creates a new log filterer instance of ICS20BankTransferApp, bound to a specific deployed contract.
func NewICS20BankTransferAppFilterer(address common.Address, filterer bind.ContractFilterer) (*ICS20BankTransferAppFilterer, error) {
	contract, err := bindICS20BankTransferApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferAppFilterer{contract: contract}, nil
}

// bindICS20BankTransferApp binds a generic wrapper to an already deployed contract.
func bindICS20BankTransferApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICS20BankTransferAppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICS20BankTransferApp *ICS20BankTransferAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICS20BankTransferApp.Contract.ICS20BankTransferAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICS20BankTransferApp *ICS20BankTransferAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.ICS20BankTransferAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICS20BankTransferApp *ICS20BankTransferAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.ICS20BankTransferAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICS20BankTransferApp *ICS20BankTransferAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICS20BankTransferApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.contract.Transact(opts, method, params...)
}

// IbcAddress is a free data retrieval call binding the contract method 0x696a9bf4.
//
// Solidity: function ibcAddress() view returns(address)
func (_ICS20BankTransferApp *ICS20BankTransferAppCaller) IbcAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICS20BankTransferApp.contract.Call(opts, &out, "ibcAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IbcAddress is a free data retrieval call binding the contract method 0x696a9bf4.
//
// Solidity: function ibcAddress() view returns(address)
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) IbcAddress() (common.Address, error) {
	return _ICS20BankTransferApp.Contract.IbcAddress(&_ICS20BankTransferApp.CallOpts)
}

// IbcAddress is a free data retrieval call binding the contract method 0x696a9bf4.
//
// Solidity: function ibcAddress() view returns(address)
func (_ICS20BankTransferApp *ICS20BankTransferAppCallerSession) IbcAddress() (common.Address, error) {
	return _ICS20BankTransferApp.Contract.IbcAddress(&_ICS20BankTransferApp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ICS20BankTransferApp *ICS20BankTransferAppCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICS20BankTransferApp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) Owner() (common.Address, error) {
	return _ICS20BankTransferApp.Contract.Owner(&_ICS20BankTransferApp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ICS20BankTransferApp *ICS20BankTransferAppCallerSession) Owner() (common.Address, error) {
	return _ICS20BankTransferApp.Contract.Owner(&_ICS20BankTransferApp.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ICS20BankTransferApp *ICS20BankTransferAppCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICS20BankTransferApp.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) PendingOwner() (common.Address, error) {
	return _ICS20BankTransferApp.Contract.PendingOwner(&_ICS20BankTransferApp.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ICS20BankTransferApp *ICS20BankTransferAppCallerSession) PendingOwner() (common.Address, error) {
	return _ICS20BankTransferApp.Contract.PendingOwner(&_ICS20BankTransferApp.CallOpts)
}

// TokenRouter is a free data retrieval call binding the contract method 0xa562bdba.
//
// Solidity: function tokenRouter() view returns(address)
func (_ICS20BankTransferApp *ICS20BankTransferAppCaller) TokenRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICS20BankTransferApp.contract.Call(opts, &out, "tokenRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenRouter is a free data retrieval call binding the contract method 0xa562bdba.
//
// Solidity: function tokenRouter() view returns(address)
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) TokenRouter() (common.Address, error) {
	return _ICS20BankTransferApp.Contract.TokenRouter(&_ICS20BankTransferApp.CallOpts)
}

// TokenRouter is a free data retrieval call binding the contract method 0xa562bdba.
//
// Solidity: function tokenRouter() view returns(address)
func (_ICS20BankTransferApp *ICS20BankTransferAppCallerSession) TokenRouter() (common.Address, error) {
	return _ICS20BankTransferApp.Contract.TokenRouter(&_ICS20BankTransferApp.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICS20BankTransferApp.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) AcceptOwnership() (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.AcceptOwnership(&_ICS20BankTransferApp.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.AcceptOwnership(&_ICS20BankTransferApp.TransactOpts)
}

// BindPort is a paid mutator transaction binding the contract method 0x9d198765.
//
// Solidity: function bindPort(address someAddr, string portId) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactor) BindPort(opts *bind.TransactOpts, someAddr common.Address, portId string) (*types.Transaction, error) {
	return _ICS20BankTransferApp.contract.Transact(opts, "bindPort", someAddr, portId)
}

// BindPort is a paid mutator transaction binding the contract method 0x9d198765.
//
// Solidity: function bindPort(address someAddr, string portId) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) BindPort(someAddr common.Address, portId string) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.BindPort(&_ICS20BankTransferApp.TransactOpts, someAddr, portId)
}

// BindPort is a paid mutator transaction binding the contract method 0x9d198765.
//
// Solidity: function bindPort(address someAddr, string portId) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorSession) BindPort(someAddr common.Address, portId string) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.BindPort(&_ICS20BankTransferApp.TransactOpts, someAddr, portId)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x63c1bec5.
//
// Solidity: function onAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes , bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet Packet, arg1 []byte, arg2 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.contract.Transact(opts, "onAcknowledgementPacket", packet, arg1, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x63c1bec5.
//
// Solidity: function onAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes , bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) OnAcknowledgementPacket(packet Packet, arg1 []byte, arg2 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.OnAcknowledgementPacket(&_ICS20BankTransferApp.TransactOpts, packet, arg1, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x63c1bec5.
//
// Solidity: function onAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes , bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorSession) OnAcknowledgementPacket(packet Packet, arg1 []byte, arg2 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.OnAcknowledgementPacket(&_ICS20BankTransferApp.TransactOpts, packet, arg1, arg2)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x977c6d33.
//
// Solidity: function onRecvPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactor) OnRecvPacket(opts *bind.TransactOpts, packet Packet, arg1 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.contract.Transact(opts, "onRecvPacket", packet, arg1)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x977c6d33.
//
// Solidity: function onRecvPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) OnRecvPacket(packet Packet, arg1 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.OnRecvPacket(&_ICS20BankTransferApp.TransactOpts, packet, arg1)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x977c6d33.
//
// Solidity: function onRecvPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorSession) OnRecvPacket(packet Packet, arg1 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.OnRecvPacket(&_ICS20BankTransferApp.TransactOpts, packet, arg1)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x3df7d406.
//
// Solidity: function onTimeoutPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactor) OnTimeoutPacket(opts *bind.TransactOpts, packet Packet, arg1 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.contract.Transact(opts, "onTimeoutPacket", packet, arg1)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x3df7d406.
//
// Solidity: function onTimeoutPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) OnTimeoutPacket(packet Packet, arg1 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.OnTimeoutPacket(&_ICS20BankTransferApp.TransactOpts, packet, arg1)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x3df7d406.
//
// Solidity: function onTimeoutPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorSession) OnTimeoutPacket(packet Packet, arg1 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.OnTimeoutPacket(&_ICS20BankTransferApp.TransactOpts, packet, arg1)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICS20BankTransferApp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) RenounceOwnership() (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.RenounceOwnership(&_ICS20BankTransferApp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.RenounceOwnership(&_ICS20BankTransferApp.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x5aa1e0ff.
//
// Solidity: function transfer(string denom, uint256 amount, bytes receiver, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes32 messageID) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactor) Transfer(opts *bind.TransactOpts, denom string, amount *big.Int, receiver []byte, sourcePort string, sourceChannel string, timeoutHeight Height, timeoutTimestamp *big.Int, messageID [32]byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.contract.Transact(opts, "transfer", denom, amount, receiver, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, messageID)
}

// Transfer is a paid mutator transaction binding the contract method 0x5aa1e0ff.
//
// Solidity: function transfer(string denom, uint256 amount, bytes receiver, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes32 messageID) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) Transfer(denom string, amount *big.Int, receiver []byte, sourcePort string, sourceChannel string, timeoutHeight Height, timeoutTimestamp *big.Int, messageID [32]byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.Transfer(&_ICS20BankTransferApp.TransactOpts, denom, amount, receiver, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, messageID)
}

// Transfer is a paid mutator transaction binding the contract method 0x5aa1e0ff.
//
// Solidity: function transfer(string denom, uint256 amount, bytes receiver, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes32 messageID) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorSession) Transfer(denom string, amount *big.Int, receiver []byte, sourcePort string, sourceChannel string, timeoutHeight Height, timeoutTimestamp *big.Int, messageID [32]byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.Transfer(&_ICS20BankTransferApp.TransactOpts, denom, amount, receiver, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, messageID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ICS20BankTransferApp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.TransferOwnership(&_ICS20BankTransferApp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.TransferOwnership(&_ICS20BankTransferApp.TransactOpts, newOwner)
}

// ICS20BankTransferAppOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppOwnershipTransferStartedIterator struct {
	Event *ICS20BankTransferAppOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *ICS20BankTransferAppOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20BankTransferAppOwnershipTransferStarted)
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
		it.Event = new(ICS20BankTransferAppOwnershipTransferStarted)
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
func (it *ICS20BankTransferAppOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20BankTransferAppOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20BankTransferAppOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ICS20BankTransferAppOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ICS20BankTransferApp.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferAppOwnershipTransferStartedIterator{contract: _ICS20BankTransferApp.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *ICS20BankTransferAppOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ICS20BankTransferApp.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20BankTransferAppOwnershipTransferStarted)
				if err := _ICS20BankTransferApp.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) ParseOwnershipTransferStarted(log types.Log) (*ICS20BankTransferAppOwnershipTransferStarted, error) {
	event := new(ICS20BankTransferAppOwnershipTransferStarted)
	if err := _ICS20BankTransferApp.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICS20BankTransferAppOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppOwnershipTransferredIterator struct {
	Event *ICS20BankTransferAppOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ICS20BankTransferAppOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20BankTransferAppOwnershipTransferred)
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
		it.Event = new(ICS20BankTransferAppOwnershipTransferred)
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
func (it *ICS20BankTransferAppOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20BankTransferAppOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20BankTransferAppOwnershipTransferred represents a OwnershipTransferred event raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ICS20BankTransferAppOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ICS20BankTransferApp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferAppOwnershipTransferredIterator{contract: _ICS20BankTransferApp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ICS20BankTransferAppOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ICS20BankTransferApp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20BankTransferAppOwnershipTransferred)
				if err := _ICS20BankTransferApp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) ParseOwnershipTransferred(log types.Log) (*ICS20BankTransferAppOwnershipTransferred, error) {
	event := new(ICS20BankTransferAppOwnershipTransferred)
	if err := _ICS20BankTransferApp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICS20BankTransferAppPacketAcknowledgedIterator is returned from FilterPacketAcknowledged and is used to iterate over the raw logs and unpacked data for PacketAcknowledged events raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppPacketAcknowledgedIterator struct {
	Event *ICS20BankTransferAppPacketAcknowledged // Event containing the contract specifics and raw log

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
func (it *ICS20BankTransferAppPacketAcknowledgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20BankTransferAppPacketAcknowledged)
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
		it.Event = new(ICS20BankTransferAppPacketAcknowledged)
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
func (it *ICS20BankTransferAppPacketAcknowledgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20BankTransferAppPacketAcknowledgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20BankTransferAppPacketAcknowledged represents a PacketAcknowledged event raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppPacketAcknowledged struct {
	Sequence           *big.Int
	SourcePort         string
	SourceChannel      string
	DestinationPort    string
	DestinationChannel string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPacketAcknowledged is a free log retrieval operation binding the contract event 0x585093c5d85b859e4f94900605b213a12e85859091fee98f5afc7b1d8395b09b.
//
// Solidity: event PacketAcknowledged(uint256 indexed sequence, string sourcePort, string sourceChannel, string destinationPort, string destinationChannel)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) FilterPacketAcknowledged(opts *bind.FilterOpts, sequence []*big.Int) (*ICS20BankTransferAppPacketAcknowledgedIterator, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ICS20BankTransferApp.contract.FilterLogs(opts, "PacketAcknowledged", sequenceRule)
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferAppPacketAcknowledgedIterator{contract: _ICS20BankTransferApp.contract, event: "PacketAcknowledged", logs: logs, sub: sub}, nil
}

// WatchPacketAcknowledged is a free log subscription operation binding the contract event 0x585093c5d85b859e4f94900605b213a12e85859091fee98f5afc7b1d8395b09b.
//
// Solidity: event PacketAcknowledged(uint256 indexed sequence, string sourcePort, string sourceChannel, string destinationPort, string destinationChannel)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) WatchPacketAcknowledged(opts *bind.WatchOpts, sink chan<- *ICS20BankTransferAppPacketAcknowledged, sequence []*big.Int) (event.Subscription, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ICS20BankTransferApp.contract.WatchLogs(opts, "PacketAcknowledged", sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20BankTransferAppPacketAcknowledged)
				if err := _ICS20BankTransferApp.contract.UnpackLog(event, "PacketAcknowledged", log); err != nil {
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

// ParsePacketAcknowledged is a log parse operation binding the contract event 0x585093c5d85b859e4f94900605b213a12e85859091fee98f5afc7b1d8395b09b.
//
// Solidity: event PacketAcknowledged(uint256 indexed sequence, string sourcePort, string sourceChannel, string destinationPort, string destinationChannel)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) ParsePacketAcknowledged(log types.Log) (*ICS20BankTransferAppPacketAcknowledged, error) {
	event := new(ICS20BankTransferAppPacketAcknowledged)
	if err := _ICS20BankTransferApp.contract.UnpackLog(event, "PacketAcknowledged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICS20BankTransferAppPacketTimeoutIterator is returned from FilterPacketTimeout and is used to iterate over the raw logs and unpacked data for PacketTimeout events raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppPacketTimeoutIterator struct {
	Event *ICS20BankTransferAppPacketTimeout // Event containing the contract specifics and raw log

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
func (it *ICS20BankTransferAppPacketTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20BankTransferAppPacketTimeout)
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
		it.Event = new(ICS20BankTransferAppPacketTimeout)
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
func (it *ICS20BankTransferAppPacketTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20BankTransferAppPacketTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20BankTransferAppPacketTimeout represents a PacketTimeout event raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppPacketTimeout struct {
	Sequence           *big.Int
	SourcePort         string
	SourceChannel      string
	DestinationPort    string
	DestinationChannel string
	Reason             string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPacketTimeout is a free log retrieval operation binding the contract event 0x57ff39a74ccd9ca587e73592bed69151fdbb1b89c99269a71aa879fa67c5f53b.
//
// Solidity: event PacketTimeout(uint256 indexed sequence, string sourcePort, string sourceChannel, string destinationPort, string destinationChannel, string reason)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) FilterPacketTimeout(opts *bind.FilterOpts, sequence []*big.Int) (*ICS20BankTransferAppPacketTimeoutIterator, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ICS20BankTransferApp.contract.FilterLogs(opts, "PacketTimeout", sequenceRule)
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferAppPacketTimeoutIterator{contract: _ICS20BankTransferApp.contract, event: "PacketTimeout", logs: logs, sub: sub}, nil
}

// WatchPacketTimeout is a free log subscription operation binding the contract event 0x57ff39a74ccd9ca587e73592bed69151fdbb1b89c99269a71aa879fa67c5f53b.
//
// Solidity: event PacketTimeout(uint256 indexed sequence, string sourcePort, string sourceChannel, string destinationPort, string destinationChannel, string reason)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) WatchPacketTimeout(opts *bind.WatchOpts, sink chan<- *ICS20BankTransferAppPacketTimeout, sequence []*big.Int) (event.Subscription, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ICS20BankTransferApp.contract.WatchLogs(opts, "PacketTimeout", sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20BankTransferAppPacketTimeout)
				if err := _ICS20BankTransferApp.contract.UnpackLog(event, "PacketTimeout", log); err != nil {
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

// ParsePacketTimeout is a log parse operation binding the contract event 0x57ff39a74ccd9ca587e73592bed69151fdbb1b89c99269a71aa879fa67c5f53b.
//
// Solidity: event PacketTimeout(uint256 indexed sequence, string sourcePort, string sourceChannel, string destinationPort, string destinationChannel, string reason)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) ParsePacketTimeout(log types.Log) (*ICS20BankTransferAppPacketTimeout, error) {
	event := new(ICS20BankTransferAppPacketTimeout)
	if err := _ICS20BankTransferApp.contract.UnpackLog(event, "PacketTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICS20BankTransferAppPortBoundIterator is returned from FilterPortBound and is used to iterate over the raw logs and unpacked data for PortBound events raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppPortBoundIterator struct {
	Event *ICS20BankTransferAppPortBound // Event containing the contract specifics and raw log

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
func (it *ICS20BankTransferAppPortBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20BankTransferAppPortBound)
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
		it.Event = new(ICS20BankTransferAppPortBound)
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
func (it *ICS20BankTransferAppPortBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20BankTransferAppPortBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20BankTransferAppPortBound represents a PortBound event raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppPortBound struct {
	Binder common.Address
	PortId string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPortBound is a free log retrieval operation binding the contract event 0xa11e5bab29d27b01501f85603d08d3e3d62334a45dab10389ecf672daf9c5f56.
//
// Solidity: event PortBound(address indexed binder, string portId)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) FilterPortBound(opts *bind.FilterOpts, binder []common.Address) (*ICS20BankTransferAppPortBoundIterator, error) {

	var binderRule []interface{}
	for _, binderItem := range binder {
		binderRule = append(binderRule, binderItem)
	}

	logs, sub, err := _ICS20BankTransferApp.contract.FilterLogs(opts, "PortBound", binderRule)
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferAppPortBoundIterator{contract: _ICS20BankTransferApp.contract, event: "PortBound", logs: logs, sub: sub}, nil
}

// WatchPortBound is a free log subscription operation binding the contract event 0xa11e5bab29d27b01501f85603d08d3e3d62334a45dab10389ecf672daf9c5f56.
//
// Solidity: event PortBound(address indexed binder, string portId)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) WatchPortBound(opts *bind.WatchOpts, sink chan<- *ICS20BankTransferAppPortBound, binder []common.Address) (event.Subscription, error) {

	var binderRule []interface{}
	for _, binderItem := range binder {
		binderRule = append(binderRule, binderItem)
	}

	logs, sub, err := _ICS20BankTransferApp.contract.WatchLogs(opts, "PortBound", binderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20BankTransferAppPortBound)
				if err := _ICS20BankTransferApp.contract.UnpackLog(event, "PortBound", log); err != nil {
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

// ParsePortBound is a log parse operation binding the contract event 0xa11e5bab29d27b01501f85603d08d3e3d62334a45dab10389ecf672daf9c5f56.
//
// Solidity: event PortBound(address indexed binder, string portId)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) ParsePortBound(log types.Log) (*ICS20BankTransferAppPortBound, error) {
	event := new(ICS20BankTransferAppPortBound)
	if err := _ICS20BankTransferApp.contract.UnpackLog(event, "PortBound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICS20BankTransferAppTokenRoutingFailedIterator is returned from FilterTokenRoutingFailed and is used to iterate over the raw logs and unpacked data for TokenRoutingFailed events raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppTokenRoutingFailedIterator struct {
	Event *ICS20BankTransferAppTokenRoutingFailed // Event containing the contract specifics and raw log

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
func (it *ICS20BankTransferAppTokenRoutingFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20BankTransferAppTokenRoutingFailed)
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
		it.Event = new(ICS20BankTransferAppTokenRoutingFailed)
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
func (it *ICS20BankTransferAppTokenRoutingFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20BankTransferAppTokenRoutingFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20BankTransferAppTokenRoutingFailed represents a TokenRoutingFailed event raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppTokenRoutingFailed struct {
	Denom  string
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenRoutingFailed is a free log retrieval operation binding the contract event 0x476f2c085b42da8c6a24daf2d3a6d5546fc2cafdd94b1ac39b635fda15cd63a6.
//
// Solidity: event TokenRoutingFailed(string denom, string reason)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) FilterTokenRoutingFailed(opts *bind.FilterOpts) (*ICS20BankTransferAppTokenRoutingFailedIterator, error) {

	logs, sub, err := _ICS20BankTransferApp.contract.FilterLogs(opts, "TokenRoutingFailed")
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferAppTokenRoutingFailedIterator{contract: _ICS20BankTransferApp.contract, event: "TokenRoutingFailed", logs: logs, sub: sub}, nil
}

// WatchTokenRoutingFailed is a free log subscription operation binding the contract event 0x476f2c085b42da8c6a24daf2d3a6d5546fc2cafdd94b1ac39b635fda15cd63a6.
//
// Solidity: event TokenRoutingFailed(string denom, string reason)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) WatchTokenRoutingFailed(opts *bind.WatchOpts, sink chan<- *ICS20BankTransferAppTokenRoutingFailed) (event.Subscription, error) {

	logs, sub, err := _ICS20BankTransferApp.contract.WatchLogs(opts, "TokenRoutingFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20BankTransferAppTokenRoutingFailed)
				if err := _ICS20BankTransferApp.contract.UnpackLog(event, "TokenRoutingFailed", log); err != nil {
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

// ParseTokenRoutingFailed is a log parse operation binding the contract event 0x476f2c085b42da8c6a24daf2d3a6d5546fc2cafdd94b1ac39b635fda15cd63a6.
//
// Solidity: event TokenRoutingFailed(string denom, string reason)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) ParseTokenRoutingFailed(log types.Log) (*ICS20BankTransferAppTokenRoutingFailed, error) {
	event := new(ICS20BankTransferAppTokenRoutingFailed)
	if err := _ICS20BankTransferApp.contract.UnpackLog(event, "TokenRoutingFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICS20BankTransferAppTokenTransferredIterator is returned from FilterTokenTransferred and is used to iterate over the raw logs and unpacked data for TokenTransferred events raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppTokenTransferredIterator struct {
	Event *ICS20BankTransferAppTokenTransferred // Event containing the contract specifics and raw log

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
func (it *ICS20BankTransferAppTokenTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20BankTransferAppTokenTransferred)
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
		it.Event = new(ICS20BankTransferAppTokenTransferred)
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
func (it *ICS20BankTransferAppTokenTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20BankTransferAppTokenTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20BankTransferAppTokenTransferred represents a TokenTransferred event raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppTokenTransferred struct {
	Denom            common.Hash
	Amount           *big.Int
	Receiver         []byte
	SourcePort       string
	SourceChannel    string
	TimeoutHeight    Height
	TimeoutTimestamp *big.Int
	MessageID        [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenTransferred is a free log retrieval operation binding the contract event 0xe693ddbb0a238dd2b07003241439cebbaf4e1d246a519bd4e1950b23431bc373.
//
// Solidity: event TokenTransferred(string indexed denom, uint256 amount, bytes receiver, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes32 indexed messageID)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) FilterTokenTransferred(opts *bind.FilterOpts, denom []string, messageID [][32]byte) (*ICS20BankTransferAppTokenTransferredIterator, error) {

	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _ICS20BankTransferApp.contract.FilterLogs(opts, "TokenTransferred", denomRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferAppTokenTransferredIterator{contract: _ICS20BankTransferApp.contract, event: "TokenTransferred", logs: logs, sub: sub}, nil
}

// WatchTokenTransferred is a free log subscription operation binding the contract event 0xe693ddbb0a238dd2b07003241439cebbaf4e1d246a519bd4e1950b23431bc373.
//
// Solidity: event TokenTransferred(string indexed denom, uint256 amount, bytes receiver, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes32 indexed messageID)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) WatchTokenTransferred(opts *bind.WatchOpts, sink chan<- *ICS20BankTransferAppTokenTransferred, denom []string, messageID [][32]byte) (event.Subscription, error) {

	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _ICS20BankTransferApp.contract.WatchLogs(opts, "TokenTransferred", denomRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20BankTransferAppTokenTransferred)
				if err := _ICS20BankTransferApp.contract.UnpackLog(event, "TokenTransferred", log); err != nil {
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

// ParseTokenTransferred is a log parse operation binding the contract event 0xe693ddbb0a238dd2b07003241439cebbaf4e1d246a519bd4e1950b23431bc373.
//
// Solidity: event TokenTransferred(string indexed denom, uint256 amount, bytes receiver, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes32 indexed messageID)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) ParseTokenTransferred(log types.Log) (*ICS20BankTransferAppTokenTransferred, error) {
	event := new(ICS20BankTransferAppTokenTransferred)
	if err := _ICS20BankTransferApp.contract.UnpackLog(event, "TokenTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
