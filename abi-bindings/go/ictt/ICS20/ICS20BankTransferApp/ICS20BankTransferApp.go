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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ibcAddr\",\"type\":\"address\"},{\"internalType\":\"contractIICS20Bank\",\"name\":\"bank\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"contractITokenRouter\",\"name\":\"tokenRouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"prefixedDenom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ibcAddress\",\"type\":\"address\"}],\"name\":\"PacketReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"PacketSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"TokenRoutingFailed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"someAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"}],\"name\":\"bindPort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibcAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"ack\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onAcknowledgementPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onRecvPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onTimeoutPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chan\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"chanAddr\",\"type\":\"address\"}],\"name\":\"setChannelEscrowAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"home\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isExternal\",\"type\":\"bool\"}],\"name\":\"setTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenRouter\",\"outputs\":[{\"internalType\":\"contractTokenRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051611fa0380380611fa083398101604081905261002f91610188565b80826001600160a01b03811661006057604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b61006981610120565b506001600160a01b0381166100d55760405162461bcd60e51b815260206004820152602c60248201527f4942434261736546756e6769626c65546f6b656e4170703a207a65726f20726f60448201526b75746572206164647265737360a01b6064820152608401610057565b6001600160a01b03908116608052600280549582166001600160a01b0319968716179055600380549482169486169490941790935560048054919093169316929092179055506101e7565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038116811461018557600080fd5b50565b6000806000806080858703121561019e57600080fd5b84516101a981610170565b60208601519094506101ba81610170565b60408601519093506101cb81610170565b60608601519092506101dc81610170565b939692955090935050565b608051611d97610209600039600081816101980152610c620152611d976000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806374a707421161007157806374a70742146101495780638da5cb5b1461015c578063977c6d331461016d5780639d19876514610180578063a562bdba14610193578063f2fde38b146101ba57600080fd5b80633df7d406146100b95780635849f2df146100e15780635aa1e0ff146100f657806363c1bec514610109578063696a9bf41461011c578063715018a614610141575b600080fd5b6100cc6100c73660046112b6565b6101cd565b60405190151581526020015b60405180910390f35b6100f46100ef36600461132e565b61023e565b005b6100f461010436600461137f565b61028a565b6100cc61011736600461145b565b61037f565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016100d8565b6100f461043c565b6100f461015736600461150f565b610450565b6000546001600160a01b0316610129565b6100cc61017b3660046112b6565b610495565b6100f461018e3660046115e6565b610745565b6101297f000000000000000000000000000000000000000000000000000000000000000081565b6100f46101c836600461162b565b6107a7565b6002546000906001600160a01b031633146102035760405162461bcd60e51b81526004016101fa90611648565b60405180910390fd5b60008360a0015180602001905181019061021d919061170d565b905061023281856020015186604001516107e5565b60019150505b92915050565b610246610904565b8060018360405161025791906117ef565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b03199092169190911790555050565b7f1ecf18a8aa27f47b640ba93dcdf4d464a88d8afc960a26a309d76ddd194e6eea8888888888866040516102c396959493929190611837565b60405180910390a1600260009054906101000a90046001600160a01b03166001600160a01b0316636052bf6f6000878787876103228f8f8f8b60405160200161030e91815260200190565b604051602081830303815290604052610931565b6040518763ffffffff1660e01b81526004016103439695949392919061189e565b600060405180830381600087803b15801561035d57600080fd5b505af1158015610371573d6000803e3d6000fd5b505050505050505050505050565b6002546000906001600160a01b031633146103ac5760405162461bcd60e51b81526004016101fa90611648565b6040805180820190915260118152707b22726573756c74223a2241513d3d227d60781b6020918201528351908401207fad671c182373c3087a2d5e877d68648c7b716d38afcdadb5298e477e2f04fb98146104325760008460a0015180602001905181019061041b919061170d565b905061043081866020015187604001516107e5565b505b5060019392505050565b610444610904565b61044e60006109c0565b565b610458610904565b60048054604051633a5383a160e11b81526001600160a01b03909116916374a7074291610343918c918c918b918d918c918c918c918c9101611901565b6000808360a001518060200190518101906104b0919061170d565b905060006104c18260000151610a10565b905060006104d786602001518760400151610a3d565b905060006104e483610acb565b90506104f08383610b3a565b610501576104fe8284610b7d565b90505b60008061050d83610c14565b915091508161056f5760405162461bcd60e51b815260206004820152602860248201527f546f6b656e20726f757465206e6f7420666f756e6420666f722070726566697860448201526765642064656e6f6d60c01b60648201526084016101fa565b8060a00151156105b1577f476f2c085b42da8c6a24daf2d3a6d5546fc2cafdd94b1ac39b635fda15cd63a6836040516105a8919061196e565b60405180910390a15b8060a00151156106035760405162461bcd60e51b815260206004820152601a60248201527f6e617469766520746f6b656e206e6f7420737570706f7274656400000000000060448201526064016101fa565b7fd3dcfe28d129b213ff54d08eede757be050cacdc923163eda54b76daa7174d328660000151848860200151896040015161064c60008c60600151610d7b90919063ffffffff16565b60808c0151336002546001600160a01b03166040516106729897969594939291906119af565b60405180910390a1600081602001519050806001600160a01b031663c1affb1265636f736d6f7360d01b89604001516040516020016106b191906117ef565b60408051601f1981840301815291905260608b01516106d1906000610d7b565b8b602001516106e38d60800151610de0565b6040518663ffffffff1660e01b8152600401610703959493929190611a37565b600060405180830381600087803b15801561071d57600080fd5b505af1158015610731573d6000803e3d6000fd5b5060019d9c50505050505050505050505050565b60405163c13b184f60e01b81526001600160a01b0383169063c13b184f90610771908490600401611a76565b600060405180830381600087803b15801561078b57600080fd5b505af115801561079f573d6000803e3d6000fd5b505050505050565b6107af610904565b6001600160a01b0381166107d957604051631e4fbdf760e01b8152600060048201526024016101fa565b6107e2816109c0565b50565b6108036107f28383610a3d565b84516107fd90610a10565b90610b3a565b61088e5761082f61081382610e64565b6040850151610823906000610d7b565b85516020870151610ef6565b6108895760405162461bcd60e51b815260206004820152602560248201527f4261736546756e6769626c65546f6b656e4170703a207472616e736665722066604482015264185a5b195960da1b60648201526084016101fa565b505050565b60408301516108ae906108a2906000610d7b565b84516020860151610f6a565b6108895760405162461bcd60e51b815260206004820152602160248201527f4261736546756e6769626c65546f6b656e4170703a206d696e74206661696c656044820152601960fa1b60648201526084016101fa565b6000546001600160a01b0316331461044e5760405163118cdaa760e01b81523360048201526024016101fa565b60606040518060a001604052808681526020018581526020016109513390565b604051602001610979919060609190911b6bffffffffffffffffffffffff1916815260140190565b6040516020818303038152906040528152602001848152602001838152506040516020016109a79190611a89565b6040516020818303038152906040529050949350505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820190915260008082526020820152610ac4610abf610a79604051806040016040528060018152602001602f60f81b815250610a10565b610ab9610abf610a8887610a10565b610ab9610abf610ab0604051806040016040528060018152602001602f60f81b815250610a10565b610ab98c610a10565b90610b7d565b610a10565b9392505050565b6060600082600001516001600160401b03811115610aeb57610aeb611055565b6040519080825280601f01601f191660200182016040528015610b15576020820181803683370190505b5090506000602082019050610b338185602001518660000151610fdb565b5092915050565b805182516000911115610b4f57506000610238565b8160200151836020015103610b6657506001610238565b508051602092830151929091015181902091201490565b80518251606091600091610b919190611b27565b6001600160401b03811115610ba857610ba8611055565b6040519080825280601f01601f191660200182016040528015610bd2576020820181803683370190505b5090506000602082019050610bf08186602001518760000151610fdb565b8451610c0c90610c009083611b27565b60208601518651610fdb565b509392505050565b6040805160e0810182526000808252602082018190529181018290526060808201526080810182905260a0810182905260c081018290526040516381f868f760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906381f868f790610c97908690600401611a76565b600060405180830381865afa925050508015610cd557506040513d6000823e601f3d908101601f19168201604052610cd29190810190611b3a565b60015b610d275750506040805160e08101825260008082526020808301829052828401829052835190810190935280835260608201929092526080810182905260a0810182905260c081018290529092909150565b6040805160e0810182526001600160a01b03988916815296881660208801529490961693850193909352606084019190915260ff166080830152151560a082015290151560c0820152600194909350915050565b6000610d88826014611b27565b83511015610dd05760405162461bcd60e51b8152602060048201526015602482015274746f416464726573735f6f75744f66426f756e647360581b60448201526064016101fa565b500160200151600160601b900490565b60008151600003610df357506000919050565b602082511115610e5c5760405162461bcd60e51b815260206004820152602e60248201527f6279746573546f427974657333323a20736f75726365206c656e67746820657860448201526d636565647320333220627974657360901b60648201526084016101fa565b506020015190565b600080600183604051610e7791906117ef565b908152604051908190036020019020546001600160a01b03169050806102385760405162461bcd60e51b815260206004820152602e60248201527f4261736546756e6769626c65546f6b656e4170703a20657363726f772061646460448201526d1c995cdcc81b9bdd08199bdd5b9960921b60648201526084016101fa565b600354604051637926e0ed60e11b81526000916001600160a01b03169063f24dc1da90610f2d908890889088908890600401611bed565b600060405180830381600087803b158015610f4757600080fd5b505af1158015610f5b573d6000803e3d6000fd5b50600198975050505050505050565b60035460405163ba7aef4360e01b81526000916001600160a01b03169063ba7aef4390610f9f90879087908790600401611c2a565b600060405180830381600087803b158015610fb957600080fd5b505af1158015610fcd573d6000803e3d6000fd5b506001979650505050505050565b602081106110135781518352610ff2602084611b27565b9250610fff602083611b27565b915061100c602082611c5e565b9050610fdb565b6000198115611042576001611029836020611c5e565b61103590610100611d55565b61103f9190611c5e565b90505b9151835183169219169190911790915250565b634e487b7160e01b600052604160045260246000fd5b60405161010081016001600160401b038111828210171561108e5761108e611055565b60405290565b60405160a081016001600160401b038111828210171561108e5761108e611055565b604051601f8201601f191681016001600160401b03811182821017156110de576110de611055565b604052919050565b60006001600160401b038211156110ff576110ff611055565b50601f01601f191660200190565b600082601f83011261111e57600080fd5b813561113161112c826110e6565b6110b6565b81815284602083860101111561114657600080fd5b816020850160208301376000918101602001919091529392505050565b60006040828403121561117557600080fd5b604051604081018181106001600160401b038211171561119757611197611055565b604052823581526020928301359281019290925250919050565b600061012082840312156111c457600080fd5b6111cc61106b565b90508135815260208201356001600160401b03808211156111ec57600080fd5b6111f88583860161110d565b6020840152604084013591508082111561121157600080fd5b61121d8583860161110d565b6040840152606084013591508082111561123657600080fd5b6112428583860161110d565b6060840152608084013591508082111561125b57600080fd5b6112678583860161110d565b608084015260a084013591508082111561128057600080fd5b5061128d8482850161110d565b60a0830152506112a08360c08401611163565b60c082015261010082013560e082015292915050565b600080604083850312156112c957600080fd5b82356001600160401b03808211156112e057600080fd5b6112ec868387016111b1565b9350602085013591508082111561130257600080fd5b5061130f8582860161110d565b9150509250929050565b6001600160a01b03811681146107e257600080fd5b6000806040838503121561134157600080fd5b82356001600160401b0381111561135757600080fd5b6113638582860161110d565b925050602083013561137481611319565b809150509250929050565b600080600080600080600080610120898b03121561139c57600080fd5b88356001600160401b03808211156113b357600080fd5b6113bf8c838d0161110d565b995060208b0135985060408b01359150808211156113dc57600080fd5b6113e88c838d0161110d565b975060608b01359150808211156113fe57600080fd5b61140a8c838d0161110d565b965060808b013591508082111561142057600080fd5b5061142d8b828c0161110d565b94505061143d8a60a08b01611163565b925060e0890135915061010089013590509295985092959890939650565b60008060006060848603121561147057600080fd5b83356001600160401b038082111561148757600080fd5b611493878388016111b1565b945060208601359150808211156114a957600080fd5b6114b58783880161110d565b935060408601359150808211156114cb57600080fd5b506114d88682870161110d565b9150509250925092565b60ff811681146107e257600080fd5b80151581146107e257600080fd5b803561150a816114f1565b919050565b600080600080600080600080610100898b03121561152c57600080fd5b88356001600160401b038082111561154357600080fd5b61154f8c838d0161110d565b995060208b0135915061156182611319565b90975060408a01359061157382611319565b90965060608a01359061158582611319565b90955060808a0135908082111561159b57600080fd5b506115a88b828c0161110d565b94505060a08901356115b9816114e2565b925060c08901356115c9816114f1565b91506115d760e08a016114ff565b90509295985092959890939650565b600080604083850312156115f957600080fd5b823561160481611319565b915060208301356001600160401b0381111561161f57600080fd5b61130f8582860161110d565b60006020828403121561163d57600080fd5b8135610ac481611319565b60208082526034908201527f4261736546756e6769626c65546f6b656e4170703a2063616c6c6572206973206040820152731b9bdd081d1a1948125090c818dbdb9d1c9858dd60621b606082015260800190565b60005b838110156116b757818101518382015260200161169f565b50506000910152565b600082601f8301126116d157600080fd5b81516116df61112c826110e6565b8181528460208386010111156116f457600080fd5b61170582602083016020870161169c565b949350505050565b60006020828403121561171f57600080fd5b81516001600160401b038082111561173657600080fd5b9083019060a0828603121561174a57600080fd5b611752611094565b82518281111561176157600080fd5b61176d878286016116c0565b8252506020830151602082015260408301518281111561178c57600080fd5b611798878286016116c0565b6040830152506060830151828111156117b057600080fd5b6117bc878286016116c0565b6060830152506080830151828111156117d457600080fd5b6117e0878286016116c0565b60808301525095945050505050565b6000825161180181846020870161169c565b9190910192915050565b6000815180845261182381602086016020860161169c565b601f01601f19169290920160200192915050565b60c08152600061184a60c083018961180b565b8760208401528281036040840152611862818861180b565b90508281036060840152611876818761180b565b9050828103608084015261188a818661180b565b9150508260a0830152979650505050505050565b86815260e0602082015260006118b760e083018861180b565b82810360408401526118c9818861180b565b905085516060840152602086015160808401528460a084015282810360c08401526118f4818561180b565b9998505050505050505050565b60006101008083526119158184018c61180b565b6001600160a01b038b811660208601528a811660408601528916606085015283810360808501529050611948818861180b565b60ff9690961660a0840152505091151560c0830152151560e09091015295945050505050565b604081526000611981604083018461180b565b8281036020840152600c81526b3730ba34bb32903a37b5b2b760a11b60208201526040810191505092915050565b60006101008083526119c38184018c61180b565b905082810360208401526119d7818b61180b565b905088604084015282810360608401526119f1818961180b565b6001600160a01b03888116608086015284820360a0860152909150611a16828861180b565b925080861660c085015280851660e085015250509998505050505050505050565b85815260a060208201526000611a5060a083018761180b565b6001600160a01b0395909516604083015250606081019290925260809091015292915050565b602081526000610ac4602083018461180b565b602081526000825160a06020840152611aa560c084018261180b565b9050602084015160408401526040840151601f1980858403016060860152611acd838361180b565b92506060860151915080858403016080860152611aea838361180b565b925060808601519150808584030160a086015250611b08828261180b565b95945050505050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561023857610238611b11565b600080600080600080600060e0888a031215611b5557600080fd5b8751611b6081611319565b6020890151909750611b7181611319565b6040890151909650611b8281611319565b60608901519095506001600160401b03811115611b9e57600080fd5b611baa8a828b016116c0565b9450506080880151611bbb816114e2565b60a0890151909350611bcc816114f1565b60c0890151909250611bdd816114f1565b8091505092959891949750929550565b6001600160a01b03858116825284166020820152608060408201819052600090611c199083018561180b565b905082606083015295945050505050565b6001600160a01b0384168152606060208201819052600090611c4e9083018561180b565b9050826040830152949350505050565b8181038181111561023857610238611b11565b600181815b80851115611cac578160001904821115611c9257611c92611b11565b80851615611c9f57918102915b93841c9390800290611c76565b509250929050565b600082611cc357506001610238565b81611cd057506000610238565b8160018114611ce65760028114611cf057611d0c565b6001915050610238565b60ff841115611d0157611d01611b11565b50506001821b610238565b5060208310610133831016604e8410600b8410161715611d2f575081810a610238565b611d398383611c71565b8060001904821115611d4d57611d4d611b11565b029392505050565b6000610ac48383611cb456fea2646970667358221220e9ad701306772a900d60fcf9d4a80c49a3dc866b13c3f2f5121d4231750f8ebd64736f6c63430008190033",
}

