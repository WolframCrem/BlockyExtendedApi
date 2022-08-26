package config

import (
	"errors"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	os "os"
)

var Config = ConfigData{}

func LoadConfigData() ConfigData {
	// read flag
	configFile := flag.String("c", "config.yml", "Supply the yaml config file")
	flag.Parse()
	// check if file exists
	if _, err := os.Stat(*configFile); errors.Is(err, os.ErrNotExist) {
		log.Fatal(fmt.Sprintf("%s does not exist", *configFile))
	}
	var data = ConfigData{}
	// load file
	content, err := os.ReadFile(*configFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(content, &data)
	return data
}
