package cmd

import (
	"os"

	"github.com/aug0cz/vmrm/internal/db"
	"github.com/aug0cz/vmrm/internal/user"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

func NewUserCommand() *cli.Command {
	return &cli.Command{
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "options for user.",
		Subcommands: []*cli.Command{
			{
				Name:      "create",
				Usage:     "create a user.",
				Args:      true,
				ArgsUsage: "[name]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "email",
						Usage:    "user email address.",
						Aliases:  []string{"e"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "enname",
						Usage:    "user english name.",
						Aliases:  []string{"en"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "department",
						Usage:    "user department with short name",
						Aliases:  []string{"d"},
						Required: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					email := ctx.String("email")
					enName := ctx.String("enname")
					departmentName := ctx.String("department")
					name := ctx.Args().Get(0)
					user := user.User{Name: name, Email: email, NameEn: enName,
						Department: user.Department{ShortName: departmentName}}

					db.DB.Create(&user)

					return nil
				},
			},
			{
				Name:    "list",
				Usage:   "show user list",
				Aliases: []string{"ls"},
				Action: func(ctx *cli.Context) error {
					var users []user.User
					db.DB.Model(&user.User{}).Preload("Department").Find(&users)
					t := table.NewWriter()
					t.SetOutputMirror(os.Stdout)
					t.AppendHeader(table.Row{"#", "name", "english name", "email", "department"})
					for _, u := range users {
						t.AppendRow(table.Row{u.ID, u.Name, u.NameEn, u.Email, u.DepartmentName})
					}
					t.Render()
					return nil
				},
			},
		},
	}
}
