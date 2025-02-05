// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenrouter

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

// TokenRouterMetaData contains all meta data concerning the TokenRouter contract.
var TokenRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"TokenConfigRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"remote\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"home\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExternal\",\"type\":\"bool\"}],\"name\":\"TokenConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"remote\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"home\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"TokenConfigValidationFailed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"denoms\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"remotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"homes\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"channels\",\"type\":\"string[]\"},{\"internalType\":\"uint8[]\",\"name\":\"decimalsArray\",\"type\":\"uint8[]\"},{\"internalType\":\"bool[]\",\"name\":\"isNativeArray\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"isExternalArray\",\"type\":\"bool[]\"}],\"name\":\"batchSetTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"sourceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"targetDecimals\",\"type\":\"uint8\"}],\"name\":\"convertDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"denomToConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"home\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isExternal\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getTokenChannel\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"channels\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getTokenConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"home\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remote\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isExternal\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"isTokenSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"removeTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"home\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isExternal\",\"type\":\"bool\"}],\"name\":\"setTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051611ad5380380611ad583398101604081905261002f916100be565b806001600160a01b03811661005e57604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b6100678161006e565b50506100ee565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100d057600080fd5b81516001600160a01b03811681146100e757600080fd5b9392505050565b6119d8806100fd6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806381f868f71161007157806381f868f71461015f5780638da5cb5b14610172578063b03171c21461018d578063b450e5d8146101ad578063c8a999df146101c0578063f2fde38b146101d357600080fd5b80630417cf8e146100b95780630d3652ec146100d8578063605ba00d146100fb5780636ae6490514610121578063715018a61461014257806374a707421461014c575b600080fd5b6100c1601281565b60405160ff90911681526020015b60405180910390f35b6100eb6100e63660046110e2565b6101e6565b60405190151581526020016100cf565b61010e6101093660046110e2565b610331565b6040516100cf979695949392919061116f565b61013461012f3660046111dd565b61041d565b6040519081526020016100cf565b61014a610495565b005b61014a61015a366004611240565b6104a9565b61010e61016d3660046110e2565b610669565b6000546040516001600160a01b0390911681526020016100cf565b6101a061019b3660046110e2565b610833565b6040516100cf9190611308565b61014a6101bb366004611501565b6109c7565b61014a6101ce3660046110e2565b610c9c565b61014a6101e1366004611641565b610dcf565b6000806001836040516101f9919061165c565b90815260408051918290036020908101832060e08401835280546001600160a01b039081168552600182015481169285019290925260028101549091169183019190915260038101805460608401919061025290611678565b80601f016020809104026020016040519081016040528092919081815260200182805461027e90611678565b80156102cb5780601f106102a0576101008083540402835291602001916102cb565b820191906000526020600020905b8154815290600101906020018083116102ae57829003601f168201915b50505091835250506004919091015460ff808216602084015261010082048116151560408401526201000090910416151560609091015280519091506001600160a01b03161580159061032a575060408101516001600160a01b031615155b9392505050565b80516020818301810180516001808352938301929094019190912092905281549082015460028301546003840180546001600160a01b039485169593851694909216929161037e90611678565b80601f01602080910402602001604051908101604052809291908181526020018280546103aa90611678565b80156103f75780601f106103cc576101008083540402835291602001916103f7565b820191906000526020600020905b8154815290600101906020018083116103da57829003601f168201915b5050506004909301549192505060ff808216916101008104821691620100009091041687565b60008160ff168360ff160361043357508261032a565b8160ff168360ff16111561046f57600061044d83856116c8565b60ff16905061045d81600a6117cb565b61046790866117d7565b91505061032a565b600061047b84846116c8565b60ff16905061048b81600a6117cb565b61046790866117f9565b61049d610e0d565b6104a76000610e3a565b565b6104b1610e0d565b6000806104c089888a88610e8a565b915091508181906104ed5760405162461bcd60e51b81526004016104e49190611308565b60405180910390fd5b506040518060e001604052808a6001600160a01b03168152602001896001600160a01b03168152602001886001600160a01b031681526020018781526020018660ff168152602001851515815260200184151581525060018b604051610553919061165c565b9081526040805160209281900383019020835181546001600160a01b03199081166001600160a01b0392831617835593850151600183018054861691831691909117905591840151600282018054909416921691909117909155606082015160038201906105c19082611861565b5060808201516004909101805460a084015160c0909401511515620100000262ff0000199415156101000261ffff1990921660ff909416939093171792909216179055604051610612908b9061165c565b60405180910390207f8da271fee9489a2d04f4b19df1eb40abce1b8139522702443aa73c5392d66c988a8a8a8a8a8a8a604051610655979695949392919061116f565b60405180910390a250505050505050505050565b60008060006060600080600080600189604051610686919061165c565b90815260408051918290036020908101832060e08401835280546001600160a01b03908116855260018201548116928501929092526002810154909116918301919091526003810180546060840191906106df90611678565b80601f016020809104026020016040519081016040528092919081815260200182805461070b90611678565b80156107585780601f1061072d57610100808354040283529160200191610758565b820191906000526020600020905b81548152906001019060200180831161073b57829003601f168201915b50505091835250506004919091015460ff80821660208085019190915261010083048216151560408086019190915262010000909304909116151560609093019290925282518151808301909252601d82527f546f6b656e526f757465723a20436f6e666967206e6f7420666f756e64000000928201929092529192506001600160a01b03166107fb5760405162461bcd60e51b81526004016104e49190611308565b508051602082015160408301516060840151608085015160a086015160c090960151949e939d50919b50995097509195509350915050565b60606000600183604051610847919061165c565b90815260408051918290036020908101832060e08401835280546001600160a01b03908116855260018201548116928501929092526002810154909116918301919091526003810180546060840191906108a090611678565b80601f01602080910402602001604051908101604052809291908181526020018280546108cc90611678565b80156109195780601f106108ee57610100808354040283529160200191610919565b820191906000526020600020905b8154815290600101906020018083116108fc57829003601f168201915b50505091835250506004919091015460ff80821660208085019190915261010083048216151560408086019190915262010000909304909116151560609093019290925282518151808301909252601d82527f546f6b656e526f757465723a20436f6e666967206e6f7420666f756e64000000928201929092529192506001600160a01b03166109bc5760405162461bcd60e51b81526004016104e49190611308565b506060015192915050565b6109cf610e0d565b865188511480156109e1575084518851145b80156109ee575085518851145b80156109fb575082518851145b8015610a08575081518851145b8015610a15575080518851145b6040518060600160405280602381526020016119806023913990610a4c5760405162461bcd60e51b81526004016104e49190611308565b5060005b8851811015610c9157600080610acc8a8481518110610a7157610a71611921565b6020026020010151898581518110610a8b57610a8b611921565b60200260200101518b8681518110610aa557610aa5611921565b6020026020010151898781518110610abf57610abf611921565b6020026020010151610e8a565b9150915081610bae578a8381518110610ae757610ae7611921565b6020026020010151604051610afc919061165c565b60405180910390207f81403746f5f147d73940fc611424ecd32b37b9ed0687cd4ef68b945f4fedcb61828c8681518110610b3857610b38611921565b60200260200101518c8781518110610b5257610b52611921565b60200260200101518c8881518110610b6c57610b6c611921565b60200260200101518b8981518110610b8657610b86611921565b6020026020010151604051610b9f959493929190611937565b60405180910390a25050610c89565b610c868b8481518110610bc357610bc3611921565b60200260200101518b8581518110610bdd57610bdd611921565b60200260200101518b8681518110610bf757610bf7611921565b60200260200101518b8781518110610c1157610c11611921565b60200260200101518b8881518110610c2b57610c2b611921565b60200260200101518b8981518110610c4557610c45611921565b60200260200101518b8a81518110610c5f57610c5f611921565b60200260200101518b8b81518110610c7957610c79611921565b60200260200101516104a9565b50505b600101610a50565b505050505050505050565b610ca4610e0d565b60006001600160a01b0316600182604051610cbf919061165c565b90815260408051918290036020908101832054838301909252601d83527f546f6b656e526f757465723a20436f6e666967206e6f7420666f756e640000009083015290916001600160a01b0390911603610d2c5760405162461bcd60e51b81526004016104e49190611308565b50600181604051610d3d919061165c565b90815260405190819003602001902080546001600160a01b031990811682556001820180548216905560028201805490911690556000610d806003830182610fdd565b50600401805462ffffff19169055604051610d9c90829061165c565b604051908190038120907f7b03ec0d67d50c23129a8d5eddce82aab209af695b966d4bf7c9b4ef0927477090600090a250565b610dd7610e0d565b6001600160a01b038116610e0157604051631e4fbdf760e01b8152600060048201526024016104e4565b610e0a81610e3a565b50565b6000546001600160a01b031633146104a75760405163118cdaa760e01b81523360048201526024016104e4565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060606001600160a01b038616610eda57505060408051808201909152601f81527f546f6b656e526f757465723a205a65726f20746f6b656e2061646472657373006020820152600090610fd4565b6001600160a01b038516610f2657505060408051808201909152601e81527f546f6b656e526f757465723a205a65726f20686f6d65206164647265737300006020820152600090610fd4565b6001600160a01b038416610f755760006040518060400160405280602081526020017f546f6b656e526f757465723a205a65726f2072656d6f7465206164647265737381525091509150610fd4565b601260ff84161115610fbf57505060408051808201909152601d81527f546f6b656e526f757465723a20496e76616c696420646563696d616c730000006020820152600090610fd4565b50506040805160208101909152600081526001905b94509492505050565b508054610fe990611678565b6000825580601f10610ff9575050565b601f016020900490600052602060002090810190610e0a91905b808211156110275760008155600101611013565b5090565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561106a5761106a61102b565b604052919050565b600082601f83011261108357600080fd5b813567ffffffffffffffff81111561109d5761109d61102b565b6110b0601f8201601f1916602001611041565b8181528460208386010111156110c557600080fd5b816020850160208301376000918101602001919091529392505050565b6000602082840312156110f457600080fd5b813567ffffffffffffffff81111561110b57600080fd5b61111784828501611072565b949350505050565b60005b8381101561113a578181015183820152602001611122565b50506000910152565b6000815180845261115b81602086016020860161111f565b601f01601f19169290920160200192915050565b6001600160a01b03888116825287811660208301528616604082015260e0606082018190526000906111a390830187611143565b60ff9590951660808301525091151560a0830152151560c090910152949350505050565b803560ff811681146111d857600080fd5b919050565b6000806000606084860312156111f257600080fd5b83359250611202602085016111c7565b9150611210604085016111c7565b90509250925092565b80356001600160a01b03811681146111d857600080fd5b803580151581146111d857600080fd5b600080600080600080600080610100898b03121561125d57600080fd5b883567ffffffffffffffff8082111561127557600080fd5b6112818c838d01611072565b995061128f60208c01611219565b985061129d60408c01611219565b97506112ab60608c01611219565b965060808b01359150808211156112c157600080fd5b506112ce8b828c01611072565b9450506112dd60a08a016111c7565b92506112eb60c08a01611230565b91506112f960e08a01611230565b90509295985092959890939650565b60208152600061032a6020830184611143565b600067ffffffffffffffff8211156113355761133561102b565b5060051b60200190565b600082601f83011261135057600080fd5b813560206113656113608361131b565b611041565b82815260059290921b8401810191818101908684111561138457600080fd5b8286015b848110156113c457803567ffffffffffffffff8111156113a85760008081fd5b6113b68986838b0101611072565b845250918301918301611388565b509695505050505050565b600082601f8301126113e057600080fd5b813560206113f06113608361131b565b8083825260208201915060208460051b87010193508684111561141257600080fd5b602086015b848110156113c45761142881611219565b8352918301918301611417565b600082601f83011261144657600080fd5b813560206114566113608361131b565b8083825260208201915060208460051b87010193508684111561147857600080fd5b602086015b848110156113c45761148e816111c7565b835291830191830161147d565b600082601f8301126114ac57600080fd5b813560206114bc6113608361131b565b8083825260208201915060208460051b8701019350868411156114de57600080fd5b602086015b848110156113c4576114f481611230565b83529183019183016114e3565b600080600080600080600080610100898b03121561151e57600080fd5b883567ffffffffffffffff8082111561153657600080fd5b6115428c838d0161133f565b995060208b013591508082111561155857600080fd5b6115648c838d016113cf565b985060408b013591508082111561157a57600080fd5b6115868c838d016113cf565b975060608b013591508082111561159c57600080fd5b6115a88c838d016113cf565b965060808b01359150808211156115be57600080fd5b6115ca8c838d0161133f565b955060a08b01359150808211156115e057600080fd5b6115ec8c838d01611435565b945060c08b013591508082111561160257600080fd5b61160e8c838d0161149b565b935060e08b013591508082111561162457600080fd5b506116318b828c0161149b565b9150509295985092959890939650565b60006020828403121561165357600080fd5b61032a82611219565b6000825161166e81846020870161111f565b9190910192915050565b600181811c9082168061168c57607f821691505b6020821081036116ac57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b60ff82811682821603908111156116e1576116e16116b2565b92915050565b600181815b80851115611722578160001904821115611708576117086116b2565b8085161561171557918102915b93841c93908002906116ec565b509250929050565b600082611739575060016116e1565b81611746575060006116e1565b816001811461175c576002811461176657611782565b60019150506116e1565b60ff841115611777576117776116b2565b50506001821b6116e1565b5060208310610133831016604e8410600b84101617156117a5575081810a6116e1565b6117af83836116e7565b80600019048211156117c3576117c36116b2565b029392505050565b600061032a838361172a565b6000826117f457634e487b7160e01b600052601260045260246000fd5b500490565b80820281158282048414176116e1576116e16116b2565b601f82111561185c576000816000526020600020601f850160051c810160208610156118395750805b601f850160051c820191505b8181101561185857828155600101611845565b5050505b505050565b815167ffffffffffffffff81111561187b5761187b61102b565b61188f816118898454611678565b84611810565b602080601f8311600181146118c457600084156118ac5750858301515b600019600386901b1c1916600185901b178555611858565b600085815260208120601f198616915b828110156118f3578886015182559484019460019091019084016118d4565b50858210156119115787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b60a08152600061194a60a0830188611143565b6001600160a01b03968716602084015294861660408301525091909316606082015260ff90921660809092019190915291905056fe546f6b656e526f757465723a204172726179206c656e67746873206d69736d61746368a264697066735822122094a13d61e5e099321b30b8e83d113f21557799e2e53c276e1fd2623e6b53e87a64736f6c63430008190033",
}

// TokenRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenRouterMetaData.ABI instead.
var TokenRouterABI = TokenRouterMetaData.ABI

// TokenRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenRouterMetaData.Bin instead.
var TokenRouterBin = TokenRouterMetaData.Bin

// DeployTokenRouter deploys a new Ethereum contract, binding an instance of TokenRouter to it.
func DeployTokenRouter(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address) (common.Address, *types.Transaction, *TokenRouter, error) {
	parsed, err := TokenRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenRouterBin), backend, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenRouter{TokenRouterCaller: TokenRouterCaller{contract: contract}, TokenRouterTransactor: TokenRouterTransactor{contract: contract}, TokenRouterFilterer: TokenRouterFilterer{contract: contract}}, nil
}

// TokenRouter is an auto generated Go binding around an Ethereum contract.
type TokenRouter struct {
	TokenRouterCaller     // Read-only binding to the contract
	TokenRouterTransactor // Write-only binding to the contract
	TokenRouterFilterer   // Log filterer for contract events
}

// TokenRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenRouterSession struct {
	Contract     *TokenRouter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenRouterCallerSession struct {
	Contract *TokenRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TokenRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenRouterTransactorSession struct {
	Contract     *TokenRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TokenRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRouterRaw struct {
	Contract *TokenRouter // Generic contract binding to access the raw methods on
}

// TokenRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenRouterCallerRaw struct {
	Contract *TokenRouterCaller // Generic read-only contract binding to access the raw methods on
}

// TokenRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenRouterTransactorRaw struct {
	Contract *TokenRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenRouter creates a new instance of TokenRouter, bound to a specific deployed contract.
func NewTokenRouter(address common.Address, backend bind.ContractBackend) (*TokenRouter, error) {
	contract, err := bindTokenRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenRouter{TokenRouterCaller: TokenRouterCaller{contract: contract}, TokenRouterTransactor: TokenRouterTransactor{contract: contract}, TokenRouterFilterer: TokenRouterFilterer{contract: contract}}, nil
}

// NewTokenRouterCaller creates a new read-only instance of TokenRouter, bound to a specific deployed contract.
func NewTokenRouterCaller(address common.Address, caller bind.ContractCaller) (*TokenRouterCaller, error) {
	contract, err := bindTokenRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRouterCaller{contract: contract}, nil
}

// NewTokenRouterTransactor creates a new write-only instance of TokenRouter, bound to a specific deployed contract.
func NewTokenRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenRouterTransactor, error) {
	contract, err := bindTokenRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRouterTransactor{contract: contract}, nil
}

