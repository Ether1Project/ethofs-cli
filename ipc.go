// Copyright 2020 The Etho.Black Development Team

package main

import (
	"context"
        "errors"
	"fmt"
        "net"
	"strings"

        "github.com/ethereum/go-ethereum/rpc"
)

type EthofsService struct {}

func (s *EthofsService) Register(name, key string) (string, error) {
	_, err := RegisterAccount(key, name)
	if err != nil {
	        return "Registration Failed", err
	}
        return "Successfully Registered " + name, nil
}

func (s *EthofsService) List(key string) ([]ContractDetails, error) {
	contractData, err := ListExistingContracts(key)
	if err != nil {
		return nil, err
	}
	return contractData, nil
}

func (s *EthofsService) Add(key string, path string, name string, blocks uint64, recursive bool) (string, error) {
        if blocks == 0 {
                return "ethofs Data Upload Failed", errors.New("block count is too low")
        }

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize ethoFS Node
	ipfs, err := initializeEthofsNode(ctx)
	if err != nil {
		return "ethofs Data Upload Failed", fmt.Errorf("Error while initializing ethoFS node: %s", err)
	}

	if recursive && path != "" {

		uploadDirectory, uploadSize, err := getUnixfsNode(path)
		if err != nil {
	                return "ethofs Data Upload Failed", errors.New("error uploading data")
		}
		cidDirectory, err := ipfs.Unixfs().Add(ctx, uploadDirectory)
		if err != nil {
	                return "ethofs Data Upload Failed", errors.New("error uploading data")
		}

		hash := strings.Split(cidDirectory.String(), "/")[2]
		contractCost := CalculateUploadCost(int32(blocks), uploadSize)
		contentHashString := "ethoFSPinningChannel_alpha11:" + hash
		contentPathString := "ethoFSPinningChannel_alpha11:/"
		UploadData(key, contractCost, hash, name, uint32(blocks), uint32(uploadSize), contentHashString, contentPathString)

		if verifyUpload(ctx, ipfs, cidDirectory) {
			return cidDirectory.String(), nil
		} else {
	                return cidDirectory.String(), errors.New("ethoFS upload failed")
		}
	} else if path != "" {

		uploadFile, uploadSize, err := getUnixfsNode(path)
		if err != nil {
	                return "ethofs Data Upload Failed", errors.New("error uploading data")
		}
		cidFile, err := ipfs.Unixfs().Add(ctx, uploadFile)
		if err != nil {
	                return "ethofs Data Upload Failed", errors.New("error uploading data")
		}

		hash := strings.Split(cidFile.String(), "/")[2]
		contractCost := CalculateUploadCost(int32(blocks), uploadSize)
		contentHashString := "ethoFSPinningChannel_alpha11:" + hash
		contentPathString := "ethoFSPinningChannel_alpha11:/"
		UploadData(key, contractCost, hash, name, uint32(blocks), uint32(uploadSize), contentHashString, contentPathString)

		if verifyUpload(ctx, ipfs, cidFile) {
			return cidFile.String(), nil
		} else {
	                return cidFile.String(), errors.New("ethoFS upload failed")
		}
	}
	return "ethofs Data Upload Failed", errors.New("params incorrect")
}

func ethofsIpc() {
        ethofs := new(EthofsService)
        server := rpc.NewServer()
        server.RegisterName("ethofs", ethofs)
        l, _ := net.ListenUnix("unix", &net.UnixAddr{Net: "unix", Name: "ethofs.ipc"})
        server.ServeListener(l)
}

/*func main() {

	} else if listFlag {

		ListExistingContracts(privateKey)

	} else if removeFlag && contractName != "" {

		contractDetails := GetContractDetails(privateKey, contractName)
		RemoveContract(privateKey, contractDetails.Address)

	} else if extendFlag && contractName != "" && contractDuration > 0 {

		contractDetails := GetContractDetails(privateKey, contractName)
		cost := CalculateUploadCost(contractDuration, int64(contractDetails.Size))
		ExtendContract(privateKey, cost, contractDetails.Address, contractDuration)

	} else if (replaceFlag || backupFlag) && contractName != "" && inputPath != "" && contractDuration > 0 {

		contractDetails := GetContractDetails(privateKey, contractName)
		nullContractDetails := ContractDetails{}
		if contractDetails != nullContractDetails {
			RemoveContract(privateKey, contractDetails.Address)
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
			contractCost := CalculateUploadCost(contractDuration, uploadSize)
			contentHashString := "ethoFSPinningChannel_alpha11:" + hash
			contentPathString := "ethoFSPinningChannel_alpha11:/"
			UploadData(privateKey, contractCost, hash, contractName, uint32(contractDuration), uint32(uploadSize), contentHashString, contentPathString)

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
			contractCost := CalculateUploadCost(contractDuration, uploadSize)
			contentHashString := "ethoFSPinningChannel_alpha11:" + hash
			contentPathString := "ethoFSPinningChannel_alpha11:/"
			UploadData(privateKey, contractCost, hash, contractName, uint32(contractDuration), uint32(uploadSize), contentHashString, contentPathString)

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

		s := spinner.StartNew("Initiating ethoFS Registration")
		time.Sleep(2 * time.Second)
		s.Stop()
		fmt.Println("✓ StInitiating ethoFS Registration: Completed")

	} else if ipcFlag {

		// Deploy IPC endpoint to wait for requests & respond
		ethofsIpc()

	}
}
*/
