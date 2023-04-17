package main

type c4mrConfig struct {
	APIKey         string `json:"apiKey"`
	OrganizationID string `json:"organizationID"`

	workingDir string
}
