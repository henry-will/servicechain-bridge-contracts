// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package callback

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

// CallbackABI is the input ABI used to generate the binding from.
const CallbackABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueOrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"RegisteredOffer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_valueOrID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"registerOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CallbackBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CallbackBinRuntime = ``

// CallbackBin is the compiled bytecode used for deploying new contracts.
var CallbackBin = "0x608060405234801561001057600080fd5b50610228806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80635ec467e214610030575b600080fd5b61004a60048036038101906100459190610128565b61004c565b005b7f6e0b5117e49b57aaf37c635363b1b78a14ad521875ec99079d95bee2838cfeb88484848460405161008194939291906101ad565b60405180910390a150505050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100bf82610094565b9050919050565b6100cf816100b4565b81146100da57600080fd5b50565b6000813590506100ec816100c6565b92915050565b6000819050919050565b610105816100f2565b811461011057600080fd5b50565b600081359050610122816100fc565b92915050565b600080600080608085870312156101425761014161008f565b5b6000610150878288016100dd565b945050602061016187828801610113565b9350506040610172878288016100dd565b925050606061018387828801610113565b91505092959194509250565b610198816100b4565b82525050565b6101a7816100f2565b82525050565b60006080820190506101c2600083018761018f565b6101cf602083018661019e565b6101dc604083018561018f565b6101e9606083018461019e565b9594505050505056fea2646970667358221220a8e1c48c8625c3976cf4f847a4707bcd870030ad62be7f8537c76db8ba2ed46764736f6c634300080e0033"

