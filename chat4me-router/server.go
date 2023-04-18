package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func initServer(startedSignal chan bool) {
	server.addr = "127.0.0.1:8888"
	server.signal = make(chan os.Signal, 1)
	signal.Notify(server.signal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	go func() {
		startedSignal <- true
		err := http.ListenAndServe(server.addr, &server)
		if err != nil {
			log.Fatalf("Error listening on %s: %s", server.addr, err.Error())
		}
	}()
	<-server.signal
}

type c4mrServer struct {
	addr   string
	signal chan os.Signal
}

func (s *c4mrServer) getRealIP(request *http.Request) string {
	ip := request.Header.Get("X-Real-Ip")
	if ip != "" {
		return ip
	}
	return request.RemoteAddr
}

func (s *c4mrServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("%s %q from %s, UA: %q",
		request.Method, request.RequestURI, s.getRealIP(request), request.UserAgent())

	writer.Header().Set("Content-Type", "application/json")
	_, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println("Error reading request body:", err.Error())
		return
	}
	encoder := json.NewEncoder(writer)
	encoder.Encode("Hello from the router!")
}
