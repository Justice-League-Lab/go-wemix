// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coin

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

// SciptMetaData contains all meta data concerning the Scipt contract.
var SciptMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"poolName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"whitelists\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ownerSetter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"breaker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"breakerSetter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"stop\",\"type\":\"bool\"}],\"name\":\"AddPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whitelist\",\"type\":\"address\"}],\"name\":\"AddPoolWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whitelist\",\"type\":\"address\"}],\"name\":\"RemovePoolWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"ResumeContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"breaker\",\"type\":\"address\"}],\"name\":\"SetPoolBreaker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"breakerSetter\",\"type\":\"address\"}],\"name\":\"SetPoolBreakerSetter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SetPoolOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ownerSetter\",\"type\":\"address\"}],\"name\":\"SetPoolOwnerSetter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"StopContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"WemixDollarBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"WemixDollarMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_poolName\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_whitelists\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ownerSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_breaker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_breakerSetter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_stop\",\"type\":\"bool\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_whitelist\",\"type\":\"address\"}],\"name\":\"addPoolWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"}],\"name\":\"getWemixDollarInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWemixDollarInfoCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"}],\"name\":\"getWhitelistAddress\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"}],\"name\":\"getWhitelistAddressCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_whitelist\",\"type\":\"address\"}],\"name\":\"removePoolWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_whitelist\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newWhitelist\",\"type\":\"address\"}],\"name\":\"replacePoolWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"}],\"name\":\"resumeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_breaker\",\"type\":\"address\"}],\"name\":\"setPoolBreaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_breakerSetter\",\"type\":\"address\"}],\"name\":\"setPoolBreakerSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setPoolOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_ownerSetter\",\"type\":\"address\"}],\"name\":\"setPoolOwnerSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"}],\"name\":\"stopContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wemixDollarInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"poolName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"poolTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ownerSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"breaker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"breakerSetter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stop\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SciptABI is the input ABI used to generate the binding from.
// Deprecated: Use SciptMetaData.ABI instead.
var SciptABI = SciptMetaData.ABI

// Scipt is an auto generated Go binding around an Ethereum contract.
type Scipt struct {
	SciptCaller     // Read-only binding to the contract
	SciptTransactor // Write-only binding to the contract
	SciptFilterer   // Log filterer for contract events
}

