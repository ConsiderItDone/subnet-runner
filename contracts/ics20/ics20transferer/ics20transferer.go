// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ics20transferer

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

// ICS20TransfererMetaData contains all meta data concerning the ICS20Transferer contract.
var ICS20TransfererMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ibcAddr\",\"type\":\"address\"},{\"internalType\":\"contractIICS20Bank\",\"name\":\"_bank\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"ack\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"OnAcknowledgementPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"OnRecvPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"someAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"}],\"name\":\"bindPort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibcAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chan\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"chanAddr\",\"type\":\"address\"}],\"name\":\"setChannelEscrowAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"channelCapability\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620023d5380380620023d58339818101604052810190620000379190620001db565b620000576200004b620000e160201b60201c565b620000e960201b60201c565b81600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000298565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050620001be8162000264565b92915050565b600081519050620001d5816200027e565b92915050565b60008060408385031215620001ef57600080fd5b6000620001ff85828601620001ad565b92505060206200021285828601620001c4565b9150509250929050565b6000620002298262000244565b9050919050565b60006200023d826200021c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200026f816200021c565b81146200027b57600080fd5b50565b620002898162000230565b81146200029557600080fd5b50565b61212d80620002a86000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063715018a611610066578063715018a61461011e57806385f7175c146101285780638da5cb5b146101585780639d19876514610176578063f2fde38b1461019257610093565b80632c49a978146100985780632c4a1bee146100c85780635849f2df146100e4578063696a9bf414610100575b600080fd5b6100b260048036038101906100ad91906115f4565b6101ae565b6040516100bf91906119b2565b60405180910390f35b6100e260048036038101906100dd919061168b565b610275565b005b6100fe60048036038101906100f991906114f3565b610314565b005b61010861037d565b604051610115919061190d565b60405180910390f35b6101266103a7565b005b610142600480360381019061013d9190611588565b6103bb565b60405161014f91906119b2565b60405180910390f35b61016061054d565b60405161016d919061190d565b60405180910390f35b610190600480360381019061018b919061149f565b610576565b005b6101ac60048036038101906101a79190611476565b6105e5565b005b60006101b8610669565b73ffffffffffffffffffffffffffffffffffffffff166101d661037d565b73ffffffffffffffffffffffffffffffffffffffff161461022c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022390611a0f565b60405180910390fd5b61023583610671565b61026a5760008460a001518060200190518101906102539190611547565b905061026881866020015187604001516106c0565b505b600190509392505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636052bf6f8787878787876040518763ffffffff1660e01b81526004016102da96959493929190611a6f565b600060405180830381600087803b1580156102f457600080fd5b505af1158015610308573d6000803e3d6000fd5b50505050505050505050565b61031c610766565b8060018360405161032d91906118f6565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6103af610766565b6103b960006107e4565b565b60006103c5610669565b73ffffffffffffffffffffffffffffffffffffffff166103e361037d565b73ffffffffffffffffffffffffffffffffffffffff1614610439576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043090611a0f565b60405180910390fd5b60008360a001518060200190518101906104539190611547565b9050600061046482600001516108a8565b9050600061049861047d876020015188604001516108d6565b61048a85600001516108a8565b6109bf90919063ffffffff16565b90506104ad8183610a5990919063ffffffff16565b6104f2576104e86104c18760400151610a6f565b6104d960008660600151610af690919063ffffffff16565b85600001518660200151610b6c565b9350505050610547565b60006105188361050a89606001518a608001516108d6565b610c0d90919063ffffffff16565b905061054061053560008660600151610af690919063ffffffff16565b828660200151610ce1565b9450505050505b92915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b8173ffffffffffffffffffffffffffffffffffffffff1663c13b184f826040518263ffffffff1660e01b81526004016105af91906119cd565b600060405180830381600087803b1580156105c957600080fd5b505af11580156105dd573d6000803e3d6000fd5b505050505050565b6105ed610766565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561065d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610654906119ef565b60405180910390fd5b610666816107e4565b50565b600033905090565b60006040518060400160405280601181526020017f7b22726573756c74223a2241513d3d227d000000000000000000000000000000815250805190602001208280519060200120149050919050565b6106e86106cd83836108d6565b6106da85600001516108a8565b610d7f90919063ffffffff16565b61072d5761071f6106f882610a6f565b61071060008660400151610af690919063ffffffff16565b85600001518660200151610b6c565b61072857600080fd5b610761565b61075761074860008560400151610af690919063ffffffff16565b84600001518560200151610ce1565b61076057600080fd5b5b505050565b61076e610669565b73ffffffffffffffffffffffffffffffffffffffff1661078c61054d565b73ffffffffffffffffffffffffffffffffffffffff16146107e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d990611a2f565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6108b0610fbe565b600060208301905060405180604001604052808451815260200182815250915050919050565b6108de610fbe565b6109b76109b26109226040518060400160405280600181526020017f2f000000000000000000000000000000000000000000000000000000000000008152506108a8565b6109a461099f610931876108a8565b61099161098c6109756040518060400160405280600181526020017f2f000000000000000000000000000000000000000000000000000000000000008152506108a8565b61097e8c6108a8565b610c0d90919063ffffffff16565b6108a8565b610c0d90919063ffffffff16565b6108a8565b610c0d90919063ffffffff16565b6108a8565b905092915050565b6109c7610fbe565b8160000151836000015110156109df57829050610a53565b6000600190508260200151846020015114610a0d578251602085015160208501518281208383201493505050505b8015610a4e57826000015184600001818151610a299190611e64565b91508181525050826000015184602001818151610a469190611baf565b915081815250505b839150505b92915050565b600080610a668484610dd8565b14905092915050565b600080600183604051610a8291906118f6565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610aed57600080fd5b80915050919050565b6000601482610b059190611baf565b83511015610b48576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3f90611a4f565b60405180910390fd5b60006c01000000000000000000000000836020860101510490508091505092915050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f24dc1da868686866040518563ffffffff1660e01b8152600401610bcf9493929190611928565b600060405180830381600087803b158015610be957600080fd5b505af1158015610bfd573d6000803e3d6000fd5b5050505060019050949350505050565b6060600082600001518460000151610c259190611baf565b67ffffffffffffffff811115610c64577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015610c965781602001600182028036833780820191505090505b5090506000602082019050610cb48186602001518760000151610f12565b610cd6856000015182610cc79190611baf565b85602001518660000151610f12565b819250505092915050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ba7aef438585856040518463ffffffff1660e01b8152600401610d4293929190611974565b600060405180830381600087803b158015610d5c57600080fd5b505af1158015610d70573d6000803e3d6000fd5b50505050600190509392505050565b6000816000015183600001511015610d9a5760009050610dd2565b816020015183602001511415610db35760019050610dd2565b6000825160208501516020850151828120838320149350505050809150505b92915050565b60008083600001519050836000015183600001511015610dfa57826000015190505b60008460200151905060008460200151905060005b83811015610ef1576000808451915083519050808214610ebd5760007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90506020871015610e9757600184886020610e679190611e64565b610e719190611baf565b6008610e7d9190611d76565b6002610e899190611c58565b610e939190611e64565b1990505b600081831682851603905060008114610eba578098505050505050505050610f0c565b50505b602085610eca9190611baf565b9450602084610ed99190611baf565b93505050602081610eea9190611baf565b9050610e0f565b5084600001518660000151610f069190611dd0565b93505050505b92915050565b5b60208110610f515781518352602083610f2c9190611baf565b9250602082610f3b9190611baf565b9150602081610f4a9190611e64565b9050610f13565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90506000821115610fa8576001826020610f8e9190611e64565b610100610f9b9190611c58565b610fa59190611e64565b90505b8019835116818551168181178652505050505050565b604051806040016040528060008152602001600081525090565b6000610feb610fe684611b0a565b611ae5565b90508281526020810184848401111561100357600080fd5b61100e848285611eea565b509392505050565b600061102961102484611b0a565b611ae5565b90508281526020810184848401111561104157600080fd5b61104c848285611ef9565b509392505050565b600061106761106284611b3b565b611ae5565b90508281526020810184848401111561107f57600080fd5b61108a848285611eea565b509392505050565b60006110a56110a084611b3b565b611ae5565b9050828152602081018484840111156110bd57600080fd5b6110c8848285611ef9565b509392505050565b6000813590506110df816120c9565b92915050565b600082601f8301126110f657600080fd5b8135611106848260208601610fd8565b91505092915050565b600082601f83011261112057600080fd5b8151611130848260208601611016565b91505092915050565b600082601f83011261114a57600080fd5b813561115a848260208601611054565b91505092915050565b600082601f83011261117457600080fd5b8151611184848260208601611092565b91505092915050565b600060a0828403121561119f57600080fd5b6111a960a0611ae5565b9050600082015167ffffffffffffffff8111156111c557600080fd5b6111d184828501611163565b60008301525060206111e584828501611461565b602083015250604082015167ffffffffffffffff81111561120557600080fd5b6112118482850161110f565b604083015250606082015167ffffffffffffffff81111561123157600080fd5b61123d8482850161110f565b606083015250608082015167ffffffffffffffff81111561125d57600080fd5b6112698482850161110f565b60808301525092915050565b60006040828403121561128757600080fd5b6112916040611ae5565b905060006112a18482850161144c565b60008301525060206112b58482850161144c565b60208301525092915050565b6000604082840312156112d357600080fd5b6112dd6040611ae5565b905060006112ed8482850161144c565b60008301525060206113018482850161144c565b60208301525092915050565b6000610120828403121561132057600080fd5b61132b610100611ae5565b9050600061133b8482850161144c565b600083015250602082013567ffffffffffffffff81111561135b57600080fd5b61136784828501611139565b602083015250604082013567ffffffffffffffff81111561138757600080fd5b61139384828501611139565b604083015250606082013567ffffffffffffffff8111156113b357600080fd5b6113bf84828501611139565b606083015250608082013567ffffffffffffffff8111156113df57600080fd5b6113eb84828501611139565b60808301525060a082013567ffffffffffffffff81111561140b57600080fd5b611417848285016110e5565b60a08301525060c061142b848285016112c1565b60c0830152506101006114408482850161144c565b60e08301525092915050565b60008135905061145b816120e0565b92915050565b600081519050611470816120e0565b92915050565b60006020828403121561148857600080fd5b6000611496848285016110d0565b91505092915050565b600080604083850312156114b257600080fd5b60006114c0858286016110d0565b925050602083013567ffffffffffffffff8111156114dd57600080fd5b6114e985828601611139565b9150509250929050565b6000806040838503121561150657600080fd5b600083013567ffffffffffffffff81111561152057600080fd5b61152c85828601611139565b925050602061153d858286016110d0565b9150509250929050565b60006020828403121561155957600080fd5b600082015167ffffffffffffffff81111561157357600080fd5b61157f8482850161118d565b91505092915050565b6000806040838503121561159b57600080fd5b600083013567ffffffffffffffff8111156115b557600080fd5b6115c18582860161130d565b925050602083013567ffffffffffffffff8111156115de57600080fd5b6115ea858286016110e5565b9150509250929050565b60008060006060848603121561160957600080fd5b600084013567ffffffffffffffff81111561162357600080fd5b61162f8682870161130d565b935050602084013567ffffffffffffffff81111561164c57600080fd5b611658868287016110e5565b925050604084013567ffffffffffffffff81111561167557600080fd5b611681868287016110e5565b9150509250925092565b60008060008060008060e087890312156116a457600080fd5b60006116b289828a0161144c565b965050602087013567ffffffffffffffff8111156116cf57600080fd5b6116db89828a01611139565b955050604087013567ffffffffffffffff8111156116f857600080fd5b61170489828a01611139565b945050606061171589828a01611275565b93505060a061172689828a0161144c565b92505060c087013567ffffffffffffffff81111561174357600080fd5b61174f89828a016110e5565b9150509295509295509295565b61176581611e98565b82525050565b61177481611eaa565b82525050565b600061178582611b6c565b61178f8185611b82565b935061179f818560208601611ef9565b6117a881611fbb565b840191505092915050565b60006117be82611b77565b6117c88185611b93565b93506117d8818560208601611ef9565b6117e181611fbb565b840191505092915050565b60006117f782611b77565b6118018185611ba4565b9350611811818560208601611ef9565b80840191505092915050565b600061182a602683611b93565b915061183582611fd9565b604082019050919050565b600061184d603483611b93565b915061185882612028565b604082019050919050565b6000611870602083611b93565b915061187b82612077565b602082019050919050565b6000611893601583611b93565b915061189e826120a0565b602082019050919050565b6040820160008201516118bf60008501826118d8565b5060208201516118d260208501826118d8565b50505050565b6118e181611ee0565b82525050565b6118f081611ee0565b82525050565b600061190282846117ec565b915081905092915050565b6000602082019050611922600083018461175c565b92915050565b600060808201905061193d600083018761175c565b61194a602083018661175c565b818103604083015261195c81856117b3565b905061196b60608301846118e7565b95945050505050565b6000606082019050611989600083018661175c565b818103602083015261199b81856117b3565b90506119aa60408301846118e7565b949350505050565b60006020820190506119c7600083018461176b565b92915050565b600060208201905081810360008301526119e781846117b3565b905092915050565b60006020820190508181036000830152611a088161181d565b9050919050565b60006020820190508181036000830152611a2881611840565b9050919050565b60006020820190508181036000830152611a4881611863565b9050919050565b60006020820190508181036000830152611a6881611886565b9050919050565b600060e082019050611a8460008301896118e7565b8181036020830152611a9681886117b3565b90508181036040830152611aaa81876117b3565b9050611ab960608301866118a9565b611ac660a08301856118e7565b81810360c0830152611ad8818461177a565b9050979650505050505050565b6000611aef611b00565b9050611afb8282611f2c565b919050565b6000604051905090565b600067ffffffffffffffff821115611b2557611b24611f8c565b5b611b2e82611fbb565b9050602081019050919050565b600067ffffffffffffffff821115611b5657611b55611f8c565b5b611b5f82611fbb565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611bba82611ee0565b9150611bc583611ee0565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611bfa57611bf9611f5d565b5b828201905092915050565b6000808291508390505b6001851115611c4f57808604811115611c2b57611c2a611f5d565b5b6001851615611c3a5780820291505b8081029050611c4885611fcc565b9450611c0f565b94509492505050565b6000611c6382611ee0565b9150611c6e83611ee0565b9250611c9b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611ca3565b905092915050565b600082611cb35760019050611d6f565b81611cc15760009050611d6f565b8160018114611cd75760028114611ce157611d10565b6001915050611d6f565b60ff841115611cf357611cf2611f5d565b5b8360020a915084821115611d0a57611d09611f5d565b5b50611d6f565b5060208310610133831016604e8410600b8410161715611d455782820a905083811115611d4057611d3f611f5d565b5b611d6f565b611d528484846001611c05565b92509050818404811115611d6957611d68611f5d565b5b81810290505b9392505050565b6000611d8182611ee0565b9150611d8c83611ee0565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611dc557611dc4611f5d565b5b828202905092915050565b6000611ddb82611eb6565b9150611de683611eb6565b9250827f800000000000000000000000000000000000000000000000000000000000000001821260008412151615611e2157611e20611f5d565b5b827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018213600084121615611e5957611e58611f5d565b5b828203905092915050565b6000611e6f82611ee0565b9150611e7a83611ee0565b925082821015611e8d57611e8c611f5d565b5b828203905092915050565b6000611ea382611ec0565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611f17578082015181840152602081019050611efc565b83811115611f26576000848401525b50505050565b611f3582611fbb565b810181811067ffffffffffffffff82111715611f5457611f53611f8c565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160011c9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f4261736546756e6769626c65546f6b656e4170703a2063616c6c65722069732060008201527f6e6f74207468652049424320636f6e7472616374000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f746f416464726573735f6f75744f66426f756e64730000000000000000000000600082015250565b6120d281611e98565b81146120dd57600080fd5b50565b6120e981611ee0565b81146120f457600080fd5b5056fea2646970667358221220d6438991728140a1c12fc70aab98fafef99128cb77cea053beee5c26e965376564736f6c63430008010033",
}