// ICS20BankTransferAppABI is the input ABI used to generate the binding from.
// Deprecated: Use ICS20BankTransferAppMetaData.ABI instead.
var ICS20BankTransferAppABI = ICS20BankTransferAppMetaData.ABI

// ICS20BankTransferAppBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ICS20BankTransferAppMetaData.Bin instead.
var ICS20BankTransferAppBin = ICS20BankTransferAppMetaData.Bin

// DeployICS20BankTransferApp deploys a new Ethereum contract, binding an instance of ICS20BankTransferApp to it.
func DeployICS20BankTransferApp(auth *bind.TransactOpts, backend bind.ContractBackend, ibcAddr common.Address, bank common.Address, initialOwner common.Address, tokenRouter common.Address) (common.Address, *types.Transaction, *ICS20BankTransferApp, error) {
	parsed, err := ICS20BankTransferAppMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ICS20BankTransferAppBin), backend, ibcAddr, bank, initialOwner, tokenRouter)
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
// Solidity: function onAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ack, bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet Packet, ack []byte, arg2 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.contract.Transact(opts, "onAcknowledgementPacket", packet, ack, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x63c1bec5.
//
// Solidity: function onAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ack, bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) OnAcknowledgementPacket(packet Packet, ack []byte, arg2 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.OnAcknowledgementPacket(&_ICS20BankTransferApp.TransactOpts, packet, ack, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x63c1bec5.
//
// Solidity: function onAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ack, bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorSession) OnAcknowledgementPacket(packet Packet, ack []byte, arg2 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.OnAcknowledgementPacket(&_ICS20BankTransferApp.TransactOpts, packet, ack, arg2)
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