// DeployCallback deploys a new Klaytn contract, binding an instance of Callback to it.
func DeployCallback(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Callback, error) {
	parsed, err := abi.JSON(strings.NewReader(CallbackABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CallbackBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Callback{CallbackCaller: CallbackCaller{contract: contract}, CallbackTransactor: CallbackTransactor{contract: contract}, CallbackFilterer: CallbackFilterer{contract: contract}}, nil
}

// Callback is an auto generated Go binding around a Klaytn contract.
type Callback struct {
	CallbackCaller     // Read-only binding to the contract
	CallbackTransactor // Write-only binding to the contract
	CallbackFilterer   // Log filterer for contract events
}

// CallbackCaller is an auto generated read-only Go binding around a Klaytn contract.
type CallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallbackTransactor is an auto generated write-only Go binding around a Klaytn contract.
type CallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallbackFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type CallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallbackSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type CallbackSession struct {
	Contract     *Callback         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallbackCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type CallbackCallerSession struct {
	Contract *CallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CallbackTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type CallbackTransactorSession struct {
	Contract     *CallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CallbackRaw is an auto generated low-level Go binding around a Klaytn contract.
type CallbackRaw struct {
	Contract *Callback // Generic contract binding to access the raw methods on
}

// CallbackCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type CallbackCallerRaw struct {
	Contract *CallbackCaller // Generic read-only contract binding to access the raw methods on
}

// CallbackTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type CallbackTransactorRaw struct {
	Contract *CallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallback creates a new instance of Callback, bound to a specific deployed contract.
func NewCallback(address common.Address, backend bind.ContractBackend) (*Callback, error) {
	contract, err := bindCallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Callback{CallbackCaller: CallbackCaller{contract: contract}, CallbackTransactor: CallbackTransactor{contract: contract}, CallbackFilterer: CallbackFilterer{contract: contract}}, nil
}

// NewCallbackCaller creates a new read-only instance of Callback, bound to a specific deployed contract.
func NewCallbackCaller(address common.Address, caller bind.ContractCaller) (*CallbackCaller, error) {
	contract, err := bindCallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallbackCaller{contract: contract}, nil
}

// NewCallbackTransactor creates a new write-only instance of Callback, bound to a specific deployed contract.
func NewCallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*CallbackTransactor, error) {
	contract, err := bindCallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallbackTransactor{contract: contract}, nil
}

// NewCallbackFilterer creates a new log filterer instance of Callback, bound to a specific deployed contract.
func NewCallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*CallbackFilterer, error) {
	contract, err := bindCallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallbackFilterer{contract: contract}, nil
}

// bindCallback binds a generic wrapper to an already deployed contract.
func bindCallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CallbackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Callback *CallbackRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Callback.Contract.CallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Callback *CallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callback.Contract.CallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Callback *CallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Callback.Contract.CallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Callback *CallbackCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Callback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Callback *CallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Callback *CallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Callback.Contract.contract.Transact(opts, method, params...)
}

// RegisterOffer is a paid mutator transaction binding the contract method 0x5ec467e2.
//
// Solidity: function registerOffer(address _owner, uint256 _valueOrID, address _tokenAddress, uint256 _price) returns()
func (_Callback *CallbackTransactor) RegisterOffer(opts *bind.TransactOpts, _owner common.Address, _valueOrID *big.Int, _tokenAddress common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Callback.contract.Transact(opts, "registerOffer", _owner, _valueOrID, _tokenAddress, _price)
}

// RegisterOffer is a paid mutator transaction binding the contract method 0x5ec467e2.
//
// Solidity: function registerOffer(address _owner, uint256 _valueOrID, address _tokenAddress, uint256 _price) returns()
func (_Callback *CallbackSession) RegisterOffer(_owner common.Address, _valueOrID *big.Int, _tokenAddress common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Callback.Contract.RegisterOffer(&_Callback.TransactOpts, _owner, _valueOrID, _tokenAddress, _price)
}

// RegisterOffer is a paid mutator transaction binding the contract method 0x5ec467e2.
//
// Solidity: function registerOffer(address _owner, uint256 _valueOrID, address _tokenAddress, uint256 _price) returns()
func (_Callback *CallbackTransactorSession) RegisterOffer(_owner common.Address, _valueOrID *big.Int, _tokenAddress common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Callback.Contract.RegisterOffer(&_Callback.TransactOpts, _owner, _valueOrID, _tokenAddress, _price)
}

// CallbackRegisteredOfferIterator is returned from FilterRegisteredOffer and is used to iterate over the raw logs and unpacked data for RegisteredOffer events raised by the Callback contract.
type CallbackRegisteredOfferIterator struct {
	Event *CallbackRegisteredOffer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CallbackRegisteredOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CallbackRegisteredOffer)
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
		it.Event = new(CallbackRegisteredOffer)
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
func (it *CallbackRegisteredOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CallbackRegisteredOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CallbackRegisteredOffer represents a RegisteredOffer event raised by the Callback contract.
type CallbackRegisteredOffer struct {
	Owner        common.Address
	ValueOrID    *big.Int
	TokenAddress common.Address
	Price        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisteredOffer is a free log retrieval operation binding the contract event 0x6e0b5117e49b57aaf37c635363b1b78a14ad521875ec99079d95bee2838cfeb8.
//
// Solidity: event RegisteredOffer(address owner, uint256 valueOrID, address tokenAddress, uint256 price)
func (_Callback *CallbackFilterer) FilterRegisteredOffer(opts *bind.FilterOpts) (*CallbackRegisteredOfferIterator, error) {

	logs, sub, err := _Callback.contract.FilterLogs(opts, "RegisteredOffer")
	if err != nil {
		return nil, err
	}
	return &CallbackRegisteredOfferIterator{contract: _Callback.contract, event: "RegisteredOffer", logs: logs, sub: sub}, nil
}

// WatchRegisteredOffer is a free log subscription operation binding the contract event 0x6e0b5117e49b57aaf37c635363b1b78a14ad521875ec99079d95bee2838cfeb8.
//
// Solidity: event RegisteredOffer(address owner, uint256 valueOrID, address tokenAddress, uint256 price)
func (_Callback *CallbackFilterer) WatchRegisteredOffer(opts *bind.WatchOpts, sink chan<- *CallbackRegisteredOffer) (event.Subscription, error) {

	logs, sub, err := _Callback.contract.WatchLogs(opts, "RegisteredOffer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CallbackRegisteredOffer)
				if err := _Callback.contract.UnpackLog(event, "RegisteredOffer", log); err != nil {
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

// ParseRegisteredOffer is a log parse operation binding the contract event 0x6e0b5117e49b57aaf37c635363b1b78a14ad521875ec99079d95bee2838cfeb8.
//
// Solidity: event RegisteredOffer(address owner, uint256 valueOrID, address tokenAddress, uint256 price)
func (_Callback *CallbackFilterer) ParseRegisteredOffer(log types.Log) (*CallbackRegisteredOffer, error) {
	event := new(CallbackRegisteredOffer)
	if err := _Callback.contract.UnpackLog(event, "RegisteredOffer", log); err != nil {
		return nil, err
	}
	return event, nil
}
