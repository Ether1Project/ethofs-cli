// Copyright 2020 The Etho.Black Development Team

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
const EthoFSControllerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"}],\"name\":\"RemoveUserOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"}],\"name\":\"CheckAccountExistence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hostingCost\",\"type\":\"uint256\"}],\"name\":\"SetHostingCost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetHostingContractExpirationBlockHeight\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetHostingContractDeployedBlockHeight\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"},{\"name\":\"AccountName\",\"type\":\"string\"}],\"name\":\"AddNewUserOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"},{\"name\":\"MainContentHash\",\"type\":\"string\"}],\"name\":\"RemoveHostingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ScrubHostingContracts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetHostingContractStorageUsed\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"},{\"name\":\"ArrayKey\",\"type\":\"uint256\"}],\"name\":\"GetHostingContractAddress\",\"outputs\":[{\"name\":\"value\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"HostingCost\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetHostingContractName\",\"outputs\":[{\"name\":\"value\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pinStorageAddress\",\"type\":\"address\"}],\"name\":\"SetPinStorageAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ethoFSHostingContractsAddress\",\"type\":\"address\"}],\"name\":\"SetEthoFSHostingContractsAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"}],\"name\":\"GetUserAccountTotalContractCount\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ethoFSDashboardAddress\",\"type\":\"address\"}],\"name\":\"SetEthoFSDashboardAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetContentHashString\",\"outputs\":[{\"name\":\"ContentHashString\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"AccountName\",\"type\":\"string\"}],\"name\":\"AddNewUserPublic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"set\",\"type\":\"address\"}],\"name\":\"SetAccountCollectionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"MainContentHash\",\"type\":\"string\"},{\"name\":\"HostingContractName\",\"type\":\"string\"},{\"name\":\"HostingContractDuration\",\"type\":\"uint32\"},{\"name\":\"TotalContractSize\",\"type\":\"uint32\"},{\"name\":\"pinSize\",\"type\":\"uint32\"},{\"name\":\"ContentHashString\",\"type\":\"string\"},{\"name\":\"ContentPathString\",\"type\":\"string\"}],\"name\":\"AddNewContract\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetContentPathString\",\"outputs\":[{\"name\":\"ContentPathString\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"RemoveUserPublic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"},{\"name\":\"HostingContractExtensionDuration\",\"type\":\"uint32\"}],\"name\":\"ExtendContract\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"}],\"name\":\"GetUserAccountActiveContractCount\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"}],\"name\":\"GetUserAccountName\",\"outputs\":[{\"name\":\"value\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetMainContentHash\",\"outputs\":[{\"name\":\"MainContentHash\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// EthoFSControllerBin is the compiled bytecode used for deploying new contracts.
const EthoFSControllerBin = `0x608060405234801561001057600080fd5b506000805433600160a060020a03199182168117835560018054909216179055678ac7230489e80000600655611a7690819061004c90396000f3006080604052600436106101695763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306394c9b811461016e5780630bbeb35b146101905780632c3027a7146101b057806331e8db11146101e6578063403169ac14610206578063415dfefa146102335780634a6b2fa31461025357806357618e1d146102735780635a58cd4c1461029357806361268dc9146102a857806366b7004c146102bd57806368e15933146102ea5780637201f738146103175780638e0976141461032c5780638f490daf14610359578063978a627a146103795780639b0a9add14610399578063aa8da20b146103b9578063af7bbf5a146103d9578063b18759de146103f9578063c8ac5fe814610419578063ccb726b114610439578063d0cd25131461044c578063d19237631461046c578063d420a7e614610481578063d4f0d92914610494578063e42f0027146104b4578063e6dc7817146104d4575b600080fd5b34801561017a57600080fd5b5061018e610189366004611490565b6104f4565b005b34801561019c57600080fd5b5061018e6101ab366004611490565b61053a565b3480156101bc57600080fd5b506101d06101cb366004611490565b6105ce565b6040516101dd91906118bc565b60405180910390f35b3480156101f257600080fd5b5061018e610201366004611709565b610671565b34801561021257600080fd5b50610226610221366004611490565b61070b565b6040516101dd919061196e565b34801561023f57600080fd5b5061022661024e366004611490565b6107a7565b34801561025f57600080fd5b5061018e61026e3660046114d4565b6107f1565b34801561027f57600080fd5b5061018e61028e3660046114d4565b610888565b34801561029f57600080fd5b5061018e610a40565b3480156102b457600080fd5b5061018e610a65565b3480156102c957600080fd5b506102dd6102d8366004611490565b610aeb565b6040516101dd919061197c565b3480156102f657600080fd5b5061030a610305366004611526565b610b87565b6040516101dd91906117c2565b34801561032357600080fd5b50610226610c2d565b34801561033857600080fd5b5061034c610347366004611490565b610c33565b6040516101dd91906118ca565b34801561036557600080fd5b5061018e610374366004611490565b610cd3565b34801561038557600080fd5b5061018e610394366004611490565b610d19565b3480156103a557600080fd5b506102dd6103b4366004611490565b610d5f565b3480156103c557600080fd5b5061018e6103d4366004611490565b610daa565b3480156103e557600080fd5b5061034c6103f4366004611490565b610df0565b34801561040557600080fd5b5061018e6104143660046115a4565b610e3a565b34801561042557600080fd5b5061018e610434366004611490565b610e84565b61018e61044736600461160e565b610fc0565b34801561045857600080fd5b5061034c610467366004611490565b61118b565b34801561047857600080fd5b5061018e6111d5565b61018e61048f366004611556565b611237565b3480156104a057600080fd5b506102dd6104af366004611490565b6112c2565b3480156104c057600080fd5b5061034c6104cf366004611490565b61130d565b3480156104e057600080fd5b5061034c6104ef366004611490565b611358565b600054600160a060020a0316331461050b57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461055157600080fd5b600480546040517f0ad8e54a000000000000000000000000000000000000000000000000000000008152600160a060020a0390911691630ad8e54a91610599918591016117c2565b600060405180830381600087803b1580156105b357600080fd5b505af11580156105c7573d6000803e3d6000fd5b5050505050565b600480546040517f2c3027a7000000000000000000000000000000000000000000000000000000008152600092600160a060020a0390921691632c3027a791610619918691016117c2565b602060405180830381600087803b15801561063357600080fd5b505af1158015610647573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061066b9190810190611586565b92915050565b600054600160a060020a0316331461068857600080fd5b6005546040517f0b6ff6dd000000000000000000000000000000000000000000000000000000008152600160a060020a0390911690630b6ff6dd906106d190849060040161196e565b600060405180830381600087803b1580156106eb57600080fd5b505af11580156106ff573d6000803e3d6000fd5b50505060069190915550565b6005546040517f403169ac000000000000000000000000000000000000000000000000000000008152600091600160a060020a03169063403169ac906107559085906004016117c2565b602060405180830381600087803b15801561076f57600080fd5b505af1158015610783573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061066b9190810190611727565b6005546040517f415dfefa000000000000000000000000000000000000000000000000000000008152600091600160a060020a03169063415dfefa906107559085906004016117c2565b600054600160a060020a0316331461080857600080fd5b600480546040517fa8f2efd8000000000000000000000000000000000000000000000000000000008152600160a060020a039091169163a8f2efd891610852918691869101611866565b600060405180830381600087803b15801561086c57600080fd5b505af1158015610880573d6000803e3d6000fd5b505050505050565b6005546002546040517f782d8e06000000000000000000000000000000000000000000000000000000008152600092600160a060020a039081169263782d8e06926108dc92339289929116906004016117eb565b602060405180830381600087803b1580156108f657600080fd5b505af115801561090a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061092e9190810190611586565b905060018115151415610a3357600480546040517f49d590ef000000000000000000000000000000000000000000000000000000008152600160a060020a03909116916349d590ef916109859133918891016117d0565b600060405180830381600087803b15801561099f57600080fd5b505af11580156109b3573d6000803e3d6000fd5b50506003546040517f3f0854a7000000000000000000000000000000000000000000000000000000008152600160a060020a039091169250633f0854a79150610a009085906004016118ca565b600060405180830381600087803b158015610a1a57600080fd5b505af1158015610a2e573d6000803e3d6000fd5b505050505b610a3b610a65565b505050565b600054600160a060020a03163314610a5757600080fd5b600054600160a060020a0316ff5b600560009054906101000a9004600160a060020a0316600160a060020a031663359a198b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b158015610ad157600080fd5b505af1158015610ae5573d6000803e3d6000fd5b50505050565b6005546040517f66b7004c000000000000000000000000000000000000000000000000000000008152600091600160a060020a0316906366b7004c90610b359085906004016117c2565b602060405180830381600087803b158015610b4f57600080fd5b505af1158015610b63573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061066b9190810190611745565b600480546040517f68e15933000000000000000000000000000000000000000000000000000000008152600092600160a060020a03909216916368e1593391610bd4918791879101611886565b602060405180830381600087803b158015610bee57600080fd5b505af1158015610c02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610c2691908101906114b6565b9392505050565b60065481565b6005546040517f8e097614000000000000000000000000000000000000000000000000000000008152606091600160a060020a031690638e09761490610c7d9085906004016117c2565b600060405180830381600087803b158015610c9757600080fd5b505af1158015610cab573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261066b91908101906115d9565b600054600160a060020a03163314610cea57600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a03163314610d3057600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600480546040517f9b0a9add000000000000000000000000000000000000000000000000000000008152600092600160a060020a0390921691639b0a9add91610b35918691016117c2565b600054600160a060020a03163314610dc157600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6005546040517faf7bbf5a000000000000000000000000000000000000000000000000000000008152606091600160a060020a03169063af7bbf5a90610c7d9085906004016117c2565b600480546040517fa8f2efd8000000000000000000000000000000000000000000000000000000008152600160a060020a039091169163a8f2efd891610599913391869101611866565b600054600160a060020a03163314610e9b57600080fd5b6005546040517fc8ac5fe8000000000000000000000000000000000000000000000000000000008152600160a060020a039091169063c8ac5fe890610ee49084906004016117c2565b600060405180830381600087803b158015610efe57600080fd5b505af1158015610f12573d6000803e3d6000fd5b5050600480546040517fc8ac5fe8000000000000000000000000000000000000000000000000000000008152600160a060020a03909116935063c8ac5fe89250610f5e918591016117c2565b600060405180830381600087803b158015610f7857600080fd5b505af1158015610f8c573d6000803e3d6000fd5b50506002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0394909416939093179092555050565b6005546040517f4b8f5af6000000000000000000000000000000000000000000000000000000008152600091600160a060020a031690634b8f5af6903490611016908c908c908c908c908b908b906004016118db565b6020604051808303818588803b15801561102f57600080fd5b505af1158015611043573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525061106891908101906114b6565b9050600160a060020a0381161561117957600480546040517f57199193000000000000000000000000000000000000000000000000000000008152600160a060020a03909116916357199193916110c991859133918d918d918d9101611813565b600060405180830381600087803b1580156110e357600080fd5b505af11580156110f7573d6000803e3d6000fd5b50506003546040517f8d036731000000000000000000000000000000000000000000000000000000008152600160a060020a039091169250638d0367319150611146908b90889060040161194e565b600060405180830381600087803b15801561116057600080fd5b505af1158015611174573d6000803e3d6000fd5b505050505b611181610a65565b5050505050505050565b6005546040517fd0cd2513000000000000000000000000000000000000000000000000000000008152606091600160a060020a03169063d0cd251390610c7d9085906004016117c2565b600480546040517f0ad8e54a000000000000000000000000000000000000000000000000000000008152600160a060020a0390911691630ad8e54a9161121d913391016117c2565b600060405180830381600087803b158015610ad157600080fd5b6005546040517fde5a8ec5000000000000000000000000000000000000000000000000000000008152600160a060020a039091169063de5a8ec590349061128490869086906004016118a1565b6000604051808303818588803b15801561129d57600080fd5b505af11580156112b1573d6000803e3d6000fd5b50505050506112be610a65565b5050565b600480546040517fd4f0d929000000000000000000000000000000000000000000000000000000008152600092600160a060020a039092169163d4f0d92991610b35918691016117c2565b600480546040517fe42f0027000000000000000000000000000000000000000000000000000000008152606092600160a060020a039092169163e42f002791610c7d918691016117c2565b6005546040517fe6dc7817000000000000000000000000000000000000000000000000000000008152606091600160a060020a03169063e6dc781790610c7d9085906004016117c2565b6000610c2682356119dd565b6000610c2682516119dd565b6000610c2682516119e9565b6000601f820183136113d757600080fd5b81356113ea6113e5826119b1565b61198a565b9150808252602083016020830185838301111561140657600080fd5b6114118382846119fa565b50505092915050565b6000601f8201831361142b57600080fd5b81516114396113e5826119b1565b9150808252602083016020830185838301111561145557600080fd5b611411838284611a06565b6000610c2682356119ee565b6000610c2682516119ee565b6000610c2682356119f1565b6000610c2682516119f1565b6000602082840312156114a257600080fd5b60006114ae84846113a2565b949350505050565b6000602082840312156114c857600080fd5b60006114ae84846113ae565b600080604083850312156114e757600080fd5b60006114f385856113a2565b925050602083013567ffffffffffffffff81111561151057600080fd5b61151c858286016113c6565b9150509250929050565b6000806040838503121561153957600080fd5b600061154585856113a2565b925050602061151c85828601611460565b6000806040838503121561156957600080fd5b600061157585856113a2565b925050602061151c85828601611478565b60006020828403121561159857600080fd5b60006114ae84846113ba565b6000602082840312156115b657600080fd5b813567ffffffffffffffff8111156115cd57600080fd5b6114ae848285016113c6565b6000602082840312156115eb57600080fd5b815167ffffffffffffffff81111561160257600080fd5b6114ae8482850161141a565b600080600080600080600060e0888a03121561162957600080fd5b873567ffffffffffffffff81111561164057600080fd5b61164c8a828b016113c6565b975050602088013567ffffffffffffffff81111561166957600080fd5b6116758a828b016113c6565b96505060406116868a828b01611478565b95505060606116978a828b01611478565b94505060806116a88a828b01611478565b93505060a088013567ffffffffffffffff8111156116c557600080fd5b6116d18a828b016113c6565b92505060c088013567ffffffffffffffff8111156116ee57600080fd5b6116fa8a828b016113c6565b91505092959891949750929550565b60006020828403121561171b57600080fd5b60006114ae8484611460565b60006020828403121561173957600080fd5b60006114ae848461146c565b60006020828403121561175757600080fd5b60006114ae8484611484565b61176c816119dd565b82525050565b61176c816119e9565b6000611786826119d9565b80845261179a816020860160208601611a06565b6117a381611a32565b9093016020019392505050565b61176c816119ee565b61176c816119f1565b6020810161066b8284611763565b604081016117de8285611763565b610c266020830184611763565b606081016117f98286611763565b6118066020830185611763565b6114ae6040830184611763565b60a081016118218288611763565b61182e6020830187611763565b8181036040830152611840818661177b565b905061184f60608301856117b9565b61185c60808301846117b9565b9695505050505050565b604081016118748285611763565b81810360208301526114ae818461177b565b604081016118948285611763565b610c2660208301846117b0565b604081016118af8285611763565b610c2660208301846117b9565b6020810161066b8284611772565b60208082528101610c26818461177b565b60c080825281016118ec818961177b565b90508181036020830152611900818861177b565b905061190f60408301876117b9565b61191c60608301866117b9565b818103608083015261192e818561177b565b905081810360a0830152611942818461177b565b98975050505050505050565b6040808252810161195f818561177b565b9050610c2660208301846117b9565b6020810161066b82846117b0565b6020810161066b82846117b9565b60405181810167ffffffffffffffff811182821017156119a957600080fd5b604052919050565b600067ffffffffffffffff8211156119c857600080fd5b506020601f91909101601f19160190565b5190565b600160a060020a031690565b151590565b90565b63ffffffff1690565b82818337506000910152565b60005b83811015611a21578181015183820152602001611a09565b83811115610ae55750506000910152565b601f01601f1916905600a265627a7a72305820ca51dfb1808e2408964eebbdccca3e3feca0cb5f07489f58b96b614dd8efefb66c6578706572696d656e74616cf50037`

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