// SetChannelEscrowAddresses is a paid mutator transaction binding the contract method 0x5849f2df.
//
// Solidity: function setChannelEscrowAddresses(string chan, address chanAddr) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactor) SetChannelEscrowAddresses(opts *bind.TransactOpts, arg0 string, chanAddr common.Address) (*types.Transaction, error) {
	return _ICS20BankTransferApp.contract.Transact(opts, "setChannelEscrowAddresses", arg0, chanAddr)
}

// SetChannelEscrowAddresses is a paid mutator transaction binding the contract method 0x5849f2df.
//
// Solidity: function setChannelEscrowAddresses(string chan, address chanAddr) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) SetChannelEscrowAddresses(arg0 string, chanAddr common.Address) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.SetChannelEscrowAddresses(&_ICS20BankTransferApp.TransactOpts, arg0, chanAddr)
}

// SetChannelEscrowAddresses is a paid mutator transaction binding the contract method 0x5849f2df.
//
// Solidity: function setChannelEscrowAddresses(string chan, address chanAddr) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorSession) SetChannelEscrowAddresses(arg0 string, chanAddr common.Address) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.SetChannelEscrowAddresses(&_ICS20BankTransferApp.TransactOpts, arg0, chanAddr)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x74a70742.
//
// Solidity: function setTokenConfig(string denom, address token, address remote, address home, string channel, uint8 decimals, bool isNative, bool isExternal) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactor) SetTokenConfig(opts *bind.TransactOpts, denom string, token common.Address, remote common.Address, home common.Address, channel string, decimals uint8, isNative bool, isExternal bool) (*types.Transaction, error) {
	return _ICS20BankTransferApp.contract.Transact(opts, "setTokenConfig", denom, token, remote, home, channel, decimals, isNative, isExternal)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x74a70742.
