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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ibcAddr\",\"type\":\"address\"},{\"internalType\":\"contractIICS20Bank\",\"name\":\"bank\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"contractITokenRouter\",\"name\":\"tokenRouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"TokenRoutingFailed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"someAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"}],\"name\":\"bindPort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibcAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"ack\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onAcknowledgementPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onRecvPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onTimeoutPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chan\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"chanAddr\",\"type\":\"address\"}],\"name\":\"setChannelEscrowAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"home\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isExternal\",\"type\":\"bool\"}],\"name\":\"setTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenRouter\",\"outputs\":[{\"internalType\":\"contractTokenRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516120bd3803806120bd83398101604081905261002f91610188565b80826001600160a01b03811661006057604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b61006981610120565b506001600160a01b0381166100d55760405162461bcd60e51b815260206004820152602c60248201527f4942434261736546756e6769626c65546f6b656e4170703a207a65726f20726f60448201526b75746572206164647265737360a01b6064820152608401610057565b6001600160a01b03908116608052600280549582166001600160a01b0319968716179055600380549482169486169490941790935560048054919093169316929092179055506101e7565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038116811461018557600080fd5b50565b6000806000806080858703121561019e57600080fd5b84516101a981610170565b60208601519094506101ba81610170565b60408601519093506101cb81610170565b60608601519092506101dc81610170565b939692955090935050565b608051611eb4610209600039600081816101980152610e580152611eb46000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806374a707421161007157806374a70742146101495780638da5cb5b1461015c578063977c6d331461016d5780639d19876514610180578063a562bdba14610193578063f2fde38b146101ba57600080fd5b80633df7d406146100b95780635849f2df146100e15780635aa1e0ff146100f657806363c1bec514610109578063696a9bf41461011c578063715018a614610141575b600080fd5b6100cc6100c736600461148b565b6101cd565b60405190151581526020015b60405180910390f35b6100f46100ef366004611503565b61023e565b005b6100f4610104366004611554565b61028a565b6100cc610117366004611630565b610395565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016100d8565b6100f4610452565b6100f46101573660046116e4565b610466565b6000546001600160a01b0316610129565b6100cc61017b36600461148b565b6104ab565b6100f461018e3660046117bb565b610731565b6101297f000000000000000000000000000000000000000000000000000000000000000081565b6100f46101c8366004611800565b610793565b6002546000906001600160a01b031633146102035760405162461bcd60e51b81526004016101fa9061181d565b60405180910390fd5b60008360a0015180602001905181019061021d91906118e2565b905061023281856020015186604001516107d1565b60019150505b92915050565b6102466108f0565b8060018360405161025791906119c4565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b03199092169190911790555050565b6102b86102a961029a878761091d565b6102a38b6109ab565b906109d8565b6102b28a6109ab565b90610a5d565b6102d6576102d0336102c986610a71565b8a8a610b03565b506102e1565b6102e1338989610b77565b600260009054906101000a90046001600160a01b03166001600160a01b0316636052bf6f6000878787876103388f8f8f8b60405160200161032491815260200190565b604051602081830303815290604052610be2565b6040518763ffffffff1660e01b815260040161035996959493929190611a0c565b600060405180830381600087803b15801561037357600080fd5b505af1158015610387573d6000803e3d6000fd5b505050505050505050505050565b6002546000906001600160a01b031633146103c25760405162461bcd60e51b81526004016101fa9061181d565b6040805180820190915260118152707b22726573756c74223a2241513d3d227d60781b6020918201528351908401207fad671c182373c3087a2d5e877d68648c7b716d38afcdadb5298e477e2f04fb98146104485760008460a0015180602001905181019061043191906118e2565b905061044681866020015187604001516107d1565b505b5060019392505050565b61045a6108f0565b6104646000610c71565b565b61046e6108f0565b60048054604051633a5383a160e11b81526001600160a01b03909116916374a7074291610359918c918c918b918d918c918c918c918c9101611a6f565b6002546000906001600160a01b031633146104d85760405162461bcd60e51b81526004016101fa9061181d565b60008360a001518060200190518101906104f291906118e2565b9050600061050382600001516109ab565b905060006105198660200151876040015161091d565b9050600061052683610cc1565b90506105328383610d30565b610543576105408284610d73565b90505b60008061054f83610e0a565b91509150816105b15760405162461bcd60e51b815260206004820152602860248201527f546f6b656e20726f757465206e6f7420666f756e6420666f722070726566697860448201526765642064656e6f6d60c01b60648201526084016101fa565b8060a0015115610644577f476f2c085b42da8c6a24daf2d3a6d5546fc2cafdd94b1ac39b635fda15cd63a6836040516105ea9190611adc565b60405180910390a18060a00151156106445760405162461bcd60e51b815260206004820152601a60248201527f6e617469766520746f6b656e206e6f7420737570706f7274656400000000000060448201526064016101fa565b602081015160408701516001600160a01b0382169063c1affb129065636f736d6f7360d01b90610675906000610f71565b60405160200161069d919060609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815291905260608b01516106bd906000610f71565b8b602001516106cf8d60800151610fd6565b6040518663ffffffff1660e01b81526004016106ef959493929190611b1d565b600060405180830381600087803b15801561070957600080fd5b505af115801561071d573d6000803e3d6000fd5b5060019d9c50505050505050505050505050565b60405163c13b184f60e01b81526001600160a01b0383169063c13b184f9061075d908490600401611b5c565b600060405180830381600087803b15801561077757600080fd5b505af115801561078b573d6000803e3d6000fd5b505050505050565b61079b6108f0565b6001600160a01b0381166107c557604051631e4fbdf760e01b8152600060048201526024016101fa565b6107ce81610c71565b50565b6107ef6107de838361091d565b84516107e9906109ab565b90610d30565b61087a5761081b6107ff82610a71565b604085015161080f906000610f71565b85516020870151610b03565b6108755760405162461bcd60e51b815260206004820152602560248201527f4261736546756e6769626c65546f6b656e4170703a207472616e736665722066604482015264185a5b195960da1b60648201526084016101fa565b505050565b604083015161089a9061088e906000610f71565b8451602086015161105a565b6108755760405162461bcd60e51b815260206004820152602160248201527f4261736546756e6769626c65546f6b656e4170703a206d696e74206661696c656044820152601960fa1b60648201526084016101fa565b6000546001600160a01b031633146104645760405163118cdaa760e01b81523360048201526024016101fa565b60408051808201909152600080825260208201526109a461099f610959604051806040016040528060018152602001602f60f81b8152506109ab565b61099961099f610968876109ab565b61099961099f610990604051806040016040528060018152602001602f60f81b8152506109ab565b6109998c6109ab565b90610d73565b6109ab565b9392505050565b60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201909152600080825260208201528151835110156109fd575081610238565b6020808301519084015160019114610a245750815160208481015190840151829020919020145b8015610a5557825184518590610a3b908390611b85565b9052508251602085018051610a51908390611b98565b9052505b509192915050565b6000610a6983836110cb565b159392505050565b600080600183604051610a8491906119c4565b908152604051908190036020019020546001600160a01b03169050806102385760405162461bcd60e51b815260206004820152602e60248201527f4261736546756e6769626c65546f6b656e4170703a20657363726f772061646460448201526d1c995cdcc81b9bdd08199bdd5b9960921b60648201526084016101fa565b600354604051637926e0ed60e11b81526000916001600160a01b03169063f24dc1da90610b3a908890889088908890600401611bab565b600060405180830381600087803b158015610b5457600080fd5b505af1158015610b68573d6000803e3d6000fd5b50600198975050505050505050565b60035460405163622db8ef60e11b81526001600160a01b039091169063c45b71de90610bab90869086908690600401611be8565b600060405180830381600087803b158015610bc557600080fd5b505af1158015610bd9573d6000803e3d6000fd5b50505050505050565b60606040518060a00160405280868152602001858152602001610c023390565b604051602001610c2a919060609190911b6bffffffffffffffffffffffff1916815260140190565b604051602081830303815290604052815260200184815260200183815250604051602001610c589190611c1c565b6040516020818303038152906040529050949350505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6060600082600001516001600160401b03811115610ce157610ce161122a565b6040519080825280601f01601f191660200182016040528015610d0b576020820181803683370190505b5090506000602082019050610d2981856020015186600001516111b0565b5092915050565b805182516000911115610d4557506000610238565b8160200151836020015103610d5c57506001610238565b508051602092830151929091015181902091201490565b80518251606091600091610d879190611b98565b6001600160401b03811115610d9e57610d9e61122a565b6040519080825280601f01601f191660200182016040528015610dc8576020820181803683370190505b5090506000602082019050610de681866020015187600001516111b0565b8451610e0290610df69083611b98565b602086015186516111b0565b509392505050565b6040805160e0810182526000808252602082018190529181018290526060808201526080810182905260a0810182905260c081018290526040516381f868f760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906381f868f790610e8d908690600401611b5c565b600060405180830381865afa925050508015610ecb57506040513d6000823e601f3d908101601f19168201604052610ec89190810190611ca4565b60015b610f1d5750506040805160e08101825260008082526020808301829052828401829052835190810190935280835260608201929092526080810182905260a0810182905260c081018290529092909150565b6040805160e0810182526001600160a01b03988916815296881660208801529490961693850193909352606084019190915260ff166080830152151560a082015290151560c0820152600194909350915050565b6000610f7e826014611b98565b83511015610fc65760405162461bcd60e51b8152602060048201526015602482015274746f416464726573735f6f75744f66426f756e647360581b60448201526064016101fa565b500160200151600160601b900490565b60008151600003610fe957506000919050565b6020825111156110525760405162461bcd60e51b815260206004820152602e60248201527f6279746573546f427974657333323a20736f75726365206c656e67746820657860448201526d636565647320333220627974657360901b60648201526084016101fa565b506020015190565b60035460405163ba7aef4360e01b81526000916001600160a01b03169063ba7aef439061108f90879087908790600401611be8565b600060405180830381600087803b1580156110a957600080fd5b505af11580156110bd573d6000803e3d6000fd5b506001979650505050505050565b81518151600091908111156110de575081515b6020808501519084015160005b83811015611197578251825180821461116757600019602087101561114657600184611118896020611b85565b6111229190611b98565b61112d906008611d57565b611138906002611e52565b6111429190611b85565b1990505b81811683821681810391146111645797506102389650505050505050565b50505b611172602086611b98565b945061117f602085611b98565b935050506020816111909190611b98565b90506110eb565b50845186516111a69190611e5e565b9695505050505050565b602081106111e857815183526111c7602084611b98565b92506111d4602083611b98565b91506111e1602082611b85565b90506111b0565b60001981156112175760016111fe836020611b85565b61120a90610100611e52565b6112149190611b85565b90505b9151835183169219169190911790915250565b634e487b7160e01b600052604160045260246000fd5b60405161010081016001600160401b03811182821017156112635761126361122a565b60405290565b60405160a081016001600160401b03811182821017156112635761126361122a565b604051601f8201601f191681016001600160401b03811182821017156112b3576112b361122a565b604052919050565b60006001600160401b038211156112d4576112d461122a565b50601f01601f191660200190565b600082601f8301126112f357600080fd5b8135611306611301826112bb565b61128b565b81815284602083860101111561131b57600080fd5b816020850160208301376000918101602001919091529392505050565b60006040828403121561134a57600080fd5b604051604081018181106001600160401b038211171561136c5761136c61122a565b604052823581526020928301359281019290925250919050565b6000610120828403121561139957600080fd5b6113a1611240565b90508135815260208201356001600160401b03808211156113c157600080fd5b6113cd858386016112e2565b602084015260408401359150808211156113e657600080fd5b6113f2858386016112e2565b6040840152606084013591508082111561140b57600080fd5b611417858386016112e2565b6060840152608084013591508082111561143057600080fd5b61143c858386016112e2565b608084015260a084013591508082111561145557600080fd5b50611462848285016112e2565b60a0830152506114758360c08401611338565b60c082015261010082013560e082015292915050565b6000806040838503121561149e57600080fd5b82356001600160401b03808211156114b557600080fd5b6114c186838701611386565b935060208501359150808211156114d757600080fd5b506114e4858286016112e2565b9150509250929050565b6001600160a01b03811681146107ce57600080fd5b6000806040838503121561151657600080fd5b82356001600160401b0381111561152c57600080fd5b611538858286016112e2565b9250506020830135611549816114ee565b809150509250929050565b600080600080600080600080610120898b03121561157157600080fd5b88356001600160401b038082111561158857600080fd5b6115948c838d016112e2565b995060208b0135985060408b01359150808211156115b157600080fd5b6115bd8c838d016112e2565b975060608b01359150808211156115d357600080fd5b6115df8c838d016112e2565b965060808b01359150808211156115f557600080fd5b506116028b828c016112e2565b9450506116128a60a08b01611338565b925060e0890135915061010089013590509295985092959890939650565b60008060006060848603121561164557600080fd5b83356001600160401b038082111561165c57600080fd5b61166887838801611386565b9450602086013591508082111561167e57600080fd5b61168a878388016112e2565b935060408601359150808211156116a057600080fd5b506116ad868287016112e2565b9150509250925092565b60ff811681146107ce57600080fd5b80151581146107ce57600080fd5b80356116df816116c6565b919050565b600080600080600080600080610100898b03121561170157600080fd5b88356001600160401b038082111561171857600080fd5b6117248c838d016112e2565b995060208b01359150611736826114ee565b90975060408a013590611748826114ee565b90965060608a01359061175a826114ee565b90955060808a0135908082111561177057600080fd5b5061177d8b828c016112e2565b94505060a089013561178e816116b7565b925060c089013561179e816116c6565b91506117ac60e08a016116d4565b90509295985092959890939650565b600080604083850312156117ce57600080fd5b82356117d9816114ee565b915060208301356001600160401b038111156117f457600080fd5b6114e4858286016112e2565b60006020828403121561181257600080fd5b81356109a4816114ee565b60208082526034908201527f4261736546756e6769626c65546f6b656e4170703a2063616c6c6572206973206040820152731b9bdd081d1a1948125090c818dbdb9d1c9858dd60621b606082015260800190565b60005b8381101561188c578181015183820152602001611874565b50506000910152565b600082601f8301126118a657600080fd5b81516118b4611301826112bb565b8181528460208386010111156118c957600080fd5b6118da826020830160208701611871565b949350505050565b6000602082840312156118f457600080fd5b81516001600160401b038082111561190b57600080fd5b9083019060a0828603121561191f57600080fd5b611927611269565b82518281111561193657600080fd5b61194287828601611895565b8252506020830151602082015260408301518281111561196157600080fd5b61196d87828601611895565b60408301525060608301518281111561198557600080fd5b61199187828601611895565b6060830152506080830151828111156119a957600080fd5b6119b587828601611895565b60808301525095945050505050565b600082516119d6818460208701611871565b9190910192915050565b600081518084526119f8816020860160208601611871565b601f01601f19169290920160200192915050565b86815260e060208201526000611a2560e08301886119e0565b8281036040840152611a3781886119e0565b905085516060840152602086015160808401528460a084015282810360c0840152611a6281856119e0565b9998505050505050505050565b6000610100808352611a838184018c6119e0565b6001600160a01b038b811660208601528a811660408601528916606085015283810360808501529050611ab681886119e0565b60ff9690961660a0840152505091151560c0830152151560e09091015295945050505050565b604081526000611aef60408301846119e0565b8281036020840152600c81526b3730ba34bb32903a37b5b2b760a11b60208201526040810191505092915050565b85815260a060208201526000611b3660a08301876119e0565b6001600160a01b0395909516604083015250606081019290925260809091015292915050565b6020815260006109a460208301846119e0565b634e487b7160e01b600052601160045260246000fd5b8181038181111561023857610238611b6f565b8082018082111561023857610238611b6f565b6001600160a01b03858116825284166020820152608060408201819052600090611bd7908301856119e0565b905082606083015295945050505050565b6001600160a01b0384168152606060208201819052600090611c0c908301856119e0565b9050826040830152949350505050565b602081526000825160a06020840152611c3860c08401826119e0565b9050602084015160408401526040840151601f1980858403016060860152611c6083836119e0565b92506060860151915080858403016080860152611c7d83836119e0565b925060808601519150808584030160a086015250611c9b82826119e0565b95945050505050565b600080600080600080600060e0888a031215611cbf57600080fd5b8751611cca816114ee565b6020890151909750611cdb816114ee565b6040890151909650611cec816114ee565b60608901519095506001600160401b03811115611d0857600080fd5b611d148a828b01611895565b9450506080880151611d25816116b7565b60a0890151909350611d36816116c6565b60c0890151909250611d47816116c6565b8091505092959891949750929550565b808202811582820484141761023857610238611b6f565b600181815b80851115611da9578160001904821115611d8f57611d8f611b6f565b80851615611d9c57918102915b93841c9390800290611d73565b509250929050565b600082611dc057506001610238565b81611dcd57506000610238565b8160018114611de35760028114611ded57611e09565b6001915050610238565b60ff841115611dfe57611dfe611b6f565b50506001821b610238565b5060208310610133831016604e8410600b8410161715611e2c575081810a610238565b611e368383611d6e565b8060001904821115611e4a57611e4a611b6f565b029392505050565b60006109a48383611db1565b8181036000831280158383131683831282161715610d2957610d29611b6f56fea2646970667358221220e82fe2a79f5d6ab4c7c15558cc6cd187241873894c481fe9d34f4fcca24c94f564736f6c63430008190033",
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
