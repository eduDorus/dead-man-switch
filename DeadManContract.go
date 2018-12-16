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

// DeadSwitchABI is the input ABI used to generate the binding from.
const DeadSwitchABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"string\"},{\"name\":\"_fileID\",\"type\":\"uint256\"}],\"name\":\"pubishKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ipfsHash\",\"type\":\"string\"},{\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"addFile\",\"outputs\":[{\"name\":\"fileId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ping\",\"type\":\"string\"},{\"name\":\"_fileID\",\"type\":\"uint256\"}],\"name\":\"ping\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"files\",\"outputs\":[{\"name\":\"fileOwner\",\"type\":\"address\"},{\"name\":\"ipfsHash\",\"type\":\"string\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"ping\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// DeadSwitchBin is the compiled bytecode used for deploying new contracts.
const DeadSwitchBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a0319163317905561083a806100326000396000f3fe60806040526004361061005b577c0100000000000000000000000000000000000000000000000000000000600035046345eeb16a8114610060578063c0a715d014610117578063e752a411146101e7578063f4c714b41461029c575b600080fd5b34801561006c57600080fd5b506101156004803603604081101561008357600080fd5b81019060208101813564010000000081111561009e57600080fd5b8201836020820111156100b057600080fd5b803590602001918460018302840111640100000000831117156100d257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610423915050565b005b34801561012357600080fd5b506101d56004803603604081101561013a57600080fd5b81019060208101813564010000000081111561015557600080fd5b82018360208201111561016757600080fd5b8035906020019184600183028401116401000000008311171561018957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050509035600160a060020a031691506104649050565b60408051918252519081900360200190f35b3480156101f357600080fd5b506101156004803603604081101561020a57600080fd5b81019060208101813564010000000081111561022557600080fd5b82018360208201111561023757600080fd5b8035906020019184600183028401116401000000008311171561025957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610550915050565b3480156102a857600080fd5b506102c6600480360360208110156102bf57600080fd5b50356105a0565b6040518085600160a060020a0316600160a060020a03168152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561032357818101518382015260200161030b565b50505050905090810190601f1680156103505780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b8381101561038357818101518382015260200161036b565b50505050905090810190601f1680156103b05780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b838110156103e35781810151838201526020016103cb565b50505050905090810190601f1680156104105780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b600054600160a060020a0316331461043a57600080fd5b6000818152600260208181526040909220845161045f93919092019190850190610773565b505050565b60008054600160a060020a0316331461047c57600080fd5b5060018054808201825560408051608081018252600160a060020a03858116825260208083018881528451808301865260008082528587019190915285518084018752818152606086015286815260028352949094208351815473ffffffffffffffffffffffffffffffffffffffff191693169290921782559251805194959294919361050e93850192910190610773565b506040820151805161052a916002840191602090910190610773565b5060608201518051610546916003840191602090910190610773565b5090505092915050565b6000818152600260205260409020548190600160a060020a0316331461057557600080fd5b6000828152600260209081526040909120845161059a92600390920191860190610773565b50505050565b6002602081815260009283526040928390208054600180830180548751601f60001994831615610100029490940190911696909604918201859004850286018501909652808552600160a060020a03909116949193928301828280156106475780601f1061061c57610100808354040283529160200191610647565b820191906000526020600020905b81548152906001019060200180831161062a57829003601f168201915b50505060028085018054604080516020601f60001961010060018716150201909416959095049283018590048502810185019091528181529596959450909250908301828280156106d95780601f106106ae576101008083540402835291602001916106d9565b820191906000526020600020905b8154815290600101906020018083116106bc57829003601f168201915b5050505060038301805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529495949350908301828280156107695780601f1061073e57610100808354040283529160200191610769565b820191906000526020600020905b81548152906001019060200180831161074c57829003601f168201915b5050505050905084565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106107b457805160ff19168380011785556107e1565b828001600101855582156107e1579182015b828111156107e15782518255916020019190600101906107c6565b506107ed9291506107f1565b5090565b61080b91905b808211156107ed57600081556001016107f7565b9056fea165627a7a723058202ce3b968bb8ebcb9ed74a55ddc4b5ade7a71f68c0c8ec436b67fe1550226a0f10029`

// DeployDeadSwitch deploys a new Ethereum contract, binding an instance of DeadSwitch to it.
func DeployDeadSwitch(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DeadSwitch, error) {
	parsed, err := abi.JSON(strings.NewReader(DeadSwitchABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DeadSwitchBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DeadSwitch{DeadSwitchCaller: DeadSwitchCaller{contract: contract}, DeadSwitchTransactor: DeadSwitchTransactor{contract: contract}, DeadSwitchFilterer: DeadSwitchFilterer{contract: contract}}, nil
}

// DeadSwitch is an auto generated Go binding around an Ethereum contract.
type DeadSwitch struct {
	DeadSwitchCaller     // Read-only binding to the contract
	DeadSwitchTransactor // Write-only binding to the contract
	DeadSwitchFilterer   // Log filterer for contract events
}

// DeadSwitchCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeadSwitchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeadSwitchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeadSwitchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeadSwitchFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeadSwitchFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeadSwitchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeadSwitchSession struct {
	Contract     *DeadSwitch       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeadSwitchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeadSwitchCallerSession struct {
	Contract *DeadSwitchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DeadSwitchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeadSwitchTransactorSession struct {
	Contract     *DeadSwitchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DeadSwitchRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeadSwitchRaw struct {
	Contract *DeadSwitch // Generic contract binding to access the raw methods on
}

// DeadSwitchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeadSwitchCallerRaw struct {
	Contract *DeadSwitchCaller // Generic read-only contract binding to access the raw methods on
}

// DeadSwitchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeadSwitchTransactorRaw struct {
	Contract *DeadSwitchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeadSwitch creates a new instance of DeadSwitch, bound to a specific deployed contract.
func NewDeadSwitch(address common.Address, backend bind.ContractBackend) (*DeadSwitch, error) {
	contract, err := bindDeadSwitch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeadSwitch{DeadSwitchCaller: DeadSwitchCaller{contract: contract}, DeadSwitchTransactor: DeadSwitchTransactor{contract: contract}, DeadSwitchFilterer: DeadSwitchFilterer{contract: contract}}, nil
}

// NewDeadSwitchCaller creates a new read-only instance of DeadSwitch, bound to a specific deployed contract.
func NewDeadSwitchCaller(address common.Address, caller bind.ContractCaller) (*DeadSwitchCaller, error) {
	contract, err := bindDeadSwitch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeadSwitchCaller{contract: contract}, nil
}

// NewDeadSwitchTransactor creates a new write-only instance of DeadSwitch, bound to a specific deployed contract.
func NewDeadSwitchTransactor(address common.Address, transactor bind.ContractTransactor) (*DeadSwitchTransactor, error) {
	contract, err := bindDeadSwitch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeadSwitchTransactor{contract: contract}, nil
}

// NewDeadSwitchFilterer creates a new log filterer instance of DeadSwitch, bound to a specific deployed contract.
func NewDeadSwitchFilterer(address common.Address, filterer bind.ContractFilterer) (*DeadSwitchFilterer, error) {
	contract, err := bindDeadSwitch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeadSwitchFilterer{contract: contract}, nil
}

// bindDeadSwitch binds a generic wrapper to an already deployed contract.
func bindDeadSwitch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeadSwitchABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeadSwitch *DeadSwitchRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DeadSwitch.Contract.DeadSwitchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeadSwitch *DeadSwitchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeadSwitch.Contract.DeadSwitchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeadSwitch *DeadSwitchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeadSwitch.Contract.DeadSwitchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeadSwitch *DeadSwitchCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DeadSwitch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeadSwitch *DeadSwitchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeadSwitch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeadSwitch *DeadSwitchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeadSwitch.Contract.contract.Transact(opts, method, params...)
}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 ) constant returns(address fileOwner, string ipfsHash, string key, string ping)
func (_DeadSwitch *DeadSwitchCaller) Files(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FileOwner common.Address
	IpfsHash  string
	Key       string
	Ping      string
}, error) {
	ret := new(struct {
		FileOwner common.Address
		IpfsHash  string
		Key       string
		Ping      string
	})
	out := ret
	err := _DeadSwitch.contract.Call(opts, out, "files", arg0)
	return *ret, err
}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 ) constant returns(address fileOwner, string ipfsHash, string key, string ping)
func (_DeadSwitch *DeadSwitchSession) Files(arg0 *big.Int) (struct {
	FileOwner common.Address
	IpfsHash  string
	Key       string
	Ping      string
}, error) {
	return _DeadSwitch.Contract.Files(&_DeadSwitch.CallOpts, arg0)
}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 ) constant returns(address fileOwner, string ipfsHash, string key, string ping)
func (_DeadSwitch *DeadSwitchCallerSession) Files(arg0 *big.Int) (struct {
	FileOwner common.Address
	IpfsHash  string
	Key       string
	Ping      string
}, error) {
	return _DeadSwitch.Contract.Files(&_DeadSwitch.CallOpts, arg0)
}

// AddFile is a paid mutator transaction binding the contract method 0xc0a715d0.
//
// Solidity: function addFile(string _ipfsHash, address _creator) returns(uint256 fileId)
func (_DeadSwitch *DeadSwitchTransactor) AddFile(opts *bind.TransactOpts, _ipfsHash string, _creator common.Address) (*types.Transaction, error) {
	return _DeadSwitch.contract.Transact(opts, "addFile", _ipfsHash, _creator)
}

// AddFile is a paid mutator transaction binding the contract method 0xc0a715d0.
//
// Solidity: function addFile(string _ipfsHash, address _creator) returns(uint256 fileId)
func (_DeadSwitch *DeadSwitchSession) AddFile(_ipfsHash string, _creator common.Address) (*types.Transaction, error) {
	return _DeadSwitch.Contract.AddFile(&_DeadSwitch.TransactOpts, _ipfsHash, _creator)
}

// AddFile is a paid mutator transaction binding the contract method 0xc0a715d0.
//
// Solidity: function addFile(string _ipfsHash, address _creator) returns(uint256 fileId)
func (_DeadSwitch *DeadSwitchTransactorSession) AddFile(_ipfsHash string, _creator common.Address) (*types.Transaction, error) {
	return _DeadSwitch.Contract.AddFile(&_DeadSwitch.TransactOpts, _ipfsHash, _creator)
}

// Ping is a paid mutator transaction binding the contract method 0xe752a411.
//
// Solidity: function ping(string _ping, uint256 _fileID) returns()
func (_DeadSwitch *DeadSwitchTransactor) Ping(opts *bind.TransactOpts, _ping string, _fileID *big.Int) (*types.Transaction, error) {
	return _DeadSwitch.contract.Transact(opts, "ping", _ping, _fileID)
}

// Ping is a paid mutator transaction binding the contract method 0xe752a411.
//
// Solidity: function ping(string _ping, uint256 _fileID) returns()
func (_DeadSwitch *DeadSwitchSession) Ping(_ping string, _fileID *big.Int) (*types.Transaction, error) {
	return _DeadSwitch.Contract.Ping(&_DeadSwitch.TransactOpts, _ping, _fileID)
}

// Ping is a paid mutator transaction binding the contract method 0xe752a411.
//
// Solidity: function ping(string _ping, uint256 _fileID) returns()
func (_DeadSwitch *DeadSwitchTransactorSession) Ping(_ping string, _fileID *big.Int) (*types.Transaction, error) {
	return _DeadSwitch.Contract.Ping(&_DeadSwitch.TransactOpts, _ping, _fileID)
}

// PubishKey is a paid mutator transaction binding the contract method 0x45eeb16a.
//
// Solidity: function pubishKey(string _key, uint256 _fileID) returns()
func (_DeadSwitch *DeadSwitchTransactor) PubishKey(opts *bind.TransactOpts, _key string, _fileID *big.Int) (*types.Transaction, error) {
	return _DeadSwitch.contract.Transact(opts, "pubishKey", _key, _fileID)
}

// PubishKey is a paid mutator transaction binding the contract method 0x45eeb16a.
//
// Solidity: function pubishKey(string _key, uint256 _fileID) returns()
func (_DeadSwitch *DeadSwitchSession) PubishKey(_key string, _fileID *big.Int) (*types.Transaction, error) {
	return _DeadSwitch.Contract.PubishKey(&_DeadSwitch.TransactOpts, _key, _fileID)
}

// PubishKey is a paid mutator transaction binding the contract method 0x45eeb16a.
//
// Solidity: function pubishKey(string _key, uint256 _fileID) returns()
func (_DeadSwitch *DeadSwitchTransactorSession) PubishKey(_key string, _fileID *big.Int) (*types.Transaction, error) {
	return _DeadSwitch.Contract.PubishKey(&_DeadSwitch.TransactOpts, _key, _fileID)
}
