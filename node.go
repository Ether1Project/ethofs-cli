// Copyright 2020 The Etho.Black Development Team

package main

import (
	"context"
	"fmt"
	"time"

	libp2p "github.com/ipfs/go-ipfs/core/node/libp2p"
	icore "github.com/ipfs/interface-go-ipfs-core"
	corepnet "github.com/libp2p/go-libp2p-core/pnet"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	"github.com/janeczku/go-spinner"
)

//var swarmkey = []byte("/key/swarm/psk/1.0.0/\n/base16/\n38307a74b2176d0054ffa2864e31ee22d0fc6c3266dd856f6d41bddf14e2ad63")

// Creates an ethoFS node and returns coreAPI
func createNode(ctx context.Context, repoPath string) (icore.CoreAPI, error) {

	corepnet.ForcePrivateNetwork = true

	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	// Construct the node

	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption,
		// Routing: libp2p.DHTClientOption, // DHT Client mode only
		Repo: repo,
	}

	node, err := core.NewNode(ctx, nodeOptions)
	if err != nil {
		return nil, err
	}

	// Attach the Core API to the constructed node
	return coreapi.NewCoreAPI(node)
}

func initializeEthofsNode(ctx context.Context) (icore.CoreAPI, error) {
	// Start ethofs node initialization
	s := spinner.StartNew("Initializing ethoFS node for upload")
	time.Sleep(3 * time.Second)
	s.Stop()
	fmt.Println("✓ Initializing ethoFS node for upload: Completed")

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	// Spawn a node using a temporary path, creating a temporary repo for the run
	s = spinner.StartNew("Initializing ethoFS data upload repo")
	time.Sleep(3 * time.Second)
	s.Stop()
	fmt.Println("✓ Initializing ethoFS data upload repo: Completed")

	s = spinner.StartNew("Finalizing ethoFS node deployment")
	ipfs, err := spawnEphemeral(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
		return nil, err
	}
	s.Stop()
	fmt.Println("✓ Finalizing ethoFS node deployment: Completed")
	fmt.Println("✓ ethoFS Node is Running")

	s = spinner.StartNew("Syncing ethoFS bootnodes with ETHO network contract")

	bootstrapNodes, err := GetBootnodeContractValues()
	if err != nil {
		panic(fmt.Errorf("failed to sync bootnodes with ether-1 network: %s", err))
		return nil, err
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
		return nil, fmt.Errorf("failed to connec to ethofs bootnodes")
	}

	return ipfs, nil
}
