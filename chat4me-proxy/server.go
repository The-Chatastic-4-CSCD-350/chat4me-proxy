package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/sashabaranov/go-openai"
	"github.com/uptrace/bunrouter"
)

func initServer(startedSignal chan bool) {
	server.addr = "127.0.0.1:8888"
	server.signal = make(chan os.Signal, 1)
	server.router = bunrouter.New(bunrouter.WithNotFoundHandler(server.serveNotFoundJSON))
	server.router.GET("/c4m/completion", server.serveCompletion)
	server.router.POST("/c4m/completion/", server.serveCompletion)
	server.router.POST("/c4m/completion", server.serveCompletion)

	signal.Notify(server.signal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	go func() {
		if startedSignal != nil {
			startedSignal <- true
		}
		err := http.ListenAndServe(server.addr, server.router)
		if err != nil {
			logger.Fatalf("Error listening on %s: %s", server.addr, err.Error())
		}
	}()
	<-server.signal
}

type c4mrServer struct {
	addr   string
	signal chan os.Signal
	router *bunrouter.Router
}

func (s *c4mrServer) getRealIP(request *http.Request) string {
	ip := request.Header.Get("X-Real-Ip")
	if ip != "" {
		return ip
	}
	return request.RemoteAddr
}

func (s *c4mrServer) serveNotFoundJSON(writer http.ResponseWriter, request bunrouter.Request) error {
	s.logAccess(http.StatusNotFound, request.Request)
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/112.0")
	writer.WriteHeader(http.StatusNotFound)
	return json.NewEncoder(writer).Encode(ErrUnrecognizedRequest)
}

func (s *c4mrServer) serveBadRequestError(writer http.ResponseWriter, request *http.Request, errorMessage string, extraLogging ...string) error {
	s.logAccess(http.StatusBadRequest, request, extraLogging...)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(writer).Encode(errorJSON{Status: "error", Message: errorMessage})
}

func (s *c4mrServer) logAccess(httpStatus int, request *http.Request, extra ...string) {
	logger.Printf("%s (%d %s) %q from %s, UA: %q %s",
		request.Method, httpStatus, http.StatusText(httpStatus),
		request.RequestURI, s.getRealIP(request), request.UserAgent(), strings.Join(extra, " "),
	)
}

func (s *c4mrServer) serveCompletion(writer http.ResponseWriter, request bunrouter.Request) error {
	writer.Header().Set("Content-Type", "application/json")
	if request.Header.Get("X-C4m") != "y" {
		return s.serveBadRequestError(writer, request.Request, http.StatusText(http.StatusBadRequest), "missing X-C4m header")
	}
	if oaiClient == nil {
		initOpenAI()
	}
	messages := request.PostFormValue("messages")
	if messages == "" {
		return s.serveBadRequestError(writer, request.Request, http.StatusText(http.StatusBadRequest), "missing messages form")
	}
	response, err := doCompletion(messages)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		encoder := json.NewEncoder(writer)
		if errors.Is(err, openai.ErrCompletionUnsupportedModel) {
			encoder.Encode(errorJSON{Status: "error", Message: "unsupported model used with completion request in proxy server configuration"})
		} else {
			encoder.Encode(errorJSON{Status: "error", Message: err.Error()})
		}
		s.logAccess(http.StatusInternalServerError, request.Request, "Error in doCompletion():", err.Error())
		return err
	}
	/*
	 * split the string by newline and only keep the first, for some reason it returns
	 * You:<expected message>
	 * Friend:<this shouldn't be here but is>
	 */
	completionText := strings.Split(strings.TrimSpace(response.Choices[0].Text), "\n")[0]

	json.NewEncoder(writer).Encode(completionText)
	s.logAccess(http.StatusOK, request.Request, "Completion request")
	return nil
}

func (s *c4mrServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.logAccess(http.StatusOK, request)
	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode("Hello from the request proxy!")
}
