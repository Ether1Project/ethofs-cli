
// Copyright 2020 The Etho.Black Development Team

package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/janeczku/go-spinner"
)

func main() {
	setFlags() // Activate user deginated options

	// Look for private key
	if privateKey == "" && !ipcFlag {
		fmt.Println("No private key detected - exiting")
		os.Exit(2) // Exit on no private key
	}

	// Set default public rpc if no custom rpc is set at runtime
	if rpcLocation == "" {
		rpcLocation = "https://rpc.ether1.org"
	}

	s := spinner.StartNew("Initializing ethofs-cli")
	time.Sleep(1 * time.Second)
	s.Stop()
	fmt.Println("✓ Initializing ethoFSfs-cli: Completed")

	// Check for upload flag
	if uploadFlag {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Initialize ethoFS Node
		ipfs, err := initializeEthofsNode(ctx)
		if err != nil {
			panic(fmt.Errorf("Error while initializing ethoFS node: %s", err))
		}

		// Check for recusrive directory upload
		if recursiveFlag && inputPath != "" && contractDuration > 0 {

			s := spinner.StartNew("Initializing ethoFS data upload")

			uploadDirectory, uploadSize, err := getUnixfsNode(inputPath)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}
			cidDirectory, err := ipfs.Unixfs().Add(ctx, uploadDirectory)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}

			hash := strings.Split(cidDirectory.String(), "/")[2]
			contractCost, err := CalculateUploadCost(contractDuration, uploadSize)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}

			contentHashString := "ethoFSPinningChannel_alpha11:" + hash
			contentPathString := "ethoFSPinningChannel_alpha11:/"
			_, err = UploadData(privateKey, contractCost, hash, contractName, uint32(contractDuration), uint32(uploadSize), contentHashString, contentPathString)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}

			if verifyUpload(ctx, ipfs, cidDirectory) {
				s.Stop()
				fmt.Println("✓ Initializing ethoFS data upload: Completed")
				fmt.Printf("ethoFS upload hash\n%s\n", cidDirectory.String())
			} else {
				s.Stop()
				fmt.Println("X Initializing ethoFS data upload: Failed")
				fmt.Printf("ethoFS upload hash\n%s\n", cidDirectory.String())
			}
		} else if inputPath != "" && contractDuration > 0 {

			s = spinner.StartNew("Initializing ethoFS data upload")

			uploadFile, uploadSize, err := getUnixfsNode(inputPath)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}
			cidFile, err := ipfs.Unixfs().Add(ctx, uploadFile)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}

			hash := strings.Split(cidFile.String(), "/")[2]
			contractCost, err := CalculateUploadCost(contractDuration, uploadSize)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}

			contentHashString := "ethoFSPinningChannel_alpha11:" + hash
			contentPathString := "ethoFSPinningChannel_alpha11:/"
			_, err = UploadData(privateKey, contractCost, hash, contractName, uint32(contractDuration), uint32(uploadSize), contentHashString, contentPathString)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}

			if verifyUpload(ctx, ipfs, cidFile) {
				s.Stop()
                	        fmt.Println("✓ Initializing ethoFS data upload: Completed")
                        	fmt.Printf("ethoFS upload hash\n%s\n", cidFile.String())
			} else {
				s.Stop()
                        	fmt.Println("X Initializing ethoFS data upload: Failed")
                	        fmt.Printf("ethoFS upload hash\n%s\n", cidFile.String())
			}
		}

		s = spinner.StartNew("Stopping ethoFS Node")
		time.Sleep(5 * time.Second)
		s.Stop()
		fmt.Println("✓ Stopping ethoFS Node: Completed")

	} else if listFlag {

		// List existing ethoFS hosting contracts
		_, err := ListExistingContracts(privateKey)
		if err != nil {
			panic(fmt.Errorf("Error finding existing contracts: %s", err))
		}

	} else if removeFlag && contractName != "" {

		// Initiate ethoFS hosting contract removal
		contractDetails, err := GetContractDetails(privateKey, contractName)
		if err != nil {
			s.Stop()
			panic(fmt.Errorf("Error removing ethoFS contract: %s", err))
		}

		_, err = RemoveContract(privateKey, contractDetails.Address)
		if err != nil {
			s.Stop()
			panic(fmt.Errorf("Error removing ethoFS contract: %s", err))
		}

	} else if extendFlag && contractName != "" && contractDuration > 0 {

		// Initiate ethoFS hosting contract extension
		contractDetails, err := GetContractDetails(privateKey, contractName)
		if err != nil {
			s.Stop()
			panic(fmt.Errorf("Error extending ethoFS contract: %s", err))
		}

		cost, err := CalculateUploadCost(contractDuration, int64(contractDetails.Size))
		if err != nil {
			s.Stop()
			panic(fmt.Errorf("Error extending ethoFS contract: %s", err))
		}

		_, err = ExtendContract(privateKey, cost, contractDetails.Address, contractDuration)
		if err != nil {
			s.Stop()
			panic(fmt.Errorf("Error extending ethoFS contract: %s", err))
		}

	} else if (replaceFlag || backupFlag) && contractName != "" && inputPath != "" && contractDuration > 0 {

		// Initiate ethoFS hosting contract replacement
		contractDetails, err := GetContractDetails(privateKey, contractName)
		if err != nil {
			s.Stop()
			panic(fmt.Errorf("Error replacing ethoFS contract: %s", err))
		}

		nullContractDetails := ContractDetails{}
		if contractDetails != nullContractDetails {
			_, err := RemoveContract(privateKey, contractDetails.Address)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error replacing ethoFS contract: %s", err))
			}
		}

		// Start ethofs node initialization
		s := spinner.StartNew("Initializing ethoFS node for upload")
		time.Sleep(3 * time.Second)
		s.Stop()
		fmt.Println("✓ Initializing ethoFS node for upload: Completed")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Spawn a node using a temporary path, creating a temporary repo for the run
		s = spinner.StartNew("Initializing ethoFS data upload repo")
		time.Sleep(3 * time.Second)
		s.Stop()
		fmt.Println("✓ Initializing ethoFS data upload repo: Completed")

		s = spinner.StartNew("Finalizing ethoFS node deployment")
		ipfs, err := spawnEphemeral(ctx)
		if err != nil {
			panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
		}
		s.Stop()
		fmt.Println("✓ Finalizing ethoFS node deployment: Completed")
		fmt.Println("✓ ethoFS Node is Running")

		s = spinner.StartNew("Syncing ethoFS bootnodes with ETHO network contract")

		bootstrapNodes, err := GetBootnodeContractValues()
		if err != nil {
			panic(fmt.Errorf("failed to sync bootnodes with ether-1 network: %s", err))
		}
		time.Sleep(3 * time.Second)
		s.Stop()
		fmt.Println("✓ Syncing ethoFS bootnodes with ETHO network contract: Completed")

		s = spinner.StartNew("Waiting for ethoFS bootnode connections")
		connectedPeers,_ := connectToPeers(ctx, ipfs, bootstrapNodes)
		time.Sleep(3 * time.Second)
		s.Stop()
		if connectedPeers > 1 {
			fmt.Println("✓ Waiting for ethoFS bootnode connections: Completed")
		} else {
			panic(fmt.Errorf("failed to connect to ethoFS bootnodes"))
		}
		// ethoFS node is completely initialized by now - full swarm

		if recursiveFlag {

			s = spinner.StartNew("Initializing ethoFS data upload")

			uploadDirectory, uploadSize, err := getUnixfsNode(inputPath)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}
			cidDirectory, err := ipfs.Unixfs().Add(ctx, uploadDirectory)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}

			hash := strings.Split(cidDirectory.String(), "/")[2]
			contractCost, err := CalculateUploadCost(contractDuration, uploadSize)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}

			contentHashString := "ethoFSPinningChannel_alpha11:" + hash
			contentPathString := "ethoFSPinningChannel_alpha11:/"
			_, err = UploadData(privateKey, contractCost, hash, contractName, uint32(contractDuration), uint32(uploadSize), contentHashString, contentPathString)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}

			if verifyUpload(ctx, ipfs, cidDirectory) {
				s.Stop()
				fmt.Println("✓ Initializing ethoFS data upload: Completed")
				fmt.Printf("ethoFS upload hash\n%s\n", cidDirectory.String())
			} else {
				s.Stop()
				fmt.Println("X Initializing ethoFS data upload: Failed")
				fmt.Printf("ethoFS upload hash\n%s\n", cidDirectory.String())
			}
		} else {

			s = spinner.StartNew("Initializing ethoFS data upload")

			uploadFile, uploadSize, err := getUnixfsNode(inputPath)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}
			cidFile, err := ipfs.Unixfs().Add(ctx, uploadFile)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}

			hash := strings.Split(cidFile.String(), "/")[2]
			contractCost, err := CalculateUploadCost(contractDuration, uploadSize)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}

			contentHashString := "ethoFSPinningChannel_alpha11:" + hash
			contentPathString := "ethoFSPinningChannel_alpha11:/"
			_, err = UploadData(privateKey, contractCost, hash, contractName, uint32(contractDuration), uint32(uploadSize), contentHashString, contentPathString)
			if err != nil {
				s.Stop()
				panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
			}

			if verifyUpload(ctx, ipfs, cidFile) {
				s.Stop()
                	        fmt.Println("✓ Initializing ethoFS data upload: Completed")
                        	fmt.Printf("ethoFS upload hash\n%s\n", cidFile.String())
			} else {
				s.Stop()
                        	fmt.Println("X Initializing ethoFS data upload: Failed")
                	        fmt.Printf("ethoFS upload hash\n%s\n", cidFile.String())
			}
		}

		s = spinner.StartNew("Stopping ethoFS Node")
		time.Sleep(5 * time.Second)
		s.Stop()
		fmt.Println("✓ Stopping ethoFS Node: Completed")
	} else if registerFlag && ethofsUsername != "" {

		// Initiate new ethoFS user registration
		s := spinner.StartNew("Initiating ethoFS Registration")
		time.Sleep(2 * time.Second)
		s.Stop()
		fmt.Println("✓ StInitiating ethoFS Registration: Completed")
		_, err := RegisterAccount(privateKey, ethofsUsername)
		if err != nil {
			s.Stop()
			panic(fmt.Errorf("Error registering ethoFS account: %s", err))
		}

	} else if ipcFlag {

		// Deploy IPC endpoint to wait for requests & respond
		ethofsIpc()

	}
}
