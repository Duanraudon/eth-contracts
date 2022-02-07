// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hello_abi

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

// HelloPolyABI is the input ABI used to generate the binding from.
const HelloPolyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_toChainId\",\"type\":\"uint64\"},{\"name\":\"_targetProxyHash\",\"type\":\"bytes\"}],\"name\":\"bindProxyHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hearSomeThing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"proxyHashMap\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_managerProxyContract\",\"type\":\"address\"}],\"name\":\"setManagerProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"managerProxyContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_toChainId\",\"type\":\"uint64\"},{\"name\":\"functionName\",\"type\":\"string\"},{\"name\":\"_somethingWoW\",\"type\":\"string\"}],\"name\":\"say\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_somethingWoW\",\"type\":\"bytes\"},{\"name\":\"_fromContractAddr\",\"type\":\"bytes\"},{\"name\":\"_toChainId\",\"type\":\"uint64\"}],\"name\":\"hear\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"SetManagerProxyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"targetProxyHash\",\"type\":\"bytes\"}],\"name\":\"BindProxyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"toContractAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"somethingWoW\",\"type\":\"bytes\"}],\"name\":\"Say\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"somethingWoW\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"fromContractAddr\",\"type\":\"bytes\"}],\"name\":\"Hear\",\"type\":\"event\"}]"

// HelloPolyFuncSigs maps the 4-byte function signature to its string representation.
var HelloPolyFuncSigs = map[string]string{
	"379b98f6": "bindProxyHash(uint64,bytes)",
	"f42e8ff0": "hear(bytes,bytes,uint64)",
	"952ccdc4": "hearSomeThing()",
	"d798f881": "managerProxyContract()",
	"9e5767aa": "proxyHashMap(uint64)",
	"eb95c85f": "say(uint64,string,string)",
	"af9980f0": "setManagerProxy(address)",
}

