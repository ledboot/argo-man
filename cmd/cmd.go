package cmd

import (
	"fmt"
	"github.com/ledboot/argo-man/config"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"runtime"
)

func InitArgs() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s has version %s build on %s @%s", config.GetBuildCfg().ServiceName, config.GetBuildCfg().Version, config.GetBuildCfg().BuildTime, runtime.Version())
	}
	commands := []*cli.Command{
		{
			Name: "app",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "app",
				},
			},
			Subcommands: []*cli.Command{
				{
					Name:  "create",
					Usage: "create an application",
					Action: func(context *cli.Context) error {
						fmt.Println("create an application", config.GetBuildCfg().ConfigFile)
						createApp()
						return nil
					},
				},
				{
					Name:  "sync",
					Usage: "sync an application. eg: sync serviceName",
					Action: func(context *cli.Context) error {
						fmt.Println("sync an application", context.Args().First())
						return nil
					},
				},
				{
					Name:  "delete",
					Usage: "delete an application. eg: delete serviceName",
					Action: func(context *cli.Context) error {
						//fmt.Println("delete an application", context.Args().First())
						deleteAll()
						return nil
					},
				},
			},
		},
	}
	app := &cli.App{
		Name: config.GetBuildCfg().ServiceName,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "config file path, eg: -c /etc/config.toml",
				Value:       "",
				Required:    true,
				Destination: &config.GetBuildCfg().ConfigFile,
			},
			&cli.BoolFlag{
				Name:     "version",
				Aliases:  []string{"v"},
				Usage:    "print the version",
				Required: false,
				Action: func(context *cli.Context, b bool) error {
					cli.VersionPrinter(context)
					os.Exit(0)
					return nil
				},
			},
		},
		Commands: commands,
		Before: func(context *cli.Context) error {
			if err := config.LoadConfig(&config.Cfg); err != nil {
				return err
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