// ICS20TransfererABI is the input ABI used to generate the binding from.
// Deprecated: Use ICS20TransfererMetaData.ABI instead.
var ICS20TransfererABI = ICS20TransfererMetaData.ABI

// ICS20TransfererBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ICS20TransfererMetaData.Bin instead.
var ICS20TransfererBin = ICS20TransfererMetaData.Bin

// DeployICS20Transferer deploys a new Ethereum contract, binding an instance of ICS20Transferer to it.
func DeployICS20Transferer(auth *bind.TransactOpts, backend bind.ContractBackend, _ibcAddr common.Address, _bank common.Address) (common.Address, *types.Transaction, *ICS20Transferer, error) {
	parsed, err := ICS20TransfererMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ICS20TransfererBin), backend, _ibcAddr, _bank)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ICS20Transferer{ICS20TransfererCaller: ICS20TransfererCaller{contract: contract}, ICS20TransfererTransactor: ICS20TransfererTransactor{contract: contract}, ICS20TransfererFilterer: ICS20TransfererFilterer{contract: contract}}, nil
}

// ICS20Transferer is an auto generated Go binding around an Ethereum contract.
type ICS20Transferer struct {
	ICS20TransfererCaller     // Read-only binding to the contract
	ICS20TransfererTransactor // Write-only binding to the contract
	ICS20TransfererFilterer   // Log filterer for contract events
}

