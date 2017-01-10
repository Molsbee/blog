package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/molsbee/blog/command"
)

func main() {
	application := cli.NewApp()
	application.Name = "blog"
	application.Usage = "Command line Utility for setting up and running Blogging Application"

	application.Before = func(ctx *cli.Context) error {
		command := ctx.Args().First()

		if command != "configure" {
			// TODO Check Application Configuration File
		}

		return nil
	}

	application.Commands = []cli.Command{
		command.Start(),
		command.Stop(),
		command.Status(),
		command.NewConfigurationCommands(),
	}
	// TODO Setup application to run as a Service

	application.Run(os.Args)
}
