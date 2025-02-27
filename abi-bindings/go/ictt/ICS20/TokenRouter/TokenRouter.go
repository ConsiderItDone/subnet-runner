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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"TokenConfigRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"remote\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"home\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExternal\",\"type\":\"bool\"}],\"name\":\"TokenConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"remote\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"home\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"TokenConfigValidationFailed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"denoms\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"remotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"homes\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"channels\",\"type\":\"string[]\"},{\"internalType\":\"uint8[]\",\"name\":\"decimalsArray\",\"type\":\"uint8[]\"},{\"internalType\":\"bool[]\",\"name\":\"isNativeArray\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"isExternalArray\",\"type\":\"bool[]\"}],\"name\":\"batchSetTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"sourceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"targetDecimals\",\"type\":\"uint8\"}],\"name\":\"convertDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"denomToConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"home\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isExternal\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getTokenChannel\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"channels\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getTokenConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"home\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remote\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isExternal\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"isTokenSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"removeTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"home\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isExternal\",\"type\":\"bool\"}],\"name\":\"setTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051611b59380380611b5983398101604081905261002e916100d7565b806001600160a01b03811661005c57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b6100658161006c565b5050610104565b600180546001600160a01b031916905561008581610088565b50565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156100e7575f80fd5b81516001600160a01b03811681146100fd575f80fd5b9392505050565b611a48806101115f395ff3fe608060405234801561000f575f80fd5b50600436106100e5575f3560e01c806381f868f711610088578063b450e5d811610063578063b450e5d8146101ee578063c8a999df14610201578063e30c397814610214578063f2fde38b14610225575f80fd5b806381f868f7146101975780638da5cb5b146101aa578063b03171c2146101ce575f80fd5b80636ae64905116100c35780636ae6490514610151578063715018a61461017257806374a707421461017c57806379ba50971461018f575f80fd5b80630417cf8e146100e95780630d3652ec14610108578063605ba00d1461012b575b5f80fd5b6100f1601281565b60405160ff90911681526020015b60405180910390f35b61011b61011636600461119a565b610238565b60405190151581526020016100ff565b61013e61013936600461119a565b610380565b6040516100ff9796959493929190611221565b61016461015f36600461128d565b61046b565b6040519081526020016100ff565b61017a6104e0565b005b61017a61018a3660046112eb565b6104f3565b61017a6106b2565b61013e6101a536600461119a565b6106f6565b5f546001600160a01b03165b6040516001600160a01b0390911681526020016100ff565b6101e16101dc36600461119a565b6108ba565b6040516100ff91906113ac565b61017a6101fc366004611595565b610a4b565b61017a61020f36600461119a565b610d1e565b6001546001600160a01b03166101b6565b61017a6102333660046116c8565b610e4e565b5f8060028360405161024a91906116e1565b90815260408051918290036020908101832060e08401835280546001600160a01b03908116855260018201548116928501929092526002810154909116918301919091526003810180546060840191906102a3906116fc565b80601f01602080910402602001604051908101604052809291908181526020018280546102cf906116fc565b801561031a5780601f106102f15761010080835404028352916020019161031a565b820191905f5260205f20905b8154815290600101906020018083116102fd57829003601f168201915b50505091835250506004919091015460ff808216602084015261010082048116151560408401526201000090910416151560609091015280519091506001600160a01b031615801590610379575060408101516001600160a01b031615155b9392505050565b8051808201602090810180516002808352938301929094019190912092905281546001830154918301546003840180546001600160a01b03938416959484169492909316926103ce906116fc565b80601f01602080910402602001604051908101604052809291908181526020018280546103fa906116fc565b80156104455780601f1061041c57610100808354040283529160200191610445565b820191905f5260205f20905b81548152906001019060200180831161042857829003601f168201915b5050506004909301549192505060ff808216916101008104821691620100009091041687565b5f8160ff168360ff1603610480575082610379565b8160ff168360ff1611156104bb575f6104998385611748565b60ff1690506104a981600a611847565b6104b39086611852565b915050610379565b5f6104c68484611748565b60ff1690506104d681600a611847565b6104b39086611871565b6104e8610ebe565b6104f15f610eea565b565b6104fb610ebe565b5f8061050989888a88610f03565b915091508181906105365760405162461bcd60e51b815260040161052d91906113ac565b60405180910390fd5b506040518060e001604052808a6001600160a01b03168152602001896001600160a01b03168152602001886001600160a01b031681526020018781526020018660ff168152602001851515815260200184151581525060028b60405161059c91906116e1565b9081526040805160209281900383019020835181546001600160a01b03199081166001600160a01b03928316178355938501516001830180548616918316919091179055918401516002820180549094169216919091179091556060820151600382019061060a90826118d4565b5060808201516004909101805460a084015160c0909401511515620100000262ff0000199415156101000261ffff1990921660ff90941693909317179290921617905560405161065b908b906116e1565b60405180910390207f8da271fee9489a2d04f4b19df1eb40abce1b8139522702443aa73c5392d66c988a8a8a8a8a8a8a60405161069e9796959493929190611221565b60405180910390a250505050505050505050565b60015433906001600160a01b031681146106ea5760405163118cdaa760e01b81526001600160a01b038216600482015260240161052d565b6106f381610eea565b50565b5f805f60605f805f8060028960405161070f91906116e1565b90815260408051918290036020908101832060e08401835280546001600160a01b0390811685526001820154811692850192909252600281015490911691830191909152600381018054606084019190610768906116fc565b80601f0160208091040260200160405190810160405280929190818152602001828054610794906116fc565b80156107df5780601f106107b6576101008083540402835291602001916107df565b820191905f5260205f20905b8154815290600101906020018083116107c257829003601f168201915b50505091835250506004919091015460ff80821660208085019190915261010083048216151560408086019190915262010000909304909116151560609093019290925282518151808301909252601d82527f546f6b656e526f757465723a20436f6e666967206e6f7420666f756e64000000928201929092529192506001600160a01b03166108825760405162461bcd60e51b815260040161052d91906113ac565b508051602082015160408301516060840151608085015160a086015160c090960151949e939d50919b50995097509195509350915050565b60605f6002836040516108cd91906116e1565b90815260408051918290036020908101832060e08401835280546001600160a01b0390811685526001820154811692850192909252600281015490911691830191909152600381018054606084019190610926906116fc565b80601f0160208091040260200160405190810160405280929190818152602001828054610952906116fc565b801561099d5780601f106109745761010080835404028352916020019161099d565b820191905f5260205f20905b81548152906001019060200180831161098057829003601f168201915b50505091835250506004919091015460ff80821660208085019190915261010083048216151560408086019190915262010000909304909116151560609093019290925282518151808301909252601d82527f546f6b656e526f757465723a20436f6e666967206e6f7420666f756e64000000928201929092529192506001600160a01b0316610a405760405162461bcd60e51b815260040161052d91906113ac565b506060015192915050565b610a53610ebe565b86518851148015610a65575084518851145b8015610a72575085518851145b8015610a7f575082518851145b8015610a8c575081518851145b8015610a99575080518851145b6040518060600160405280602381526020016119f06023913990610ad05760405162461bcd60e51b815260040161052d91906113ac565b505f5b8851811015610d13575f80610b4e8a8481518110610af357610af3611994565b6020026020010151898581518110610b0d57610b0d611994565b60200260200101518b8681518110610b2757610b27611994565b6020026020010151898781518110610b4157610b41611994565b6020026020010151610f03565b9150915081610c30578a8381518110610b6957610b69611994565b6020026020010151604051610b7e91906116e1565b60405180910390207f81403746f5f147d73940fc611424ecd32b37b9ed0687cd4ef68b945f4fedcb61828c8681518110610bba57610bba611994565b60200260200101518c8781518110610bd457610bd4611994565b60200260200101518c8881518110610bee57610bee611994565b60200260200101518b8981518110610c0857610c08611994565b6020026020010151604051610c219594939291906119a8565b60405180910390a25050610d0b565b610d088b8481518110610c4557610c45611994565b60200260200101518b8581518110610c5f57610c5f611994565b60200260200101518b8681518110610c7957610c79611994565b60200260200101518b8781518110610c9357610c93611994565b60200260200101518b8881518110610cad57610cad611994565b60200260200101518b8981518110610cc757610cc7611994565b60200260200101518b8a81518110610ce157610ce1611994565b60200260200101518b8b81518110610cfb57610cfb611994565b60200260200101516104f3565b50505b600101610ad3565b505050505050505050565b610d26610ebe565b5f6001600160a01b0316600282604051610d4091906116e1565b90815260408051918290036020908101832054838301909252601d83527f546f6b656e526f757465723a20436f6e666967206e6f7420666f756e640000009083015290916001600160a01b0390911603610dad5760405162461bcd60e51b815260040161052d91906113ac565b50600281604051610dbe91906116e1565b90815260405190819003602001902080546001600160a01b031990811682556001820180548216905560028201805490911690555f610e00600383018261109f565b50600401805462ffffff19169055604051610e1c9082906116e1565b604051908190038120907f7b03ec0d67d50c23129a8d5eddce82aab209af695b966d4bf7c9b4ef09274770905f90a250565b610e56610ebe565b600180546001600160a01b0383166001600160a01b03199091168117909155610e865f546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b5f546001600160a01b031633146104f15760405163118cdaa760e01b815233600482015260240161052d565b600180546001600160a01b03191690556106f381611050565b5f60606001600160a01b038616610f5157505060408051808201909152601f81527f546f6b656e526f757465723a205a65726f20746f6b656e20616464726573730060208201525f90611047565b6001600160a01b038516610f9c57505060408051808201909152601e81527f546f6b656e526f757465723a205a65726f20686f6d652061646472657373000060208201525f90611047565b6001600160a01b038416610fea575f6040518060400160405280602081526020017f546f6b656e526f757465723a205a65726f2072656d6f7465206164647265737381525091509150611047565b601260ff8416111561103357505060408051808201909152601d81527f546f6b656e526f757465723a20496e76616c696420646563696d616c7300000060208201525f90611047565b505060408051602081019091525f81526001905b94509492505050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5080546110ab906116fc565b5f825580601f106110ba575050565b601f0160209004905f5260205f20908101906106f391905b808211156110e5575f81556001016110d2565b5090565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611126576111266110e9565b604052919050565b5f82601f83011261113d575f80fd5b813567ffffffffffffffff811115611157576111576110e9565b61116a601f8201601f19166020016110fd565b81815284602083860101111561117e575f80fd5b816020850160208301375f918101602001919091529392505050565b5f602082840312156111aa575f80fd5b813567ffffffffffffffff8111156111c0575f80fd5b6111cc8482850161112e565b949350505050565b5f5b838110156111ee5781810151838201526020016111d6565b50505f910152565b5f815180845261120d8160208601602086016111d4565b601f01601f19169290920160200192915050565b6001600160a01b03888116825287811660208301528616604082015260e0606082018190525f90611254908301876111f6565b60ff9590951660808301525091151560a0830152151560c090910152949350505050565b803560ff81168114611288575f80fd5b919050565b5f805f6060848603121561129f575f80fd5b833592506112af60208501611278565b91506112bd60408501611278565b90509250925092565b80356001600160a01b0381168114611288575f80fd5b80358015158114611288575f80fd5b5f805f805f805f80610100898b031215611303575f80fd5b883567ffffffffffffffff8082111561131a575f80fd5b6113268c838d0161112e565b995061133460208c016112c6565b985061134260408c016112c6565b975061135060608c016112c6565b965060808b0135915080821115611365575f80fd5b506113728b828c0161112e565b94505061138160a08a01611278565b925061138f60c08a016112dc565b915061139d60e08a016112dc565b90509295985092959890939650565b602081525f61037960208301846111f6565b5f67ffffffffffffffff8211156113d7576113d76110e9565b5060051b60200190565b5f82601f8301126113f0575f80fd5b81356020611405611400836113be565b6110fd565b82815260059290921b84018101918181019086841115611423575f80fd5b8286015b8481101561146157803567ffffffffffffffff811115611445575f80fd5b6114538986838b010161112e565b845250918301918301611427565b509695505050505050565b5f82601f83011261147b575f80fd5b8135602061148b611400836113be565b8083825260208201915060208460051b8701019350868411156114ac575f80fd5b602086015b84811015611461576114c2816112c6565b83529183019183016114b1565b5f82601f8301126114de575f80fd5b813560206114ee611400836113be565b8083825260208201915060208460051b87010193508684111561150f575f80fd5b602086015b848110156114615761152581611278565b8352918301918301611514565b5f82601f830112611541575f80fd5b81356020611551611400836113be565b8083825260208201915060208460051b870101935086841115611572575f80fd5b602086015b8481101561146157611588816112dc565b8352918301918301611577565b5f805f805f805f80610100898b0312156115ad575f80fd5b883567ffffffffffffffff808211156115c4575f80fd5b6115d08c838d016113e1565b995060208b01359150808211156115e5575f80fd5b6115f18c838d0161146c565b985060408b0135915080821115611606575f80fd5b6116128c838d0161146c565b975060608b0135915080821115611627575f80fd5b6116338c838d0161146c565b965060808b0135915080821115611648575f80fd5b6116548c838d016113e1565b955060a08b0135915080821115611669575f80fd5b6116758c838d016114cf565b945060c08b013591508082111561168a575f80fd5b6116968c838d01611532565b935060e08b01359150808211156116ab575f80fd5b506116b88b828c01611532565b9150509295985092959890939650565b5f602082840312156116d8575f80fd5b610379826112c6565b5f82516116f28184602087016111d4565b9190910192915050565b600181811c9082168061171057607f821691505b60208210810361172e57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b60ff828116828216039081111561176157611761611734565b92915050565b600181815b808511156117a157815f190482111561178757611787611734565b8085161561179457918102915b93841c939080029061176c565b509250929050565b5f826117b757506001611761565b816117c357505f611761565b81600181146117d957600281146117e3576117ff565b6001915050611761565b60ff8411156117f4576117f4611734565b50506001821b611761565b5060208310610133831016604e8410600b8410161715611822575081810a611761565b61182c8383611767565b805f190482111561183f5761183f611734565b029392505050565b5f61037983836117a9565b5f8261186c57634e487b7160e01b5f52601260045260245ffd5b500490565b808202811582820484141761176157611761611734565b601f8211156118cf57805f5260205f20601f840160051c810160208510156118ad5750805b601f840160051c820191505b818110156118cc575f81556001016118b9565b50505b505050565b815167ffffffffffffffff8111156118ee576118ee6110e9565b611902816118fc84546116fc565b84611888565b602080601f831160018114611935575f841561191e5750858301515b5f19600386901b1c1916600185901b17855561198c565b5f85815260208120601f198616915b8281101561196357888601518255948401946001909101908401611944565b508582101561198057878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b634e487b7160e01b5f52603260045260245ffd5b60a081525f6119ba60a08301886111f6565b6001600160a01b03968716602084015294861660408301525091909316606082015260ff90921660809092019190915291905056fe546f6b656e526f757465723a204172726179206c656e67746873206d69736d61746368a2646970667358221220fd437a7f70d0289fc8b37d66582ebecf3624ebbe439bd81262220bab9b46f61664736f6c63430008190033",
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

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TokenRouter *TokenRouterCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TokenRouter *TokenRouterSession) PendingOwner() (common.Address, error) {
	return _TokenRouter.Contract.PendingOwner(&_TokenRouter.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TokenRouter *TokenRouterCallerSession) PendingOwner() (common.Address, error) {
	return _TokenRouter.Contract.PendingOwner(&_TokenRouter.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenRouter *TokenRouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenRouter *TokenRouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenRouter.Contract.AcceptOwnership(&_TokenRouter.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenRouter *TokenRouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenRouter.Contract.AcceptOwnership(&_TokenRouter.TransactOpts)
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

// TokenRouterOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the TokenRouter contract.
type TokenRouterOwnershipTransferStartedIterator struct {
	Event *TokenRouterOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *TokenRouterOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRouterOwnershipTransferStarted)
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
		it.Event = new(TokenRouterOwnershipTransferStarted)
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
func (it *TokenRouterOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRouterOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRouterOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the TokenRouter contract.
type TokenRouterOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_TokenRouter *TokenRouterFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenRouterOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenRouter.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenRouterOwnershipTransferStartedIterator{contract: _TokenRouter.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_TokenRouter *TokenRouterFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *TokenRouterOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenRouter.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRouterOwnershipTransferStarted)
				if err := _TokenRouter.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_TokenRouter *TokenRouterFilterer) ParseOwnershipTransferStarted(log types.Log) (*TokenRouterOwnershipTransferStarted, error) {
	event := new(TokenRouterOwnershipTransferStarted)
	if err := _TokenRouter.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
