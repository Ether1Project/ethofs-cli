package main

import (
	"flag"
)

var uploadFlag bool
var recursiveFlag bool
var inputPath string

func setFlags() {
//	uploadPtr := flag.Bool("upload", false, "Upload to ethoFS")
//	recursivePtr := flag.Bool("r", false, "Recursive upload")
	flag.BoolVar(&uploadFlag, "upload", false, "Upload to ethoFS")
	flag.BoolVar(&recursiveFlag, "r", false, "Recursive upload")
	flag.StringVar(&inputPath, "path", "", "Data upload path")

//	uploadFlag = *uploadPtr
//	recursiveFlag = *recursivePtr
	flag.Parse()
}