// NewTokenRouterFilterer creates a new log filterer instance of TokenRouter, bound to a specific deployed contract.
func NewTokenRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenRouterFilterer, error) {
	contract, err := bindTokenRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenRouterFilterer{contract: contract}, nil
}

// bindTokenRouter binds a generic wrapper to an already deployed contract.
func bindTokenRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRouter *TokenRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRouter.Contract.TokenRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRouter *TokenRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRouter.Contract.TokenRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRouter *TokenRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRouter.Contract.TokenRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRouter *TokenRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRouter *TokenRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRouter *TokenRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRouter.Contract.contract.Transact(opts, method, params...)
}

// MAXDECIMALS is a free data retrieval call binding the contract method 0x0417cf8e.
//
// Solidity: function MAX_DECIMALS() view returns(uint8)
func (_TokenRouter *TokenRouterCaller) MAXDECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "MAX_DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXDECIMALS is a free data retrieval call binding the contract method 0x0417cf8e.
//
// Solidity: function MAX_DECIMALS() view returns(uint8)
func (_TokenRouter *TokenRouterSession) MAXDECIMALS() (uint8, error) {
	return _TokenRouter.Contract.MAXDECIMALS(&_TokenRouter.CallOpts)
}

// MAXDECIMALS is a free data retrieval call binding the contract method 0x0417cf8e.
//
// Solidity: function MAX_DECIMALS() view returns(uint8)
func (_TokenRouter *TokenRouterCallerSession) MAXDECIMALS() (uint8, error) {
	return _TokenRouter.Contract.MAXDECIMALS(&_TokenRouter.CallOpts)
}

