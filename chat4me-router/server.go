package main

import "net/http"

func initServer() {
	var s http.Handler
	http.ListenAndServe("127.0.0.1")
}

type c4mrServer struct {
}

func (s *c4mrServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}
