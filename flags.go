package main

import (
	"flag"
)

var uploadFlag bool
var listFlag bool
var recursiveFlag bool
var inputPath string
var privateKey string
var contractName string
var contractDuration int32
var removeFlag bool
var extendFlag bool

func setFlags() {
	flag.BoolVar(&uploadFlag, "upload", false, "Upload to ethoFS")
	flag.BoolVar(&removeFlag, "remove", false, "Remove ethofs Hosting Contract")
	flag.BoolVar(&removeFlag, "extend", false, "Extend ethofs Hosting Contract")
	flag.BoolVar(&listFlag, "list", false, "List ethoFS Upload Contracts")
	flag.BoolVar(&recursiveFlag, "r", false, "Recursive Upload")
	flag.StringVar(&inputPath, "path", "", "Data Upload Path")
	flag.StringVar(&privateKey, "key", "", "Private Key")
	flag.StringVar(&contractName, "name", "", "Hosting Contract Name")

	duration := uint(0)
	flag.UintVar(&duration, "blocks", 0, "Hosting Contract Duration (In Blocks)")

	flag.Parse()

	contractDuration = int32(duration)
}