// ConvertDecimals is a free data retrieval call binding the contract method 0x6ae64905.
//
// Solidity: function convertDecimals(uint256 amount, uint8 sourceDecimals, uint8 targetDecimals) pure returns(uint256)
func (_TokenRouter *TokenRouterCaller) ConvertDecimals(opts *bind.CallOpts, amount *big.Int, sourceDecimals uint8, targetDecimals uint8) (*big.Int, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "convertDecimals", amount, sourceDecimals, targetDecimals)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertDecimals is a free data retrieval call binding the contract method 0x6ae64905.
//
// Solidity: function convertDecimals(uint256 amount, uint8 sourceDecimals, uint8 targetDecimals) pure returns(uint256)
func (_TokenRouter *TokenRouterSession) ConvertDecimals(amount *big.Int, sourceDecimals uint8, targetDecimals uint8) (*big.Int, error) {
	return _TokenRouter.Contract.ConvertDecimals(&_TokenRouter.CallOpts, amount, sourceDecimals, targetDecimals)
}

// ConvertDecimals is a free data retrieval call binding the contract method 0x6ae64905.
//
// Solidity: function convertDecimals(uint256 amount, uint8 sourceDecimals, uint8 targetDecimals) pure returns(uint256)
func (_TokenRouter *TokenRouterCallerSession) ConvertDecimals(amount *big.Int, sourceDecimals uint8, targetDecimals uint8) (*big.Int, error) {
	return _TokenRouter.Contract.ConvertDecimals(&_TokenRouter.CallOpts, amount, sourceDecimals, targetDecimals)
}

