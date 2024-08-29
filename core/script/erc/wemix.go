// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc

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

// CoreMetaData contains all meta data concerning the Core contract.
var CoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WWEMIX\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"AddLiquidityReturn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Receive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityReturn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WWEMIX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWEMIXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityWEMIX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWEMIX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWEMIXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityWEMIX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWEMIX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWEMIXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWEMIXWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWEMIX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForWEMIX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactWEMIXForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactWEMIX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapWEMIXForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// CoreABI is the input ABI used to generate the binding from.
// Deprecated: Use CoreMetaData.ABI instead.
var CoreABI = CoreMetaData.ABI

// Core is an auto generated Go binding around an Ethereum contract.
type Core struct {
	CoreCaller     // Read-only binding to the contract
	CoreTransactor // Write-only binding to the contract
	CoreFilterer   // Log filterer for contract events
}

// CoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

func  (c CoreTransactor) GetContract() *bind.BoundContract   {
		return c.contract
}	

// CoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoreSession struct {
	Contract     *Core             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoreCallerSession struct {
	Contract *CoreCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoreTransactorSession struct {
	Contract     *CoreTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoreRaw struct {
	Contract *Core // Generic contract binding to access the raw methods on
}

// CoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoreCallerRaw struct {
	Contract *CoreCaller // Generic read-only contract binding to access the raw methods on
}

// CoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoreTransactorRaw struct {
	Contract *CoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCore creates a new instance of Core, bound to a specific deployed contract.
func NewCore(address common.Address, backend bind.ContractBackend) (*Core, error) {
	contract, err := bindCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Core{CoreCaller: CoreCaller{contract: contract}, CoreTransactor: CoreTransactor{contract: contract}, CoreFilterer: CoreFilterer{contract: contract}}, nil
}

// NewCoreCaller creates a new read-only instance of Core, bound to a specific deployed contract.
func NewCoreCaller(address common.Address, caller bind.ContractCaller) (*CoreCaller, error) {
	contract, err := bindCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoreCaller{contract: contract}, nil
}

// NewCoreTransactor creates a new write-only instance of Core, bound to a specific deployed contract.
func NewCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*CoreTransactor, error) {
	contract, err := bindCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoreTransactor{contract: contract}, nil
}

// NewCoreFilterer creates a new log filterer instance of Core, bound to a specific deployed contract.
func NewCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*CoreFilterer, error) {
	contract, err := bindCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoreFilterer{contract: contract}, nil
}

// bindCore binds a generic wrapper to an already deployed contract.
func bindCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Core *CoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Core.Contract.CoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Core *CoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.Contract.CoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Core *CoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Core.Contract.CoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Core *CoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Core.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Core *CoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Core *CoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Core.Contract.contract.Transact(opts, method, params...)
}

// WWEMIX is a free data retrieval call binding the contract method 0x8e8b0e15.
//
// Solidity: function WWEMIX() view returns(address)
func (_Core *CoreCaller) WWEMIX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "WWEMIX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WWEMIX is a free data retrieval call binding the contract method 0x8e8b0e15.
//
// Solidity: function WWEMIX() view returns(address)
func (_Core *CoreSession) WWEMIX() (common.Address, error) {
	return _Core.Contract.WWEMIX(&_Core.CallOpts)
}

