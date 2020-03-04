// Copyright 2020 The Etho.Black Team

package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
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
	Name string		`json:"name"`
	Address common.Address  `json:"address"`
	MainHash string         `json:"mainhash"`
	Size uint32             `json;"size"`
}

//WaitForTx returns on tx verification
func WaitForTx(client *ethclient.Client, hash common.Hash) (bool, error) {
        s := spinner.StartNew("Waiting for transaction confirmation")

	_, err := waitForTxConfirmations(client, hash, 1)
	if err != nil {
        	s.Stop()
	        fmt.Println("X Waiting for transaction confirmation: Failed")
		return false, err
	}
        s.Stop()
        fmt.Println("✓ Waiting for transaction confirmation: Confirmed")
	return true, nil
}

//CheckAccountBalance verifies sufficient balance at specified address
func CheckAccountBalance(address common.Address, amount *big.Int) (bool, error) {
	client, err := ethclient.Dial(rpcLocation)
	if err != nil {
		return false, err
	}

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return false, err
	}

	if balance.Cmp(amount) >= 0 {
		return true, nil
	}
	return false, nil
}

//CalculateUploadCost returns the cost to upload to ethoFS based on provided upload size
func CalculateUploadCost(contractDuration int32, uploadSize int64) (*big.Int, error) {

/*	fmt.Printf("\nInitiating Hosting Cost Calculation - Duration: %d Size: %d\n", contractDuration, uploadSize)

	client, err := ethclient.Dial(rpcLocation)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(controllerContractAddress)
	instance, err := NewEthoFSController(address, client)
	if err != nil {
		return nil, err
	}

	// Get hosting cost
	//hostingCost, err := instance.HostingCost(&bind.CallOpts{})
	hostingCost, err := instance.HostingCost()
	if err != nil {
		return nil, err
	}
*/
	hostingCost := big.NewInt(1e+18)
	num1 := new(big.Int).Mul(big.NewInt(uploadSize), hostingCost)
	num2 := new(big.Int).Mul(num1, big.NewInt(int64(contractDuration)))
	num3 := new(big.Int).Div(num2, big.NewInt(1048576))
	cost := new(big.Int).Div(num3, big.NewInt(46522))

	//fmt.Printf("Hosting Cost Calculated: %d\n", cost)

	return cost, nil
}

//CheckAccountExistence verifies an ethoFS account has been registered
func CheckAccountExistence(accountAddress common.Address) (bool, error) {
	client, err := ethclient.Dial(rpcLocation)
	if err != nil {
		return false, err
	}

	address := common.HexToAddress(controllerContractAddress)
	instance, err := NewEthoFSController(address, client)
	if err != nil {
		return false, err
	}

	// Check account existence
	exists, err := instance.CheckAccountExistence(&bind.CallOpts{}, accountAddress)
	if err != nil {
		return false, err
	}

	return exists, nil
}

//UploadData initiates the ethoFS upload tx
func UploadData(key string, contractCost *big.Int, mainHash string, contractName string, contractDuration uint32, uploadSize uint32, contentHashString string, contentPathString string) (bool, error) {

	//fmt.Printf("Upload Data - Cost: %d Hash: %s Name: %s Duration: %d Size: %d HashString: %s PathString: %s\n", contractCost, mainHash, contractName, contractDuration, uploadSize, contentHashString, contentPathString)

        s := spinner.StartNew("Sending ethoFS upload transaction")

	client, err := ethclient.Dial(rpcLocation)
	if err != nil {
		return false, err
	}
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return false, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	exists, err := CheckAccountExistence(fromAddress)
	if err != nil {
		return false, err
	}

	if exists {

		sufficientBalance, err := CheckAccountBalance(fromAddress, contractCost)
		if err != nil {
			return false, err
		}

		if sufficientBalance {
			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				return false, err
			}

			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				return false, err
			}

			auth := bind.NewKeyedTransactor(privateKey)
			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = contractCost // in wei
			auth.GasLimit = uint64(3000000) // in units
			auth.GasPrice = gasPrice

			address := common.HexToAddress(controllerContractAddress)
			instance, err := NewEthoFSController(address, client)
			if err != nil {
				return false, err
			}

			// Initiaite upload tx
			tx, err := instance.AddNewContract(auth, mainHash, contractName, contractDuration, uploadSize, uploadSize, contentHashString, contentPathString)
			if err != nil {
				return false, err
			}

		        s.Stop()
        		fmt.Println("✓ Sending ethoFS upload transaction: Complete")

			// Wait for tx confirmation
			_, err = WaitForTx(client, tx.Hash())
			if err != nil {
				return false, err
			}


		} else {
			fmt.Println("Insufficient balance for upload")
			return false, fmt.Errorf("Insufficient balance for upload")
		}
	} else {
		fmt.Println("Unable to find valid hosting account, please register your address")
		return false, fmt.Errorf("Unable to find valid hosting account, please register your address")
	}

	return true, nil
}

