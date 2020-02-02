package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

func checkResponse(httpResponse *http.Response, err error) {
	if httpResponse != nil {
		if httpResponse.StatusCode > 201 {
			body, respErr := ioutil.ReadAll(httpResponse.Body)
			if respErr != nil {
				log.Println("Error reading response body")
				os.Exit(1)
			}
			log.Println("HTTP Status code: ", httpResponse.StatusCode)
			log.Println("HTTP Body: ", string(body))
		}
	}

	if err != nil {
		log.Println("Something went wrong sending your message: ", err)
		os.Exit(1)
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

	cfg, err := ini.Load("emissary.ini")
	if err != nil {
		log.Fatal("Fail to read configuration file: ", err)
		os.Exit(1)
	}

	if len(opts.message) > 0 && opts.stdin {
		fmt.Println("[-] Please take input from either STDIN or cli argument, not both.")
		os.Exit(1)
	}

	if opts.message == "" && !opts.stdin {
		fmt.Println("[-] You forgot to set message...")
		os.Exit(1)
	}

	messages := make([]string, 21)

	if opts.stdin {
		count := 0
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			msg := sc.Text()
			if msg != "" {
				if count < 20 {
					messages[count] = msg
					count++
				} else {
					messages[count] = "Sent 20 domains, there are more on the server."
					break
				}

			} else {
				break
			}
		}

		opts.message = strings.Join(messages[:], "\n")
	}

	if opts.telegram {
		telegramAPIKey := cfg.Section("Telegram").Key("api_key").String()
		telegramUser := cfg.Section("Telegram").Key("chat_id").String()

		resp, err := Telegram(telegramUser, telegramAPIKey, opts.message)

		checkResponse(resp, err)
	}

	if opts.slack {
		slackWebhook := cfg.Section("Slack").Key("webhook").String()

		resp, err := Slack(opts.message, slackWebhook)

		checkResponse(resp, err)
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
