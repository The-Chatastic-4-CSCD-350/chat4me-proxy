package main

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func initTestServer() {
	if server.signal == nil {
		waitForServer := make(chan bool)
		go initServer(waitForServer)
		<-waitForServer
	}
}

func TestServerRequestOK(t *testing.T) {
	initTestServer()
	t.Log("Server started, continuing with test")
	req, err := http.NewRequest("GET", "http://127.0.0.1:8888/c4m/completion", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Got status code %s", resp.Status)
	}
	ba, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("Got response: %q", string(ba))
	server.signal <- os.Interrupt
}

func TestRealIP(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8888", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	expectedRealIP := "192.168.56.1:1234"
	req.RemoteAddr = "127.0.0.1:1234"
	req.Header.Set("X-Real-IP", expectedRealIP)
	realIP := server.getRealIP(req)
	if realIP != expectedRealIP {
		t.Fatalf(
			"server.getRealIP(req) not using the header field X-Real-Ip (expected %q, got %q)",
			expectedRealIP, realIP)
	}
}
