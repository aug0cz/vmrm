package main

import (
	"log"
	"os"

	"github.com/aug0cz/vmrm/cmd/cluster"
	"github.com/aug0cz/vmrm/internal/service/common"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			cluster.NewClusterCommand(),
		},
	}

	ctx := common.NewSqliteConnectWithContext()

	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}