// CheckAccountExistence is a free data retrieval call binding the contract method 0x2c3027a7.
//
// Solidity: function CheckAccountExistence(address UserAddress) constant returns(bool)
func (_EthoFSController *EthoFSControllerCaller) CheckAccountExistence(opts *bind.CallOpts, UserAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthoFSController.contract.Call(opts, out, "CheckAccountExistence", UserAddress)
	return *ret0, err
}

// CheckAccountExistence is a free data retrieval call binding the contract method 0x2c3027a7.
//
// Solidity: function CheckAccountExistence(address UserAddress) constant returns(bool)
func (_EthoFSController *EthoFSControllerSession) CheckAccountExistence(UserAddress common.Address) (bool, error) {
	return _EthoFSController.Contract.CheckAccountExistence(&_EthoFSController.CallOpts, UserAddress)
}

// CheckAccountExistence is a free data retrieval call binding the contract method 0x2c3027a7.
//
// Solidity: function CheckAccountExistence(address UserAddress) constant returns(bool)
func (_EthoFSController *EthoFSControllerCallerSession) CheckAccountExistence(UserAddress common.Address) (bool, error) {
	return _EthoFSController.Contract.CheckAccountExistence(&_EthoFSController.CallOpts, UserAddress)
}

// GetContentHashString is a free data retrieval call binding the contract method 0xaf7bbf5a.
//
// Solidity: function GetContentHashString(address HostingContractAddress) constant returns(string ContentHashString)
func (_EthoFSController *EthoFSControllerCaller) GetContentHashString(opts *bind.CallOpts, HostingContractAddress common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EthoFSController.contract.Call(opts, out, "GetContentHashString", HostingContractAddress)
	return *ret0, err
}

