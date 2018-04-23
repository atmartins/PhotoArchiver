package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atmartins/PhotoArchiver/archive"
	"github.com/atmartins/PhotoArchiver/config"
	"github.com/atmartins/PhotoArchiver/db"
)

var constants config.Constants
var constantsFile = "./constants.json"
var flagNameSrc = "src"
var flagNameDest = "dest"
var emptyString = ""

func main() {
	constants = config.LoadConstants(constantsFile)

	srcPtr := flag.String(flagNameSrc, emptyString, "Source directory to read photos FROM")
	destPtr := flag.String(flagNameDest, emptyString, "Destination directory to archive photos TO")

	flag.Parse()

	if *srcPtr == emptyString || *destPtr == emptyString {
		flag.Usage()
		os.Exit(1)
	}

	err := db.Connect(constants.DbAddr, constants.DbName)
	if err != nil {
		panic(err)
	}

	stats, err := archive.All(*srcPtr, *destPtr)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("stats %+v", stats))

	db.Disconnect()

}
