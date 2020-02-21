// Copyright 2020 The Etho.Black Development Team

package main

import (
	"context"

	libp2p "github.com/ipfs/go-ipfs/core/node/libp2p"
	icore "github.com/ipfs/interface-go-ipfs-core"
	corepnet "github.com/libp2p/go-libp2p-core/pnet"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
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
