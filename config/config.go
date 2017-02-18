package config

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Configuration struct {
  DirPhotosSrc string `json:"DIR_PHOTOS_SRC"`
  DirPhotosDest string `json:"DIR_PHOTOS_DEST"`
}

func Load(configFile string) Configuration {
  raw, err := ioutil.ReadFile(configFile)
  if err != nil {
      fmt.Println(err.Error())
      os.Exit(1)
  }
  var c Configuration
  json.Unmarshal(raw, &c)
  return c
}
