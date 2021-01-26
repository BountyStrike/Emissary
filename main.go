package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"
)

func checkResponse(httpResponse *http.Response, err error) {
	if httpResponse != nil {
		if httpResponse.StatusCode >= 400 {
			body, respErr := ioutil.ReadAll(httpResponse.Body)
			if respErr != nil {
				log.Println("Error reading response body:", respErr)
			}
			log.Println("Response HTTP Status code: ", httpResponse.StatusCode)
			log.Println("Response HTTP Body: ", string(body))
		}
	}

	if err != nil {
		log.Printf("Something went wrong sending your message: %s\n", err)
	}
}

func main() {
	opts := processArgs()

	if opts.version {
		fmt.Printf("Emissary version: %s\n", VERSION)
		os.Exit(0)
	}
	if len(os.Args) <= 1 {
		flag.Usage()
		os.Exit(0)
	}

	User, err := user.Current()

	if err != nil {
		log.Fatal("Something went wrong trying to figure out your home directory", err)
	}

	configPath := filepath.FromSlash(User.HomeDir + "/.config/emissary.ini")

	cfg, err := ini.Load(configPath)
	if err != nil {
		log.Fatal("Fail to read configuration file: ", err)
	}

	if len(opts.message) > 0 && opts.stdin {
		fmt.Println("[-] Please take input from either STDIN or cli argument, not both.")
		os.Exit(1)
	}

	if opts.message == "" && !opts.stdin {
		fmt.Println("[-] You forgot to set message...")
		os.Exit(1)
	}

	if len(opts.inline.hooks) > 0 {

		for _, val := range opts.inline.hooks {
			if val.webhook == "" {
				log.Fatal("[-] Inline webhook does not contain webhook...")
			}

			if val.textField == "" {
				val.textField = "text"
			}

			resp, err := WebhookRequest(val.webhook, opts.message, val.textField, val.data)

			checkResponse(resp, err)
		}

	}

	messages := []string{"Data from Emissary\n--------"}

	if opts.stdin {
		count := 0
		sc := bufio.NewScanner(os.Stdin)
		msg := ""
		for sc.Scan() {
			msg = sc.Text()
			if opts.rows == 0 {
				messages = append(messages, msg)
			} else {
				if count < opts.rows {
					messages = append(messages, msg)
				} else {
					break
				}
			}
			count++
		}

		messages = append(messages, fmt.Sprintf("--------\nSent %d lines", count))
		fmt.Println(messages)

		opts.message = strings.Join(messages[:], "\n")
	}

	if len(opts.channel) != 0 {
		for _, ch := range opts.channel {
			webhook := cfg.Section(ch).Key("webhook").String()
			textField := cfg.Section(ch).Key("textField").MustString("text")
			additionalData := cfg.Section(ch).Key("data").String()

			if webhook == "" {
				fmt.Printf("[-] Channel %s does not contain a webhook...", ch)
				os.Exit(1)
			}

			// MS Teams hack for properly showing rows
			if strings.HasPrefix(webhook, "https://outlook.office.com") {
				split := strings.Split(opts.message, "\n")
				newMessage := ""
				for _, v := range split {
					newMessage += v + "\n\n"
				}
				opts.message = newMessage
			}

			resp, err := WebhookRequest(webhook, opts.message, textField, additionalData)

			checkResponse(resp, err)

		}
	}

	if opts.email {
		emailUsername := cfg.Section("Email").Key("username").String()
		emailPassword := cfg.Section("Email").Key("password").String()
		emailRecipient := cfg.Section("Email").Key("recipient").String()
		emailPort := cfg.Section("Email").Key("port").String()
		emailServer := cfg.Section("Email").Key("server").String()
		emailSubject := cfg.Section("Email").Key("subject").String()

		emailConfig := EmailConfig{username: emailUsername, password: emailPassword,
			recipient: emailRecipient, port: emailPort, server: emailServer, subject: emailSubject,
			message: opts.message}

		checkResponse(nil, Email(emailConfig))

	}

}
