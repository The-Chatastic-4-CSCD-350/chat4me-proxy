package main

import (
	"log"
	"net/http"
)

func initServer() {
	server.addr = "127.0.0.1:8888"
	err := http.ListenAndServe(server.addr, &server)
	if err != nil {
		log.Fatalf("Error listening on %s: %s", server.addr, err.Error())
	}
}

type c4mrServer struct {
	addr string
}

func (s *c4mrServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Hello from the router!"))
}
