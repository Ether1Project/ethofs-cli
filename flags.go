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
var replaceFlag bool
var backupFlag bool
var registerFlag bool
var ethofsUsername string
var ipcFlag bool

func setFlags() {
	flag.BoolVar(&ipcFlag, "ipc", false, "Launch IPC Endpoint - Abilty To Respond to Json Requests")
	flag.BoolVar(&uploadFlag, "upload", false, "Upload to ethoFS")
	flag.BoolVar(&removeFlag, "remove", false, "Remove ethofs Hosting Contract")
	flag.BoolVar(&removeFlag, "extend", false, "Extend ethofs Hosting Contract")
	flag.BoolVar(&replaceFlag, "replace", false, "Replace ethofs Hosting Contract (Will Automatically Replace Same Named Contract if Existing)")
	flag.BoolVar(&backupFlag, "backup", false, "Backup Data to ethofs Network (Will Automatically Replace Same Named Contract if Existing)")
	flag.BoolVar(&listFlag, "list", false, "List ethoFS Upload Contracts")
	flag.BoolVar(&registerFlag, "register", false, "Register New ethoFS Address/Account")
	flag.BoolVar(&recursiveFlag, "r", false, "Recursive Upload")
	flag.StringVar(&inputPath, "path", "", "Data Upload Path")
	flag.StringVar(&privateKey, "key", "", "Private Key")
	flag.StringVar(&contractName, "name", "", "Hosting Contract Name")
	flag.StringVar(&ethofsUsername, "username", "", "ethoFS Registration Username")

	duration := uint(0)
	flag.UintVar(&duration, "blocks", 0, "Hosting Contract Duration (In Blocks)")

	flag.Parse()

	contractDuration = int32(duration)
}