//RegisterAccount initates the ethoFS registration tx
func RegisterAccount(key string, name string) (bool, error) {
	s := spinner.StartNew("Sending ethoFS registration transaction")

	client, err := ethclient.Dial(rpcLocation)
	if err != nil {
		return false, err
	}
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return false, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	exists, err := CheckAccountExistence(fromAddress)
	if err != nil {
		return false, err
	}

	if exists {
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return false, err
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return false, err
		}

		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasLimit = uint64(3000000) // in units
		auth.GasPrice = gasPrice

		address := common.HexToAddress(controllerContractAddress)
		instance, err := NewEthoFSController(address, client)
		if err != nil {
			return false, err
		}

		// Initiaite registration tx
		tx, err := instance.AddNewUserPublic(auth, name)
		if err != nil {
			return false, err
		}
	        s.Stop()
        	fmt.Println("✓ Sending ethoFS registration transaction: Completed")
		fmt.Printf("Registration Tx Sent: %s", tx.Hash().Hex())

		_, err = WaitForTx(client, tx.Hash())
		if err != nil {
			return false, err
		}

	} else {
	        s.Stop()
        	fmt.Println("X Sending ethoFS registration transaction: Failed")

		fmt.Println("ethoFS hosting account already registered")
		return false, fmt.Errorf("ethofs hosting account already registered")
	}
	return true, nil
}

//ListExistingContracts lists active ethoFS hosting contracts for current user
func ListExistingContracts(key string) ([]ContractDetails, error) {

	var contractDetails []ContractDetails

	client, err := ethclient.Dial(rpcLocation)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	address := common.HexToAddress(controllerContractAddress)
	instance, err := NewEthoFSController(address, client)
	if err != nil {
		return nil, err
	}

	// Get account user name
	name, err := instance.GetUserAccountName(&bind.CallOpts{}, accountAddress)
	if err != nil {
		return nil, err
	}

	// Get existing contract count
	count, err := instance.GetUserAccountTotalContractCount(&bind.CallOpts{}, accountAddress)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\nExisting Hosting Contracts - %s\n", name)
	fmt.Println("Contract Name,Contract Address,Contract Hash,Deployment Block,Expiration Block")

	for i := uint32(0); i < count; i++ {

		// Get hosting contract address
		contractAddress, err := instance.GetHostingContractAddress(&bind.CallOpts{}, accountAddress, big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}

		// Get hosting contract deployed block height
		deploymentBlock, err := instance.GetHostingContractDeployedBlockHeight(&bind.CallOpts{}, contractAddress)
		if err != nil {
			return nil, err
		}

		// Get hosting contract expiration block height
		expirationBlock, err := instance.GetHostingContractExpirationBlockHeight(&bind.CallOpts{}, contractAddress)
		if err != nil {
			return nil, err
		}

		// Get hosting contract name
		contractName, err := instance.GetHostingContractName(&bind.CallOpts{}, contractAddress)
		if err != nil {
			return nil, err
		}

		// Get hosting contract main hash
		contractMainHash, err := instance.GetMainContentHash(&bind.CallOpts{}, contractAddress)
		if err != nil {
			return nil, err
		}

		// Get hosting contract storage used
		contractStorageUsed, err := instance.GetHostingContractStorageUsed(&bind.CallOpts{}, contractAddress)
		if err != nil {
			return nil, err
		}

		contractDetail := ContractDetails{Address: contractAddress, Name: contractName, MainHash: contractMainHash, Size: contractStorageUsed}
		contractDetails = append(contractDetails, contractDetail)

		if deploymentBlock.Int64() > 0 && expirationBlock.Int64() > 0 {
			fmt.Printf("%s,%s,%s,%d,%d\n", contractName, contractAddress.String(), contractMainHash, deploymentBlock, expirationBlock)
		}
	}

	return contractDetails, nil
}

