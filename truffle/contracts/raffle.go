// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package raffle

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BCBABI is the input ABI used to generate the binding from.
const BCBABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferWithData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TransferWithData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// BCBFuncSigs maps the 4-byte function signature to its string representation.
var BCBFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"9dc29fac": "burn(address,uint256)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"8f32d59b": "isOwner()",
	"40c10f19": "mint(address,uint256)",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
	"2535f762": "transferWithData(address,uint256,bytes)",
}

// BCBBin is the compiled bytecode used for deploying new contracts.
var BCBBin = "0x60806040526006805460ff191660121790553480156200001e57600080fd5b506040516200129538038062001295833981810160405260408110156200004457600080fd5b81019080805160405193929190846401000000008211156200006557600080fd5b9083019060208201858111156200007b57600080fd5b82516401000000008111828201881017156200009657600080fd5b82525081516020918201929091019080838360005b83811015620000c5578181015183820152602001620000ab565b50505050905090810190601f168015620000f35780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011757600080fd5b9083019060208201858111156200012d57600080fd5b82516401000000008111828201881017156200014857600080fd5b82525081516020918201929091019080838360005b83811015620001775781810151838201526020016200015d565b50505050905090810190601f168015620001a55780820380516001836020036101000a031916815260200191505b506040819052600380546001600160a01b0319163317908190556001600160a01b03169350600092507f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091508290a381516200020990600490602085019062000228565b5080516200021f90600590602084019062000228565b505050620002cd565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200026b57805160ff19168380011785556200029b565b828001600101855582156200029b579182015b828111156200029b5782518255916020019190600101906200027e565b50620002a9929150620002ad565b5090565b620002ca91905b80821115620002a95760008155600101620002b4565b90565b610fb880620002dd6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063715018a6116100a25780639dc29fac116100715780639dc29fac146103bd578063a457c2d7146103e9578063a9059cbb14610415578063dd62ed3e14610441578063f2fde38b1461046f57610116565b8063715018a61461037f5780638da5cb5b146103895780638f32d59b146103ad57806395d89b41146103b557610116565b80632535f762116100e95780632535f76214610228578063313ce567146102e3578063395093511461030157806340c10f191461032d57806370a082311461035957610116565b806306fdde031461011b578063095ea7b31461019857806318160ddd146101d857806323b872dd146101f2575b600080fd5b610123610495565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561015d578181015183820152602001610145565b50505050905090810190601f16801561018a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101c4600480360360408110156101ae57600080fd5b506001600160a01b038135169060200135610523565b604080519115158252519081900360200190f35b6101e0610539565b60408051918252519081900360200190f35b6101c46004803603606081101561020857600080fd5b506001600160a01b0381358116916020810135909116906040013561053f565b6101c46004803603606081101561023e57600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561026e57600080fd5b82018360208201111561028057600080fd5b803590602001918460018302840111640100000000831117156102a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610596945050505050565b6102eb61065f565b6040805160ff9092168252519081900360200190f35b6101c46004803603604081101561031757600080fd5b506001600160a01b038135169060200135610668565b6101c46004803603604081101561034357600080fd5b506001600160a01b0381351690602001356106a4565b6101e06004803603602081101561036f57600080fd5b50356001600160a01b03166106f7565b610387610712565b005b6103916107a3565b604080516001600160a01b039092168252519081900360200190f35b6101c46107b2565b6101236107c3565b6101c4600480360360408110156103d357600080fd5b506001600160a01b03813516906020013561081e565b6101c4600480360360408110156103ff57600080fd5b506001600160a01b038135169060200135610871565b6101c46004803603604081101561042b57600080fd5b506001600160a01b0381351690602001356108ad565b6101e06004803603604081101561045757600080fd5b506001600160a01b03813581169160200135166108ba565b6103876004803603602081101561048557600080fd5b50356001600160a01b03166108e5565b6004805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561051b5780601f106104f05761010080835404028352916020019161051b565b820191906000526020600020905b8154815290600101906020018083116104fe57829003601f168201915b505050505081565b6000610530338484610938565b50600192915050565b60025490565b600061054c848484610a24565b6001600160a01b03841660009081526001602090815260408083203380855292529091205461058c918691610587908663ffffffff610b6616565b610938565b5060019392505050565b6000836001600160a01b0316336001600160a01b03167fe68ca1ec8e8e022357047aae1f96036cbb808c6dc2bbbfbd3bde507ab21098c485856040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156106125781810151838201526020016105fa565b50505050905090810190601f16801561063f5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a361065784846108ad565b949350505050565b60065460ff1681565b3360008181526001602090815260408083206001600160a01b03871684529091528120549091610530918590610587908663ffffffff610bc316565b60006106ae6107b2565b6106ed576040805162461bcd60e51b81526020600482018190526024820152600080516020610efa833981519152604482015290519081900360640190fd5b6105308383610c24565b6001600160a01b031660009081526020819052604090205490565b61071a6107b2565b610759576040805162461bcd60e51b81526020600482018190526024820152600080516020610efa833981519152604482015290519081900360640190fd5b6003546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600380546001600160a01b0319169055565b6003546001600160a01b031690565b6003546001600160a01b0316331490565b6005805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561051b5780601f106104f05761010080835404028352916020019161051b565b60006108286107b2565b610867576040805162461bcd60e51b81526020600482018190526024820152600080516020610efa833981519152604482015290519081900360640190fd5b6105308383610d14565b3360008181526001602090815260408083206001600160a01b03871684529091528120549091610530918590610587908663ffffffff610b6616565b6000610530338484610a24565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6108ed6107b2565b61092c576040805162461bcd60e51b81526020600482018190526024820152600080516020610efa833981519152604482015290519081900360640190fd5b61093581610ded565b50565b6001600160a01b03831661097d5760405162461bcd60e51b8152600401808060200182810382526024815260200180610f606024913960400191505060405180910390fd5b6001600160a01b0382166109c25760405162461bcd60e51b8152600401808060200182810382526022815260200180610ed86022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316610a695760405162461bcd60e51b8152600401808060200182810382526025815260200180610f3b6025913960400191505060405180910390fd5b6001600160a01b038216610aae5760405162461bcd60e51b8152600401808060200182810382526023815260200180610e8f6023913960400191505060405180910390fd5b6001600160a01b038316600090815260208190526040902054610ad7908263ffffffff610b6616565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610b0c908263ffffffff610bc316565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600082821115610bbd576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015610c1d576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b038216610c7f576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b600254610c92908263ffffffff610bc316565b6002556001600160a01b038216600090815260208190526040902054610cbe908263ffffffff610bc316565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6001600160a01b038216610d595760405162461bcd60e51b8152600401808060200182810382526021815260200180610f1a6021913960400191505060405180910390fd5b600254610d6c908263ffffffff610b6616565b6002556001600160a01b038216600090815260208190526040902054610d98908263ffffffff610b6616565b6001600160a01b038316600081815260208181526040808320949094558351858152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35050565b6001600160a01b038116610e325760405162461bcd60e51b8152600401808060200182810382526026815260200180610eb26026913960400191505060405180910390fd5b6003546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600380546001600160a01b0319166001600160a01b039290921691909117905556fe45524332303a207472616e7366657220746f20746865207a65726f20616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657245524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373a265627a7a7231582012ed0d0fa909d4d3f19a5425f5845b242c264601488acb0544ede9dc56d83b3164736f6c634300050b0032"