// HelloPolyBin is the compiled bytecode used for deploying new contracts.
var HelloPolyBin = "0x608060405234801561001057600080fd5b50610c1d806100206000396000f3006080604052600436106100695763ffffffff60e060020a600035041663379b98f6811461006e578063952ccdc4146100ea5780639e5767aa14610174578063af9980f014610196578063d798f881146101b9578063eb95c85f146101ea578063f42e8ff014610290575b600080fd5b34801561007a57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100d695833567ffffffffffffffff169536956044949193909101919081908401838280828437509497506103339650505050505050565b604080519115158252519081900360200190f35b3480156100f657600080fd5b506100ff61041b565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610139578181015183820152602001610121565b50505050905090810190601f1680156101665780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561018057600080fd5b506100ff67ffffffffffffffff600435166104a6565b3480156101a257600080fd5b506101b7600160a060020a036004351661050d565b005b3480156101c557600080fd5b506101ce610574565b60408051600160a060020a039092168252519081900360200190f35b3480156101f657600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100d695833567ffffffffffffffff1695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506105839650505050505050565b34801561029c57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100d694369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497505050923567ffffffffffffffff169350610a3592505050565b67ffffffffffffffff821660009081526001602090815260408220835161035c92850190610b56565b507fdacd7d303272a3b58aec6620d6d1fb588f4996a5b46858ed437f1c34348f2d0f8383604051808367ffffffffffffffff1667ffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156103d75781810151838201526020016103bf565b50505050905090810190601f1680156104045780820380516001836020036101000a031916815260200191505b50935050505060405180910390a150600192915050565b6002805460408051602060018416156101000260001901909316849004601f8101849004840282018401909252818152929183018282801561049e5780601f106104735761010080835404028352916020019161049e565b820191906000526020600020905b81548152906001019060200180831161048157829003601f168201915b505050505081565b60016020818152600092835260409283902080548451600294821615610100026000190190911693909304601f810183900483028401830190945283835291929083018282801561049e5780601f106104735761010080835404028352916020019161049e565b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03838116919091179182905560408051929091168252517f43b1a8ec337adb61e8311ed025d99c80db65c02fe5c5027c1b6a93b40970cec4916020908290030190a150565b600054600160a060020a031681565b60008060008060606000809054906101000a9004600160a060020a0316935083600160a060020a03166387939a7f6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156105e057600080fd5b505af11580156105f4573d6000803e3d6000fd5b505050506040513d602081101561060a57600080fd5b505167ffffffffffffffff891660009081526001602081815260409283902080548451600294821615610100026000190190911693909304601f8101839004830284018301909452838352939650869550909291908301828280156106b05780601f10610685576101008083540402835291602001916106b0565b820191906000526020600020905b81548152906001019060200180831161069357829003601f168201915b5050505050905081600160a060020a031663bd5cf62589838a8a6040518563ffffffff1660e060020a028152600401808567ffffffffffffffff1667ffffffffffffffff168152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561073b578181015183820152602001610723565b50505050905090810190601f1680156107685780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b8381101561079b578181015183820152602001610783565b50505050905090810190601f1680156107c85780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b838110156107fb5781810151838201526020016107e3565b50505050905090810190601f1680156108285780820380516001836020036101000a031916815260200191505b50975050505050505050602060405180830381600087803b15801561084c57600080fd5b505af1158015610860573d6000803e3d6000fd5b505050506040513d602081101561087657600080fd5b5051151561090b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f45746843726f7373436861696e4d616e616765722063726f7373436861696e2060448201527f6578656375746564206572726f72210000000000000000000000000000000000606482015290519081900360840190fd5b7f6d79831d873f007520be5fc48cecfe5ecb3fbc99963f511184d52f0f9e8654af888288604051808467ffffffffffffffff1667ffffffffffffffff1681526020018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561098a578181015183820152602001610972565b50505050905090810190601f1680156109b75780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156109ea5781810151838201526020016109d2565b50505050905090810190601f168015610a175780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a1506001979650505050505050565b8251600090610a4b906002906020870190610b56565b507fde3dff0487572889d19ed051ddecb40eff7add6d5c728a1290a3803b2f3e465c8484604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610ab0578181015183820152602001610a98565b50505050905090810190601f168015610add5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610b10578181015183820152602001610af8565b50505050905090810190601f168015610b3d5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a15060019392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b9757805160ff1916838001178555610bc4565b82800160010185558215610bc4579182015b82811115610bc4578251825591602001919060010190610ba9565b50610bd0929150610bd4565b5090565b610bee91905b80821115610bd05760008155600101610bda565b905600a165627a7a723058203f980f3831ff8982f6770d30b6839185bdfaa64bf17bd407213c483f691112c30029"

// DeployHelloPoly deploys a new Ethereum contract, binding an instance of HelloPoly to it.
func DeployHelloPoly(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HelloPoly, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloPolyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HelloPolyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HelloPoly{HelloPolyCaller: HelloPolyCaller{contract: contract}, HelloPolyTransactor: HelloPolyTransactor{contract: contract}, HelloPolyFilterer: HelloPolyFilterer{contract: contract}}, nil
}

// HelloPoly is an auto generated Go binding around an Ethereum contract.
type HelloPoly struct {
	HelloPolyCaller     // Read-only binding to the contract
	HelloPolyTransactor // Write-only binding to the contract
	HelloPolyFilterer   // Log filterer for contract events
}

// HelloPolyCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloPolyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloPolyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloPolyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloPolyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloPolyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloPolySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloPolySession struct {
	Contract     *HelloPoly        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloPolyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloPolyCallerSession struct {
	Contract *HelloPolyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HelloPolyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloPolyTransactorSession struct {
	Contract     *HelloPolyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HelloPolyRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloPolyRaw struct {
	Contract *HelloPoly // Generic contract binding to access the raw methods on
}

// HelloPolyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloPolyCallerRaw struct {
	Contract *HelloPolyCaller // Generic read-only contract binding to access the raw methods on
}

// HelloPolyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloPolyTransactorRaw struct {
	Contract *HelloPolyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelloPoly creates a new instance of HelloPoly, bound to a specific deployed contract.
func NewHelloPoly(address common.Address, backend bind.ContractBackend) (*HelloPoly, error) {
	contract, err := bindHelloPoly(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HelloPoly{HelloPolyCaller: HelloPolyCaller{contract: contract}, HelloPolyTransactor: HelloPolyTransactor{contract: contract}, HelloPolyFilterer: HelloPolyFilterer{contract: contract}}, nil
}

// NewHelloPolyCaller creates a new read-only instance of HelloPoly, bound to a specific deployed contract.
func NewHelloPolyCaller(address common.Address, caller bind.ContractCaller) (*HelloPolyCaller, error) {
	contract, err := bindHelloPoly(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloPolyCaller{contract: contract}, nil
}

// NewHelloPolyTransactor creates a new write-only instance of HelloPoly, bound to a specific deployed contract.
func NewHelloPolyTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloPolyTransactor, error) {
	contract, err := bindHelloPoly(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloPolyTransactor{contract: contract}, nil
}

// NewHelloPolyFilterer creates a new log filterer instance of HelloPoly, bound to a specific deployed contract.
func NewHelloPolyFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloPolyFilterer, error) {
	contract, err := bindHelloPoly(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloPolyFilterer{contract: contract}, nil
}

// bindHelloPoly binds a generic wrapper to an already deployed contract.
func bindHelloPoly(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloPolyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloPoly *HelloPolyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HelloPoly.Contract.HelloPolyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloPoly *HelloPolyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloPoly.Contract.HelloPolyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloPoly *HelloPolyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloPoly.Contract.HelloPolyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloPoly *HelloPolyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HelloPoly.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloPoly *HelloPolyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloPoly.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloPoly *HelloPolyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloPoly.Contract.contract.Transact(opts, method, params...)
}

// HearSomeThing is a free data retrieval call binding the contract method 0x952ccdc4.
//
// Solidity: function hearSomeThing() view returns(bytes)
func (_HelloPoly *HelloPolyCaller) HearSomeThing(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _HelloPoly.contract.Call(opts, &out, "hearSomeThing")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// HearSomeThing is a free data retrieval call binding the contract method 0x952ccdc4.
//
// Solidity: function hearSomeThing() view returns(bytes)
func (_HelloPoly *HelloPolySession) HearSomeThing() ([]byte, error) {
	return _HelloPoly.Contract.HearSomeThing(&_HelloPoly.CallOpts)
}

// HearSomeThing is a free data retrieval call binding the contract method 0x952ccdc4.
//
// Solidity: function hearSomeThing() view returns(bytes)
func (_HelloPoly *HelloPolyCallerSession) HearSomeThing() ([]byte, error) {
	return _HelloPoly.Contract.HearSomeThing(&_HelloPoly.CallOpts)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_HelloPoly *HelloPolyCaller) ManagerProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HelloPoly.contract.Call(opts, &out, "managerProxyContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_HelloPoly *HelloPolySession) ManagerProxyContract() (common.Address, error) {
	return _HelloPoly.Contract.ManagerProxyContract(&_HelloPoly.CallOpts)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_HelloPoly *HelloPolyCallerSession) ManagerProxyContract() (common.Address, error) {
	return _HelloPoly.Contract.ManagerProxyContract(&_HelloPoly.CallOpts)
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_HelloPoly *HelloPolyCaller) ProxyHashMap(opts *bind.CallOpts, arg0 uint64) ([]byte, error) {
	var out []interface{}
	err := _HelloPoly.contract.Call(opts, &out, "proxyHashMap", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_HelloPoly *HelloPolySession) ProxyHashMap(arg0 uint64) ([]byte, error) {
	return _HelloPoly.Contract.ProxyHashMap(&_HelloPoly.CallOpts, arg0)
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_HelloPoly *HelloPolyCallerSession) ProxyHashMap(arg0 uint64) ([]byte, error) {
	return _HelloPoly.Contract.ProxyHashMap(&_HelloPoly.CallOpts, arg0)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 _toChainId, bytes _targetProxyHash) returns(bool)
func (_HelloPoly *HelloPolyTransactor) BindProxyHash(opts *bind.TransactOpts, _toChainId uint64, _targetProxyHash []byte) (*types.Transaction, error) {
	return _HelloPoly.contract.Transact(opts, "bindProxyHash", _toChainId, _targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 _toChainId, bytes _targetProxyHash) returns(bool)
func (_HelloPoly *HelloPolySession) BindProxyHash(_toChainId uint64, _targetProxyHash []byte) (*types.Transaction, error) {
	return _HelloPoly.Contract.BindProxyHash(&_HelloPoly.TransactOpts, _toChainId, _targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 _toChainId, bytes _targetProxyHash) returns(bool)
func (_HelloPoly *HelloPolyTransactorSession) BindProxyHash(_toChainId uint64, _targetProxyHash []byte) (*types.Transaction, error) {
	return _HelloPoly.Contract.BindProxyHash(&_HelloPoly.TransactOpts, _toChainId, _targetProxyHash)
}

// Hear is a paid mutator transaction binding the contract method 0xf42e8ff0.
//
// Solidity: function hear(bytes _somethingWoW, bytes _fromContractAddr, uint64 _toChainId) returns(bool)
func (_HelloPoly *HelloPolyTransactor) Hear(opts *bind.TransactOpts, _somethingWoW []byte, _fromContractAddr []byte, _toChainId uint64) (*types.Transaction, error) {
	return _HelloPoly.contract.Transact(opts, "hear", _somethingWoW, _fromContractAddr, _toChainId)
}

// Hear is a paid mutator transaction binding the contract method 0xf42e8ff0.
//
// Solidity: function hear(bytes _somethingWoW, bytes _fromContractAddr, uint64 _toChainId) returns(bool)
func (_HelloPoly *HelloPolySession) Hear(_somethingWoW []byte, _fromContractAddr []byte, _toChainId uint64) (*types.Transaction, error) {
	return _HelloPoly.Contract.Hear(&_HelloPoly.TransactOpts, _somethingWoW, _fromContractAddr, _toChainId)
}

// Hear is a paid mutator transaction binding the contract method 0xf42e8ff0.
//
// Solidity: function hear(bytes _somethingWoW, bytes _fromContractAddr, uint64 _toChainId) returns(bool)
func (_HelloPoly *HelloPolyTransactorSession) Hear(_somethingWoW []byte, _fromContractAddr []byte, _toChainId uint64) (*types.Transaction, error) {
	return _HelloPoly.Contract.Hear(&_HelloPoly.TransactOpts, _somethingWoW, _fromContractAddr, _toChainId)
}

// Say is a paid mutator transaction binding the contract method 0xeb95c85f.
//
// Solidity: function say(uint64 _toChainId, string functionName, string _somethingWoW) returns(bool)
func (_HelloPoly *HelloPolyTransactor) Say(opts *bind.TransactOpts, _toChainId uint64, functionName string, _somethingWoW string) (*types.Transaction, error) {
	return _HelloPoly.contract.Transact(opts, "say", _toChainId, functionName, _somethingWoW)
}

// Say is a paid mutator transaction binding the contract method 0xeb95c85f.
//
// Solidity: function say(uint64 _toChainId, string functionName, string _somethingWoW) returns(bool)
func (_HelloPoly *HelloPolySession) Say(_toChainId uint64, functionName string, _somethingWoW string) (*types.Transaction, error) {
	return _HelloPoly.Contract.Say(&_HelloPoly.TransactOpts, _toChainId, functionName, _somethingWoW)
}

// Say is a paid mutator transaction binding the contract method 0xeb95c85f.
//
// Solidity: function say(uint64 _toChainId, string functionName, string _somethingWoW) returns(bool)
func (_HelloPoly *HelloPolyTransactorSession) Say(_toChainId uint64, functionName string, _somethingWoW string) (*types.Transaction, error) {
	return _HelloPoly.Contract.Say(&_HelloPoly.TransactOpts, _toChainId, functionName, _somethingWoW)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address _managerProxyContract) returns()
func (_HelloPoly *HelloPolyTransactor) SetManagerProxy(opts *bind.TransactOpts, _managerProxyContract common.Address) (*types.Transaction, error) {
	return _HelloPoly.contract.Transact(opts, "setManagerProxy", _managerProxyContract)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address _managerProxyContract) returns()
func (_HelloPoly *HelloPolySession) SetManagerProxy(_managerProxyContract common.Address) (*types.Transaction, error) {
	return _HelloPoly.Contract.SetManagerProxy(&_HelloPoly.TransactOpts, _managerProxyContract)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address _managerProxyContract) returns()
func (_HelloPoly *HelloPolyTransactorSession) SetManagerProxy(_managerProxyContract common.Address) (*types.Transaction, error) {
	return _HelloPoly.Contract.SetManagerProxy(&_HelloPoly.TransactOpts, _managerProxyContract)
}

// HelloPolyBindProxyEventIterator is returned from FilterBindProxyEvent and is used to iterate over the raw logs and unpacked data for BindProxyEvent events raised by the HelloPoly contract.
type HelloPolyBindProxyEventIterator struct {
	Event *HelloPolyBindProxyEvent // Event containing the contract specifics and raw log

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
func (it *HelloPolyBindProxyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloPolyBindProxyEvent)
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
		it.Event = new(HelloPolyBindProxyEvent)
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
func (it *HelloPolyBindProxyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloPolyBindProxyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloPolyBindProxyEvent represents a BindProxyEvent event raised by the HelloPoly contract.
type HelloPolyBindProxyEvent struct {
	ToChainId       uint64
	TargetProxyHash []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBindProxyEvent is a free log retrieval operation binding the contract event 0xdacd7d303272a3b58aec6620d6d1fb588f4996a5b46858ed437f1c34348f2d0f.
//
// Solidity: event BindProxyEvent(uint64 toChainId, bytes targetProxyHash)
func (_HelloPoly *HelloPolyFilterer) FilterBindProxyEvent(opts *bind.FilterOpts) (*HelloPolyBindProxyEventIterator, error) {

	logs, sub, err := _HelloPoly.contract.FilterLogs(opts, "BindProxyEvent")
	if err != nil {
		return nil, err
	}
	return &HelloPolyBindProxyEventIterator{contract: _HelloPoly.contract, event: "BindProxyEvent", logs: logs, sub: sub}, nil
}

// WatchBindProxyEvent is a free log subscription operation binding the contract event 0xdacd7d303272a3b58aec6620d6d1fb588f4996a5b46858ed437f1c34348f2d0f.
//
// Solidity: event BindProxyEvent(uint64 toChainId, bytes targetProxyHash)
func (_HelloPoly *HelloPolyFilterer) WatchBindProxyEvent(opts *bind.WatchOpts, sink chan<- *HelloPolyBindProxyEvent) (event.Subscription, error) {

	logs, sub, err := _HelloPoly.contract.WatchLogs(opts, "BindProxyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloPolyBindProxyEvent)
				if err := _HelloPoly.contract.UnpackLog(event, "BindProxyEvent", log); err != nil {
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

// ParseBindProxyEvent is a log parse operation binding the contract event 0xdacd7d303272a3b58aec6620d6d1fb588f4996a5b46858ed437f1c34348f2d0f.
//
// Solidity: event BindProxyEvent(uint64 toChainId, bytes targetProxyHash)
func (_HelloPoly *HelloPolyFilterer) ParseBindProxyEvent(log types.Log) (*HelloPolyBindProxyEvent, error) {
	event := new(HelloPolyBindProxyEvent)
	if err := _HelloPoly.contract.UnpackLog(event, "BindProxyEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HelloPolyHearIterator is returned from FilterHear and is used to iterate over the raw logs and unpacked data for Hear events raised by the HelloPoly contract.
type HelloPolyHearIterator struct {
	Event *HelloPolyHear // Event containing the contract specifics and raw log

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
func (it *HelloPolyHearIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloPolyHear)
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
		it.Event = new(HelloPolyHear)
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
func (it *HelloPolyHearIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloPolyHearIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloPolyHear represents a Hear event raised by the HelloPoly contract.
type HelloPolyHear struct {
	SomethingWoW     []byte
	FromContractAddr []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterHear is a free log retrieval operation binding the contract event 0xde3dff0487572889d19ed051ddecb40eff7add6d5c728a1290a3803b2f3e465c.
//
// Solidity: event Hear(bytes somethingWoW, bytes fromContractAddr)
func (_HelloPoly *HelloPolyFilterer) FilterHear(opts *bind.FilterOpts) (*HelloPolyHearIterator, error) {

	logs, sub, err := _HelloPoly.contract.FilterLogs(opts, "Hear")
	if err != nil {
		return nil, err
	}
	return &HelloPolyHearIterator{contract: _HelloPoly.contract, event: "Hear", logs: logs, sub: sub}, nil
}

// WatchHear is a free log subscription operation binding the contract event 0xde3dff0487572889d19ed051ddecb40eff7add6d5c728a1290a3803b2f3e465c.
//
// Solidity: event Hear(bytes somethingWoW, bytes fromContractAddr)
func (_HelloPoly *HelloPolyFilterer) WatchHear(opts *bind.WatchOpts, sink chan<- *HelloPolyHear) (event.Subscription, error) {

	logs, sub, err := _HelloPoly.contract.WatchLogs(opts, "Hear")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloPolyHear)
				if err := _HelloPoly.contract.UnpackLog(event, "Hear", log); err != nil {
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

// ParseHear is a log parse operation binding the contract event 0xde3dff0487572889d19ed051ddecb40eff7add6d5c728a1290a3803b2f3e465c.
//
// Solidity: event Hear(bytes somethingWoW, bytes fromContractAddr)
func (_HelloPoly *HelloPolyFilterer) ParseHear(log types.Log) (*HelloPolyHear, error) {
	event := new(HelloPolyHear)
	if err := _HelloPoly.contract.UnpackLog(event, "Hear", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HelloPolySayIterator is returned from FilterSay and is used to iterate over the raw logs and unpacked data for Say events raised by the HelloPoly contract.
type HelloPolySayIterator struct {
	Event *HelloPolySay // Event containing the contract specifics and raw log

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
func (it *HelloPolySayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloPolySay)
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
		it.Event = new(HelloPolySay)
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
func (it *HelloPolySayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloPolySayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloPolySay represents a Say event raised by the HelloPoly contract.
type HelloPolySay struct {
	ToChainId         uint64
	ToContractAddress []byte
	SomethingWoW      []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSay is a free log retrieval operation binding the contract event 0x6d79831d873f007520be5fc48cecfe5ecb3fbc99963f511184d52f0f9e8654af.
//
// Solidity: event Say(uint64 toChainId, bytes toContractAddress, bytes somethingWoW)
func (_HelloPoly *HelloPolyFilterer) FilterSay(opts *bind.FilterOpts) (*HelloPolySayIterator, error) {

	logs, sub, err := _HelloPoly.contract.FilterLogs(opts, "Say")
	if err != nil {
		return nil, err
	}
	return &HelloPolySayIterator{contract: _HelloPoly.contract, event: "Say", logs: logs, sub: sub}, nil
}

// WatchSay is a free log subscription operation binding the contract event 0x6d79831d873f007520be5fc48cecfe5ecb3fbc99963f511184d52f0f9e8654af.
//
// Solidity: event Say(uint64 toChainId, bytes toContractAddress, bytes somethingWoW)
func (_HelloPoly *HelloPolyFilterer) WatchSay(opts *bind.WatchOpts, sink chan<- *HelloPolySay) (event.Subscription, error) {

	logs, sub, err := _HelloPoly.contract.WatchLogs(opts, "Say")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloPolySay)
				if err := _HelloPoly.contract.UnpackLog(event, "Say", log); err != nil {
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

// ParseSay is a log parse operation binding the contract event 0x6d79831d873f007520be5fc48cecfe5ecb3fbc99963f511184d52f0f9e8654af.
//
// Solidity: event Say(uint64 toChainId, bytes toContractAddress, bytes somethingWoW)
func (_HelloPoly *HelloPolyFilterer) ParseSay(log types.Log) (*HelloPolySay, error) {
	event := new(HelloPolySay)
	if err := _HelloPoly.contract.UnpackLog(event, "Say", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HelloPolySetManagerProxyEventIterator is returned from FilterSetManagerProxyEvent and is used to iterate over the raw logs and unpacked data for SetManagerProxyEvent events raised by the HelloPoly contract.
type HelloPolySetManagerProxyEventIterator struct {
	Event *HelloPolySetManagerProxyEvent // Event containing the contract specifics and raw log

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
func (it *HelloPolySetManagerProxyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloPolySetManagerProxyEvent)
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
		it.Event = new(HelloPolySetManagerProxyEvent)
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
func (it *HelloPolySetManagerProxyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloPolySetManagerProxyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloPolySetManagerProxyEvent represents a SetManagerProxyEvent event raised by the HelloPoly contract.
type HelloPolySetManagerProxyEvent struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetManagerProxyEvent is a free log retrieval operation binding the contract event 0x43b1a8ec337adb61e8311ed025d99c80db65c02fe5c5027c1b6a93b40970cec4.
//
// Solidity: event SetManagerProxyEvent(address manager)
func (_HelloPoly *HelloPolyFilterer) FilterSetManagerProxyEvent(opts *bind.FilterOpts) (*HelloPolySetManagerProxyEventIterator, error) {

	logs, sub, err := _HelloPoly.contract.FilterLogs(opts, "SetManagerProxyEvent")
	if err != nil {
		return nil, err
	}
	return &HelloPolySetManagerProxyEventIterator{contract: _HelloPoly.contract, event: "SetManagerProxyEvent", logs: logs, sub: sub}, nil
}

// WatchSetManagerProxyEvent is a free log subscription operation binding the contract event 0x43b1a8ec337adb61e8311ed025d99c80db65c02fe5c5027c1b6a93b40970cec4.
//
// Solidity: event SetManagerProxyEvent(address manager)
func (_HelloPoly *HelloPolyFilterer) WatchSetManagerProxyEvent(opts *bind.WatchOpts, sink chan<- *HelloPolySetManagerProxyEvent) (event.Subscription, error) {

	logs, sub, err := _HelloPoly.contract.WatchLogs(opts, "SetManagerProxyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloPolySetManagerProxyEvent)
				if err := _HelloPoly.contract.UnpackLog(event, "SetManagerProxyEvent", log); err != nil {
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

// ParseSetManagerProxyEvent is a log parse operation binding the contract event 0x43b1a8ec337adb61e8311ed025d99c80db65c02fe5c5027c1b6a93b40970cec4.
//
// Solidity: event SetManagerProxyEvent(address manager)
func (_HelloPoly *HelloPolyFilterer) ParseSetManagerProxyEvent(log types.Log) (*HelloPolySetManagerProxyEvent, error) {
	event := new(HelloPolySetManagerProxyEvent)
	if err := _HelloPoly.contract.UnpackLog(event, "SetManagerProxyEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEthCrossChainManagerABI is the input ABI used to generate the binding from.
const IEthCrossChainManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_toChainId\",\"type\":\"uint64\"},{\"name\":\"_toContract\",\"type\":\"bytes\"},{\"name\":\"_method\",\"type\":\"bytes\"},{\"name\":\"_txData\",\"type\":\"bytes\"}],\"name\":\"crossChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IEthCrossChainManagerFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainManagerFuncSigs = map[string]string{
	"bd5cf625": "crossChain(uint64,bytes,bytes,bytes)",
}

// IEthCrossChainManager is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainManager struct {
	IEthCrossChainManagerCaller     // Read-only binding to the contract
	IEthCrossChainManagerTransactor // Write-only binding to the contract
	IEthCrossChainManagerFilterer   // Log filterer for contract events
}

// IEthCrossChainManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainManagerSession struct {
	Contract     *IEthCrossChainManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainManagerCallerSession struct {
	Contract *IEthCrossChainManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IEthCrossChainManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainManagerTransactorSession struct {
	Contract     *IEthCrossChainManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainManagerRaw struct {
	Contract *IEthCrossChainManager // Generic contract binding to access the raw methods on
}

// IEthCrossChainManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerCallerRaw struct {
	Contract *IEthCrossChainManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerTransactorRaw struct {
	Contract *IEthCrossChainManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainManager creates a new instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManager(address common.Address, backend bind.ContractBackend) (*IEthCrossChainManager, error) {
	contract, err := bindIEthCrossChainManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManager{IEthCrossChainManagerCaller: IEthCrossChainManagerCaller{contract: contract}, IEthCrossChainManagerTransactor: IEthCrossChainManagerTransactor{contract: contract}, IEthCrossChainManagerFilterer: IEthCrossChainManagerFilterer{contract: contract}}, nil
}

// NewIEthCrossChainManagerCaller creates a new read-only instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainManagerCaller, error) {
	contract, err := bindIEthCrossChainManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerCaller{contract: contract}, nil
}

// NewIEthCrossChainManagerTransactor creates a new write-only instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainManagerTransactor, error) {
	contract, err := bindIEthCrossChainManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerTransactor{contract: contract}, nil
}

// NewIEthCrossChainManagerFilterer creates a new log filterer instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainManagerFilterer, error) {
	contract, err := bindIEthCrossChainManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerFilterer{contract: contract}, nil
}

// bindIEthCrossChainManager binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManager *IEthCrossChainManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.contract.Transact(opts, method, params...)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerTransactor) CrossChain(opts *bind.TransactOpts, _toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.contract.Transact(opts, "crossChain", _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.CrossChain(&_IEthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.CrossChain(&_IEthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// IEthCrossChainManagerProxyABI is the input ABI used to generate the binding from.
const IEthCrossChainManagerProxyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getEthCrossChainManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IEthCrossChainManagerProxyFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainManagerProxyFuncSigs = map[string]string{
	"87939a7f": "getEthCrossChainManager()",
}

// IEthCrossChainManagerProxy is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainManagerProxy struct {
	IEthCrossChainManagerProxyCaller     // Read-only binding to the contract
	IEthCrossChainManagerProxyTransactor // Write-only binding to the contract
	IEthCrossChainManagerProxyFilterer   // Log filterer for contract events
}

// IEthCrossChainManagerProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainManagerProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainManagerProxySession struct {
	Contract     *IEthCrossChainManagerProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainManagerProxyCallerSession struct {
	Contract *IEthCrossChainManagerProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// IEthCrossChainManagerProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainManagerProxyTransactorSession struct {
	Contract     *IEthCrossChainManagerProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyRaw struct {
	Contract *IEthCrossChainManagerProxy // Generic contract binding to access the raw methods on
}

// IEthCrossChainManagerProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyCallerRaw struct {
	Contract *IEthCrossChainManagerProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainManagerProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyTransactorRaw struct {
	Contract *IEthCrossChainManagerProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainManagerProxy creates a new instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxy(address common.Address, backend bind.ContractBackend) (*IEthCrossChainManagerProxy, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxy{IEthCrossChainManagerProxyCaller: IEthCrossChainManagerProxyCaller{contract: contract}, IEthCrossChainManagerProxyTransactor: IEthCrossChainManagerProxyTransactor{contract: contract}, IEthCrossChainManagerProxyFilterer: IEthCrossChainManagerProxyFilterer{contract: contract}}, nil
}

// NewIEthCrossChainManagerProxyCaller creates a new read-only instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxyCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainManagerProxyCaller, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxyCaller{contract: contract}, nil
}

// NewIEthCrossChainManagerProxyTransactor creates a new write-only instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainManagerProxyTransactor, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxyTransactor{contract: contract}, nil
}

// NewIEthCrossChainManagerProxyFilterer creates a new log filterer instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainManagerProxyFilterer, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxyFilterer{contract: contract}, nil
}

// bindIEthCrossChainManagerProxy binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainManagerProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainManagerProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManagerProxy.Contract.IEthCrossChainManagerProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.IEthCrossChainManagerProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.IEthCrossChainManagerProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManagerProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.contract.Transact(opts, method, params...)
}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyCaller) GetEthCrossChainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEthCrossChainManagerProxy.contract.Call(opts, &out, "getEthCrossChainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxySession) GetEthCrossChainManager() (common.Address, error) {
	return _IEthCrossChainManagerProxy.Contract.GetEthCrossChainManager(&_IEthCrossChainManagerProxy.CallOpts)
}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyCallerSession) GetEthCrossChainManager() (common.Address, error) {
	return _IEthCrossChainManagerProxy.Contract.GetEthCrossChainManager(&_IEthCrossChainManagerProxy.CallOpts)
}
