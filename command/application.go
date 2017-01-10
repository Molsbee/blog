package command

import "github.com/urfave/cli"

func Start() cli.Command {
	return cli.Command{
		Name: "start",
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