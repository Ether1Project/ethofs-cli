// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// PinStorageABI is the input ABI used to generate the binding from.
const PinStorageABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"pin\",\"type\":\"string\"}],\"name\":\"PinRemove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pinToAdd\",\"type\":\"string\"},{\"name\":\"pinSize\",\"type\":\"uint32\"}],\"name\":\"PinAdd\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PinStorageBin is the compiled bytecode used for deploying new contracts.
const PinStorageBin = `0x`

// DeployPinStorage deploys a new Ethereum contract, binding an instance of PinStorage to it.
func DeployPinStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PinStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(PinStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PinStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PinStorage{PinStorageCaller: PinStorageCaller{contract: contract}, PinStorageTransactor: PinStorageTransactor{contract: contract}, PinStorageFilterer: PinStorageFilterer{contract: contract}}, nil
}

// PinStorage is an auto generated Go binding around an Ethereum contract.
type PinStorage struct {
	PinStorageCaller     // Read-only binding to the contract
	PinStorageTransactor // Write-only binding to the contract
	PinStorageFilterer   // Log filterer for contract events
}

// PinStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type PinStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PinStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PinStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PinStorageSession struct {
	Contract     *PinStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PinStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PinStorageCallerSession struct {
	Contract *PinStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PinStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PinStorageTransactorSession struct {
	Contract     *PinStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PinStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type PinStorageRaw struct {
	Contract *PinStorage // Generic contract binding to access the raw methods on
}

// PinStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PinStorageCallerRaw struct {
	Contract *PinStorageCaller // Generic read-only contract binding to access the raw methods on
}

// PinStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PinStorageTransactorRaw struct {
	Contract *PinStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPinStorage creates a new instance of PinStorage, bound to a specific deployed contract.
func NewPinStorage(address common.Address, backend bind.ContractBackend) (*PinStorage, error) {
	contract, err := bindPinStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PinStorage{PinStorageCaller: PinStorageCaller{contract: contract}, PinStorageTransactor: PinStorageTransactor{contract: contract}, PinStorageFilterer: PinStorageFilterer{contract: contract}}, nil
}

// NewPinStorageCaller creates a new read-only instance of PinStorage, bound to a specific deployed contract.
func NewPinStorageCaller(address common.Address, caller bind.ContractCaller) (*PinStorageCaller, error) {
	contract, err := bindPinStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PinStorageCaller{contract: contract}, nil
}

// NewPinStorageTransactor creates a new write-only instance of PinStorage, bound to a specific deployed contract.
func NewPinStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*PinStorageTransactor, error) {
	contract, err := bindPinStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PinStorageTransactor{contract: contract}, nil
}

// NewPinStorageFilterer creates a new log filterer instance of PinStorage, bound to a specific deployed contract.
func NewPinStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*PinStorageFilterer, error) {
	contract, err := bindPinStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PinStorageFilterer{contract: contract}, nil
}

// bindPinStorage binds a generic wrapper to an already deployed contract.
func bindPinStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PinStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PinStorage *PinStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PinStorage.Contract.PinStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PinStorage *PinStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PinStorage.Contract.PinStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PinStorage *PinStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PinStorage.Contract.PinStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PinStorage *PinStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PinStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PinStorage *PinStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PinStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PinStorage *PinStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PinStorage.Contract.contract.Transact(opts, method, params...)
}

// PinAdd is a paid mutator transaction binding the contract method 0x8d036731.
//
// Solidity: function PinAdd(string pinToAdd, uint32 pinSize) returns()
func (_PinStorage *PinStorageTransactor) PinAdd(opts *bind.TransactOpts, pinToAdd string, pinSize uint32) (*types.Transaction, error) {
	return _PinStorage.contract.Transact(opts, "PinAdd", pinToAdd, pinSize)
}

// PinAdd is a paid mutator transaction binding the contract method 0x8d036731.
//
// Solidity: function PinAdd(string pinToAdd, uint32 pinSize) returns()
func (_PinStorage *PinStorageSession) PinAdd(pinToAdd string, pinSize uint32) (*types.Transaction, error) {
	return _PinStorage.Contract.PinAdd(&_PinStorage.TransactOpts, pinToAdd, pinSize)
}

// PinAdd is a paid mutator transaction binding the contract method 0x8d036731.
//
// Solidity: function PinAdd(string pinToAdd, uint32 pinSize) returns()
func (_PinStorage *PinStorageTransactorSession) PinAdd(pinToAdd string, pinSize uint32) (*types.Transaction, error) {
	return _PinStorage.Contract.PinAdd(&_PinStorage.TransactOpts, pinToAdd, pinSize)
}

// PinRemove is a paid mutator transaction binding the contract method 0x3f0854a7.
//
// Solidity: function PinRemove(string pin) returns()
func (_PinStorage *PinStorageTransactor) PinRemove(opts *bind.TransactOpts, pin string) (*types.Transaction, error) {
	return _PinStorage.contract.Transact(opts, "PinRemove", pin)
}

