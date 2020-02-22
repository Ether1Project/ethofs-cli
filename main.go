// Copyright 2020 The Etho.Black Development Team

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	setFlags() // Activate user deginated options

	// Work in progress
	fmt.Println("Initializing ethoFS Node For Upload ")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Spawn a node using a temporary path, creating a temporary repo for the run
	fmt.Println("Spawning node on a temporary repo")
	ipfs, err := spawnEphemeral(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
	}

	fmt.Println("ethoFS node is running")
	time.Sleep(5 * time.Second)

	fmt.Println("Syncing ethoFS bootnodes with ETHO network contract")
	bootstrapNodes, err := GetBootnodeContractValues()
	if err != nil {
		panic(fmt.Errorf("failed to sync bootnodes with ether-1 network: %s", err))
	}

	fmt.Println("Waiting for ethoFS bootnode connections")
	connectedPeers,_ := connectToPeers(ctx, ipfs, bootstrapNodes)
	if connectedPeers > 1 {
		fmt.Println("ethoFS bootnode connections successful")
	} else {
		panic(fmt.Errorf("failed to connect to ethoFS bootnodes"))
	}

	if uploadFlag && recursiveFlag && inputPath != "" {

		uploadDirectory,_, err := getUnixfsNode(inputPath)
		if err != nil {
			panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
		}
		cidDirectory, err := ipfs.Unixfs().Add(ctx, uploadDirectory)
		if err != nil {
			panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
		}
		fmt.Printf("ethoFS data upload complete\n\nUpload Hash\n%s\n", cidDirectory.String())

	} else if uploadFlag && inputPath != "" {

		uploadFile,_, err := getUnixfsNode(inputPath)
		if err != nil {
			panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
		}
		cidFile, err := ipfs.Unixfs().Add(ctx, uploadFile)
		if err != nil {
			panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
		}
		fmt.Printf("ethoFS data upload complete\n\nUpload Hash\n%s\n", cidFile.String())

	}

	time.Sleep(5 * time.Second)
	fmt.Println("Stopping ethoFS node")
	time.Sleep(1 * time.Second)
	fmt.Println("Exiting..")
}
