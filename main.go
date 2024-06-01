package main

import (
	"log"
	"os"

	"github.com/aug0cz/vmrm/cmd"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "vmrmctl",
		Usage: "virtual machine resources management.",
		Commands: []*cli.Command{
			cmd.NewUserCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