// WWEMIX is a free data retrieval call binding the contract method 0x8e8b0e15.
//
// Solidity: function WWEMIX() view returns(address)
func (_Core *CoreCallerSession) WWEMIX() (common.Address, error) {
	return _Core.Contract.WWEMIX(&_Core.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Core *CoreCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Core *CoreSession) Factory() (common.Address, error) {
	return _Core.Contract.Factory(&_Core.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Core *CoreCallerSession) Factory() (common.Address, error) {
	return _Core.Contract.Factory(&_Core.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Core *CoreCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Core *CoreSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Core.Contract.GetAmountIn(&_Core.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Core *CoreCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Core.Contract.GetAmountIn(&_Core.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Core *CoreCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Core *CoreSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Core.Contract.GetAmountOut(&_Core.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Core *CoreCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Core.Contract.GetAmountOut(&_Core.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Core *CoreCaller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Core *CoreSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Core.Contract.GetAmountsIn(&_Core.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Core *CoreCallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Core.Contract.GetAmountsIn(&_Core.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Core *CoreCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Core *CoreSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Core.Contract.GetAmountsOut(&_Core.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Core *CoreCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Core.Contract.GetAmountsOut(&_Core.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Core *CoreCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Core *CoreSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Core.Contract.Quote(&_Core.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Core *CoreCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Core.Contract.Quote(&_Core.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Core *CoreTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Core *CoreSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.AddLiquidity(&_Core.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Core *CoreTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.AddLiquidity(&_Core.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityWEMIX is a paid mutator transaction binding the contract method 0xf15a282f.
//
// Solidity: function addLiquidityWEMIX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountWEMIXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountWEMIX, uint256 liquidity)
func (_Core *CoreTransactor) AddLiquidityWEMIX(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountWEMIXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "addLiquidityWEMIX", token, amountTokenDesired, amountTokenMin, amountWEMIXMin, to, deadline)
}

// AddLiquidityWEMIX is a paid mutator transaction binding the contract method 0xf15a282f.
//
// Solidity: function addLiquidityWEMIX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountWEMIXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountWEMIX, uint256 liquidity)
func (_Core *CoreSession) AddLiquidityWEMIX(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountWEMIXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.AddLiquidityWEMIX(&_Core.TransactOpts, token, amountTokenDesired, amountTokenMin, amountWEMIXMin, to, deadline)
}

// AddLiquidityWEMIX is a paid mutator transaction binding the contract method 0xf15a282f.
//
// Solidity: function addLiquidityWEMIX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountWEMIXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountWEMIX, uint256 liquidity)
func (_Core *CoreTransactorSession) AddLiquidityWEMIX(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountWEMIXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.AddLiquidityWEMIX(&_Core.TransactOpts, token, amountTokenDesired, amountTokenMin, amountWEMIXMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Core *CoreTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Core *CoreSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RemoveLiquidity(&_Core.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Core *CoreTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RemoveLiquidity(&_Core.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityWEMIX is a paid mutator transaction binding the contract method 0xb168df13.
//
// Solidity: function removeLiquidityWEMIX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountWEMIXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountWEMIX)
func (_Core *CoreTransactor) RemoveLiquidityWEMIX(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountWEMIXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "removeLiquidityWEMIX", token, liquidity, amountTokenMin, amountWEMIXMin, to, deadline)
}

// RemoveLiquidityWEMIX is a paid mutator transaction binding the contract method 0xb168df13.
//
// Solidity: function removeLiquidityWEMIX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountWEMIXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountWEMIX)
func (_Core *CoreSession) RemoveLiquidityWEMIX(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountWEMIXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RemoveLiquidityWEMIX(&_Core.TransactOpts, token, liquidity, amountTokenMin, amountWEMIXMin, to, deadline)
}

// RemoveLiquidityWEMIX is a paid mutator transaction binding the contract method 0xb168df13.
//
// Solidity: function removeLiquidityWEMIX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountWEMIXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountWEMIX)
func (_Core *CoreTransactorSession) RemoveLiquidityWEMIX(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountWEMIXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RemoveLiquidityWEMIX(&_Core.TransactOpts, token, liquidity, amountTokenMin, amountWEMIXMin, to, deadline)
}

// RemoveLiquidityWEMIXWithPermit is a paid mutator transaction binding the contract method 0x81733276.
//
// Solidity: function removeLiquidityWEMIXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountWEMIXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountWEMIX)
func (_Core *CoreTransactor) RemoveLiquidityWEMIXWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountWEMIXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "removeLiquidityWEMIXWithPermit", token, liquidity, amountTokenMin, amountWEMIXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWEMIXWithPermit is a paid mutator transaction binding the contract method 0x81733276.
//
// Solidity: function removeLiquidityWEMIXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountWEMIXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountWEMIX)
func (_Core *CoreSession) RemoveLiquidityWEMIXWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountWEMIXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Core.Contract.RemoveLiquidityWEMIXWithPermit(&_Core.TransactOpts, token, liquidity, amountTokenMin, amountWEMIXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWEMIXWithPermit is a paid mutator transaction binding the contract method 0x81733276.
//
// Solidity: function removeLiquidityWEMIXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountWEMIXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountWEMIX)
func (_Core *CoreTransactorSession) RemoveLiquidityWEMIXWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountWEMIXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Core.Contract.RemoveLiquidityWEMIXWithPermit(&_Core.TransactOpts, token, liquidity, amountTokenMin, amountWEMIXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Core *CoreTransactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Core *CoreSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Core.Contract.RemoveLiquidityWithPermit(&_Core.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Core *CoreTransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Core.Contract.RemoveLiquidityWithPermit(&_Core.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Core *CoreTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Core *CoreSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SwapExactTokensForTokens(&_Core.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Core *CoreTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SwapExactTokensForTokens(&_Core.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForWEMIX is a paid mutator transaction binding the contract method 0x41876647.
//
// Solidity: function swapExactTokensForWEMIX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Core *CoreTransactor) SwapExactTokensForWEMIX(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "swapExactTokensForWEMIX", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForWEMIX is a paid mutator transaction binding the contract method 0x41876647.
//
// Solidity: function swapExactTokensForWEMIX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Core *CoreSession) SwapExactTokensForWEMIX(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SwapExactTokensForWEMIX(&_Core.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForWEMIX is a paid mutator transaction binding the contract method 0x41876647.
//
// Solidity: function swapExactTokensForWEMIX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Core *CoreTransactorSession) SwapExactTokensForWEMIX(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SwapExactTokensForWEMIX(&_Core.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactWEMIXForTokens is a paid mutator transaction binding the contract method 0x06fd4ac5.
//
// Solidity: function swapExactWEMIXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Core *CoreTransactor) SwapExactWEMIXForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "swapExactWEMIXForTokens", amountOutMin, path, to, deadline)
}

// SwapExactWEMIXForTokens is a paid mutator transaction binding the contract method 0x06fd4ac5.
//
// Solidity: function swapExactWEMIXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Core *CoreSession) SwapExactWEMIXForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SwapExactWEMIXForTokens(&_Core.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactWEMIXForTokens is a paid mutator transaction binding the contract method 0x06fd4ac5.
//
// Solidity: function swapExactWEMIXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Core *CoreTransactorSession) SwapExactWEMIXForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SwapExactWEMIXForTokens(&_Core.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Core *CoreTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Core *CoreSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SwapTokensForExactTokens(&_Core.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Core *CoreTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SwapTokensForExactTokens(&_Core.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactWEMIX is a paid mutator transaction binding the contract method 0x95c89bc9.
//
// Solidity: function swapTokensForExactWEMIX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Core *CoreTransactor) SwapTokensForExactWEMIX(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "swapTokensForExactWEMIX", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactWEMIX is a paid mutator transaction binding the contract method 0x95c89bc9.
//
// Solidity: function swapTokensForExactWEMIX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Core *CoreSession) SwapTokensForExactWEMIX(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SwapTokensForExactWEMIX(&_Core.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactWEMIX is a paid mutator transaction binding the contract method 0x95c89bc9.
//
// Solidity: function swapTokensForExactWEMIX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Core *CoreTransactorSession) SwapTokensForExactWEMIX(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SwapTokensForExactWEMIX(&_Core.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapWEMIXForExactTokens is a paid mutator transaction binding the contract method 0x3cabe617.
//
// Solidity: function swapWEMIXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Core *CoreTransactor) SwapWEMIXForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "swapWEMIXForExactTokens", amountOut, path, to, deadline)
}

// SwapWEMIXForExactTokens is a paid mutator transaction binding the contract method 0x3cabe617.
//
// Solidity: function swapWEMIXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Core *CoreSession) SwapWEMIXForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SwapWEMIXForExactTokens(&_Core.TransactOpts, amountOut, path, to, deadline)
}

// SwapWEMIXForExactTokens is a paid mutator transaction binding the contract method 0x3cabe617.
//
// Solidity: function swapWEMIXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Core *CoreTransactorSession) SwapWEMIXForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Core.Contract.SwapWEMIXForExactTokens(&_Core.TransactOpts, amountOut, path, to, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Core *CoreTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Core *CoreSession) Receive() (*types.Transaction, error) {
	return _Core.Contract.Receive(&_Core.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Core *CoreTransactorSession) Receive() (*types.Transaction, error) {
	return _Core.Contract.Receive(&_Core.TransactOpts)
}

// CoreAddLiquidityReturnIterator is returned from FilterAddLiquidityReturn and is used to iterate over the raw logs and unpacked data for AddLiquidityReturn events raised by the Core contract.
type CoreAddLiquidityReturnIterator struct {
	Event *CoreAddLiquidityReturn // Event containing the contract specifics and raw log

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
func (it *CoreAddLiquidityReturnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreAddLiquidityReturn)
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
		it.Event = new(CoreAddLiquidityReturn)
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
func (it *CoreAddLiquidityReturnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreAddLiquidityReturnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreAddLiquidityReturn represents a AddLiquidityReturn event raised by the Core contract.
type CoreAddLiquidityReturn struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidityReturn is a free log retrieval operation binding the contract event 0xde049815f82eebfd0877021cfb381850ab76db3c5e18fb38577f5ea9130d6748.
//
// Solidity: event AddLiquidityReturn(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Core *CoreFilterer) FilterAddLiquidityReturn(opts *bind.FilterOpts) (*CoreAddLiquidityReturnIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "AddLiquidityReturn")
	if err != nil {
		return nil, err
	}
	return &CoreAddLiquidityReturnIterator{contract: _Core.contract, event: "AddLiquidityReturn", logs: logs, sub: sub}, nil
}

// WatchAddLiquidityReturn is a free log subscription operation binding the contract event 0xde049815f82eebfd0877021cfb381850ab76db3c5e18fb38577f5ea9130d6748.
//
// Solidity: event AddLiquidityReturn(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Core *CoreFilterer) WatchAddLiquidityReturn(opts *bind.WatchOpts, sink chan<- *CoreAddLiquidityReturn) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "AddLiquidityReturn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreAddLiquidityReturn)
				if err := _Core.contract.UnpackLog(event, "AddLiquidityReturn", log); err != nil {
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

// ParseAddLiquidityReturn is a log parse operation binding the contract event 0xde049815f82eebfd0877021cfb381850ab76db3c5e18fb38577f5ea9130d6748.
//
// Solidity: event AddLiquidityReturn(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Core *CoreFilterer) ParseAddLiquidityReturn(log types.Log) (*CoreAddLiquidityReturn, error) {
	event := new(CoreAddLiquidityReturn)
	if err := _Core.contract.UnpackLog(event, "AddLiquidityReturn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreReceiveIterator is returned from FilterReceive and is used to iterate over the raw logs and unpacked data for Receive events raised by the Core contract.
type CoreReceiveIterator struct {
	Event *CoreReceive // Event containing the contract specifics and raw log

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
func (it *CoreReceiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreReceive)
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
		it.Event = new(CoreReceive)
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
func (it *CoreReceiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreReceiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreReceive represents a Receive event raised by the Core contract.
type CoreReceive struct {
	Who   common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReceive is a free log retrieval operation binding the contract event 0xd6717f327e0cb88b4a97a7f67a453e9258252c34937ccbdd86de7cb840e7def3.
//
// Solidity: event Receive(address indexed who, uint256 value)
func (_Core *CoreFilterer) FilterReceive(opts *bind.FilterOpts, who []common.Address) (*CoreReceiveIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "Receive", whoRule)
	if err != nil {
		return nil, err
	}
	return &CoreReceiveIterator{contract: _Core.contract, event: "Receive", logs: logs, sub: sub}, nil
}

// WatchReceive is a free log subscription operation binding the contract event 0xd6717f327e0cb88b4a97a7f67a453e9258252c34937ccbdd86de7cb840e7def3.
//
// Solidity: event Receive(address indexed who, uint256 value)
func (_Core *CoreFilterer) WatchReceive(opts *bind.WatchOpts, sink chan<- *CoreReceive, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "Receive", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreReceive)
				if err := _Core.contract.UnpackLog(event, "Receive", log); err != nil {
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

// ParseReceive is a log parse operation binding the contract event 0xd6717f327e0cb88b4a97a7f67a453e9258252c34937ccbdd86de7cb840e7def3.
//
// Solidity: event Receive(address indexed who, uint256 value)
func (_Core *CoreFilterer) ParseReceive(log types.Log) (*CoreReceive, error) {
	event := new(CoreReceive)
	if err := _Core.contract.UnpackLog(event, "Receive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRemoveLiquidityReturnIterator is returned from FilterRemoveLiquidityReturn and is used to iterate over the raw logs and unpacked data for RemoveLiquidityReturn events raised by the Core contract.
type CoreRemoveLiquidityReturnIterator struct {
	Event *CoreRemoveLiquidityReturn // Event containing the contract specifics and raw log

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
func (it *CoreRemoveLiquidityReturnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRemoveLiquidityReturn)
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
		it.Event = new(CoreRemoveLiquidityReturn)
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
func (it *CoreRemoveLiquidityReturnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRemoveLiquidityReturnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRemoveLiquidityReturn represents a RemoveLiquidityReturn event raised by the Core contract.
type CoreRemoveLiquidityReturn struct {
	AmountA *big.Int
	AmountB *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityReturn is a free log retrieval operation binding the contract event 0x1caa46c34f31da794a4d73c18bbfb8fcd02affabf15c482c052b4141371f57b3.
//
// Solidity: event RemoveLiquidityReturn(uint256 amountA, uint256 amountB)
func (_Core *CoreFilterer) FilterRemoveLiquidityReturn(opts *bind.FilterOpts) (*CoreRemoveLiquidityReturnIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "RemoveLiquidityReturn")
	if err != nil {
		return nil, err
	}
	return &CoreRemoveLiquidityReturnIterator{contract: _Core.contract, event: "RemoveLiquidityReturn", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityReturn is a free log subscription operation binding the contract event 0x1caa46c34f31da794a4d73c18bbfb8fcd02affabf15c482c052b4141371f57b3.
//
// Solidity: event RemoveLiquidityReturn(uint256 amountA, uint256 amountB)
func (_Core *CoreFilterer) WatchRemoveLiquidityReturn(opts *bind.WatchOpts, sink chan<- *CoreRemoveLiquidityReturn) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "RemoveLiquidityReturn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRemoveLiquidityReturn)
				if err := _Core.contract.UnpackLog(event, "RemoveLiquidityReturn", log); err != nil {
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

// ParseRemoveLiquidityReturn is a log parse operation binding the contract event 0x1caa46c34f31da794a4d73c18bbfb8fcd02affabf15c482c052b4141371f57b3.
//
// Solidity: event RemoveLiquidityReturn(uint256 amountA, uint256 amountB)
func (_Core *CoreFilterer) ParseRemoveLiquidityReturn(log types.Log) (*CoreRemoveLiquidityReturn, error) {
	event := new(CoreRemoveLiquidityReturn)
	if err := _Core.contract.UnpackLog(event, "RemoveLiquidityReturn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