//ExtendContract initates the ethoFS contract extension tx
func ExtendContract(key string, extensionCost *big.Int, contractAddress common.Address, duration int32) (bool, error) {
	s := spinner.StartNew("Sending ethoFS contract extension transaction")

	client, err := ethclient.Dial(rpcLocation)
	if err != nil {
		return false, err
	}
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return false, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("error casting public key in ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return false, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return false, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = extensionCost // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(controllerContractAddress)
	instance, err := NewEthoFSController(address, client)
	if err != nil {
		return false, err
	}

	// Initiaite extension tx
	tx, err := instance.ExtendContract(auth, contractAddress, uint32(duration))
	if err != nil {
		return false, err
	}

        s.Stop()
        fmt.Println("✓ Sending ethoFS contract extension transaction: Completed")

	fmt.Printf("Contract Extension Tx Sent: %s", tx.Hash().Hex())

	_, err = WaitForTx(client, tx.Hash())
	if err != nil {
		return false, err
	}

	return true, nil
}

//RemoveContract initates the ethoFS contract removal tx
func RemoveContract(key string, contractAddress common.Address) (bool, error) {
	s := spinner.StartNew("Sending ethoFS contract removal transaction")

	client, err := ethclient.Dial(rpcLocation)
	if err != nil {
		return false, err
	}
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return false, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return false, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return false, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(controllerContractAddress)
	instance, err := NewEthoFSController(address, client)
	if err != nil {
		return false, err
	}

	// Get hosting contract main hash
	contractMainHash, err := instance.GetMainContentHash(&bind.CallOpts{}, contractAddress)
	if err != nil {
		return false, err
	}

	// Initiaite removal tx
	tx, err := instance.RemoveHostingContract(auth, contractAddress, contractMainHash)
	if err != nil {
		return false, err
	}

        s.Stop()
        fmt.Println("✓ Sending ethoFS contract removal transaction: Completed")

	fmt.Printf("Contract Removal Tx Sent: %s", tx.Hash().Hex())

	_, err = WaitForTx(client, tx.Hash())
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetContractDetails(key string, name string) (ContractDetails, error) {

	contractDetails := ContractDetails{}

	client, err := ethclient.Dial(rpcLocation)
	if err != nil {
		return contractDetails, err
	}

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return contractDetails, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return contractDetails, fmt.Errorf("error casting public key to ECDSA")
	}

	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	address := common.HexToAddress(controllerContractAddress)
	instance, err := NewEthoFSController(address, client)
	if err != nil {
		return contractDetails, err
	}

	// Get existing contract count
	count, err := instance.GetUserAccountTotalContractCount(&bind.CallOpts{}, accountAddress)
	if err != nil {
		return contractDetails, err
	}

	for i := uint32(0); i < count; i++ {

		// Get hosting contract address
		contractAddress, err := instance.GetHostingContractAddress(&bind.CallOpts{}, accountAddress, big.NewInt(int64(i)))
		if err != nil {
			return contractDetails, err
		}

		// Get hosting contract name
		contractName, err := instance.GetHostingContractName(&bind.CallOpts{}, contractAddress)
		if err != nil {
			return contractDetails, err
		}

		if contractName == name {

			// Get hosting contract deployed block height
			deploymentBlock, err := instance.GetHostingContractDeployedBlockHeight(&bind.CallOpts{}, contractAddress)
			if err != nil {
				return contractDetails, err
			}

			// Get hosting contract expiration block height
			expirationBlock, err := instance.GetHostingContractExpirationBlockHeight(&bind.CallOpts{}, contractAddress)
			if err != nil {
				return contractDetails, err
			}

			if deploymentBlock.Int64() > 0 && expirationBlock.Int64() > 0 {

				// Get hosting contract main hash
				contractMainHash, err := instance.GetMainContentHash(&bind.CallOpts{}, contractAddress)
				if err != nil {
					return contractDetails, err
				}

				// Get hosting contract storage used
				contractStorageUsed, err := instance.GetHostingContractStorageUsed(&bind.CallOpts{}, contractAddress)
				if err != nil {
					return contractDetails, err
				}

				contractDetails.Name = contractName
				contractDetails.Address = contractAddress
				contractDetails.MainHash = contractMainHash
				contractDetails.Size = contractStorageUsed
				fmt.Printf("Existing Contract Found - Name: %s Address: %s, Hash: %s\n", contractName, contractAddress.String(), contractMainHash)

			}
		}

	}
	return contractDetails, nil
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
