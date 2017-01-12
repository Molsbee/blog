package command

import (
	"github.com/urfave/cli"
	"github.com/molsbee/blog/service"
	"fmt"
	"github.com/molsbee/blog/app"
)

func Start() cli.Command {
	return cli.Command{
		Name: "start",
		Usage: "[blogName]",
		Description: "starts blog application with configuration file associated with the name provided",
		Action: func(ctx *cli.Context) error {
			blogName := ctx.Args().First();
			if blogName == "" {
				return cli.ShowCommandHelp(ctx, "start")
			}

			configuration, err := service.GetConfiguration(blogName)
			if err != nil {
				return err;
			}

			// TODO Start application
			fmt.Println(configuration.ToString())
			app.Start(*configuration)

			return nil
		},
	}
}

func Stop() cli.Command {
	return cli.Command{
		Name: "stop",
	}
}

func Status() cli.Command {
	return cli.Command{
		Name: "status",
	}
}