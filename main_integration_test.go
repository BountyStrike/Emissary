package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMainApp(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	webhook := "webhook:=" + ts.URL + "§data:={'chat_id': 'xxxx'}"
	os.Args = []string{"main", "-in", webhook, "-m", "hacker"}

	main()

	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", 200, res.StatusCode)
	}
}

func TestMainWithoutWebhook(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	webhook := ts.URL + "§data:={'chat_id': 'xxxx'}"
	os.Args = []string{"main", "-in", webhook, "-m", "hacker"}

	var buf bytes.Buffer
	log.SetOutput(&buf)

	main()

	fmt.Println(res.Request.Method)

	err = res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	if res.Request.Method != "" {
		t.Errorf("Expected no request to server, got a %s request", res.Request.Method)
	}

}
