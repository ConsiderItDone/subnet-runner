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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ibcAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"contractITokenRouter\",\"name\":\"tokenRouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"TokenRoutingFailed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"someAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"}],\"name\":\"bindPort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibcAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onAcknowledgementPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onRecvPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationPort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destinationChannel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPacket\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onTimeoutPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chan\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"chanAddr\",\"type\":\"address\"}],\"name\":\"setChannelEscrowAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenRouter\",\"outputs\":[{\"internalType\":\"contractTokenRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"revisionNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revisionHeight\",\"type\":\"uint256\"}],\"internalType\":\"structHeight\",\"name\":\"timeoutHeight\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timeoutTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b506040516117f13803806117f183398101604081905261002e91610173565b80826001600160a01b03811661005e57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b6100678161010d565b506001600160a01b0381166100d35760405162461bcd60e51b815260206004820152602c60248201527f4942434261736546756e6769626c65546f6b656e4170703a207a65726f20726f60448201526b75746572206164647265737360a01b6064820152608401610055565b6001600160a01b03908116608052600280549482166001600160a01b031995861617905560038054929091169190931617909155506101bd565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0381168114610170575f80fd5b50565b5f805f60608486031215610185575f80fd5b83516101908161015c565b60208501519093506101a18161015c565b60408501519092506101b28161015c565b809150509250925092565b6080516116156101dc5f395f8181610175015261096c01526116155ff3fe608060405234801561000f575f80fd5b50600436106100a6575f3560e01c8063715018a61161006e578063715018a6146101325780638da5cb5b1461013a578063977c6d331461014a5780639d1987651461015d578063a562bdba14610170578063f2fde38b14610197575f80fd5b80633df7d406146100aa5780635849f2df146100d25780635aa1e0ff146100e757806363c1bec5146100fa578063696a9bf41461010d575b5f80fd5b6100bd6100b8366004610e31565b6101aa565b60405190151581526020015b60405180910390f35b6100e56100e0366004610ea4565b6101e9565b005b6100e56100f5366004610ef2565b610235565b6100bd610108366004610fc5565b6102e2565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016100c9565b6100e5610318565b5f546001600160a01b031661011a565b6100bd610158366004610e31565b61032b565b6100e561016b366004611046565b610582565b61011a7f000000000000000000000000000000000000000000000000000000000000000081565b6100e56101a5366004611088565b6105df565b6002545f906001600160a01b031633146101df5760405162461bcd60e51b81526004016101d6906110a3565b60405180910390fd5b5060015b92915050565b6101f161061c565b806001836040516102029190611119565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b03199092169190911790555050565b60025f9054906101000a90046001600160a01b03166001600160a01b0316636052bf6f5f8787878761028a8f8f8f8b60405160200161027691815260200190565b604051602081830303815290604052610648565b6040518763ffffffff1660e01b81526004016102ab9695949392919061115f565b5f604051808303815f87803b1580156102c2575f80fd5b505af11580156102d4573d5f803e3d5ffd5b505050505050505050505050565b6002545f906001600160a01b0316331461030e5760405162461bcd60e51b81526004016101d6906110a3565b5060019392505050565b61032061061c565b6103295f6106d7565b565b6002545f906001600160a01b031633146103575760405162461bcd60e51b81526004016101d6906110a3565b5f8360a00151806020019051810190610370919061120b565b90505f61037f825f0151610726565b90505f61039486602001518760400151610752565b90505f6103a0836107df565b90506103ac838361084a565b6103bd576103ba828461088b565b90505b5f806103c88361091f565b915091508161042a5760405162461bcd60e51b815260206004820152602860248201527f546f6b656e20726f757465206e6f7420666f756e6420666f722070726566697860448201526765642064656e6f6d60c01b60648201526084016101d6565b8060a001511561046c577f476f2c085b42da8c6a24daf2d3a6d5546fc2cafdd94b1ac39b635fda15cd63a68360405161046391906112e5565b60405180910390a15b8060a00151156104be5760405162461bcd60e51b815260206004820152601a60248201527f6e617469766520746f6b656e206e6f7420737570706f7274656400000000000060448201526064016101d6565b5f81602001519050806001600160a01b031663c1affb1265636f736d6f7360d01b89604001516040516020016104f49190611119565b60408051601f1981840301815291905260608b0151610513905f610a82565b8b602001516105258d60800151610ae6565b6040518663ffffffff1660e01b8152600401610545959493929190611325565b5f604051808303815f87803b15801561055c575f80fd5b505af115801561056e573d5f803e3d5ffd5b5060019d9c50505050505050505050505050565b60405163c13b184f60e01b81526001600160a01b0383169063c13b184f906105ae908490600401611363565b5f604051808303815f87803b1580156105c5575f80fd5b505af11580156105d7573d5f803e3d5ffd5b505050505050565b6105e761061c565b6001600160a01b03811661061057604051631e4fbdf760e01b81525f60048201526024016101d6565b610619816106d7565b50565b5f546001600160a01b031633146103295760405163118cdaa760e01b81523360048201526024016101d6565b60606040518060a001604052808681526020018581526020016106683390565b604051602001610690919060609190911b6bffffffffffffffffffffffff1916815260140190565b6040516020818303038152906040528152602001848152602001838152506040516020016106be9190611375565b6040516020818303038152906040529050949350505050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082019091525f80825260208201526107d86107d361078d604051806040016040528060018152602001602f60f81b815250610726565b6107cd6107d361079c87610726565b6107cd6107d36107c4604051806040016040528060018152602001602f60f81b815250610726565b6107cd8c610726565b9061088b565b610726565b9392505050565b60605f825f01516001600160401b038111156107fd576107fd610be0565b6040519080825280601f01601f191660200182016040528015610827576020820181803683370190505b5090505f602082019050610843818560200151865f0151610b67565b5092915050565b805182515f91111561085d57505f6101e3565b8160200151836020015103610874575060016101e3565b508051602092830151929091015181902091201490565b805182516060915f9161089e9190611410565b6001600160401b038111156108b5576108b5610be0565b6040519080825280601f01601f1916602001820160405280156108df576020820181803683370190505b5090505f6020820190506108fb818660200151875f0151610b67565b84516109179061090b9083611410565b60208601518651610b67565b509392505050565b6040805160e0810182525f808252602082018190529181018290526060808201526080810182905260a0810182905260c081018290526040516381f868f760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906381f868f7906109a1908690600401611363565b5f60405180830381865afa9250505080156109dd57506040513d5f823e601f3d908101601f191682016040526109da9190810190611437565b60015b610a2e5750506040805160e0810182525f8082526020808301829052828401829052835190810190935280835260608201929092526080810182905260a0810182905260c081018290529092909150565b6040805160e0810182526001600160a01b03988916815296881660208801529490961693850193909352606084019190915260ff166080830152151560a082015290151560c0820152600194909350915050565b5f610a8e826014611410565b83511015610ad65760405162461bcd60e51b8152602060048201526015602482015274746f416464726573735f6f75744f66426f756e647360581b60448201526064016101d6565b500160200151600160601b900490565b5f81515f03610af657505f919050565b602082511115610b5f5760405162461bcd60e51b815260206004820152602e60248201527f6279746573546f427974657333323a20736f75726365206c656e67746820657860448201526d636565647320333220627974657360901b60648201526084016101d6565b506020015190565b60208110610b9f5781518352610b7e602084611410565b9250610b8b602083611410565b9150610b986020826114e1565b9050610b67565b5f198115610bcd576001610bb48360206114e1565b610bc0906101006115d4565b610bca91906114e1565b90505b9151835183169219169190911790915250565b634e487b7160e01b5f52604160045260245ffd5b60405161010081016001600160401b0381118282101715610c1757610c17610be0565b60405290565b60405160a081016001600160401b0381118282101715610c1757610c17610be0565b604051601f8201601f191681016001600160401b0381118282101715610c6757610c67610be0565b604052919050565b5f6001600160401b03821115610c8757610c87610be0565b50601f01601f191660200190565b5f82601f830112610ca4575f80fd5b8135610cb7610cb282610c6f565b610c3f565b818152846020838601011115610ccb575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60408284031215610cf7575f80fd5b604051604081018181106001600160401b0382111715610d1957610d19610be0565b604052823581526020928301359281019290925250919050565b5f6101208284031215610d44575f80fd5b610d4c610bf4565b90508135815260208201356001600160401b0380821115610d6b575f80fd5b610d7785838601610c95565b60208401526040840135915080821115610d8f575f80fd5b610d9b85838601610c95565b60408401526060840135915080821115610db3575f80fd5b610dbf85838601610c95565b60608401526080840135915080821115610dd7575f80fd5b610de385838601610c95565b608084015260a0840135915080821115610dfb575f80fd5b50610e0884828501610c95565b60a083015250610e1b8360c08401610ce7565b60c082015261010082013560e082015292915050565b5f8060408385031215610e42575f80fd5b82356001600160401b0380821115610e58575f80fd5b610e6486838701610d33565b93506020850135915080821115610e79575f80fd5b50610e8685828601610c95565b9150509250929050565b6001600160a01b0381168114610619575f80fd5b5f8060408385031215610eb5575f80fd5b82356001600160401b03811115610eca575f80fd5b610ed685828601610c95565b9250506020830135610ee781610e90565b809150509250929050565b5f805f805f805f80610120898b031215610f0a575f80fd5b88356001600160401b0380821115610f20575f80fd5b610f2c8c838d01610c95565b995060208b0135985060408b0135915080821115610f48575f80fd5b610f548c838d01610c95565b975060608b0135915080821115610f69575f80fd5b610f758c838d01610c95565b965060808b0135915080821115610f8a575f80fd5b50610f978b828c01610c95565b945050610fa78a60a08b01610ce7565b925060e0890135915061010089013590509295985092959890939650565b5f805f60608486031215610fd7575f80fd5b83356001600160401b0380821115610fed575f80fd5b610ff987838801610d33565b9450602086013591508082111561100e575f80fd5b61101a87838801610c95565b9350604086013591508082111561102f575f80fd5b5061103c86828701610c95565b9150509250925092565b5f8060408385031215611057575f80fd5b823561106281610e90565b915060208301356001600160401b0381111561107c575f80fd5b610e8685828601610c95565b5f60208284031215611098575f80fd5b81356107d881610e90565b60208082526034908201527f4261736546756e6769626c65546f6b656e4170703a2063616c6c6572206973206040820152731b9bdd081d1a1948125090c818dbdb9d1c9858dd60621b606082015260800190565b5f5b838110156111115781810151838201526020016110f9565b50505f910152565b5f825161112a8184602087016110f7565b9190910192915050565b5f815180845261114b8160208601602086016110f7565b601f01601f19169290920160200192915050565b86815260e060208201525f61117760e0830188611134565b82810360408401526111898188611134565b905085516060840152602086015160808401528460a084015282810360c08401526111b48185611134565b9998505050505050505050565b5f82601f8301126111d0575f80fd5b81516111de610cb282610c6f565b8181528460208386010111156111f2575f80fd5b6112038260208301602087016110f7565b949350505050565b5f6020828403121561121b575f80fd5b81516001600160401b0380821115611231575f80fd5b9083019060a08286031215611244575f80fd5b61124c610c1d565b82518281111561125a575f80fd5b611266878286016111c1565b82525060208301516020820152604083015182811115611284575f80fd5b611290878286016111c1565b6040830152506060830151828111156112a7575f80fd5b6112b3878286016111c1565b6060830152506080830151828111156112ca575f80fd5b6112d6878286016111c1565b60808301525095945050505050565b604081525f6112f76040830184611134565b8281036020840152600c81526b3730ba34bb32903a37b5b2b760a11b60208201526040810191505092915050565b85815260a060208201525f61133d60a0830187611134565b6001600160a01b0395909516604083015250606081019290925260809091015292915050565b602081525f6107d86020830184611134565b602081525f825160a0602084015261139060c0840182611134565b9050602084015160408401526040840151601f19808584030160608601526113b88383611134565b925060608601519150808584030160808601526113d58383611134565b925060808601519150808584030160a0860152506113f38282611134565b95945050505050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156101e3576101e36113fc565b80518015158114611432575f80fd5b919050565b5f805f805f805f60e0888a03121561144d575f80fd5b875161145881610e90565b602089015190975061146981610e90565b604089015190965061147a81610e90565b60608901519095506001600160401b03811115611495575f80fd5b6114a18a828b016111c1565b945050608088015160ff811681146114b7575f80fd5b92506114c560a08901611423565b91506114d360c08901611423565b905092959891949750929550565b818103818111156101e3576101e36113fc565b600181815b8085111561152e57815f1904821115611514576115146113fc565b8085161561152157918102915b93841c93908002906114f9565b509250929050565b5f82611544575060016101e3565b8161155057505f6101e3565b816001811461156657600281146115705761158c565b60019150506101e3565b60ff841115611581576115816113fc565b50506001821b6101e3565b5060208310610133831016604e8410600b84101617156115af575081810a6101e3565b6115b983836114f4565b805f19048211156115cc576115cc6113fc565b029392505050565b5f6107d8838361153656fea26469706673582212205319068d66aef68c357d452415f32ac157c51e2d8ee9f136f9cc431f5a9e2f0c64736f6c63430008190033",
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
// Solidity: function onAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) , bytes , bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, arg0 Packet, arg1 []byte, arg2 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.contract.Transact(opts, "onAcknowledgementPacket", arg0, arg1, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x63c1bec5.
//
// Solidity: function onAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) , bytes , bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) OnAcknowledgementPacket(arg0 Packet, arg1 []byte, arg2 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.OnAcknowledgementPacket(&_ICS20BankTransferApp.TransactOpts, arg0, arg1, arg2)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x63c1bec5.
//
// Solidity: function onAcknowledgementPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) , bytes , bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorSession) OnAcknowledgementPacket(arg0 Packet, arg1 []byte, arg2 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.OnAcknowledgementPacket(&_ICS20BankTransferApp.TransactOpts, arg0, arg1, arg2)
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
// Solidity: function onTimeoutPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) , bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactor) OnTimeoutPacket(opts *bind.TransactOpts, arg0 Packet, arg1 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.contract.Transact(opts, "onTimeoutPacket", arg0, arg1)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x3df7d406.
//
// Solidity: function onTimeoutPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) , bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppSession) OnTimeoutPacket(arg0 Packet, arg1 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.OnTimeoutPacket(&_ICS20BankTransferApp.TransactOpts, arg0, arg1)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x3df7d406.
//
// Solidity: function onTimeoutPacket((uint256,string,string,string,string,bytes,(uint256,uint256),uint256) , bytes ) returns(bool)
func (_ICS20BankTransferApp *ICS20BankTransferAppTransactorSession) OnTimeoutPacket(arg0 Packet, arg1 []byte) (*types.Transaction, error) {
	return _ICS20BankTransferApp.Contract.OnTimeoutPacket(&_ICS20BankTransferApp.TransactOpts, arg0, arg1)
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