// GetContentHashString is a free data retrieval call binding the contract method 0xaf7bbf5a.
//
// Solidity: function GetContentHashString(address HostingContractAddress) constant returns(string ContentHashString)
func (_EthoFSController *EthoFSControllerSession) GetContentHashString(HostingContractAddress common.Address) (string, error) {
	return _EthoFSController.Contract.GetContentHashString(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetContentHashString is a free data retrieval call binding the contract method 0xaf7bbf5a.
//
// Solidity: function GetContentHashString(address HostingContractAddress) constant returns(string ContentHashString)
func (_EthoFSController *EthoFSControllerCallerSession) GetContentHashString(HostingContractAddress common.Address) (string, error) {
	return _EthoFSController.Contract.GetContentHashString(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetContentPathString is a free data retrieval call binding the contract method 0xd0cd2513.
//
// Solidity: function GetContentPathString(address HostingContractAddress) constant returns(string ContentPathString)
func (_EthoFSController *EthoFSControllerCaller) GetContentPathString(opts *bind.CallOpts, HostingContractAddress common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EthoFSController.contract.Call(opts, out, "GetContentPathString", HostingContractAddress)
	return *ret0, err
}

// GetContentPathString is a free data retrieval call binding the contract method 0xd0cd2513.
//
// Solidity: function GetContentPathString(address HostingContractAddress) constant returns(string ContentPathString)
func (_EthoFSController *EthoFSControllerSession) GetContentPathString(HostingContractAddress common.Address) (string, error) {
	return _EthoFSController.Contract.GetContentPathString(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetContentPathString is a free data retrieval call binding the contract method 0xd0cd2513.
//
// Solidity: function GetContentPathString(address HostingContractAddress) constant returns(string ContentPathString)
func (_EthoFSController *EthoFSControllerCallerSession) GetContentPathString(HostingContractAddress common.Address) (string, error) {
	return _EthoFSController.Contract.GetContentPathString(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetHostingContractAddress is a free data retrieval call binding the contract method 0x68e15933.
//
// Solidity: function GetHostingContractAddress(address UserAddress, uint256 ArrayKey) constant returns(address value)
func (_EthoFSController *EthoFSControllerCaller) GetHostingContractAddress(opts *bind.CallOpts, UserAddress common.Address, ArrayKey *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthoFSController.contract.Call(opts, out, "GetHostingContractAddress", UserAddress, ArrayKey)
	return *ret0, err
}

// GetHostingContractAddress is a free data retrieval call binding the contract method 0x68e15933.
//
// Solidity: function GetHostingContractAddress(address UserAddress, uint256 ArrayKey) constant returns(address value)
func (_EthoFSController *EthoFSControllerSession) GetHostingContractAddress(UserAddress common.Address, ArrayKey *big.Int) (common.Address, error) {
	return _EthoFSController.Contract.GetHostingContractAddress(&_EthoFSController.CallOpts, UserAddress, ArrayKey)
}

// GetHostingContractAddress is a free data retrieval call binding the contract method 0x68e15933.
//
// Solidity: function GetHostingContractAddress(address UserAddress, uint256 ArrayKey) constant returns(address value)
func (_EthoFSController *EthoFSControllerCallerSession) GetHostingContractAddress(UserAddress common.Address, ArrayKey *big.Int) (common.Address, error) {
	return _EthoFSController.Contract.GetHostingContractAddress(&_EthoFSController.CallOpts, UserAddress, ArrayKey)
}

// GetHostingContractDeployedBlockHeight is a free data retrieval call binding the contract method 0x415dfefa.
//
// Solidity: function GetHostingContractDeployedBlockHeight(address HostingContractAddress) constant returns(uint256 value)
func (_EthoFSController *EthoFSControllerCaller) GetHostingContractDeployedBlockHeight(opts *bind.CallOpts, HostingContractAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthoFSController.contract.Call(opts, out, "GetHostingContractDeployedBlockHeight", HostingContractAddress)
	return *ret0, err
}

// GetHostingContractDeployedBlockHeight is a free data retrieval call binding the contract method 0x415dfefa.
//
// Solidity: function GetHostingContractDeployedBlockHeight(address HostingContractAddress) constant returns(uint256 value)
func (_EthoFSController *EthoFSControllerSession) GetHostingContractDeployedBlockHeight(HostingContractAddress common.Address) (*big.Int, error) {
	return _EthoFSController.Contract.GetHostingContractDeployedBlockHeight(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetHostingContractDeployedBlockHeight is a free data retrieval call binding the contract method 0x415dfefa.
//
// Solidity: function GetHostingContractDeployedBlockHeight(address HostingContractAddress) constant returns(uint256 value)
func (_EthoFSController *EthoFSControllerCallerSession) GetHostingContractDeployedBlockHeight(HostingContractAddress common.Address) (*big.Int, error) {
	return _EthoFSController.Contract.GetHostingContractDeployedBlockHeight(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetHostingContractExpirationBlockHeight is a free data retrieval call binding the contract method 0x403169ac.
//
// Solidity: function GetHostingContractExpirationBlockHeight(address HostingContractAddress) constant returns(uint256 value)
func (_EthoFSController *EthoFSControllerCaller) GetHostingContractExpirationBlockHeight(opts *bind.CallOpts, HostingContractAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthoFSController.contract.Call(opts, out, "GetHostingContractExpirationBlockHeight", HostingContractAddress)
	return *ret0, err
}

// GetHostingContractExpirationBlockHeight is a free data retrieval call binding the contract method 0x403169ac.
//
// Solidity: function GetHostingContractExpirationBlockHeight(address HostingContractAddress) constant returns(uint256 value)
func (_EthoFSController *EthoFSControllerSession) GetHostingContractExpirationBlockHeight(HostingContractAddress common.Address) (*big.Int, error) {
	return _EthoFSController.Contract.GetHostingContractExpirationBlockHeight(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetHostingContractExpirationBlockHeight is a free data retrieval call binding the contract method 0x403169ac.
//
// Solidity: function GetHostingContractExpirationBlockHeight(address HostingContractAddress) constant returns(uint256 value)
func (_EthoFSController *EthoFSControllerCallerSession) GetHostingContractExpirationBlockHeight(HostingContractAddress common.Address) (*big.Int, error) {
	return _EthoFSController.Contract.GetHostingContractExpirationBlockHeight(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetHostingContractName is a free data retrieval call binding the contract method 0x8e097614.
//
// Solidity: function GetHostingContractName(address HostingContractAddress) constant returns(string value)
func (_EthoFSController *EthoFSControllerCaller) GetHostingContractName(opts *bind.CallOpts, HostingContractAddress common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EthoFSController.contract.Call(opts, out, "GetHostingContractName", HostingContractAddress)
	return *ret0, err
}

// GetHostingContractName is a free data retrieval call binding the contract method 0x8e097614.
//
// Solidity: function GetHostingContractName(address HostingContractAddress) constant returns(string value)
func (_EthoFSController *EthoFSControllerSession) GetHostingContractName(HostingContractAddress common.Address) (string, error) {
	return _EthoFSController.Contract.GetHostingContractName(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetHostingContractName is a free data retrieval call binding the contract method 0x8e097614.
//
// Solidity: function GetHostingContractName(address HostingContractAddress) constant returns(string value)
func (_EthoFSController *EthoFSControllerCallerSession) GetHostingContractName(HostingContractAddress common.Address) (string, error) {
	return _EthoFSController.Contract.GetHostingContractName(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetHostingContractStorageUsed is a free data retrieval call binding the contract method 0x66b7004c.
//
// Solidity: function GetHostingContractStorageUsed(address HostingContractAddress) constant returns(uint32 value)
func (_EthoFSController *EthoFSControllerCaller) GetHostingContractStorageUsed(opts *bind.CallOpts, HostingContractAddress common.Address) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _EthoFSController.contract.Call(opts, out, "GetHostingContractStorageUsed", HostingContractAddress)
	return *ret0, err
}

// GetHostingContractStorageUsed is a free data retrieval call binding the contract method 0x66b7004c.
//
// Solidity: function GetHostingContractStorageUsed(address HostingContractAddress) constant returns(uint32 value)
func (_EthoFSController *EthoFSControllerSession) GetHostingContractStorageUsed(HostingContractAddress common.Address) (uint32, error) {
	return _EthoFSController.Contract.GetHostingContractStorageUsed(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetHostingContractStorageUsed is a free data retrieval call binding the contract method 0x66b7004c.
//
// Solidity: function GetHostingContractStorageUsed(address HostingContractAddress) constant returns(uint32 value)
func (_EthoFSController *EthoFSControllerCallerSession) GetHostingContractStorageUsed(HostingContractAddress common.Address) (uint32, error) {
	return _EthoFSController.Contract.GetHostingContractStorageUsed(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetMainContentHash is a free data retrieval call binding the contract method 0xe6dc7817.
//
// Solidity: function GetMainContentHash(address HostingContractAddress) constant returns(string MainContentHash)
func (_EthoFSController *EthoFSControllerCaller) GetMainContentHash(opts *bind.CallOpts, HostingContractAddress common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EthoFSController.contract.Call(opts, out, "GetMainContentHash", HostingContractAddress)
	return *ret0, err
}

// GetMainContentHash is a free data retrieval call binding the contract method 0xe6dc7817.
//
// Solidity: function GetMainContentHash(address HostingContractAddress) constant returns(string MainContentHash)
func (_EthoFSController *EthoFSControllerSession) GetMainContentHash(HostingContractAddress common.Address) (string, error) {
	return _EthoFSController.Contract.GetMainContentHash(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetMainContentHash is a free data retrieval call binding the contract method 0xe6dc7817.
//
// Solidity: function GetMainContentHash(address HostingContractAddress) constant returns(string MainContentHash)
func (_EthoFSController *EthoFSControllerCallerSession) GetMainContentHash(HostingContractAddress common.Address) (string, error) {
	return _EthoFSController.Contract.GetMainContentHash(&_EthoFSController.CallOpts, HostingContractAddress)
}

// GetUserAccountActiveContractCount is a free data retrieval call binding the contract method 0xd4f0d929.
//
// Solidity: function GetUserAccountActiveContractCount(address UserAddress) constant returns(uint32 value)
func (_EthoFSController *EthoFSControllerCaller) GetUserAccountActiveContractCount(opts *bind.CallOpts, UserAddress common.Address) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _EthoFSController.contract.Call(opts, out, "GetUserAccountActiveContractCount", UserAddress)
	return *ret0, err
}

// GetUserAccountActiveContractCount is a free data retrieval call binding the contract method 0xd4f0d929.
//
// Solidity: function GetUserAccountActiveContractCount(address UserAddress) constant returns(uint32 value)
func (_EthoFSController *EthoFSControllerSession) GetUserAccountActiveContractCount(UserAddress common.Address) (uint32, error) {
	return _EthoFSController.Contract.GetUserAccountActiveContractCount(&_EthoFSController.CallOpts, UserAddress)
}

// GetUserAccountActiveContractCount is a free data retrieval call binding the contract method 0xd4f0d929.
//
// Solidity: function GetUserAccountActiveContractCount(address UserAddress) constant returns(uint32 value)
func (_EthoFSController *EthoFSControllerCallerSession) GetUserAccountActiveContractCount(UserAddress common.Address) (uint32, error) {
	return _EthoFSController.Contract.GetUserAccountActiveContractCount(&_EthoFSController.CallOpts, UserAddress)
}

// GetUserAccountName is a free data retrieval call binding the contract method 0xe42f0027.
//
// Solidity: function GetUserAccountName(address UserAddress) constant returns(string value)
func (_EthoFSController *EthoFSControllerCaller) GetUserAccountName(opts *bind.CallOpts, UserAddress common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EthoFSController.contract.Call(opts, out, "GetUserAccountName", UserAddress)
	return *ret0, err
}

// GetUserAccountName is a free data retrieval call binding the contract method 0xe42f0027.
//
// Solidity: function GetUserAccountName(address UserAddress) constant returns(string value)
func (_EthoFSController *EthoFSControllerSession) GetUserAccountName(UserAddress common.Address) (string, error) {
	return _EthoFSController.Contract.GetUserAccountName(&_EthoFSController.CallOpts, UserAddress)
}

// GetUserAccountName is a free data retrieval call binding the contract method 0xe42f0027.
//
// Solidity: function GetUserAccountName(address UserAddress) constant returns(string value)
func (_EthoFSController *EthoFSControllerCallerSession) GetUserAccountName(UserAddress common.Address) (string, error) {
	return _EthoFSController.Contract.GetUserAccountName(&_EthoFSController.CallOpts, UserAddress)
}

// GetUserAccountTotalContractCount is a free data retrieval call binding the contract method 0x9b0a9add.
//
// Solidity: function GetUserAccountTotalContractCount(address UserAddress) constant returns(uint32 value)
func (_EthoFSController *EthoFSControllerCaller) GetUserAccountTotalContractCount(opts *bind.CallOpts, UserAddress common.Address) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _EthoFSController.contract.Call(opts, out, "GetUserAccountTotalContractCount", UserAddress)
	return *ret0, err
}

// GetUserAccountTotalContractCount is a free data retrieval call binding the contract method 0x9b0a9add.
//
// Solidity: function GetUserAccountTotalContractCount(address UserAddress) constant returns(uint32 value)
func (_EthoFSController *EthoFSControllerSession) GetUserAccountTotalContractCount(UserAddress common.Address) (uint32, error) {
	return _EthoFSController.Contract.GetUserAccountTotalContractCount(&_EthoFSController.CallOpts, UserAddress)
}

// GetUserAccountTotalContractCount is a free data retrieval call binding the contract method 0x9b0a9add.
//
// Solidity: function GetUserAccountTotalContractCount(address UserAddress) constant returns(uint32 value)
func (_EthoFSController *EthoFSControllerCallerSession) GetUserAccountTotalContractCount(UserAddress common.Address) (uint32, error) {
	return _EthoFSController.Contract.GetUserAccountTotalContractCount(&_EthoFSController.CallOpts, UserAddress)
}

// HostingCost is a free data retrieval call binding the contract method 0x7201f738.
//
// Solidity: function HostingCost() constant returns(uint256)
func (_EthoFSController *EthoFSControllerCaller) HostingCost(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthoFSController.contract.Call(opts, out, "HostingCost")
	return *ret0, err
}

// HostingCost is a free data retrieval call binding the contract method 0x7201f738.
//
// Solidity: function HostingCost() constant returns(uint256)
func (_EthoFSController *EthoFSControllerSession) HostingCost() (*big.Int, error) {
	return _EthoFSController.Contract.HostingCost(&_EthoFSController.CallOpts)
}

// HostingCost is a free data retrieval call binding the contract method 0x7201f738.
//
// Solidity: function HostingCost() constant returns(uint256)
func (_EthoFSController *EthoFSControllerCallerSession) HostingCost() (*big.Int, error) {
	return _EthoFSController.Contract.HostingCost(&_EthoFSController.CallOpts)
}

// AddNewContract is a paid mutator transaction binding the contract method 0xccb726b1.
//
// Solidity: function AddNewContract(string MainContentHash, string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize, uint32 pinSize, string ContentHashString, string ContentPathString) returns()
func (_EthoFSController *EthoFSControllerTransactor) AddNewContract(opts *bind.TransactOpts, MainContentHash string, HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32, pinSize uint32, ContentHashString string, ContentPathString string) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "AddNewContract", MainContentHash, HostingContractName, HostingContractDuration, TotalContractSize, pinSize, ContentHashString, ContentPathString)
}

// AddNewContract is a paid mutator transaction binding the contract method 0xccb726b1.
//
// Solidity: function AddNewContract(string MainContentHash, string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize, uint32 pinSize, string ContentHashString, string ContentPathString) returns()
func (_EthoFSController *EthoFSControllerSession) AddNewContract(MainContentHash string, HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32, pinSize uint32, ContentHashString string, ContentPathString string) (*types.Transaction, error) {
	return _EthoFSController.Contract.AddNewContract(&_EthoFSController.TransactOpts, MainContentHash, HostingContractName, HostingContractDuration, TotalContractSize, pinSize, ContentHashString, ContentPathString)
}

// AddNewContract is a paid mutator transaction binding the contract method 0xccb726b1.
//
// Solidity: function AddNewContract(string MainContentHash, string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize, uint32 pinSize, string ContentHashString, string ContentPathString) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) AddNewContract(MainContentHash string, HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32, pinSize uint32, ContentHashString string, ContentPathString string) (*types.Transaction, error) {
	return _EthoFSController.Contract.AddNewContract(&_EthoFSController.TransactOpts, MainContentHash, HostingContractName, HostingContractDuration, TotalContractSize, pinSize, ContentHashString, ContentPathString)
}

// AddNewUserOwner is a paid mutator transaction binding the contract method 0x4a6b2fa3.
//
// Solidity: function AddNewUserOwner(address UserAddress, string AccountName) returns()
func (_EthoFSController *EthoFSControllerTransactor) AddNewUserOwner(opts *bind.TransactOpts, UserAddress common.Address, AccountName string) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "AddNewUserOwner", UserAddress, AccountName)
}

// AddNewUserOwner is a paid mutator transaction binding the contract method 0x4a6b2fa3.
//
// Solidity: function AddNewUserOwner(address UserAddress, string AccountName) returns()
func (_EthoFSController *EthoFSControllerSession) AddNewUserOwner(UserAddress common.Address, AccountName string) (*types.Transaction, error) {
	return _EthoFSController.Contract.AddNewUserOwner(&_EthoFSController.TransactOpts, UserAddress, AccountName)
}

// AddNewUserOwner is a paid mutator transaction binding the contract method 0x4a6b2fa3.
//
// Solidity: function AddNewUserOwner(address UserAddress, string AccountName) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) AddNewUserOwner(UserAddress common.Address, AccountName string) (*types.Transaction, error) {
	return _EthoFSController.Contract.AddNewUserOwner(&_EthoFSController.TransactOpts, UserAddress, AccountName)
}

// AddNewUserPublic is a paid mutator transaction binding the contract method 0xb18759de.
//
// Solidity: function AddNewUserPublic(string AccountName) returns()
func (_EthoFSController *EthoFSControllerTransactor) AddNewUserPublic(opts *bind.TransactOpts, AccountName string) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "AddNewUserPublic", AccountName)
}

// AddNewUserPublic is a paid mutator transaction binding the contract method 0xb18759de.
//
// Solidity: function AddNewUserPublic(string AccountName) returns()
func (_EthoFSController *EthoFSControllerSession) AddNewUserPublic(AccountName string) (*types.Transaction, error) {
	return _EthoFSController.Contract.AddNewUserPublic(&_EthoFSController.TransactOpts, AccountName)
}

// AddNewUserPublic is a paid mutator transaction binding the contract method 0xb18759de.
//
// Solidity: function AddNewUserPublic(string AccountName) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) AddNewUserPublic(AccountName string) (*types.Transaction, error) {
	return _EthoFSController.Contract.AddNewUserPublic(&_EthoFSController.TransactOpts, AccountName)
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

// RemoveUserOwner is a paid mutator transaction binding the contract method 0x0bbeb35b.
//
// Solidity: function RemoveUserOwner(address UserAddress) returns()
func (_EthoFSController *EthoFSControllerTransactor) RemoveUserOwner(opts *bind.TransactOpts, UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "RemoveUserOwner", UserAddress)
}

// RemoveUserOwner is a paid mutator transaction binding the contract method 0x0bbeb35b.
//
// Solidity: function RemoveUserOwner(address UserAddress) returns()
func (_EthoFSController *EthoFSControllerSession) RemoveUserOwner(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.RemoveUserOwner(&_EthoFSController.TransactOpts, UserAddress)
}

// RemoveUserOwner is a paid mutator transaction binding the contract method 0x0bbeb35b.
//
// Solidity: function RemoveUserOwner(address UserAddress) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) RemoveUserOwner(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.RemoveUserOwner(&_EthoFSController.TransactOpts, UserAddress)
}

// RemoveUserPublic is a paid mutator transaction binding the contract method 0xd1923763.
//
// Solidity: function RemoveUserPublic() returns()
func (_EthoFSController *EthoFSControllerTransactor) RemoveUserPublic(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "RemoveUserPublic")
}

// RemoveUserPublic is a paid mutator transaction binding the contract method 0xd1923763.
//
// Solidity: function RemoveUserPublic() returns()
func (_EthoFSController *EthoFSControllerSession) RemoveUserPublic() (*types.Transaction, error) {
	return _EthoFSController.Contract.RemoveUserPublic(&_EthoFSController.TransactOpts)
}

// RemoveUserPublic is a paid mutator transaction binding the contract method 0xd1923763.
//
// Solidity: function RemoveUserPublic() returns()
func (_EthoFSController *EthoFSControllerTransactorSession) RemoveUserPublic() (*types.Transaction, error) {
	return _EthoFSController.Contract.RemoveUserPublic(&_EthoFSController.TransactOpts)
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

// SetAccountCollectionAddress is a paid mutator transaction binding the contract method 0xc8ac5fe8.
//
// Solidity: function SetAccountCollectionAddress(address set) returns()
func (_EthoFSController *EthoFSControllerTransactor) SetAccountCollectionAddress(opts *bind.TransactOpts, set common.Address) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "SetAccountCollectionAddress", set)
}

// SetAccountCollectionAddress is a paid mutator transaction binding the contract method 0xc8ac5fe8.
//
// Solidity: function SetAccountCollectionAddress(address set) returns()
func (_EthoFSController *EthoFSControllerSession) SetAccountCollectionAddress(set common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.SetAccountCollectionAddress(&_EthoFSController.TransactOpts, set)
}

// SetAccountCollectionAddress is a paid mutator transaction binding the contract method 0xc8ac5fe8.
//
// Solidity: function SetAccountCollectionAddress(address set) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) SetAccountCollectionAddress(set common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.SetAccountCollectionAddress(&_EthoFSController.TransactOpts, set)
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

// SetEthoFSHostingContractsAddress is a paid mutator transaction binding the contract method 0x978a627a.
//
// Solidity: function SetEthoFSHostingContractsAddress(address ethoFSHostingContractsAddress) returns()
func (_EthoFSController *EthoFSControllerTransactor) SetEthoFSHostingContractsAddress(opts *bind.TransactOpts, ethoFSHostingContractsAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.contract.Transact(opts, "SetEthoFSHostingContractsAddress", ethoFSHostingContractsAddress)
}

// SetEthoFSHostingContractsAddress is a paid mutator transaction binding the contract method 0x978a627a.
//
// Solidity: function SetEthoFSHostingContractsAddress(address ethoFSHostingContractsAddress) returns()
func (_EthoFSController *EthoFSControllerSession) SetEthoFSHostingContractsAddress(ethoFSHostingContractsAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.SetEthoFSHostingContractsAddress(&_EthoFSController.TransactOpts, ethoFSHostingContractsAddress)
}

// SetEthoFSHostingContractsAddress is a paid mutator transaction binding the contract method 0x978a627a.
//
// Solidity: function SetEthoFSHostingContractsAddress(address ethoFSHostingContractsAddress) returns()
func (_EthoFSController *EthoFSControllerTransactorSession) SetEthoFSHostingContractsAddress(ethoFSHostingContractsAddress common.Address) (*types.Transaction, error) {
	return _EthoFSController.Contract.SetEthoFSHostingContractsAddress(&_EthoFSController.TransactOpts, ethoFSHostingContractsAddress)
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
const EthoFSDashboardABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"}],\"name\":\"RemoveUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"}],\"name\":\"GetUserAccountAddress\",\"outputs\":[{\"name\":\"value\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"}],\"name\":\"CheckAccountExistence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"},{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"RemoveHostingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newContractAddress\",\"type\":\"address\"},{\"name\":\"UserAddress\",\"type\":\"address\"},{\"name\":\"HostingContractName\",\"type\":\"string\"},{\"name\":\"HostingContractDuration\",\"type\":\"uint32\"},{\"name\":\"TotalContractSize\",\"type\":\"uint32\"}],\"name\":\"AddHostingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"},{\"name\":\"ArrayKey\",\"type\":\"uint256\"}],\"name\":\"GetHostingContractAddress\",\"outputs\":[{\"name\":\"value\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"}],\"name\":\"GetUserAccountTotalContractCount\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"},{\"name\":\"AccountName\",\"type\":\"string\"}],\"name\":\"AddNewUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"set\",\"type\":\"address\"}],\"name\":\"SetAccountCollectionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"}],\"name\":\"GetUserAccountActiveContractCount\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"UserAddress\",\"type\":\"address\"}],\"name\":\"GetUserAccountName\",\"outputs\":[{\"name\":\"value\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// AddHostingContract is a paid mutator transaction binding the contract method 0x57199193.
//
// Solidity: function AddHostingContract(address newContractAddress, address UserAddress, string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactor) AddHostingContract(opts *bind.TransactOpts, newContractAddress common.Address, UserAddress common.Address, HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "AddHostingContract", newContractAddress, UserAddress, HostingContractName, HostingContractDuration, TotalContractSize)
}

// AddHostingContract is a paid mutator transaction binding the contract method 0x57199193.
//
// Solidity: function AddHostingContract(address newContractAddress, address UserAddress, string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize) returns()
func (_EthoFSDashboard *EthoFSDashboardSession) AddHostingContract(newContractAddress common.Address, UserAddress common.Address, HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.AddHostingContract(&_EthoFSDashboard.TransactOpts, newContractAddress, UserAddress, HostingContractName, HostingContractDuration, TotalContractSize)
}

// AddHostingContract is a paid mutator transaction binding the contract method 0x57199193.
//
// Solidity: function AddHostingContract(address newContractAddress, address UserAddress, string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) AddHostingContract(newContractAddress common.Address, UserAddress common.Address, HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.AddHostingContract(&_EthoFSDashboard.TransactOpts, newContractAddress, UserAddress, HostingContractName, HostingContractDuration, TotalContractSize)
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

// CheckAccountExistence is a paid mutator transaction binding the contract method 0x2c3027a7.
//
// Solidity: function CheckAccountExistence(address UserAddress) returns(bool)
func (_EthoFSDashboard *EthoFSDashboardTransactor) CheckAccountExistence(opts *bind.TransactOpts, UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "CheckAccountExistence", UserAddress)
}

// CheckAccountExistence is a paid mutator transaction binding the contract method 0x2c3027a7.
//
// Solidity: function CheckAccountExistence(address UserAddress) returns(bool)
func (_EthoFSDashboard *EthoFSDashboardSession) CheckAccountExistence(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.CheckAccountExistence(&_EthoFSDashboard.TransactOpts, UserAddress)
}

// CheckAccountExistence is a paid mutator transaction binding the contract method 0x2c3027a7.
//
// Solidity: function CheckAccountExistence(address UserAddress) returns(bool)
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) CheckAccountExistence(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.CheckAccountExistence(&_EthoFSDashboard.TransactOpts, UserAddress)
}

// GetHostingContractAddress is a paid mutator transaction binding the contract method 0x68e15933.
//
// Solidity: function GetHostingContractAddress(address UserAddress, uint256 ArrayKey) returns(address value)
func (_EthoFSDashboard *EthoFSDashboardTransactor) GetHostingContractAddress(opts *bind.TransactOpts, UserAddress common.Address, ArrayKey *big.Int) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "GetHostingContractAddress", UserAddress, ArrayKey)
}

