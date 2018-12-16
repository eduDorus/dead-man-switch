package actions

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

var contract *DeadSwitch
var auth *bind.TransactOpts
var userAuth *bind.TransactOpts
var sim *backends.SimulatedBackend

func init() {

	auth = createTransactor()
	userAuth = createTransactor()

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1337000000000000000)}
	alloc[userAuth.From] = core.GenesisAccount{Balance: big.NewInt(1337000000000000000)}
	sim = backends.NewSimulatedBackend(alloc, uint64(100000000000000))

	var transaction *types.Transaction
	var err error
	_, transaction, contract, err = DeployDeadSwitch(auth, sim)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mining...")
	sim.Commit()
	fmt.Println("Gas used: ", transaction.Gas())
}

func UploadFileToBlockchain(ipfsHash, address string) error {
	fmt.Println("Generating Address from Hex ", address)
	toAddress := common.HexToAddress(address)
	fmt.Println(toAddress)
	fmt.Println(contract)

	fmt.Println("Add ipfs hash", ipfsHash, "to blackchain...")
	_, err := contract.AddFile(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 23816230000000,
		Value:    big.NewInt(10),
	}, ipfsHash, auth.From)
	if err != nil {
		log.Println(err)
		return errors.WithStack(err)
	}
	fmt.Println("Mining...")
	sim.Commit()
	return nil
}

func createTransactor() *bind.TransactOpts {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	return bind.NewKeyedTransactor(key)
}
