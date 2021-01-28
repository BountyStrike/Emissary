package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPreparePayload(t *testing.T) {

	message := "msg"
	textField := "content"
	additionalData := "{\"someKey\":\"someValue\",\"someOtherKey\":\"someOtherValue\"}"
	json, err := PreparePayload(message, textField, additionalData)
	if err != nil {
		t.Fatal("PreparePayload returned error: ", err)
	}
	strJSON := string(json)
	expectedValue := `{"content":"msg","someKey":"someValue","someOtherKey":"someOtherValue"}`
	if strJSON != expectedValue {
		t.Errorf("Expected %s, got %s", expectedValue, strJSON)
	}
}

func TestSendRequest(t *testing.T) {
	// 1. start web server

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	webhook := ts.URL
	json := []byte(`{"content":"msg","someKey":"someValue","someOtherKey":"someOtherValue"}`)
	resp, err := SendRequest(webhook, json)

	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	if err != nil {
		t.Fatal("SendRequest returned error: ", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", 200, resp.StatusCode)
	}

}
