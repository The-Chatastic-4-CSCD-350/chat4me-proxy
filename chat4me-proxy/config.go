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
	logWriter        io.Writer
	logger           *log.Logger
)

type c4mrConfig struct {
	APIKey         string `json:"apiKey"`
	OrganizationID string `json:"organizationID"`
	LogDir         string `json:"logDir"`
	Verbose        bool   `json:"verbose"`

	workingDir string
	logFile    *os.File
}

func (c *c4mrConfig) validate() error {
	if c.APIKey == "" {
		return ErrMissingAPIKey
	}
	if c.LogDir == "" {
		c.LogDir = "."
	}
	return nil
}

func (c *c4mrConfig) close() {
	if c.logFile != nil {
		c.logFile.Close()
	}
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
	logPath := path.Join(cfg.LogDir, "c4m-proxy.log")
	if cfg.logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0640); err != nil {
		log.Fatalf("Error opening log file %s: %s", logPath, err.Error())
	}
	logWriter = cfg.logFile
	if cfg.Verbose {
		logWriter = io.MultiWriter(cfg.logFile, os.Stdout)
	}
	logger = log.New(logWriter, "[c4m]", log.Flags())
}
