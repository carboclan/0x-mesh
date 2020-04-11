// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

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

// MarketContractProxyABI is the input ABI used to generate the binding from.
const MarketContractProxyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DURATION_DAYS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINTER_BRIDGE_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COLLATERAL_TOKEN_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ORACLE_URL\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ORACLE_STATISTIC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketContractSpecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"HONEY_LEMON_ORACLE_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_marketContractFactoryMPX\",\"type\":\"address\"},{\"name\":\"_honeyLemonOracle\",\"type\":\"address\"},{\"name\":\"_minterBridge\",\"type\":\"address\"},{\"name\":\"_imBTCTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"makerAddresses\",\"type\":\"address[]\"}],\"name\":\"getFillableAmounts\",\"outputs\":[{\"name\":\"fillableAmounts\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"makerAddress\",\"type\":\"address\"}],\"name\":\"getFillableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLatestMarketContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLatestMarketCollateralPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateRequiredCollateral\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getAllMarketContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"currentIndexValue\",\"type\":\"uint256\"},{\"name\":\"lookbackIndexValue\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"dailySettlement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"qtyToMint\",\"type\":\"uint256\"},{\"name\":\"longTokenRecipient\",\"type\":\"address\"},{\"name\":\"shortTokenRecipient\",\"type\":\"address\"}],\"name\":\"mintPositionTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_honeyLemonOracleAddress\",\"type\":\"address\"}],\"name\":\"setOracleAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minterBridgeAddress\",\"type\":\"address\"}],\"name\":\"setMinterBridgeAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_params\",\"type\":\"uint256[7]\"}],\"name\":\"setMarketContractSpecs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"settleLatestMarketContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deployContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"generateContractSpecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[7]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"generateContractNames\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[3]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketContractProxy is an auto generated Go binding around an Ethereum contract.
type MarketContractProxy struct {
	MarketContractProxyCaller     // Read-only binding to the contract
	MarketContractProxyTransactor // Write-only binding to the contract
	MarketContractProxyFilterer   // Log filterer for contract events
}

// MarketContractProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketContractProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketContractProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketContractProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketContractProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketContractProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketContractProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketContractProxySession struct {
	Contract     *MarketContractProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MarketContractProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketContractProxyCallerSession struct {
	Contract *MarketContractProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MarketContractProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketContractProxyTransactorSession struct {
	Contract     *MarketContractProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MarketContractProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketContractProxyRaw struct {
	Contract *MarketContractProxy // Generic contract binding to access the raw methods on
}

// MarketContractProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketContractProxyCallerRaw struct {
	Contract *MarketContractProxyCaller // Generic read-only contract binding to access the raw methods on
}

// MarketContractProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketContractProxyTransactorRaw struct {
	Contract *MarketContractProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketContractProxy creates a new instance of MarketContractProxy, bound to a specific deployed contract.
func NewMarketContractProxy(address common.Address, backend bind.ContractBackend) (*MarketContractProxy, error) {
	contract, err := bindMarketContractProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketContractProxy{MarketContractProxyCaller: MarketContractProxyCaller{contract: contract}, MarketContractProxyTransactor: MarketContractProxyTransactor{contract: contract}, MarketContractProxyFilterer: MarketContractProxyFilterer{contract: contract}}, nil
}

// NewMarketContractProxyCaller creates a new read-only instance of MarketContractProxy, bound to a specific deployed contract.
func NewMarketContractProxyCaller(address common.Address, caller bind.ContractCaller) (*MarketContractProxyCaller, error) {
	contract, err := bindMarketContractProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketContractProxyCaller{contract: contract}, nil
}

// NewMarketContractProxyTransactor creates a new write-only instance of MarketContractProxy, bound to a specific deployed contract.
func NewMarketContractProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketContractProxyTransactor, error) {
	contract, err := bindMarketContractProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketContractProxyTransactor{contract: contract}, nil
}

// NewMarketContractProxyFilterer creates a new log filterer instance of MarketContractProxy, bound to a specific deployed contract.
func NewMarketContractProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketContractProxyFilterer, error) {
	contract, err := bindMarketContractProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketContractProxyFilterer{contract: contract}, nil
}

// bindMarketContractProxy binds a generic wrapper to an already deployed contract.
func bindMarketContractProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketContractProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketContractProxy *MarketContractProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MarketContractProxy.Contract.MarketContractProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketContractProxy *MarketContractProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.MarketContractProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketContractProxy *MarketContractProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.MarketContractProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketContractProxy *MarketContractProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MarketContractProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketContractProxy *MarketContractProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketContractProxy *MarketContractProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.contract.Transact(opts, method, params...)
}