//
// Solidity: function setTokenConfig(string denom, address token, address remote, address home, string channel, uint8 decimals, bool isNative, bool isExternal) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) SetTokenConfig(denom string, token common.Address, remote common.Address, home common.Address, channel string, decimals uint8, isNative bool, isExternal bool) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.SetTokenConfig(&_ICS20BankTransferApp.TransactOpts, denom, token, remote, home, channel, decimals, isNative, isExternal)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x74a70742.
//
// Solidity: function setTokenConfig(string denom, address token, address remote, address home, string channel, uint8 decimals, bool isNative, bool isExternal) returns()
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorSession) SetTokenConfig(denom string, token common.Address, remote common.Address, home common.Address, channel string, decimals uint8, isNative bool, isExternal bool) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.SetTokenConfig(&_ICS20BankTransferApp.TransactOpts, denom, token, remote, home, channel, decimals, isNative, isExternal)
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

// ICS20BankTransferAppPacketReceivedIterator is returned from FilterPacketReceived and is used to iterate over the raw logs and unpacked data for PacketReceived events raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppPacketReceivedIterator struct {
	Event *ICS20BankTransferAppPacketReceived // Event containing the contract specifics and raw log

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
func (it *ICS20BankTransferAppPacketReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20BankTransferAppPacketReceived)
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
		it.Event = new(ICS20BankTransferAppPacketReceived)
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
func (it *ICS20BankTransferAppPacketReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20BankTransferAppPacketReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20BankTransferAppPacketReceived represents a PacketReceived event raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppPacketReceived struct {
	Denom         string
	PrefixedDenom string
	Amount        *big.Int
	Sender        []byte
	Receiver      common.Address
	Memo          []byte
	MsgSender     common.Address
	IbcAddress    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPacketReceived is a free log retrieval operation binding the contract event 0xd3dcfe28d129b213ff54d08eede757be050cacdc923163eda54b76daa7174d32.
//
// Solidity: event PacketReceived(string denom, string prefixedDenom, uint256 amount, bytes sender, address receiver, bytes memo, address msgSender, address ibcAddress)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) FilterPacketReceived(opts *bind.FilterOpts) (*ICS20BankTransferAppPacketReceivedIterator, error) {

	logs, sub, err := _ICS20BankTransferApp.contract.FilterLogs(opts, "PacketReceived")
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferAppPacketReceivedIterator{contract: _ICS20BankTransferApp.contract, event: "PacketReceived", logs: logs, sub: sub}, nil
}

// WatchPacketReceived is a free log subscription operation binding the contract event 0xd3dcfe28d129b213ff54d08eede757be050cacdc923163eda54b76daa7174d32.
//
// Solidity: event PacketReceived(string denom, string prefixedDenom, uint256 amount, bytes sender, address receiver, bytes memo, address msgSender, address ibcAddress)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) WatchPacketReceived(opts *bind.WatchOpts, sink chan<- *ICS20BankTransferAppPacketReceived) (event.Subscription, error) {

	logs, sub, err := _ICS20BankTransferApp.contract.WatchLogs(opts, "PacketReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20BankTransferAppPacketReceived)
				if err := _ICS20BankTransferApp.contract.UnpackLog(event, "PacketReceived", log); err != nil {
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

// ParsePacketReceived is a log parse operation binding the contract event 0xd3dcfe28d129b213ff54d08eede757be050cacdc923163eda54b76daa7174d32.
//
// Solidity: event PacketReceived(string denom, string prefixedDenom, uint256 amount, bytes sender, address receiver, bytes memo, address msgSender, address ibcAddress)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) ParsePacketReceived(log types.Log) (*ICS20BankTransferAppPacketReceived, error) {
	event := new(ICS20BankTransferAppPacketReceived)
	if err := _ICS20BankTransferApp.contract.UnpackLog(event, "PacketReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICS20BankTransferAppPacketSentIterator is returned from FilterPacketSent and is used to iterate over the raw logs and unpacked data for PacketSent events raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppPacketSentIterator struct {
	Event *ICS20BankTransferAppPacketSent // Event containing the contract specifics and raw log

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
func (it *ICS20BankTransferAppPacketSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20BankTransferAppPacketSent)
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
		it.Event = new(ICS20BankTransferAppPacketSent)
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
func (it *ICS20BankTransferAppPacketSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20BankTransferAppPacketSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20BankTransferAppPacketSent represents a PacketSent event raised by the ICS20BankTransferApp contract.
type ICS20BankTransferAppPacketSent struct {
	Denom         string
	Amount        *big.Int
	Receiver      []byte
	SourcePort    string
	SourceChannel string
	MessageID     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPacketSent is a free log retrieval operation binding the contract event 0x1ecf18a8aa27f47b640ba93dcdf4d464a88d8afc960a26a309d76ddd194e6eea.
//
// Solidity: event PacketSent(string denom, uint256 amount, bytes receiver, string sourcePort, string sourceChannel, bytes32 messageID)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) FilterPacketSent(opts *bind.FilterOpts) (*ICS20BankTransferAppPacketSentIterator, error) {

	logs, sub, err := _ICS20BankTransferApp.contract.FilterLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return &ICS20BankTransferAppPacketSentIterator{contract: _ICS20BankTransferApp.contract, event: "PacketSent", logs: logs, sub: sub}, nil
}

// WatchPacketSent is a free log subscription operation binding the contract event 0x1ecf18a8aa27f47b640ba93dcdf4d464a88d8afc960a26a309d76ddd194e6eea.
//
// Solidity: event PacketSent(string denom, uint256 amount, bytes receiver, string sourcePort, string sourceChannel, bytes32 messageID)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) WatchPacketSent(opts *bind.WatchOpts, sink chan<- *ICS20BankTransferAppPacketSent) (event.Subscription, error) {

	logs, sub, err := _ICS20BankTransferApp.contract.WatchLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20BankTransferAppPacketSent)
				if err := _ICS20BankTransferApp.contract.UnpackLog(event, "PacketSent", log); err != nil {
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

// ParsePacketSent is a log parse operation binding the contract event 0x1ecf18a8aa27f47b640ba93dcdf4d464a88d8afc960a26a309d76ddd194e6eea.
//
// Solidity: event PacketSent(string denom, uint256 amount, bytes receiver, string sourcePort, string sourceChannel, bytes32 messageID)
func (_ICS20BankTransferApp *ICS20BankTransferAppFilterer) ParsePacketSent(log types.Log) (*ICS20BankTransferAppPacketSent, error) {
	event := new(ICS20BankTransferAppPacketSent)
	if err := _ICS20BankTransferApp.contract.UnpackLog(event, "PacketSent", log); err != nil {
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
