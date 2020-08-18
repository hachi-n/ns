package options

import "github.com/urfave/cli/v2"

func MessageFlag(destination *string, required bool, defaultValue string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "message",
		Value:       defaultValue,
		Usage:       "",
		Required:    required,
		Destination: destination,
	}
}