// COLLATERALTOKENADDRESS is a free data retrieval call binding the contract method 0x323bb775.
//
// Solidity: function COLLATERAL_TOKEN_ADDRESS() constant returns(address)
func (_MarketContractProxy *MarketContractProxyCaller) COLLATERALTOKENADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "COLLATERAL_TOKEN_ADDRESS")
	return *ret0, err
}

// COLLATERALTOKENADDRESS is a free data retrieval call binding the contract method 0x323bb775.
//
// Solidity: function COLLATERAL_TOKEN_ADDRESS() constant returns(address)
func (_MarketContractProxy *MarketContractProxySession) COLLATERALTOKENADDRESS() (common.Address, error) {
	return _MarketContractProxy.Contract.COLLATERALTOKENADDRESS(&_MarketContractProxy.CallOpts)
}

// COLLATERALTOKENADDRESS is a free data retrieval call binding the contract method 0x323bb775.
//
// Solidity: function COLLATERAL_TOKEN_ADDRESS() constant returns(address)
func (_MarketContractProxy *MarketContractProxyCallerSession) COLLATERALTOKENADDRESS() (common.Address, error) {
	return _MarketContractProxy.Contract.COLLATERALTOKENADDRESS(&_MarketContractProxy.CallOpts)
}

// CONTRACTDURATION is a free data retrieval call binding the contract method 0x17ac56b6.
//
// Solidity: function CONTRACT_DURATION() constant returns(uint256)
func (_MarketContractProxy *MarketContractProxyCaller) CONTRACTDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "CONTRACT_DURATION")
	return *ret0, err
}

// CONTRACTDURATION is a free data retrieval call binding the contract method 0x17ac56b6.
//
// Solidity: function CONTRACT_DURATION() constant returns(uint256)
func (_MarketContractProxy *MarketContractProxySession) CONTRACTDURATION() (*big.Int, error) {
	return _MarketContractProxy.Contract.CONTRACTDURATION(&_MarketContractProxy.CallOpts)
}

// CONTRACTDURATION is a free data retrieval call binding the contract method 0x17ac56b6.
//
// Solidity: function CONTRACT_DURATION() constant returns(uint256)
func (_MarketContractProxy *MarketContractProxyCallerSession) CONTRACTDURATION() (*big.Int, error) {
	return _MarketContractProxy.Contract.CONTRACTDURATION(&_MarketContractProxy.CallOpts)
}

// CONTRACTDURATIONDAYS is a free data retrieval call binding the contract method 0x12fb183e.
//
// Solidity: function CONTRACT_DURATION_DAYS() constant returns(uint256)
func (_MarketContractProxy *MarketContractProxyCaller) CONTRACTDURATIONDAYS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "CONTRACT_DURATION_DAYS")
	return *ret0, err
}

// CONTRACTDURATIONDAYS is a free data retrieval call binding the contract method 0x12fb183e.
//
// Solidity: function CONTRACT_DURATION_DAYS() constant returns(uint256)
func (_MarketContractProxy *MarketContractProxySession) CONTRACTDURATIONDAYS() (*big.Int, error) {
	return _MarketContractProxy.Contract.CONTRACTDURATIONDAYS(&_MarketContractProxy.CallOpts)
}

// CONTRACTDURATIONDAYS is a free data retrieval call binding the contract method 0x12fb183e.
//
// Solidity: function CONTRACT_DURATION_DAYS() constant returns(uint256)
func (_MarketContractProxy *MarketContractProxyCallerSession) CONTRACTDURATIONDAYS() (*big.Int, error) {
	return _MarketContractProxy.Contract.CONTRACTDURATIONDAYS(&_MarketContractProxy.CallOpts)
}

// HONEYLEMONORACLEADDRESS is a free data retrieval call binding the contract method 0xec443c3a.
//
// Solidity: function HONEY_LEMON_ORACLE_ADDRESS() constant returns(address)
func (_MarketContractProxy *MarketContractProxyCaller) HONEYLEMONORACLEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "HONEY_LEMON_ORACLE_ADDRESS")
	return *ret0, err
}

// HONEYLEMONORACLEADDRESS is a free data retrieval call binding the contract method 0xec443c3a.
//
// Solidity: function HONEY_LEMON_ORACLE_ADDRESS() constant returns(address)
func (_MarketContractProxy *MarketContractProxySession) HONEYLEMONORACLEADDRESS() (common.Address, error) {
	return _MarketContractProxy.Contract.HONEYLEMONORACLEADDRESS(&_MarketContractProxy.CallOpts)
}

