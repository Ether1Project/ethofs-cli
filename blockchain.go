package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

// GetLatestEther1Block fetches ether-1 block from rpc endpoint
func GetLatestEther1Block() string {
	c, _ := ethclient.Dial(rpcLocation)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if c != nil && ctx != nil {
		latestBlock, err := c.BlockByNumber(ctx, nil)
		if err != nil {
			log.Println("Ether-1 Network Connection Error: ", err)
			c.Close()
			return "0"
		}
		c.Close()
		return latestBlock.Number().String()
	}
	return "0"
}

// GetBootnodeContractValues fetches ethoFS bootnode entries from etho smart contract
func GetBootnodeContractValues() ([]string, error) {
	var BootnodeArray []string
	c, _ := ethclient.Dial(rpcLocation)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if c != nil && ctx != nil {
		address := common.HexToAddress("0xd5Cc0D9031bD7CA18A320A777A9aCCFbEdFA8709")
		contract, err := NewNodesStorage(address, c)
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
	return BootnodeArray, nil
}
