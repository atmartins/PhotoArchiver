package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*Constants runtime configuration values */
type Constants struct {
	DirNamePhotos          string `json:"DirNamePhotos"`
	DirNameVideos          string `json:"DirNameVideos"`
	ClassifyOnly           bool   `json:"ClassifyOnly"`
	MaxFileSizeBytesToTest int64  `json:"MaxFileSizeBytesToTest"`
	DbAddr                 string `json:"DbAddr"`
	DbName                 string `json:"DbName"`
}

/*LoadConstants load constants from json file*/
func LoadConstants(jsonFilePath string) Constants {
	raw, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c Constants
	json.Unmarshal(raw, &c)
	return c
}