// ICS20TransfererCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICS20TransfererCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS20TransfererTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICS20TransfererTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS20TransfererFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICS20TransfererFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS20TransfererSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICS20TransfererSession struct {
	Contract     *ICS20Transferer  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICS20TransfererCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICS20TransfererCallerSession struct {
	Contract *ICS20TransfererCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ICS20TransfererTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICS20TransfererTransactorSession struct {
	Contract     *ICS20TransfererTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ICS20TransfererRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICS20TransfererRaw struct {
	Contract *ICS20Transferer // Generic contract binding to access the raw methods on
}

// ICS20TransfererCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICS20TransfererCallerRaw struct {
	Contract *ICS20TransfererCaller // Generic read-only contract binding to access the raw methods on
}

// ICS20TransfererTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICS20TransfererTransactorRaw struct {
	Contract *ICS20TransfererTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICS20Transferer creates a new instance of ICS20Transferer, bound to a specific deployed contract.
func NewICS20Transferer(address common.Address, backend bind.ContractBackend) (*ICS20Transferer, error) {
	contract, err := bindICS20Transferer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICS20Transferer{ICS20TransfererCaller: ICS20TransfererCaller{contract: contract}, ICS20TransfererTransactor: ICS20TransfererTransactor{contract: contract}, ICS20TransfererFilterer: ICS20TransfererFilterer{contract: contract}}, nil
}

// NewICS20TransfererCaller creates a new read-only instance of ICS20Transferer, bound to a specific deployed contract.
func NewICS20TransfererCaller(address common.Address, caller bind.ContractCaller) (*ICS20TransfererCaller, error) {
	contract, err := bindICS20Transferer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICS20TransfererCaller{contract: contract}, nil
}

// NewICS20TransfererTransactor creates a new write-only instance of ICS20Transferer, bound to a specific deployed contract.
func NewICS20TransfererTransactor(address common.Address, transactor bind.ContractTransactor) (*ICS20TransfererTransactor, error) {
	contract, err := bindICS20Transferer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICS20TransfererTransactor{contract: contract}, nil
}

// NewICS20TransfererFilterer creates a new log filterer instance of ICS20Transferer, bound to a specific deployed contract.
func NewICS20TransfererFilterer(address common.Address, filterer bind.ContractFilterer) (*ICS20TransfererFilterer, error) {
	contract, err := bindICS20Transferer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICS20TransfererFilterer{contract: contract}, nil
}

// bindICS20Transferer binds a generic wrapper to an already deployed contract.
func bindICS20Transferer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICS20TransfererMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICS20Transferer *ICS20TransfererRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICS20Transferer.Contract.ICS20TransfererCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICS20Transferer *ICS20TransfererRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.ICS20TransfererTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICS20Transferer *ICS20TransfererRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.ICS20TransfererTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICS20Transferer *ICS20TransfererCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICS20Transferer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICS20Transferer *ICS20TransfererTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICS20Transferer *ICS20TransfererTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.contract.Transact(opts, method, params...)
}

