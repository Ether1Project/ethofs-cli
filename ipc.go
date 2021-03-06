// Copyright 2020 The Etho.Black Development Team

package main

import (
	"context"
	"encoding/json"
        "errors"
	"fmt"
        "net"
	"strings"

        "github.com/ethereum/go-ethereum/rpc"
        "github.com/ethereum/go-ethereum/common"
)

type EthofsService struct {}

// Register sets up new hosting user on etho chain - ipc access
func (s *EthofsService) Register(name, key string) (string, error) {
	_, err := RegisterAccount(key, name)
	if err != nil {
	        return "Registration Failed", err
	}
        return "Successfully Registered " + name, nil
}

// Remove allows for ethoFS hosting contract removal - ipc access
func (s *EthofsService) Remove(address string, key string) (string, error) {
	_, err := RemoveContract(key, common.HexToAddress(address))
	if err != nil {
	        return "Contract Removal Failed", err
	}
        return "Contract Successfully Removed " + address, nil
}

// Extend allows ethoFS contract extension - ipc access
func (s *EthofsService) Extend(name string, duration int32, key string) (string, error) {

	contractDetails, err := GetContractDetails(key, contractName)
	if err != nil {
	        return "Contract Extension Failed", err
	}

	extensionCost, err := CalculateUploadCost(duration, int64(contractDetails.Size))
	if err != nil {
	        return "Contract Extension Failed", err
	}

	_, err = ExtendContract(key, extensionCost, contractDetails.Address, duration)
	if err != nil {
	        return "Contract Extension Failed", err
	}
        return "Contract Successfully Extended: " + name, nil
}

// List returns list of currently active ethoFS hosting contracts - ipc access
func (s *EthofsService) List(key string) (string, error) {
	contractData, err := ListExistingContracts(key)
	if err != nil {
		return "", err
	}

	d, err := json.Marshal(contractData)
	if err != nil {
		return "", err
	}

	return string(d), nil
}

// Add allows for ethoFS hosting contract addition - ipc access
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
		contractCost, err := CalculateUploadCost(int32(blocks), uploadSize)
		if err != nil {
	                return "ethofs Data Upload Failed", errors.New("error uploading data")
		}

		contentHashString := "ethoFSPinningChannel_alpha11:" + hash
		contentPathString := "ethoFSPinningChannel_alpha11:/"
		_, err = UploadData(key, contractCost, hash, name, uint32(blocks), uint32(uploadSize), contentHashString, contentPathString)
		if err != nil {
	                return "ethofs Data Upload Failed", errors.New("error uploading data")
		}

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
		contractCost, err := CalculateUploadCost(int32(blocks), uploadSize)
		if err != nil {
	                return "ethofs Data Upload Failed", errors.New("error uploading data")
		}

		contentHashString := "ethoFSPinningChannel_alpha11:" + hash
		contentPathString := "ethoFSPinningChannel_alpha11:/"
		_, err = UploadData(key, contractCost, hash, name, uint32(blocks), uint32(uploadSize), contentHashString, contentPathString)
		if err != nil {
	                return "ethofs Data Upload Failed", errors.New("error uploading data")
		}

		if verifyUpload(ctx, ipfs, cidFile) {
			return cidFile.String(), nil
		} else {
	                return cidFile.String(), errors.New("ethoFS upload failed")
		}
	}
	return "ethofs Data Upload Failed", errors.New("params incorrect")
}

// ethofsIpc sets up new ipc end point
func ethofsIpc() {
        ethofs := new(EthofsService)
        server := rpc.NewServer()
        server.RegisterName("ethofs", ethofs)
        l, _ := net.ListenUnix("unix", &net.UnixAddr{Net: "unix", Name: "ethofs.ipc"})
        server.ServeListener(l)
}