// PinRemove is a paid mutator transaction binding the contract method 0x3f0854a7.
//
// Solidity: function PinRemove(string pin) returns()
func (_PinStorage *PinStorageSession) PinRemove(pin string) (*types.Transaction, error) {
	return _PinStorage.Contract.PinRemove(&_PinStorage.TransactOpts, pin)
}

// PinRemove is a paid mutator transaction binding the contract method 0x3f0854a7.
//
// Solidity: function PinRemove(string pin) returns()
func (_PinStorage *PinStorageTransactorSession) PinRemove(pin string) (*types.Transaction, error) {
	return _PinStorage.Contract.PinRemove(&_PinStorage.TransactOpts, pin)
}

// EthoFSControllerABI is the input ABI used to generate the binding from.
const EthoFSControllerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"}],\"name\":\"RemoveUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hostingCost\",\"type\":\"uint256\"}],\"name\":\"SetHostingCost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"},{\"name\":\"MainContentHash\",\"type\":\"string\"}],\"name\":\"RemoveHostingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ScrubHostingContracts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pinStorageAddress\",\"type\":\"address\"}],\"name\":\"SetPinStorageAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"},{\"name\":\"AccountName\",\"type\":\"string\"}],\"name\":\"AddNewUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ethoFSDashboardAddress\",\"type\":\"address\"}],\"name\":\"SetEthoFSDashboardAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"},{\"name\":\"HostingContractExtensionDuration\",\"type\":\"uint32\"}],\"name\":\"ExtendContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"HostingContractName\",\"type\":\"string\"},{\"name\":\"HostingContractDuration\",\"type\":\"uint32\"},{\"name\":\"TotalContractSize\",\"type\":\"uint32\"},{\"name\":\"MainContentHash\",\"type\":\"string\"},{\"name\":\"pinSize\",\"type\":\"uint32\"}],\"name\":\"AddNewContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// EthoFSControllerBin is the compiled bytecode used for deploying new contracts.
const EthoFSControllerBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a031990811633908117835560018054909216179055600455610b00806100436000396000f3006080604052600436106100955763ffffffff60e060020a60003504166306394c9b811461009a5780630ad8e54a146100bc57806331e8db11146100dc57806357618e1d146100fc5780635a58cd4c1461011c57806361268dc9146101315780638f490daf14610146578063a8f2efd814610166578063aa8da20b14610186578063d420a7e6146101a6578063d632972a146101c6575b600080fd5b3480156100a657600080fd5b506100ba6100b5366004610749565b6101e6565b005b3480156100c857600080fd5b506100ba6100d7366004610749565b61022c565b3480156100e857600080fd5b506100ba6100f7366004610895565b6102b2565b34801561010857600080fd5b506100ba61011736600461076f565b6102ce565b34801561012857600080fd5b506100ba6103d2565b34801561013d57600080fd5b506100ba6103f7565b34801561015257600080fd5b506100ba610161366004610749565b610464565b34801561017257600080fd5b506100ba61018136600461076f565b6104aa565b34801561019257600080fd5b506100ba6101a1366004610749565b6104f5565b3480156101b257600080fd5b506100ba6101c13660046107c1565b61053b565b3480156101d257600080fd5b506100ba6101e13660046107f1565b610588565b600054600160a060020a031633146101fd57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6003546040517f0ad8e54a000000000000000000000000000000000000000000000000000000008152600160a060020a0390911690630ad8e54a90610275908490600401610909565b600060405180830381600087803b15801561028f57600080fd5b505af11580156102a3573d6000803e3d6000fd5b505050506102af6103f7565b50565b600054600160a060020a031633146102c957600080fd5b600455565b6003546040517f49d590ef000000000000000000000000000000000000000000000000000000008152600160a060020a03909116906349d590ef90610319903390869060040161091d565b600060405180830381600087803b15801561033357600080fd5b505af1158015610347573d6000803e3d6000fd5b50506002546040517f3f0854a7000000000000000000000000000000000000000000000000000000008152600160a060020a039091169250633f0854a791506103949084906004016109e8565b600060405180830381600087803b1580156103ae57600080fd5b505af11580156103c2573d6000803e3d6000fd5b505050506103ce6103f7565b5050565b600054600160a060020a031633146103e957600080fd5b600054600160a060020a0316ff5b600360009054906101000a9004600160a060020a0316600160a060020a031663359a198b6040518163ffffffff1660e060020a028152600401600060405180830381600087803b15801561044a57600080fd5b505af115801561045e573d6000803e3d6000fd5b50505050565b600054600160a060020a0316331461047b57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6003546040517fa8f2efd8000000000000000000000000000000000000000000000000000000008152600160a060020a039091169063a8f2efd8906103949085908590600401610960565b600054600160a060020a0316331461050c57600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6003546040517f0cb2a1e6000000000000000000000000000000000000000000000000000000008152600160a060020a0390911690630cb2a1e69061039490339086908690600401610938565b600061b5ba63ffffffff86160463ffffffff16600454621000008463ffffffff168115156105b257fe5b0463ffffffff1602029050600360009054906101000a9004600160a060020a0316600160a060020a03166395c0e7313388888888876040518763ffffffff1660e060020a02815260040161060b96959493929190610980565b600060405180830381600087803b15801561062557600080fd5b505af1158015610639573d6000803e3d6000fd5b50506002546040517f8d036731000000000000000000000000000000000000000000000000000000008152600160a060020a039091169250638d036731915061068890869086906004016109f9565b600060405180830381600087803b1580156106a257600080fd5b505af11580156106b6573d6000803e3d6000fd5b505050506106c26103f7565b505050505050565b60006106d68235610a6c565b9392505050565b6000601f820183136106ee57600080fd5b81356107016106fc82610a40565b610a19565b9150808252602083016020830185838301111561071d57600080fd5b610728838284610a84565b50505092915050565b60006106d68235610a78565b60006106d68235610a7b565b60006020828403121561075b57600080fd5b600061076784846106ca565b949350505050565b6000806040838503121561078257600080fd5b600061078e85856106ca565b925050602083013567ffffffffffffffff8111156107ab57600080fd5b6107b7858286016106dd565b9150509250929050565b600080604083850312156107d457600080fd5b60006107e085856106ca565b92505060206107b78582860161073d565b600080600080600060a0868803121561080957600080fd5b853567ffffffffffffffff81111561082057600080fd5b61082c888289016106dd565b955050602061083d8882890161073d565b945050604061084e8882890161073d565b935050606086013567ffffffffffffffff81111561086b57600080fd5b610877888289016106dd565b92505060806108888882890161073d565b9150509295509295909350565b6000602082840312156108a757600080fd5b60006107678484610731565b6108bc81610a6c565b82525050565b60006108cd82610a68565b8084526108e1816020860160208601610a90565b6108ea81610abc565b9093016020019392505050565b6108bc81610a78565b6108bc81610a7b565b6020810161091782846108b3565b92915050565b6040810161092b82856108b3565b6106d660208301846108b3565b6060810161094682866108b3565b61095360208301856108b3565b6107676040830184610900565b6040810161096e82856108b3565b818103602083015261076781846108c2565b60c0810161098e82896108b3565b81810360208301526109a081886108c2565b90506109af6040830187610900565b6109bc6060830186610900565b81810360808301526109ce81856108c2565b90506109dd60a08301846108f7565b979650505050505050565b602080825281016106d681846108c2565b60408082528101610a0a81856108c2565b90506106d66020830184610900565b60405181810167ffffffffffffffff81118282101715610a3857600080fd5b604052919050565b600067ffffffffffffffff821115610a5757600080fd5b506020601f91909101601f19160190565b5190565b600160a060020a031690565b90565b63ffffffff1690565b82818337506000910152565b60005b83811015610aab578181015183820152602001610a93565b8381111561045e5750506000910152565b601f01601f1916905600a265627a7a72305820f61394914ac7b5575472099bcffc02deb8d4a75e974c20286e572be6132898726c6578706572696d656e74616cf50037`

// DeployEthoFSController deploys a new Ethereum contract, binding an instance of EthoFSController to it.
func DeployEthoFSController(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthoFSController, error) {
	parsed, err := abi.JSON(strings.NewReader(EthoFSControllerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthoFSControllerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthoFSController{EthoFSControllerCaller: EthoFSControllerCaller{contract: contract}, EthoFSControllerTransactor: EthoFSControllerTransactor{contract: contract}, EthoFSControllerFilterer: EthoFSControllerFilterer{contract: contract}}, nil
}

// EthoFSController is an auto generated Go binding around an Ethereum contract.
type EthoFSController struct {
	EthoFSControllerCaller     // Read-only binding to the contract
	EthoFSControllerTransactor // Write-only binding to the contract
	EthoFSControllerFilterer   // Log filterer for contract events
}

// EthoFSControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthoFSControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthoFSControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthoFSControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthoFSControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthoFSControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthoFSControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthoFSControllerSession struct {
	Contract     *EthoFSController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthoFSControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthoFSControllerCallerSession struct {
	Contract *EthoFSControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EthoFSControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthoFSControllerTransactorSession struct {
	Contract     *EthoFSControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EthoFSControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthoFSControllerRaw struct {
	Contract *EthoFSController // Generic contract binding to access the raw methods on
}

// EthoFSControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthoFSControllerCallerRaw struct {
	Contract *EthoFSControllerCaller // Generic read-only contract binding to access the raw methods on
}

// EthoFSControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthoFSControllerTransactorRaw struct {
	Contract *EthoFSControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthoFSController creates a new instance of EthoFSController, bound to a specific deployed contract.
func NewEthoFSController(address common.Address, backend bind.ContractBackend) (*EthoFSController, error) {
	contract, err := bindEthoFSController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthoFSController{EthoFSControllerCaller: EthoFSControllerCaller{contract: contract}, EthoFSControllerTransactor: EthoFSControllerTransactor{contract: contract}, EthoFSControllerFilterer: EthoFSControllerFilterer{contract: contract}}, nil
}

// NewEthoFSControllerCaller creates a new read-only instance of EthoFSController, bound to a specific deployed contract.
func NewEthoFSControllerCaller(address common.Address, caller bind.ContractCaller) (*EthoFSControllerCaller, error) {
	contract, err := bindEthoFSController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthoFSControllerCaller{contract: contract}, nil
}

// NewEthoFSControllerTransactor creates a new write-only instance of EthoFSController, bound to a specific deployed contract.
func NewEthoFSControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*EthoFSControllerTransactor, error) {
	contract, err := bindEthoFSController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthoFSControllerTransactor{contract: contract}, nil
}

// NewEthoFSControllerFilterer creates a new log filterer instance of EthoFSController, bound to a specific deployed contract.
func NewEthoFSControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*EthoFSControllerFilterer, error) {
	contract, err := bindEthoFSController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthoFSControllerFilterer{contract: contract}, nil
}

// bindEthoFSController binds a generic wrapper to an already deployed contract.
func bindEthoFSController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthoFSControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthoFSController *EthoFSControllerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthoFSController.Contract.EthoFSControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthoFSController *EthoFSControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthoFSController.Contract.EthoFSControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthoFSController *EthoFSControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthoFSController.Contract.EthoFSControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthoFSController *EthoFSControllerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthoFSController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthoFSController *EthoFSControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthoFSController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthoFSController *EthoFSControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthoFSController.Contract.contract.Transact(opts, method, params...)
}

// AddNewContract is a paid mutator transaction binding the contract method 0xd632972a.
//
// Solidity: function AddNewContract(string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize, string MainContentHash, uint32 pinSize) returns()
func (_EthoFSController *EthoFSControllerTransactor) AddNewContract(opts *bind.TransactOpts, HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32, MainContentHash string, pinSize uint32) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "AddNewContract", HostingContractName, HostingContractDuration, TotalContractSize, MainContentHash, pinSize)
}

// AddNewContract is a paid mutator transaction binding the contract method 0xd632972a.
//
// Solidity: function AddNewContract(string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize, string MainContentHash, uint32 pinSize) returns()
func (_EthoFSController *EthoFSControllerSession) AddNewContract(HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32, MainContentHash string, pinSize uint32) (*types.Transaction, error) {
	return _EthoFSController.Contract.AddNewContract(&_EthoFSController.TransactOpts, HostingContractName, HostingContractDuration, TotalContractSize, MainContentHash, pinSize)
}

// AddNewContract is a paid mutator transaction binding the contract method 0xd632972a.
//
// Solidity: function AddNewContract(string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize, string MainContentHash, uint32 pinSize) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) AddNewContract(HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32, MainContentHash string, pinSize uint32) (*types.Transaction, error) {
	return _EthoFSController.Contract.AddNewContract(&_EthoFSController.TransactOpts, HostingContractName, HostingContractDuration, TotalContractSize, MainContentHash, pinSize)
}

// AddNewUser is a paid mutator transaction binding the contract method 0xa8f2efd8.
//
// Solidity: function AddNewUser(address UserAddress, string AccountName) returns()
func (_EthoFSController *EthoFSControllerTransactor) AddNewUser(opts *bind.TransactOpts, UserAddress common.Address, AccountName string) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "AddNewUser", UserAddress, AccountName)
}

// AddNewUser is a paid mutator transaction binding the contract method 0xa8f2efd8.
//
// Solidity: function AddNewUser(address UserAddress, string AccountName) returns()
func (_EthoFSController *EthoFSControllerSession) AddNewUser(UserAddress common.Address, AccountName string) (*types.Transaction, error) {
	return _EthoFSController.Contract.AddNewUser(&_EthoFSController.TransactOpts, UserAddress, AccountName)
}

// AddNewUser is a paid mutator transaction binding the contract method 0xa8f2efd8.
//
// Solidity: function AddNewUser(address UserAddress, string AccountName) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) AddNewUser(UserAddress common.Address, AccountName string) (*types.Transaction, error) {
	return _EthoFSController.Contract.AddNewUser(&_EthoFSController.TransactOpts, UserAddress, AccountName)
}

// ExtendContract is a paid mutator transaction binding the contract method 0xd420a7e6.
//
// Solidity: function ExtendContract(address HostingContractAddress, uint32 HostingContractExtensionDuration) returns()
func (_EthoFSController *EthoFSControllerTransactor) ExtendContract(opts *bind.TransactOpts, HostingContractAddress common.Address, HostingContractExtensionDuration uint32) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "ExtendContract", HostingContractAddress, HostingContractExtensionDuration)
}

// ExtendContract is a paid mutator transaction binding the contract method 0xd420a7e6.
//
// Solidity: function ExtendContract(address HostingContractAddress, uint32 HostingContractExtensionDuration) returns()
func (_EthoFSController *EthoFSControllerSession) ExtendContract(HostingContractAddress common.Address, HostingContractExtensionDuration uint32) (*types.Transaction, error) {
	return _EthoFSController.Contract.ExtendContract(&_EthoFSController.TransactOpts, HostingContractAddress, HostingContractExtensionDuration)
}

// ExtendContract is a paid mutator transaction binding the contract method 0xd420a7e6.
//
// Solidity: function ExtendContract(address HostingContractAddress, uint32 HostingContractExtensionDuration) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) ExtendContract(HostingContractAddress common.Address, HostingContractExtensionDuration uint32) (*types.Transaction, error) {
	return _EthoFSController.Contract.ExtendContract(&_EthoFSController.TransactOpts, HostingContractAddress, HostingContractExtensionDuration)
}

// RemoveHostingContract is a paid mutator transaction binding the contract method 0x57618e1d.
//
// Solidity: function RemoveHostingContract(address HostingContractAddress, string MainContentHash) returns()
func (_EthoFSController *EthoFSControllerTransactor) RemoveHostingContract(opts *bind.TransactOpts, HostingContractAddress common.Address, MainContentHash string) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "RemoveHostingContract", HostingContractAddress, MainContentHash)
}

// RemoveHostingContract is a paid mutator transaction binding the contract method 0x57618e1d.
//
// Solidity: function RemoveHostingContract(address HostingContractAddress, string MainContentHash) returns()
func (_EthoFSController *EthoFSControllerSession) RemoveHostingContract(HostingContractAddress common.Address, MainContentHash string) (*types.Transaction, error) {
	return _EthoFSController.Contract.RemoveHostingContract(&_EthoFSController.TransactOpts, HostingContractAddress, MainContentHash)
}

// RemoveHostingContract is a paid mutator transaction binding the contract method 0x57618e1d.
//
// Solidity: function RemoveHostingContract(address HostingContractAddress, string MainContentHash) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) RemoveHostingContract(HostingContractAddress common.Address, MainContentHash string) (*types.Transaction, error) {
	return _EthoFSController.Contract.RemoveHostingContract(&_EthoFSController.TransactOpts, HostingContractAddress, MainContentHash)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x0ad8e54a.
//
// Solidity: function RemoveUser(address UserAddress) returns()
func (_EthoFSController *EthoFSControllerTransactor) RemoveUser(opts *bind.TransactOpts, UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "RemoveUser", UserAddress)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x0ad8e54a.
//
// Solidity: function RemoveUser(address UserAddress) returns()
func (_EthoFSController *EthoFSControllerSession) RemoveUser(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.RemoveUser(&_EthoFSController.TransactOpts, UserAddress)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x0ad8e54a.
//
// Solidity: function RemoveUser(address UserAddress) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) RemoveUser(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.RemoveUser(&_EthoFSController.TransactOpts, UserAddress)
}

// ScrubHostingContracts is a paid mutator transaction binding the contract method 0x61268dc9.
//
// Solidity: function ScrubHostingContracts() returns()
func (_EthoFSController *EthoFSControllerTransactor) ScrubHostingContracts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "ScrubHostingContracts")
}

// ScrubHostingContracts is a paid mutator transaction binding the contract method 0x61268dc9.
//
// Solidity: function ScrubHostingContracts() returns()
func (_EthoFSController *EthoFSControllerSession) ScrubHostingContracts() (*types.Transaction, error) {
	return _EthoFSController.Contract.ScrubHostingContracts(&_EthoFSController.TransactOpts)
}

// ScrubHostingContracts is a paid mutator transaction binding the contract method 0x61268dc9.
//
// Solidity: function ScrubHostingContracts() returns()
func (_EthoFSController *EthoFSControllerTransactorSession) ScrubHostingContracts() (*types.Transaction, error) {
	return _EthoFSController.Contract.ScrubHostingContracts(&_EthoFSController.TransactOpts)
}

// SetEthoFSDashboardAddress is a paid mutator transaction binding the contract method 0xaa8da20b.
//
// Solidity: function SetEthoFSDashboardAddress(address ethoFSDashboardAddress) returns()
func (_EthoFSController *EthoFSControllerTransactor) SetEthoFSDashboardAddress(opts *bind.TransactOpts, ethoFSDashboardAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "SetEthoFSDashboardAddress", ethoFSDashboardAddress)
}

// SetEthoFSDashboardAddress is a paid mutator transaction binding the contract method 0xaa8da20b.
//
// Solidity: function SetEthoFSDashboardAddress(address ethoFSDashboardAddress) returns()
func (_EthoFSController *EthoFSControllerSession) SetEthoFSDashboardAddress(ethoFSDashboardAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.SetEthoFSDashboardAddress(&_EthoFSController.TransactOpts, ethoFSDashboardAddress)
}

// SetEthoFSDashboardAddress is a paid mutator transaction binding the contract method 0xaa8da20b.
//
// Solidity: function SetEthoFSDashboardAddress(address ethoFSDashboardAddress) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) SetEthoFSDashboardAddress(ethoFSDashboardAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.SetEthoFSDashboardAddress(&_EthoFSController.TransactOpts, ethoFSDashboardAddress)
}

// SetHostingCost is a paid mutator transaction binding the contract method 0x31e8db11.
//
// Solidity: function SetHostingCost(uint256 hostingCost) returns()
func (_EthoFSController *EthoFSControllerTransactor) SetHostingCost(opts *bind.TransactOpts, hostingCost *big.Int) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "SetHostingCost", hostingCost)
}

// SetHostingCost is a paid mutator transaction binding the contract method 0x31e8db11.
//
// Solidity: function SetHostingCost(uint256 hostingCost) returns()
func (_EthoFSController *EthoFSControllerSession) SetHostingCost(hostingCost *big.Int) (*types.Transaction, error) {
	return _EthoFSController.Contract.SetHostingCost(&_EthoFSController.TransactOpts, hostingCost)
}

// SetHostingCost is a paid mutator transaction binding the contract method 0x31e8db11.
//
// Solidity: function SetHostingCost(uint256 hostingCost) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) SetHostingCost(hostingCost *big.Int) (*types.Transaction, error) {
	return _EthoFSController.Contract.SetHostingCost(&_EthoFSController.TransactOpts, hostingCost)
}

// SetPinStorageAddress is a paid mutator transaction binding the contract method 0x8f490daf.
//
// Solidity: function SetPinStorageAddress(address pinStorageAddress) returns()
func (_EthoFSController *EthoFSControllerTransactor) SetPinStorageAddress(opts *bind.TransactOpts, pinStorageAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "SetPinStorageAddress", pinStorageAddress)
}

// SetPinStorageAddress is a paid mutator transaction binding the contract method 0x8f490daf.
//
// Solidity: function SetPinStorageAddress(address pinStorageAddress) returns()
func (_EthoFSController *EthoFSControllerSession) SetPinStorageAddress(pinStorageAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.SetPinStorageAddress(&_EthoFSController.TransactOpts, pinStorageAddress)
}

// SetPinStorageAddress is a paid mutator transaction binding the contract method 0x8f490daf.
//
// Solidity: function SetPinStorageAddress(address pinStorageAddress) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) SetPinStorageAddress(pinStorageAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.SetPinStorageAddress(&_EthoFSController.TransactOpts, pinStorageAddress)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address newOperator) returns()
func (_EthoFSController *EthoFSControllerTransactor) ChangeOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "changeOperator", newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address newOperator) returns()
func (_EthoFSController *EthoFSControllerSession) ChangeOperator(newOperator common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.ChangeOperator(&_EthoFSController.TransactOpts, newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address newOperator) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) ChangeOperator(newOperator common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.ChangeOperator(&_EthoFSController.TransactOpts, newOperator)
}

// DeleteContract is a paid mutator transaction binding the contract method 0x5a58cd4c.
//
// Solidity: function deleteContract() returns()
func (_EthoFSController *EthoFSControllerTransactor) DeleteContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "deleteContract")
}

// DeleteContract is a paid mutator transaction binding the contract method 0x5a58cd4c.
//
// Solidity: function deleteContract() returns()
func (_EthoFSController *EthoFSControllerSession) DeleteContract() (*types.Transaction, error) {
	return _EthoFSController.Contract.DeleteContract(&_EthoFSController.TransactOpts)
}

// DeleteContract is a paid mutator transaction binding the contract method 0x5a58cd4c.
//
// Solidity: function deleteContract() returns()
func (_EthoFSController *EthoFSControllerTransactorSession) DeleteContract() (*types.Transaction, error) {
	return _EthoFSController.Contract.DeleteContract(&_EthoFSController.TransactOpts)
}

// EthoFSDashboardABI is the input ABI used to generate the binding from.
const EthoFSDashboardABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"}],\"name\":\"RemoveUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"},{\"name\":\"HostingContractAddress\",\"type\":\"address\"},{\"name\":\"HostingContractExtensionDuration\",\"type\":\"uint32\"}],\"name\":\"ExtendHostingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ScrubContractList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"},{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"RemoveHostingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"},{\"name\":\"HostingContractName\",\"type\":\"string\"},{\"name\":\"HostingContractDuration\",\"type\":\"uint32\"},{\"name\":\"TotalContractSize\",\"type\":\"uint32\"},{\"name\":\"MainContentHash\",\"type\":\"string\"},{\"name\":\"ContractCost\",\"type\":\"uint256\"}],\"name\":\"AddHostingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"},{\"name\":\"AccountName\",\"type\":\"string\"}],\"name\":\"AddNewUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthoFSDashboardBin is the compiled bytecode used for deploying new contracts.
const EthoFSDashboardBin = `0x`

// DeployEthoFSDashboard deploys a new Ethereum contract, binding an instance of EthoFSDashboard to it.
func DeployEthoFSDashboard(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthoFSDashboard, error) {
	parsed, err := abi.JSON(strings.NewReader(EthoFSDashboardABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthoFSDashboardBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthoFSDashboard{EthoFSDashboardCaller: EthoFSDashboardCaller{contract: contract}, EthoFSDashboardTransactor: EthoFSDashboardTransactor{contract: contract}, EthoFSDashboardFilterer: EthoFSDashboardFilterer{contract: contract}}, nil
}

// EthoFSDashboard is an auto generated Go binding around an Ethereum contract.
type EthoFSDashboard struct {
	EthoFSDashboardCaller     // Read-only binding to the contract
	EthoFSDashboardTransactor // Write-only binding to the contract
	EthoFSDashboardFilterer   // Log filterer for contract events
}

// EthoFSDashboardCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthoFSDashboardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthoFSDashboardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthoFSDashboardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthoFSDashboardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthoFSDashboardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthoFSDashboardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthoFSDashboardSession struct {
	Contract     *EthoFSDashboard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthoFSDashboardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthoFSDashboardCallerSession struct {
	Contract *EthoFSDashboardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EthoFSDashboardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthoFSDashboardTransactorSession struct {
	Contract     *EthoFSDashboardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EthoFSDashboardRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthoFSDashboardRaw struct {
	Contract *EthoFSDashboard // Generic contract binding to access the raw methods on
}

// EthoFSDashboardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthoFSDashboardCallerRaw struct {
	Contract *EthoFSDashboardCaller // Generic read-only contract binding to access the raw methods on
}

// EthoFSDashboardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthoFSDashboardTransactorRaw struct {
	Contract *EthoFSDashboardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthoFSDashboard creates a new instance of EthoFSDashboard, bound to a specific deployed contract.
func NewEthoFSDashboard(address common.Address, backend bind.ContractBackend) (*EthoFSDashboard, error) {
	contract, err := bindEthoFSDashboard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthoFSDashboard{EthoFSDashboardCaller: EthoFSDashboardCaller{contract: contract}, EthoFSDashboardTransactor: EthoFSDashboardTransactor{contract: contract}, EthoFSDashboardFilterer: EthoFSDashboardFilterer{contract: contract}}, nil
}

// NewEthoFSDashboardCaller creates a new read-only instance of EthoFSDashboard, bound to a specific deployed contract.
func NewEthoFSDashboardCaller(address common.Address, caller bind.ContractCaller) (*EthoFSDashboardCaller, error) {
	contract, err := bindEthoFSDashboard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthoFSDashboardCaller{contract: contract}, nil
}

// NewEthoFSDashboardTransactor creates a new write-only instance of EthoFSDashboard, bound to a specific deployed contract.
func NewEthoFSDashboardTransactor(address common.Address, transactor bind.ContractTransactor) (*EthoFSDashboardTransactor, error) {
	contract, err := bindEthoFSDashboard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthoFSDashboardTransactor{contract: contract}, nil
}

// NewEthoFSDashboardFilterer creates a new log filterer instance of EthoFSDashboard, bound to a specific deployed contract.
func NewEthoFSDashboardFilterer(address common.Address, filterer bind.ContractFilterer) (*EthoFSDashboardFilterer, error) {
	contract, err := bindEthoFSDashboard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthoFSDashboardFilterer{contract: contract}, nil
}

// bindEthoFSDashboard binds a generic wrapper to an already deployed contract.
func bindEthoFSDashboard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthoFSDashboardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthoFSDashboard *EthoFSDashboardRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthoFSDashboard.Contract.EthoFSDashboardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthoFSDashboard *EthoFSDashboardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.EthoFSDashboardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthoFSDashboard *EthoFSDashboardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.EthoFSDashboardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthoFSDashboard *EthoFSDashboardCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthoFSDashboard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthoFSDashboard *EthoFSDashboardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthoFSDashboard *EthoFSDashboardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.contract.Transact(opts, method, params...)
}

// AddHostingContract is a paid mutator transaction binding the contract method 0x95c0e731.
//
// Solidity: function AddHostingContract(address UserAddress, string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize, string MainContentHash, uint256 ContractCost) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactor) AddHostingContract(opts *bind.TransactOpts, UserAddress common.Address, HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32, MainContentHash string, ContractCost *big.Int) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "AddHostingContract", UserAddress, HostingContractName, HostingContractDuration, TotalContractSize, MainContentHash, ContractCost)
}

// AddHostingContract is a paid mutator transaction binding the contract method 0x95c0e731.
//
// Solidity: function AddHostingContract(address UserAddress, string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize, string MainContentHash, uint256 ContractCost) returns()
func (_EthoFSDashboard *EthoFSDashboardSession) AddHostingContract(UserAddress common.Address, HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32, MainContentHash string, ContractCost *big.Int) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.AddHostingContract(&_EthoFSDashboard.TransactOpts, UserAddress, HostingContractName, HostingContractDuration, TotalContractSize, MainContentHash, ContractCost)
}

// AddHostingContract is a paid mutator transaction binding the contract method 0x95c0e731.
//
// Solidity: function AddHostingContract(address UserAddress, string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize, string MainContentHash, uint256 ContractCost) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) AddHostingContract(UserAddress common.Address, HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32, MainContentHash string, ContractCost *big.Int) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.AddHostingContract(&_EthoFSDashboard.TransactOpts, UserAddress, HostingContractName, HostingContractDuration, TotalContractSize, MainContentHash, ContractCost)
}

// AddNewUser is a paid mutator transaction binding the contract method 0xa8f2efd8.
//
// Solidity: function AddNewUser(address UserAddress, string AccountName) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactor) AddNewUser(opts *bind.TransactOpts, UserAddress common.Address, AccountName string) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "AddNewUser", UserAddress, AccountName)
}

// AddNewUser is a paid mutator transaction binding the contract method 0xa8f2efd8.
//
// Solidity: function AddNewUser(address UserAddress, string AccountName) returns()
func (_EthoFSDashboard *EthoFSDashboardSession) AddNewUser(UserAddress common.Address, AccountName string) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.AddNewUser(&_EthoFSDashboard.TransactOpts, UserAddress, AccountName)
}

// AddNewUser is a paid mutator transaction binding the contract method 0xa8f2efd8.
//
// Solidity: function AddNewUser(address UserAddress, string AccountName) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) AddNewUser(UserAddress common.Address, AccountName string) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.AddNewUser(&_EthoFSDashboard.TransactOpts, UserAddress, AccountName)
}

// ExtendHostingContract is a paid mutator transaction binding the contract method 0x0cb2a1e6.
//
// Solidity: function ExtendHostingContract(address UserAddress, address HostingContractAddress, uint32 HostingContractExtensionDuration) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactor) ExtendHostingContract(opts *bind.TransactOpts, UserAddress common.Address, HostingContractAddress common.Address, HostingContractExtensionDuration uint32) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "ExtendHostingContract", UserAddress, HostingContractAddress, HostingContractExtensionDuration)
}

// ExtendHostingContract is a paid mutator transaction binding the contract method 0x0cb2a1e6.
//
// Solidity: function ExtendHostingContract(address UserAddress, address HostingContractAddress, uint32 HostingContractExtensionDuration) returns()
func (_EthoFSDashboard *EthoFSDashboardSession) ExtendHostingContract(UserAddress common.Address, HostingContractAddress common.Address, HostingContractExtensionDuration uint32) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.ExtendHostingContract(&_EthoFSDashboard.TransactOpts, UserAddress, HostingContractAddress, HostingContractExtensionDuration)
}

// ExtendHostingContract is a paid mutator transaction binding the contract method 0x0cb2a1e6.
//
// Solidity: function ExtendHostingContract(address UserAddress, address HostingContractAddress, uint32 HostingContractExtensionDuration) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) ExtendHostingContract(UserAddress common.Address, HostingContractAddress common.Address, HostingContractExtensionDuration uint32) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.ExtendHostingContract(&_EthoFSDashboard.TransactOpts, UserAddress, HostingContractAddress, HostingContractExtensionDuration)
}

// RemoveHostingContract is a paid mutator transaction binding the contract method 0x49d590ef.
//
// Solidity: function RemoveHostingContract(address UserAddress, address HostingContractAddress) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactor) RemoveHostingContract(opts *bind.TransactOpts, UserAddress common.Address, HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "RemoveHostingContract", UserAddress, HostingContractAddress)
}

// RemoveHostingContract is a paid mutator transaction binding the contract method 0x49d590ef.
//
// Solidity: function RemoveHostingContract(address UserAddress, address HostingContractAddress) returns()
func (_EthoFSDashboard *EthoFSDashboardSession) RemoveHostingContract(UserAddress common.Address, HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.RemoveHostingContract(&_EthoFSDashboard.TransactOpts, UserAddress, HostingContractAddress)
}

// RemoveHostingContract is a paid mutator transaction binding the contract method 0x49d590ef.
//
// Solidity: function RemoveHostingContract(address UserAddress, address HostingContractAddress) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) RemoveHostingContract(UserAddress common.Address, HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.RemoveHostingContract(&_EthoFSDashboard.TransactOpts, UserAddress, HostingContractAddress)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x0ad8e54a.
//
// Solidity: function RemoveUser(address UserAddress) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactor) RemoveUser(opts *bind.TransactOpts, UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "RemoveUser", UserAddress)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x0ad8e54a.
//
// Solidity: function RemoveUser(address UserAddress) returns()
func (_EthoFSDashboard *EthoFSDashboardSession) RemoveUser(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.RemoveUser(&_EthoFSDashboard.TransactOpts, UserAddress)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x0ad8e54a.
//
// Solidity: function RemoveUser(address UserAddress) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) RemoveUser(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.RemoveUser(&_EthoFSDashboard.TransactOpts, UserAddress)
}

// ScrubContractList is a paid mutator transaction binding the contract method 0x359a198b.
//
// Solidity: function ScrubContractList() returns()
func (_EthoFSDashboard *EthoFSDashboardTransactor) ScrubContractList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "ScrubContractList")
}

// ScrubContractList is a paid mutator transaction binding the contract method 0x359a198b.
//
// Solidity: function ScrubContractList() returns()
func (_EthoFSDashboard *EthoFSDashboardSession) ScrubContractList() (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.ScrubContractList(&_EthoFSDashboard.TransactOpts)
}

// ScrubContractList is a paid mutator transaction binding the contract method 0x359a198b.
//
// Solidity: function ScrubContractList() returns()
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) ScrubContractList() (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.ScrubContractList(&_EthoFSDashboard.TransactOpts)
}
