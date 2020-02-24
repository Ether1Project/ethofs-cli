package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
)

var controllerContractAddress = "0xc38B47169950D8A28bC77a6Fa7467464f25ADAFc"

func CalculateUploadCost(contractDuration int32, uploadSize int32) int32 {
	client, err := ethclient.Dial(rpcLocation)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(controllerContractAddress)
	instance, err := NewEthoFSController(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// Get hosting cost
	hostingCost, err := instance.HostingCost(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	cost := ((uploadSize / 1048576) * int32(hostingCost.Int64())) * (contractDuration / 46522)

	return cost
}

func CheckAccountExistence(accountAddress common.Address) bool {
	client, err := ethclient.Dial(rpcLocation)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(controllerContractAddress)
	instance, err := NewEthoFSController(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// Check account existence
	exists, err := instance.CheckAccountExistence(&bind.CallOpts{}, accountAddress)
	if err != nil {
		log.Fatal(err)
	}

	return exists
}

func UploadData(key string, contractCost uint64, mainHash string, contractName string, contractDuration uint32, uploadSize uint32, contentHashString string, contentPathString string) {
	client, err := ethclient.Dial(rpcLocation)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
        	log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	if CheckAccountExistence(fromAddress) {
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = new(big.Int).Mul(big.NewInt(int64(contractCost)), big.NewInt(1e+18)) // in wei
		auth.GasLimit = uint64(3000000) // in units
		auth.GasPrice = gasPrice

		address := common.HexToAddress(controllerContractAddress)
		instance, err := NewEthoFSController(address, client)
		if err != nil {
			log.Fatal(err)
		}

		// Initiaite upload tx
		tx, err := instance.AddNewContract(auth, mainHash, contractName, contractDuration, uploadSize, uploadSize, contentHashString, contentPathString)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Upload Tx Sent: %s", tx.Hash().Hex())
		fmt.Println("\n")
	} else {
		fmt.Println("Unable to find valid hosting account, please register your address")
		log.Fatal("\n")
	}
}
