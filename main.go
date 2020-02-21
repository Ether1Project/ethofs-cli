// Copyright 2020 The Etho.Black Development Team

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Work in progress
	fmt.Println("Initializing ethoFS Node For Upload ")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Spawn a node using a temporary path, creating a temporary repo for the run
	fmt.Println("Spawning node on a temporary repo")
	//ipfs, err := spawnEphemeral(ctx)
	_, err := spawnEphemeral(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
	}

	fmt.Println("ethoFS node is running")
	time.Sleep(5 * time.Second)
	fmt.Println("Stopping ethoFS node")
	time.Sleep(1 * time.Second)
	fmt.Println("Exiting..")
}
