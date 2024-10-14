package main

import (
	"io"
	"net/http"
	"testing"
	"time"
)

func SendRequest(url string) (*http.Response, error) {
	response, err := http.Get(url)
	return response, err
}
func TestServer(t *testing.T) {
	go RunServer(":8081")
	time.Sleep(time.Second * 2)
	resp, err := SendRequest("http://localhost:8081/")
	if err != nil {
		t.Errorf("server doesn't respond: %s", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("uncorrect status code: %d, expected: 200", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("error reading content")
	}

	if string(body) != "<body>testpage</body>" {
		t.Errorf("unexpected content, got: %s, expected: <body>testpage</body>", string(body))
	}

	cookies := resp.Cookies()
}
