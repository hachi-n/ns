package main

import (
	"github.com/hachi-n/pigeon"
	"github.com/hachi-n/pigeon/cli/pigeon/internal/options"
	"github.com/hachi-n/pigeon/lib/static"
	"github.com/urfave/cli/v2"
)

func slackCommand() *cli.Command {
	var message string
	return &cli.Command{
		Name:  "slack",
		Usage: "slack post message.",
		Flags: []cli.Flag{
			options.MessageFlag(&message, true, ""),
		},
		Action: func(c *cli.Context) error {
			yml, err := static.ApplicationYamlLoad()
			if err != nil {
				return err
			}

			slack := pigeon.NewSlack(
				message,
				yml.Slack.Url,
				yml.Slack.Channel,
				yml.Slack.Messenger,
				yml.Slack.IconEmoji,
				yml.Slack.IconUrl,
			)
			return pigeon.Throw(slack)
		},
	}
}
