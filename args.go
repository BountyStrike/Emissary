package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type channels []string

func (c *channels) String() string {
	return ""
}

func (c *channels) Set(value string) error {
	*c = append(*c, value)
	return nil
}

type inlines struct {
	hooks []inline
}

type inline struct {
	webhook   string
	textField string
	data      string
}

func (i *inlines) String() string {
	return "{'lol':1}"
}

func (i *inlines) Set(value string) error {
	split := strings.Split(value, "§")
	mul := inlines{}
	final := inline{}
	for _, val := range split {
		s := strings.Split(val, ":=")

		if strings.ToLower(s[0]) == "webhook" {
			final.webhook = s[1]
		}

		if strings.ToLower(s[0]) == "textField" {
			final.textField = s[1]
		}

		if strings.ToLower(s[0]) == "data" {
			final.data = s[1]
		}

	}
	mul.hooks = append(mul.hooks, final)
	*i = mul

	return nil
}

type cliOptions struct {
	email   bool
	version bool
	stdin   bool
	message string
	channel channels
	inline  inlines
	data    string
	text    string
	rows    int
}

func processArgs() cliOptions {

	opts := cliOptions{}
	flag.Var(&opts.channel, "channel", "Specify a custom channel you have defined in ~/.config/emissary.ini")
	flag.Var(&opts.channel, "ch", "Specify a custom channel you have defined in ~/.config/emissary.ini")
	flag.Var(&opts.inline, "inline", "Specify channel directly in the command line")
	flag.Var(&opts.inline, "in", "Specify channel directly in the command line")
	flag.StringVar(&opts.text, "text", "", "Specify the field that contains the message. Default is 'text'")
	flag.StringVar(&opts.text, "txt", "", "Specify the field that contains the message. Default is 'text'")
	flag.StringVar(&opts.data, "data", "", "Specify json data that should be sent")
	flag.StringVar(&opts.data, "d", "", "Specify json data that should be sent")
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
		h += "  -s,   --slack        Send via Slack\n"
		h += "  -t,   --telegram     Send via Telegram\n"
		h += "  -e,   --email        Send via Email\n"
		h += "  -ms,  --teams        Send via Microsoft Teams\n"
		h += "  -si,  --stdin        Get message from stdin\n"
		h += "  -m,   --message      Message to send\n"
		h += "  -ch,  --channel      Specify a custom channel you have defined emissary.ini\n"
		h += "  -in,  --inline       Specify channel directly in the commandline\n"
		h += "  -txt, --text         Specify the field that contains the message. Default is 'message'\n"
		h += "  -d,   --data         Specify additional data in json format that should be sent\n"
		h += "  -r,   --rows         Max rows/lines to send, 0 for unlimited. Default 20\n"
		h += "  -v,   --version      Show version\n"

		h += "\nExamples:\n"
		h += "  emissary -telegram --message \"Hello telegram\"\n"
		h += "  cat domins.txt | emissary --slack --stdin \n\n"

		fmt.Fprintf(os.Stderr, h)
	}
}
