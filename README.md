<h1 align="center">Emissary</h1>
<p align="center">
  <img width="300" height="300" src="./img.png">
  <p align="center">
  Send notifications via different channels such as Slack, Telegram or Teams in your bug bounty flow. 
  </p>
</p>

## Motivation
The idea is to hook Emissary into https://github.com/BountyStrike/Bountystrike-sh which will notify me on Telegram when new domains have been found.

## Usage

```
$ emissary
Send data through chat channels. Made by @dubs3c.

Usage:
  emissary [channel] [message]

Options:
  -s,   --slack        Send via Slack
  -t,   --telegram     Send via Telegram
  -e,   --email        Send via Email
  -ms,  --teams        Send via Microsoft Teams
  -si,  --stdin        Get message from stdin
  -m,   --message      Message to send
  -ch,  --channel      Specify a custom channel you have defined emissary.ini
  -in,  --inline       Specify channel directly in the command line
  -txt, --text         Specify the field that contains the message. Default is 'message'
  -d,   --data         Specify additional data in json format that should be sent
  -r,   --rows         Max rows/lines to send, 0 for unlimited. Default 20
  -v,   --version      Show version

Examples:
  emissary -telegram --message "Hello telegram"
  cat domins.txt | emissary --slack --stdin
  emissary --channel Discord -m "It works!!!"
```

**Create ~/.config/emissary.ini with the following:**
```
[Telegram]
webhook=https://api.telegram.org/botxxxxxx:xxxxxxxxxxxxxxxxxxxxxxxxxxxxx/sendMessage
textField=text
data={"chat_id": "xxxxxx"}

[Slack]
webhook=https://hooks.slack.com/services/xxxxxxxxxx/xxxxxxxxxx/xxxxxxxxxx

[Teams]
webhook=https://outlook.office.com/webhook/xxxxxxxxxxxxxxxxxxxxx

[Email]
username=
password=
recipient=
server=smtp.gmail.com
port=587
subject="New domains found!"
```
*When using gmail, you need to activate less secure apps on your account: [https://myaccount.google.com/lesssecureapps](https://myaccount.google.com/lesssecureapps)*

Now you can start using emissary :)


**Custom Webhooks**

It's possible to add your own channels as well, adding Discord as a custom channel looks like this:

```
[Discord]
webhook=https://discord.com/api/webhooks/xxxxxxxxxxxxxxxxxxxxxxxxxx
textField=content
```

And can be executed with `emissary --channel Discord -m "It works!!!"`.

The following fields can be used for a given channel:

| field     | description                                                                                                                                                |
| --------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| webhook   | The actual webhook to send data to                                                                                                                         |
| textField | Some API's have a specific json key where the message goes, here you can define that. Default key is `text`, e.g. `{"text": "Your message"}`.              |
| data      | If you want to send additional data, you can specify that here as a json formatted string, e.g. `data={"someKey": "someValue", "otherKey": "otherValue"}`. |


**Pipe data via stdin:**
```
$ cat domains.txt | emissary --telegram --stdin
```

**Specify a message as an argument:**
```
$ emissary --telegram --message "This is a very cool message"
```

**Send to multiple channels:**
```
$ cat domains.txt | emissary -t -s -si
```

**Send only 10 lines:**
```
$ cat domains.txt | emissary -t -si --rows 10
```

**Send everything from the file:**
```
$ cat domains.txt | emissary -t -si -r 0
```

Emissary will only send 20 rows by default, this is to protect against accidentally sending a gazillion domains :) It can be overwritten with `--rows 0` which means unlimited rows. 

## Contributing
Any feedback or ideas are welcome! Want to improve something? Create a pull request!

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D