// Copyright 2020 The Etho.Black Development Team

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/janeczku/go-spinner"
)

func main() {
	setFlags() // Activate user deginated options

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

	if uploadFlag && recursiveFlag && inputPath != "" {

		s = spinner.StartNew("Initializing ethoFS data upload")

		uploadDirectory,_, err := getUnixfsNode(inputPath)
		if err != nil {
			s.Stop()
			panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
		}
		cidDirectory, err := ipfs.Unixfs().Add(ctx, uploadDirectory)
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
	} else if uploadFlag && inputPath != "" {

		s = spinner.StartNew("Initializing ethoFS data upload")

		uploadFile,_, err := getUnixfsNode(inputPath)
		if err != nil {
			s.Stop()
			panic(fmt.Errorf("Error uploading data to ethoFS: %s", err))
		}
		cidFile, err := ipfs.Unixfs().Add(ctx, uploadFile)
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
}
