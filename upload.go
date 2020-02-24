package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

var controllerContractAddress = "0xc38B47169950D8A28bC77a6Fa7467464f25ADAFc"

func UploadData(key string, contractCost uint64, mainHash string, contractName string, contractDuration uint64, uploadSize uint64, contentHashString string, contentPathString string) {
	client, err := ethclient.Dial(rpcAddress)
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
	instance, err := NewControllerContract(address, client)
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
}
