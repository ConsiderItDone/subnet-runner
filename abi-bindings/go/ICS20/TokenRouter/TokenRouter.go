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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchSetTokenConfig\",\"inputs\":[{\"name\":\"denoms\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"remotes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"homes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"channels\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"decimalsArray\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"},{\"name\":\"isNativeArray\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"isExternalArray\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertDecimals\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"targetDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"denomToConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"home\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isNative\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isExternal\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenChannel\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"channels\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenConfig\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"home\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isNative\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isExternal\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTokenSupported\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeTokenConfig\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenConfig\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"home\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"channel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isNative\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isExternal\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenConfigRemoved\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenConfigSet\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"remote\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"home\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"channel\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"isNative\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"isExternal\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenConfigValidationFailed\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"remote\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"home\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051611a2c380380611a2c83398101604081905261002e916100bb565b806001600160a01b03811661005c57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b6100658161006c565b50506100e8565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156100cb575f80fd5b81516001600160a01b03811681146100e1575f80fd5b9392505050565b611937806100f55f395ff3fe608060405234801561000f575f80fd5b50600436106100b1575f3560e01c806381f868f71161006e57806381f868f71461015b5780638da5cb5b1461016e578063b03171c214610188578063b450e5d8146101a8578063c8a999df146101bb578063f2fde38b146101ce575f80fd5b80630417cf8e146100b55780630d3652ec146100d4578063605ba00d146100f75780636ae649051461011d578063715018a61461013e57806374a7074214610148575b5f80fd5b6100bd601281565b60405160ff90911681526020015b60405180910390f35b6100e76100e23660046110b2565b6101e1565b60405190151581526020016100cb565b61010a6101053660046110b2565b610329565b6040516100cb9796959493929190611139565b61013061012b3660046111a5565b610413565b6040519081526020016100cb565b610146610488565b005b610146610156366004611203565b61049b565b61010a6101693660046110b2565b61065a565b5f546040516001600160a01b0390911681526020016100cb565b61019b6101963660046110b2565b61081e565b6040516100cb91906112c4565b6101466101b63660046114ad565b6109af565b6101466101c93660046110b2565b610c82565b6101466101dc3660046115e0565b610db2565b5f806001836040516101f391906115f9565b90815260408051918290036020908101832060e08401835280546001600160a01b039081168552600182015481169285019290925260028101549091169183019190915260038101805460608401919061024c90611614565b80601f016020809104026020016040519081016040528092919081815260200182805461027890611614565b80156102c35780601f1061029a576101008083540402835291602001916102c3565b820191905f5260205f20905b8154815290600101906020018083116102a657829003601f168201915b50505091835250506004919091015460ff808216602084015261010082048116151560408401526201000090910416151560609091015280519091506001600160a01b031615801590610322575060408101516001600160a01b031615155b9392505050565b80516020818301810180516001808352938301929094019190912092905281549082015460028301546003840180546001600160a01b039485169593851694909216929161037690611614565b80601f01602080910402602001604051908101604052809291908181526020018280546103a290611614565b80156103ed5780601f106103c4576101008083540402835291602001916103ed565b820191905f5260205f20905b8154815290600101906020018083116103d057829003601f168201915b5050506004909301549192505060ff808216916101008104821691620100009091041687565b5f8160ff168360ff1603610428575082610322565b8160ff168360ff161115610463575f6104418385611660565b60ff16905061045181600a61175f565b61045b908661176a565b915050610322565b5f61046e8484611660565b60ff16905061047e81600a61175f565b61045b9086611789565b610490610def565b6104995f610e1b565b565b6104a3610def565b5f806104b189888a88610e6a565b915091508181906104de5760405162461bcd60e51b81526004016104d591906112c4565b60405180910390fd5b506040518060e001604052808a6001600160a01b03168152602001896001600160a01b03168152602001886001600160a01b031681526020018781526020018660ff168152602001851515815260200184151581525060018b60405161054491906115f9565b9081526040805160209281900383019020835181546001600160a01b03199081166001600160a01b0392831617835593850151600183018054861691831691909117905591840151600282018054909416921691909117909155606082015160038201906105b290826117ec565b5060808201516004909101805460a084015160c0909401511515620100000262ff0000199415156101000261ffff1990921660ff909416939093171792909216179055604051610603908b906115f9565b60405180910390207f8da271fee9489a2d04f4b19df1eb40abce1b8139522702443aa73c5392d66c988a8a8a8a8a8a8a6040516106469796959493929190611139565b60405180910390a250505050505050505050565b5f805f60605f805f8060018960405161067391906115f9565b90815260408051918290036020908101832060e08401835280546001600160a01b03908116855260018201548116928501929092526002810154909116918301919091526003810180546060840191906106cc90611614565b80601f01602080910402602001604051908101604052809291908181526020018280546106f890611614565b80156107435780601f1061071a57610100808354040283529160200191610743565b820191905f5260205f20905b81548152906001019060200180831161072657829003601f168201915b50505091835250506004919091015460ff80821660208085019190915261010083048216151560408086019190915262010000909304909116151560609093019290925282518151808301909252601d82527f546f6b656e526f757465723a20436f6e666967206e6f7420666f756e64000000928201929092529192506001600160a01b03166107e65760405162461bcd60e51b81526004016104d591906112c4565b508051602082015160408301516060840151608085015160a086015160c090960151949e939d50919b50995097509195509350915050565b60605f60018360405161083191906115f9565b90815260408051918290036020908101832060e08401835280546001600160a01b039081168552600182015481169285019290925260028101549091169183019190915260038101805460608401919061088a90611614565b80601f01602080910402602001604051908101604052809291908181526020018280546108b690611614565b80156109015780601f106108d857610100808354040283529160200191610901565b820191905f5260205f20905b8154815290600101906020018083116108e457829003601f168201915b50505091835250506004919091015460ff80821660208085019190915261010083048216151560408086019190915262010000909304909116151560609093019290925282518151808301909252601d82527f546f6b656e526f757465723a20436f6e666967206e6f7420666f756e64000000928201929092529192506001600160a01b03166109a45760405162461bcd60e51b81526004016104d591906112c4565b506060015192915050565b6109b7610def565b865188511480156109c9575084518851145b80156109d6575085518851145b80156109e3575082518851145b80156109f0575081518851145b80156109fd575080518851145b6040518060600160405280602381526020016119086023913990610a345760405162461bcd60e51b81526004016104d591906112c4565b505f5b8851811015610c77575f80610ab28a8481518110610a5757610a576118ac565b6020026020010151898581518110610a7157610a716118ac565b60200260200101518b8681518110610a8b57610a8b6118ac565b6020026020010151898781518110610aa557610aa56118ac565b6020026020010151610e6a565b9150915081610b94578a8381518110610acd57610acd6118ac565b6020026020010151604051610ae291906115f9565b60405180910390207f81403746f5f147d73940fc611424ecd32b37b9ed0687cd4ef68b945f4fedcb61828c8681518110610b1e57610b1e6118ac565b60200260200101518c8781518110610b3857610b386118ac565b60200260200101518c8881518110610b5257610b526118ac565b60200260200101518b8981518110610b6c57610b6c6118ac565b6020026020010151604051610b859594939291906118c0565b60405180910390a25050610c6f565b610c6c8b8481518110610ba957610ba96118ac565b60200260200101518b8581518110610bc357610bc36118ac565b60200260200101518b8681518110610bdd57610bdd6118ac565b60200260200101518b8781518110610bf757610bf76118ac565b60200260200101518b8881518110610c1157610c116118ac565b60200260200101518b8981518110610c2b57610c2b6118ac565b60200260200101518b8a81518110610c4557610c456118ac565b60200260200101518b8b81518110610c5f57610c5f6118ac565b602002602001015161049b565b50505b600101610a37565b505050505050505050565b610c8a610def565b5f6001600160a01b0316600182604051610ca491906115f9565b90815260408051918290036020908101832054838301909252601d83527f546f6b656e526f757465723a20436f6e666967206e6f7420666f756e640000009083015290916001600160a01b0390911603610d115760405162461bcd60e51b81526004016104d591906112c4565b50600181604051610d2291906115f9565b90815260405190819003602001902080546001600160a01b031990811682556001820180548216905560028201805490911690555f610d646003830182610fb7565b50600401805462ffffff19169055604051610d809082906115f9565b604051908190038120907f7b03ec0d67d50c23129a8d5eddce82aab209af695b966d4bf7c9b4ef09274770905f90a250565b610dba610def565b6001600160a01b038116610de357604051631e4fbdf760e01b81525f60048201526024016104d5565b610dec81610e1b565b50565b5f546001600160a01b031633146104995760405163118cdaa760e01b81523360048201526024016104d5565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f60606001600160a01b038616610eb857505060408051808201909152601f81527f546f6b656e526f757465723a205a65726f20746f6b656e20616464726573730060208201525f90610fae565b6001600160a01b038516610f0357505060408051808201909152601e81527f546f6b656e526f757465723a205a65726f20686f6d652061646472657373000060208201525f90610fae565b6001600160a01b038416610f51575f6040518060400160405280602081526020017f546f6b656e526f757465723a205a65726f2072656d6f7465206164647265737381525091509150610fae565b601260ff84161115610f9a57505060408051808201909152601d81527f546f6b656e526f757465723a20496e76616c696420646563696d616c7300000060208201525f90610fae565b505060408051602081019091525f81526001905b94509492505050565b508054610fc390611614565b5f825580601f10610fd2575050565b601f0160209004905f5260205f2090810190610dec91905b80821115610ffd575f8155600101610fea565b5090565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561103e5761103e611001565b604052919050565b5f82601f830112611055575f80fd5b813567ffffffffffffffff81111561106f5761106f611001565b611082601f8201601f1916602001611015565b818152846020838601011115611096575f80fd5b816020850160208301375f918101602001919091529392505050565b5f602082840312156110c2575f80fd5b813567ffffffffffffffff8111156110d8575f80fd5b6110e484828501611046565b949350505050565b5f5b838110156111065781810151838201526020016110ee565b50505f910152565b5f81518084526111258160208601602086016110ec565b601f01601f19169290920160200192915050565b6001600160a01b03888116825287811660208301528616604082015260e0606082018190525f9061116c9083018761110e565b60ff9590951660808301525091151560a0830152151560c090910152949350505050565b803560ff811681146111a0575f80fd5b919050565b5f805f606084860312156111b7575f80fd5b833592506111c760208501611190565b91506111d560408501611190565b90509250925092565b80356001600160a01b03811681146111a0575f80fd5b803580151581146111a0575f80fd5b5f805f805f805f80610100898b03121561121b575f80fd5b883567ffffffffffffffff80821115611232575f80fd5b61123e8c838d01611046565b995061124c60208c016111de565b985061125a60408c016111de565b975061126860608c016111de565b965060808b013591508082111561127d575f80fd5b5061128a8b828c01611046565b94505061129960a08a01611190565b92506112a760c08a016111f4565b91506112b560e08a016111f4565b90509295985092959890939650565b602081525f610322602083018461110e565b5f67ffffffffffffffff8211156112ef576112ef611001565b5060051b60200190565b5f82601f830112611308575f80fd5b8135602061131d611318836112d6565b611015565b82815260059290921b8401810191818101908684111561133b575f80fd5b8286015b8481101561137957803567ffffffffffffffff81111561135d575f80fd5b61136b8986838b0101611046565b84525091830191830161133f565b509695505050505050565b5f82601f830112611393575f80fd5b813560206113a3611318836112d6565b8083825260208201915060208460051b8701019350868411156113c4575f80fd5b602086015b84811015611379576113da816111de565b83529183019183016113c9565b5f82601f8301126113f6575f80fd5b81356020611406611318836112d6565b8083825260208201915060208460051b870101935086841115611427575f80fd5b602086015b848110156113795761143d81611190565b835291830191830161142c565b5f82601f830112611459575f80fd5b81356020611469611318836112d6565b8083825260208201915060208460051b87010193508684111561148a575f80fd5b602086015b84811015611379576114a0816111f4565b835291830191830161148f565b5f805f805f805f80610100898b0312156114c5575f80fd5b883567ffffffffffffffff808211156114dc575f80fd5b6114e88c838d016112f9565b995060208b01359150808211156114fd575f80fd5b6115098c838d01611384565b985060408b013591508082111561151e575f80fd5b61152a8c838d01611384565b975060608b013591508082111561153f575f80fd5b61154b8c838d01611384565b965060808b0135915080821115611560575f80fd5b61156c8c838d016112f9565b955060a08b0135915080821115611581575f80fd5b61158d8c838d016113e7565b945060c08b01359150808211156115a2575f80fd5b6115ae8c838d0161144a565b935060e08b01359150808211156115c3575f80fd5b506115d08b828c0161144a565b9150509295985092959890939650565b5f602082840312156115f0575f80fd5b610322826111de565b5f825161160a8184602087016110ec565b9190910192915050565b600181811c9082168061162857607f821691505b60208210810361164657634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b60ff82811682821603908111156116795761167961164c565b92915050565b600181815b808511156116b957815f190482111561169f5761169f61164c565b808516156116ac57918102915b93841c9390800290611684565b509250929050565b5f826116cf57506001611679565b816116db57505f611679565b81600181146116f157600281146116fb57611717565b6001915050611679565b60ff84111561170c5761170c61164c565b50506001821b611679565b5060208310610133831016604e8410600b841016171561173a575081810a611679565b611744838361167f565b805f19048211156117575761175761164c565b029392505050565b5f61032283836116c1565b5f8261178457634e487b7160e01b5f52601260045260245ffd5b500490565b80820281158282048414176116795761167961164c565b601f8211156117e757805f5260205f20601f840160051c810160208510156117c55750805b601f840160051c820191505b818110156117e4575f81556001016117d1565b50505b505050565b815167ffffffffffffffff81111561180657611806611001565b61181a816118148454611614565b846117a0565b602080601f83116001811461184d575f84156118365750858301515b5f19600386901b1c1916600185901b1785556118a4565b5f85815260208120601f198616915b8281101561187b5788860151825594840194600190910190840161185c565b508582101561189857878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b634e487b7160e01b5f52603260045260245ffd5b60a081525f6118d260a083018861110e565b6001600160a01b03968716602084015294861660408301525091909316606082015260ff90921660809092019190915291905056fe546f6b656e526f757465723a204172726179206c656e67746873206d69736d61746368a164736f6c6343000819000a",
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
