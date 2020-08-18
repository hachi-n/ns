//go:generate pkger -o generated

package main

import (
	_ "github.com/hachi-n/pigeon/generated"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:        "pigeon",
		Usage:       "notification system.",
		Description: "notification system.",
		Commands: []*cli.Command{
			slackCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
