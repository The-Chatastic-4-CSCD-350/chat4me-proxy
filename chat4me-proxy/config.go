package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"path"
)

var (
	ErrMissingAPIKey = errors.New("missing API key from config file")
)

type c4mrConfig struct {
	APIKey         string `json:"apiKey"`
	OrganizationID string `json:"organizationID"`

	workingDir string
}

func (c *c4mrConfig) validate() error {
	if c.APIKey == "" {
		return ErrMissingAPIKey
	}
	return nil
}

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

	if err = cfg.validate(); err != nil {
		log.Fatalf("Error validating config file %s: %s", cfgFilePath, err.Error())
	}
}