// GetHostingContractAddress is a paid mutator transaction binding the contract method 0x68e15933.
//
// Solidity: function GetHostingContractAddress(address UserAddress, uint256 ArrayKey) returns(address value)
func (_EthoFSDashboard *EthoFSDashboardSession) GetHostingContractAddress(UserAddress common.Address, ArrayKey *big.Int) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.GetHostingContractAddress(&_EthoFSDashboard.TransactOpts, UserAddress, ArrayKey)
}

// GetHostingContractAddress is a paid mutator transaction binding the contract method 0x68e15933.
//
// Solidity: function GetHostingContractAddress(address UserAddress, uint256 ArrayKey) returns(address value)
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) GetHostingContractAddress(UserAddress common.Address, ArrayKey *big.Int) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.GetHostingContractAddress(&_EthoFSDashboard.TransactOpts, UserAddress, ArrayKey)
}

// GetUserAccountActiveContractCount is a paid mutator transaction binding the contract method 0xd4f0d929.
//
// Solidity: function GetUserAccountActiveContractCount(address UserAddress) returns(uint32 value)
func (_EthoFSDashboard *EthoFSDashboardTransactor) GetUserAccountActiveContractCount(opts *bind.TransactOpts, UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "GetUserAccountActiveContractCount", UserAddress)
}

