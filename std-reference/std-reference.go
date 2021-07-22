// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stdreference

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IStdReferenceReferenceData is an auto generated low-level Go binding around an user-defined struct.
type IStdReferenceReferenceData struct {
	Rate             *big.Int
	LastUpdatedBase  *big.Int
	LastUpdatedQuote *big.Int
}

// StdreferenceABI is the input ABI used to generate the binding from.
const StdreferenceABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_base\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_quote\",\"type\":\"string\"}],\"name\":\"getReferenceData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedQuote\",\"type\":\"uint256\"}],\"internalType\":\"structIStdReference.ReferenceData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_bases\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_quotes\",\"type\":\"string[]\"}],\"name\":\"getReferenceDataBulk\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedQuote\",\"type\":\"uint256\"}],\"internalType\":\"structIStdReference.ReferenceData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Stdreference is an auto generated Go binding around an Ethereum contract.
type Stdreference struct {
	StdreferenceCaller     // Read-only binding to the contract
	StdreferenceTransactor // Write-only binding to the contract
	StdreferenceFilterer   // Log filterer for contract events
}

// StdreferenceCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdreferenceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdreferenceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdreferenceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdreferenceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdreferenceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdreferenceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdreferenceSession struct {
	Contract     *Stdreference     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdreferenceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdreferenceCallerSession struct {
	Contract *StdreferenceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StdreferenceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdreferenceTransactorSession struct {
	Contract     *StdreferenceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StdreferenceRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdreferenceRaw struct {
	Contract *Stdreference // Generic contract binding to access the raw methods on
}

// StdreferenceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdreferenceCallerRaw struct {
	Contract *StdreferenceCaller // Generic read-only contract binding to access the raw methods on
}

// StdreferenceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdreferenceTransactorRaw struct {
	Contract *StdreferenceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdreference creates a new instance of Stdreference, bound to a specific deployed contract.
func NewStdreference(address common.Address, backend bind.ContractBackend) (*Stdreference, error) {
	contract, err := bindStdreference(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stdreference{StdreferenceCaller: StdreferenceCaller{contract: contract}, StdreferenceTransactor: StdreferenceTransactor{contract: contract}, StdreferenceFilterer: StdreferenceFilterer{contract: contract}}, nil
}

// NewStdreferenceCaller creates a new read-only instance of Stdreference, bound to a specific deployed contract.
func NewStdreferenceCaller(address common.Address, caller bind.ContractCaller) (*StdreferenceCaller, error) {
	contract, err := bindStdreference(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdreferenceCaller{contract: contract}, nil
}

// NewStdreferenceTransactor creates a new write-only instance of Stdreference, bound to a specific deployed contract.
func NewStdreferenceTransactor(address common.Address, transactor bind.ContractTransactor) (*StdreferenceTransactor, error) {
	contract, err := bindStdreference(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdreferenceTransactor{contract: contract}, nil
}

// NewStdreferenceFilterer creates a new log filterer instance of Stdreference, bound to a specific deployed contract.
func NewStdreferenceFilterer(address common.Address, filterer bind.ContractFilterer) (*StdreferenceFilterer, error) {
	contract, err := bindStdreference(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdreferenceFilterer{contract: contract}, nil
}

// bindStdreference binds a generic wrapper to an already deployed contract.
func bindStdreference(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StdreferenceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stdreference *StdreferenceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stdreference.Contract.StdreferenceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stdreference *StdreferenceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stdreference.Contract.StdreferenceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stdreference *StdreferenceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stdreference.Contract.StdreferenceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stdreference *StdreferenceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stdreference.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stdreference *StdreferenceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stdreference.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stdreference *StdreferenceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stdreference.Contract.contract.Transact(opts, method, params...)
}

// GetReferenceData is a free data retrieval call binding the contract method 0x65555bcc.
//
// Solidity: function getReferenceData(string _base, string _quote) view returns((uint256,uint256,uint256))
func (_Stdreference *StdreferenceCaller) GetReferenceData(opts *bind.CallOpts, _base string, _quote string) (IStdReferenceReferenceData, error) {
	var out []interface{}
	err := _Stdreference.contract.Call(opts, &out, "getReferenceData", _base, _quote)

	if err != nil {
		return *new(IStdReferenceReferenceData), err
	}

	out0 := *abi.ConvertType(out[0], new(IStdReferenceReferenceData)).(*IStdReferenceReferenceData)

	return out0, err

}

// GetReferenceData is a free data retrieval call binding the contract method 0x65555bcc.
//
// Solidity: function getReferenceData(string _base, string _quote) view returns((uint256,uint256,uint256))
func (_Stdreference *StdreferenceSession) GetReferenceData(_base string, _quote string) (IStdReferenceReferenceData, error) {
	return _Stdreference.Contract.GetReferenceData(&_Stdreference.CallOpts, _base, _quote)
}

// GetReferenceData is a free data retrieval call binding the contract method 0x65555bcc.
//
// Solidity: function getReferenceData(string _base, string _quote) view returns((uint256,uint256,uint256))
func (_Stdreference *StdreferenceCallerSession) GetReferenceData(_base string, _quote string) (IStdReferenceReferenceData, error) {
	return _Stdreference.Contract.GetReferenceData(&_Stdreference.CallOpts, _base, _quote)
}

// GetReferenceDataBulk is a free data retrieval call binding the contract method 0xe42a071b.
//
// Solidity: function getReferenceDataBulk(string[] _bases, string[] _quotes) view returns((uint256,uint256,uint256)[])
func (_Stdreference *StdreferenceCaller) GetReferenceDataBulk(opts *bind.CallOpts, _bases []string, _quotes []string) ([]IStdReferenceReferenceData, error) {
	var out []interface{}
	err := _Stdreference.contract.Call(opts, &out, "getReferenceDataBulk", _bases, _quotes)

	if err != nil {
		return *new([]IStdReferenceReferenceData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStdReferenceReferenceData)).(*[]IStdReferenceReferenceData)

	return out0, err

}

// GetReferenceDataBulk is a free data retrieval call binding the contract method 0xe42a071b.
//
// Solidity: function getReferenceDataBulk(string[] _bases, string[] _quotes) view returns((uint256,uint256,uint256)[])
func (_Stdreference *StdreferenceSession) GetReferenceDataBulk(_bases []string, _quotes []string) ([]IStdReferenceReferenceData, error) {
	return _Stdreference.Contract.GetReferenceDataBulk(&_Stdreference.CallOpts, _bases, _quotes)
}

// GetReferenceDataBulk is a free data retrieval call binding the contract method 0xe42a071b.
//
// Solidity: function getReferenceDataBulk(string[] _bases, string[] _quotes) view returns((uint256,uint256,uint256)[])
func (_Stdreference *StdreferenceCallerSession) GetReferenceDataBulk(_bases []string, _quotes []string) ([]IStdReferenceReferenceData, error) {
	return _Stdreference.Contract.GetReferenceDataBulk(&_Stdreference.CallOpts, _bases, _quotes)
}