// IbcAddress is a free data retrieval call binding the contract method 0x696a9bf4.
//
// Solidity: function ibcAddress() view returns(address)
func (_ICS20Transferer *ICS20TransfererCaller) IbcAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICS20Transferer.contract.Call(opts, &out, "ibcAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IbcAddress is a free data retrieval call binding the contract method 0x696a9bf4.
//
// Solidity: function ibcAddress() view returns(address)
func (_ICS20Transferer *ICS20TransfererSession) IbcAddress() (common.Address, error) {
	return _ICS20Transferer.Contract.IbcAddress(&_ICS20Transferer.CallOpts)
}

// IbcAddress is a free data retrieval call binding the contract method 0x696a9bf4.
//
// Solidity: function ibcAddress() view returns(address)
func (_ICS20Transferer *ICS20TransfererCallerSession) IbcAddress() (common.Address, error) {
	return _ICS20Transferer.Contract.IbcAddress(&_ICS20Transferer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ICS20Transferer *ICS20TransfererCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICS20Transferer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ICS20Transferer *ICS20TransfererSession) Owner() (common.Address, error) {
	return _ICS20Transferer.Contract.Owner(&_ICS20Transferer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ICS20Transferer *ICS20TransfererCallerSession) Owner() (common.Address, error) {
	return _ICS20Transferer.Contract.Owner(&_ICS20Transferer.CallOpts)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x2c49a978.
//
// Solidity: function OnAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ack, bytes ) returns(bool)
func (_ICS20Transferer *ICS20TransfererTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet Packet, ack []byte, arg2 []byte) (*types.Transaction, error) {
	return _ICS20Transferer.contract.Transact(opts, "OnAcknowledgementPacket", packet, ack, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x2c49a978.
//
// Solidity: function OnAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ack, bytes ) returns(bool)
func (_ICS20Transferer *ICS20TransfererSession) OnAcknowledgementPacket(packet Packet, ack []byte, arg2 []byte) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.OnAcknowledgementPacket(&_ICS20Transferer.TransactOpts, packet, ack, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x2c49a978.
//
// Solidity: function OnAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ack, bytes ) returns(bool)
func (_ICS20Transferer *ICS20TransfererTransactorSession) OnAcknowledgementPacket(packet Packet, ack []byte, arg2 []byte) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.OnAcknowledgementPacket(&_ICS20Transferer.TransactOpts, packet, ack, arg2)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x85f7175c.
//
// Solidity: function OnRecvPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ) returns(bool)
func (_ICS20Transferer *ICS20TransfererTransactor) OnRecvPacket(opts *bind.TransactOpts, packet Packet, arg1 []byte) (*types.Transaction, error) {
	return _ICS20Transferer.contract.Transact(opts, "OnRecvPacket", packet, arg1)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x85f7175c.
//
// Solidity: function OnRecvPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ) returns(bool)
func (_ICS20Transferer *ICS20TransfererSession) OnRecvPacket(packet Packet, arg1 []byte) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.OnRecvPacket(&_ICS20Transferer.TransactOpts, packet, arg1)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x85f7175c.
//
// Solidity: function OnRecvPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) packet, bytes ) returns(bool)
func (_ICS20Transferer *ICS20TransfererTransactorSession) OnRecvPacket(packet Packet, arg1 []byte) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.OnRecvPacket(&_ICS20Transferer.TransactOpts, packet, arg1)
}

// BindPort is a paid mutator transaction binding the contract method 0x9d198765.
//
// Solidity: function bindPort(address someAddr, string portId) returns()
func (_ICS20Transferer *ICS20TransfererTransactor) BindPort(opts *bind.TransactOpts, someAddr common.Address, portId string) (*types.Transaction, error) {
	return _ICS20Transferer.contract.Transact(opts, "bindPort", someAddr, portId)
}

// BindPort is a paid mutator transaction binding the contract method 0x9d198765.
//
// Solidity: function bindPort(address someAddr, string portId) returns()
func (_ICS20Transferer *ICS20TransfererSession) BindPort(someAddr common.Address, portId string) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.BindPort(&_ICS20Transferer.TransactOpts, someAddr, portId)
}

// BindPort is a paid mutator transaction binding the contract method 0x9d198765.
//
// Solidity: function bindPort(address someAddr, string portId) returns()
func (_ICS20Transferer *ICS20TransfererTransactorSession) BindPort(someAddr common.Address, portId string) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.BindPort(&_ICS20Transferer.TransactOpts, someAddr, portId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ICS20Transferer *ICS20TransfererTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICS20Transferer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ICS20Transferer *ICS20TransfererSession) RenounceOwnership() (*types.Transaction, error) {
	return _ICS20Transferer.Contract.RenounceOwnership(&_ICS20Transferer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ICS20Transferer *ICS20TransfererTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ICS20Transferer.Contract.RenounceOwnership(&_ICS20Transferer.TransactOpts)
}

// SetChannelEscrowAddresses is a paid mutator transaction binding the contract method 0x5849f2df.
//
// Solidity: function setChannelEscrowAddresses(string chan, address chanAddr) returns()
func (_ICS20Transferer *ICS20TransfererTransactor) SetChannelEscrowAddresses(opts *bind.TransactOpts, arg0 string, chanAddr common.Address) (*types.Transaction, error) {
	return _ICS20Transferer.contract.Transact(opts, "setChannelEscrowAddresses", arg0, chanAddr)
}

// SetChannelEscrowAddresses is a paid mutator transaction binding the contract method 0x5849f2df.
//
// Solidity: function setChannelEscrowAddresses(string chan, address chanAddr) returns()
func (_ICS20Transferer *ICS20TransfererSession) SetChannelEscrowAddresses(arg0 string, chanAddr common.Address) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.SetChannelEscrowAddresses(&_ICS20Transferer.TransactOpts, arg0, chanAddr)
}

// SetChannelEscrowAddresses is a paid mutator transaction binding the contract method 0x5849f2df.
//
// Solidity: function setChannelEscrowAddresses(string chan, address chanAddr) returns()
func (_ICS20Transferer *ICS20TransfererTransactorSession) SetChannelEscrowAddresses(arg0 string, chanAddr common.Address) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.SetChannelEscrowAddresses(&_ICS20Transferer.TransactOpts, arg0, chanAddr)
}

// Transfer is a paid mutator transaction binding the contract method 0x2c4a1bee.
//
// Solidity: function transfer(uint256 channelCapability, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes data) returns()
func (_ICS20Transferer *ICS20TransfererTransactor) Transfer(opts *bind.TransactOpts, channelCapability *big.Int, sourcePort string, sourceChannel string, timeoutHeight Height, timeoutTimestamp *big.Int, data []byte) (*types.Transaction, error) {
	return _ICS20Transferer.contract.Transact(opts, "transfer", channelCapability, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// Transfer is a paid mutator transaction binding the contract method 0x2c4a1bee.
//
// Solidity: function transfer(uint256 channelCapability, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes data) returns()
func (_ICS20Transferer *ICS20TransfererSession) Transfer(channelCapability *big.Int, sourcePort string, sourceChannel string, timeoutHeight Height, timeoutTimestamp *big.Int, data []byte) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.Transfer(&_ICS20Transferer.TransactOpts, channelCapability, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// Transfer is a paid mutator transaction binding the contract method 0x2c4a1bee.
//
// Solidity: function transfer(uint256 channelCapability, string sourcePort, string sourceChannel, (uint256,uint256) timeoutHeight, uint256 timeoutTimestamp, bytes data) returns()
func (_ICS20Transferer *ICS20TransfererTransactorSession) Transfer(channelCapability *big.Int, sourcePort string, sourceChannel string, timeoutHeight Height, timeoutTimestamp *big.Int, data []byte) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.Transfer(&_ICS20Transferer.TransactOpts, channelCapability, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ICS20Transferer *ICS20TransfererTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ICS20Transferer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ICS20Transferer *ICS20TransfererSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.TransferOwnership(&_ICS20Transferer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ICS20Transferer *ICS20TransfererTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ICS20Transferer.Contract.TransferOwnership(&_ICS20Transferer.TransactOpts, newOwner)
}

// ICS20TransfererOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ICS20Transferer contract.
type ICS20TransfererOwnershipTransferredIterator struct {
	Event *ICS20TransfererOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ICS20TransfererOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICS20TransfererOwnershipTransferred)
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
		it.Event = new(ICS20TransfererOwnershipTransferred)
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
func (it *ICS20TransfererOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICS20TransfererOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICS20TransfererOwnershipTransferred represents a OwnershipTransferred event raised by the ICS20Transferer contract.
type ICS20TransfererOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ICS20Transferer *ICS20TransfererFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ICS20TransfererOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ICS20Transferer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ICS20TransfererOwnershipTransferredIterator{contract: _ICS20Transferer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ICS20Transferer *ICS20TransfererFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ICS20TransfererOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ICS20Transferer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICS20TransfererOwnershipTransferred)
				if err := _ICS20Transferer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ICS20Transferer *ICS20TransfererFilterer) ParseOwnershipTransferred(log types.Log) (*ICS20TransfererOwnershipTransferred, error) {
	event := new(ICS20TransfererOwnershipTransferred)
	if err := _ICS20Transferer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
