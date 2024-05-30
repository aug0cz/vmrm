package cluster

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func NewClusterCreateCommand() *cli.Command {
	return &cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage:   "craete a new cluster",
		Action: func(ctx *cli.Context) error {
			fmt.Println("create a new cluster")
			return nil
		},
	}
}
