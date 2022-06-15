// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kip13

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

// Kip13MetaData contains all meta data concerning the Kip13 contract.
var Kip13MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Kip13ABI is the input ABI used to generate the binding from.
// Deprecated: Use Kip13MetaData.ABI instead.
var Kip13ABI = Kip13MetaData.ABI

// Kip13 is an auto generated Go binding around an Ethereum contract.
type Kip13 struct {
	Kip13Caller     // Read-only binding to the contract
	Kip13Transactor // Write-only binding to the contract
	Kip13Filterer   // Log filterer for contract events
}

// Kip13Caller is an auto generated read-only Go binding around an Ethereum contract.
type Kip13Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Kip13Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Kip13Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Kip13Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Kip13Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Kip13Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Kip13Session struct {
	Contract     *Kip13            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Kip13CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Kip13CallerSession struct {
	Contract *Kip13Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Kip13TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Kip13TransactorSession struct {
	Contract     *Kip13Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Kip13Raw is an auto generated low-level Go binding around an Ethereum contract.
type Kip13Raw struct {
	Contract *Kip13 // Generic contract binding to access the raw methods on
}

// Kip13CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Kip13CallerRaw struct {
	Contract *Kip13Caller // Generic read-only contract binding to access the raw methods on
}

// Kip13TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Kip13TransactorRaw struct {
	Contract *Kip13Transactor // Generic write-only contract binding to access the raw methods on
}

// NewKip13 creates a new instance of Kip13, bound to a specific deployed contract.
func NewKip13(address common.Address, backend bind.ContractBackend) (*Kip13, error) {
	contract, err := bindKip13(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kip13{Kip13Caller: Kip13Caller{contract: contract}, Kip13Transactor: Kip13Transactor{contract: contract}, Kip13Filterer: Kip13Filterer{contract: contract}}, nil
}

// NewKip13Caller creates a new read-only instance of Kip13, bound to a specific deployed contract.
func NewKip13Caller(address common.Address, caller bind.ContractCaller) (*Kip13Caller, error) {
	contract, err := bindKip13(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Kip13Caller{contract: contract}, nil
}

// NewKip13Transactor creates a new write-only instance of Kip13, bound to a specific deployed contract.
func NewKip13Transactor(address common.Address, transactor bind.ContractTransactor) (*Kip13Transactor, error) {
	contract, err := bindKip13(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Kip13Transactor{contract: contract}, nil
}

// NewKip13Filterer creates a new log filterer instance of Kip13, bound to a specific deployed contract.
func NewKip13Filterer(address common.Address, filterer bind.ContractFilterer) (*Kip13Filterer, error) {
	contract, err := bindKip13(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Kip13Filterer{contract: contract}, nil
}

// bindKip13 binds a generic wrapper to an already deployed contract.
func bindKip13(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Kip13ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kip13 *Kip13Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kip13.Contract.Kip13Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kip13 *Kip13Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kip13.Contract.Kip13Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kip13 *Kip13Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kip13.Contract.Kip13Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kip13 *Kip13CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kip13.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kip13 *Kip13TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kip13.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kip13 *Kip13TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kip13.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Kip13 *Kip13Caller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Kip13.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Kip13 *Kip13Session) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Kip13.Contract.SupportsInterface(&_Kip13.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Kip13 *Kip13CallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Kip13.Contract.SupportsInterface(&_Kip13.CallOpts, interfaceID)
}
