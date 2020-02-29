// Copyright 2020 The Etho.Black Development Team

package main

import (
        "errors"
        "net"
        "github.com/ethereum/go-ethereum/rpc"
)

type EthofsService struct {}

func (s *EthofsService) Register(name, key string) string {
        return "Registered " + name
}

func (s *EthofsService) Add(path string, key string, blocks uint64) (string, error) {
        if blocks == 0 {
                return "ethofs Data Upload Failed", errors.New("block count is too low")
        }
        return "ethoFS Data Upload Complete", nil
}

func ethofsIpc() {
        ethofs := new(EthofsService)
        server := rpc.NewServer()
        server.RegisterName("ethofs", ethofs)
        l, _ := net.ListenUnix("unix", &net.UnixAddr{Net: "unix", Name: "ethofs.ipc"})
        server.ServeListener(l)
}