// SciptCaller is an auto generated read-only Go binding around an Ethereum contract.
type SciptCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SciptTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SciptTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SciptFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SciptFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SciptSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SciptSession struct {
	Contract     *Scipt            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SciptCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SciptCallerSession struct {
	Contract *SciptCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SciptTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SciptTransactorSession struct {
	Contract     *SciptTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SciptRaw is an auto generated low-level Go binding around an Ethereum contract.
type SciptRaw struct {
	Contract *Scipt // Generic contract binding to access the raw methods on
}

// SciptCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SciptCallerRaw struct {
	Contract *SciptCaller // Generic read-only contract binding to access the raw methods on
}

// SciptTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SciptTransactorRaw struct {
	Contract *SciptTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScipt creates a new instance of Scipt, bound to a specific deployed contract.
func NewScipt(address common.Address, backend bind.ContractBackend) (*Scipt, error) {
	contract, err := bindScipt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Scipt{SciptCaller: SciptCaller{contract: contract}, SciptTransactor: SciptTransactor{contract: contract}, SciptFilterer: SciptFilterer{contract: contract}}, nil
}

// NewSciptCaller creates a new read-only instance of Scipt, bound to a specific deployed contract.
func NewSciptCaller(address common.Address, caller bind.ContractCaller) (*SciptCaller, error) {
	contract, err := bindScipt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SciptCaller{contract: contract}, nil
}

// NewSciptTransactor creates a new write-only instance of Scipt, bound to a specific deployed contract.
func NewSciptTransactor(address common.Address, transactor bind.ContractTransactor) (*SciptTransactor, error) {
	contract, err := bindScipt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SciptTransactor{contract: contract}, nil
}

// NewSciptFilterer creates a new log filterer instance of Scipt, bound to a specific deployed contract.
func NewSciptFilterer(address common.Address, filterer bind.ContractFilterer) (*SciptFilterer, error) {
	contract, err := bindScipt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SciptFilterer{contract: contract}, nil
}

// bindScipt binds a generic wrapper to an already deployed contract.
func bindScipt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SciptMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scipt *SciptRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scipt.Contract.SciptCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scipt *SciptRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scipt.Contract.SciptTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scipt *SciptRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scipt.Contract.SciptTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scipt *SciptCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scipt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scipt *SciptTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scipt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scipt *SciptTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scipt.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Scipt *SciptCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Scipt.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Scipt *SciptSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Scipt.Contract.Allowance(&_Scipt.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Scipt *SciptCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Scipt.Contract.Allowance(&_Scipt.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Scipt *SciptCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Scipt.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Scipt *SciptSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Scipt.Contract.BalanceOf(&_Scipt.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Scipt *SciptCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Scipt.Contract.BalanceOf(&_Scipt.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Scipt *SciptCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Scipt.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Scipt *SciptSession) Decimals() (uint8, error) {
	return _Scipt.Contract.Decimals(&_Scipt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Scipt *SciptCallerSession) Decimals() (uint8, error) {
	return _Scipt.Contract.Decimals(&_Scipt.CallOpts)
}

// GetWemixDollarInfo is a free data retrieval call binding the contract method 0x828298bd.
//
// Solidity: function getWemixDollarInfo(uint256 _poolId) view returns(uint256, string, uint256, address[], address, address, address, address, bool)
func (_Scipt *SciptCaller) GetWemixDollarInfo(opts *bind.CallOpts, _poolId *big.Int) (*big.Int, string, *big.Int, []common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	var out []interface{}
	err := _Scipt.contract.Call(opts, &out, "getWemixDollarInfo", _poolId)

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
func (_Scipt *SciptSession) GetWemixDollarInfo(_poolId *big.Int) (*big.Int, string, *big.Int, []common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	return _Scipt.Contract.GetWemixDollarInfo(&_Scipt.CallOpts, _poolId)
}

// GetWemixDollarInfo is a free data retrieval call binding the contract method 0x828298bd.
//
// Solidity: function getWemixDollarInfo(uint256 _poolId) view returns(uint256, string, uint256, address[], address, address, address, address, bool)
func (_Scipt *SciptCallerSession) GetWemixDollarInfo(_poolId *big.Int) (*big.Int, string, *big.Int, []common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	return _Scipt.Contract.GetWemixDollarInfo(&_Scipt.CallOpts, _poolId)
}

// GetWemixDollarInfoCount is a free data retrieval call binding the contract method 0xae0ce498.
//
// Solidity: function getWemixDollarInfoCount() view returns(uint256)
func (_Scipt *SciptCaller) GetWemixDollarInfoCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Scipt.contract.Call(opts, &out, "getWemixDollarInfoCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWemixDollarInfoCount is a free data retrieval call binding the contract method 0xae0ce498.
//
// Solidity: function getWemixDollarInfoCount() view returns(uint256)
func (_Scipt *SciptSession) GetWemixDollarInfoCount() (*big.Int, error) {
	return _Scipt.Contract.GetWemixDollarInfoCount(&_Scipt.CallOpts)
}

// GetWemixDollarInfoCount is a free data retrieval call binding the contract method 0xae0ce498.
//
// Solidity: function getWemixDollarInfoCount() view returns(uint256)
func (_Scipt *SciptCallerSession) GetWemixDollarInfoCount() (*big.Int, error) {
	return _Scipt.Contract.GetWemixDollarInfoCount(&_Scipt.CallOpts)
}

// GetWhitelistAddress is a free data retrieval call binding the contract method 0x565853af.
//
// Solidity: function getWhitelistAddress(uint256 _poolId) view returns(address[])
func (_Scipt *SciptCaller) GetWhitelistAddress(opts *bind.CallOpts, _poolId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Scipt.contract.Call(opts, &out, "getWhitelistAddress", _poolId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWhitelistAddress is a free data retrieval call binding the contract method 0x565853af.
//
// Solidity: function getWhitelistAddress(uint256 _poolId) view returns(address[])
func (_Scipt *SciptSession) GetWhitelistAddress(_poolId *big.Int) ([]common.Address, error) {
	return _Scipt.Contract.GetWhitelistAddress(&_Scipt.CallOpts, _poolId)
}

// GetWhitelistAddress is a free data retrieval call binding the contract method 0x565853af.
//
// Solidity: function getWhitelistAddress(uint256 _poolId) view returns(address[])
func (_Scipt *SciptCallerSession) GetWhitelistAddress(_poolId *big.Int) ([]common.Address, error) {
	return _Scipt.Contract.GetWhitelistAddress(&_Scipt.CallOpts, _poolId)
}

// GetWhitelistAddressCount is a free data retrieval call binding the contract method 0x9c63739a.
//
// Solidity: function getWhitelistAddressCount(uint256 _poolId) view returns(uint256)
func (_Scipt *SciptCaller) GetWhitelistAddressCount(opts *bind.CallOpts, _poolId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Scipt.contract.Call(opts, &out, "getWhitelistAddressCount", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWhitelistAddressCount is a free data retrieval call binding the contract method 0x9c63739a.
//
// Solidity: function getWhitelistAddressCount(uint256 _poolId) view returns(uint256)
func (_Scipt *SciptSession) GetWhitelistAddressCount(_poolId *big.Int) (*big.Int, error) {
	return _Scipt.Contract.GetWhitelistAddressCount(&_Scipt.CallOpts, _poolId)
}

// GetWhitelistAddressCount is a free data retrieval call binding the contract method 0x9c63739a.
//
// Solidity: function getWhitelistAddressCount(uint256 _poolId) view returns(uint256)
func (_Scipt *SciptCallerSession) GetWhitelistAddressCount(_poolId *big.Int) (*big.Int, error) {
	return _Scipt.Contract.GetWhitelistAddressCount(&_Scipt.CallOpts, _poolId)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(uint256)
func (_Scipt *SciptCaller) IsWhitelist(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Scipt.contract.Call(opts, &out, "isWhitelist", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(uint256)
func (_Scipt *SciptSession) IsWhitelist(arg0 common.Address) (*big.Int, error) {
	return _Scipt.Contract.IsWhitelist(&_Scipt.CallOpts, arg0)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(uint256)
func (_Scipt *SciptCallerSession) IsWhitelist(arg0 common.Address) (*big.Int, error) {
	return _Scipt.Contract.IsWhitelist(&_Scipt.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Scipt *SciptCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Scipt.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Scipt *SciptSession) Name() (string, error) {
	return _Scipt.Contract.Name(&_Scipt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Scipt *SciptCallerSession) Name() (string, error) {
	return _Scipt.Contract.Name(&_Scipt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Scipt *SciptCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Scipt.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Scipt *SciptSession) Owner() (common.Address, error) {
	return _Scipt.Contract.Owner(&_Scipt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Scipt *SciptCallerSession) Owner() (common.Address, error) {
	return _Scipt.Contract.Owner(&_Scipt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Scipt *SciptCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Scipt.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Scipt *SciptSession) Symbol() (string, error) {
	return _Scipt.Contract.Symbol(&_Scipt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Scipt *SciptCallerSession) Symbol() (string, error) {
	return _Scipt.Contract.Symbol(&_Scipt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Scipt *SciptCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Scipt.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Scipt *SciptSession) TotalSupply() (*big.Int, error) {
	return _Scipt.Contract.TotalSupply(&_Scipt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Scipt *SciptCallerSession) TotalSupply() (*big.Int, error) {
	return _Scipt.Contract.TotalSupply(&_Scipt.CallOpts)
}

// WemixDollarInfos is a free data retrieval call binding the contract method 0x2acabf3e.
//
// Solidity: function wemixDollarInfos(uint256 ) view returns(uint256 poolId, string poolName, uint256 poolTotalSupply, address owner, address ownerSetter, address breaker, address breakerSetter, bool stop)
func (_Scipt *SciptCaller) WemixDollarInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Scipt.contract.Call(opts, &out, "wemixDollarInfos", arg0)

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
func (_Scipt *SciptSession) WemixDollarInfos(arg0 *big.Int) (struct {
	PoolId          *big.Int
	PoolName        string
	PoolTotalSupply *big.Int
	Owner           common.Address
	OwnerSetter     common.Address
	Breaker         common.Address
	BreakerSetter   common.Address
	Stop            bool
}, error) {
	return _Scipt.Contract.WemixDollarInfos(&_Scipt.CallOpts, arg0)
}

// WemixDollarInfos is a free data retrieval call binding the contract method 0x2acabf3e.
//
// Solidity: function wemixDollarInfos(uint256 ) view returns(uint256 poolId, string poolName, uint256 poolTotalSupply, address owner, address ownerSetter, address breaker, address breakerSetter, bool stop)
func (_Scipt *SciptCallerSession) WemixDollarInfos(arg0 *big.Int) (struct {
	PoolId          *big.Int
	PoolName        string
	PoolTotalSupply *big.Int
	Owner           common.Address
	OwnerSetter     common.Address
	Breaker         common.Address
	BreakerSetter   common.Address
	Stop            bool
}, error) {
	return _Scipt.Contract.WemixDollarInfos(&_Scipt.CallOpts, arg0)
}

// AddPool is a paid mutator transaction binding the contract method 0x03d19032.
//
// Solidity: function addPool(string _poolName, address[] _whitelists, address _owner, address _ownerSetter, address _breaker, address _breakerSetter, bool _stop) returns()
func (_Scipt *SciptTransactor) AddPool(opts *bind.TransactOpts, _poolName string, _whitelists []common.Address, _owner common.Address, _ownerSetter common.Address, _breaker common.Address, _breakerSetter common.Address, _stop bool) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "addPool", _poolName, _whitelists, _owner, _ownerSetter, _breaker, _breakerSetter, _stop)
}

// AddPool is a paid mutator transaction binding the contract method 0x03d19032.
//
// Solidity: function addPool(string _poolName, address[] _whitelists, address _owner, address _ownerSetter, address _breaker, address _breakerSetter, bool _stop) returns()
func (_Scipt *SciptSession) AddPool(_poolName string, _whitelists []common.Address, _owner common.Address, _ownerSetter common.Address, _breaker common.Address, _breakerSetter common.Address, _stop bool) (*types.Transaction, error) {
	return _Scipt.Contract.AddPool(&_Scipt.TransactOpts, _poolName, _whitelists, _owner, _ownerSetter, _breaker, _breakerSetter, _stop)
}

// AddPool is a paid mutator transaction binding the contract method 0x03d19032.
//
// Solidity: function addPool(string _poolName, address[] _whitelists, address _owner, address _ownerSetter, address _breaker, address _breakerSetter, bool _stop) returns()
func (_Scipt *SciptTransactorSession) AddPool(_poolName string, _whitelists []common.Address, _owner common.Address, _ownerSetter common.Address, _breaker common.Address, _breakerSetter common.Address, _stop bool) (*types.Transaction, error) {
	return _Scipt.Contract.AddPool(&_Scipt.TransactOpts, _poolName, _whitelists, _owner, _ownerSetter, _breaker, _breakerSetter, _stop)
}

// AddPoolWhitelist is a paid mutator transaction binding the contract method 0x3fab3b91.
//
// Solidity: function addPoolWhitelist(uint256 _poolId, address _whitelist) returns()
func (_Scipt *SciptTransactor) AddPoolWhitelist(opts *bind.TransactOpts, _poolId *big.Int, _whitelist common.Address) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "addPoolWhitelist", _poolId, _whitelist)
}

// AddPoolWhitelist is a paid mutator transaction binding the contract method 0x3fab3b91.
//
// Solidity: function addPoolWhitelist(uint256 _poolId, address _whitelist) returns()
func (_Scipt *SciptSession) AddPoolWhitelist(_poolId *big.Int, _whitelist common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.AddPoolWhitelist(&_Scipt.TransactOpts, _poolId, _whitelist)
}

// AddPoolWhitelist is a paid mutator transaction binding the contract method 0x3fab3b91.
//
// Solidity: function addPoolWhitelist(uint256 _poolId, address _whitelist) returns()
func (_Scipt *SciptTransactorSession) AddPoolWhitelist(_poolId *big.Int, _whitelist common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.AddPoolWhitelist(&_Scipt.TransactOpts, _poolId, _whitelist)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Scipt *SciptTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Scipt *SciptSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.Approve(&_Scipt.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Scipt *SciptTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.Approve(&_Scipt.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_Scipt *SciptTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_Scipt *SciptSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.Burn(&_Scipt.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_Scipt *SciptTransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.Burn(&_Scipt.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Scipt *SciptTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Scipt *SciptSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.DecreaseAllowance(&_Scipt.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Scipt *SciptTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.DecreaseAllowance(&_Scipt.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Scipt *SciptTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Scipt *SciptSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.IncreaseAllowance(&_Scipt.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Scipt *SciptTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.IncreaseAllowance(&_Scipt.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Scipt *SciptTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Scipt *SciptSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.Mint(&_Scipt.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Scipt *SciptTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.Mint(&_Scipt.TransactOpts, _to, _amount)
}

// RemovePoolWhitelist is a paid mutator transaction binding the contract method 0xec4f052a.
//
// Solidity: function removePoolWhitelist(uint256 _poolId, address _whitelist) returns()
func (_Scipt *SciptTransactor) RemovePoolWhitelist(opts *bind.TransactOpts, _poolId *big.Int, _whitelist common.Address) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "removePoolWhitelist", _poolId, _whitelist)
}

// RemovePoolWhitelist is a paid mutator transaction binding the contract method 0xec4f052a.
//
// Solidity: function removePoolWhitelist(uint256 _poolId, address _whitelist) returns()
func (_Scipt *SciptSession) RemovePoolWhitelist(_poolId *big.Int, _whitelist common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.RemovePoolWhitelist(&_Scipt.TransactOpts, _poolId, _whitelist)
}

// RemovePoolWhitelist is a paid mutator transaction binding the contract method 0xec4f052a.
//
// Solidity: function removePoolWhitelist(uint256 _poolId, address _whitelist) returns()
func (_Scipt *SciptTransactorSession) RemovePoolWhitelist(_poolId *big.Int, _whitelist common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.RemovePoolWhitelist(&_Scipt.TransactOpts, _poolId, _whitelist)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Scipt *SciptTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Scipt *SciptSession) RenounceOwnership() (*types.Transaction, error) {
	return _Scipt.Contract.RenounceOwnership(&_Scipt.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Scipt *SciptTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Scipt.Contract.RenounceOwnership(&_Scipt.TransactOpts)
}

// ReplacePoolWhitelist is a paid mutator transaction binding the contract method 0x874e5299.
//
// Solidity: function replacePoolWhitelist(uint256 _poolId, address _whitelist, address _newWhitelist) returns()
func (_Scipt *SciptTransactor) ReplacePoolWhitelist(opts *bind.TransactOpts, _poolId *big.Int, _whitelist common.Address, _newWhitelist common.Address) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "replacePoolWhitelist", _poolId, _whitelist, _newWhitelist)
}

// ReplacePoolWhitelist is a paid mutator transaction binding the contract method 0x874e5299.
//
// Solidity: function replacePoolWhitelist(uint256 _poolId, address _whitelist, address _newWhitelist) returns()
func (_Scipt *SciptSession) ReplacePoolWhitelist(_poolId *big.Int, _whitelist common.Address, _newWhitelist common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.ReplacePoolWhitelist(&_Scipt.TransactOpts, _poolId, _whitelist, _newWhitelist)
}

// ReplacePoolWhitelist is a paid mutator transaction binding the contract method 0x874e5299.
//
// Solidity: function replacePoolWhitelist(uint256 _poolId, address _whitelist, address _newWhitelist) returns()
func (_Scipt *SciptTransactorSession) ReplacePoolWhitelist(_poolId *big.Int, _whitelist common.Address, _newWhitelist common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.ReplacePoolWhitelist(&_Scipt.TransactOpts, _poolId, _whitelist, _newWhitelist)
}

// ResumeContract is a paid mutator transaction binding the contract method 0xa78e2f2d.
//
// Solidity: function resumeContract(uint256 _poolId) returns()
func (_Scipt *SciptTransactor) ResumeContract(opts *bind.TransactOpts, _poolId *big.Int) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "resumeContract", _poolId)
}

// ResumeContract is a paid mutator transaction binding the contract method 0xa78e2f2d.
//
// Solidity: function resumeContract(uint256 _poolId) returns()
func (_Scipt *SciptSession) ResumeContract(_poolId *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.ResumeContract(&_Scipt.TransactOpts, _poolId)
}

// ResumeContract is a paid mutator transaction binding the contract method 0xa78e2f2d.
//
// Solidity: function resumeContract(uint256 _poolId) returns()
func (_Scipt *SciptTransactorSession) ResumeContract(_poolId *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.ResumeContract(&_Scipt.TransactOpts, _poolId)
}

// SetPoolBreaker is a paid mutator transaction binding the contract method 0x7d3824d9.
//
// Solidity: function setPoolBreaker(uint256 _poolId, address _breaker) returns()
func (_Scipt *SciptTransactor) SetPoolBreaker(opts *bind.TransactOpts, _poolId *big.Int, _breaker common.Address) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "setPoolBreaker", _poolId, _breaker)
}

// SetPoolBreaker is a paid mutator transaction binding the contract method 0x7d3824d9.
//
// Solidity: function setPoolBreaker(uint256 _poolId, address _breaker) returns()
func (_Scipt *SciptSession) SetPoolBreaker(_poolId *big.Int, _breaker common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.SetPoolBreaker(&_Scipt.TransactOpts, _poolId, _breaker)
}

// SetPoolBreaker is a paid mutator transaction binding the contract method 0x7d3824d9.
//
// Solidity: function setPoolBreaker(uint256 _poolId, address _breaker) returns()
func (_Scipt *SciptTransactorSession) SetPoolBreaker(_poolId *big.Int, _breaker common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.SetPoolBreaker(&_Scipt.TransactOpts, _poolId, _breaker)
}

// SetPoolBreakerSetter is a paid mutator transaction binding the contract method 0x6365530e.
//
// Solidity: function setPoolBreakerSetter(uint256 _poolId, address _breakerSetter) returns()
func (_Scipt *SciptTransactor) SetPoolBreakerSetter(opts *bind.TransactOpts, _poolId *big.Int, _breakerSetter common.Address) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "setPoolBreakerSetter", _poolId, _breakerSetter)
}

// SetPoolBreakerSetter is a paid mutator transaction binding the contract method 0x6365530e.
//
// Solidity: function setPoolBreakerSetter(uint256 _poolId, address _breakerSetter) returns()
func (_Scipt *SciptSession) SetPoolBreakerSetter(_poolId *big.Int, _breakerSetter common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.SetPoolBreakerSetter(&_Scipt.TransactOpts, _poolId, _breakerSetter)
}

// SetPoolBreakerSetter is a paid mutator transaction binding the contract method 0x6365530e.
//
// Solidity: function setPoolBreakerSetter(uint256 _poolId, address _breakerSetter) returns()
func (_Scipt *SciptTransactorSession) SetPoolBreakerSetter(_poolId *big.Int, _breakerSetter common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.SetPoolBreakerSetter(&_Scipt.TransactOpts, _poolId, _breakerSetter)
}

// SetPoolOwner is a paid mutator transaction binding the contract method 0x52202b0a.
//
// Solidity: function setPoolOwner(uint256 _poolId, address _owner) returns()
func (_Scipt *SciptTransactor) SetPoolOwner(opts *bind.TransactOpts, _poolId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "setPoolOwner", _poolId, _owner)
}

// SetPoolOwner is a paid mutator transaction binding the contract method 0x52202b0a.
//
// Solidity: function setPoolOwner(uint256 _poolId, address _owner) returns()
func (_Scipt *SciptSession) SetPoolOwner(_poolId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.SetPoolOwner(&_Scipt.TransactOpts, _poolId, _owner)
}

// SetPoolOwner is a paid mutator transaction binding the contract method 0x52202b0a.
//
// Solidity: function setPoolOwner(uint256 _poolId, address _owner) returns()
func (_Scipt *SciptTransactorSession) SetPoolOwner(_poolId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.SetPoolOwner(&_Scipt.TransactOpts, _poolId, _owner)
}

// SetPoolOwnerSetter is a paid mutator transaction binding the contract method 0x50f63fe7.
//
// Solidity: function setPoolOwnerSetter(uint256 _poolId, address _ownerSetter) returns()
func (_Scipt *SciptTransactor) SetPoolOwnerSetter(opts *bind.TransactOpts, _poolId *big.Int, _ownerSetter common.Address) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "setPoolOwnerSetter", _poolId, _ownerSetter)
}

// SetPoolOwnerSetter is a paid mutator transaction binding the contract method 0x50f63fe7.
//
// Solidity: function setPoolOwnerSetter(uint256 _poolId, address _ownerSetter) returns()
func (_Scipt *SciptSession) SetPoolOwnerSetter(_poolId *big.Int, _ownerSetter common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.SetPoolOwnerSetter(&_Scipt.TransactOpts, _poolId, _ownerSetter)
}

// SetPoolOwnerSetter is a paid mutator transaction binding the contract method 0x50f63fe7.
//
// Solidity: function setPoolOwnerSetter(uint256 _poolId, address _ownerSetter) returns()
func (_Scipt *SciptTransactorSession) SetPoolOwnerSetter(_poolId *big.Int, _ownerSetter common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.SetPoolOwnerSetter(&_Scipt.TransactOpts, _poolId, _ownerSetter)
}

// StopContract is a paid mutator transaction binding the contract method 0xbc59d503.
//
// Solidity: function stopContract(uint256 _poolId) returns()
func (_Scipt *SciptTransactor) StopContract(opts *bind.TransactOpts, _poolId *big.Int) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "stopContract", _poolId)
}

// StopContract is a paid mutator transaction binding the contract method 0xbc59d503.
//
// Solidity: function stopContract(uint256 _poolId) returns()
func (_Scipt *SciptSession) StopContract(_poolId *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.StopContract(&_Scipt.TransactOpts, _poolId)
}

// StopContract is a paid mutator transaction binding the contract method 0xbc59d503.
//
// Solidity: function stopContract(uint256 _poolId) returns()
func (_Scipt *SciptTransactorSession) StopContract(_poolId *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.StopContract(&_Scipt.TransactOpts, _poolId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Scipt *SciptTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Scipt *SciptSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.Transfer(&_Scipt.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Scipt *SciptTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.Transfer(&_Scipt.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Scipt *SciptTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Scipt *SciptSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.TransferFrom(&_Scipt.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Scipt *SciptTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scipt.Contract.TransferFrom(&_Scipt.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Scipt *SciptTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Scipt.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Scipt *SciptSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.TransferOwnership(&_Scipt.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Scipt *SciptTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Scipt.Contract.TransferOwnership(&_Scipt.TransactOpts, newOwner)
}

// SciptAddPoolIterator is returned from FilterAddPool and is used to iterate over the raw logs and unpacked data for AddPool events raised by the Scipt contract.
type SciptAddPoolIterator struct {
	Event *SciptAddPool // Event containing the contract specifics and raw log

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
func (it *SciptAddPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptAddPool)
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
		it.Event = new(SciptAddPool)
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
func (it *SciptAddPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptAddPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptAddPool represents a AddPool event raised by the Scipt contract.
type SciptAddPool struct {
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
func (_Scipt *SciptFilterer) FilterAddPool(opts *bind.FilterOpts) (*SciptAddPoolIterator, error) {

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "AddPool")
	if err != nil {
		return nil, err
	}
	return &SciptAddPoolIterator{contract: _Scipt.contract, event: "AddPool", logs: logs, sub: sub}, nil
}

// WatchAddPool is a free log subscription operation binding the contract event 0x9329e5b15842a818b22b4f72136262697e6a2acb38fabe63f72d9dd079a7bbff.
//
// Solidity: event AddPool(uint256 poolId, string poolName, address[] whitelists, address owner, address ownerSetter, address breaker, address breakerSetter, bool stop)
func (_Scipt *SciptFilterer) WatchAddPool(opts *bind.WatchOpts, sink chan<- *SciptAddPool) (event.Subscription, error) {

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "AddPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptAddPool)
				if err := _Scipt.contract.UnpackLog(event, "AddPool", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseAddPool(log types.Log) (*SciptAddPool, error) {
	event := new(SciptAddPool)
	if err := _Scipt.contract.UnpackLog(event, "AddPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SciptAddPoolWhitelistIterator is returned from FilterAddPoolWhitelist and is used to iterate over the raw logs and unpacked data for AddPoolWhitelist events raised by the Scipt contract.
type SciptAddPoolWhitelistIterator struct {
	Event *SciptAddPoolWhitelist // Event containing the contract specifics and raw log

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
func (it *SciptAddPoolWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptAddPoolWhitelist)
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
		it.Event = new(SciptAddPoolWhitelist)
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
func (it *SciptAddPoolWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptAddPoolWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptAddPoolWhitelist represents a AddPoolWhitelist event raised by the Scipt contract.
type SciptAddPoolWhitelist struct {
	PoolId    *big.Int
	Whitelist common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddPoolWhitelist is a free log retrieval operation binding the contract event 0x5f7bf5a6535d2f3b9b2649a02c1afaa7b46a85ac2bc0bac26dc979747762b96b.
//
// Solidity: event AddPoolWhitelist(uint256 poolId, address whitelist)
func (_Scipt *SciptFilterer) FilterAddPoolWhitelist(opts *bind.FilterOpts) (*SciptAddPoolWhitelistIterator, error) {

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "AddPoolWhitelist")
	if err != nil {
		return nil, err
	}
	return &SciptAddPoolWhitelistIterator{contract: _Scipt.contract, event: "AddPoolWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddPoolWhitelist is a free log subscription operation binding the contract event 0x5f7bf5a6535d2f3b9b2649a02c1afaa7b46a85ac2bc0bac26dc979747762b96b.
//
// Solidity: event AddPoolWhitelist(uint256 poolId, address whitelist)
func (_Scipt *SciptFilterer) WatchAddPoolWhitelist(opts *bind.WatchOpts, sink chan<- *SciptAddPoolWhitelist) (event.Subscription, error) {

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "AddPoolWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptAddPoolWhitelist)
				if err := _Scipt.contract.UnpackLog(event, "AddPoolWhitelist", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseAddPoolWhitelist(log types.Log) (*SciptAddPoolWhitelist, error) {
	event := new(SciptAddPoolWhitelist)
	if err := _Scipt.contract.UnpackLog(event, "AddPoolWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SciptApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Scipt contract.
type SciptApprovalIterator struct {
	Event *SciptApproval // Event containing the contract specifics and raw log

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
func (it *SciptApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptApproval)
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
		it.Event = new(SciptApproval)
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
func (it *SciptApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptApproval represents a Approval event raised by the Scipt contract.
type SciptApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Scipt *SciptFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SciptApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SciptApprovalIterator{contract: _Scipt.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Scipt *SciptFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SciptApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptApproval)
				if err := _Scipt.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseApproval(log types.Log) (*SciptApproval, error) {
	event := new(SciptApproval)
	if err := _Scipt.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SciptOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Scipt contract.
type SciptOwnershipTransferredIterator struct {
	Event *SciptOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SciptOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptOwnershipTransferred)
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
		it.Event = new(SciptOwnershipTransferred)
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
func (it *SciptOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptOwnershipTransferred represents a OwnershipTransferred event raised by the Scipt contract.
type SciptOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Scipt *SciptFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SciptOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SciptOwnershipTransferredIterator{contract: _Scipt.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Scipt *SciptFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SciptOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptOwnershipTransferred)
				if err := _Scipt.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseOwnershipTransferred(log types.Log) (*SciptOwnershipTransferred, error) {
	event := new(SciptOwnershipTransferred)
	if err := _Scipt.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SciptRemovePoolWhitelistIterator is returned from FilterRemovePoolWhitelist and is used to iterate over the raw logs and unpacked data for RemovePoolWhitelist events raised by the Scipt contract.
type SciptRemovePoolWhitelistIterator struct {
	Event *SciptRemovePoolWhitelist // Event containing the contract specifics and raw log

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
func (it *SciptRemovePoolWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptRemovePoolWhitelist)
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
		it.Event = new(SciptRemovePoolWhitelist)
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
func (it *SciptRemovePoolWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptRemovePoolWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptRemovePoolWhitelist represents a RemovePoolWhitelist event raised by the Scipt contract.
type SciptRemovePoolWhitelist struct {
	PoolId    *big.Int
	Whitelist common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemovePoolWhitelist is a free log retrieval operation binding the contract event 0xe1938c7173044f254b823afeeca186e2d18319fb30e37e0e4de6277a52d29106.
//
// Solidity: event RemovePoolWhitelist(uint256 poolId, address whitelist)
func (_Scipt *SciptFilterer) FilterRemovePoolWhitelist(opts *bind.FilterOpts) (*SciptRemovePoolWhitelistIterator, error) {

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "RemovePoolWhitelist")
	if err != nil {
		return nil, err
	}
	return &SciptRemovePoolWhitelistIterator{contract: _Scipt.contract, event: "RemovePoolWhitelist", logs: logs, sub: sub}, nil
}

// WatchRemovePoolWhitelist is a free log subscription operation binding the contract event 0xe1938c7173044f254b823afeeca186e2d18319fb30e37e0e4de6277a52d29106.
//
// Solidity: event RemovePoolWhitelist(uint256 poolId, address whitelist)
func (_Scipt *SciptFilterer) WatchRemovePoolWhitelist(opts *bind.WatchOpts, sink chan<- *SciptRemovePoolWhitelist) (event.Subscription, error) {

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "RemovePoolWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptRemovePoolWhitelist)
				if err := _Scipt.contract.UnpackLog(event, "RemovePoolWhitelist", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseRemovePoolWhitelist(log types.Log) (*SciptRemovePoolWhitelist, error) {
	event := new(SciptRemovePoolWhitelist)
	if err := _Scipt.contract.UnpackLog(event, "RemovePoolWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SciptResumeContractIterator is returned from FilterResumeContract and is used to iterate over the raw logs and unpacked data for ResumeContract events raised by the Scipt contract.
type SciptResumeContractIterator struct {
	Event *SciptResumeContract // Event containing the contract specifics and raw log

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
func (it *SciptResumeContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptResumeContract)
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
		it.Event = new(SciptResumeContract)
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
func (it *SciptResumeContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptResumeContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptResumeContract represents a ResumeContract event raised by the Scipt contract.
type SciptResumeContract struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResumeContract is a free log retrieval operation binding the contract event 0xd3b27aa7927629b0fe961a7f5ac65c9a1dd990a0bdfa1962108af3ab4aa22d8a.
//
// Solidity: event ResumeContract(uint256 poolId)
func (_Scipt *SciptFilterer) FilterResumeContract(opts *bind.FilterOpts) (*SciptResumeContractIterator, error) {

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "ResumeContract")
	if err != nil {
		return nil, err
	}
	return &SciptResumeContractIterator{contract: _Scipt.contract, event: "ResumeContract", logs: logs, sub: sub}, nil
}

// WatchResumeContract is a free log subscription operation binding the contract event 0xd3b27aa7927629b0fe961a7f5ac65c9a1dd990a0bdfa1962108af3ab4aa22d8a.
//
// Solidity: event ResumeContract(uint256 poolId)
func (_Scipt *SciptFilterer) WatchResumeContract(opts *bind.WatchOpts, sink chan<- *SciptResumeContract) (event.Subscription, error) {

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "ResumeContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptResumeContract)
				if err := _Scipt.contract.UnpackLog(event, "ResumeContract", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseResumeContract(log types.Log) (*SciptResumeContract, error) {
	event := new(SciptResumeContract)
	if err := _Scipt.contract.UnpackLog(event, "ResumeContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SciptSetPoolBreakerIterator is returned from FilterSetPoolBreaker and is used to iterate over the raw logs and unpacked data for SetPoolBreaker events raised by the Scipt contract.
type SciptSetPoolBreakerIterator struct {
	Event *SciptSetPoolBreaker // Event containing the contract specifics and raw log

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
func (it *SciptSetPoolBreakerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptSetPoolBreaker)
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
		it.Event = new(SciptSetPoolBreaker)
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
func (it *SciptSetPoolBreakerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptSetPoolBreakerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptSetPoolBreaker represents a SetPoolBreaker event raised by the Scipt contract.
type SciptSetPoolBreaker struct {
	PoolId  *big.Int
	Breaker common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetPoolBreaker is a free log retrieval operation binding the contract event 0x70d216e53bba32c29a3424ecf8aa15aad290eb6525997dd33ea324dfd9cd5339.
//
// Solidity: event SetPoolBreaker(uint256 poolId, address breaker)
func (_Scipt *SciptFilterer) FilterSetPoolBreaker(opts *bind.FilterOpts) (*SciptSetPoolBreakerIterator, error) {

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "SetPoolBreaker")
	if err != nil {
		return nil, err
	}
	return &SciptSetPoolBreakerIterator{contract: _Scipt.contract, event: "SetPoolBreaker", logs: logs, sub: sub}, nil
}

// WatchSetPoolBreaker is a free log subscription operation binding the contract event 0x70d216e53bba32c29a3424ecf8aa15aad290eb6525997dd33ea324dfd9cd5339.
//
// Solidity: event SetPoolBreaker(uint256 poolId, address breaker)
func (_Scipt *SciptFilterer) WatchSetPoolBreaker(opts *bind.WatchOpts, sink chan<- *SciptSetPoolBreaker) (event.Subscription, error) {

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "SetPoolBreaker")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptSetPoolBreaker)
				if err := _Scipt.contract.UnpackLog(event, "SetPoolBreaker", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseSetPoolBreaker(log types.Log) (*SciptSetPoolBreaker, error) {
	event := new(SciptSetPoolBreaker)
	if err := _Scipt.contract.UnpackLog(event, "SetPoolBreaker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SciptSetPoolBreakerSetterIterator is returned from FilterSetPoolBreakerSetter and is used to iterate over the raw logs and unpacked data for SetPoolBreakerSetter events raised by the Scipt contract.
type SciptSetPoolBreakerSetterIterator struct {
	Event *SciptSetPoolBreakerSetter // Event containing the contract specifics and raw log

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
func (it *SciptSetPoolBreakerSetterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptSetPoolBreakerSetter)
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
		it.Event = new(SciptSetPoolBreakerSetter)
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
func (it *SciptSetPoolBreakerSetterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptSetPoolBreakerSetterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptSetPoolBreakerSetter represents a SetPoolBreakerSetter event raised by the Scipt contract.
type SciptSetPoolBreakerSetter struct {
	PoolId        *big.Int
	BreakerSetter common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetPoolBreakerSetter is a free log retrieval operation binding the contract event 0x7fc7b82c46c9f322f2a9c3ab999a39ba8230b89e3173adfde268967405ebb2df.
//
// Solidity: event SetPoolBreakerSetter(uint256 poolId, address breakerSetter)
func (_Scipt *SciptFilterer) FilterSetPoolBreakerSetter(opts *bind.FilterOpts) (*SciptSetPoolBreakerSetterIterator, error) {

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "SetPoolBreakerSetter")
	if err != nil {
		return nil, err
	}
	return &SciptSetPoolBreakerSetterIterator{contract: _Scipt.contract, event: "SetPoolBreakerSetter", logs: logs, sub: sub}, nil
}

// WatchSetPoolBreakerSetter is a free log subscription operation binding the contract event 0x7fc7b82c46c9f322f2a9c3ab999a39ba8230b89e3173adfde268967405ebb2df.
//
// Solidity: event SetPoolBreakerSetter(uint256 poolId, address breakerSetter)
func (_Scipt *SciptFilterer) WatchSetPoolBreakerSetter(opts *bind.WatchOpts, sink chan<- *SciptSetPoolBreakerSetter) (event.Subscription, error) {

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "SetPoolBreakerSetter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptSetPoolBreakerSetter)
				if err := _Scipt.contract.UnpackLog(event, "SetPoolBreakerSetter", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseSetPoolBreakerSetter(log types.Log) (*SciptSetPoolBreakerSetter, error) {
	event := new(SciptSetPoolBreakerSetter)
	if err := _Scipt.contract.UnpackLog(event, "SetPoolBreakerSetter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SciptSetPoolOwnerIterator is returned from FilterSetPoolOwner and is used to iterate over the raw logs and unpacked data for SetPoolOwner events raised by the Scipt contract.
type SciptSetPoolOwnerIterator struct {
	Event *SciptSetPoolOwner // Event containing the contract specifics and raw log

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
func (it *SciptSetPoolOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptSetPoolOwner)
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
		it.Event = new(SciptSetPoolOwner)
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
func (it *SciptSetPoolOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptSetPoolOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptSetPoolOwner represents a SetPoolOwner event raised by the Scipt contract.
type SciptSetPoolOwner struct {
	PoolId *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPoolOwner is a free log retrieval operation binding the contract event 0xf87e07c03fe1220193f248d7937e558545020f9304702037cc5412451abfc363.
//
// Solidity: event SetPoolOwner(uint256 poolId, address owner)
func (_Scipt *SciptFilterer) FilterSetPoolOwner(opts *bind.FilterOpts) (*SciptSetPoolOwnerIterator, error) {

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "SetPoolOwner")
	if err != nil {
		return nil, err
	}
	return &SciptSetPoolOwnerIterator{contract: _Scipt.contract, event: "SetPoolOwner", logs: logs, sub: sub}, nil
}

// WatchSetPoolOwner is a free log subscription operation binding the contract event 0xf87e07c03fe1220193f248d7937e558545020f9304702037cc5412451abfc363.
//
// Solidity: event SetPoolOwner(uint256 poolId, address owner)
func (_Scipt *SciptFilterer) WatchSetPoolOwner(opts *bind.WatchOpts, sink chan<- *SciptSetPoolOwner) (event.Subscription, error) {

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "SetPoolOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptSetPoolOwner)
				if err := _Scipt.contract.UnpackLog(event, "SetPoolOwner", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseSetPoolOwner(log types.Log) (*SciptSetPoolOwner, error) {
	event := new(SciptSetPoolOwner)
	if err := _Scipt.contract.UnpackLog(event, "SetPoolOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SciptSetPoolOwnerSetterIterator is returned from FilterSetPoolOwnerSetter and is used to iterate over the raw logs and unpacked data for SetPoolOwnerSetter events raised by the Scipt contract.
type SciptSetPoolOwnerSetterIterator struct {
	Event *SciptSetPoolOwnerSetter // Event containing the contract specifics and raw log

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
func (it *SciptSetPoolOwnerSetterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptSetPoolOwnerSetter)
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
		it.Event = new(SciptSetPoolOwnerSetter)
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
func (it *SciptSetPoolOwnerSetterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptSetPoolOwnerSetterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptSetPoolOwnerSetter represents a SetPoolOwnerSetter event raised by the Scipt contract.
type SciptSetPoolOwnerSetter struct {
	PoolId      *big.Int
	OwnerSetter common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetPoolOwnerSetter is a free log retrieval operation binding the contract event 0x7bd3dbd617f5f18a97d85d4a0b17ca7f1f567a7dacbadbbec5b9cea34f46af3c.
//
// Solidity: event SetPoolOwnerSetter(uint256 poolId, address ownerSetter)
func (_Scipt *SciptFilterer) FilterSetPoolOwnerSetter(opts *bind.FilterOpts) (*SciptSetPoolOwnerSetterIterator, error) {

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "SetPoolOwnerSetter")
	if err != nil {
		return nil, err
	}
	return &SciptSetPoolOwnerSetterIterator{contract: _Scipt.contract, event: "SetPoolOwnerSetter", logs: logs, sub: sub}, nil
}

// WatchSetPoolOwnerSetter is a free log subscription operation binding the contract event 0x7bd3dbd617f5f18a97d85d4a0b17ca7f1f567a7dacbadbbec5b9cea34f46af3c.
//
// Solidity: event SetPoolOwnerSetter(uint256 poolId, address ownerSetter)
func (_Scipt *SciptFilterer) WatchSetPoolOwnerSetter(opts *bind.WatchOpts, sink chan<- *SciptSetPoolOwnerSetter) (event.Subscription, error) {

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "SetPoolOwnerSetter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptSetPoolOwnerSetter)
				if err := _Scipt.contract.UnpackLog(event, "SetPoolOwnerSetter", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseSetPoolOwnerSetter(log types.Log) (*SciptSetPoolOwnerSetter, error) {
	event := new(SciptSetPoolOwnerSetter)
	if err := _Scipt.contract.UnpackLog(event, "SetPoolOwnerSetter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SciptStopContractIterator is returned from FilterStopContract and is used to iterate over the raw logs and unpacked data for StopContract events raised by the Scipt contract.
type SciptStopContractIterator struct {
	Event *SciptStopContract // Event containing the contract specifics and raw log

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
func (it *SciptStopContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptStopContract)
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
		it.Event = new(SciptStopContract)
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
func (it *SciptStopContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptStopContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptStopContract represents a StopContract event raised by the Scipt contract.
type SciptStopContract struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStopContract is a free log retrieval operation binding the contract event 0x94ef9a71caf532caf20b1f9820d83dc5d252c51ec150f608e4974a955cb33f47.
//
// Solidity: event StopContract(uint256 poolId)
func (_Scipt *SciptFilterer) FilterStopContract(opts *bind.FilterOpts) (*SciptStopContractIterator, error) {

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "StopContract")
	if err != nil {
		return nil, err
	}
	return &SciptStopContractIterator{contract: _Scipt.contract, event: "StopContract", logs: logs, sub: sub}, nil
}

// WatchStopContract is a free log subscription operation binding the contract event 0x94ef9a71caf532caf20b1f9820d83dc5d252c51ec150f608e4974a955cb33f47.
//
// Solidity: event StopContract(uint256 poolId)
func (_Scipt *SciptFilterer) WatchStopContract(opts *bind.WatchOpts, sink chan<- *SciptStopContract) (event.Subscription, error) {

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "StopContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptStopContract)
				if err := _Scipt.contract.UnpackLog(event, "StopContract", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseStopContract(log types.Log) (*SciptStopContract, error) {
	event := new(SciptStopContract)
	if err := _Scipt.contract.UnpackLog(event, "StopContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SciptTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Scipt contract.
type SciptTransferIterator struct {
	Event *SciptTransfer // Event containing the contract specifics and raw log

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
func (it *SciptTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptTransfer)
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
		it.Event = new(SciptTransfer)
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
func (it *SciptTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptTransfer represents a Transfer event raised by the Scipt contract.
type SciptTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Scipt *SciptFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SciptTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SciptTransferIterator{contract: _Scipt.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Scipt *SciptFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SciptTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptTransfer)
				if err := _Scipt.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseTransfer(log types.Log) (*SciptTransfer, error) {
	event := new(SciptTransfer)
	if err := _Scipt.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SciptWemixDollarBurnedIterator is returned from FilterWemixDollarBurned and is used to iterate over the raw logs and unpacked data for WemixDollarBurned events raised by the Scipt contract.
type SciptWemixDollarBurnedIterator struct {
	Event *SciptWemixDollarBurned // Event containing the contract specifics and raw log

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
func (it *SciptWemixDollarBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptWemixDollarBurned)
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
		it.Event = new(SciptWemixDollarBurned)
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
func (it *SciptWemixDollarBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptWemixDollarBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptWemixDollarBurned represents a WemixDollarBurned event raised by the Scipt contract.
type SciptWemixDollarBurned struct {
	From   common.Address
	Amount *big.Int
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWemixDollarBurned is a free log retrieval operation binding the contract event 0xa3724ce1d4d5f628abdec7ee4d063f2111d074566122a66514e20939c9ebd693.
//
// Solidity: event WemixDollarBurned(address from, uint256 amount, uint256 poolId)
func (_Scipt *SciptFilterer) FilterWemixDollarBurned(opts *bind.FilterOpts) (*SciptWemixDollarBurnedIterator, error) {

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "WemixDollarBurned")
	if err != nil {
		return nil, err
	}
	return &SciptWemixDollarBurnedIterator{contract: _Scipt.contract, event: "WemixDollarBurned", logs: logs, sub: sub}, nil
}

// WatchWemixDollarBurned is a free log subscription operation binding the contract event 0xa3724ce1d4d5f628abdec7ee4d063f2111d074566122a66514e20939c9ebd693.
//
// Solidity: event WemixDollarBurned(address from, uint256 amount, uint256 poolId)
func (_Scipt *SciptFilterer) WatchWemixDollarBurned(opts *bind.WatchOpts, sink chan<- *SciptWemixDollarBurned) (event.Subscription, error) {

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "WemixDollarBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptWemixDollarBurned)
				if err := _Scipt.contract.UnpackLog(event, "WemixDollarBurned", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseWemixDollarBurned(log types.Log) (*SciptWemixDollarBurned, error) {
	event := new(SciptWemixDollarBurned)
	if err := _Scipt.contract.UnpackLog(event, "WemixDollarBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SciptWemixDollarMintedIterator is returned from FilterWemixDollarMinted and is used to iterate over the raw logs and unpacked data for WemixDollarMinted events raised by the Scipt contract.
type SciptWemixDollarMintedIterator struct {
	Event *SciptWemixDollarMinted // Event containing the contract specifics and raw log

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
func (it *SciptWemixDollarMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SciptWemixDollarMinted)
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
		it.Event = new(SciptWemixDollarMinted)
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
func (it *SciptWemixDollarMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SciptWemixDollarMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SciptWemixDollarMinted represents a WemixDollarMinted event raised by the Scipt contract.
type SciptWemixDollarMinted struct {
	To     common.Address
	Amount *big.Int
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWemixDollarMinted is a free log retrieval operation binding the contract event 0x13880874a82a6b61d937203f39eb8a30bf92ef5a373f913c3bada342438b4a80.
//
// Solidity: event WemixDollarMinted(address to, uint256 amount, uint256 poolId)
func (_Scipt *SciptFilterer) FilterWemixDollarMinted(opts *bind.FilterOpts) (*SciptWemixDollarMintedIterator, error) {

	logs, sub, err := _Scipt.contract.FilterLogs(opts, "WemixDollarMinted")
	if err != nil {
		return nil, err
	}
	return &SciptWemixDollarMintedIterator{contract: _Scipt.contract, event: "WemixDollarMinted", logs: logs, sub: sub}, nil
}

// WatchWemixDollarMinted is a free log subscription operation binding the contract event 0x13880874a82a6b61d937203f39eb8a30bf92ef5a373f913c3bada342438b4a80.
//
// Solidity: event WemixDollarMinted(address to, uint256 amount, uint256 poolId)
func (_Scipt *SciptFilterer) WatchWemixDollarMinted(opts *bind.WatchOpts, sink chan<- *SciptWemixDollarMinted) (event.Subscription, error) {

	logs, sub, err := _Scipt.contract.WatchLogs(opts, "WemixDollarMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SciptWemixDollarMinted)
				if err := _Scipt.contract.UnpackLog(event, "WemixDollarMinted", log); err != nil {
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
func (_Scipt *SciptFilterer) ParseWemixDollarMinted(log types.Log) (*SciptWemixDollarMinted, error) {
	event := new(SciptWemixDollarMinted)
	if err := _Scipt.contract.UnpackLog(event, "WemixDollarMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
