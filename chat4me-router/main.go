package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path"
)

var (
	cfg    c4mrConfig
	server c4mrServer
)

func initConfig() {
	cfgFile, err := os.Open("config.json")
	cfgFilePath := path.Join(cfg.workingDir, "config.json")
	if err != nil {
		log.Fatalf("Error opening config file %s: %s", cfgFilePath, err.Error())
	}
	defer cfgFile.Close()
	ba, err := io.ReadAll(cfgFile)
	if err != nil {
		log.Fatalf("Error reading config file %s: %s", cfgFilePath, err.Error())
	}
	if err = json.Unmarshal(ba, &cfg); err != nil {
		log.Fatalf("Error parsing config file %s: %s", cfgFilePath, err.Error())
	}
}

func main() {
	log.Println("Initializing Chat 4 Me request router")
	cfg = c4mrConfig{}
	var err error
	if cfg.workingDir, err = os.Getwd(); err != nil {
		log.Fatalln("Error getting working directory:", err.Error())
	}
	initConfig()
}
