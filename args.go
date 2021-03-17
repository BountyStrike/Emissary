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
	header  string
	rows    int
}

func processArgs() cliOptions {

	opts := cliOptions{}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.Var(&opts.channel, "channel", "Specify a custom channel you have defined in ~/.config/emissary.ini")
	flag.Var(&opts.channel, "ch", "Specify a custom channel you have defined in ~/.config/emissary.ini")
	flag.Var(&opts.inline, "inline", "Specify channel directly in the command line")
	flag.Var(&opts.inline, "in", "Specify channel directly in the command line")
	flag.BoolVar(&opts.email, "email", false, "Send via smtp")
	flag.BoolVar(&opts.email, "e", false, "Send via smtp")
	flag.StringVar(&opts.header, "header", "Data from Emissary", "Specify custom header to send. Default 'Data from Emissary'")
	flag.StringVar(&opts.header, "h", "Data from Emissary", "Specify custom header to send. Default 'Data from Emissary'")
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
		h += "  emissary [options] [message]\n\n"

		h += "Options:\n"
		h += "  -ch,  --channel      Specify a custom channel you have defined emissary.ini\n"
		h += "  -in,  --inline       Specify channel directly in the commandline\n"
		h += "  -m,   --message      Message to send\n"
		h += "  -h,   --header       Custom header\n"
		h += "  -si,  --stdin        Get message from stdin\n"
		h += "  -e,   --email        Send via Email\n"
		h += "  -txt, --text         Specify the field that contains the message. Default is 'message'\n"
		h += "  -d,   --data         Specify additional data in json format that should be sent\n"
		h += "  -r,   --rows         Max rows/lines to send, 0 for unlimited. Default 20\n"
		h += "  -v,   --version      Show version\n"

		h += "\nExamples:\n"
		h += "  emissary --channel Telegram --message \"Hello telegram\"\n"
		h += "  cat domins.txt | emissary -ch Slack --stdin --header \"New subdomains from Google!\"\n"
		h += "  emissary -ch Discord -ch Telegram -m \"Your message\" \n"
		h += "  emissary -in \"webhook:=https://api.telegram.org/botxxxxx/sendMessage§data:={'chat_id': 'xxxx'}\" -in \"webhook:=https://hooks.slack.com/services/xxxxx\" -m \"Hack the planet!\" \n"

		fmt.Fprintf(os.Stderr, h)
	}
}
