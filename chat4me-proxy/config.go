package main

import "errors"

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
