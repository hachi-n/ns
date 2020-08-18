# pigeon
pigeon is a simple message library.

## HOW TO INSTALL?
### Please Exec Command
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