// HONEYLEMONORACLEADDRESS is a free data retrieval call binding the contract method 0xec443c3a.
//
// Solidity: function HONEY_LEMON_ORACLE_ADDRESS() constant returns(address)
func (_MarketContractProxy *MarketContractProxyCallerSession) HONEYLEMONORACLEADDRESS() (common.Address, error) {
	return _MarketContractProxy.Contract.HONEYLEMONORACLEADDRESS(&_MarketContractProxy.CallOpts)
}

// MINTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x1523f83f.
//
// Solidity: function MINTER_BRIDGE_ADDRESS() constant returns(address)
func (_MarketContractProxy *MarketContractProxyCaller) MINTERBRIDGEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "MINTER_BRIDGE_ADDRESS")
	return *ret0, err
}

// MINTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x1523f83f.
//
// Solidity: function MINTER_BRIDGE_ADDRESS() constant returns(address)
func (_MarketContractProxy *MarketContractProxySession) MINTERBRIDGEADDRESS() (common.Address, error) {
	return _MarketContractProxy.Contract.MINTERBRIDGEADDRESS(&_MarketContractProxy.CallOpts)
}

// MINTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x1523f83f.
//
// Solidity: function MINTER_BRIDGE_ADDRESS() constant returns(address)
func (_MarketContractProxy *MarketContractProxyCallerSession) MINTERBRIDGEADDRESS() (common.Address, error) {
	return _MarketContractProxy.Contract.MINTERBRIDGEADDRESS(&_MarketContractProxy.CallOpts)
}

// ORACLESTATISTIC is a free data retrieval call binding the contract method 0x4ac77f22.
//
// Solidity: function ORACLE_STATISTIC() constant returns(string)
func (_MarketContractProxy *MarketContractProxyCaller) ORACLESTATISTIC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "ORACLE_STATISTIC")
	return *ret0, err
}

// ORACLESTATISTIC is a free data retrieval call binding the contract method 0x4ac77f22.
//
// Solidity: function ORACLE_STATISTIC() constant returns(string)
func (_MarketContractProxy *MarketContractProxySession) ORACLESTATISTIC() (string, error) {
	return _MarketContractProxy.Contract.ORACLESTATISTIC(&_MarketContractProxy.CallOpts)
}

// ORACLESTATISTIC is a free data retrieval call binding the contract method 0x4ac77f22.
//
// Solidity: function ORACLE_STATISTIC() constant returns(string)
func (_MarketContractProxy *MarketContractProxyCallerSession) ORACLESTATISTIC() (string, error) {
	return _MarketContractProxy.Contract.ORACLESTATISTIC(&_MarketContractProxy.CallOpts)
}

// ORACLEURL is a free data retrieval call binding the contract method 0x3d168eda.
//
// Solidity: function ORACLE_URL() constant returns(string)
func (_MarketContractProxy *MarketContractProxyCaller) ORACLEURL(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "ORACLE_URL")
	return *ret0, err
}

// ORACLEURL is a free data retrieval call binding the contract method 0x3d168eda.
//
// Solidity: function ORACLE_URL() constant returns(string)
func (_MarketContractProxy *MarketContractProxySession) ORACLEURL() (string, error) {
	return _MarketContractProxy.Contract.ORACLEURL(&_MarketContractProxy.CallOpts)
}

// ORACLEURL is a free data retrieval call binding the contract method 0x3d168eda.
//
// Solidity: function ORACLE_URL() constant returns(string)
func (_MarketContractProxy *MarketContractProxyCallerSession) ORACLEURL() (string, error) {
	return _MarketContractProxy.Contract.ORACLEURL(&_MarketContractProxy.CallOpts)
}

// CalculateRequiredCollateral is a free data retrieval call binding the contract method 0x34931b4b.
//
// Solidity: function calculateRequiredCollateral(uint256 amount) constant returns(uint256)
func (_MarketContractProxy *MarketContractProxyCaller) CalculateRequiredCollateral(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "calculateRequiredCollateral", amount)
	return *ret0, err
}

// CalculateRequiredCollateral is a free data retrieval call binding the contract method 0x34931b4b.
//
// Solidity: function calculateRequiredCollateral(uint256 amount) constant returns(uint256)
func (_MarketContractProxy *MarketContractProxySession) CalculateRequiredCollateral(amount *big.Int) (*big.Int, error) {
	return _MarketContractProxy.Contract.CalculateRequiredCollateral(&_MarketContractProxy.CallOpts, amount)
}

