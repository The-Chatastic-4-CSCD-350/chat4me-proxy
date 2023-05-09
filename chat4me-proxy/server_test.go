package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

func initTestServer() {
	if logger == nil {
		logger = log.New(os.Stderr, "[test]", log.Default().Flags())
	}
	if server.signal == nil {
		waitForServer := make(chan bool)
		go initServer(waitForServer)
		<-waitForServer
	}
}

func makeCompletionRequest(form url.Values, t *testing.T) *http.Request {
	req, err := http.NewRequest("POST", "http://127.0.0.1:8888/c4m/completion", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req
}

func TestServerHeaders(t *testing.T) {
	initTestServer()
	defer func() {
		server.signal <- os.Interrupt
	}()
	t.Log("Server started, continuing with test")

	postForm := url.Values{}
	req := makeCompletionRequest(postForm, t)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("Got status code %s, expected 400", resp.Status)
	}

	req = makeCompletionRequest(postForm, t)
	req.Header.Set("X-C4m", "y")
	if resp, err = http.DefaultClient.Do(req); err != nil {
		t.Fatal(err.Error())
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("Got status code %s, expected 400", resp.Status)
	}

	postForm = url.Values{"messages": {"You:How are you today?\nFriend:I am fine, how about you?"}}
	req = makeCompletionRequest(postForm, t)
	req.Header.Set("X-C4m", "y")
	if resp, err = http.DefaultClient.Do(req); err != nil {
		t.Fatal(err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Got status code %s, expected 200", resp.Status)
	}

	ba, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("Got response: %s", string(ba))

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