// GetUserAccountActiveContractCount is a paid mutator transaction binding the contract method 0xd4f0d929.
//
// Solidity: function GetUserAccountActiveContractCount(address UserAddress) returns(uint32 value)
func (_EthoFSDashboard *EthoFSDashboardSession) GetUserAccountActiveContractCount(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.GetUserAccountActiveContractCount(&_EthoFSDashboard.TransactOpts, UserAddress)
}

// GetUserAccountActiveContractCount is a paid mutator transaction binding the contract method 0xd4f0d929.
//
// Solidity: function GetUserAccountActiveContractCount(address UserAddress) returns(uint32 value)
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) GetUserAccountActiveContractCount(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.GetUserAccountActiveContractCount(&_EthoFSDashboard.TransactOpts, UserAddress)
}

// GetUserAccountAddress is a paid mutator transaction binding the contract method 0x2c1ca1ac.
//
// Solidity: function GetUserAccountAddress(address UserAddress) returns(address value)
func (_EthoFSDashboard *EthoFSDashboardTransactor) GetUserAccountAddress(opts *bind.TransactOpts, UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "GetUserAccountAddress", UserAddress)
}

// GetUserAccountAddress is a paid mutator transaction binding the contract method 0x2c1ca1ac.
//
// Solidity: function GetUserAccountAddress(address UserAddress) returns(address value)
func (_EthoFSDashboard *EthoFSDashboardSession) GetUserAccountAddress(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.GetUserAccountAddress(&_EthoFSDashboard.TransactOpts, UserAddress)
}

