// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"contentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"assetId\",\"type\":\"uint64\"}],\"name\":\"CreateAsset\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"assetById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"contentId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"contentAssetMapping\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_contentId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"createAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_contentId\",\"type\":\"uint64\"}],\"name\":\"getAssetId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_assetId\",\"type\":\"uint64\"}],\"name\":\"isQualified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paymentRecord\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_assetId\",\"type\":\"uint64\"}],\"name\":\"purchaseAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610c47806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80639e281a98116100665780639e281a9814610169578063bd2aaf7f1461017e578063e48a8fcf14610191578063fc41585114610210578063fe12e6a31461023357600080fd5b80631359b0a3146100985780631eb1a724146100d6578063223dc2c71461012257806389432e5e14610135575b600080fd5b6100c36100a6366004610a0f565b600360209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b61010a6100e4366004610a59565b60016020908152600092835260408084209091529082529020546001600160401b031681565b6040516001600160401b0390911681526020016100cd565b61010a610130366004610a59565b610246565b61010a610143366004610a83565b60026020908152600092835260408084209091529082529020546001600160401b031681565b61017c610177366004610a9f565b61027c565b005b61017c61018c366004610ac9565b610332565b6101dc61019f366004610ac9565b60006020819052908152604090208054600182015460028301546003909301546001600160a01b039283169391909216916001600160401b031684565b604080516001600160a01b039586168152949093166020850152918301526001600160401b031660608201526080016100cd565b61022361021e366004610a59565b6104ce565b60405190151581526020016100cd565b61017c610241366004610aeb565b610517565b6001600160a01b03821660009081526001602090815260408083206001600160401b038086168552925290912054165b92915050565b3360009081526003602090815260408083206001600160a01b03861684529091529020548111156102e15760405162461bcd60e51b815260206004820152600a6024820152696e6f2062616c616e636560b01b60448201526064015b60405180910390fd5b3360009081526003602090815260408083206001600160a01b038616845290915281208054839290610314908490610b3d565b9091555061032e90506001600160a01b03831633836106eb565b5050565b6001600160401b0381811660009081526020818152604091829020825160808101845281546001600160a01b0390811682526001830154169281018390526002820154818501819052600390920154909416606085015291516370a0823160e01b81523360048201526370a0823190602401602060405180830381865afa1580156103c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e59190610b50565b116104275760405162461bcd60e51b81526020600482015260126024820152710c4c2d8c2dcc6ca40dcdee840cadcdeeaced60731b60448201526064016102d8565b61044f3330836040015184602001516001600160a01b0316610753909392919063ffffffff16565b60408082015182516001600160a01b0390811660009081526003602090815284822081870151909316825291909152918220805491929091610492908490610b69565b9091555050506001600160401b031660009081526002602090815260408083203384529091529020805467ffffffffffffffff19166001179055565b6001600160401b0380821660009081526002602090815260408083206001600160a01b038716845290915281205490911660010361050e57506001610276565b50600092915050565b610525600480546001019055565b6001600160401b038061053760045490565b11156105855760405162461bcd60e51b815260206004820152601c60248201527f76616c756520646f65736e27742066697420696e20363420626974730000000060448201526064016102d8565b3360009081526001602090815260408083206001600160401b03808916855292529091205416156105e95760405162461bcd60e51b815260206004820152600e60248201526d616c72656164792063726561746560901b60448201526064016102d8565b60006105f460045490565b60408051608081018252338082526001600160a01b0388811660208085019182528486018a81526001600160401b038d811660608089018281528b841660008181528088528c81209b518c54908b166001600160a01b0319918216178d5598516001808e01805492909c1691909a1617909955945160028b015551600390990180549990931667ffffffffffffffff19998a1617909255868652938352878520848652835293879020805490961681179095558551938452830152928101919091529192507f269cf5c71029c03bbfc885db18cdb5e3044db3d4ebe27d52df30f4fa490eb280910160405180910390a15050505050565b6040516001600160a01b03831660248201526044810182905261074e90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610791565b505050565b6040516001600160a01b038085166024830152831660448201526064810182905261078b9085906323b872dd60e01b90608401610717565b50505050565b60006107e6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166108639092919063ffffffff16565b80519091501561074e57808060200190518101906108049190610b7c565b61074e5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016102d8565b6060610872848460008561087a565b949350505050565b6060824710156108db5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016102d8565b600080866001600160a01b031685876040516108f79190610bc2565b60006040518083038185875af1925050503d8060008114610934576040519150601f19603f3d011682016040523d82523d6000602084013e610939565b606091505b509150915061094a87838387610955565b979650505050505050565b606083156109c45782516000036109bd576001600160a01b0385163b6109bd5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102d8565b5081610872565b61087283838151156109d95781518083602001fd5b8060405162461bcd60e51b81526004016102d89190610bde565b80356001600160a01b0381168114610a0a57600080fd5b919050565b60008060408385031215610a2257600080fd5b610a2b836109f3565b9150610a39602084016109f3565b90509250929050565b80356001600160401b0381168114610a0a57600080fd5b60008060408385031215610a6c57600080fd5b610a75836109f3565b9150610a3960208401610a42565b60008060408385031215610a9657600080fd5b610a2b83610a42565b60008060408385031215610ab257600080fd5b610abb836109f3565b946020939093013593505050565b600060208284031215610adb57600080fd5b610ae482610a42565b9392505050565b600080600060608486031215610b0057600080fd5b610b0984610a42565b9250610b17602085016109f3565b9150604084013590509250925092565b634e487b7160e01b600052601160045260246000fd5b8181038181111561027657610276610b27565b600060208284031215610b6257600080fd5b5051919050565b8082018082111561027657610276610b27565b600060208284031215610b8e57600080fd5b81518015158114610ae457600080fd5b60005b83811015610bb9578181015183820152602001610ba1565b50506000910152565b60008251610bd4818460208701610b9e565b9190910192915050565b6020815260008251806020840152610bfd816040850160208701610b9e565b601f01601f1916919091016040019291505056fea2646970667358221220c4b383e59c80733cefdffd67c4bae10348968367c171d430c9ff526d83f74fd364736f6c63430008110033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// AssetById is a free data retrieval call binding the contract method 0xe48a8fcf.
//
// Solidity: function assetById(uint64 ) view returns(address author, address paymentToken, uint256 amount, uint64 contentId)
func (_Contracts *ContractsCaller) AssetById(opts *bind.CallOpts, arg0 uint64) (struct {
	Author       common.Address
	PaymentToken common.Address
	Amount       *big.Int
	ContentId    uint64
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "assetById", arg0)

	outstruct := new(struct {
		Author       common.Address
		PaymentToken common.Address
		Amount       *big.Int
		ContentId    uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Author = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PaymentToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ContentId = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// AssetById is a free data retrieval call binding the contract method 0xe48a8fcf.
//
// Solidity: function assetById(uint64 ) view returns(address author, address paymentToken, uint256 amount, uint64 contentId)
func (_Contracts *ContractsSession) AssetById(arg0 uint64) (struct {
	Author       common.Address
	PaymentToken common.Address
	Amount       *big.Int
	ContentId    uint64
}, error) {
	return _Contracts.Contract.AssetById(&_Contracts.CallOpts, arg0)
}

// AssetById is a free data retrieval call binding the contract method 0xe48a8fcf.
//
// Solidity: function assetById(uint64 ) view returns(address author, address paymentToken, uint256 amount, uint64 contentId)
func (_Contracts *ContractsCallerSession) AssetById(arg0 uint64) (struct {
	Author       common.Address
	PaymentToken common.Address
	Amount       *big.Int
	ContentId    uint64
}, error) {
	return _Contracts.Contract.AssetById(&_Contracts.CallOpts, arg0)
}

// AuthorBalance is a free data retrieval call binding the contract method 0x1359b0a3.
//
// Solidity: function authorBalance(address , address ) view returns(uint256)
func (_Contracts *ContractsCaller) AuthorBalance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "authorBalance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuthorBalance is a free data retrieval call binding the contract method 0x1359b0a3.
//
// Solidity: function authorBalance(address , address ) view returns(uint256)
func (_Contracts *ContractsSession) AuthorBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contracts.Contract.AuthorBalance(&_Contracts.CallOpts, arg0, arg1)
}

// AuthorBalance is a free data retrieval call binding the contract method 0x1359b0a3.
//
// Solidity: function authorBalance(address , address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) AuthorBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contracts.Contract.AuthorBalance(&_Contracts.CallOpts, arg0, arg1)
}

// ContentAssetMapping is a free data retrieval call binding the contract method 0x1eb1a724.
//
// Solidity: function contentAssetMapping(address , uint64 ) view returns(uint64)
func (_Contracts *ContractsCaller) ContentAssetMapping(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "contentAssetMapping", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ContentAssetMapping is a free data retrieval call binding the contract method 0x1eb1a724.
//
// Solidity: function contentAssetMapping(address , uint64 ) view returns(uint64)
func (_Contracts *ContractsSession) ContentAssetMapping(arg0 common.Address, arg1 uint64) (uint64, error) {
	return _Contracts.Contract.ContentAssetMapping(&_Contracts.CallOpts, arg0, arg1)
}

// ContentAssetMapping is a free data retrieval call binding the contract method 0x1eb1a724.
//
// Solidity: function contentAssetMapping(address , uint64 ) view returns(uint64)
func (_Contracts *ContractsCallerSession) ContentAssetMapping(arg0 common.Address, arg1 uint64) (uint64, error) {
	return _Contracts.Contract.ContentAssetMapping(&_Contracts.CallOpts, arg0, arg1)
}

// GetAssetId is a free data retrieval call binding the contract method 0x223dc2c7.
//
// Solidity: function getAssetId(address _creator, uint64 _contentId) view returns(uint64)
func (_Contracts *ContractsCaller) GetAssetId(opts *bind.CallOpts, _creator common.Address, _contentId uint64) (uint64, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAssetId", _creator, _contentId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAssetId is a free data retrieval call binding the contract method 0x223dc2c7.
//
// Solidity: function getAssetId(address _creator, uint64 _contentId) view returns(uint64)
func (_Contracts *ContractsSession) GetAssetId(_creator common.Address, _contentId uint64) (uint64, error) {
	return _Contracts.Contract.GetAssetId(&_Contracts.CallOpts, _creator, _contentId)
}

// GetAssetId is a free data retrieval call binding the contract method 0x223dc2c7.
//
// Solidity: function getAssetId(address _creator, uint64 _contentId) view returns(uint64)
func (_Contracts *ContractsCallerSession) GetAssetId(_creator common.Address, _contentId uint64) (uint64, error) {
	return _Contracts.Contract.GetAssetId(&_Contracts.CallOpts, _creator, _contentId)
}

// IsQualified is a free data retrieval call binding the contract method 0xfc415851.
//
// Solidity: function isQualified(address account, uint64 _assetId) view returns(bool)
func (_Contracts *ContractsCaller) IsQualified(opts *bind.CallOpts, account common.Address, _assetId uint64) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isQualified", account, _assetId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsQualified is a free data retrieval call binding the contract method 0xfc415851.
//
// Solidity: function isQualified(address account, uint64 _assetId) view returns(bool)
func (_Contracts *ContractsSession) IsQualified(account common.Address, _assetId uint64) (bool, error) {
	return _Contracts.Contract.IsQualified(&_Contracts.CallOpts, account, _assetId)
}

// IsQualified is a free data retrieval call binding the contract method 0xfc415851.
//
// Solidity: function isQualified(address account, uint64 _assetId) view returns(bool)
func (_Contracts *ContractsCallerSession) IsQualified(account common.Address, _assetId uint64) (bool, error) {
	return _Contracts.Contract.IsQualified(&_Contracts.CallOpts, account, _assetId)
}

// PaymentRecord is a free data retrieval call binding the contract method 0x89432e5e.
//
// Solidity: function paymentRecord(uint64 , address ) view returns(uint64)
func (_Contracts *ContractsCaller) PaymentRecord(opts *bind.CallOpts, arg0 uint64, arg1 common.Address) (uint64, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "paymentRecord", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PaymentRecord is a free data retrieval call binding the contract method 0x89432e5e.
//
// Solidity: function paymentRecord(uint64 , address ) view returns(uint64)
func (_Contracts *ContractsSession) PaymentRecord(arg0 uint64, arg1 common.Address) (uint64, error) {
	return _Contracts.Contract.PaymentRecord(&_Contracts.CallOpts, arg0, arg1)
}

// PaymentRecord is a free data retrieval call binding the contract method 0x89432e5e.
//
// Solidity: function paymentRecord(uint64 , address ) view returns(uint64)
func (_Contracts *ContractsCallerSession) PaymentRecord(arg0 uint64, arg1 common.Address) (uint64, error) {
	return _Contracts.Contract.PaymentRecord(&_Contracts.CallOpts, arg0, arg1)
}

// CreateAsset is a paid mutator transaction binding the contract method 0xfe12e6a3.
//
// Solidity: function createAsset(uint64 _contentId, address _paymentToken, uint256 _amount) returns()
func (_Contracts *ContractsTransactor) CreateAsset(opts *bind.TransactOpts, _contentId uint64, _paymentToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createAsset", _contentId, _paymentToken, _amount)
}

// CreateAsset is a paid mutator transaction binding the contract method 0xfe12e6a3.
//
// Solidity: function createAsset(uint64 _contentId, address _paymentToken, uint256 _amount) returns()
func (_Contracts *ContractsSession) CreateAsset(_contentId uint64, _paymentToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateAsset(&_Contracts.TransactOpts, _contentId, _paymentToken, _amount)
}

// CreateAsset is a paid mutator transaction binding the contract method 0xfe12e6a3.
//
// Solidity: function createAsset(uint64 _contentId, address _paymentToken, uint256 _amount) returns()
func (_Contracts *ContractsTransactorSession) CreateAsset(_contentId uint64, _paymentToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateAsset(&_Contracts.TransactOpts, _contentId, _paymentToken, _amount)
}

// PurchaseAsset is a paid mutator transaction binding the contract method 0xbd2aaf7f.
//
// Solidity: function purchaseAsset(uint64 _assetId) returns()
func (_Contracts *ContractsTransactor) PurchaseAsset(opts *bind.TransactOpts, _assetId uint64) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "purchaseAsset", _assetId)
}

// PurchaseAsset is a paid mutator transaction binding the contract method 0xbd2aaf7f.
//
// Solidity: function purchaseAsset(uint64 _assetId) returns()
func (_Contracts *ContractsSession) PurchaseAsset(_assetId uint64) (*types.Transaction, error) {
	return _Contracts.Contract.PurchaseAsset(&_Contracts.TransactOpts, _assetId)
}

// PurchaseAsset is a paid mutator transaction binding the contract method 0xbd2aaf7f.
//
// Solidity: function purchaseAsset(uint64 _assetId) returns()
func (_Contracts *ContractsTransactorSession) PurchaseAsset(_assetId uint64) (*types.Transaction, error) {
	return _Contracts.Contract.PurchaseAsset(&_Contracts.TransactOpts, _assetId)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address _tokenAddress, uint256 _amount) returns()
func (_Contracts *ContractsTransactor) WithdrawToken(opts *bind.TransactOpts, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawToken", _tokenAddress, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address _tokenAddress, uint256 _amount) returns()
func (_Contracts *ContractsSession) WithdrawToken(_tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawToken(&_Contracts.TransactOpts, _tokenAddress, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address _tokenAddress, uint256 _amount) returns()
func (_Contracts *ContractsTransactorSession) WithdrawToken(_tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawToken(&_Contracts.TransactOpts, _tokenAddress, _amount)
}

// ContractsCreateAssetIterator is returned from FilterCreateAsset and is used to iterate over the raw logs and unpacked data for CreateAsset events raised by the Contracts contract.
type ContractsCreateAssetIterator struct {
	Event *ContractsCreateAsset // Event containing the contract specifics and raw log

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
func (it *ContractsCreateAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCreateAsset)
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
		it.Event = new(ContractsCreateAsset)
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
func (it *ContractsCreateAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCreateAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCreateAsset represents a CreateAsset event raised by the Contracts contract.
type ContractsCreateAsset struct {
	Author    common.Address
	ContentId uint64
	AssetId   uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateAsset is a free log retrieval operation binding the contract event 0x269cf5c71029c03bbfc885db18cdb5e3044db3d4ebe27d52df30f4fa490eb280.
//
// Solidity: event CreateAsset(address author, uint64 contentId, uint64 assetId)
func (_Contracts *ContractsFilterer) FilterCreateAsset(opts *bind.FilterOpts) (*ContractsCreateAssetIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CreateAsset")
	if err != nil {
		return nil, err
	}
	return &ContractsCreateAssetIterator{contract: _Contracts.contract, event: "CreateAsset", logs: logs, sub: sub}, nil
}

// WatchCreateAsset is a free log subscription operation binding the contract event 0x269cf5c71029c03bbfc885db18cdb5e3044db3d4ebe27d52df30f4fa490eb280.
//
// Solidity: event CreateAsset(address author, uint64 contentId, uint64 assetId)
func (_Contracts *ContractsFilterer) WatchCreateAsset(opts *bind.WatchOpts, sink chan<- *ContractsCreateAsset) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CreateAsset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCreateAsset)
				if err := _Contracts.contract.UnpackLog(event, "CreateAsset", log); err != nil {
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

// ParseCreateAsset is a log parse operation binding the contract event 0x269cf5c71029c03bbfc885db18cdb5e3044db3d4ebe27d52df30f4fa490eb280.
//
// Solidity: event CreateAsset(address author, uint64 contentId, uint64 assetId)
func (_Contracts *ContractsFilterer) ParseCreateAsset(log types.Log) (*ContractsCreateAsset, error) {
	event := new(ContractsCreateAsset)
	if err := _Contracts.contract.UnpackLog(event, "CreateAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
