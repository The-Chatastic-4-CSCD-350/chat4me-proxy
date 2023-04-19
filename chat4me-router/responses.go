package main

var (
	ErrUnrecognizedRequest     = errorJSON{Status: "error", Message: "unrecognized request"}
	ErrOAIClientNotInitialized = errorJSON{Status: "error", Message: "unable to send OpenAI request: client not initialized"}
)

type errorJSON struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (e *errorJSON) Error() string {
	return e.Message
}
