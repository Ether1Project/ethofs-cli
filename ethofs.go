// Copyright 2020 The Etho.Black Team

package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/janeczku/go-spinner"
)

var controllerContractAddress = "0xc38B47169950D8A28bC77a6Fa7467464f25ADAFc"

type ContractDetails struct {
	Name string
	Address common.Address
	MainHash string
}

//WaitForTx returns on tx verification
func WaitForTx(client *ethclient.Client, hash common.Hash) bool {
        s := spinner.StartNew("Waiting for transaction confirmation")

	_, err := waitForTxConfirmations(client, hash, 1)
	if err != nil {
        	s.Stop()
	        fmt.Println("X Waiting for transaction confirmation: Failed")
		log.Fatal(err)
	}
        s.Stop()
        fmt.Println("✓ Waiting for transaction confirmation: Confirmed")
	return true
}

//CheckAccountBalance verifies sufficient balance at specified address
func CheckAccountBalance(address common.Address, amount *big.Int) bool {
	client, err := ethclient.Dial(rpcLocation)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	if balance.Cmp(amount) >= 0 {
		return true
	}
	return false
}

//CalculateUploadCost returns the cost to upload to ethoFS based on provided upload size
func CalculateUploadCost(contractDuration int32, uploadSize int64) *big.Int {

/*	fmt.Printf("\nInitiating Hosting Cost Calculation - Duration: %d Size: %d\n", contractDuration, uploadSize)

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
	//hostingCost, err := instance.HostingCost(&bind.CallOpts{})
	hostingCost, err := instance.HostingCost()
	if err != nil {
		log.Fatal(err)
	}
*/
	hostingCost := big.NewInt(1e+18)
	num1 := new(big.Int).Mul(big.NewInt(uploadSize), hostingCost)
	num2 := new(big.Int).Mul(num1, big.NewInt(int64(contractDuration)))
	num3 := new(big.Int).Div(num2, big.NewInt(1048576))
	cost := new(big.Int).Div(num3, big.NewInt(46522))

	//fmt.Printf("Hosting Cost Calculated: %d\n", cost)

	return cost
}

//CheckAccountExistence verifies an ethoFS account has been registered
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

//UploadData initiates the ethoFS upload tx
func UploadData(key string, contractCost *big.Int, mainHash string, contractName string, contractDuration uint32, uploadSize uint32, contentHashString string, contentPathString string) {

	//fmt.Printf("Upload Data - Cost: %d Hash: %s Name: %s Duration: %d Size: %d HashString: %s PathString: %s\n", contractCost, mainHash, contractName, contractDuration, uploadSize, contentHashString, contentPathString)

        s := spinner.StartNew("Sending ethoFS upload transaction")

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
		if CheckAccountBalance(fromAddress, contractCost) {
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
			auth.Value = contractCost // in wei
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

		        s.Stop()
        		fmt.Println("✓ Sending ethoFS upload transaction: Complete")
			//fmt.Printf("Upload Tx Sent: %s", tx.Hash().Hex())
			//fmt.Println("\n")

			// Wait for tx confirmation
			WaitForTx(client, tx.Hash())

		} else {
			fmt.Println("Insufficient balance for upload")
			log.Fatal("\n")
		}
	} else {
		fmt.Println("Unable to find valid hosting account, please register your address")
		log.Fatal("\n")
	}
}

//RegisterAccount initates the ethoFS registration tx
func RegisterAccount(key string, name string) {
	s := spinner.StartNew("Sending ethoFS registration transaction")

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

	if !CheckAccountExistence(fromAddress) {
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
		auth.GasLimit = uint64(3000000) // in units
		auth.GasPrice = gasPrice

		address := common.HexToAddress(controllerContractAddress)
		instance, err := NewEthoFSController(address, client)
		if err != nil {
			log.Fatal(err)
		}

		// Initiaite registration tx
		tx, err := instance.AddNewUserPublic(auth, name)
		if err != nil {
			log.Fatal(err)
		}
	        s.Stop()
        	fmt.Println("✓ Sending ethoFS registration transaction: Completed")
		fmt.Printf("Registration Tx Sent: %s", tx.Hash().Hex())

		WaitForTx(client, tx.Hash())

	} else {
	        s.Stop()
        	fmt.Println("X Sending ethoFS registration transaction: Failed")

		fmt.Println("ethoFS hosting account already registered")
		log.Fatal("\n")
	}
}

