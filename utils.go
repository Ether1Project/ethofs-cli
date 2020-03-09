// Copyright 2020 The Etho.Black Development Team

package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	options "github.com/ipfs/interface-go-ipfs-core/options"
	path "github.com/ipfs/interface-go-ipfs-core/path"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/libp2p/go-libp2p-core/peer"
)

// connectToPeers connect to specific peer via id
func connectToPeers(ctxOLD context.Context, ipfs icore.CoreAPI, peers []string) (int64, error) {

	connectedPeers := int64(0)

	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	var wg sync.WaitGroup
	peerInfos := make(map[peer.ID]*peerstore.PeerInfo, len(peers))
	for _, addrStr := range peers {
		addr, err := ma.NewMultiaddr(addrStr)
		if err != nil {
			return connectedPeers, err
		}
		pii, err := peerstore.InfoFromP2pAddr(addr)
		if err != nil {
			return connectedPeers, err
		}
		pi, ok := peerInfos[pii.ID]
		if !ok {
			pi = &peerstore.PeerInfo{ID: pii.ID}
			peerInfos[pi.ID] = pi
		}
		pi.Addrs = append(pi.Addrs, pii.Addrs...)
	}

	wg.Add(len(peerInfos))
	for _, peerInfo := range peerInfos {
		go func(peerInfo *peerstore.PeerInfo) {
			defer wg.Done()
			err := ipfs.Swarm().Connect(ctx, *peerInfo)
			if err != nil {
				//log.Printf("failed to connect to %s: %s", peerInfo.ID, err)
			} else {
				connectedPeers++
				//fmt.Println("ethoFS Node connection successful!")
			}
		}(peerInfo)
	}
	wg.Wait()
	return connectedPeers, nil
}

// verifyUpload will retreive uplode validaiton using the findProvs function with go-ipfs
func verifyUpload(ctx context.Context, ipfs icore.CoreAPI, cid path.Path) bool {
	providerCount := int64(0)
	dhtApi := ipfs.Dht()
	out, err := dhtApi.FindProviders(ctx, cid, options.Dht.NumProviders(10))
	if err != nil {
		fmt.Printf("Error finding providers: %s\n", cid.String())
		return false
	}

	for {
		select {
     		case provider := <-out:
			if provider.ID.String() == "" {
				providerCount = 0
				out, err = dhtApi.FindProviders(ctx, cid, options.Dht.NumProviders(10))
				if err != nil {
					return false
				}
				time.Sleep(1 * time.Second)
			} else {
				providerCount++
				if providerCount > 3 {
					return true
				} else {
					time.Sleep(1 * time.Second)
				}
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
	return false
}

// getUnixfsFile returns ipfs file data
func getUnixfsFile(path string) (files.File, int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	st, err := file.Stat()
	if err != nil {
		return nil, 0, err
	}
	//fmt.Printf("Data upload size: %d\n", st.Size())

	f, err := files.NewReaderPathFile(path, file, st)
	if err != nil {
		return nil, 0, err
	}

	return f, st.Size(), nil
}

// getUnixfsNode returns full ipfs node/data
func getUnixfsNode(path string) (files.Node, int64, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, 0, err
	}
	//fmt.Printf("Data upload size: %d\n", st.Size())

	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return nil, 0, err
	}

	return f, st.Size(), nil
}
