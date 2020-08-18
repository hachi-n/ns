package main

import (
	"github.com/hachi-n/pigeon"
	"github.com/hachi-n/pigeon/cli/pigeon/internal/options"
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

			slack := pigeon.NewSlack(message)
			return pigeon.Deliver(slack)
		},
	}
}
