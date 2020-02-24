package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

func UploadData(mainHash string, contractName string, contractDuration uint64, uploadSize uint64, contentHashString string, contentPathString string) {
	c, _ := ethclient.Dial(rpcLocation)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if c != nil && ctx != nil {
		address := common.HexToAddress("")
		contract, err := NewController(address, c)
		if err != nil {
			log.Println("Ether-1 Network Connection Error: ", err)
			return nil, err
		} else {
			NodeCount, err := contract.NodeID(nil)
			if err != nil {
				log.Println("Ether-1 Network Connection Error: (NodeCount) ", err)
				return nil, err
			} else {
				for j := uint32(1); j <= NodeCount; j++ {
					Bootnode, err := contract.Nodes(nil, j)
					if err != nil {
						log.Println("Ether-1 Network Connection Error: (Bootnodes) ", err)
						return nil, err
					} else {
						BootnodeArray = append(BootnodeArray, Bootnode)
					}
				}
			}
			c.Close()
		}
	}
}
