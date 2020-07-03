package main

import (
	"flag"
	"fmt"
	"os"
)

type cliOptions struct {
	telegram bool
	discord  bool
	slack    bool
	email    bool
	teams    bool
	version  bool
	stdin    bool
	message  string
	rows     int
}

func processArgs() cliOptions {

	opts := cliOptions{}
	flag.BoolVar(&opts.telegram, "telegram", false, "Send via telegram")
	flag.BoolVar(&opts.telegram, "t", false, "Send via telegram")
	flag.BoolVar(&opts.slack, "slack", false, "Send via slack")
	flag.BoolVar(&opts.slack, "s", false, "Send via slack")
	flag.BoolVar(&opts.email, "email", false, "Send via smtp")
	flag.BoolVar(&opts.email, "e", false, "Send via smtp")
	flag.BoolVar(&opts.teams, "teams", false, "Send via Microsoft Teams")
	flag.BoolVar(&opts.teams, "ms", false, "Send via Send via Microsoft Teams")
	flag.BoolVar(&opts.version, "version", false, "Show version number")
	flag.BoolVar(&opts.version, "v", false, "Show version number")
	flag.BoolVar(&opts.stdin, "stdin", false, "Take input from stdin")
	flag.BoolVar(&opts.stdin, "si", false, "Take input from stdin")
	flag.StringVar(&opts.message, "message", "", "The message you want to send")
	flag.StringVar(&opts.message, "m", "", "The message you want to send")
	flag.IntVar(&opts.rows, "rows", 20, "Max rows/lines to send, 0 for unlimited. Default 20")
	flag.IntVar(&opts.rows, "r", 20, "Max rows/lines to send, 0 for unlimited. Default 20")
	flag.Parse()

	return opts

}

func init() {
	flag.Usage = func() {
		h := "\nSend data through chat channels. Made by @dubs3c.\n\n"

		h += "Usage:\n"
		h += "  emissary [channel] [message]\n\n"

		h += "Options:\n"
		h += "  -s,  --slack        Send via Slack\n"
		h += "  -t,  --telegram     Send via Telegram\n"
		h += "  -e,  --email        Send via Email\n"
		h += "  -ms, --teams        Send via Microsoft Teams\n"
		h += "  -si, --stdin        Get message from stdin\n"
		h += "  -m,  --message      Message to send\n"
		h += "  -r,  --rows         Max rows/lines to send, 0 for unlimited. Default 20\n"
		h += "  -v,  --version      Show version\n"

		h += "\nExamples:\n"
		h += "  emissary -telegram --message \"Hello telegram\"\n"
		h += "  cat domins.txt | emissary --slack --stdin \n\n"

		fmt.Fprintf(os.Stderr, h)
	}
}
