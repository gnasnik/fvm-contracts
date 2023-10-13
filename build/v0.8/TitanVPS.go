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
	_ = abi.ConvertType
)

// TitanVPSOrder is an auto generated low-level Go binding around an user-defined struct.
type TitanVPSOrder struct {
	OrderID    string
	BuyerAddr  common.Address
	Price      *big.Int
	Expiration *big.Int
	Timestamp  *big.Int
}

// TitanVPSMetaData contains all meta data concerning the TitanVPS contract.
var TitanVPSMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"orderID\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orderID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_buyerAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"newOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orderID\",\"type\":\"string\"}],\"name\":\"queryOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"orderID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"buyerAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structTitanVPS.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"orderID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"buyerAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TitanVPSABI is the input ABI used to generate the binding from.
// Deprecated: Use TitanVPSMetaData.ABI instead.
var TitanVPSABI = TitanVPSMetaData.ABI

// TitanVPS is an auto generated Go binding around an Ethereum contract.
type TitanVPS struct {
	TitanVPSCaller     // Read-only binding to the contract
	TitanVPSTransactor // Write-only binding to the contract
	TitanVPSFilterer   // Log filterer for contract events
}

// TitanVPSCaller is an auto generated read-only Go binding around an Ethereum contract.
type TitanVPSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TitanVPSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TitanVPSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TitanVPSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TitanVPSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TitanVPSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TitanVPSSession struct {
	Contract     *TitanVPS         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TitanVPSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TitanVPSCallerSession struct {
	Contract *TitanVPSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TitanVPSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TitanVPSTransactorSession struct {
	Contract     *TitanVPSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TitanVPSRaw is an auto generated low-level Go binding around an Ethereum contract.
type TitanVPSRaw struct {
	Contract *TitanVPS // Generic contract binding to access the raw methods on
}

// TitanVPSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TitanVPSCallerRaw struct {
	Contract *TitanVPSCaller // Generic read-only contract binding to access the raw methods on
}

// TitanVPSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TitanVPSTransactorRaw struct {
	Contract *TitanVPSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTitanVPS creates a new instance of TitanVPS, bound to a specific deployed contract.
func NewTitanVPS(address common.Address, backend bind.ContractBackend) (*TitanVPS, error) {
	contract, err := bindTitanVPS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TitanVPS{TitanVPSCaller: TitanVPSCaller{contract: contract}, TitanVPSTransactor: TitanVPSTransactor{contract: contract}, TitanVPSFilterer: TitanVPSFilterer{contract: contract}}, nil
}

// NewTitanVPSCaller creates a new read-only instance of TitanVPS, bound to a specific deployed contract.
func NewTitanVPSCaller(address common.Address, caller bind.ContractCaller) (*TitanVPSCaller, error) {
	contract, err := bindTitanVPS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TitanVPSCaller{contract: contract}, nil
}

// NewTitanVPSTransactor creates a new write-only instance of TitanVPS, bound to a specific deployed contract.
func NewTitanVPSTransactor(address common.Address, transactor bind.ContractTransactor) (*TitanVPSTransactor, error) {
	contract, err := bindTitanVPS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TitanVPSTransactor{contract: contract}, nil
}

// NewTitanVPSFilterer creates a new log filterer instance of TitanVPS, bound to a specific deployed contract.
func NewTitanVPSFilterer(address common.Address, filterer bind.ContractFilterer) (*TitanVPSFilterer, error) {
	contract, err := bindTitanVPS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TitanVPSFilterer{contract: contract}, nil
}

// bindTitanVPS binds a generic wrapper to an already deployed contract.
func bindTitanVPS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TitanVPSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TitanVPS *TitanVPSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TitanVPS.Contract.TitanVPSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TitanVPS *TitanVPSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TitanVPS.Contract.TitanVPSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TitanVPS *TitanVPSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TitanVPS.Contract.TitanVPSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TitanVPS *TitanVPSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TitanVPS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TitanVPS *TitanVPSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TitanVPS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TitanVPS *TitanVPSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TitanVPS.Contract.contract.Transact(opts, method, params...)
}

// Orders is a free data retrieval call binding the contract method 0x1a948947.
//
// Solidity: function orders(string ) view returns(bool)
func (_TitanVPS *TitanVPSCaller) Orders(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _TitanVPS.contract.Call(opts, &out, "orders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Orders is a free data retrieval call binding the contract method 0x1a948947.
//
// Solidity: function orders(string ) view returns(bool)
func (_TitanVPS *TitanVPSSession) Orders(arg0 string) (bool, error) {
	return _TitanVPS.Contract.Orders(&_TitanVPS.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0x1a948947.
//
// Solidity: function orders(string ) view returns(bool)
func (_TitanVPS *TitanVPSCallerSession) Orders(arg0 string) (bool, error) {
	return _TitanVPS.Contract.Orders(&_TitanVPS.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TitanVPS *TitanVPSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TitanVPS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TitanVPS *TitanVPSSession) Owner() (common.Address, error) {
	return _TitanVPS.Contract.Owner(&_TitanVPS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TitanVPS *TitanVPSCallerSession) Owner() (common.Address, error) {
	return _TitanVPS.Contract.Owner(&_TitanVPS.CallOpts)
}

// QueryOrder is a free data retrieval call binding the contract method 0x92d18c04.
//
// Solidity: function queryOrder(string _orderID) view returns((string,address,uint256,uint256,uint256))
func (_TitanVPS *TitanVPSCaller) QueryOrder(opts *bind.CallOpts, _orderID string) (TitanVPSOrder, error) {
	var out []interface{}
	err := _TitanVPS.contract.Call(opts, &out, "queryOrder", _orderID)

	if err != nil {
		return *new(TitanVPSOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(TitanVPSOrder)).(*TitanVPSOrder)

	return out0, err

}

// QueryOrder is a free data retrieval call binding the contract method 0x92d18c04.
//
// Solidity: function queryOrder(string _orderID) view returns((string,address,uint256,uint256,uint256))
func (_TitanVPS *TitanVPSSession) QueryOrder(_orderID string) (TitanVPSOrder, error) {
	return _TitanVPS.Contract.QueryOrder(&_TitanVPS.CallOpts, _orderID)
}

// QueryOrder is a free data retrieval call binding the contract method 0x92d18c04.
//
// Solidity: function queryOrder(string _orderID) view returns((string,address,uint256,uint256,uint256))
func (_TitanVPS *TitanVPSCallerSession) QueryOrder(_orderID string) (TitanVPSOrder, error) {
	return _TitanVPS.Contract.QueryOrder(&_TitanVPS.CallOpts, _orderID)
}

// Records is a free data retrieval call binding the contract method 0x541e771d.
//
// Solidity: function records(string ) view returns(string orderID, address buyerAddr, uint256 price, uint256 expiration, uint256 timestamp)
func (_TitanVPS *TitanVPSCaller) Records(opts *bind.CallOpts, arg0 string) (struct {
	OrderID    string
	BuyerAddr  common.Address
	Price      *big.Int
	Expiration *big.Int
	Timestamp  *big.Int
}, error) {
	var out []interface{}
	err := _TitanVPS.contract.Call(opts, &out, "records", arg0)

	outstruct := new(struct {
		OrderID    string
		BuyerAddr  common.Address
		Price      *big.Int
		Expiration *big.Int
		Timestamp  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderID = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.BuyerAddr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Expiration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Records is a free data retrieval call binding the contract method 0x541e771d.
//
// Solidity: function records(string ) view returns(string orderID, address buyerAddr, uint256 price, uint256 expiration, uint256 timestamp)
func (_TitanVPS *TitanVPSSession) Records(arg0 string) (struct {
	OrderID    string
	BuyerAddr  common.Address
	Price      *big.Int
	Expiration *big.Int
	Timestamp  *big.Int
}, error) {
	return _TitanVPS.Contract.Records(&_TitanVPS.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x541e771d.
//
// Solidity: function records(string ) view returns(string orderID, address buyerAddr, uint256 price, uint256 expiration, uint256 timestamp)
func (_TitanVPS *TitanVPSCallerSession) Records(arg0 string) (struct {
	OrderID    string
	BuyerAddr  common.Address
	Price      *big.Int
	Expiration *big.Int
	Timestamp  *big.Int
}, error) {
	return _TitanVPS.Contract.Records(&_TitanVPS.CallOpts, arg0)
}

// NewOrder is a paid mutator transaction binding the contract method 0x2afa5be8.
//
// Solidity: function newOrder(string _orderID, address _buyerAddr, uint256 _price, uint256 _expiration, uint256 _timestamp) returns()
func (_TitanVPS *TitanVPSTransactor) NewOrder(opts *bind.TransactOpts, _orderID string, _buyerAddr common.Address, _price *big.Int, _expiration *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _TitanVPS.contract.Transact(opts, "newOrder", _orderID, _buyerAddr, _price, _expiration, _timestamp)
}

// NewOrder is a paid mutator transaction binding the contract method 0x2afa5be8.
//
// Solidity: function newOrder(string _orderID, address _buyerAddr, uint256 _price, uint256 _expiration, uint256 _timestamp) returns()
func (_TitanVPS *TitanVPSSession) NewOrder(_orderID string, _buyerAddr common.Address, _price *big.Int, _expiration *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _TitanVPS.Contract.NewOrder(&_TitanVPS.TransactOpts, _orderID, _buyerAddr, _price, _expiration, _timestamp)
}

// NewOrder is a paid mutator transaction binding the contract method 0x2afa5be8.
//
// Solidity: function newOrder(string _orderID, address _buyerAddr, uint256 _price, uint256 _expiration, uint256 _timestamp) returns()
func (_TitanVPS *TitanVPSTransactorSession) NewOrder(_orderID string, _buyerAddr common.Address, _price *big.Int, _expiration *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _TitanVPS.Contract.NewOrder(&_TitanVPS.TransactOpts, _orderID, _buyerAddr, _price, _expiration, _timestamp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TitanVPS *TitanVPSTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TitanVPS.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TitanVPS *TitanVPSSession) RenounceOwnership() (*types.Transaction, error) {
	return _TitanVPS.Contract.RenounceOwnership(&_TitanVPS.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TitanVPS *TitanVPSTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TitanVPS.Contract.RenounceOwnership(&_TitanVPS.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TitanVPS *TitanVPSTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TitanVPS.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TitanVPS *TitanVPSSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TitanVPS.Contract.TransferOwnership(&_TitanVPS.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TitanVPS *TitanVPSTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TitanVPS.Contract.TransferOwnership(&_TitanVPS.TransactOpts, newOwner)
}

// TitanVPSOrderEventIterator is returned from FilterOrderEvent and is used to iterate over the raw logs and unpacked data for OrderEvent events raised by the TitanVPS contract.
type TitanVPSOrderEventIterator struct {
	Event *TitanVPSOrderEvent // Event containing the contract specifics and raw log

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
func (it *TitanVPSOrderEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TitanVPSOrderEvent)
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
		it.Event = new(TitanVPSOrderEvent)
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
func (it *TitanVPSOrderEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TitanVPSOrderEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TitanVPSOrderEvent represents a OrderEvent event raised by the TitanVPS contract.
type TitanVPSOrderEvent struct {
	OrderID    common.Hash
	BuyerAddr  common.Address
	Price      *big.Int
	Expiration *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderEvent is a free log retrieval operation binding the contract event 0xcd7da7825e80a4ce02d420d47c3cf6a3d9625b10b609555c83cf96c0c666149d.
//
// Solidity: event OrderEvent(string indexed orderID, address indexed buyerAddr, uint256 price, uint256 expiration, uint256 timestamp)
func (_TitanVPS *TitanVPSFilterer) FilterOrderEvent(opts *bind.FilterOpts, orderID []string, buyerAddr []common.Address) (*TitanVPSOrderEventIterator, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}
	var buyerAddrRule []interface{}
	for _, buyerAddrItem := range buyerAddr {
		buyerAddrRule = append(buyerAddrRule, buyerAddrItem)
	}

	logs, sub, err := _TitanVPS.contract.FilterLogs(opts, "OrderEvent", orderIDRule, buyerAddrRule)
	if err != nil {
		return nil, err
	}
	return &TitanVPSOrderEventIterator{contract: _TitanVPS.contract, event: "OrderEvent", logs: logs, sub: sub}, nil
}

// WatchOrderEvent is a free log subscription operation binding the contract event 0xcd7da7825e80a4ce02d420d47c3cf6a3d9625b10b609555c83cf96c0c666149d.
//
// Solidity: event OrderEvent(string indexed orderID, address indexed buyerAddr, uint256 price, uint256 expiration, uint256 timestamp)
func (_TitanVPS *TitanVPSFilterer) WatchOrderEvent(opts *bind.WatchOpts, sink chan<- *TitanVPSOrderEvent, orderID []string, buyerAddr []common.Address) (event.Subscription, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}
	var buyerAddrRule []interface{}
	for _, buyerAddrItem := range buyerAddr {
		buyerAddrRule = append(buyerAddrRule, buyerAddrItem)
	}

	logs, sub, err := _TitanVPS.contract.WatchLogs(opts, "OrderEvent", orderIDRule, buyerAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TitanVPSOrderEvent)
				if err := _TitanVPS.contract.UnpackLog(event, "OrderEvent", log); err != nil {
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

// ParseOrderEvent is a log parse operation binding the contract event 0xcd7da7825e80a4ce02d420d47c3cf6a3d9625b10b609555c83cf96c0c666149d.
//
// Solidity: event OrderEvent(string indexed orderID, address indexed buyerAddr, uint256 price, uint256 expiration, uint256 timestamp)
func (_TitanVPS *TitanVPSFilterer) ParseOrderEvent(log types.Log) (*TitanVPSOrderEvent, error) {
	event := new(TitanVPSOrderEvent)
	if err := _TitanVPS.contract.UnpackLog(event, "OrderEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TitanVPSOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TitanVPS contract.
type TitanVPSOwnershipTransferredIterator struct {
	Event *TitanVPSOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TitanVPSOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TitanVPSOwnershipTransferred)
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
		it.Event = new(TitanVPSOwnershipTransferred)
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
func (it *TitanVPSOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TitanVPSOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TitanVPSOwnershipTransferred represents a OwnershipTransferred event raised by the TitanVPS contract.
type TitanVPSOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TitanVPS *TitanVPSFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TitanVPSOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TitanVPS.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TitanVPSOwnershipTransferredIterator{contract: _TitanVPS.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TitanVPS *TitanVPSFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TitanVPSOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TitanVPS.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TitanVPSOwnershipTransferred)
				if err := _TitanVPS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TitanVPS *TitanVPSFilterer) ParseOwnershipTransferred(log types.Log) (*TitanVPSOwnershipTransferred, error) {
	event := new(TitanVPSOwnershipTransferred)
	if err := _TitanVPS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
