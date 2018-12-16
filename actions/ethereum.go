package actions

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
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

type file struct {
	FileOwner common.Address
	IpfsHash  string
	Key       string
	Ping      string
}

func init() {

	auth, _ = createTransactor()
	userAuth, userKey := createTransactor()

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

	fmt.Println("PrivateKey:", hex.EncodeToString(userKey.D.Bytes()))
	fmt.Println("PublicKey:", crypto.PubkeyToAddress(userKey.PublicKey).Hex())

	// pk := hex.EncodeToString(userKey.D.Bytes())
	// bpk, _ := hex.DecodeString(pk)
	// prk, _ := crypto.ToECDSA(bpk)
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
	}, ipfsHash, toAddress)
	if err != nil {
		log.Println(err)
		return errors.WithStack(err)
	}
	fmt.Println("Mining...")
	sim.Commit()

	file, err := contract.Files(&bind.CallOpts{
		From: auth.From,
	}, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}
	jsonObject, _ := json.Marshal(file)
	fmt.Println(string(jsonObject))
	return nil
}

func ReadFilesFromBlockchain() []file {

	var files []file
	fileIndex := 0

	for {
		file := readFileFromBlockchain(big.NewInt(int64(fileIndex)))
		fileIndex++
		if file.IpfsHash == "" {
			break
		} else {
			files = append(files, file)
		}
	}

	return files
}

func readFileFromBlockchain(id *big.Int) file {
	file, err := contract.Files(&bind.CallOpts{
		From: auth.From,
	}, id)
	if err != nil {
		log.Fatal(err)
	}
	jsonObject, _ := json.Marshal(file)
	fmt.Println(string(jsonObject))

	return file
}

func createTransactor() (*bind.TransactOpts, *ecdsa.PrivateKey) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	return bind.NewKeyedTransactor(key), key
}
