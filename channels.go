package main

import (
	"encoding/json"
	"net/http"
	"net/smtp"
	"strings"
	"time"
)

// EmailConfig contains email data
type EmailConfig struct {
	username  string
	password  string
	server    string
	port      string
	recipient string
	subject   string
	message   string
}

// Telegram Send messages via telegram
func Telegram(chatID string, apiKey string, message string) (*http.Response, error) {

	jayson := map[string]interface{}{
		"chat_id": chatID,
		"text":    message,
	}
	js, _ := json.Marshal(jayson)
	endpoint := "https://api.telegram.org/bot" + apiKey + "/sendMessage"

	return request(endpoint, string(js))
}

// Slack Send messages via Slack
func Slack(message string, webhook string) (*http.Response, error) {
	js := `{"text":"` + message + `"}`
	return request(webhook, js)
}

// Email Send messages via email
func Email(email EmailConfig) error {
	// Set up authentication information.
	auth := smtp.PlainAuth("", email.username, email.password, email.server)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{email.recipient}
	msg := []byte("To: " + email.recipient + "\r\n" +
		"Subject: " + email.subject + "\r\n" +
		"\r\n" +
		email.message + "\r\n")
	err := smtp.SendMail(email.server+":"+email.port, auth, email.username, to, msg)
	if err != nil {
		return err
	}

	return nil
}

func request(endpoint string, data string) (*http.Response, error) {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	var resp *http.Response
	var err error

	client := &http.Client{Transport: tr}

	resp, err = client.Post(endpoint, "application/json", strings.NewReader(data))

	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	return resp, nil
}