// CalculateRequiredCollateral is a free data retrieval call binding the contract method 0x34931b4b.
//
// Solidity: function calculateRequiredCollateral(uint256 amount) constant returns(uint256)
func (_MarketContractProxy *MarketContractProxyCallerSession) CalculateRequiredCollateral(amount *big.Int) (*big.Int, error) {
	return _MarketContractProxy.Contract.CalculateRequiredCollateral(&_MarketContractProxy.CallOpts, amount)
}

// GetFillableAmount is a free data retrieval call binding the contract method 0xb12a3cf3.
//
// Solidity: function getFillableAmount(address makerAddress) constant returns(uint256)
func (_MarketContractProxy *MarketContractProxyCaller) GetFillableAmount(opts *bind.CallOpts, makerAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "getFillableAmount", makerAddress)
	return *ret0, err
}

// GetFillableAmount is a free data retrieval call binding the contract method 0xb12a3cf3.
//
// Solidity: function getFillableAmount(address makerAddress) constant returns(uint256)
func (_MarketContractProxy *MarketContractProxySession) GetFillableAmount(makerAddress common.Address) (*big.Int, error) {
	return _MarketContractProxy.Contract.GetFillableAmount(&_MarketContractProxy.CallOpts, makerAddress)
}

// GetFillableAmount is a free data retrieval call binding the contract method 0xb12a3cf3.
//
// Solidity: function getFillableAmount(address makerAddress) constant returns(uint256)
func (_MarketContractProxy *MarketContractProxyCallerSession) GetFillableAmount(makerAddress common.Address) (*big.Int, error) {
	return _MarketContractProxy.Contract.GetFillableAmount(&_MarketContractProxy.CallOpts, makerAddress)
}

// GetFillableAmounts is a free data retrieval call binding the contract method 0xd7258608.
//
// Solidity: function getFillableAmounts(address[] makerAddresses) constant returns(uint256[] fillableAmounts)
func (_MarketContractProxy *MarketContractProxyCaller) GetFillableAmounts(opts *bind.CallOpts, makerAddresses []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "getFillableAmounts", makerAddresses)
	return *ret0, err
}

// GetFillableAmounts is a free data retrieval call binding the contract method 0xd7258608.
//
// Solidity: function getFillableAmounts(address[] makerAddresses) constant returns(uint256[] fillableAmounts)
func (_MarketContractProxy *MarketContractProxySession) GetFillableAmounts(makerAddresses []common.Address) ([]*big.Int, error) {
	return _MarketContractProxy.Contract.GetFillableAmounts(&_MarketContractProxy.CallOpts, makerAddresses)
}

// GetFillableAmounts is a free data retrieval call binding the contract method 0xd7258608.
//
// Solidity: function getFillableAmounts(address[] makerAddresses) constant returns(uint256[] fillableAmounts)
func (_MarketContractProxy *MarketContractProxyCallerSession) GetFillableAmounts(makerAddresses []common.Address) ([]*big.Int, error) {
	return _MarketContractProxy.Contract.GetFillableAmounts(&_MarketContractProxy.CallOpts, makerAddresses)
}

// GetLatestMarketCollateralPool is a free data retrieval call binding the contract method 0x0f604d3b.
//
// Solidity: function getLatestMarketCollateralPool() constant returns(address)
func (_MarketContractProxy *MarketContractProxyCaller) GetLatestMarketCollateralPool(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "getLatestMarketCollateralPool")
	return *ret0, err
}

// GetLatestMarketCollateralPool is a free data retrieval call binding the contract method 0x0f604d3b.
//
// Solidity: function getLatestMarketCollateralPool() constant returns(address)
func (_MarketContractProxy *MarketContractProxySession) GetLatestMarketCollateralPool() (common.Address, error) {
	return _MarketContractProxy.Contract.GetLatestMarketCollateralPool(&_MarketContractProxy.CallOpts)
}

// GetLatestMarketCollateralPool is a free data retrieval call binding the contract method 0x0f604d3b.
//
// Solidity: function getLatestMarketCollateralPool() constant returns(address)
func (_MarketContractProxy *MarketContractProxyCallerSession) GetLatestMarketCollateralPool() (common.Address, error) {
	return _MarketContractProxy.Contract.GetLatestMarketCollateralPool(&_MarketContractProxy.CallOpts)
}

// GetLatestMarketContract is a free data retrieval call binding the contract method 0x93678279.
//
// Solidity: function getLatestMarketContract() constant returns(address)
func (_MarketContractProxy *MarketContractProxyCaller) GetLatestMarketContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "getLatestMarketContract")
	return *ret0, err
}

