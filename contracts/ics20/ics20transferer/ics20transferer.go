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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ibcAddr\",\"type\":\"address\"},{\"internalType\":\"contractIICS20Bank\",\"name\":\"_bank\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"ack\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"OnAcknowledgementPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"OnRecvPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"someAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"}],\"name\":\"bindPort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibcAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chan\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"chanAddr\",\"type\":\"address\"}],\"name\":\"setChannelEscrowAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620020ff380380620020ff8339818101604052810190620000379190620001db565b620000576200004b620000e160201b60201c565b620000e960201b60201c565b81600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000298565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050620001be8162000264565b92915050565b600081519050620001d5816200027e565b92915050565b60008060408385031215620001ef57600080fd5b6000620001ff85828601620001ad565b92505060206200021285828601620001c4565b9150509250929050565b6000620002298262000244565b9050919050565b60006200023d826200021c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200026f816200021c565b81146200027b57600080fd5b50565b620002898162000230565b81146200029557600080fd5b50565b611e5780620002a86000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806385f7175c1161005b57806385f7175c146101015780638da5cb5b146101315780639d1987651461014f578063f2fde38b1461016b57610088565b80632c49a9781461008d5780635849f2df146100bd578063696a9bf4146100d9578063715018a6146100f7575b600080fd5b6100a760048036038101906100a291906114f8565b610187565b6040516100b4919061176e565b60405180910390f35b6100d760048036038101906100d291906113f7565b61024e565b005b6100e16102b7565b6040516100ee91906116c9565b60405180910390f35b6100ff6102e1565b005b61011b6004803603810190610116919061148c565b6102f5565b604051610128919061176e565b60405180910390f35b610139610487565b60405161014691906116c9565b60405180910390f35b610169600480360381019061016491906113a3565b6104b0565b005b6101856004803603810190610180919061137a565b61051f565b005b60006101916105a3565b73ffffffffffffffffffffffffffffffffffffffff166101af6102b7565b73ffffffffffffffffffffffffffffffffffffffff1614610205576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101fc906117cb565b60405180910390fd5b61020e836105ab565b6102435760008460a0015180602001905181019061022c919061144b565b90506102418186602001518760400151610610565b505b600190509392505050565b6102566106b6565b8060018360405161026791906116b2565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6102e96106b6565b6102f36000610734565b565b60006102ff6105a3565b73ffffffffffffffffffffffffffffffffffffffff1661031d6102b7565b73ffffffffffffffffffffffffffffffffffffffff1614610373576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036a906117cb565b60405180910390fd5b60008360a0015180602001905181019061038d919061144b565b9050600061039e82600001516107f8565b905060006103d26103b787602001518860400151610826565b6103c485600001516107f8565b61090f90919063ffffffff16565b90506103e781836109a990919063ffffffff16565b61042c576104226103fb87604001516109bf565b61041360008660600151610a4690919063ffffffff16565b85600001518660200151610abc565b9350505050610481565b60006104528361044489606001518a60800151610826565b610b5d90919063ffffffff16565b905061047a61046f60008660600151610a4690919063ffffffff16565b828660200151610c31565b9450505050505b92915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b8173ffffffffffffffffffffffffffffffffffffffff1663c13b184f826040518263ffffffff1660e01b81526004016104e99190611789565b600060405180830381600087803b15801561050357600080fd5b505af1158015610517573d6000803e3d6000fd5b505050505050565b6105276106b6565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610597576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058e906117ab565b60405180910390fd5b6105a081610734565b50565b600033905090565b600060018251146105bb57600080fd5b6001826000815181106105f7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b60f81c60ff16149050919050565b61063861061d8383610826565b61062a85600001516107f8565b610ccf90919063ffffffff16565b61067d5761066f610648826109bf565b61066060008660400151610a4690919063ffffffff16565b85600001518660200151610abc565b61067857600080fd5b6106b1565b6106a761069860008560400151610a4690919063ffffffff16565b84600001518560200151610c31565b6106b057600080fd5b5b505050565b6106be6105a3565b73ffffffffffffffffffffffffffffffffffffffff166106dc610487565b73ffffffffffffffffffffffffffffffffffffffff1614610732576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610729906117eb565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610800610f0e565b600060208301905060405180604001604052808451815260200182815250915050919050565b61082e610f0e565b6109076109026108726040518060400160405280600181526020017f2f000000000000000000000000000000000000000000000000000000000000008152506107f8565b6108f46108ef610881876107f8565b6108e16108dc6108c56040518060400160405280600181526020017f2f000000000000000000000000000000000000000000000000000000000000008152506107f8565b6108ce8c6107f8565b610b5d90919063ffffffff16565b6107f8565b610b5d90919063ffffffff16565b6107f8565b610b5d90919063ffffffff16565b6107f8565b905092915050565b610917610f0e565b81600001518360000151101561092f578290506109a3565b600060019050826020015184602001511461095d578251602085015160208501518281208383201493505050505b801561099e578260000151846000018181516109799190611b8e565b9150818152505082600001518460200181815161099691906118d9565b915081815250505b839150505b92915050565b6000806109b68484610d28565b14905092915050565b6000806001836040516109d291906116b2565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a3d57600080fd5b80915050919050565b6000601482610a5591906118d9565b83511015610a98576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8f9061180b565b60405180910390fd5b60006c01000000000000000000000000836020860101510490508091505092915050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f24dc1da868686866040518563ffffffff1660e01b8152600401610b1f94939291906116e4565b600060405180830381600087803b158015610b3957600080fd5b505af1158015610b4d573d6000803e3d6000fd5b5050505060019050949350505050565b6060600082600001518460000151610b7591906118d9565b67ffffffffffffffff811115610bb4577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015610be65781602001600182028036833780820191505090505b5090506000602082019050610c048186602001518760000151610e62565b610c26856000015182610c1791906118d9565b85602001518660000151610e62565b819250505092915050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ba7aef438585856040518463ffffffff1660e01b8152600401610c9293929190611730565b600060405180830381600087803b158015610cac57600080fd5b505af1158015610cc0573d6000803e3d6000fd5b50505050600190509392505050565b6000816000015183600001511015610cea5760009050610d22565b816020015183602001511415610d035760019050610d22565b6000825160208501516020850151828120838320149350505050809150505b92915050565b60008083600001519050836000015183600001511015610d4a57826000015190505b60008460200151905060008460200151905060005b83811015610e41576000808451915083519050808214610e0d5760007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90506020871015610de757600184886020610db79190611b8e565b610dc191906118d9565b6008610dcd9190611aa0565b6002610dd99190611982565b610de39190611b8e565b1990505b600081831682851603905060008114610e0a578098505050505050505050610e5c565b50505b602085610e1a91906118d9565b9450602084610e2991906118d9565b93505050602081610e3a91906118d9565b9050610d5f565b5084600001518660000151610e569190611afa565b93505050505b92915050565b5b60208110610ea15781518352602083610e7c91906118d9565b9250602082610e8b91906118d9565b9150602081610e9a9190611b8e565b9050610e63565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90506000821115610ef8576001826020610ede9190611b8e565b610100610eeb9190611982565b610ef59190611b8e565b90505b8019835116818551168181178652505050505050565b604051806040016040528060008152602001600081525090565b6000610f3b610f3684611850565b61182b565b905082815260208101848484011115610f5357600080fd5b610f5e848285611c14565b509392505050565b6000610f79610f7484611850565b61182b565b905082815260208101848484011115610f9157600080fd5b610f9c848285611c23565b509392505050565b6000610fb7610fb284611881565b61182b565b905082815260208101848484011115610fcf57600080fd5b610fda848285611c14565b509392505050565b6000610ff5610ff084611881565b61182b565b90508281526020810184848401111561100d57600080fd5b611018848285611c23565b509392505050565b60008135905061102f81611df3565b92915050565b600082601f83011261104657600080fd5b8135611056848260208601610f28565b91505092915050565b600082601f83011261107057600080fd5b8151611080848260208601610f66565b91505092915050565b600082601f83011261109a57600080fd5b81356110aa848260208601610fa4565b91505092915050565b600082601f8301126110c457600080fd5b81516110d4848260208601610fe2565b91505092915050565b600060a082840312156110ef57600080fd5b6110f960a061182b565b9050600082015167ffffffffffffffff81111561111557600080fd5b611121848285016110b3565b600083015250602061113584828501611365565b602083015250604082015167ffffffffffffffff81111561115557600080fd5b6111618482850161105f565b604083015250606082015167ffffffffffffffff81111561118157600080fd5b61118d8482850161105f565b606083015250608082015167ffffffffffffffff8111156111ad57600080fd5b6111b98482850161105f565b60808301525092915050565b6000604082840312156111d757600080fd5b6111e1604061182b565b905060006111f184828501611350565b600083015250602061120584828501611350565b60208301525092915050565b6000610120828403121561122457600080fd5b61122f61010061182b565b9050600061123f84828501611350565b600083015250602082013567ffffffffffffffff81111561125f57600080fd5b61126b84828501611089565b602083015250604082013567ffffffffffffffff81111561128b57600080fd5b61129784828501611089565b604083015250606082013567ffffffffffffffff8111156112b757600080fd5b6112c384828501611089565b606083015250608082013567ffffffffffffffff8111156112e357600080fd5b6112ef84828501611089565b60808301525060a082013567ffffffffffffffff81111561130f57600080fd5b61131b84828501611035565b60a08301525060c061132f848285016111c5565b60c08301525061010061134484828501611350565b60e08301525092915050565b60008135905061135f81611e0a565b92915050565b60008151905061137481611e0a565b92915050565b60006020828403121561138c57600080fd5b600061139a84828501611020565b91505092915050565b600080604083850312156113b657600080fd5b60006113c485828601611020565b925050602083013567ffffffffffffffff8111156113e157600080fd5b6113ed85828601611089565b9150509250929050565b6000806040838503121561140a57600080fd5b600083013567ffffffffffffffff81111561142457600080fd5b61143085828601611089565b925050602061144185828601611020565b9150509250929050565b60006020828403121561145d57600080fd5b600082015167ffffffffffffffff81111561147757600080fd5b611483848285016110dd565b91505092915050565b6000806040838503121561149f57600080fd5b600083013567ffffffffffffffff8111156114b957600080fd5b6114c585828601611211565b925050602083013567ffffffffffffffff8111156114e257600080fd5b6114ee85828601611035565b9150509250929050565b60008060006060848603121561150d57600080fd5b600084013567ffffffffffffffff81111561152757600080fd5b61153386828701611211565b935050602084013567ffffffffffffffff81111561155057600080fd5b61155c86828701611035565b925050604084013567ffffffffffffffff81111561157957600080fd5b61158586828701611035565b9150509250925092565b61159881611bc2565b82525050565b6115a781611bd4565b82525050565b60006115b8826118b2565b6115c281856118bd565b93506115d2818560208601611c23565b6115db81611ce5565b840191505092915050565b60006115f1826118b2565b6115fb81856118ce565b935061160b818560208601611c23565b80840191505092915050565b60006116246026836118bd565b915061162f82611d03565b604082019050919050565b60006116476034836118bd565b915061165282611d52565b604082019050919050565b600061166a6020836118bd565b915061167582611da1565b602082019050919050565b600061168d6015836118bd565b915061169882611dca565b602082019050919050565b6116ac81611c0a565b82525050565b60006116be82846115e6565b915081905092915050565b60006020820190506116de600083018461158f565b92915050565b60006080820190506116f9600083018761158f565b611706602083018661158f565b818103604083015261171881856115ad565b905061172760608301846116a3565b95945050505050565b6000606082019050611745600083018661158f565b818103602083015261175781856115ad565b905061176660408301846116a3565b949350505050565b6000602082019050611783600083018461159e565b92915050565b600060208201905081810360008301526117a381846115ad565b905092915050565b600060208201905081810360008301526117c481611617565b9050919050565b600060208201905081810360008301526117e48161163a565b9050919050565b600060208201905081810360008301526118048161165d565b9050919050565b6000602082019050818103600083015261182481611680565b9050919050565b6000611835611846565b90506118418282611c56565b919050565b6000604051905090565b600067ffffffffffffffff82111561186b5761186a611cb6565b5b61187482611ce5565b9050602081019050919050565b600067ffffffffffffffff82111561189c5761189b611cb6565b5b6118a582611ce5565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006118e482611c0a565b91506118ef83611c0a565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561192457611923611c87565b5b828201905092915050565b6000808291508390505b60018511156119795780860481111561195557611954611c87565b5b60018516156119645780820291505b808102905061197285611cf6565b9450611939565b94509492505050565b600061198d82611c0a565b915061199883611c0a565b92506119c57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846119cd565b905092915050565b6000826119dd5760019050611a99565b816119eb5760009050611a99565b8160018114611a015760028114611a0b57611a3a565b6001915050611a99565b60ff841115611a1d57611a1c611c87565b5b8360020a915084821115611a3457611a33611c87565b5b50611a99565b5060208310610133831016604e8410600b8410161715611a6f5782820a905083811115611a6a57611a69611c87565b5b611a99565b611a7c848484600161192f565b92509050818404811115611a9357611a92611c87565b5b81810290505b9392505050565b6000611aab82611c0a565b9150611ab683611c0a565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611aef57611aee611c87565b5b828202905092915050565b6000611b0582611be0565b9150611b1083611be0565b9250827f800000000000000000000000000000000000000000000000000000000000000001821260008412151615611b4b57611b4a611c87565b5b827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018213600084121615611b8357611b82611c87565b5b828203905092915050565b6000611b9982611c0a565b9150611ba483611c0a565b925082821015611bb757611bb6611c87565b5b828203905092915050565b6000611bcd82611bea565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611c41578082015181840152602081019050611c26565b83811115611c50576000848401525b50505050565b611c5f82611ce5565b810181811067ffffffffffffffff82111715611c7e57611c7d611cb6565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160011c9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f4261736546756e6769626c65546f6b656e4170703a2063616c6c65722069732060008201527f6e6f74207468652049424320636f6e7472616374000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f746f416464726573735f6f75744f66426f756e64730000000000000000000000600082015250565b611dfc81611bc2565b8114611e0757600080fd5b50565b611e1381611c0a565b8114611e1e57600080fd5b5056fea26469706673582212202880de9312cec9fd80920ffa79a2f663083969c45453735acd2bc69d4c684c4e64736f6c63430008010033",
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
