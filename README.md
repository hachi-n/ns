# pigeon
Pigeon is a simple message notification system.
Pegeon's Command Line Interface is a single binary
In other words, just deploy the built program.
Not only that, this program is also available as a library.

If you have a request for the notification function, please create an Issue.

## HOW TO INSTALL?
### When used as a Command
#### Download
``` 
git clone git@github.com:hachi-n/pigeon.git
``` 
#### Setup commands.
```
cp -pr config/application.yml.sample config/application.yml
vim config/application.yml
---- sample -----
slack:
  url:         "https://hooks.slack.com/services/XXX"
  channel:     "#sample"
  messenger:   "slack-bot"
  icon_emoji:  "https://slack.com/img/icons/app-57.png"
  icon_url:    ":ghost:"
----------------

cd ./installer
./install.sh
```

* Will be set `/usr/local/bin/pegion`

### When used as a Library
```
go get github.com/hachi-n/pigeon
```

## HOW TO USE?
### When used as a Command
#### Slack
```
pigeon slack --message "sample message"
```

### When used as a Library
#### Slack
* NewStruct
```
func NewSlack(message, url, channel, messenger, iconEmoji, iconUrl string) *Slack {}
```

* example
```
slack := pigeon.NewSlack(
	"sample message,
	"https://hooks.slack.com/services/XXX",
	"#sample",
	"slack-bot",
	"https://slack.com/img/icons/app-57.png",
	":ghost:",
)
pigeon.Throw(slack)
```