// GetLatestMarketContract is a free data retrieval call binding the contract method 0x93678279.
//
// Solidity: function getLatestMarketContract() constant returns(address)
func (_MarketContractProxy *MarketContractProxySession) GetLatestMarketContract() (common.Address, error) {
	return _MarketContractProxy.Contract.GetLatestMarketContract(&_MarketContractProxy.CallOpts)
}

// GetLatestMarketContract is a free data retrieval call binding the contract method 0x93678279.
//
// Solidity: function getLatestMarketContract() constant returns(address)
func (_MarketContractProxy *MarketContractProxyCallerSession) GetLatestMarketContract() (common.Address, error) {
	return _MarketContractProxy.Contract.GetLatestMarketContract(&_MarketContractProxy.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_MarketContractProxy *MarketContractProxyCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_MarketContractProxy *MarketContractProxySession) IsOwner() (bool, error) {
	return _MarketContractProxy.Contract.IsOwner(&_MarketContractProxy.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_MarketContractProxy *MarketContractProxyCallerSession) IsOwner() (bool, error) {
	return _MarketContractProxy.Contract.IsOwner(&_MarketContractProxy.CallOpts)
}

// MarketContractSpecs is a free data retrieval call binding the contract method 0x9a7f60c8.
//
// Solidity: function marketContractSpecs(uint256 ) constant returns(uint256)
func (_MarketContractProxy *MarketContractProxyCaller) MarketContractSpecs(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "marketContractSpecs", arg0)
	return *ret0, err
}

// MarketContractSpecs is a free data retrieval call binding the contract method 0x9a7f60c8.
//
// Solidity: function marketContractSpecs(uint256 ) constant returns(uint256)
func (_MarketContractProxy *MarketContractProxySession) MarketContractSpecs(arg0 *big.Int) (*big.Int, error) {
	return _MarketContractProxy.Contract.MarketContractSpecs(&_MarketContractProxy.CallOpts, arg0)
}

// MarketContractSpecs is a free data retrieval call binding the contract method 0x9a7f60c8.
//
// Solidity: function marketContractSpecs(uint256 ) constant returns(uint256)
func (_MarketContractProxy *MarketContractProxyCallerSession) MarketContractSpecs(arg0 *big.Int) (*big.Int, error) {
	return _MarketContractProxy.Contract.MarketContractSpecs(&_MarketContractProxy.CallOpts, arg0)
}

// MarketContracts is a free data retrieval call binding the contract method 0x7a1de44e.
//
// Solidity: function marketContracts(uint256 ) constant returns(address)
func (_MarketContractProxy *MarketContractProxyCaller) MarketContracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "marketContracts", arg0)
	return *ret0, err
}

// MarketContracts is a free data retrieval call binding the contract method 0x7a1de44e.
//
// Solidity: function marketContracts(uint256 ) constant returns(address)
func (_MarketContractProxy *MarketContractProxySession) MarketContracts(arg0 *big.Int) (common.Address, error) {
	return _MarketContractProxy.Contract.MarketContracts(&_MarketContractProxy.CallOpts, arg0)
}

// MarketContracts is a free data retrieval call binding the contract method 0x7a1de44e.
//
// Solidity: function marketContracts(uint256 ) constant returns(address)
func (_MarketContractProxy *MarketContractProxyCallerSession) MarketContracts(arg0 *big.Int) (common.Address, error) {
	return _MarketContractProxy.Contract.MarketContracts(&_MarketContractProxy.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MarketContractProxy *MarketContractProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MarketContractProxy.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MarketContractProxy *MarketContractProxySession) Owner() (common.Address, error) {
	return _MarketContractProxy.Contract.Owner(&_MarketContractProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MarketContractProxy *MarketContractProxyCallerSession) Owner() (common.Address, error) {
	return _MarketContractProxy.Contract.Owner(&_MarketContractProxy.CallOpts)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256)
func (_MarketContractProxy *MarketContractProxyTransactor) BalanceOf(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "balanceOf", _owner)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256)
func (_MarketContractProxy *MarketContractProxySession) BalanceOf(_owner common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.BalanceOf(&_MarketContractProxy.TransactOpts, _owner)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256)
func (_MarketContractProxy *MarketContractProxyTransactorSession) BalanceOf(_owner common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.BalanceOf(&_MarketContractProxy.TransactOpts, _owner)
}

// DailySettlement is a paid mutator transaction binding the contract method 0xfe5f923e.
//
// Solidity: function dailySettlement(uint256 currentIndexValue, uint256 lookbackIndexValue, uint256 timestamp) returns()
func (_MarketContractProxy *MarketContractProxyTransactor) DailySettlement(opts *bind.TransactOpts, currentIndexValue *big.Int, lookbackIndexValue *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "dailySettlement", currentIndexValue, lookbackIndexValue, timestamp)
}

