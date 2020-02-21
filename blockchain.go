package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

var rpcLocation ="https://rpc.ether1.org"

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

/*func GetPinContractValues(ContractPinTrackingMap map[string][]string) map[string][]string {
	InternalContractPinTrackingMap := make(map[string][]string)
	c, _ := ethclient.Dial(rpcLocation)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if c != nil && ctx != nil {
		address := common.HexToAddress("0xD3b80c611999D46895109d75322494F7A49D742F")
		contract, err := NewPinStorage(address, c)
		if err != nil {
			log.Println("Ether-1 Network Connection Error: ", err)
		} else {
			ReplicationFactor, err := contract.ReplicationFactor(nil)
			if err != nil {
				log.Println("Ether-1 Network Connection Error: (ReplicationFactor) ", err)
			} else {
				repFactor = int(ReplicationFactor)
			}
			PinCount, err := contract.PinCount(nil)
			if err != nil {
				log.Println("Ether-1 Network Connection Error: (PinCount) ", err)
			} else {
				var ContractPinArray []string
				for j := uint64(0); j < uint64(PinCount); j++ {
					i := new(big.Int).SetUint64(j)
					ContractPin, err := contract.Pins(nil, i)
					if err != nil {
						log.Println("Ether-1 Network Connection Error: (ContractPin) ", err)
					} else {
						if _, ok := ContractPinTrackingMap[ContractPin]; ok {
							var tempPinArray []string
							for l := uint64(0); uint64(len(ContractPinTrackingMap[ContractPin])) > l; l++ {
								ContractPinArray = append(ContractPinArray, ContractPinTrackingMap[ContractPin][l])
								tempPinArray = append(tempPinArray, ContractPinTrackingMap[ContractPin][l])
							}
							InternalContractPinTrackingMap[ContractPin] = tempPinArray
						} else {
							ContractPinArray = append(ContractPinArray, ContractPin)
							//NOW CAT FOR SERIALISED PIN LIST IN IPFS
							catsh := shell.NewShell("localhost:" + apiPort)
							catsh.SetTimeout(1 * time.Second)
							resp, err := catsh.Cat(ContractPin)
							if err != nil {
							} else {
								buf := new(bytes.Buffer)
								buf.ReadFrom(resp)
								catString := buf.String()
								IPFSPinListArray := strings.Split(catString, ":")
								var tempPinArray []string
								if len(IPFSPinListArray) > 1 {
									if IPFSPinListArray[0] == mainChannelString {
										for k := uint64(1); k < uint64(len(IPFSPinListArray)); k++ {
											tempPinArray = append(tempPinArray, IPFSPinListArray[k])
											ContractPinArray = append(ContractPinArray, IPFSPinListArray[k])
										}
									}
								}
								InternalContractPinTrackingMap[ContractPin] = tempPinArray
							}
						}
					}
				}
				MasterPinArray = ContractPinArray
			}
		}
		c.Close()
	}
	return InternalContractPinTrackingMap
}*/
