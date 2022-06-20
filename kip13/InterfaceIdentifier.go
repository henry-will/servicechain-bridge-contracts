// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kip13

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IERC20BridgeReceiverABI is the input ABI used to generate the binding from.
const IERC20BridgeReceiverABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"onERC20Received\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20BridgeReceiverBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IERC20BridgeReceiverBinRuntime = ``

// IERC20BridgeReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IERC20BridgeReceiverFuncSigs = map[string]string{
	"f1656e53": "onERC20Received(address,address,uint256,uint256,bytes)",
}

// IERC20BridgeReceiver is an auto generated Go binding around a Klaytn contract.
type IERC20BridgeReceiver struct {
	IERC20BridgeReceiverCaller     // Read-only binding to the contract
	IERC20BridgeReceiverTransactor // Write-only binding to the contract
	IERC20BridgeReceiverFilterer   // Log filterer for contract events
}

// IERC20BridgeReceiverCaller is an auto generated read-only Go binding around a Klaytn contract.
type IERC20BridgeReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20BridgeReceiverTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IERC20BridgeReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20BridgeReceiverFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IERC20BridgeReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20BridgeReceiverSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IERC20BridgeReceiverSession struct {
	Contract     *IERC20BridgeReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IERC20BridgeReceiverCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IERC20BridgeReceiverCallerSession struct {
	Contract *IERC20BridgeReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IERC20BridgeReceiverTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IERC20BridgeReceiverTransactorSession struct {
	Contract     *IERC20BridgeReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IERC20BridgeReceiverRaw is an auto generated low-level Go binding around a Klaytn contract.
type IERC20BridgeReceiverRaw struct {
	Contract *IERC20BridgeReceiver // Generic contract binding to access the raw methods on
}

// IERC20BridgeReceiverCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IERC20BridgeReceiverCallerRaw struct {
	Contract *IERC20BridgeReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20BridgeReceiverTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IERC20BridgeReceiverTransactorRaw struct {
	Contract *IERC20BridgeReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20BridgeReceiver creates a new instance of IERC20BridgeReceiver, bound to a specific deployed contract.
func NewIERC20BridgeReceiver(address common.Address, backend bind.ContractBackend) (*IERC20BridgeReceiver, error) {
	contract, err := bindIERC20BridgeReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20BridgeReceiver{IERC20BridgeReceiverCaller: IERC20BridgeReceiverCaller{contract: contract}, IERC20BridgeReceiverTransactor: IERC20BridgeReceiverTransactor{contract: contract}, IERC20BridgeReceiverFilterer: IERC20BridgeReceiverFilterer{contract: contract}}, nil
}

// NewIERC20BridgeReceiverCaller creates a new read-only instance of IERC20BridgeReceiver, bound to a specific deployed contract.
func NewIERC20BridgeReceiverCaller(address common.Address, caller bind.ContractCaller) (*IERC20BridgeReceiverCaller, error) {
	contract, err := bindIERC20BridgeReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20BridgeReceiverCaller{contract: contract}, nil
}

// NewIERC20BridgeReceiverTransactor creates a new write-only instance of IERC20BridgeReceiver, bound to a specific deployed contract.
func NewIERC20BridgeReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20BridgeReceiverTransactor, error) {
	contract, err := bindIERC20BridgeReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20BridgeReceiverTransactor{contract: contract}, nil
}

// NewIERC20BridgeReceiverFilterer creates a new log filterer instance of IERC20BridgeReceiver, bound to a specific deployed contract.
func NewIERC20BridgeReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20BridgeReceiverFilterer, error) {
	contract, err := bindIERC20BridgeReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20BridgeReceiverFilterer{contract: contract}, nil
}

// bindIERC20BridgeReceiver binds a generic wrapper to an already deployed contract.
func bindIERC20BridgeReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20BridgeReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20BridgeReceiver *IERC20BridgeReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20BridgeReceiver.Contract.IERC20BridgeReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20BridgeReceiver *IERC20BridgeReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20BridgeReceiver.Contract.IERC20BridgeReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20BridgeReceiver *IERC20BridgeReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20BridgeReceiver.Contract.IERC20BridgeReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20BridgeReceiver *IERC20BridgeReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20BridgeReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20BridgeReceiver *IERC20BridgeReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20BridgeReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20BridgeReceiver *IERC20BridgeReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20BridgeReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf1656e53.
//
// Solidity: function onERC20Received(address _from, address _to, uint256 _amount, uint256 _feeLimit, bytes _extraData) returns()
func (_IERC20BridgeReceiver *IERC20BridgeReceiverTransactor) OnERC20Received(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _IERC20BridgeReceiver.contract.Transact(opts, "onERC20Received", _from, _to, _amount, _feeLimit, _extraData)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf1656e53.
//
// Solidity: function onERC20Received(address _from, address _to, uint256 _amount, uint256 _feeLimit, bytes _extraData) returns()
func (_IERC20BridgeReceiver *IERC20BridgeReceiverSession) OnERC20Received(_from common.Address, _to common.Address, _amount *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _IERC20BridgeReceiver.Contract.OnERC20Received(&_IERC20BridgeReceiver.TransactOpts, _from, _to, _amount, _feeLimit, _extraData)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf1656e53.
//
// Solidity: function onERC20Received(address _from, address _to, uint256 _amount, uint256 _feeLimit, bytes _extraData) returns()
func (_IERC20BridgeReceiver *IERC20BridgeReceiverTransactorSession) OnERC20Received(_from common.Address, _to common.Address, _amount *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _IERC20BridgeReceiver.Contract.OnERC20Received(&_IERC20BridgeReceiver.TransactOpts, _from, _to, _amount, _feeLimit, _extraData)
}

// IERC20MintABI is the input ABI used to generate the binding from.
const IERC20MintABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20MintBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IERC20MintBinRuntime = ``

// IERC20MintFuncSigs maps the 4-byte function signature to its string representation.
var IERC20MintFuncSigs = map[string]string{
	"40c10f19": "mint(address,uint256)",
}

// IERC20Mint is an auto generated Go binding around a Klaytn contract.
type IERC20Mint struct {
	IERC20MintCaller     // Read-only binding to the contract
	IERC20MintTransactor // Write-only binding to the contract
	IERC20MintFilterer   // Log filterer for contract events
}

// IERC20MintCaller is an auto generated read-only Go binding around a Klaytn contract.
type IERC20MintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MintTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IERC20MintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MintFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IERC20MintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MintSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IERC20MintSession struct {
	Contract     *IERC20Mint       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MintCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IERC20MintCallerSession struct {
	Contract *IERC20MintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IERC20MintTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IERC20MintTransactorSession struct {
	Contract     *IERC20MintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IERC20MintRaw is an auto generated low-level Go binding around a Klaytn contract.
type IERC20MintRaw struct {
	Contract *IERC20Mint // Generic contract binding to access the raw methods on
}

// IERC20MintCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IERC20MintCallerRaw struct {
	Contract *IERC20MintCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MintTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IERC20MintTransactorRaw struct {
	Contract *IERC20MintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Mint creates a new instance of IERC20Mint, bound to a specific deployed contract.
func NewIERC20Mint(address common.Address, backend bind.ContractBackend) (*IERC20Mint, error) {
	contract, err := bindIERC20Mint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Mint{IERC20MintCaller: IERC20MintCaller{contract: contract}, IERC20MintTransactor: IERC20MintTransactor{contract: contract}, IERC20MintFilterer: IERC20MintFilterer{contract: contract}}, nil
}

// NewIERC20MintCaller creates a new read-only instance of IERC20Mint, bound to a specific deployed contract.
func NewIERC20MintCaller(address common.Address, caller bind.ContractCaller) (*IERC20MintCaller, error) {
	contract, err := bindIERC20Mint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MintCaller{contract: contract}, nil
}

// NewIERC20MintTransactor creates a new write-only instance of IERC20Mint, bound to a specific deployed contract.
func NewIERC20MintTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MintTransactor, error) {
	contract, err := bindIERC20Mint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MintTransactor{contract: contract}, nil
}

// NewIERC20MintFilterer creates a new log filterer instance of IERC20Mint, bound to a specific deployed contract.
func NewIERC20MintFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MintFilterer, error) {
	contract, err := bindIERC20Mint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MintFilterer{contract: contract}, nil
}

// bindIERC20Mint binds a generic wrapper to an already deployed contract.
func bindIERC20Mint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20MintABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Mint *IERC20MintRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20Mint.Contract.IERC20MintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Mint *IERC20MintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Mint.Contract.IERC20MintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Mint *IERC20MintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Mint.Contract.IERC20MintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Mint *IERC20MintCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20Mint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Mint *IERC20MintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Mint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Mint *IERC20MintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Mint.Contract.contract.Transact(opts, method, params...)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IERC20Mint *IERC20MintTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mint.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IERC20Mint *IERC20MintSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mint.Contract.Mint(&_IERC20Mint.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IERC20Mint *IERC20MintTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mint.Contract.Mint(&_IERC20Mint.TransactOpts, to, amount)
}

// InterfaceIdentifierABI is the input ABI used to generate the binding from.
const InterfaceIdentifierABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"onERC20Received\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// InterfaceIdentifierBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const InterfaceIdentifierBinRuntime = ``

// InterfaceIdentifierFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceIdentifierFuncSigs = map[string]string{
	"f1656e53": "onERC20Received(address,address,uint256,uint256,bytes)",
	"01ffc9a7": "supportsInterface(bytes4)",
}

// InterfaceIdentifier is an auto generated Go binding around a Klaytn contract.
type InterfaceIdentifier struct {
	InterfaceIdentifierCaller     // Read-only binding to the contract
	InterfaceIdentifierTransactor // Write-only binding to the contract
	InterfaceIdentifierFilterer   // Log filterer for contract events
}

// InterfaceIdentifierCaller is an auto generated read-only Go binding around a Klaytn contract.
type InterfaceIdentifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceIdentifierTransactor is an auto generated write-only Go binding around a Klaytn contract.
type InterfaceIdentifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceIdentifierFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type InterfaceIdentifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceIdentifierSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type InterfaceIdentifierSession struct {
	Contract     *InterfaceIdentifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// InterfaceIdentifierCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type InterfaceIdentifierCallerSession struct {
	Contract *InterfaceIdentifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// InterfaceIdentifierTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type InterfaceIdentifierTransactorSession struct {
	Contract     *InterfaceIdentifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// InterfaceIdentifierRaw is an auto generated low-level Go binding around a Klaytn contract.
type InterfaceIdentifierRaw struct {
	Contract *InterfaceIdentifier // Generic contract binding to access the raw methods on
}

// InterfaceIdentifierCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type InterfaceIdentifierCallerRaw struct {
	Contract *InterfaceIdentifierCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceIdentifierTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type InterfaceIdentifierTransactorRaw struct {
	Contract *InterfaceIdentifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceIdentifier creates a new instance of InterfaceIdentifier, bound to a specific deployed contract.
func NewInterfaceIdentifier(address common.Address, backend bind.ContractBackend) (*InterfaceIdentifier, error) {
	contract, err := bindInterfaceIdentifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceIdentifier{InterfaceIdentifierCaller: InterfaceIdentifierCaller{contract: contract}, InterfaceIdentifierTransactor: InterfaceIdentifierTransactor{contract: contract}, InterfaceIdentifierFilterer: InterfaceIdentifierFilterer{contract: contract}}, nil
}

// NewInterfaceIdentifierCaller creates a new read-only instance of InterfaceIdentifier, bound to a specific deployed contract.
func NewInterfaceIdentifierCaller(address common.Address, caller bind.ContractCaller) (*InterfaceIdentifierCaller, error) {
	contract, err := bindInterfaceIdentifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceIdentifierCaller{contract: contract}, nil
}

// NewInterfaceIdentifierTransactor creates a new write-only instance of InterfaceIdentifier, bound to a specific deployed contract.
func NewInterfaceIdentifierTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceIdentifierTransactor, error) {
	contract, err := bindInterfaceIdentifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceIdentifierTransactor{contract: contract}, nil
}

// NewInterfaceIdentifierFilterer creates a new log filterer instance of InterfaceIdentifier, bound to a specific deployed contract.
func NewInterfaceIdentifierFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceIdentifierFilterer, error) {
	contract, err := bindInterfaceIdentifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceIdentifierFilterer{contract: contract}, nil
}

// bindInterfaceIdentifier binds a generic wrapper to an already deployed contract.
func bindInterfaceIdentifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceIdentifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceIdentifier *InterfaceIdentifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InterfaceIdentifier.Contract.InterfaceIdentifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceIdentifier *InterfaceIdentifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceIdentifier.Contract.InterfaceIdentifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceIdentifier *InterfaceIdentifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceIdentifier.Contract.InterfaceIdentifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceIdentifier *InterfaceIdentifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InterfaceIdentifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceIdentifier *InterfaceIdentifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceIdentifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceIdentifier *InterfaceIdentifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceIdentifier.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_InterfaceIdentifier *InterfaceIdentifierCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _InterfaceIdentifier.contract.Call(opts, out, "supportsInterface", interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_InterfaceIdentifier *InterfaceIdentifierSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _InterfaceIdentifier.Contract.SupportsInterface(&_InterfaceIdentifier.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_InterfaceIdentifier *InterfaceIdentifierCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _InterfaceIdentifier.Contract.SupportsInterface(&_InterfaceIdentifier.CallOpts, interfaceID)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf1656e53.
//
// Solidity: function onERC20Received(address _from, address _to, uint256 _amount, uint256 _feeLimit, bytes _extraData) returns()
func (_InterfaceIdentifier *InterfaceIdentifierTransactor) OnERC20Received(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _InterfaceIdentifier.contract.Transact(opts, "onERC20Received", _from, _to, _amount, _feeLimit, _extraData)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf1656e53.
//
// Solidity: function onERC20Received(address _from, address _to, uint256 _amount, uint256 _feeLimit, bytes _extraData) returns()
func (_InterfaceIdentifier *InterfaceIdentifierSession) OnERC20Received(_from common.Address, _to common.Address, _amount *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _InterfaceIdentifier.Contract.OnERC20Received(&_InterfaceIdentifier.TransactOpts, _from, _to, _amount, _feeLimit, _extraData)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf1656e53.
//
// Solidity: function onERC20Received(address _from, address _to, uint256 _amount, uint256 _feeLimit, bytes _extraData) returns()
func (_InterfaceIdentifier *InterfaceIdentifierTransactorSession) OnERC20Received(_from common.Address, _to common.Address, _amount *big.Int, _feeLimit *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _InterfaceIdentifier.Contract.OnERC20Received(&_InterfaceIdentifier.TransactOpts, _from, _to, _amount, _feeLimit, _extraData)
}
