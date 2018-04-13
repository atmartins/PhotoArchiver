package main

import (
	"github.com/atmartins/PhotoArchiver/archive"
	"github.com/atmartins/PhotoArchiver/classify"
	"github.com/atmartins/PhotoArchiver/config"
)

var constants config.Constants

func main() {
	constants = config.LoadConstants("./constants.json")

	classify.All(constants)
	archive.All(constants)
}
