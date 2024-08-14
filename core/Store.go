// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package core

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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"AddPool\",\"type\":\"event\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"poolName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"whitelists\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"ownerSetter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"breaker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"breakerSetter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"stop\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"name\":\"AddPoolWhitelist\",\"type\":\"event\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"whitelist\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"Approval\",\"type\":\"event\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"OwnershipTransferred\",\"type\":\"event\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"RemovePoolWhitelist\",\"type\":\"event\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"whitelist\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"ResumeContract\",\"type\":\"event\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"SetPoolBreaker\",\"type\":\"event\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"breaker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"SetPoolBreakerSetter\",\"type\":\"event\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"breakerSetter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"SetPoolOwner\",\"type\":\"event\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"SetPoolOwnerSetter\",\"type\":\"event\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ownerSetter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"StopContract\",\"type\":\"event\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"Transfer\",\"type\":\"event\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"WemixDollarBurned\",\"type\":\"event\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"WemixDollarMinted\",\"type\":\"event\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"addPool\",\"type\":\"function\",\"inputs\":[{\"name\":\"_poolName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_whitelists\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ownerSetter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_breaker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_breakerSetter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stop\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"addPoolWhitelist\",\"type\":\"function\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_whitelist\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"allowance\",\"type\":\"function\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"approve\",\"type\":\"function\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"name\":\"balanceOf\",\"type\":\"function\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"burn\",\"type\":\"function\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"decimals\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"name\":\"decreaseAllowance\",\"type\":\"function\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"name\":\"getWemixDollarInfo\",\"type\":\"function\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"getWemixDollarInfoCount\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"getWhitelistAddress\",\"type\":\"function\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"name\":\"getWhitelistAddressCount\",\"type\":\"function\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"increaseAllowance\",\"type\":\"function\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"name\":\"isWhitelist\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"mint\",\"type\":\"function\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"name\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"name\":\"owner\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"removePoolWhitelist\",\"type\":\"function\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_whitelist\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"renounceOwnership\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"replacePoolWhitelist\",\"type\":\"function\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_whitelist\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_newWhitelist\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"resumeContract\",\"type\":\"function\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"setPoolBreaker\",\"type\":\"function\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_breaker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"setPoolBreakerSetter\",\"type\":\"function\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_breakerSetter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"setPoolOwner\",\"type\":\"function\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"setPoolOwnerSetter\",\"type\":\"function\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_ownerSetter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"stopContract\",\"type\":\"function\",\"inputs\":[{\"name\":\"_poolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"symbol\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"name\":\"totalSupply\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"transfer\",\"type\":\"function\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"name\":\"transferFrom\",\"type\":\"function\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"name\":\"transferOwnership\",\"type\":\"function\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"wemixDollarInfos\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"poolId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"poolName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"poolTotalSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ownerSetter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"breaker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"breakerSetter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stop\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"}]",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Store *StoreCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Store *StoreSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Store.Contract.Allowance(&_Store.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Store *StoreCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Store.Contract.Allowance(&_Store.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Store *StoreCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Store *StoreSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Store *StoreCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Store *StoreCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Store *StoreSession) Decimals() (uint8, error) {
	return _Store.Contract.Decimals(&_Store.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Store *StoreCallerSession) Decimals() (uint8, error) {
	return _Store.Contract.Decimals(&_Store.CallOpts)
}