// DailySettlement is a paid mutator transaction binding the contract method 0xfe5f923e.
//
// Solidity: function dailySettlement(uint256 currentIndexValue, uint256 lookbackIndexValue, uint256 timestamp) returns()
func (_MarketContractProxy *MarketContractProxySession) DailySettlement(currentIndexValue *big.Int, lookbackIndexValue *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.DailySettlement(&_MarketContractProxy.TransactOpts, currentIndexValue, lookbackIndexValue, timestamp)
}

// DailySettlement is a paid mutator transaction binding the contract method 0xfe5f923e.
//
// Solidity: function dailySettlement(uint256 currentIndexValue, uint256 lookbackIndexValue, uint256 timestamp) returns()
func (_MarketContractProxy *MarketContractProxyTransactorSession) DailySettlement(currentIndexValue *big.Int, lookbackIndexValue *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.DailySettlement(&_MarketContractProxy.TransactOpts, currentIndexValue, lookbackIndexValue, timestamp)
}

// DeployContract is a paid mutator transaction binding the contract method 0x6cd5c39b.
//
// Solidity: function deployContract() returns(address)
func (_MarketContractProxy *MarketContractProxyTransactor) DeployContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "deployContract")
}

// DeployContract is a paid mutator transaction binding the contract method 0x6cd5c39b.
//
// Solidity: function deployContract() returns(address)
func (_MarketContractProxy *MarketContractProxySession) DeployContract() (*types.Transaction, error) {
	return _MarketContractProxy.Contract.DeployContract(&_MarketContractProxy.TransactOpts)
}

// DeployContract is a paid mutator transaction binding the contract method 0x6cd5c39b.
//
// Solidity: function deployContract() returns(address)
func (_MarketContractProxy *MarketContractProxyTransactorSession) DeployContract() (*types.Transaction, error) {
	return _MarketContractProxy.Contract.DeployContract(&_MarketContractProxy.TransactOpts)
}

// GenerateContractNames is a paid mutator transaction binding the contract method 0x6bb3a3ee.
//
// Solidity: function generateContractNames() returns(bytes32[3])
func (_MarketContractProxy *MarketContractProxyTransactor) GenerateContractNames(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "generateContractNames")
}

// GenerateContractNames is a paid mutator transaction binding the contract method 0x6bb3a3ee.
//
// Solidity: function generateContractNames() returns(bytes32[3])
func (_MarketContractProxy *MarketContractProxySession) GenerateContractNames() (*types.Transaction, error) {
	return _MarketContractProxy.Contract.GenerateContractNames(&_MarketContractProxy.TransactOpts)
}

// GenerateContractNames is a paid mutator transaction binding the contract method 0x6bb3a3ee.
//
// Solidity: function generateContractNames() returns(bytes32[3])
func (_MarketContractProxy *MarketContractProxyTransactorSession) GenerateContractNames() (*types.Transaction, error) {
	return _MarketContractProxy.Contract.GenerateContractNames(&_MarketContractProxy.TransactOpts)
}

// GenerateContractSpecs is a paid mutator transaction binding the contract method 0xf14798ca.
//
// Solidity: function generateContractSpecs() returns(uint256[7])
func (_MarketContractProxy *MarketContractProxyTransactor) GenerateContractSpecs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "generateContractSpecs")
}

// GenerateContractSpecs is a paid mutator transaction binding the contract method 0xf14798ca.
//
// Solidity: function generateContractSpecs() returns(uint256[7])
func (_MarketContractProxy *MarketContractProxySession) GenerateContractSpecs() (*types.Transaction, error) {
	return _MarketContractProxy.Contract.GenerateContractSpecs(&_MarketContractProxy.TransactOpts)
}

// GenerateContractSpecs is a paid mutator transaction binding the contract method 0xf14798ca.
//
// Solidity: function generateContractSpecs() returns(uint256[7])
func (_MarketContractProxy *MarketContractProxyTransactorSession) GenerateContractSpecs() (*types.Transaction, error) {
	return _MarketContractProxy.Contract.GenerateContractSpecs(&_MarketContractProxy.TransactOpts)
}

// GetAllMarketContracts is a paid mutator transaction binding the contract method 0xce79d803.
//
// Solidity: function getAllMarketContracts() returns(address[])
func (_MarketContractProxy *MarketContractProxyTransactor) GetAllMarketContracts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "getAllMarketContracts")
}

// GetAllMarketContracts is a paid mutator transaction binding the contract method 0xce79d803.
//
// Solidity: function getAllMarketContracts() returns(address[])
func (_MarketContractProxy *MarketContractProxySession) GetAllMarketContracts() (*types.Transaction, error) {
	return _MarketContractProxy.Contract.GetAllMarketContracts(&_MarketContractProxy.TransactOpts)
}