// GetUserAccountAddress is a paid mutator transaction binding the contract method 0x2c1ca1ac.
//
// Solidity: function GetUserAccountAddress(address UserAddress) returns(address value)
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) GetUserAccountAddress(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.GetUserAccountAddress(&_EthoFSDashboard.TransactOpts, UserAddress)
}

// GetUserAccountName is a paid mutator transaction binding the contract method 0xe42f0027.
//
// Solidity: function GetUserAccountName(address UserAddress) returns(string value)
func (_EthoFSDashboard *EthoFSDashboardTransactor) GetUserAccountName(opts *bind.TransactOpts, UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "GetUserAccountName", UserAddress)
}

// GetUserAccountName is a paid mutator transaction binding the contract method 0xe42f0027.
//
// Solidity: function GetUserAccountName(address UserAddress) returns(string value)
func (_EthoFSDashboard *EthoFSDashboardSession) GetUserAccountName(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.GetUserAccountName(&_EthoFSDashboard.TransactOpts, UserAddress)
}

// GetUserAccountName is a paid mutator transaction binding the contract method 0xe42f0027.
//
// Solidity: function GetUserAccountName(address UserAddress) returns(string value)
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) GetUserAccountName(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.GetUserAccountName(&_EthoFSDashboard.TransactOpts, UserAddress)
}

// GetUserAccountTotalContractCount is a paid mutator transaction binding the contract method 0x9b0a9add.
//
// Solidity: function GetUserAccountTotalContractCount(address UserAddress) returns(uint32 value)
func (_EthoFSDashboard *EthoFSDashboardTransactor) GetUserAccountTotalContractCount(opts *bind.TransactOpts, UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "GetUserAccountTotalContractCount", UserAddress)
}

// GetUserAccountTotalContractCount is a paid mutator transaction binding the contract method 0x9b0a9add.
//
// Solidity: function GetUserAccountTotalContractCount(address UserAddress) returns(uint32 value)
func (_EthoFSDashboard *EthoFSDashboardSession) GetUserAccountTotalContractCount(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.GetUserAccountTotalContractCount(&_EthoFSDashboard.TransactOpts, UserAddress)
}

// GetUserAccountTotalContractCount is a paid mutator transaction binding the contract method 0x9b0a9add.
//
// Solidity: function GetUserAccountTotalContractCount(address UserAddress) returns(uint32 value)
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) GetUserAccountTotalContractCount(UserAddress common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.GetUserAccountTotalContractCount(&_EthoFSDashboard.TransactOpts, UserAddress)
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

// SetAccountCollectionAddress is a paid mutator transaction binding the contract method 0xc8ac5fe8.
//
// Solidity: function SetAccountCollectionAddress(address set) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactor) SetAccountCollectionAddress(opts *bind.TransactOpts, set common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.contract.Transact(opts, "SetAccountCollectionAddress", set)
}

// SetAccountCollectionAddress is a paid mutator transaction binding the contract method 0xc8ac5fe8.
//
// Solidity: function SetAccountCollectionAddress(address set) returns()
func (_EthoFSDashboard *EthoFSDashboardSession) SetAccountCollectionAddress(set common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.SetAccountCollectionAddress(&_EthoFSDashboard.TransactOpts, set)
}

// SetAccountCollectionAddress is a paid mutator transaction binding the contract method 0xc8ac5fe8.
//
// Solidity: function SetAccountCollectionAddress(address set) returns()
func (_EthoFSDashboard *EthoFSDashboardTransactorSession) SetAccountCollectionAddress(set common.Address) (*types.Transaction, error) {
	return _EthoFSDashboard.Contract.SetAccountCollectionAddress(&_EthoFSDashboard.TransactOpts, set)
}

// EthoFSHostingContractsABI is the input ABI used to generate the binding from.
const EthoFSHostingContractsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"set\",\"type\":\"uint256\"}],\"name\":\"SetHostingContractCost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ScrubContractList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetHostingContractExpirationBlockHeight\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetHostingContractDeployedBlockHeight\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"MainContentHash\",\"type\":\"string\"},{\"name\":\"HostingContractName\",\"type\":\"string\"},{\"name\":\"HostingContractDuration\",\"type\":\"uint32\"},{\"name\":\"TotalContractSize\",\"type\":\"uint32\"},{\"name\":\"ContentHashString\",\"type\":\"string\"},{\"name\":\"ContentPathString\",\"type\":\"string\"}],\"name\":\"AddHostingContract\",\"outputs\":[{\"name\":\"value\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetHostingContractStorageUsed\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"CustomerAddress\",\"type\":\"address\"},{\"name\":\"HostingContractAddress\",\"type\":\"address\"},{\"name\":\"AccountCollectionAddress\",\"type\":\"address\"}],\"name\":\"RemoveHostingContract\",\"outputs\":[{\"name\":\"value\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetHostingContractName\",\"outputs\":[{\"name\":\"value\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetContentHashString\",\"outputs\":[{\"name\":\"ContentHashString\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"AccountCollectionAddress\",\"type\":\"address\"}],\"name\":\"SetAccountCollectionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetContentPathString\",\"outputs\":[{\"name\":\"ContentPathString\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"},{\"name\":\"HostingContractExtensionDuration\",\"type\":\"uint32\"}],\"name\":\"ExtendHostingContract\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"HostingContractAddress\",\"type\":\"address\"}],\"name\":\"GetMainContentHash\",\"outputs\":[{\"name\":\"MainContentHash\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthoFSHostingContractsBin is the compiled bytecode used for deploying new contracts.
const EthoFSHostingContractsBin = `0x`

// DeployEthoFSHostingContracts deploys a new Ethereum contract, binding an instance of EthoFSHostingContracts to it.
func DeployEthoFSHostingContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthoFSHostingContracts, error) {
	parsed, err := abi.JSON(strings.NewReader(EthoFSHostingContractsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthoFSHostingContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthoFSHostingContracts{EthoFSHostingContractsCaller: EthoFSHostingContractsCaller{contract: contract}, EthoFSHostingContractsTransactor: EthoFSHostingContractsTransactor{contract: contract}, EthoFSHostingContractsFilterer: EthoFSHostingContractsFilterer{contract: contract}}, nil
}

// EthoFSHostingContracts is an auto generated Go binding around an Ethereum contract.
type EthoFSHostingContracts struct {
	EthoFSHostingContractsCaller     // Read-only binding to the contract
	EthoFSHostingContractsTransactor // Write-only binding to the contract
	EthoFSHostingContractsFilterer   // Log filterer for contract events
}

// EthoFSHostingContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthoFSHostingContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthoFSHostingContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthoFSHostingContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthoFSHostingContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthoFSHostingContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthoFSHostingContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthoFSHostingContractsSession struct {
	Contract     *EthoFSHostingContracts // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EthoFSHostingContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthoFSHostingContractsCallerSession struct {
	Contract *EthoFSHostingContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// EthoFSHostingContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthoFSHostingContractsTransactorSession struct {
	Contract     *EthoFSHostingContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// EthoFSHostingContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthoFSHostingContractsRaw struct {
	Contract *EthoFSHostingContracts // Generic contract binding to access the raw methods on
}

// EthoFSHostingContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthoFSHostingContractsCallerRaw struct {
	Contract *EthoFSHostingContractsCaller // Generic read-only contract binding to access the raw methods on
}

// EthoFSHostingContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthoFSHostingContractsTransactorRaw struct {
	Contract *EthoFSHostingContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthoFSHostingContracts creates a new instance of EthoFSHostingContracts, bound to a specific deployed contract.
func NewEthoFSHostingContracts(address common.Address, backend bind.ContractBackend) (*EthoFSHostingContracts, error) {
	contract, err := bindEthoFSHostingContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthoFSHostingContracts{EthoFSHostingContractsCaller: EthoFSHostingContractsCaller{contract: contract}, EthoFSHostingContractsTransactor: EthoFSHostingContractsTransactor{contract: contract}, EthoFSHostingContractsFilterer: EthoFSHostingContractsFilterer{contract: contract}}, nil
}

// NewEthoFSHostingContractsCaller creates a new read-only instance of EthoFSHostingContracts, bound to a specific deployed contract.
func NewEthoFSHostingContractsCaller(address common.Address, caller bind.ContractCaller) (*EthoFSHostingContractsCaller, error) {
	contract, err := bindEthoFSHostingContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthoFSHostingContractsCaller{contract: contract}, nil
}

// NewEthoFSHostingContractsTransactor creates a new write-only instance of EthoFSHostingContracts, bound to a specific deployed contract.
func NewEthoFSHostingContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*EthoFSHostingContractsTransactor, error) {
	contract, err := bindEthoFSHostingContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthoFSHostingContractsTransactor{contract: contract}, nil
}

// NewEthoFSHostingContractsFilterer creates a new log filterer instance of EthoFSHostingContracts, bound to a specific deployed contract.
func NewEthoFSHostingContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*EthoFSHostingContractsFilterer, error) {
	contract, err := bindEthoFSHostingContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthoFSHostingContractsFilterer{contract: contract}, nil
}

// bindEthoFSHostingContracts binds a generic wrapper to an already deployed contract.
func bindEthoFSHostingContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthoFSHostingContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthoFSHostingContracts *EthoFSHostingContractsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthoFSHostingContracts.Contract.EthoFSHostingContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthoFSHostingContracts *EthoFSHostingContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.EthoFSHostingContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthoFSHostingContracts *EthoFSHostingContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.EthoFSHostingContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthoFSHostingContracts *EthoFSHostingContractsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthoFSHostingContracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.contract.Transact(opts, method, params...)
}

// AddHostingContract is a paid mutator transaction binding the contract method 0x4b8f5af6.
//
// Solidity: function AddHostingContract(string MainContentHash, string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize, string ContentHashString, string ContentPathString) returns(address value)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactor) AddHostingContract(opts *bind.TransactOpts, MainContentHash string, HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32, ContentHashString string, ContentPathString string) (*types.Transaction, error) {
	return _EthoFSHostingContracts.contract.Transact(opts, "AddHostingContract", MainContentHash, HostingContractName, HostingContractDuration, TotalContractSize, ContentHashString, ContentPathString)
}

// AddHostingContract is a paid mutator transaction binding the contract method 0x4b8f5af6.
//
// Solidity: function AddHostingContract(string MainContentHash, string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize, string ContentHashString, string ContentPathString) returns(address value)
func (_EthoFSHostingContracts *EthoFSHostingContractsSession) AddHostingContract(MainContentHash string, HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32, ContentHashString string, ContentPathString string) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.AddHostingContract(&_EthoFSHostingContracts.TransactOpts, MainContentHash, HostingContractName, HostingContractDuration, TotalContractSize, ContentHashString, ContentPathString)
}

// AddHostingContract is a paid mutator transaction binding the contract method 0x4b8f5af6.
//
// Solidity: function AddHostingContract(string MainContentHash, string HostingContractName, uint32 HostingContractDuration, uint32 TotalContractSize, string ContentHashString, string ContentPathString) returns(address value)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorSession) AddHostingContract(MainContentHash string, HostingContractName string, HostingContractDuration uint32, TotalContractSize uint32, ContentHashString string, ContentPathString string) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.AddHostingContract(&_EthoFSHostingContracts.TransactOpts, MainContentHash, HostingContractName, HostingContractDuration, TotalContractSize, ContentHashString, ContentPathString)
}

// ExtendHostingContract is a paid mutator transaction binding the contract method 0xde5a8ec5.
//
// Solidity: function ExtendHostingContract(address HostingContractAddress, uint32 HostingContractExtensionDuration) returns()
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactor) ExtendHostingContract(opts *bind.TransactOpts, HostingContractAddress common.Address, HostingContractExtensionDuration uint32) (*types.Transaction, error) {
	return _EthoFSHostingContracts.contract.Transact(opts, "ExtendHostingContract", HostingContractAddress, HostingContractExtensionDuration)
}

// ExtendHostingContract is a paid mutator transaction binding the contract method 0xde5a8ec5.
//
// Solidity: function ExtendHostingContract(address HostingContractAddress, uint32 HostingContractExtensionDuration) returns()
func (_EthoFSHostingContracts *EthoFSHostingContractsSession) ExtendHostingContract(HostingContractAddress common.Address, HostingContractExtensionDuration uint32) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.ExtendHostingContract(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress, HostingContractExtensionDuration)
}

// ExtendHostingContract is a paid mutator transaction binding the contract method 0xde5a8ec5.
//
// Solidity: function ExtendHostingContract(address HostingContractAddress, uint32 HostingContractExtensionDuration) returns()
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorSession) ExtendHostingContract(HostingContractAddress common.Address, HostingContractExtensionDuration uint32) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.ExtendHostingContract(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress, HostingContractExtensionDuration)
}

// GetContentHashString is a paid mutator transaction binding the contract method 0xaf7bbf5a.
//
// Solidity: function GetContentHashString(address HostingContractAddress) returns(string ContentHashString)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactor) GetContentHashString(opts *bind.TransactOpts, HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.contract.Transact(opts, "GetContentHashString", HostingContractAddress)
}