// DeployBCB deploys a new Ethereum contract, binding an instance of BCB to it.
func DeployBCB(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *BCB, error) {
	parsed, err := abi.JSON(strings.NewReader(BCBABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BCBBin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BCB{BCBCaller: BCBCaller{contract: contract}, BCBTransactor: BCBTransactor{contract: contract}, BCBFilterer: BCBFilterer{contract: contract}}, nil
}

// BCB is an auto generated Go binding around an Ethereum contract.
type BCB struct {
	BCBCaller     // Read-only binding to the contract
	BCBTransactor // Write-only binding to the contract
	BCBFilterer   // Log filterer for contract events
}

// BCBCaller is an auto generated read-only Go binding around an Ethereum contract.
type BCBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BCBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BCBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BCBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BCBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BCBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BCBSession struct {
	Contract     *BCB              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BCBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BCBCallerSession struct {
	Contract *BCBCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BCBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BCBTransactorSession struct {
	Contract     *BCBTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BCBRaw is an auto generated low-level Go binding around an Ethereum contract.
type BCBRaw struct {
	Contract *BCB // Generic contract binding to access the raw methods on
}

// BCBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BCBCallerRaw struct {
	Contract *BCBCaller // Generic read-only contract binding to access the raw methods on
}

// BCBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BCBTransactorRaw struct {
	Contract *BCBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBCB creates a new instance of BCB, bound to a specific deployed contract.
func NewBCB(address common.Address, backend bind.ContractBackend) (*BCB, error) {
	contract, err := bindBCB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BCB{BCBCaller: BCBCaller{contract: contract}, BCBTransactor: BCBTransactor{contract: contract}, BCBFilterer: BCBFilterer{contract: contract}}, nil
}

// NewBCBCaller creates a new read-only instance of BCB, bound to a specific deployed contract.
func NewBCBCaller(address common.Address, caller bind.ContractCaller) (*BCBCaller, error) {
	contract, err := bindBCB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BCBCaller{contract: contract}, nil
}

// NewBCBTransactor creates a new write-only instance of BCB, bound to a specific deployed contract.
func NewBCBTransactor(address common.Address, transactor bind.ContractTransactor) (*BCBTransactor, error) {
	contract, err := bindBCB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BCBTransactor{contract: contract}, nil
}

// NewBCBFilterer creates a new log filterer instance of BCB, bound to a specific deployed contract.
func NewBCBFilterer(address common.Address, filterer bind.ContractFilterer) (*BCBFilterer, error) {
	contract, err := bindBCB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BCBFilterer{contract: contract}, nil
}

// bindBCB binds a generic wrapper to an already deployed contract.
func bindBCB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BCBABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BCB *BCBRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BCB.Contract.BCBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BCB *BCBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BCB.Contract.BCBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BCB *BCBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BCB.Contract.BCBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BCB *BCBCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BCB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BCB *BCBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BCB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BCB *BCBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BCB.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_BCB *BCBCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BCB.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_BCB *BCBSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BCB.Contract.Allowance(&_BCB.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_BCB *BCBCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BCB.Contract.Allowance(&_BCB.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_BCB *BCBCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BCB.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_BCB *BCBSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BCB.Contract.BalanceOf(&_BCB.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_BCB *BCBCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BCB.Contract.BalanceOf(&_BCB.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BCB *BCBCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BCB.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BCB *BCBSession) Decimals() (uint8, error) {
	return _BCB.Contract.Decimals(&_BCB.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BCB *BCBCallerSession) Decimals() (uint8, error) {
	return _BCB.Contract.Decimals(&_BCB.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BCB *BCBCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BCB.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BCB *BCBSession) IsOwner() (bool, error) {
	return _BCB.Contract.IsOwner(&_BCB.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BCB *BCBCallerSession) IsOwner() (bool, error) {
	return _BCB.Contract.IsOwner(&_BCB.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BCB *BCBCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BCB.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BCB *BCBSession) Name() (string, error) {
	return _BCB.Contract.Name(&_BCB.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BCB *BCBCallerSession) Name() (string, error) {
	return _BCB.Contract.Name(&_BCB.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BCB *BCBCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BCB.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BCB *BCBSession) Owner() (common.Address, error) {
	return _BCB.Contract.Owner(&_BCB.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BCB *BCBCallerSession) Owner() (common.Address, error) {
	return _BCB.Contract.Owner(&_BCB.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BCB *BCBCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BCB.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BCB *BCBSession) Symbol() (string, error) {
	return _BCB.Contract.Symbol(&_BCB.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BCB *BCBCallerSession) Symbol() (string, error) {
	return _BCB.Contract.Symbol(&_BCB.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BCB *BCBCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BCB.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BCB *BCBSession) TotalSupply() (*big.Int, error) {
	return _BCB.Contract.TotalSupply(&_BCB.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BCB *BCBCallerSession) TotalSupply() (*big.Int, error) {
	return _BCB.Contract.TotalSupply(&_BCB.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BCB *BCBTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BCB.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BCB *BCBSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.Approve(&_BCB.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BCB *BCBTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.Approve(&_BCB.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 value) returns(bool)
func (_BCB *BCBTransactor) Burn(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _BCB.contract.Transact(opts, "burn", from, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 value) returns(bool)
func (_BCB *BCBSession) Burn(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.Burn(&_BCB.TransactOpts, from, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 value) returns(bool)
func (_BCB *BCBTransactorSession) Burn(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.Burn(&_BCB.TransactOpts, from, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BCB *BCBTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BCB.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BCB *BCBSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.DecreaseAllowance(&_BCB.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BCB *BCBTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.DecreaseAllowance(&_BCB.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BCB *BCBTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BCB.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BCB *BCBSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.IncreaseAllowance(&_BCB.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BCB *BCBTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.IncreaseAllowance(&_BCB.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_BCB *BCBTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BCB.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_BCB *BCBSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.Mint(&_BCB.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_BCB *BCBTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.Mint(&_BCB.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BCB *BCBTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BCB.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BCB *BCBSession) RenounceOwnership() (*types.Transaction, error) {
	return _BCB.Contract.RenounceOwnership(&_BCB.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BCB *BCBTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BCB.Contract.RenounceOwnership(&_BCB.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BCB *BCBTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BCB.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BCB *BCBSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.Transfer(&_BCB.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BCB *BCBTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.Transfer(&_BCB.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BCB *BCBTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BCB.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BCB *BCBSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.TransferFrom(&_BCB.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BCB *BCBTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BCB.Contract.TransferFrom(&_BCB.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BCB *BCBTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BCB.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BCB *BCBSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BCB.Contract.TransferOwnership(&_BCB.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BCB *BCBTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BCB.Contract.TransferOwnership(&_BCB.TransactOpts, newOwner)
}

// TransferWithData is a paid mutator transaction binding the contract method 0x2535f762.
//
// Solidity: function transferWithData(address to, uint256 value, bytes data) returns(bool)
func (_BCB *BCBTransactor) TransferWithData(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BCB.contract.Transact(opts, "transferWithData", to, value, data)
}

// TransferWithData is a paid mutator transaction binding the contract method 0x2535f762.
//
// Solidity: function transferWithData(address to, uint256 value, bytes data) returns(bool)
func (_BCB *BCBSession) TransferWithData(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BCB.Contract.TransferWithData(&_BCB.TransactOpts, to, value, data)
}

// TransferWithData is a paid mutator transaction binding the contract method 0x2535f762.
//
// Solidity: function transferWithData(address to, uint256 value, bytes data) returns(bool)
func (_BCB *BCBTransactorSession) TransferWithData(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BCB.Contract.TransferWithData(&_BCB.TransactOpts, to, value, data)
}

// BCBApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BCB contract.
type BCBApprovalIterator struct {
	Event *BCBApproval // Event containing the contract specifics and raw log

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
func (it *BCBApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BCBApproval)
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
		it.Event = new(BCBApproval)
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
func (it *BCBApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BCBApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BCBApproval represents a Approval event raised by the BCB contract.
type BCBApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BCB *BCBFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BCBApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BCB.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BCBApprovalIterator{contract: _BCB.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BCB *BCBFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BCBApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BCB.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BCBApproval)
				if err := _BCB.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BCB *BCBFilterer) ParseApproval(log types.Log) (*BCBApproval, error) {
	event := new(BCBApproval)
	if err := _BCB.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BCBOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BCB contract.
type BCBOwnershipTransferredIterator struct {
	Event *BCBOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BCBOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BCBOwnershipTransferred)
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
		it.Event = new(BCBOwnershipTransferred)
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
func (it *BCBOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BCBOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BCBOwnershipTransferred represents a OwnershipTransferred event raised by the BCB contract.
type BCBOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BCB *BCBFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BCBOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BCB.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BCBOwnershipTransferredIterator{contract: _BCB.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BCB *BCBFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BCBOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BCB.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BCBOwnershipTransferred)
				if err := _BCB.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BCB *BCBFilterer) ParseOwnershipTransferred(log types.Log) (*BCBOwnershipTransferred, error) {
	event := new(BCBOwnershipTransferred)
	if err := _BCB.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BCBTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BCB contract.
type BCBTransferIterator struct {
	Event *BCBTransfer // Event containing the contract specifics and raw log

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
func (it *BCBTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BCBTransfer)
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
		it.Event = new(BCBTransfer)
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
func (it *BCBTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BCBTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BCBTransfer represents a Transfer event raised by the BCB contract.
type BCBTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BCB *BCBFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BCBTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BCB.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BCBTransferIterator{contract: _BCB.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BCB *BCBFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BCBTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BCB.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BCBTransfer)
				if err := _BCB.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BCB *BCBFilterer) ParseTransfer(log types.Log) (*BCBTransfer, error) {
	event := new(BCBTransfer)
	if err := _BCB.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BCBTransferWithDataIterator is returned from FilterTransferWithData and is used to iterate over the raw logs and unpacked data for TransferWithData events raised by the BCB contract.
type BCBTransferWithDataIterator struct {
	Event *BCBTransferWithData // Event containing the contract specifics and raw log

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
func (it *BCBTransferWithDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BCBTransferWithData)
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
		it.Event = new(BCBTransferWithData)
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
func (it *BCBTransferWithDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BCBTransferWithDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BCBTransferWithData represents a TransferWithData event raised by the BCB contract.
type BCBTransferWithData struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferWithData is a free log retrieval operation binding the contract event 0xe68ca1ec8e8e022357047aae1f96036cbb808c6dc2bbbfbd3bde507ab21098c4.
//
// Solidity: event TransferWithData(address indexed from, address indexed to, uint256 value, bytes data)
func (_BCB *BCBFilterer) FilterTransferWithData(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BCBTransferWithDataIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BCB.contract.FilterLogs(opts, "TransferWithData", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BCBTransferWithDataIterator{contract: _BCB.contract, event: "TransferWithData", logs: logs, sub: sub}, nil
}

// WatchTransferWithData is a free log subscription operation binding the contract event 0xe68ca1ec8e8e022357047aae1f96036cbb808c6dc2bbbfbd3bde507ab21098c4.
//
// Solidity: event TransferWithData(address indexed from, address indexed to, uint256 value, bytes data)
func (_BCB *BCBFilterer) WatchTransferWithData(opts *bind.WatchOpts, sink chan<- *BCBTransferWithData, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BCB.contract.WatchLogs(opts, "TransferWithData", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BCBTransferWithData)
				if err := _BCB.contract.UnpackLog(event, "TransferWithData", log); err != nil {
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

// ParseTransferWithData is a log parse operation binding the contract event 0xe68ca1ec8e8e022357047aae1f96036cbb808c6dc2bbbfbd3bde507ab21098c4.
//
// Solidity: event TransferWithData(address indexed from, address indexed to, uint256 value, bytes data)
func (_BCB *BCBFilterer) ParseTransferWithData(log types.Log) (*BCBTransferWithData, error) {
	event := new(BCBTransferWithData)
	if err := _BCB.contract.UnpackLog(event, "TransferWithData", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Bin is the compiled bytecode used for deploying new contracts.
var ERC20Bin = "0x608060405234801561001057600080fd5b506106e2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a0823114610149578063a457c2d71461016f578063a9059cbb1461019b578063dd62ed3e146101c757610088565b8063095ea7b31461008d57806318160ddd146100cd57806323b872dd146100e7578063395093511461011d575b600080fd5b6100b9600480360360408110156100a357600080fd5b506001600160a01b0381351690602001356101f5565b604080519115158252519081900360200190f35b6100d561020b565b60408051918252519081900360200190f35b6100b9600480360360608110156100fd57600080fd5b506001600160a01b03813581169160208101359091169060400135610211565b6100b96004803603604081101561013357600080fd5b506001600160a01b038135169060200135610268565b6100d56004803603602081101561015f57600080fd5b50356001600160a01b03166102a4565b6100b96004803603604081101561018557600080fd5b506001600160a01b0381351690602001356102bf565b6100b9600480360360408110156101b157600080fd5b506001600160a01b0381351690602001356102fb565b6100d5600480360360408110156101dd57600080fd5b506001600160a01b0381358116916020013516610308565b6000610202338484610333565b50600192915050565b60025490565b600061021e84848461041f565b6001600160a01b03841660009081526001602090815260408083203380855292529091205461025e918691610259908663ffffffff61056116565b610333565b5060019392505050565b3360008181526001602090815260408083206001600160a01b03871684529091528120549091610202918590610259908663ffffffff6105be16565b6001600160a01b031660009081526020819052604090205490565b3360008181526001602090815260408083206001600160a01b03871684529091528120549091610202918590610259908663ffffffff61056116565b600061020233848461041f565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166103785760405162461bcd60e51b815260040180806020018281038252602481526020018061068a6024913960400191505060405180910390fd5b6001600160a01b0382166103bd5760405162461bcd60e51b81526004018080602001828103825260228152602001806106436022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166104645760405162461bcd60e51b81526004018080602001828103825260258152602001806106656025913960400191505060405180910390fd5b6001600160a01b0382166104a95760405162461bcd60e51b81526004018080602001828103825260238152602001806106206023913960400191505060405180910390fd5b6001600160a01b0383166000908152602081905260409020546104d2908263ffffffff61056116565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610507908263ffffffff6105be16565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000828211156105b8576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015610618576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373a265627a7a72315820e0837422aa76026f487cc7aaf5c95836b4a47a19df7711d725845ac644d321f964736f6c634300050b0032"

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RaffleABI is the input ABI used to generate the binding from.
const RaffleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newDrawing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"potMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_charities\",\"type\":\"address[]\"}],\"name\":\"setCharities\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rounds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ticketCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pot\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ticket\",\"type\":\"uint256\"}],\"name\":\"getTicket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_player\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_player\",\"type\":\"address\"}],\"name\":\"createTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setTokenContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentPotLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_potLimit\",\"type\":\"uint256\"}],\"name\":\"setCurrentPotLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"charities\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"TicketPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ticketNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DrawingCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"potLimit\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RaffleFuncSigs maps the 4-byte function signature to its string representation.
var RaffleFuncSigs = map[string]string{
	"fb709d2e": "charities(uint256)",
	"a5763a1f": "createTicket(address)",
	"d129299c": "currentPotLimit()",
	"8a19c8bc": "currentRound()",
	"a426e4c8": "getTicket(uint256,uint256)",
	"8f32d59b": "isOwner()",
	"06fdde03": "name()",
	"1b61a8d0": "newDrawing()",
	"8da5cb5b": "owner()",
	"599b6efe": "potMax()",
	"715018a6": "renounceOwnership()",
	"8c65c81f": "rounds(uint256)",
	"7be126b0": "setCharities(address[])",
	"ed67f845": "setCurrentPotLimit(uint256)",
	"bbcd5bbe": "setTokenContract(address)",
	"fc0c546a": "token()",
	"f2fde38b": "transferOwnership(address)",
	"2e1a7d4d": "withdraw(uint256)",
}

// RaffleBin is the compiled bytecode used for deploying new contracts.
var RaffleBin = "0x60806040526000600255600560035560a060045534801561001f57600080fd5b50600080546001600160a01b03191633178082556040516001600160a01b039190911691907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a36040805180820190915260168082527f4275747465726375702035302f353020526166666c650000000000000000000060209092019182526100ad91600191610109565b5060408051606081018252600080825260208083018281526005848601908152600280548552600790935294909220925183546001600160a01b0319166001600160a01b039091161783559051600183015591519101556101a4565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014a57805160ff1916838001178555610177565b82800160010185558215610177579182015b8281111561017757825182559160200191906001019061015c565b50610183929150610187565b5090565b6101a191905b80821115610183576000815560010161018d565b90565b610ea7806101b36000396000f3fe6080604052600436106101095760003560e01c80638f32d59b11610095578063d129299c11610064578063d129299c14610417578063ed67f8451461042c578063f2fde38b14610456578063fb709d2e14610489578063fc0c546a146104b357610109565b80638f32d59b14610358578063a426e4c814610381578063a5763a1f146103b1578063bbcd5bbe146103e457610109565b8063715018a6116100dc578063715018a6146101fb5780637be126b0146102105780638a19c8bc146102c05780638c65c81f146102d55780638da5cb5b1461032757610109565b806306fdde031461010b5780631b61a8d0146101955780632e1a7d4d146101aa578063599b6efe146101d4575b005b34801561011757600080fd5b506101206104c8565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561015a578181015183820152602001610142565b50505050905090810190601f1680156101875780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101a157600080fd5b50610109610555565b3480156101b657600080fd5b50610109600480360360208110156101cd57600080fd5b50356108c5565b3480156101e057600080fd5b506101e961093d565b60408051918252519081900360200190f35b34801561020757600080fd5b50610109610943565b34801561021c57600080fd5b506101096004803603602081101561023357600080fd5b81019060208101813564010000000081111561024e57600080fd5b82018360208201111561026057600080fd5b8035906020019184602083028401116401000000008311171561028257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506109d4945050505050565b3480156102cc57600080fd5b506101e9610a2e565b3480156102e157600080fd5b506102ff600480360360208110156102f857600080fd5b5035610a34565b604080516001600160a01b039094168452602084019290925282820152519081900360600190f35b34801561033357600080fd5b5061033c610a5f565b604080516001600160a01b039092168252519081900360200190f35b34801561036457600080fd5b5061036d610a6f565b604080519115158252519081900360200190f35b34801561038d57600080fd5b5061033c600480360360408110156103a457600080fd5b5080359060200135610a80565b3480156103bd57600080fd5b506101e9600480360360208110156103d457600080fd5b50356001600160a01b0316610aaa565b3480156103f057600080fd5b506101096004803603602081101561040757600080fd5b50356001600160a01b0316610bbf565b34801561042357600080fd5b506101e9610c28565b34801561043857600080fd5b506101096004803603602081101561044f57600080fd5b5035610c2e565b34801561046257600080fd5b506101096004803603602081101561047957600080fd5b50356001600160a01b0316610c7a565b34801561049557600080fd5b5061033c600480360360208110156104ac57600080fd5b5035610ccd565b3480156104bf57600080fd5b5061033c610cf4565b60018054604080516020600284861615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561054d5780601f106105225761010080835404028352916020019161054d565b820191906000526020600020905b81548152906001019060200180831161053057829003601f168201915b505050505081565b61055d610a6f565b61059c576040805162461bcd60e51b81526020600482018190526024820152600080516020610e53833981519152604482015290519081900360640190fd5b600254600090815260076020526040812060010154600019430140816105be57fe5b6002546000818152600760208181526040808420969095068084526003870182528584205494845291815285546001600160a01b0319166001600160a01b039485161790955560055484516370a0823160e01b8152306004820152945191965090921693909284926370a0823192602480840193919291829003018186803b15801561064957600080fd5b505afa15801561065d573d6000803e3d6000fd5b505050506040513d602081101561067357600080fd5b5051600280546000908152600760209081526040808320888452600301825280832054815163a9059cbb60e01b81526001600160a01b0391821660048201529486046024860152905194955086169363a9059cbb93604480820194918390030190829087803b1580156106e557600080fd5b505af11580156106f9573d6000803e3d6000fd5b505050506040513d602081101561070f57600080fd5b50600090505b6006548110156107e557826001600160a01b031663a9059cbb6006838154811061073b57fe5b6000918252602090912001546006546001600160a01b0390911690600286048161076157fe5b046040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b1580156107b157600080fd5b505af11580156107c5573d6000803e3d6000fd5b505050506040513d60208110156107db57600080fd5b5050600101610715565b506002805460009081526007602090815260408083208784526003018252918290205482518781526001600160a01b03909116918101919091528082018690529183046060830152517f94d9474e9279e038c29a59d8ddc94349c518eef152a547838f6f0bd4bd83fcb89181900360800190a16004546003541061086d576005600355610877565b6003805460020290555b600280546001019081905560035460408051928352602083019190915280517f1237ad121dfe1e3abc3b4254161693eb4add51b458114b8c15502bea94a6b02b9281900390910190a1505050565b6108cd610a6f565b61090c576040805162461bcd60e51b81526020600482018190526024820152600080516020610e53833981519152604482015290519081900360640190fd5b604051339082156108fc029083906000818181858888f19350505050158015610939573d6000803e3d6000fd5b5050565b60045481565b61094b610a6f565b61098a576040805162461bcd60e51b81526020600482018190526024820152600080516020610e53833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6109dc610a6f565b610a1b576040805162461bcd60e51b81526020600482018190526024820152600080516020610e53833981519152604482015290519081900360640190fd5b8051610939906006906020840190610da3565b60025481565b6007602052600090815260409020805460018201546002909201546001600160a01b03909116919083565b6000546001600160a01b03165b90565b6000546001600160a01b0316331490565b6000918252600760209081526040808420928452600390920190529020546001600160a01b031690565b6000610ab4610a6f565b610af3576040805162461bcd60e51b81526020600482018190526024820152600080516020610e53833981519152604482015290519081900360640190fd5b6002805460009081526007602081815260408084206001808201548087526003909201845282862080546001600160a01b0319166001600160a01b038b169081179091558754875294845282862081018054820190558654865294829020909501805490940190935582518481529081019190915281517f785f2865353ae960010300f584fc6c577b757a853af69a99c769df3c8232d3de929181900390910190a1600354600280546000908152600760205260409020015410610bb957610bb9610555565b92915050565b610bc7610a6f565b610c06576040805162461bcd60e51b81526020600482018190526024820152600080516020610e53833981519152604482015290519081900360640190fd5b600580546001600160a01b0319166001600160a01b0392909216919091179055565b60035481565b610c36610a6f565b610c75576040805162461bcd60e51b81526020600482018190526024820152600080516020610e53833981519152604482015290519081900360640190fd5b600355565b610c82610a6f565b610cc1576040805162461bcd60e51b81526020600482018190526024820152600080516020610e53833981519152604482015290519081900360640190fd5b610cca81610d03565b50565b60068181548110610cda57fe5b6000918252602090912001546001600160a01b0316905081565b6005546001600160a01b031681565b6001600160a01b038116610d485760405162461bcd60e51b8152600401808060200182810382526026815260200180610e2d6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054828255906000526020600020908101928215610df8579160200282015b82811115610df857825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610dc3565b50610e04929150610e08565b5090565b610a6c91905b80821115610e045780546001600160a01b0319168155600101610e0e56fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a265627a7a7231582032557e961132a5e3d31bcc1c3c4c4c8976be6d3655f6b0ffe31269c9038a379664736f6c634300050b0032"

// DeployRaffle deploys a new Ethereum contract, binding an instance of Raffle to it.
func DeployRaffle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Raffle, error) {
	parsed, err := abi.JSON(strings.NewReader(RaffleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RaffleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Raffle{RaffleCaller: RaffleCaller{contract: contract}, RaffleTransactor: RaffleTransactor{contract: contract}, RaffleFilterer: RaffleFilterer{contract: contract}}, nil
}

// Raffle is an auto generated Go binding around an Ethereum contract.
type Raffle struct {
	RaffleCaller     // Read-only binding to the contract
	RaffleTransactor // Write-only binding to the contract
	RaffleFilterer   // Log filterer for contract events
}

// RaffleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RaffleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RaffleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RaffleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RaffleSession struct {
	Contract     *Raffle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RaffleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RaffleCallerSession struct {
	Contract *RaffleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RaffleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RaffleTransactorSession struct {
	Contract     *RaffleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RaffleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RaffleRaw struct {
	Contract *Raffle // Generic contract binding to access the raw methods on
}

// RaffleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RaffleCallerRaw struct {
	Contract *RaffleCaller // Generic read-only contract binding to access the raw methods on
}

// RaffleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RaffleTransactorRaw struct {
	Contract *RaffleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRaffle creates a new instance of Raffle, bound to a specific deployed contract.
func NewRaffle(address common.Address, backend bind.ContractBackend) (*Raffle, error) {
	contract, err := bindRaffle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Raffle{RaffleCaller: RaffleCaller{contract: contract}, RaffleTransactor: RaffleTransactor{contract: contract}, RaffleFilterer: RaffleFilterer{contract: contract}}, nil
}

// NewRaffleCaller creates a new read-only instance of Raffle, bound to a specific deployed contract.
func NewRaffleCaller(address common.Address, caller bind.ContractCaller) (*RaffleCaller, error) {
	contract, err := bindRaffle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RaffleCaller{contract: contract}, nil
}

// NewRaffleTransactor creates a new write-only instance of Raffle, bound to a specific deployed contract.
func NewRaffleTransactor(address common.Address, transactor bind.ContractTransactor) (*RaffleTransactor, error) {
	contract, err := bindRaffle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RaffleTransactor{contract: contract}, nil
}

// NewRaffleFilterer creates a new log filterer instance of Raffle, bound to a specific deployed contract.
func NewRaffleFilterer(address common.Address, filterer bind.ContractFilterer) (*RaffleFilterer, error) {
	contract, err := bindRaffle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RaffleFilterer{contract: contract}, nil
}

// bindRaffle binds a generic wrapper to an already deployed contract.
func bindRaffle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RaffleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Raffle *RaffleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Raffle.Contract.RaffleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Raffle *RaffleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Raffle.Contract.RaffleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Raffle *RaffleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Raffle.Contract.RaffleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Raffle *RaffleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Raffle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Raffle *RaffleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Raffle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Raffle *RaffleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Raffle.Contract.contract.Transact(opts, method, params...)
}

// Charities is a free data retrieval call binding the contract method 0xfb709d2e.
//
// Solidity: function charities(uint256 ) constant returns(address)
func (_Raffle *RaffleCaller) Charities(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Raffle.contract.Call(opts, out, "charities", arg0)
	return *ret0, err
}

// Charities is a free data retrieval call binding the contract method 0xfb709d2e.
//
// Solidity: function charities(uint256 ) constant returns(address)
func (_Raffle *RaffleSession) Charities(arg0 *big.Int) (common.Address, error) {
	return _Raffle.Contract.Charities(&_Raffle.CallOpts, arg0)
}

// Charities is a free data retrieval call binding the contract method 0xfb709d2e.
//
// Solidity: function charities(uint256 ) constant returns(address)
func (_Raffle *RaffleCallerSession) Charities(arg0 *big.Int) (common.Address, error) {
	return _Raffle.Contract.Charities(&_Raffle.CallOpts, arg0)
}

// CurrentPotLimit is a free data retrieval call binding the contract method 0xd129299c.
//
// Solidity: function currentPotLimit() constant returns(uint256)
func (_Raffle *RaffleCaller) CurrentPotLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Raffle.contract.Call(opts, out, "currentPotLimit")
	return *ret0, err
}

// CurrentPotLimit is a free data retrieval call binding the contract method 0xd129299c.
//
// Solidity: function currentPotLimit() constant returns(uint256)
func (_Raffle *RaffleSession) CurrentPotLimit() (*big.Int, error) {
	return _Raffle.Contract.CurrentPotLimit(&_Raffle.CallOpts)
}

// CurrentPotLimit is a free data retrieval call binding the contract method 0xd129299c.
//
// Solidity: function currentPotLimit() constant returns(uint256)
func (_Raffle *RaffleCallerSession) CurrentPotLimit() (*big.Int, error) {
	return _Raffle.Contract.CurrentPotLimit(&_Raffle.CallOpts)
}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() constant returns(uint256)
func (_Raffle *RaffleCaller) CurrentRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Raffle.contract.Call(opts, out, "currentRound")
	return *ret0, err
}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() constant returns(uint256)
func (_Raffle *RaffleSession) CurrentRound() (*big.Int, error) {
	return _Raffle.Contract.CurrentRound(&_Raffle.CallOpts)
}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() constant returns(uint256)
func (_Raffle *RaffleCallerSession) CurrentRound() (*big.Int, error) {
	return _Raffle.Contract.CurrentRound(&_Raffle.CallOpts)
}

// GetTicket is a free data retrieval call binding the contract method 0xa426e4c8.
//
// Solidity: function getTicket(uint256 _round, uint256 _ticket) constant returns(address _player)
func (_Raffle *RaffleCaller) GetTicket(opts *bind.CallOpts, _round *big.Int, _ticket *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Raffle.contract.Call(opts, out, "getTicket", _round, _ticket)
	return *ret0, err
}

// GetTicket is a free data retrieval call binding the contract method 0xa426e4c8.
//
// Solidity: function getTicket(uint256 _round, uint256 _ticket) constant returns(address _player)
func (_Raffle *RaffleSession) GetTicket(_round *big.Int, _ticket *big.Int) (common.Address, error) {
	return _Raffle.Contract.GetTicket(&_Raffle.CallOpts, _round, _ticket)
}

// GetTicket is a free data retrieval call binding the contract method 0xa426e4c8.
//
// Solidity: function getTicket(uint256 _round, uint256 _ticket) constant returns(address _player)
func (_Raffle *RaffleCallerSession) GetTicket(_round *big.Int, _ticket *big.Int) (common.Address, error) {
	return _Raffle.Contract.GetTicket(&_Raffle.CallOpts, _round, _ticket)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Raffle *RaffleCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Raffle.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Raffle *RaffleSession) IsOwner() (bool, error) {
	return _Raffle.Contract.IsOwner(&_Raffle.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Raffle *RaffleCallerSession) IsOwner() (bool, error) {
	return _Raffle.Contract.IsOwner(&_Raffle.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Raffle *RaffleCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Raffle.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Raffle *RaffleSession) Name() (string, error) {
	return _Raffle.Contract.Name(&_Raffle.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Raffle *RaffleCallerSession) Name() (string, error) {
	return _Raffle.Contract.Name(&_Raffle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Raffle *RaffleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Raffle.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Raffle *RaffleSession) Owner() (common.Address, error) {
	return _Raffle.Contract.Owner(&_Raffle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Raffle *RaffleCallerSession) Owner() (common.Address, error) {
	return _Raffle.Contract.Owner(&_Raffle.CallOpts)
}

// PotMax is a free data retrieval call binding the contract method 0x599b6efe.
//
// Solidity: function potMax() constant returns(uint256)
func (_Raffle *RaffleCaller) PotMax(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Raffle.contract.Call(opts, out, "potMax")
	return *ret0, err
}

// PotMax is a free data retrieval call binding the contract method 0x599b6efe.
//
// Solidity: function potMax() constant returns(uint256)
func (_Raffle *RaffleSession) PotMax() (*big.Int, error) {
	return _Raffle.Contract.PotMax(&_Raffle.CallOpts)
}

// PotMax is a free data retrieval call binding the contract method 0x599b6efe.
//
// Solidity: function potMax() constant returns(uint256)
func (_Raffle *RaffleCallerSession) PotMax() (*big.Int, error) {
	return _Raffle.Contract.PotMax(&_Raffle.CallOpts)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) constant returns(address winner, uint256 ticketCount, uint256 pot)
func (_Raffle *RaffleCaller) Rounds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Winner      common.Address
	TicketCount *big.Int
	Pot         *big.Int
}, error) {
	ret := new(struct {
		Winner      common.Address
		TicketCount *big.Int
		Pot         *big.Int
	})
	out := ret
	err := _Raffle.contract.Call(opts, out, "rounds", arg0)
	return *ret, err
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) constant returns(address winner, uint256 ticketCount, uint256 pot)
func (_Raffle *RaffleSession) Rounds(arg0 *big.Int) (struct {
	Winner      common.Address
	TicketCount *big.Int
	Pot         *big.Int
}, error) {
	return _Raffle.Contract.Rounds(&_Raffle.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) constant returns(address winner, uint256 ticketCount, uint256 pot)
func (_Raffle *RaffleCallerSession) Rounds(arg0 *big.Int) (struct {
	Winner      common.Address
	TicketCount *big.Int
	Pot         *big.Int
}, error) {
	return _Raffle.Contract.Rounds(&_Raffle.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Raffle *RaffleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Raffle.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Raffle *RaffleSession) Token() (common.Address, error) {
	return _Raffle.Contract.Token(&_Raffle.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Raffle *RaffleCallerSession) Token() (common.Address, error) {
	return _Raffle.Contract.Token(&_Raffle.CallOpts)
}

// CreateTicket is a paid mutator transaction binding the contract method 0xa5763a1f.
//
// Solidity: function createTicket(address _player) returns(uint256)
func (_Raffle *RaffleTransactor) CreateTicket(opts *bind.TransactOpts, _player common.Address) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "createTicket", _player)
}

// CreateTicket is a paid mutator transaction binding the contract method 0xa5763a1f.
//
// Solidity: function createTicket(address _player) returns(uint256)
func (_Raffle *RaffleSession) CreateTicket(_player common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.CreateTicket(&_Raffle.TransactOpts, _player)
}

// CreateTicket is a paid mutator transaction binding the contract method 0xa5763a1f.
//
// Solidity: function createTicket(address _player) returns(uint256)
func (_Raffle *RaffleTransactorSession) CreateTicket(_player common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.CreateTicket(&_Raffle.TransactOpts, _player)
}

// NewDrawing is a paid mutator transaction binding the contract method 0x1b61a8d0.
//
// Solidity: function newDrawing() returns()
func (_Raffle *RaffleTransactor) NewDrawing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "newDrawing")
}

// NewDrawing is a paid mutator transaction binding the contract method 0x1b61a8d0.
//
// Solidity: function newDrawing() returns()
func (_Raffle *RaffleSession) NewDrawing() (*types.Transaction, error) {
	return _Raffle.Contract.NewDrawing(&_Raffle.TransactOpts)
}

// NewDrawing is a paid mutator transaction binding the contract method 0x1b61a8d0.
//
// Solidity: function newDrawing() returns()
func (_Raffle *RaffleTransactorSession) NewDrawing() (*types.Transaction, error) {
	return _Raffle.Contract.NewDrawing(&_Raffle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Raffle *RaffleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Raffle *RaffleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Raffle.Contract.RenounceOwnership(&_Raffle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Raffle *RaffleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Raffle.Contract.RenounceOwnership(&_Raffle.TransactOpts)
}

// SetCharities is a paid mutator transaction binding the contract method 0x7be126b0.
//
// Solidity: function setCharities(address[] _charities) returns()
func (_Raffle *RaffleTransactor) SetCharities(opts *bind.TransactOpts, _charities []common.Address) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "setCharities", _charities)
}

// SetCharities is a paid mutator transaction binding the contract method 0x7be126b0.
//
// Solidity: function setCharities(address[] _charities) returns()
func (_Raffle *RaffleSession) SetCharities(_charities []common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.SetCharities(&_Raffle.TransactOpts, _charities)
}

// SetCharities is a paid mutator transaction binding the contract method 0x7be126b0.
//
// Solidity: function setCharities(address[] _charities) returns()
func (_Raffle *RaffleTransactorSession) SetCharities(_charities []common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.SetCharities(&_Raffle.TransactOpts, _charities)
}

// SetCurrentPotLimit is a paid mutator transaction binding the contract method 0xed67f845.
//
// Solidity: function setCurrentPotLimit(uint256 _potLimit) returns()
func (_Raffle *RaffleTransactor) SetCurrentPotLimit(opts *bind.TransactOpts, _potLimit *big.Int) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "setCurrentPotLimit", _potLimit)
}

// SetCurrentPotLimit is a paid mutator transaction binding the contract method 0xed67f845.
//
// Solidity: function setCurrentPotLimit(uint256 _potLimit) returns()
func (_Raffle *RaffleSession) SetCurrentPotLimit(_potLimit *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.SetCurrentPotLimit(&_Raffle.TransactOpts, _potLimit)
}

// SetCurrentPotLimit is a paid mutator transaction binding the contract method 0xed67f845.
//
// Solidity: function setCurrentPotLimit(uint256 _potLimit) returns()
func (_Raffle *RaffleTransactorSession) SetCurrentPotLimit(_potLimit *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.SetCurrentPotLimit(&_Raffle.TransactOpts, _potLimit)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address _contract) returns()
func (_Raffle *RaffleTransactor) SetTokenContract(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "setTokenContract", _contract)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address _contract) returns()
func (_Raffle *RaffleSession) SetTokenContract(_contract common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.SetTokenContract(&_Raffle.TransactOpts, _contract)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address _contract) returns()
func (_Raffle *RaffleTransactorSession) SetTokenContract(_contract common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.SetTokenContract(&_Raffle.TransactOpts, _contract)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Raffle *RaffleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Raffle *RaffleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.TransferOwnership(&_Raffle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Raffle *RaffleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.TransferOwnership(&_Raffle.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Raffle *RaffleTransactor) Withdraw(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "withdraw", value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Raffle *RaffleSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.Withdraw(&_Raffle.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Raffle *RaffleTransactorSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.Withdraw(&_Raffle.TransactOpts, value)
}

// RaffleDrawingCompletedIterator is returned from FilterDrawingCompleted and is used to iterate over the raw logs and unpacked data for DrawingCompleted events raised by the Raffle contract.
type RaffleDrawingCompletedIterator struct {
	Event *RaffleDrawingCompleted // Event containing the contract specifics and raw log

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
func (it *RaffleDrawingCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleDrawingCompleted)
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
		it.Event = new(RaffleDrawingCompleted)
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
func (it *RaffleDrawingCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleDrawingCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleDrawingCompleted represents a DrawingCompleted event raised by the Raffle contract.
type RaffleDrawingCompleted struct {
	Round        *big.Int
	Winner       common.Address
	TicketNumber *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDrawingCompleted is a free log retrieval operation binding the contract event 0x94d9474e9279e038c29a59d8ddc94349c518eef152a547838f6f0bd4bd83fcb8.
//
// Solidity: event DrawingCompleted(uint256 round, address winner, uint256 ticketNumber, uint256 amount)
func (_Raffle *RaffleFilterer) FilterDrawingCompleted(opts *bind.FilterOpts) (*RaffleDrawingCompletedIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "DrawingCompleted")
	if err != nil {
		return nil, err
	}
	return &RaffleDrawingCompletedIterator{contract: _Raffle.contract, event: "DrawingCompleted", logs: logs, sub: sub}, nil
}

// WatchDrawingCompleted is a free log subscription operation binding the contract event 0x94d9474e9279e038c29a59d8ddc94349c518eef152a547838f6f0bd4bd83fcb8.
//
// Solidity: event DrawingCompleted(uint256 round, address winner, uint256 ticketNumber, uint256 amount)
func (_Raffle *RaffleFilterer) WatchDrawingCompleted(opts *bind.WatchOpts, sink chan<- *RaffleDrawingCompleted) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "DrawingCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleDrawingCompleted)
				if err := _Raffle.contract.UnpackLog(event, "DrawingCompleted", log); err != nil {
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

// ParseDrawingCompleted is a log parse operation binding the contract event 0x94d9474e9279e038c29a59d8ddc94349c518eef152a547838f6f0bd4bd83fcb8.
//
// Solidity: event DrawingCompleted(uint256 round, address winner, uint256 ticketNumber, uint256 amount)
func (_Raffle *RaffleFilterer) ParseDrawingCompleted(log types.Log) (*RaffleDrawingCompleted, error) {
	event := new(RaffleDrawingCompleted)
	if err := _Raffle.contract.UnpackLog(event, "DrawingCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RaffleNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the Raffle contract.
type RaffleNewRoundIterator struct {
	Event *RaffleNewRound // Event containing the contract specifics and raw log

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
func (it *RaffleNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleNewRound)
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
		it.Event = new(RaffleNewRound)
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
func (it *RaffleNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleNewRound represents a NewRound event raised by the Raffle contract.
type RaffleNewRound struct {
	Round    *big.Int
	PotLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x1237ad121dfe1e3abc3b4254161693eb4add51b458114b8c15502bea94a6b02b.
//
// Solidity: event NewRound(uint256 round, uint256 potLimit)
func (_Raffle *RaffleFilterer) FilterNewRound(opts *bind.FilterOpts) (*RaffleNewRoundIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "NewRound")
	if err != nil {
		return nil, err
	}
	return &RaffleNewRoundIterator{contract: _Raffle.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x1237ad121dfe1e3abc3b4254161693eb4add51b458114b8c15502bea94a6b02b.
//
// Solidity: event NewRound(uint256 round, uint256 potLimit)
func (_Raffle *RaffleFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *RaffleNewRound) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "NewRound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleNewRound)
				if err := _Raffle.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0x1237ad121dfe1e3abc3b4254161693eb4add51b458114b8c15502bea94a6b02b.
//
// Solidity: event NewRound(uint256 round, uint256 potLimit)
func (_Raffle *RaffleFilterer) ParseNewRound(log types.Log) (*RaffleNewRound, error) {
	event := new(RaffleNewRound)
	if err := _Raffle.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RaffleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Raffle contract.
type RaffleOwnershipTransferredIterator struct {
	Event *RaffleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RaffleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleOwnershipTransferred)
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
		it.Event = new(RaffleOwnershipTransferred)
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
func (it *RaffleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleOwnershipTransferred represents a OwnershipTransferred event raised by the Raffle contract.
type RaffleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Raffle *RaffleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RaffleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RaffleOwnershipTransferredIterator{contract: _Raffle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Raffle *RaffleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RaffleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleOwnershipTransferred)
				if err := _Raffle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Raffle *RaffleFilterer) ParseOwnershipTransferred(log types.Log) (*RaffleOwnershipTransferred, error) {
	event := new(RaffleOwnershipTransferred)
	if err := _Raffle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RaffleTicketPurchasedIterator is returned from FilterTicketPurchased and is used to iterate over the raw logs and unpacked data for TicketPurchased events raised by the Raffle contract.
type RaffleTicketPurchasedIterator struct {
	Event *RaffleTicketPurchased // Event containing the contract specifics and raw log

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
func (it *RaffleTicketPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleTicketPurchased)
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
		it.Event = new(RaffleTicketPurchased)
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
func (it *RaffleTicketPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleTicketPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleTicketPurchased represents a TicketPurchased event raised by the Raffle contract.
type RaffleTicketPurchased struct {
	Id     *big.Int
	Player common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTicketPurchased is a free log retrieval operation binding the contract event 0x785f2865353ae960010300f584fc6c577b757a853af69a99c769df3c8232d3de.
//
// Solidity: event TicketPurchased(uint256 id, address player)
func (_Raffle *RaffleFilterer) FilterTicketPurchased(opts *bind.FilterOpts) (*RaffleTicketPurchasedIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "TicketPurchased")
	if err != nil {
		return nil, err
	}
	return &RaffleTicketPurchasedIterator{contract: _Raffle.contract, event: "TicketPurchased", logs: logs, sub: sub}, nil
}

// WatchTicketPurchased is a free log subscription operation binding the contract event 0x785f2865353ae960010300f584fc6c577b757a853af69a99c769df3c8232d3de.
//
// Solidity: event TicketPurchased(uint256 id, address player)
func (_Raffle *RaffleFilterer) WatchTicketPurchased(opts *bind.WatchOpts, sink chan<- *RaffleTicketPurchased) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "TicketPurchased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleTicketPurchased)
				if err := _Raffle.contract.UnpackLog(event, "TicketPurchased", log); err != nil {
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

// ParseTicketPurchased is a log parse operation binding the contract event 0x785f2865353ae960010300f584fc6c577b757a853af69a99c769df3c8232d3de.
//
// Solidity: event TicketPurchased(uint256 id, address player)
func (_Raffle *RaffleFilterer) ParseTicketPurchased(log types.Log) (*RaffleTicketPurchased, error) {
	event := new(RaffleTicketPurchased)
	if err := _Raffle.contract.UnpackLog(event, "TicketPurchased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209f9b9b406ef3455a9d088fbc5c2f718921f778512010e6949db518bae49fd8f264736f6c634300050b0032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
