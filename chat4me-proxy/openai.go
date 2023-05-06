package main

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

var (
	oaiClient *openai.Client
	oaiConfig openai.ClientConfig
)

func doCompletion(promptString string) (openai.CompletionResponse, error) {
	req := openai.CompletionRequest{
		Model:            openai.GPT3Dot5Turbo,
		Prompt:           promptString + "\nYou:",
		Temperature:      0.5,
		MaxTokens:        60,
		TopP:             1.0,
		FrequencyPenalty: 1.0,
		PresencePenalty:  0.0,
		Stop:             []string{"You:"},
	}
	return oaiClient.CreateCompletion(context.Background(), req)
}

func initOpenAI() {
	oaiConfig = openai.DefaultConfig(cfg.APIKey)
	oaiConfig.OrgID = cfg.OrganizationID
	oaiClient = openai.NewClientWithConfig(oaiConfig)
}