// GetAllMarketContracts is a paid mutator transaction binding the contract method 0xce79d803.
//
// Solidity: function getAllMarketContracts() returns(address[])
func (_MarketContractProxy *MarketContractProxyTransactorSession) GetAllMarketContracts() (*types.Transaction, error) {
	return _MarketContractProxy.Contract.GetAllMarketContracts(&_MarketContractProxy.TransactOpts)
}

// GetTime is a paid mutator transaction binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() returns(uint256)
func (_MarketContractProxy *MarketContractProxyTransactor) GetTime(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "getTime")
}

// GetTime is a paid mutator transaction binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() returns(uint256)
func (_MarketContractProxy *MarketContractProxySession) GetTime() (*types.Transaction, error) {
	return _MarketContractProxy.Contract.GetTime(&_MarketContractProxy.TransactOpts)
}

// GetTime is a paid mutator transaction binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() returns(uint256)
func (_MarketContractProxy *MarketContractProxyTransactorSession) GetTime() (*types.Transaction, error) {
	return _MarketContractProxy.Contract.GetTime(&_MarketContractProxy.TransactOpts)
}

// MintPositionTokens is a paid mutator transaction binding the contract method 0x1d161ee7.
//
// Solidity: function mintPositionTokens(uint256 qtyToMint, address longTokenRecipient, address shortTokenRecipient) returns()
func (_MarketContractProxy *MarketContractProxyTransactor) MintPositionTokens(opts *bind.TransactOpts, qtyToMint *big.Int, longTokenRecipient common.Address, shortTokenRecipient common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "mintPositionTokens", qtyToMint, longTokenRecipient, shortTokenRecipient)
}

// MintPositionTokens is a paid mutator transaction binding the contract method 0x1d161ee7.
//
// Solidity: function mintPositionTokens(uint256 qtyToMint, address longTokenRecipient, address shortTokenRecipient) returns()
func (_MarketContractProxy *MarketContractProxySession) MintPositionTokens(qtyToMint *big.Int, longTokenRecipient common.Address, shortTokenRecipient common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.MintPositionTokens(&_MarketContractProxy.TransactOpts, qtyToMint, longTokenRecipient, shortTokenRecipient)
}

// MintPositionTokens is a paid mutator transaction binding the contract method 0x1d161ee7.
//
// Solidity: function mintPositionTokens(uint256 qtyToMint, address longTokenRecipient, address shortTokenRecipient) returns()
func (_MarketContractProxy *MarketContractProxyTransactorSession) MintPositionTokens(qtyToMint *big.Int, longTokenRecipient common.Address, shortTokenRecipient common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.MintPositionTokens(&_MarketContractProxy.TransactOpts, qtyToMint, longTokenRecipient, shortTokenRecipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketContractProxy *MarketContractProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketContractProxy *MarketContractProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketContractProxy.Contract.RenounceOwnership(&_MarketContractProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketContractProxy *MarketContractProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketContractProxy.Contract.RenounceOwnership(&_MarketContractProxy.TransactOpts)
}

// SetMarketContractSpecs is a paid mutator transaction binding the contract method 0x5b5161d5.
//
// Solidity: function setMarketContractSpecs(uint256[7] _params) returns()
func (_MarketContractProxy *MarketContractProxyTransactor) SetMarketContractSpecs(opts *bind.TransactOpts, _params [7]*big.Int) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "setMarketContractSpecs", _params)
}

// SetMarketContractSpecs is a paid mutator transaction binding the contract method 0x5b5161d5.
//
// Solidity: function setMarketContractSpecs(uint256[7] _params) returns()
func (_MarketContractProxy *MarketContractProxySession) SetMarketContractSpecs(_params [7]*big.Int) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.SetMarketContractSpecs(&_MarketContractProxy.TransactOpts, _params)
}

// SetMarketContractSpecs is a paid mutator transaction binding the contract method 0x5b5161d5.
//
// Solidity: function setMarketContractSpecs(uint256[7] _params) returns()
func (_MarketContractProxy *MarketContractProxyTransactorSession) SetMarketContractSpecs(_params [7]*big.Int) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.SetMarketContractSpecs(&_MarketContractProxy.TransactOpts, _params)
}

// SetMinterBridgeAddress is a paid mutator transaction binding the contract method 0x5abe304d.
//
// Solidity: function setMinterBridgeAddress(address _minterBridgeAddress) returns()
func (_MarketContractProxy *MarketContractProxyTransactor) SetMinterBridgeAddress(opts *bind.TransactOpts, _minterBridgeAddress common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "setMinterBridgeAddress", _minterBridgeAddress)
}

