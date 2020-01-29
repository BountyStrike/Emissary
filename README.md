# Emissary
Send notifications via different channels such as Slack, Telegram, Discord etc.

## Motivation
The idea is to hook Emissary into https://github.com/BountyStrike/Bountystrike-sh which will notify me on Telegram when new domains have been found.

## Usage

**Create a configuration file:**
```
[Telegram]
chat_id=xxxxxx
api_key=xxxxxxx:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

[Slack]
webhook=https://hooks.slack.com/services/xxxxxxxxxx/xxxxxxxxxx/xxxxxxxxxx
```

**Pipe data via stdin:**
```
$ cat domains.txt | emissary -telegram -stdin
```

**Specify a message as an argument:**
```
$ emissary -telegram -message "This is a very cool message"
```

**Send to multiple channels:**
```
$ cat domains.txt | emissary -telegram -slack -stdin
```

Right now the Emissary will only deliver 20 rows, to protect against accidentally sending a gazillion domains :) 

## Todo
Some stuff that I plan to implement:
- [X] Slack
- [X] Telegram
- [ ] Discord
- [ ] Email
- [ ] Let user decide max rows to be sent
- [ ] Place config file in ~/.config/emissary.ini

## Contributing
Any feedback or ideas are welcome! Want to improve something? Create a pull request!

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D