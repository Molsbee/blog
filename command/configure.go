package command

import (
	"bufio"
	"fmt"
	"github.com/molsbee/blog/service"
	"github.com/urfave/cli"
	"os"
)

func NewConfigurationCommands() cli.Command {
	return cli.Command{
		Name:  "configuration",
		Usage: "Configure application with endpoints/credentials",
		Subcommands: []cli.Command{
			set(),
			get(),
		},
	}
}

func get() cli.Command {
	return cli.Command{
		Name:  "get",
		Usage: "{{ application name }} - returns all data that has been configured with application",
		Action: func(ctx *cli.Context) error {
			blogName := ctx.Args().First()
			if blogName == "" {
				return cli.ShowCommandHelp(ctx, "get")
			}

			configuration, err := service.GetConfiguration(blogName)
			fmt.Println(configuration.ToString())

			return err
		},
	}
}

func set() cli.Command {
	return cli.Command{
		Name:        "set",
		Usage:       "Calling function will start a set of command prompts to facilitate proper configuration",
		Description: "Provides the ability to configure application with appropriate dependencies.",
		Action: func(ctx *cli.Context) error {
			var (
				blogName        string
				applicationPort string
				username        string
				password        string
				hostName        string
				databasePort    string
			)

			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Blog name: ")
			fmt.Fscanln(reader, &blogName)

			fmt.Print("Application Port: ")
			fmt.Fscanln(reader, &applicationPort)

			fmt.Print("Database Username: ")
			fmt.Fscanln(reader, &username)

			fmt.Print("Database Password: ")
			fmt.Fscanln(reader, &password) // TODO Replace with gopass

			fmt.Print("Database HostName: ")
			fmt.Fscanln(reader, &hostName)

			fmt.Print("Database Port: ")
			fmt.Fscanln(reader, &databasePort)

			fmt.Printf("Blog name: %s port: %s username: %s password: %s hostName: %s port: %s\n", blogName, applicationPort, username, password, hostName, databasePort)
			return service.NewConfiguration(blogName, username, password, hostName, applicationPort, databasePort).Save()
		},
	}
}
