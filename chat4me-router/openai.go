package main

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

var (
	oaiClient *openai.Client
	oaiConfig openai.ClientConfig
)

type completionRequest struct {
	Messages []string `json:"messages"`
	YourName string   `json:"yourName"`
}

func doCompletion(promptString string) (openai.CompletionResponse, error) {
	req := openai.CompletionRequest{
		Model:            openai.GPT3TextAda001,
		Prompt:           promptString, //"You: What have you been up to?\nFriend: Watching old movies.\nYou: Did you watch anything interesting?\nFriend:",
		Temperature:      0.5,
		MaxTokens:        30,
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
