package cluster

import "github.com/urfave/cli/v2"

func NewClusterCommand() *cli.Command {
	return &cli.Command{
		Name:    "cluster",
		Aliases: []string{"c"},
		Usage:   "options for cluster",
		Subcommands: []*cli.Command{
			NewClusterCreateCommand(),
		},
	}
}