// DenomToConfig is a free data retrieval call binding the contract method 0x605ba00d.
//
// Solidity: function denomToConfig(string ) view returns(address token, address remote, address home, string channel, uint8 decimals, bool isNative, bool isExternal)
func (_TokenRouter *TokenRouterCaller) DenomToConfig(opts *bind.CallOpts, arg0 string) (struct {
	Token      common.Address
	Remote     common.Address
	Home       common.Address
	Channel    string
	Decimals   uint8
	IsNative   bool
	IsExternal bool
}, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "denomToConfig", arg0)

	outstruct := new(struct {
		Token      common.Address
		Remote     common.Address
		Home       common.Address
		Channel    string
		Decimals   uint8
		IsNative   bool
		IsExternal bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Remote = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Home = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Channel = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Decimals = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.IsNative = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.IsExternal = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// DenomToConfig is a free data retrieval call binding the contract method 0x605ba00d.
//
// Solidity: function denomToConfig(string ) view returns(address token, address remote, address home, string channel, uint8 decimals, bool isNative, bool isExternal)
func (_TokenRouter *TokenRouterSession) DenomToConfig(arg0 string) (struct {
	Token      common.Address
	Remote     common.Address
	Home       common.Address
	Channel    string
	Decimals   uint8
	IsNative   bool
	IsExternal bool
}, error) {
	return _TokenRouter.Contract.DenomToConfig(&_TokenRouter.CallOpts, arg0)
}

// DenomToConfig is a free data retrieval call binding the contract method 0x605ba00d.
//
// Solidity: function denomToConfig(string ) view returns(address token, address remote, address home, string channel, uint8 decimals, bool isNative, bool isExternal)
func (_TokenRouter *TokenRouterCallerSession) DenomToConfig(arg0 string) (struct {
	Token      common.Address
	Remote     common.Address
	Home       common.Address
	Channel    string
	Decimals   uint8
	IsNative   bool
	IsExternal bool
}, error) {
	return _TokenRouter.Contract.DenomToConfig(&_TokenRouter.CallOpts, arg0)
}

// GetTokenChannel is a free data retrieval call binding the contract method 0xb03171c2.
//
// Solidity: function getTokenChannel(string denom) view returns(string channels)
func (_TokenRouter *TokenRouterCaller) GetTokenChannel(opts *bind.CallOpts, denom string) (string, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "getTokenChannel", denom)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenChannel is a free data retrieval call binding the contract method 0xb03171c2.
//
// Solidity: function getTokenChannel(string denom) view returns(string channels)
func (_TokenRouter *TokenRouterSession) GetTokenChannel(denom string) (string, error) {
	return _TokenRouter.Contract.GetTokenChannel(&_TokenRouter.CallOpts, denom)
}

// GetTokenChannel is a free data retrieval call binding the contract method 0xb03171c2.
//
// Solidity: function getTokenChannel(string denom) view returns(string channels)
func (_TokenRouter *TokenRouterCallerSession) GetTokenChannel(denom string) (string, error) {
	return _TokenRouter.Contract.GetTokenChannel(&_TokenRouter.CallOpts, denom)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x81f868f7.
//
// Solidity: function getTokenConfig(string denom) view returns(address token, address home, address remote, string channel, uint8 decimals, bool isNative, bool isExternal)
func (_TokenRouter *TokenRouterCaller) GetTokenConfig(opts *bind.CallOpts, denom string) (struct {
	Token      common.Address
	Home       common.Address
	Remote     common.Address
	Channel    string
	Decimals   uint8
	IsNative   bool
	IsExternal bool
}, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "getTokenConfig", denom)

	outstruct := new(struct {
		Token      common.Address
		Home       common.Address
		Remote     common.Address
		Channel    string
		Decimals   uint8
		IsNative   bool
		IsExternal bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Home = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Remote = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Channel = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Decimals = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.IsNative = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.IsExternal = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// GetTokenConfig is a free data retrieval call binding the contract method 0x81f868f7.
//
// Solidity: function getTokenConfig(string denom) view returns(address token, address home, address remote, string channel, uint8 decimals, bool isNative, bool isExternal)
func (_TokenRouter *TokenRouterSession) GetTokenConfig(denom string) (struct {
	Token      common.Address
	Home       common.Address
	Remote     common.Address
	Channel    string
	Decimals   uint8
	IsNative   bool
	IsExternal bool
}, error) {
	return _TokenRouter.Contract.GetTokenConfig(&_TokenRouter.CallOpts, denom)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x81f868f7.
//
// Solidity: function getTokenConfig(string denom) view returns(address token, address home, address remote, string channel, uint8 decimals, bool isNative, bool isExternal)
func (_TokenRouter *TokenRouterCallerSession) GetTokenConfig(denom string) (struct {
	Token      common.Address
	Home       common.Address
	Remote     common.Address
	Channel    string
	Decimals   uint8
	IsNative   bool
	IsExternal bool
}, error) {
	return _TokenRouter.Contract.GetTokenConfig(&_TokenRouter.CallOpts, denom)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x0d3652ec.
//
// Solidity: function isTokenSupported(string denom) view returns(bool)
func (_TokenRouter *TokenRouterCaller) IsTokenSupported(opts *bind.CallOpts, denom string) (bool, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "isTokenSupported", denom)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenSupported is a free data retrieval call binding the contract method 0x0d3652ec.
//
// Solidity: function isTokenSupported(string denom) view returns(bool)
func (_TokenRouter *TokenRouterSession) IsTokenSupported(denom string) (bool, error) {
	return _TokenRouter.Contract.IsTokenSupported(&_TokenRouter.CallOpts, denom)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x0d3652ec.
//
// Solidity: function isTokenSupported(string denom) view returns(bool)
func (_TokenRouter *TokenRouterCallerSession) IsTokenSupported(denom string) (bool, error) {
	return _TokenRouter.Contract.IsTokenSupported(&_TokenRouter.CallOpts, denom)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenRouter *TokenRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenRouter *TokenRouterSession) Owner() (common.Address, error) {
	return _TokenRouter.Contract.Owner(&_TokenRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenRouter *TokenRouterCallerSession) Owner() (common.Address, error) {
	return _TokenRouter.Contract.Owner(&_TokenRouter.CallOpts)
}

// BatchSetTokenConfig is a paid mutator transaction binding the contract method 0xb450e5d8.
//
// Solidity: function batchSetTokenConfig(string[] denoms, address[] tokens, address[] remotes, address[] homes, string[] channels, uint8[] decimalsArray, bool[] isNativeArray, bool[] isExternalArray) returns()
func (_TokenRouter *TokenRouterTransactor) BatchSetTokenConfig(opts *bind.TransactOpts, denoms []string, tokens []common.Address, remotes []common.Address, homes []common.Address, channels []string, decimalsArray []uint8, isNativeArray []bool, isExternalArray []bool) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "batchSetTokenConfig", denoms, tokens, remotes, homes, channels, decimalsArray, isNativeArray, isExternalArray)
}

// BatchSetTokenConfig is a paid mutator transaction binding the contract method 0xb450e5d8.
//
// Solidity: function batchSetTokenConfig(string[] denoms, address[] tokens, address[] remotes, address[] homes, string[] channels, uint8[] decimalsArray, bool[] isNativeArray, bool[] isExternalArray) returns()
func (_TokenRouter *TokenRouterSession) BatchSetTokenConfig(denoms []string, tokens []common.Address, remotes []common.Address, homes []common.Address, channels []string, decimalsArray []uint8, isNativeArray []bool, isExternalArray []bool) (*types.Transaction, error) {
	return _TokenRouter.Contract.BatchSetTokenConfig(&_TokenRouter.TransactOpts, denoms, tokens, remotes, homes, channels, decimalsArray, isNativeArray, isExternalArray)
}

// BatchSetTokenConfig is a paid mutator transaction binding the contract method 0xb450e5d8.
//
// Solidity: function batchSetTokenConfig(string[] denoms, address[] tokens, address[] remotes, address[] homes, string[] channels, uint8[] decimalsArray, bool[] isNativeArray, bool[] isExternalArray) returns()
func (_TokenRouter *TokenRouterTransactorSession) BatchSetTokenConfig(denoms []string, tokens []common.Address, remotes []common.Address, homes []common.Address, channels []string, decimalsArray []uint8, isNativeArray []bool, isExternalArray []bool) (*types.Transaction, error) {
	return _TokenRouter.Contract.BatchSetTokenConfig(&_TokenRouter.TransactOpts, denoms, tokens, remotes, homes, channels, decimalsArray, isNativeArray, isExternalArray)
}

// RemoveTokenConfig is a paid mutator transaction binding the contract method 0xc8a999df.
//
// Solidity: function removeTokenConfig(string denom) returns()
func (_TokenRouter *TokenRouterTransactor) RemoveTokenConfig(opts *bind.TransactOpts, denom string) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "removeTokenConfig", denom)
}

// RemoveTokenConfig is a paid mutator transaction binding the contract method 0xc8a999df.
//
// Solidity: function removeTokenConfig(string denom) returns()
func (_TokenRouter *TokenRouterSession) RemoveTokenConfig(denom string) (*types.Transaction, error) {
	return _TokenRouter.Contract.RemoveTokenConfig(&_TokenRouter.TransactOpts, denom)
}

// RemoveTokenConfig is a paid mutator transaction binding the contract method 0xc8a999df.
//
// Solidity: function removeTokenConfig(string denom) returns()
func (_TokenRouter *TokenRouterTransactorSession) RemoveTokenConfig(denom string) (*types.Transaction, error) {
	return _TokenRouter.Contract.RemoveTokenConfig(&_TokenRouter.TransactOpts, denom)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenRouter *TokenRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenRouter *TokenRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenRouter.Contract.RenounceOwnership(&_TokenRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenRouter *TokenRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenRouter.Contract.RenounceOwnership(&_TokenRouter.TransactOpts)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x74a70742.
//
// Solidity: function setTokenConfig(string denom, address token, address remote, address home, string channel, uint8 decimals, bool isNative, bool isExternal) returns()
func (_TokenRouter *TokenRouterTransactor) SetTokenConfig(opts *bind.TransactOpts, denom string, token common.Address, remote common.Address, home common.Address, channel string, decimals uint8, isNative bool, isExternal bool) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "setTokenConfig", denom, token, remote, home, channel, decimals, isNative, isExternal)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x74a70742.
//
// Solidity: function setTokenConfig(string denom, address token, address remote, address home, string channel, uint8 decimals, bool isNative, bool isExternal) returns()
func (_TokenRouter *TokenRouterSession) SetTokenConfig(denom string, token common.Address, remote common.Address, home common.Address, channel string, decimals uint8, isNative bool, isExternal bool) (*types.Transaction, error) {
	return _TokenRouter.Contract.SetTokenConfig(&_TokenRouter.TransactOpts, denom, token, remote, home, channel, decimals, isNative, isExternal)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x74a70742.
//
// Solidity: function setTokenConfig(string denom, address token, address remote, address home, string channel, uint8 decimals, bool isNative, bool isExternal) returns()
func (_TokenRouter *TokenRouterTransactorSession) SetTokenConfig(denom string, token common.Address, remote common.Address, home common.Address, channel string, decimals uint8, isNative bool, isExternal bool) (*types.Transaction, error) {
	return _TokenRouter.Contract.SetTokenConfig(&_TokenRouter.TransactOpts, denom, token, remote, home, channel, decimals, isNative, isExternal)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenRouter *TokenRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenRouter *TokenRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenRouter.Contract.TransferOwnership(&_TokenRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenRouter *TokenRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenRouter.Contract.TransferOwnership(&_TokenRouter.TransactOpts, newOwner)
}

// TokenRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenRouter contract.
type TokenRouterOwnershipTransferredIterator struct {
	Event *TokenRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRouterOwnershipTransferred)
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
		it.Event = new(TokenRouterOwnershipTransferred)
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
func (it *TokenRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRouterOwnershipTransferred represents a OwnershipTransferred event raised by the TokenRouter contract.
type TokenRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenRouter *TokenRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenRouterOwnershipTransferredIterator{contract: _TokenRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenRouter *TokenRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRouterOwnershipTransferred)
				if err := _TokenRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenRouter *TokenRouterFilterer) ParseOwnershipTransferred(log types.Log) (*TokenRouterOwnershipTransferred, error) {
	event := new(TokenRouterOwnershipTransferred)
	if err := _TokenRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRouterTokenConfigRemovedIterator is returned from FilterTokenConfigRemoved and is used to iterate over the raw logs and unpacked data for TokenConfigRemoved events raised by the TokenRouter contract.
type TokenRouterTokenConfigRemovedIterator struct {
	Event *TokenRouterTokenConfigRemoved // Event containing the contract specifics and raw log

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
func (it *TokenRouterTokenConfigRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRouterTokenConfigRemoved)
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
		it.Event = new(TokenRouterTokenConfigRemoved)
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
func (it *TokenRouterTokenConfigRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRouterTokenConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRouterTokenConfigRemoved represents a TokenConfigRemoved event raised by the TokenRouter contract.
type TokenRouterTokenConfigRemoved struct {
	Denom common.Hash
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenConfigRemoved is a free log retrieval operation binding the contract event 0x7b03ec0d67d50c23129a8d5eddce82aab209af695b966d4bf7c9b4ef09274770.
//
// Solidity: event TokenConfigRemoved(string indexed denom)
func (_TokenRouter *TokenRouterFilterer) FilterTokenConfigRemoved(opts *bind.FilterOpts, denom []string) (*TokenRouterTokenConfigRemovedIterator, error) {

	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}

	logs, sub, err := _TokenRouter.contract.FilterLogs(opts, "TokenConfigRemoved", denomRule)
	if err != nil {
		return nil, err
	}
	return &TokenRouterTokenConfigRemovedIterator{contract: _TokenRouter.contract, event: "TokenConfigRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenConfigRemoved is a free log subscription operation binding the contract event 0x7b03ec0d67d50c23129a8d5eddce82aab209af695b966d4bf7c9b4ef09274770.
//
// Solidity: event TokenConfigRemoved(string indexed denom)
func (_TokenRouter *TokenRouterFilterer) WatchTokenConfigRemoved(opts *bind.WatchOpts, sink chan<- *TokenRouterTokenConfigRemoved, denom []string) (event.Subscription, error) {

	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}

	logs, sub, err := _TokenRouter.contract.WatchLogs(opts, "TokenConfigRemoved", denomRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRouterTokenConfigRemoved)
				if err := _TokenRouter.contract.UnpackLog(event, "TokenConfigRemoved", log); err != nil {
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

// ParseTokenConfigRemoved is a log parse operation binding the contract event 0x7b03ec0d67d50c23129a8d5eddce82aab209af695b966d4bf7c9b4ef09274770.
//
// Solidity: event TokenConfigRemoved(string indexed denom)
func (_TokenRouter *TokenRouterFilterer) ParseTokenConfigRemoved(log types.Log) (*TokenRouterTokenConfigRemoved, error) {
	event := new(TokenRouterTokenConfigRemoved)
	if err := _TokenRouter.contract.UnpackLog(event, "TokenConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRouterTokenConfigSetIterator is returned from FilterTokenConfigSet and is used to iterate over the raw logs and unpacked data for TokenConfigSet events raised by the TokenRouter contract.
type TokenRouterTokenConfigSetIterator struct {
	Event *TokenRouterTokenConfigSet // Event containing the contract specifics and raw log

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
func (it *TokenRouterTokenConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRouterTokenConfigSet)
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
		it.Event = new(TokenRouterTokenConfigSet)
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
func (it *TokenRouterTokenConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRouterTokenConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRouterTokenConfigSet represents a TokenConfigSet event raised by the TokenRouter contract.
type TokenRouterTokenConfigSet struct {
	Denom      common.Hash
	Token      common.Address
	Remote     common.Address
	Home       common.Address
	Channel    string
	Decimals   uint8
	IsNative   bool
	IsExternal bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenConfigSet is a free log retrieval operation binding the contract event 0x8da271fee9489a2d04f4b19df1eb40abce1b8139522702443aa73c5392d66c98.
//
// Solidity: event TokenConfigSet(string indexed denom, address token, address remote, address home, string channel, uint8 decimals, bool isNative, bool isExternal)
func (_TokenRouter *TokenRouterFilterer) FilterTokenConfigSet(opts *bind.FilterOpts, denom []string) (*TokenRouterTokenConfigSetIterator, error) {

	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}

	logs, sub, err := _TokenRouter.contract.FilterLogs(opts, "TokenConfigSet", denomRule)
	if err != nil {
		return nil, err
	}
	return &TokenRouterTokenConfigSetIterator{contract: _TokenRouter.contract, event: "TokenConfigSet", logs: logs, sub: sub}, nil
}

// WatchTokenConfigSet is a free log subscription operation binding the contract event 0x8da271fee9489a2d04f4b19df1eb40abce1b8139522702443aa73c5392d66c98.
//
// Solidity: event TokenConfigSet(string indexed denom, address token, address remote, address home, string channel, uint8 decimals, bool isNative, bool isExternal)
func (_TokenRouter *TokenRouterFilterer) WatchTokenConfigSet(opts *bind.WatchOpts, sink chan<- *TokenRouterTokenConfigSet, denom []string) (event.Subscription, error) {

	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}

	logs, sub, err := _TokenRouter.contract.WatchLogs(opts, "TokenConfigSet", denomRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRouterTokenConfigSet)
				if err := _TokenRouter.contract.UnpackLog(event, "TokenConfigSet", log); err != nil {
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

// ParseTokenConfigSet is a log parse operation binding the contract event 0x8da271fee9489a2d04f4b19df1eb40abce1b8139522702443aa73c5392d66c98.
//
// Solidity: event TokenConfigSet(string indexed denom, address token, address remote, address home, string channel, uint8 decimals, bool isNative, bool isExternal)
func (_TokenRouter *TokenRouterFilterer) ParseTokenConfigSet(log types.Log) (*TokenRouterTokenConfigSet, error) {
	event := new(TokenRouterTokenConfigSet)
	if err := _TokenRouter.contract.UnpackLog(event, "TokenConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRouterTokenConfigValidationFailedIterator is returned from FilterTokenConfigValidationFailed and is used to iterate over the raw logs and unpacked data for TokenConfigValidationFailed events raised by the TokenRouter contract.
type TokenRouterTokenConfigValidationFailedIterator struct {
	Event *TokenRouterTokenConfigValidationFailed // Event containing the contract specifics and raw log

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
func (it *TokenRouterTokenConfigValidationFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRouterTokenConfigValidationFailed)
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
		it.Event = new(TokenRouterTokenConfigValidationFailed)
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
func (it *TokenRouterTokenConfigValidationFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRouterTokenConfigValidationFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRouterTokenConfigValidationFailed represents a TokenConfigValidationFailed event raised by the TokenRouter contract.
type TokenRouterTokenConfigValidationFailed struct {
	Denom    common.Hash
	Reason   string
	Token    common.Address
	Remote   common.Address
	Home     common.Address
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenConfigValidationFailed is a free log retrieval operation binding the contract event 0x81403746f5f147d73940fc611424ecd32b37b9ed0687cd4ef68b945f4fedcb61.
//
// Solidity: event TokenConfigValidationFailed(string indexed denom, string reason, address token, address remote, address home, uint8 decimals)
func (_TokenRouter *TokenRouterFilterer) FilterTokenConfigValidationFailed(opts *bind.FilterOpts, denom []string) (*TokenRouterTokenConfigValidationFailedIterator, error) {

	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}

	logs, sub, err := _TokenRouter.contract.FilterLogs(opts, "TokenConfigValidationFailed", denomRule)
	if err != nil {
		return nil, err
	}
	return &TokenRouterTokenConfigValidationFailedIterator{contract: _TokenRouter.contract, event: "TokenConfigValidationFailed", logs: logs, sub: sub}, nil
}

// WatchTokenConfigValidationFailed is a free log subscription operation binding the contract event 0x81403746f5f147d73940fc611424ecd32b37b9ed0687cd4ef68b945f4fedcb61.
//
// Solidity: event TokenConfigValidationFailed(string indexed denom, string reason, address token, address remote, address home, uint8 decimals)
func (_TokenRouter *TokenRouterFilterer) WatchTokenConfigValidationFailed(opts *bind.WatchOpts, sink chan<- *TokenRouterTokenConfigValidationFailed, denom []string) (event.Subscription, error) {

	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}

	logs, sub, err := _TokenRouter.contract.WatchLogs(opts, "TokenConfigValidationFailed", denomRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRouterTokenConfigValidationFailed)
				if err := _TokenRouter.contract.UnpackLog(event, "TokenConfigValidationFailed", log); err != nil {
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

// ParseTokenConfigValidationFailed is a log parse operation binding the contract event 0x81403746f5f147d73940fc611424ecd32b37b9ed0687cd4ef68b945f4fedcb61.
//
// Solidity: event TokenConfigValidationFailed(string indexed denom, string reason, address token, address remote, address home, uint8 decimals)
func (_TokenRouter *TokenRouterFilterer) ParseTokenConfigValidationFailed(log types.Log) (*TokenRouterTokenConfigValidationFailed, error) {
	event := new(TokenRouterTokenConfigValidationFailed)
	if err := _TokenRouter.contract.UnpackLog(event, "TokenConfigValidationFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
