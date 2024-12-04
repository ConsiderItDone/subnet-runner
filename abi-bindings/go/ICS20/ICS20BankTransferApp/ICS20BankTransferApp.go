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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ibcAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bank\",\"type\":\"address\",\"internalType\":\"contractIICS20Bank\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenRouter\",\"type\":\"address\",\"internalType\":\"contractITokenRouter\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bindPort\",\"inputs\":[{\"name\":\"someAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ibcAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onAcknowledgementPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structPacket\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revisionHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ack\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structPacket\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revisionHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structPacket\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revisionHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChannelEscrowAddresses\",\"inputs\":[{\"name\":\"chan\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chanAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenConfig\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"home\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isNative\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isExternal\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractTokenRouter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceChannel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revisionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revisionHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"messageID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenRoutingFailed\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051611fdc380380611fdc83398101604081905261002e91610184565b80826001600160a01b03811661005e57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b6100678161011e565b506001600160a01b0381166100d35760405162461bcd60e51b815260206004820152602c60248201527f4942434261736546756e6769626c65546f6b656e4170703a207a65726f20726f60448201526b75746572206164647265737360a01b6064820152608401610055565b6001600160a01b03908116608052600280549582166001600160a01b0319968716179055600380549482169486169490941790935560048054919093169316929092179055506101e0565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0381168114610181575f80fd5b50565b5f805f8060808587031215610197575f80fd5b84516101a28161016d565b60208601519094506101b38161016d565b60408601519093506101c48161016d565b60608601519092506101d58161016d565b939692955090935050565b608051611ddd6101ff5f395f81816101930152610e150152611ddd5ff3fe608060405234801561000f575f80fd5b50600436106100b1575f3560e01c806374a707421161006e57806374a70742146101455780638da5cb5b14610158578063977c6d33146101685780639d1987651461017b578063a562bdba1461018e578063f2fde38b146101b5575f80fd5b80633df7d406146100b55780635849f2df146100dd5780635aa1e0ff146100f257806363c1bec514610105578063696a9bf414610118578063715018a61461013d575b5f80fd5b6100c86100c3366004611427565b6101c8565b60405190151581526020015b60405180910390f35b6100f06100eb36600461149a565b610237565b005b6100f06101003660046114e8565b610283565b6100c86101133660046115bb565b610387565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016100d4565b6100f0610442565b6100f0610153366004611667565b610455565b5f546001600160a01b0316610125565b6100c8610176366004611427565b61049a565b6100f0610189366004611737565b610712565b6101257f000000000000000000000000000000000000000000000000000000000000000081565b6100f06101c3366004611779565b61076f565b6002545f906001600160a01b031633146101fd5760405162461bcd60e51b81526004016101f490611794565b60405180910390fd5b5f8360a001518060200190518101906102169190611854565b905061022b81856020015186604001516107ac565b60019150505b92915050565b61023f6108c9565b80600183604051610250919061192e565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b03199092169190911790555050565b6102b16102a261029387876108f5565b61029c8b610982565b906109ae565b6102ab8a610982565b90610a32565b6102cf576102c9336102c286610a45565b8a8a610ad6565b506102da565b6102da338989610b44565b60025f9054906101000a90046001600160a01b03166001600160a01b0316636052bf6f5f8787878761032f8f8f8f8b60405160200161031b91815260200190565b604051602081830303815290604052610baa565b6040518763ffffffff1660e01b815260040161035096959493929190611974565b5f604051808303815f87803b158015610367575f80fd5b505af1158015610379573d5f803e3d5ffd5b505050505050505050505050565b6002545f906001600160a01b031633146103b35760405162461bcd60e51b81526004016101f490611794565b6040805180820190915260118152707b22726573756c74223a2241513d3d227d60781b6020918201528351908401207fad671c182373c3087a2d5e877d68648c7b716d38afcdadb5298e477e2f04fb9814610438575f8460a001518060200190518101906104219190611854565b905061043681866020015187604001516107ac565b505b5060019392505050565b61044a6108c9565b6104535f610c39565b565b61045d6108c9565b60048054604051633a5383a160e11b81526001600160a01b03909116916374a7074291610350918c918c918b918d918c918c918c918c91016119d6565b6002545f906001600160a01b031633146104c65760405162461bcd60e51b81526004016101f490611794565b5f8360a001518060200190518101906104df9190611854565b90505f6104ee825f0151610982565b90505f610503866020015187604001516108f5565b90505f61050f83610c88565b905061051b8383610cf3565b61052c576105298284610d34565b90505b5f8061053783610dc8565b91509150816105995760405162461bcd60e51b815260206004820152602860248201527f546f6b656e20726f757465206e6f7420666f756e6420666f722070726566697860448201526765642064656e6f6d60c01b60648201526084016101f4565b8060a001511561062c577f476f2c085b42da8c6a24daf2d3a6d5546fc2cafdd94b1ac39b635fda15cd63a6836040516105d29190611a42565b60405180910390a18060a001511561062c5760405162461bcd60e51b815260206004820152601a60248201527f6e617469766520746f6b656e206e6f7420737570706f7274656400000000000060448201526064016101f4565b602081015160408701516001600160a01b0382169063c1affb129065636f736d6f7360d01b9061065c905f610f2b565b604051602001610684919060609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815291905260608b01516106a3905f610f2b565b8b602001516106b58d60800151610f8f565b6040518663ffffffff1660e01b81526004016106d5959493929190611a82565b5f604051808303815f87803b1580156106ec575f80fd5b505af11580156106fe573d5f803e3d5ffd5b5060019d9c50505050505050505050505050565b60405163c13b184f60e01b81526001600160a01b0383169063c13b184f9061073e908490600401611ac0565b5f604051808303815f87803b158015610755575f80fd5b505af1158015610767573d5f803e3d5ffd5b505050505050565b6107776108c9565b6001600160a01b0381166107a057604051631e4fbdf760e01b81525f60048201526024016101f4565b6107a981610c39565b50565b6107ca6107b983836108f5565b84516107c490610982565b90610cf3565b610854576107f56107da82610a45565b60408501516107e9905f610f2b565b85516020870151610ad6565b61084f5760405162461bcd60e51b815260206004820152602560248201527f4261736546756e6769626c65546f6b656e4170703a207472616e736665722066604482015264185a5b195960da1b60648201526084016101f4565b505050565b604083015161087390610867905f610f2b565b84516020860151611010565b61084f5760405162461bcd60e51b815260206004820152602160248201527f4261736546756e6769626c65546f6b656e4170703a206d696e74206661696c656044820152601960fa1b60648201526084016101f4565b5f546001600160a01b031633146104535760405163118cdaa760e01b81523360048201526024016101f4565b604080518082019091525f808252602082015261097b610976610930604051806040016040528060018152602001602f60f81b815250610982565b61097061097661093f87610982565b610970610976610967604051806040016040528060018152602001602f60f81b815250610982565b6109708c610982565b90610d34565b610982565b9392505050565b6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082019091525f80825260208201528151835110156109d2575081610231565b60208083015190840151600191146109f95750815160208481015190840151829020919020145b8015610a2a57825184518590610a10908390611ae6565b9052508251602085018051610a26908390611af9565b9052505b509192915050565b5f610a3d838361107b565b159392505050565b5f80600183604051610a57919061192e565b908152604051908190036020019020546001600160a01b03169050806102315760405162461bcd60e51b815260206004820152602e60248201527f4261736546756e6769626c65546f6b656e4170703a20657363726f772061646460448201526d1c995cdcc81b9bdd08199bdd5b9960921b60648201526084016101f4565b600354604051637926e0ed60e11b81525f916001600160a01b03169063f24dc1da90610b0c908890889088908890600401611b0c565b5f604051808303815f87803b158015610b23575f80fd5b505af1158015610b35573d5f803e3d5ffd5b50600198975050505050505050565b60035460405163622db8ef60e11b81526001600160a01b039091169063c45b71de90610b7890869086908690600401611b48565b5f604051808303815f87803b158015610b8f575f80fd5b505af1158015610ba1573d5f803e3d5ffd5b50505050505050565b60606040518060a00160405280868152602001858152602001610bca3390565b604051602001610bf2919060609190911b6bffffffffffffffffffffffff1916815260140190565b604051602081830303815290604052815260200184815260200183815250604051602001610c209190611b7b565b6040516020818303038152906040529050949350505050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60605f825f01516001600160401b03811115610ca657610ca66111d6565b6040519080825280601f01601f191660200182016040528015610cd0576020820181803683370190505b5090505f602082019050610cec818560200151865f015161115d565b5092915050565b805182515f911115610d0657505f610231565b8160200151836020015103610d1d57506001610231565b508051602092830151929091015181902091201490565b805182516060915f91610d479190611af9565b6001600160401b03811115610d5e57610d5e6111d6565b6040519080825280601f01601f191660200182016040528015610d88576020820181803683370190505b5090505f602082019050610da4818660200151875f015161115d565b8451610dc090610db49083611af9565b6020860151865161115d565b509392505050565b6040805160e0810182525f808252602082018190529181018290526060808201526080810182905260a0810182905260c081018290526040516381f868f760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906381f868f790610e4a908690600401611ac0565b5f60405180830381865afa925050508015610e8657506040513d5f823e601f3d908101601f19168201604052610e839190810190611c02565b60015b610ed75750506040805160e0810182525f8082526020808301829052828401829052835190810190935280835260608201929092526080810182905260a0810182905260c081018290529092909150565b6040805160e0810182526001600160a01b03988916815296881660208801529490961693850193909352606084019190915260ff166080830152151560a082015290151560c0820152600194909350915050565b5f610f37826014611af9565b83511015610f7f5760405162461bcd60e51b8152602060048201526015602482015274746f416464726573735f6f75744f66426f756e647360581b60448201526064016101f4565b500160200151600160601b900490565b5f81515f03610f9f57505f919050565b6020825111156110085760405162461bcd60e51b815260206004820152602e60248201527f6279746573546f427974657333323a20736f75726365206c656e67746820657860448201526d636565647320333220627974657360901b60648201526084016101f4565b506020015190565b60035460405163ba7aef4360e01b81525f916001600160a01b03169063ba7aef439061104490879087908790600401611b48565b5f604051808303815f87803b15801561105b575f80fd5b505af115801561106d573d5f803e3d5ffd5b506001979650505050505050565b815181515f919081111561108d575081515b602080850151908401515f5b838110156111445782518251808214611114575f1960208710156110f3576001846110c5896020611ae6565b6110cf9190611af9565b6110da906008611caf565b6110e5906002611da6565b6110ef9190611ae6565b1990505b81811683821681810391146111115797506102319650505050505050565b50505b61111f602086611af9565b945061112c602085611af9565b9350505060208161113d9190611af9565b9050611099565b50845186516111539190611db1565b9695505050505050565b602081106111955781518352611174602084611af9565b9250611181602083611af9565b915061118e602082611ae6565b905061115d565b5f1981156111c35760016111aa836020611ae6565b6111b690610100611da6565b6111c09190611ae6565b90505b9151835183169219169190911790915250565b634e487b7160e01b5f52604160045260245ffd5b60405161010081016001600160401b038111828210171561120d5761120d6111d6565b60405290565b60405160a081016001600160401b038111828210171561120d5761120d6111d6565b604051601f8201601f191681016001600160401b038111828210171561125d5761125d6111d6565b604052919050565b5f6001600160401b0382111561127d5761127d6111d6565b50601f01601f191660200190565b5f82601f83011261129a575f80fd5b81356112ad6112a882611265565b611235565b8181528460208386010111156112c1575f80fd5b816020850160208301375f918101602001919091529392505050565b5f604082840312156112ed575f80fd5b604051604081018181106001600160401b038211171561130f5761130f6111d6565b604052823581526020928301359281019290925250919050565b5f610120828403121561133a575f80fd5b6113426111ea565b90508135815260208201356001600160401b0380821115611361575f80fd5b61136d8583860161128b565b60208401526040840135915080821115611385575f80fd5b6113918583860161128b565b604084015260608401359150808211156113a9575f80fd5b6113b58583860161128b565b606084015260808401359150808211156113cd575f80fd5b6113d98583860161128b565b608084015260a08401359150808211156113f1575f80fd5b506113fe8482850161128b565b60a0830152506114118360c084016112dd565b60c082015261010082013560e082015292915050565b5f8060408385031215611438575f80fd5b82356001600160401b038082111561144e575f80fd5b61145a86838701611329565b9350602085013591508082111561146f575f80fd5b5061147c8582860161128b565b9150509250929050565b6001600160a01b03811681146107a9575f80fd5b5f80604083850312156114ab575f80fd5b82356001600160401b038111156114c0575f80fd5b6114cc8582860161128b565b92505060208301356114dd81611486565b809150509250929050565b5f805f805f805f80610120898b031215611500575f80fd5b88356001600160401b0380821115611516575f80fd5b6115228c838d0161128b565b995060208b0135985060408b013591508082111561153e575f80fd5b61154a8c838d0161128b565b975060608b013591508082111561155f575f80fd5b61156b8c838d0161128b565b965060808b0135915080821115611580575f80fd5b5061158d8b828c0161128b565b94505061159d8a60a08b016112dd565b925060e0890135915061010089013590509295985092959890939650565b5f805f606084860312156115cd575f80fd5b83356001600160401b03808211156115e3575f80fd5b6115ef87838801611329565b94506020860135915080821115611604575f80fd5b6116108783880161128b565b93506040860135915080821115611625575f80fd5b506116328682870161128b565b9150509250925092565b60ff811681146107a9575f80fd5b80151581146107a9575f80fd5b80356116628161164a565b919050565b5f805f805f805f80610100898b03121561167f575f80fd5b88356001600160401b0380821115611695575f80fd5b6116a18c838d0161128b565b995060208b013591506116b382611486565b90975060408a0135906116c582611486565b90965060608a0135906116d782611486565b90955060808a013590808211156116ec575f80fd5b506116f98b828c0161128b565b94505060a089013561170a8161163c565b925060c089013561171a8161164a565b915061172860e08a01611657565b90509295985092959890939650565b5f8060408385031215611748575f80fd5b823561175381611486565b915060208301356001600160401b0381111561176d575f80fd5b61147c8582860161128b565b5f60208284031215611789575f80fd5b813561097b81611486565b60208082526034908201527f4261736546756e6769626c65546f6b656e4170703a2063616c6c6572206973206040820152731b9bdd081d1a1948125090c818dbdb9d1c9858dd60621b606082015260800190565b5f5b838110156118025781810151838201526020016117ea565b50505f910152565b5f82601f830112611819575f80fd5b81516118276112a882611265565b81815284602083860101111561183b575f80fd5b61184c8260208301602087016117e8565b949350505050565b5f60208284031215611864575f80fd5b81516001600160401b038082111561187a575f80fd5b9083019060a0828603121561188d575f80fd5b611895611213565b8251828111156118a3575f80fd5b6118af8782860161180a565b825250602083015160208201526040830151828111156118cd575f80fd5b6118d98782860161180a565b6040830152506060830151828111156118f0575f80fd5b6118fc8782860161180a565b606083015250608083015182811115611913575f80fd5b61191f8782860161180a565b60808301525095945050505050565b5f825161193f8184602087016117e8565b9190910192915050565b5f81518084526119608160208601602086016117e8565b601f01601f19169290920160200192915050565b86815260e060208201525f61198c60e0830188611949565b828103604084015261199e8188611949565b905085516060840152602086015160808401528460a084015282810360c08401526119c98185611949565b9998505050505050505050565b5f6101008083526119e98184018c611949565b6001600160a01b038b811660208601528a811660408601528916606085015283810360808501529050611a1c8188611949565b60ff9690961660a0840152505091151560c0830152151560e09091015295945050505050565b604081525f611a546040830184611949565b8281036020840152600c81526b3730ba34bb32903a37b5b2b760a11b60208201526040810191505092915050565b85815260a060208201525f611a9a60a0830187611949565b6001600160a01b0395909516604083015250606081019290925260809091015292915050565b602081525f61097b6020830184611949565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561023157610231611ad2565b8082018082111561023157610231611ad2565b6001600160a01b038581168252841660208201526080604082018190525f90611b3790830185611949565b905082606083015295945050505050565b6001600160a01b03841681526060602082018190525f90611b6b90830185611949565b9050826040830152949350505050565b602081525f825160a06020840152611b9660c0840182611949565b9050602084015160408401526040840151601f1980858403016060860152611bbe8383611949565b92506060860151915080858403016080860152611bdb8383611949565b925060808601519150808584030160a086015250611bf98282611949565b95945050505050565b5f805f805f805f60e0888a031215611c18575f80fd5b8751611c2381611486565b6020890151909750611c3481611486565b6040890151909650611c4581611486565b60608901519095506001600160401b03811115611c60575f80fd5b611c6c8a828b0161180a565b9450506080880151611c7d8161163c565b60a0890151909350611c8e8161164a565b60c0890151909250611c9f8161164a565b8091505092959891949750929550565b808202811582820484141761023157610231611ad2565b600181815b80851115611d0057815f1904821115611ce657611ce6611ad2565b80851615611cf357918102915b93841c9390800290611ccb565b509250929050565b5f82611d1657506001610231565b81611d2257505f610231565b8160018114611d385760028114611d4257611d5e565b6001915050610231565b60ff841115611d5357611d53611ad2565b50506001821b610231565b5060208310610133831016604e8410600b8410161715611d81575081810a610231565b611d8b8383611cc6565b805f1904821115611d9e57611d9e611ad2565b029392505050565b5f61097b8383611d08565b8181035f831280158383131683831282161715610cec57610cec611ad256fea164736f6c6343000819000a",
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