// GetContentHashString is a paid mutator transaction binding the contract method 0xaf7bbf5a.
//
// Solidity: function GetContentHashString(address HostingContractAddress) returns(string ContentHashString)
func (_EthoFSHostingContracts *EthoFSHostingContractsSession) GetContentHashString(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetContentHashString(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// GetContentHashString is a paid mutator transaction binding the contract method 0xaf7bbf5a.
//
// Solidity: function GetContentHashString(address HostingContractAddress) returns(string ContentHashString)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorSession) GetContentHashString(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetContentHashString(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// GetContentPathString is a paid mutator transaction binding the contract method 0xd0cd2513.
//
// Solidity: function GetContentPathString(address HostingContractAddress) returns(string ContentPathString)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactor) GetContentPathString(opts *bind.TransactOpts, HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.contract.Transact(opts, "GetContentPathString", HostingContractAddress)
}

// GetContentPathString is a paid mutator transaction binding the contract method 0xd0cd2513.
//
// Solidity: function GetContentPathString(address HostingContractAddress) returns(string ContentPathString)
func (_EthoFSHostingContracts *EthoFSHostingContractsSession) GetContentPathString(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetContentPathString(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// GetContentPathString is a paid mutator transaction binding the contract method 0xd0cd2513.
//
// Solidity: function GetContentPathString(address HostingContractAddress) returns(string ContentPathString)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorSession) GetContentPathString(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetContentPathString(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// GetHostingContractDeployedBlockHeight is a paid mutator transaction binding the contract method 0x415dfefa.
//
// Solidity: function GetHostingContractDeployedBlockHeight(address HostingContractAddress) returns(uint256 value)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactor) GetHostingContractDeployedBlockHeight(opts *bind.TransactOpts, HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.contract.Transact(opts, "GetHostingContractDeployedBlockHeight", HostingContractAddress)
}

// GetHostingContractDeployedBlockHeight is a paid mutator transaction binding the contract method 0x415dfefa.
//
// Solidity: function GetHostingContractDeployedBlockHeight(address HostingContractAddress) returns(uint256 value)
func (_EthoFSHostingContracts *EthoFSHostingContractsSession) GetHostingContractDeployedBlockHeight(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetHostingContractDeployedBlockHeight(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// GetHostingContractDeployedBlockHeight is a paid mutator transaction binding the contract method 0x415dfefa.
//
// Solidity: function GetHostingContractDeployedBlockHeight(address HostingContractAddress) returns(uint256 value)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorSession) GetHostingContractDeployedBlockHeight(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetHostingContractDeployedBlockHeight(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// GetHostingContractExpirationBlockHeight is a paid mutator transaction binding the contract method 0x403169ac.
//
// Solidity: function GetHostingContractExpirationBlockHeight(address HostingContractAddress) returns(uint256 value)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactor) GetHostingContractExpirationBlockHeight(opts *bind.TransactOpts, HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.contract.Transact(opts, "GetHostingContractExpirationBlockHeight", HostingContractAddress)
}

// GetHostingContractExpirationBlockHeight is a paid mutator transaction binding the contract method 0x403169ac.
//
// Solidity: function GetHostingContractExpirationBlockHeight(address HostingContractAddress) returns(uint256 value)
func (_EthoFSHostingContracts *EthoFSHostingContractsSession) GetHostingContractExpirationBlockHeight(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetHostingContractExpirationBlockHeight(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// GetHostingContractExpirationBlockHeight is a paid mutator transaction binding the contract method 0x403169ac.
//
// Solidity: function GetHostingContractExpirationBlockHeight(address HostingContractAddress) returns(uint256 value)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorSession) GetHostingContractExpirationBlockHeight(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetHostingContractExpirationBlockHeight(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// GetHostingContractName is a paid mutator transaction binding the contract method 0x8e097614.
//
// Solidity: function GetHostingContractName(address HostingContractAddress) returns(string value)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactor) GetHostingContractName(opts *bind.TransactOpts, HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.contract.Transact(opts, "GetHostingContractName", HostingContractAddress)
}

// GetHostingContractName is a paid mutator transaction binding the contract method 0x8e097614.
//
// Solidity: function GetHostingContractName(address HostingContractAddress) returns(string value)
func (_EthoFSHostingContracts *EthoFSHostingContractsSession) GetHostingContractName(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetHostingContractName(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// GetHostingContractName is a paid mutator transaction binding the contract method 0x8e097614.
//
// Solidity: function GetHostingContractName(address HostingContractAddress) returns(string value)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorSession) GetHostingContractName(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetHostingContractName(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// GetHostingContractStorageUsed is a paid mutator transaction binding the contract method 0x66b7004c.
//
// Solidity: function GetHostingContractStorageUsed(address HostingContractAddress) returns(uint32 value)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactor) GetHostingContractStorageUsed(opts *bind.TransactOpts, HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.contract.Transact(opts, "GetHostingContractStorageUsed", HostingContractAddress)
}

// GetHostingContractStorageUsed is a paid mutator transaction binding the contract method 0x66b7004c.
//
// Solidity: function GetHostingContractStorageUsed(address HostingContractAddress) returns(uint32 value)
func (_EthoFSHostingContracts *EthoFSHostingContractsSession) GetHostingContractStorageUsed(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetHostingContractStorageUsed(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// GetHostingContractStorageUsed is a paid mutator transaction binding the contract method 0x66b7004c.
//
// Solidity: function GetHostingContractStorageUsed(address HostingContractAddress) returns(uint32 value)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorSession) GetHostingContractStorageUsed(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetHostingContractStorageUsed(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// GetMainContentHash is a paid mutator transaction binding the contract method 0xe6dc7817.
//
// Solidity: function GetMainContentHash(address HostingContractAddress) returns(string MainContentHash)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactor) GetMainContentHash(opts *bind.TransactOpts, HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.contract.Transact(opts, "GetMainContentHash", HostingContractAddress)
}

// GetMainContentHash is a paid mutator transaction binding the contract method 0xe6dc7817.
//
// Solidity: function GetMainContentHash(address HostingContractAddress) returns(string MainContentHash)
func (_EthoFSHostingContracts *EthoFSHostingContractsSession) GetMainContentHash(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetMainContentHash(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// GetMainContentHash is a paid mutator transaction binding the contract method 0xe6dc7817.
//
// Solidity: function GetMainContentHash(address HostingContractAddress) returns(string MainContentHash)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorSession) GetMainContentHash(HostingContractAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.GetMainContentHash(&_EthoFSHostingContracts.TransactOpts, HostingContractAddress)
}

// RemoveHostingContract is a paid mutator transaction binding the contract method 0x782d8e06.
//
// Solidity: function RemoveHostingContract(address CustomerAddress, address HostingContractAddress, address AccountCollectionAddress) returns(bool value)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactor) RemoveHostingContract(opts *bind.TransactOpts, CustomerAddress common.Address, HostingContractAddress common.Address, AccountCollectionAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.contract.Transact(opts, "RemoveHostingContract", CustomerAddress, HostingContractAddress, AccountCollectionAddress)
}

// RemoveHostingContract is a paid mutator transaction binding the contract method 0x782d8e06.
//
// Solidity: function RemoveHostingContract(address CustomerAddress, address HostingContractAddress, address AccountCollectionAddress) returns(bool value)
func (_EthoFSHostingContracts *EthoFSHostingContractsSession) RemoveHostingContract(CustomerAddress common.Address, HostingContractAddress common.Address, AccountCollectionAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.RemoveHostingContract(&_EthoFSHostingContracts.TransactOpts, CustomerAddress, HostingContractAddress, AccountCollectionAddress)
}

// RemoveHostingContract is a paid mutator transaction binding the contract method 0x782d8e06.
//
// Solidity: function RemoveHostingContract(address CustomerAddress, address HostingContractAddress, address AccountCollectionAddress) returns(bool value)
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorSession) RemoveHostingContract(CustomerAddress common.Address, HostingContractAddress common.Address, AccountCollectionAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.RemoveHostingContract(&_EthoFSHostingContracts.TransactOpts, CustomerAddress, HostingContractAddress, AccountCollectionAddress)
}

// ScrubContractList is a paid mutator transaction binding the contract method 0x359a198b.
//
// Solidity: function ScrubContractList() returns()
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactor) ScrubContractList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthoFSHostingContracts.contract.Transact(opts, "ScrubContractList")
}

// ScrubContractList is a paid mutator transaction binding the contract method 0x359a198b.
//
// Solidity: function ScrubContractList() returns()
func (_EthoFSHostingContracts *EthoFSHostingContractsSession) ScrubContractList() (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.ScrubContractList(&_EthoFSHostingContracts.TransactOpts)
}

// ScrubContractList is a paid mutator transaction binding the contract method 0x359a198b.
//
// Solidity: function ScrubContractList() returns()
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorSession) ScrubContractList() (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.ScrubContractList(&_EthoFSHostingContracts.TransactOpts)
}

// SetAccountCollectionAddress is a paid mutator transaction binding the contract method 0xc8ac5fe8.
//
// Solidity: function SetAccountCollectionAddress(address AccountCollectionAddress) returns()
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactor) SetAccountCollectionAddress(opts *bind.TransactOpts, AccountCollectionAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.contract.Transact(opts, "SetAccountCollectionAddress", AccountCollectionAddress)
}

// SetAccountCollectionAddress is a paid mutator transaction binding the contract method 0xc8ac5fe8.
//
// Solidity: function SetAccountCollectionAddress(address AccountCollectionAddress) returns()
func (_EthoFSHostingContracts *EthoFSHostingContractsSession) SetAccountCollectionAddress(AccountCollectionAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.SetAccountCollectionAddress(&_EthoFSHostingContracts.TransactOpts, AccountCollectionAddress)
}

// SetAccountCollectionAddress is a paid mutator transaction binding the contract method 0xc8ac5fe8.
//
// Solidity: function SetAccountCollectionAddress(address AccountCollectionAddress) returns()
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorSession) SetAccountCollectionAddress(AccountCollectionAddress common.Address) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.SetAccountCollectionAddress(&_EthoFSHostingContracts.TransactOpts, AccountCollectionAddress)
}

// SetHostingContractCost is a paid mutator transaction binding the contract method 0x0b6ff6dd.
//
// Solidity: function SetHostingContractCost(uint256 set) returns()
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactor) SetHostingContractCost(opts *bind.TransactOpts, set *big.Int) (*types.Transaction, error) {
	return _EthoFSHostingContracts.contract.Transact(opts, "SetHostingContractCost", set)
}

// SetHostingContractCost is a paid mutator transaction binding the contract method 0x0b6ff6dd.
//
// Solidity: function SetHostingContractCost(uint256 set) returns()
func (_EthoFSHostingContracts *EthoFSHostingContractsSession) SetHostingContractCost(set *big.Int) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.SetHostingContractCost(&_EthoFSHostingContracts.TransactOpts, set)
}

// SetHostingContractCost is a paid mutator transaction binding the contract method 0x0b6ff6dd.
//
// Solidity: function SetHostingContractCost(uint256 set) returns()
func (_EthoFSHostingContracts *EthoFSHostingContractsTransactorSession) SetHostingContractCost(set *big.Int) (*types.Transaction, error) {
	return _EthoFSHostingContracts.Contract.SetHostingContractCost(&_EthoFSHostingContracts.TransactOpts, set)
}