// SetMinterBridgeAddress is a paid mutator transaction binding the contract method 0x5abe304d.
//
// Solidity: function setMinterBridgeAddress(address _minterBridgeAddress) returns()
func (_MarketContractProxy *MarketContractProxySession) SetMinterBridgeAddress(_minterBridgeAddress common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.SetMinterBridgeAddress(&_MarketContractProxy.TransactOpts, _minterBridgeAddress)
}

// SetMinterBridgeAddress is a paid mutator transaction binding the contract method 0x5abe304d.
//
// Solidity: function setMinterBridgeAddress(address _minterBridgeAddress) returns()
func (_MarketContractProxy *MarketContractProxyTransactorSession) SetMinterBridgeAddress(_minterBridgeAddress common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.SetMinterBridgeAddress(&_MarketContractProxy.TransactOpts, _minterBridgeAddress)
}

// SetOracleAddress is a paid mutator transaction binding the contract method 0x4c69c00f.
//
// Solidity: function setOracleAddress(address _honeyLemonOracleAddress) returns()
func (_MarketContractProxy *MarketContractProxyTransactor) SetOracleAddress(opts *bind.TransactOpts, _honeyLemonOracleAddress common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "setOracleAddress", _honeyLemonOracleAddress)
}

// SetOracleAddress is a paid mutator transaction binding the contract method 0x4c69c00f.
//
// Solidity: function setOracleAddress(address _honeyLemonOracleAddress) returns()
func (_MarketContractProxy *MarketContractProxySession) SetOracleAddress(_honeyLemonOracleAddress common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.SetOracleAddress(&_MarketContractProxy.TransactOpts, _honeyLemonOracleAddress)
}

// SetOracleAddress is a paid mutator transaction binding the contract method 0x4c69c00f.
//
// Solidity: function setOracleAddress(address _honeyLemonOracleAddress) returns()
func (_MarketContractProxy *MarketContractProxyTransactorSession) SetOracleAddress(_honeyLemonOracleAddress common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.SetOracleAddress(&_MarketContractProxy.TransactOpts, _honeyLemonOracleAddress)
}

// SettleLatestMarketContract is a paid mutator transaction binding the contract method 0x9985c4cc.
//
// Solidity: function settleLatestMarketContract(uint256 price) returns()
func (_MarketContractProxy *MarketContractProxyTransactor) SettleLatestMarketContract(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "settleLatestMarketContract", price)
}

// SettleLatestMarketContract is a paid mutator transaction binding the contract method 0x9985c4cc.
//
// Solidity: function settleLatestMarketContract(uint256 price) returns()
func (_MarketContractProxy *MarketContractProxySession) SettleLatestMarketContract(price *big.Int) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.SettleLatestMarketContract(&_MarketContractProxy.TransactOpts, price)
}

// SettleLatestMarketContract is a paid mutator transaction binding the contract method 0x9985c4cc.
//
// Solidity: function settleLatestMarketContract(uint256 price) returns()
func (_MarketContractProxy *MarketContractProxyTransactorSession) SettleLatestMarketContract(price *big.Int) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.SettleLatestMarketContract(&_MarketContractProxy.TransactOpts, price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketContractProxy *MarketContractProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketContractProxy *MarketContractProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.TransferOwnership(&_MarketContractProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketContractProxy *MarketContractProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketContractProxy.Contract.TransferOwnership(&_MarketContractProxy.TransactOpts, newOwner)
}

// MarketContractProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MarketContractProxy contract.
type MarketContractProxyOwnershipTransferredIterator struct {
	Event *MarketContractProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketContractProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketContractProxyOwnershipTransferred)
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
		it.Event = new(MarketContractProxyOwnershipTransferred)
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
func (it *MarketContractProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketContractProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketContractProxyOwnershipTransferred represents a OwnershipTransferred event raised by the MarketContractProxy contract.
type MarketContractProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketContractProxy *MarketContractProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketContractProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketContractProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketContractProxyOwnershipTransferredIterator{contract: _MarketContractProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketContractProxy *MarketContractProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketContractProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketContractProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketContractProxyOwnershipTransferred)
				if err := _MarketContractProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MarketContractProxy *MarketContractProxyFilterer) ParseOwnershipTransferred(log types.Log) (*MarketContractProxyOwnershipTransferred, error) {
	event := new(MarketContractProxyOwnershipTransferred)
	if err := _MarketContractProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
