package main

import (
	"errors"
	"testing"
)

func TestValidConfig(t *testing.T) {
	cfg := c4mrConfig{
		APIKey: "key",
	}
	if cfg.validate() != nil {
		t.Fatal("cfg.validate() returned an error despite being valid")
	}
}

func TestInvalidConfig(t *testing.T) {
	var cfg c4mrConfig
	if !errors.Is(cfg.validate(), ErrMissingAPIKey) {
		t.Fatal("cfg has an empty APIKey property but cfg.validate() r")
	}
}
