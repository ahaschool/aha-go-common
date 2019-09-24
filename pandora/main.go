package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	b build
	r run
	obj = [2]string{"api", "srv"}
)

func main() {
	app := cli.NewApp()
	app.Name = "pandora"
	app.Usage = "pandora 工具集"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "new pandora project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "o",
					Value:       "",
					Usage:       "project owner for create project",
					Destination: &p.Owner,
				},
				cli.StringFlag{
					Name:        "d",
					Value:       "",
					Usage:       "project directory for create project",
					Destination: &p.Path,
				},
			},
			Action:  runNew,
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build pandora project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "name",
					Value:       "server",
					Usage:       "build name for this project",
					Destination: &b.Name,
				},
			},
			Action:  buildAction,
		},
		{
			Name:    "gateway",
			Aliases: []string{"g"},
			Usage:   "start micro gateway",
			Action: runGatewayAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "port",
					Value:       "8181",
					Usage:       "start name for this project",
					Destination: &r.Port,
				},
				cli.StringFlag{
					Name:        "registry_address",
					Value:       "127.0.0.1:8500",
					Usage:       "registry address",
					Destination: &r.RegistryAddress,
				},
				cli.StringFlag{
					Name:        "namespace",
					Value:       "cn.com.ahaschool.api",
					Usage:       "namespace",
					Destination: &r.Namespace,
				},
			},
		},
		{
			Name:    "start",
			Aliases: []string{"start"},
			Usage:   "start pandora project at product",
			Action: runStartAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "name",
					Value:       "server",
					Usage:       "start name for this project",
					Destination: &b.Name,
				},
			},
		},
		{
			Name:    "stop",
			Aliases: []string{"stop"},
			Usage:   "stop pandora project at product",
			Action: runStopAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "name",
					Value:       "server",
					Usage:       "stop name for this project",
					Destination: &b.Name,
				},
			},
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "pandora version",
			Action: func(c *cli.Context) error {
				fmt.Println(Version)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