//ListExistingContracts lists active ethoFS hosting contracts for current user
func ListExistingContracts(key string) {

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

	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	address := common.HexToAddress(controllerContractAddress)
	instance, err := NewEthoFSController(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// Get account user name
	name, err := instance.GetUserAccountName(&bind.CallOpts{}, accountAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Get existing contract count
	count, err := instance.GetUserAccountTotalContractCount(&bind.CallOpts{}, accountAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nExisting Hosting Contracts - %s\n", name)
	fmt.Println("Contract Name,Contract Address,Contract Hash,Deployment Block,Expiration Block")

	for i := uint32(0); i < count; i++ {

		// Get hosting contract address
		contractAddress, err := instance.GetHostingContractAddress(&bind.CallOpts{}, accountAddress, big.NewInt(int64(i)))
		if err != nil {
			log.Fatal(err)
		}

		// Get hosting contract deployed block height
		deploymentBlock, err := instance.GetHostingContractDeployedBlockHeight(&bind.CallOpts{}, contractAddress)
		if err != nil {
			log.Fatal(err)
		}

		// Get hosting contract expiration block height
		expirationBlock, err := instance.GetHostingContractExpirationBlockHeight(&bind.CallOpts{}, contractAddress)
		if err != nil {
			log.Fatal(err)
		}

		// Get hosting contract name
		contractName, err := instance.GetHostingContractName(&bind.CallOpts{}, contractAddress)
		if err != nil {
			log.Fatal(err)
		}

		// Get hosting contract main hash
		contractMainHash, err := instance.GetMainContentHash(&bind.CallOpts{}, contractAddress)
		if err != nil {
			log.Fatal(err)
		}

		if deploymentBlock.Int64() > 0 && expirationBlock.Int64() > 0 {
			fmt.Printf("%s,%s,%s,%d,%d\n", contractName, contractAddress.String(), contractMainHash, deploymentBlock, expirationBlock)
		}
	}
}

//ExtendContract initates the ethoFS contract extension tx
func ExtendContract(key string, extensionCost *big.Int, contractAddress string, duration uint32) {
	s := spinner.StartNew("Sending ethoFS contract extension transaction")

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
	auth.Value = extensionCost // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(controllerContractAddress)
	instance, err := NewEthoFSController(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// Initiaite extension tx
	tx, err := instance.ExtendContract(auth, common.HexToAddress(contractAddress), duration)
	if err != nil {
		log.Fatal(err)
	}

        s.Stop()
        fmt.Println("✓ Sending ethoFS contract extension transaction: Completed")

	fmt.Printf("Contract Extension Tx Sent: %s", tx.Hash().Hex())

	WaitForTx(client, tx.Hash())
}

//RemoveContract initates the ethoFS contract removal tx
func RemoveContract(key string, contractAddress string) {
	s := spinner.StartNew("Sending ethoFS contract removal transaction")

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
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(controllerContractAddress)
	instance, err := NewEthoFSController(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// Get hosting contract main hash
	contractMainHash, err := instance.GetMainContentHash(&bind.CallOpts{}, common.HexToAddress(contractAddress))
	if err != nil {
		log.Fatal(err)
	}

	// Initiaite removal tx
	tx, err := instance.RemoveHostingContract(auth, common.HexToAddress(contractAddress), contractMainHash)
	if err != nil {
		log.Fatal(err)
	}

        s.Stop()
        fmt.Println("✓ Sending ethoFS contract removal transaction: Completed")

	fmt.Printf("Contract Removal Tx Sent: %s", tx.Hash().Hex())

	WaitForTx(client, tx.Hash())
}

func GetContractDetails(privateKey string, name string) ContractDetails {

	contractDetails := ContractDetails{}

	return contractDetails
}

func waitForTxConfirmations(client *ethclient.Client, txHash common.Hash, n uint64) (*types.Receipt, error) {
	rpcTimeout := 120 * time.Second
	var (
		receipt    *types.Receipt
		startBlock *types.Block
		err        error
	)

	for i := 0; i < 90; i++ {
		ctx, _ := context.WithTimeout(context.Background(), rpcTimeout)
		receipt, err = client.TransactionReceipt(ctx, txHash)
		if err != nil && err != ethereum.NotFound {
			return nil, err
		}
		if receipt != nil {
			break
		}
		time.Sleep(time.Second)
	}

	if receipt == nil {
		return nil, ethereum.NotFound
	}

	ctx, _ := context.WithTimeout(context.Background(), rpcTimeout)
	if startBlock, err = client.BlockByNumber(ctx, nil); err != nil {
		return nil, err
	}

	for i := 0; i < 90; i++ {
		ctx, _ := context.WithTimeout(context.Background(), rpcTimeout)
		currentBlock, err := client.BlockByNumber(ctx, nil)
		if err != nil {
			return nil, err
		}

		if startBlock.NumberU64()+n >= currentBlock.NumberU64() {
			ctx, _ := context.WithTimeout(context.Background(), rpcTimeout)
			if checkReceipt, err := client.TransactionReceipt(ctx, txHash); checkReceipt != nil {
				if bytes.Compare(receipt.PostState, checkReceipt.PostState) == 0 {
					return receipt, nil
				} else { // chain reorg
					waitForTxConfirmations(client, txHash, n)
				}
			} else {
				return nil, err
			}
		}

		time.Sleep(time.Second)
	}

	return nil, ethereum.NotFound
}