// GetWemixDollarInfo is a free data retrieval call binding the contract method 0x828298bd.
//
// Solidity: function getWemixDollarInfo(uint256 _poolId) view returns(uint256, string, uint256, address[], address, address, address, address, bool)
func (_Store *StoreCaller) GetWemixDollarInfo(opts *bind.CallOpts, _poolId *big.Int) (*big.Int, string, *big.Int, []common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getWemixDollarInfo", _poolId)

	if err != nil {
		return *new(*big.Int), *new(string), *new(*big.Int), *new([]common.Address), *new(common.Address), *new(common.Address), *new(common.Address), *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	out6 := *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	out7 := *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	out8 := *abi.ConvertType(out[8], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetWemixDollarInfo is a free data retrieval call binding the contract method 0x828298bd.
//
// Solidity: function getWemixDollarInfo(uint256 _poolId) view returns(uint256, string, uint256, address[], address, address, address, address, bool)
func (_Store *StoreSession) GetWemixDollarInfo(_poolId *big.Int) (*big.Int, string, *big.Int, []common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	return _Store.Contract.GetWemixDollarInfo(&_Store.CallOpts, _poolId)
}

// GetWemixDollarInfo is a free data retrieval call binding the contract method 0x828298bd.
//
// Solidity: function getWemixDollarInfo(uint256 _poolId) view returns(uint256, string, uint256, address[], address, address, address, address, bool)
func (_Store *StoreCallerSession) GetWemixDollarInfo(_poolId *big.Int) (*big.Int, string, *big.Int, []common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	return _Store.Contract.GetWemixDollarInfo(&_Store.CallOpts, _poolId)
}

// GetWemixDollarInfoCount is a free data retrieval call binding the contract method 0xae0ce498.
//
// Solidity: function getWemixDollarInfoCount() view returns(uint256)
func (_Store *StoreCaller) GetWemixDollarInfoCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getWemixDollarInfoCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWemixDollarInfoCount is a free data retrieval call binding the contract method 0xae0ce498.
//
// Solidity: function getWemixDollarInfoCount() view returns(uint256)
func (_Store *StoreSession) GetWemixDollarInfoCount() (*big.Int, error) {
	return _Store.Contract.GetWemixDollarInfoCount(&_Store.CallOpts)
}

// GetWemixDollarInfoCount is a free data retrieval call binding the contract method 0xae0ce498.
//
// Solidity: function getWemixDollarInfoCount() view returns(uint256)
func (_Store *StoreCallerSession) GetWemixDollarInfoCount() (*big.Int, error) {
	return _Store.Contract.GetWemixDollarInfoCount(&_Store.CallOpts)
}

// GetWhitelistAddress is a free data retrieval call binding the contract method 0x565853af.
//
// Solidity: function getWhitelistAddress(uint256 _poolId) view returns(address[])
func (_Store *StoreCaller) GetWhitelistAddress(opts *bind.CallOpts, _poolId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getWhitelistAddress", _poolId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWhitelistAddress is a free data retrieval call binding the contract method 0x565853af.
//
// Solidity: function getWhitelistAddress(uint256 _poolId) view returns(address[])
func (_Store *StoreSession) GetWhitelistAddress(_poolId *big.Int) ([]common.Address, error) {
	return _Store.Contract.GetWhitelistAddress(&_Store.CallOpts, _poolId)
}

// GetWhitelistAddress is a free data retrieval call binding the contract method 0x565853af.
//
// Solidity: function getWhitelistAddress(uint256 _poolId) view returns(address[])
func (_Store *StoreCallerSession) GetWhitelistAddress(_poolId *big.Int) ([]common.Address, error) {
	return _Store.Contract.GetWhitelistAddress(&_Store.CallOpts, _poolId)
}

// GetWhitelistAddressCount is a free data retrieval call binding the contract method 0x9c63739a.
//
// Solidity: function getWhitelistAddressCount(uint256 _poolId) view returns(uint256)
func (_Store *StoreCaller) GetWhitelistAddressCount(opts *bind.CallOpts, _poolId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getWhitelistAddressCount", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWhitelistAddressCount is a free data retrieval call binding the contract method 0x9c63739a.
//
// Solidity: function getWhitelistAddressCount(uint256 _poolId) view returns(uint256)
func (_Store *StoreSession) GetWhitelistAddressCount(_poolId *big.Int) (*big.Int, error) {
	return _Store.Contract.GetWhitelistAddressCount(&_Store.CallOpts, _poolId)
}

// GetWhitelistAddressCount is a free data retrieval call binding the contract method 0x9c63739a.
//
// Solidity: function getWhitelistAddressCount(uint256 _poolId) view returns(uint256)
func (_Store *StoreCallerSession) GetWhitelistAddressCount(_poolId *big.Int) (*big.Int, error) {
	return _Store.Contract.GetWhitelistAddressCount(&_Store.CallOpts, _poolId)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(uint256)
func (_Store *StoreCaller) IsWhitelist(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "isWhitelist", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(uint256)
func (_Store *StoreSession) IsWhitelist(arg0 common.Address) (*big.Int, error) {
	return _Store.Contract.IsWhitelist(&_Store.CallOpts, arg0)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(uint256)
func (_Store *StoreCallerSession) IsWhitelist(arg0 common.Address) (*big.Int, error) {
	return _Store.Contract.IsWhitelist(&_Store.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreCallerSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCallerSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreCallerSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Store *StoreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Store *StoreSession) TotalSupply() (*big.Int, error) {
	return _Store.Contract.TotalSupply(&_Store.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Store *StoreCallerSession) TotalSupply() (*big.Int, error) {
	return _Store.Contract.TotalSupply(&_Store.CallOpts)
}

// WemixDollarInfos is a free data retrieval call binding the contract method 0x2acabf3e.
//
// Solidity: function wemixDollarInfos(uint256 ) view returns(uint256 poolId, string poolName, uint256 poolTotalSupply, address owner, address ownerSetter, address breaker, address breakerSetter, bool stop)
func (_Store *StoreCaller) WemixDollarInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PoolId          *big.Int
	PoolName        string
	PoolTotalSupply *big.Int
	Owner           common.Address
	OwnerSetter     common.Address
	Breaker         common.Address
	BreakerSetter   common.Address
	Stop            bool
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "wemixDollarInfos", arg0)

	outstruct := new(struct {
		PoolId          *big.Int
		PoolName        string
		PoolTotalSupply *big.Int
		Owner           common.Address
		OwnerSetter     common.Address
		Breaker         common.Address
		BreakerSetter   common.Address
		Stop            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PoolId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PoolName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.PoolTotalSupply = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.OwnerSetter = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Breaker = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.BreakerSetter = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Stop = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// WemixDollarInfos is a free data retrieval call binding the contract method 0x2acabf3e.
//
// Solidity: function wemixDollarInfos(uint256 ) view returns(uint256 poolId, string poolName, uint256 poolTotalSupply, address owner, address ownerSetter, address breaker, address breakerSetter, bool stop)
func (_Store *StoreSession) WemixDollarInfos(arg0 *big.Int) (struct {
	PoolId          *big.Int
	PoolName        string
	PoolTotalSupply *big.Int
	Owner           common.Address
	OwnerSetter     common.Address
	Breaker         common.Address
	BreakerSetter   common.Address
	Stop            bool
}, error) {
	return _Store.Contract.WemixDollarInfos(&_Store.CallOpts, arg0)
}

// WemixDollarInfos is a free data retrieval call binding the contract method 0x2acabf3e.
//
// Solidity: function wemixDollarInfos(uint256 ) view returns(uint256 poolId, string poolName, uint256 poolTotalSupply, address owner, address ownerSetter, address breaker, address breakerSetter, bool stop)
func (_Store *StoreCallerSession) WemixDollarInfos(arg0 *big.Int) (struct {
	PoolId          *big.Int
	PoolName        string
	PoolTotalSupply *big.Int
	Owner           common.Address
	OwnerSetter     common.Address
	Breaker         common.Address
	BreakerSetter   common.Address
	Stop            bool
}, error) {
	return _Store.Contract.WemixDollarInfos(&_Store.CallOpts, arg0)
}

// AddPool is a paid mutator transaction binding the contract method 0x03d19032.
//
// Solidity: function addPool(string _poolName, address[] _whitelists, address _owner, address _ownerSetter, address _breaker, address _breakerSetter, bool _stop) returns()
func (_Store *StoreTransactor) AddPool(opts *bind.TransactOpts, _poolName string, _whitelists []common.Address, _owner common.Address, _ownerSetter common.Address, _breaker common.Address, _breakerSetter common.Address, _stop bool) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addPool", _poolName, _whitelists, _owner, _ownerSetter, _breaker, _breakerSetter, _stop)
}

// AddPool is a paid mutator transaction binding the contract method 0x03d19032.
//
// Solidity: function addPool(string _poolName, address[] _whitelists, address _owner, address _ownerSetter, address _breaker, address _breakerSetter, bool _stop) returns()
func (_Store *StoreSession) AddPool(_poolName string, _whitelists []common.Address, _owner common.Address, _ownerSetter common.Address, _breaker common.Address, _breakerSetter common.Address, _stop bool) (*types.Transaction, error) {
	return _Store.Contract.AddPool(&_Store.TransactOpts, _poolName, _whitelists, _owner, _ownerSetter, _breaker, _breakerSetter, _stop)
}

// AddPool is a paid mutator transaction binding the contract method 0x03d19032.
//
// Solidity: function addPool(string _poolName, address[] _whitelists, address _owner, address _ownerSetter, address _breaker, address _breakerSetter, bool _stop) returns()
func (_Store *StoreTransactorSession) AddPool(_poolName string, _whitelists []common.Address, _owner common.Address, _ownerSetter common.Address, _breaker common.Address, _breakerSetter common.Address, _stop bool) (*types.Transaction, error) {
	return _Store.Contract.AddPool(&_Store.TransactOpts, _poolName, _whitelists, _owner, _ownerSetter, _breaker, _breakerSetter, _stop)
}

// AddPoolWhitelist is a paid mutator transaction binding the contract method 0x3fab3b91.
//
// Solidity: function addPoolWhitelist(uint256 _poolId, address _whitelist) returns()
func (_Store *StoreTransactor) AddPoolWhitelist(opts *bind.TransactOpts, _poolId *big.Int, _whitelist common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addPoolWhitelist", _poolId, _whitelist)
}

// AddPoolWhitelist is a paid mutator transaction binding the contract method 0x3fab3b91.
//
// Solidity: function addPoolWhitelist(uint256 _poolId, address _whitelist) returns()
func (_Store *StoreSession) AddPoolWhitelist(_poolId *big.Int, _whitelist common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddPoolWhitelist(&_Store.TransactOpts, _poolId, _whitelist)
}

// AddPoolWhitelist is a paid mutator transaction binding the contract method 0x3fab3b91.
//
// Solidity: function addPoolWhitelist(uint256 _poolId, address _whitelist) returns()
func (_Store *StoreTransactorSession) AddPoolWhitelist(_poolId *big.Int, _whitelist common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddPoolWhitelist(&_Store.TransactOpts, _poolId, _whitelist)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Store *StoreTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Store *StoreSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Store *StoreTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_Store *StoreTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_Store *StoreSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Burn(&_Store.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_Store *StoreTransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Burn(&_Store.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Store *StoreTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Store *StoreSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DecreaseAllowance(&_Store.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Store *StoreTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DecreaseAllowance(&_Store.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Store *StoreTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Store *StoreSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.IncreaseAllowance(&_Store.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Store *StoreTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.IncreaseAllowance(&_Store.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Store *StoreTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Store *StoreSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Mint(&_Store.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Store *StoreTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Mint(&_Store.TransactOpts, _to, _amount)
}

// RemovePoolWhitelist is a paid mutator transaction binding the contract method 0xec4f052a.
//
// Solidity: function removePoolWhitelist(uint256 _poolId, address _whitelist) returns()
func (_Store *StoreTransactor) RemovePoolWhitelist(opts *bind.TransactOpts, _poolId *big.Int, _whitelist common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "removePoolWhitelist", _poolId, _whitelist)
}

// RemovePoolWhitelist is a paid mutator transaction binding the contract method 0xec4f052a.
//
// Solidity: function removePoolWhitelist(uint256 _poolId, address _whitelist) returns()
func (_Store *StoreSession) RemovePoolWhitelist(_poolId *big.Int, _whitelist common.Address) (*types.Transaction, error) {
	return _Store.Contract.RemovePoolWhitelist(&_Store.TransactOpts, _poolId, _whitelist)
}

// RemovePoolWhitelist is a paid mutator transaction binding the contract method 0xec4f052a.
//
// Solidity: function removePoolWhitelist(uint256 _poolId, address _whitelist) returns()
func (_Store *StoreTransactorSession) RemovePoolWhitelist(_poolId *big.Int, _whitelist common.Address) (*types.Transaction, error) {
	return _Store.Contract.RemovePoolWhitelist(&_Store.TransactOpts, _poolId, _whitelist)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _Store.Contract.RenounceOwnership(&_Store.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Store.Contract.RenounceOwnership(&_Store.TransactOpts)
}

// ReplacePoolWhitelist is a paid mutator transaction binding the contract method 0x874e5299.
//
// Solidity: function replacePoolWhitelist(uint256 _poolId, address _whitelist, address _newWhitelist) returns()
func (_Store *StoreTransactor) ReplacePoolWhitelist(opts *bind.TransactOpts, _poolId *big.Int, _whitelist common.Address, _newWhitelist common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "replacePoolWhitelist", _poolId, _whitelist, _newWhitelist)
}

// ReplacePoolWhitelist is a paid mutator transaction binding the contract method 0x874e5299.
//
// Solidity: function replacePoolWhitelist(uint256 _poolId, address _whitelist, address _newWhitelist) returns()
func (_Store *StoreSession) ReplacePoolWhitelist(_poolId *big.Int, _whitelist common.Address, _newWhitelist common.Address) (*types.Transaction, error) {
	return _Store.Contract.ReplacePoolWhitelist(&_Store.TransactOpts, _poolId, _whitelist, _newWhitelist)
}

// ReplacePoolWhitelist is a paid mutator transaction binding the contract method 0x874e5299.
//
// Solidity: function replacePoolWhitelist(uint256 _poolId, address _whitelist, address _newWhitelist) returns()
func (_Store *StoreTransactorSession) ReplacePoolWhitelist(_poolId *big.Int, _whitelist common.Address, _newWhitelist common.Address) (*types.Transaction, error) {
	return _Store.Contract.ReplacePoolWhitelist(&_Store.TransactOpts, _poolId, _whitelist, _newWhitelist)
}

// ResumeContract is a paid mutator transaction binding the contract method 0xa78e2f2d.
//
// Solidity: function resumeContract(uint256 _poolId) returns()
func (_Store *StoreTransactor) ResumeContract(opts *bind.TransactOpts, _poolId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "resumeContract", _poolId)
}

// ResumeContract is a paid mutator transaction binding the contract method 0xa78e2f2d.
//
// Solidity: function resumeContract(uint256 _poolId) returns()
func (_Store *StoreSession) ResumeContract(_poolId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ResumeContract(&_Store.TransactOpts, _poolId)
}

// ResumeContract is a paid mutator transaction binding the contract method 0xa78e2f2d.
//
// Solidity: function resumeContract(uint256 _poolId) returns()
func (_Store *StoreTransactorSession) ResumeContract(_poolId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ResumeContract(&_Store.TransactOpts, _poolId)
}

// SetPoolBreaker is a paid mutator transaction binding the contract method 0x7d3824d9.
//
// Solidity: function setPoolBreaker(uint256 _poolId, address _breaker) returns()
func (_Store *StoreTransactor) SetPoolBreaker(opts *bind.TransactOpts, _poolId *big.Int, _breaker common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setPoolBreaker", _poolId, _breaker)
}

// SetPoolBreaker is a paid mutator transaction binding the contract method 0x7d3824d9.
//
// Solidity: function setPoolBreaker(uint256 _poolId, address _breaker) returns()
func (_Store *StoreSession) SetPoolBreaker(_poolId *big.Int, _breaker common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetPoolBreaker(&_Store.TransactOpts, _poolId, _breaker)
}

// SetPoolBreaker is a paid mutator transaction binding the contract method 0x7d3824d9.
//
// Solidity: function setPoolBreaker(uint256 _poolId, address _breaker) returns()
func (_Store *StoreTransactorSession) SetPoolBreaker(_poolId *big.Int, _breaker common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetPoolBreaker(&_Store.TransactOpts, _poolId, _breaker)
}

// SetPoolBreakerSetter is a paid mutator transaction binding the contract method 0x6365530e.
//
// Solidity: function setPoolBreakerSetter(uint256 _poolId, address _breakerSetter) returns()
func (_Store *StoreTransactor) SetPoolBreakerSetter(opts *bind.TransactOpts, _poolId *big.Int, _breakerSetter common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setPoolBreakerSetter", _poolId, _breakerSetter)
}

// SetPoolBreakerSetter is a paid mutator transaction binding the contract method 0x6365530e.
//
// Solidity: function setPoolBreakerSetter(uint256 _poolId, address _breakerSetter) returns()
func (_Store *StoreSession) SetPoolBreakerSetter(_poolId *big.Int, _breakerSetter common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetPoolBreakerSetter(&_Store.TransactOpts, _poolId, _breakerSetter)
}

// SetPoolBreakerSetter is a paid mutator transaction binding the contract method 0x6365530e.
//
// Solidity: function setPoolBreakerSetter(uint256 _poolId, address _breakerSetter) returns()
func (_Store *StoreTransactorSession) SetPoolBreakerSetter(_poolId *big.Int, _breakerSetter common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetPoolBreakerSetter(&_Store.TransactOpts, _poolId, _breakerSetter)
}

// SetPoolOwner is a paid mutator transaction binding the contract method 0x52202b0a.
//
// Solidity: function setPoolOwner(uint256 _poolId, address _owner) returns()
func (_Store *StoreTransactor) SetPoolOwner(opts *bind.TransactOpts, _poolId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setPoolOwner", _poolId, _owner)
}

// SetPoolOwner is a paid mutator transaction binding the contract method 0x52202b0a.
//
// Solidity: function setPoolOwner(uint256 _poolId, address _owner) returns()
func (_Store *StoreSession) SetPoolOwner(_poolId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetPoolOwner(&_Store.TransactOpts, _poolId, _owner)
}

// SetPoolOwner is a paid mutator transaction binding the contract method 0x52202b0a.
//
// Solidity: function setPoolOwner(uint256 _poolId, address _owner) returns()
func (_Store *StoreTransactorSession) SetPoolOwner(_poolId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetPoolOwner(&_Store.TransactOpts, _poolId, _owner)
}

// SetPoolOwnerSetter is a paid mutator transaction binding the contract method 0x50f63fe7.
//
// Solidity: function setPoolOwnerSetter(uint256 _poolId, address _ownerSetter) returns()
func (_Store *StoreTransactor) SetPoolOwnerSetter(opts *bind.TransactOpts, _poolId *big.Int, _ownerSetter common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setPoolOwnerSetter", _poolId, _ownerSetter)
}

// SetPoolOwnerSetter is a paid mutator transaction binding the contract method 0x50f63fe7.
//
// Solidity: function setPoolOwnerSetter(uint256 _poolId, address _ownerSetter) returns()
func (_Store *StoreSession) SetPoolOwnerSetter(_poolId *big.Int, _ownerSetter common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetPoolOwnerSetter(&_Store.TransactOpts, _poolId, _ownerSetter)
}

// SetPoolOwnerSetter is a paid mutator transaction binding the contract method 0x50f63fe7.
//
// Solidity: function setPoolOwnerSetter(uint256 _poolId, address _ownerSetter) returns()
func (_Store *StoreTransactorSession) SetPoolOwnerSetter(_poolId *big.Int, _ownerSetter common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetPoolOwnerSetter(&_Store.TransactOpts, _poolId, _ownerSetter)
}

// StopContract is a paid mutator transaction binding the contract method 0xbc59d503.
//
// Solidity: function stopContract(uint256 _poolId) returns()
func (_Store *StoreTransactor) StopContract(opts *bind.TransactOpts, _poolId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "stopContract", _poolId)
}

// StopContract is a paid mutator transaction binding the contract method 0xbc59d503.
//
// Solidity: function stopContract(uint256 _poolId) returns()
func (_Store *StoreSession) StopContract(_poolId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.StopContract(&_Store.TransactOpts, _poolId)
}

// StopContract is a paid mutator transaction binding the contract method 0xbc59d503.
//
// Solidity: function stopContract(uint256 _poolId) returns()
func (_Store *StoreTransactorSession) StopContract(_poolId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.StopContract(&_Store.TransactOpts, _poolId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Store *StoreTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Store *StoreSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Transfer(&_Store.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Store *StoreTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Transfer(&_Store.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Store *StoreTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Store *StoreSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Store *StoreTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Store *StoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Store *StoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Store *StoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// StoreAddPoolIterator is returned from FilterAddPool and is used to iterate over the raw logs and unpacked data for AddPool events raised by the Store contract.
type StoreAddPoolIterator struct {
	Event *StoreAddPool // Event containing the contract specifics and raw log

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
func (it *StoreAddPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreAddPool)
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
		it.Event = new(StoreAddPool)
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
func (it *StoreAddPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreAddPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreAddPool represents a AddPool event raised by the Store contract.
type StoreAddPool struct {
	PoolId        *big.Int
	PoolName      string
	Whitelists    []common.Address
	Owner         common.Address
	OwnerSetter   common.Address
	Breaker       common.Address
	BreakerSetter common.Address
	Stop          bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddPool is a free log retrieval operation binding the contract event 0x9329e5b15842a818b22b4f72136262697e6a2acb38fabe63f72d9dd079a7bbff.
//
// Solidity: event AddPool(uint256 poolId, string poolName, address[] whitelists, address owner, address ownerSetter, address breaker, address breakerSetter, bool stop)
func (_Store *StoreFilterer) FilterAddPool(opts *bind.FilterOpts) (*StoreAddPoolIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "AddPool")
	if err != nil {
		return nil, err
	}
	return &StoreAddPoolIterator{contract: _Store.contract, event: "AddPool", logs: logs, sub: sub}, nil
}

// WatchAddPool is a free log subscription operation binding the contract event 0x9329e5b15842a818b22b4f72136262697e6a2acb38fabe63f72d9dd079a7bbff.
//
// Solidity: event AddPool(uint256 poolId, string poolName, address[] whitelists, address owner, address ownerSetter, address breaker, address breakerSetter, bool stop)
func (_Store *StoreFilterer) WatchAddPool(opts *bind.WatchOpts, sink chan<- *StoreAddPool) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "AddPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreAddPool)
				if err := _Store.contract.UnpackLog(event, "AddPool", log); err != nil {
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

// ParseAddPool is a log parse operation binding the contract event 0x9329e5b15842a818b22b4f72136262697e6a2acb38fabe63f72d9dd079a7bbff.
//
// Solidity: event AddPool(uint256 poolId, string poolName, address[] whitelists, address owner, address ownerSetter, address breaker, address breakerSetter, bool stop)
func (_Store *StoreFilterer) ParseAddPool(log types.Log) (*StoreAddPool, error) {
	event := new(StoreAddPool)
	if err := _Store.contract.UnpackLog(event, "AddPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreAddPoolWhitelistIterator is returned from FilterAddPoolWhitelist and is used to iterate over the raw logs and unpacked data for AddPoolWhitelist events raised by the Store contract.
type StoreAddPoolWhitelistIterator struct {
	Event *StoreAddPoolWhitelist // Event containing the contract specifics and raw log

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
func (it *StoreAddPoolWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreAddPoolWhitelist)
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
		it.Event = new(StoreAddPoolWhitelist)
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
func (it *StoreAddPoolWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreAddPoolWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreAddPoolWhitelist represents a AddPoolWhitelist event raised by the Store contract.
type StoreAddPoolWhitelist struct {
	PoolId    *big.Int
	Whitelist common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddPoolWhitelist is a free log retrieval operation binding the contract event 0x5f7bf5a6535d2f3b9b2649a02c1afaa7b46a85ac2bc0bac26dc979747762b96b.
//
// Solidity: event AddPoolWhitelist(uint256 poolId, address whitelist)
func (_Store *StoreFilterer) FilterAddPoolWhitelist(opts *bind.FilterOpts) (*StoreAddPoolWhitelistIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "AddPoolWhitelist")
	if err != nil {
		return nil, err
	}
	return &StoreAddPoolWhitelistIterator{contract: _Store.contract, event: "AddPoolWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddPoolWhitelist is a free log subscription operation binding the contract event 0x5f7bf5a6535d2f3b9b2649a02c1afaa7b46a85ac2bc0bac26dc979747762b96b.
//
// Solidity: event AddPoolWhitelist(uint256 poolId, address whitelist)
func (_Store *StoreFilterer) WatchAddPoolWhitelist(opts *bind.WatchOpts, sink chan<- *StoreAddPoolWhitelist) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "AddPoolWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreAddPoolWhitelist)
				if err := _Store.contract.UnpackLog(event, "AddPoolWhitelist", log); err != nil {
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

// ParseAddPoolWhitelist is a log parse operation binding the contract event 0x5f7bf5a6535d2f3b9b2649a02c1afaa7b46a85ac2bc0bac26dc979747762b96b.
//
// Solidity: event AddPoolWhitelist(uint256 poolId, address whitelist)
func (_Store *StoreFilterer) ParseAddPoolWhitelist(log types.Log) (*StoreAddPoolWhitelist, error) {
	event := new(StoreAddPoolWhitelist)
	if err := _Store.contract.UnpackLog(event, "AddPoolWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Store contract.
type StoreApprovalIterator struct {
	Event *StoreApproval // Event containing the contract specifics and raw log

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
func (it *StoreApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreApproval)
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
		it.Event = new(StoreApproval)
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
func (it *StoreApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreApproval represents a Approval event raised by the Store contract.
type StoreApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Store *StoreFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StoreApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StoreApprovalIterator{contract: _Store.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Store *StoreFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StoreApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreApproval)
				if err := _Store.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Store *StoreFilterer) ParseApproval(log types.Log) (*StoreApproval, error) {
	event := new(StoreApproval)
	if err := _Store.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Store contract.
type StoreOwnershipTransferredIterator struct {
	Event *StoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOwnershipTransferred)
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
		it.Event = new(StoreOwnershipTransferred)
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
func (it *StoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOwnershipTransferred represents a OwnershipTransferred event raised by the Store contract.
type StoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Store *StoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StoreOwnershipTransferredIterator{contract: _Store.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Store *StoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOwnershipTransferred)
				if err := _Store.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Store *StoreFilterer) ParseOwnershipTransferred(log types.Log) (*StoreOwnershipTransferred, error) {
	event := new(StoreOwnershipTransferred)
	if err := _Store.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreRemovePoolWhitelistIterator is returned from FilterRemovePoolWhitelist and is used to iterate over the raw logs and unpacked data for RemovePoolWhitelist events raised by the Store contract.
type StoreRemovePoolWhitelistIterator struct {
	Event *StoreRemovePoolWhitelist // Event containing the contract specifics and raw log

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
func (it *StoreRemovePoolWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreRemovePoolWhitelist)
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
		it.Event = new(StoreRemovePoolWhitelist)
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
func (it *StoreRemovePoolWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreRemovePoolWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreRemovePoolWhitelist represents a RemovePoolWhitelist event raised by the Store contract.
type StoreRemovePoolWhitelist struct {
	PoolId    *big.Int
	Whitelist common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemovePoolWhitelist is a free log retrieval operation binding the contract event 0xe1938c7173044f254b823afeeca186e2d18319fb30e37e0e4de6277a52d29106.
//
// Solidity: event RemovePoolWhitelist(uint256 poolId, address whitelist)
func (_Store *StoreFilterer) FilterRemovePoolWhitelist(opts *bind.FilterOpts) (*StoreRemovePoolWhitelistIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "RemovePoolWhitelist")
	if err != nil {
		return nil, err
	}
	return &StoreRemovePoolWhitelistIterator{contract: _Store.contract, event: "RemovePoolWhitelist", logs: logs, sub: sub}, nil
}

// WatchRemovePoolWhitelist is a free log subscription operation binding the contract event 0xe1938c7173044f254b823afeeca186e2d18319fb30e37e0e4de6277a52d29106.
//
// Solidity: event RemovePoolWhitelist(uint256 poolId, address whitelist)
func (_Store *StoreFilterer) WatchRemovePoolWhitelist(opts *bind.WatchOpts, sink chan<- *StoreRemovePoolWhitelist) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "RemovePoolWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreRemovePoolWhitelist)
				if err := _Store.contract.UnpackLog(event, "RemovePoolWhitelist", log); err != nil {
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

// ParseRemovePoolWhitelist is a log parse operation binding the contract event 0xe1938c7173044f254b823afeeca186e2d18319fb30e37e0e4de6277a52d29106.
//
// Solidity: event RemovePoolWhitelist(uint256 poolId, address whitelist)
func (_Store *StoreFilterer) ParseRemovePoolWhitelist(log types.Log) (*StoreRemovePoolWhitelist, error) {
	event := new(StoreRemovePoolWhitelist)
	if err := _Store.contract.UnpackLog(event, "RemovePoolWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreResumeContractIterator is returned from FilterResumeContract and is used to iterate over the raw logs and unpacked data for ResumeContract events raised by the Store contract.
type StoreResumeContractIterator struct {
	Event *StoreResumeContract // Event containing the contract specifics and raw log

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
func (it *StoreResumeContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreResumeContract)
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
		it.Event = new(StoreResumeContract)
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
func (it *StoreResumeContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreResumeContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreResumeContract represents a ResumeContract event raised by the Store contract.
type StoreResumeContract struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResumeContract is a free log retrieval operation binding the contract event 0xd3b27aa7927629b0fe961a7f5ac65c9a1dd990a0bdfa1962108af3ab4aa22d8a.
//
// Solidity: event ResumeContract(uint256 poolId)
func (_Store *StoreFilterer) FilterResumeContract(opts *bind.FilterOpts) (*StoreResumeContractIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ResumeContract")
	if err != nil {
		return nil, err
	}
	return &StoreResumeContractIterator{contract: _Store.contract, event: "ResumeContract", logs: logs, sub: sub}, nil
}

// WatchResumeContract is a free log subscription operation binding the contract event 0xd3b27aa7927629b0fe961a7f5ac65c9a1dd990a0bdfa1962108af3ab4aa22d8a.
//
// Solidity: event ResumeContract(uint256 poolId)
func (_Store *StoreFilterer) WatchResumeContract(opts *bind.WatchOpts, sink chan<- *StoreResumeContract) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ResumeContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreResumeContract)
				if err := _Store.contract.UnpackLog(event, "ResumeContract", log); err != nil {
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

// ParseResumeContract is a log parse operation binding the contract event 0xd3b27aa7927629b0fe961a7f5ac65c9a1dd990a0bdfa1962108af3ab4aa22d8a.
//
// Solidity: event ResumeContract(uint256 poolId)
func (_Store *StoreFilterer) ParseResumeContract(log types.Log) (*StoreResumeContract, error) {
	event := new(StoreResumeContract)
	if err := _Store.contract.UnpackLog(event, "ResumeContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreSetPoolBreakerIterator is returned from FilterSetPoolBreaker and is used to iterate over the raw logs and unpacked data for SetPoolBreaker events raised by the Store contract.
type StoreSetPoolBreakerIterator struct {
	Event *StoreSetPoolBreaker // Event containing the contract specifics and raw log

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
func (it *StoreSetPoolBreakerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreSetPoolBreaker)
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
		it.Event = new(StoreSetPoolBreaker)
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
func (it *StoreSetPoolBreakerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreSetPoolBreakerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreSetPoolBreaker represents a SetPoolBreaker event raised by the Store contract.
type StoreSetPoolBreaker struct {
	PoolId  *big.Int
	Breaker common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetPoolBreaker is a free log retrieval operation binding the contract event 0x70d216e53bba32c29a3424ecf8aa15aad290eb6525997dd33ea324dfd9cd5339.
//
// Solidity: event SetPoolBreaker(uint256 poolId, address breaker)
func (_Store *StoreFilterer) FilterSetPoolBreaker(opts *bind.FilterOpts) (*StoreSetPoolBreakerIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "SetPoolBreaker")
	if err != nil {
		return nil, err
	}
	return &StoreSetPoolBreakerIterator{contract: _Store.contract, event: "SetPoolBreaker", logs: logs, sub: sub}, nil
}

// WatchSetPoolBreaker is a free log subscription operation binding the contract event 0x70d216e53bba32c29a3424ecf8aa15aad290eb6525997dd33ea324dfd9cd5339.
//
// Solidity: event SetPoolBreaker(uint256 poolId, address breaker)
func (_Store *StoreFilterer) WatchSetPoolBreaker(opts *bind.WatchOpts, sink chan<- *StoreSetPoolBreaker) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "SetPoolBreaker")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreSetPoolBreaker)
				if err := _Store.contract.UnpackLog(event, "SetPoolBreaker", log); err != nil {
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

// ParseSetPoolBreaker is a log parse operation binding the contract event 0x70d216e53bba32c29a3424ecf8aa15aad290eb6525997dd33ea324dfd9cd5339.
//
// Solidity: event SetPoolBreaker(uint256 poolId, address breaker)
func (_Store *StoreFilterer) ParseSetPoolBreaker(log types.Log) (*StoreSetPoolBreaker, error) {
	event := new(StoreSetPoolBreaker)
	if err := _Store.contract.UnpackLog(event, "SetPoolBreaker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreSetPoolBreakerSetterIterator is returned from FilterSetPoolBreakerSetter and is used to iterate over the raw logs and unpacked data for SetPoolBreakerSetter events raised by the Store contract.
type StoreSetPoolBreakerSetterIterator struct {
	Event *StoreSetPoolBreakerSetter // Event containing the contract specifics and raw log

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
func (it *StoreSetPoolBreakerSetterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreSetPoolBreakerSetter)
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
		it.Event = new(StoreSetPoolBreakerSetter)
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
func (it *StoreSetPoolBreakerSetterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreSetPoolBreakerSetterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreSetPoolBreakerSetter represents a SetPoolBreakerSetter event raised by the Store contract.
type StoreSetPoolBreakerSetter struct {
	PoolId        *big.Int
	BreakerSetter common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetPoolBreakerSetter is a free log retrieval operation binding the contract event 0x7fc7b82c46c9f322f2a9c3ab999a39ba8230b89e3173adfde268967405ebb2df.
//
// Solidity: event SetPoolBreakerSetter(uint256 poolId, address breakerSetter)
func (_Store *StoreFilterer) FilterSetPoolBreakerSetter(opts *bind.FilterOpts) (*StoreSetPoolBreakerSetterIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "SetPoolBreakerSetter")
	if err != nil {
		return nil, err
	}
	return &StoreSetPoolBreakerSetterIterator{contract: _Store.contract, event: "SetPoolBreakerSetter", logs: logs, sub: sub}, nil
}

// WatchSetPoolBreakerSetter is a free log subscription operation binding the contract event 0x7fc7b82c46c9f322f2a9c3ab999a39ba8230b89e3173adfde268967405ebb2df.
//
// Solidity: event SetPoolBreakerSetter(uint256 poolId, address breakerSetter)
func (_Store *StoreFilterer) WatchSetPoolBreakerSetter(opts *bind.WatchOpts, sink chan<- *StoreSetPoolBreakerSetter) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "SetPoolBreakerSetter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreSetPoolBreakerSetter)
				if err := _Store.contract.UnpackLog(event, "SetPoolBreakerSetter", log); err != nil {
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

// ParseSetPoolBreakerSetter is a log parse operation binding the contract event 0x7fc7b82c46c9f322f2a9c3ab999a39ba8230b89e3173adfde268967405ebb2df.
//
// Solidity: event SetPoolBreakerSetter(uint256 poolId, address breakerSetter)
func (_Store *StoreFilterer) ParseSetPoolBreakerSetter(log types.Log) (*StoreSetPoolBreakerSetter, error) {
	event := new(StoreSetPoolBreakerSetter)
	if err := _Store.contract.UnpackLog(event, "SetPoolBreakerSetter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreSetPoolOwnerIterator is returned from FilterSetPoolOwner and is used to iterate over the raw logs and unpacked data for SetPoolOwner events raised by the Store contract.
type StoreSetPoolOwnerIterator struct {
	Event *StoreSetPoolOwner // Event containing the contract specifics and raw log

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
func (it *StoreSetPoolOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreSetPoolOwner)
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
		it.Event = new(StoreSetPoolOwner)
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
func (it *StoreSetPoolOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreSetPoolOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreSetPoolOwner represents a SetPoolOwner event raised by the Store contract.
type StoreSetPoolOwner struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPoolOwner is a free log retrieval operation binding the contract event 0xf87e07c03fe1220193f248d7937e558545020f9304702037cc5412451abfc363.
//
// Solidity: event SetPoolOwner(uint256 poolId, address owner)
func (_Store *StoreFilterer) FilterSetPoolOwner(opts *bind.FilterOpts) (*StoreSetPoolOwnerIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "SetPoolOwner")
	if err != nil {
		return nil, err
	}
	return &StoreSetPoolOwnerIterator{contract: _Store.contract, event: "SetPoolOwner", logs: logs, sub: sub}, nil
}

// WatchSetPoolOwner is a free log subscription operation binding the contract event 0xf87e07c03fe1220193f248d7937e558545020f9304702037cc5412451abfc363.
//
// Solidity: event SetPoolOwner(uint256 poolId, address owner)
func (_Store *StoreFilterer) WatchSetPoolOwner(opts *bind.WatchOpts, sink chan<- *StoreSetPoolOwner) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "SetPoolOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreSetPoolOwner)
				if err := _Store.contract.UnpackLog(event, "SetPoolOwner", log); err != nil {
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

// ParseSetPoolOwner is a log parse operation binding the contract event 0xf87e07c03fe1220193f248d7937e558545020f9304702037cc5412451abfc363.
//
// Solidity: event SetPoolOwner(uint256 poolId, address owner)
func (_Store *StoreFilterer) ParseSetPoolOwner(log types.Log) (*StoreSetPoolOwner, error) {
	event := new(StoreSetPoolOwner)
	if err := _Store.contract.UnpackLog(event, "SetPoolOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreSetPoolOwnerSetterIterator is returned from FilterSetPoolOwnerSetter and is used to iterate over the raw logs and unpacked data for SetPoolOwnerSetter events raised by the Store contract.
type StoreSetPoolOwnerSetterIterator struct {
	Event *StoreSetPoolOwnerSetter // Event containing the contract specifics and raw log

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
func (it *StoreSetPoolOwnerSetterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreSetPoolOwnerSetter)
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
		it.Event = new(StoreSetPoolOwnerSetter)
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
func (it *StoreSetPoolOwnerSetterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreSetPoolOwnerSetterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreSetPoolOwnerSetter represents a SetPoolOwnerSetter event raised by the Store contract.
type StoreSetPoolOwnerSetter struct {
	PoolId      *big.Int
	OwnerSetter common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetPoolOwnerSetter is a free log retrieval operation binding the contract event 0x7bd3dbd617f5f18a97d85d4a0b17ca7f1f567a7dacbadbbec5b9cea34f46af3c.
//
// Solidity: event SetPoolOwnerSetter(uint256 poolId, address ownerSetter)
func (_Store *StoreFilterer) FilterSetPoolOwnerSetter(opts *bind.FilterOpts) (*StoreSetPoolOwnerSetterIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "SetPoolOwnerSetter")
	if err != nil {
		return nil, err
	}
	return &StoreSetPoolOwnerSetterIterator{contract: _Store.contract, event: "SetPoolOwnerSetter", logs: logs, sub: sub}, nil
}

// WatchSetPoolOwnerSetter is a free log subscription operation binding the contract event 0x7bd3dbd617f5f18a97d85d4a0b17ca7f1f567a7dacbadbbec5b9cea34f46af3c.
//
// Solidity: event SetPoolOwnerSetter(uint256 poolId, address ownerSetter)
func (_Store *StoreFilterer) WatchSetPoolOwnerSetter(opts *bind.WatchOpts, sink chan<- *StoreSetPoolOwnerSetter) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "SetPoolOwnerSetter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreSetPoolOwnerSetter)
				if err := _Store.contract.UnpackLog(event, "SetPoolOwnerSetter", log); err != nil {
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

// ParseSetPoolOwnerSetter is a log parse operation binding the contract event 0x7bd3dbd617f5f18a97d85d4a0b17ca7f1f567a7dacbadbbec5b9cea34f46af3c.
//
// Solidity: event SetPoolOwnerSetter(uint256 poolId, address ownerSetter)
func (_Store *StoreFilterer) ParseSetPoolOwnerSetter(log types.Log) (*StoreSetPoolOwnerSetter, error) {
	event := new(StoreSetPoolOwnerSetter)
	if err := _Store.contract.UnpackLog(event, "SetPoolOwnerSetter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreStopContractIterator is returned from FilterStopContract and is used to iterate over the raw logs and unpacked data for StopContract events raised by the Store contract.
type StoreStopContractIterator struct {
	Event *StoreStopContract // Event containing the contract specifics and raw log

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
func (it *StoreStopContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreStopContract)
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
		it.Event = new(StoreStopContract)
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
func (it *StoreStopContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreStopContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreStopContract represents a StopContract event raised by the Store contract.
type StoreStopContract struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStopContract is a free log retrieval operation binding the contract event 0x94ef9a71caf532caf20b1f9820d83dc5d252c51ec150f608e4974a955cb33f47.
//
// Solidity: event StopContract(uint256 poolId)
func (_Store *StoreFilterer) FilterStopContract(opts *bind.FilterOpts) (*StoreStopContractIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "StopContract")
	if err != nil {
		return nil, err
	}
	return &StoreStopContractIterator{contract: _Store.contract, event: "StopContract", logs: logs, sub: sub}, nil
}

// WatchStopContract is a free log subscription operation binding the contract event 0x94ef9a71caf532caf20b1f9820d83dc5d252c51ec150f608e4974a955cb33f47.
//
// Solidity: event StopContract(uint256 poolId)
func (_Store *StoreFilterer) WatchStopContract(opts *bind.WatchOpts, sink chan<- *StoreStopContract) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "StopContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreStopContract)
				if err := _Store.contract.UnpackLog(event, "StopContract", log); err != nil {
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

// ParseStopContract is a log parse operation binding the contract event 0x94ef9a71caf532caf20b1f9820d83dc5d252c51ec150f608e4974a955cb33f47.
//
// Solidity: event StopContract(uint256 poolId)
func (_Store *StoreFilterer) ParseStopContract(log types.Log) (*StoreStopContract, error) {
	event := new(StoreStopContract)
	if err := _Store.contract.UnpackLog(event, "StopContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Store contract.
type StoreTransferIterator struct {
	Event *StoreTransfer // Event containing the contract specifics and raw log

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
func (it *StoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTransfer)
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
		it.Event = new(StoreTransfer)
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
func (it *StoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTransfer represents a Transfer event raised by the Store contract.
type StoreTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Store *StoreFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StoreTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StoreTransferIterator{contract: _Store.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Store *StoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StoreTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTransfer)
				if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Store *StoreFilterer) ParseTransfer(log types.Log) (*StoreTransfer, error) {
	event := new(StoreTransfer)
	if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreWemixDollarBurnedIterator is returned from FilterWemixDollarBurned and is used to iterate over the raw logs and unpacked data for WemixDollarBurned events raised by the Store contract.
type StoreWemixDollarBurnedIterator struct {
	Event *StoreWemixDollarBurned // Event containing the contract specifics and raw log

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
func (it *StoreWemixDollarBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreWemixDollarBurned)
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
		it.Event = new(StoreWemixDollarBurned)
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
func (it *StoreWemixDollarBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreWemixDollarBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreWemixDollarBurned represents a WemixDollarBurned event raised by the Store contract.
type StoreWemixDollarBurned struct {
	From   common.Address
	Amount *big.Int
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWemixDollarBurned is a free log retrieval operation binding the contract event 0xa3724ce1d4d5f628abdec7ee4d063f2111d074566122a66514e20939c9ebd693.
//
// Solidity: event WemixDollarBurned(address from, uint256 amount, uint256 poolId)
func (_Store *StoreFilterer) FilterWemixDollarBurned(opts *bind.FilterOpts) (*StoreWemixDollarBurnedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "WemixDollarBurned")
	if err != nil {
		return nil, err
	}
	return &StoreWemixDollarBurnedIterator{contract: _Store.contract, event: "WemixDollarBurned", logs: logs, sub: sub}, nil
}

// WatchWemixDollarBurned is a free log subscription operation binding the contract event 0xa3724ce1d4d5f628abdec7ee4d063f2111d074566122a66514e20939c9ebd693.
//
// Solidity: event WemixDollarBurned(address from, uint256 amount, uint256 poolId)
func (_Store *StoreFilterer) WatchWemixDollarBurned(opts *bind.WatchOpts, sink chan<- *StoreWemixDollarBurned) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "WemixDollarBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreWemixDollarBurned)
				if err := _Store.contract.UnpackLog(event, "WemixDollarBurned", log); err != nil {
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

// ParseWemixDollarBurned is a log parse operation binding the contract event 0xa3724ce1d4d5f628abdec7ee4d063f2111d074566122a66514e20939c9ebd693.
//
// Solidity: event WemixDollarBurned(address from, uint256 amount, uint256 poolId)
func (_Store *StoreFilterer) ParseWemixDollarBurned(log types.Log) (*StoreWemixDollarBurned, error) {
	event := new(StoreWemixDollarBurned)
	if err := _Store.contract.UnpackLog(event, "WemixDollarBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreWemixDollarMintedIterator is returned from FilterWemixDollarMinted and is used to iterate over the raw logs and unpacked data for WemixDollarMinted events raised by the Store contract.
type StoreWemixDollarMintedIterator struct {
	Event *StoreWemixDollarMinted // Event containing the contract specifics and raw log

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
func (it *StoreWemixDollarMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreWemixDollarMinted)
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
		it.Event = new(StoreWemixDollarMinted)
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
func (it *StoreWemixDollarMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreWemixDollarMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreWemixDollarMinted represents a WemixDollarMinted event raised by the Store contract.
type StoreWemixDollarMinted struct {
	To     common.Address
	Amount *big.Int
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWemixDollarMinted is a free log retrieval operation binding the contract event 0x13880874a82a6b61d937203f39eb8a30bf92ef5a373f913c3bada342438b4a80.
//
// Solidity: event WemixDollarMinted(address to, uint256 amount, uint256 poolId)
func (_Store *StoreFilterer) FilterWemixDollarMinted(opts *bind.FilterOpts) (*StoreWemixDollarMintedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "WemixDollarMinted")
	if err != nil {
		return nil, err
	}
	return &StoreWemixDollarMintedIterator{contract: _Store.contract, event: "WemixDollarMinted", logs: logs, sub: sub}, nil
}

// WatchWemixDollarMinted is a free log subscription operation binding the contract event 0x13880874a82a6b61d937203f39eb8a30bf92ef5a373f913c3bada342438b4a80.
//
// Solidity: event WemixDollarMinted(address to, uint256 amount, uint256 poolId)
func (_Store *StoreFilterer) WatchWemixDollarMinted(opts *bind.WatchOpts, sink chan<- *StoreWemixDollarMinted) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "WemixDollarMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreWemixDollarMinted)
				if err := _Store.contract.UnpackLog(event, "WemixDollarMinted", log); err != nil {
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

// ParseWemixDollarMinted is a log parse operation binding the contract event 0x13880874a82a6b61d937203f39eb8a30bf92ef5a373f913c3bada342438b4a80.
//
// Solidity: event WemixDollarMinted(address to, uint256 amount, uint256 poolId)
func (_Store *StoreFilterer) ParseWemixDollarMinted(log types.Log) (*StoreWemixDollarMinted, error) {
	event := new(StoreWemixDollarMinted)
	if err := _Store.contract.UnpackLog(event, "WemixDollarMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
