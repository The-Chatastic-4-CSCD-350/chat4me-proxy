package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/uptrace/bunrouter"
)

func initServer(startedSignal chan bool) {
	server.addr = "127.0.0.1:8888"
	server.signal = make(chan os.Signal, 1)
	server.router = bunrouter.New(bunrouter.WithNotFoundHandler(server.serveNotFoundJSON))
	server.router.GET("/c4m/completion", server.serveCompletion)

	signal.Notify(server.signal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	go func() {
		if startedSignal != nil {
			startedSignal <- true
		}
		err := http.ListenAndServe(server.addr, server.router)
		if err != nil {
			log.Fatalf("Error listening on %s: %s", server.addr, err.Error())
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
	writer.WriteHeader(http.StatusNotFound)
	return json.NewEncoder(writer).Encode(ErrUnrecognizedRequest)
}

func (s *c4mrServer) serveBadRequestError(writer http.ResponseWriter, request *http.Request, errorMessage string) error {
	s.logAccess(http.StatusBadRequest, request)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(writer).Encode(errorJSON{Status: "error", Message: errorMessage})
}

func (s *c4mrServer) logAccess(httpStatus int, request *http.Request, extra ...string) {
	log.Printf("%s (%d %s) %q from %s, UA: %q",
		request.Method, httpStatus, http.StatusText(httpStatus),
		request.RequestURI, s.getRealIP(request), request.UserAgent(),
	)
}

func (s *c4mrServer) serveCompletion(writer http.ResponseWriter, request bunrouter.Request) error {
	writer.Header().Set("Content-Type", "application/json")
	if request.Header.Get("X-C4m") != "y" {
		return s.serveBadRequestError(writer, request.Request, "")
	}
	if oaiClient == nil {
		writer.WriteHeader(http.StatusInternalServerError)
		s.logAccess(http.StatusInternalServerError, request.Request)
		return json.NewEncoder(writer).Encode(ErrOAIClientNotInitialized)
	}
	return nil
}

func (s *c4mrServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.logAccess(http.StatusOK, request)
	writer.Header().Set("Content-Type", "application/json")
	// var req completionRequest
	// err := json.NewDecoder(request.Body).Decode(&req)

	encoder := json.NewEncoder(writer)
	encoder.Encode("Hello from the router!")
}
