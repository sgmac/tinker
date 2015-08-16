package main

import "github.com/codegangsta/cli"

var commands = []cli.Command{
	{
		Name:   "add",
		Usage:  "add idea topic ",
		Action: func(c *cli.Context) {},
	},
	{
		Name:   "list",
		Usage:  "list ideas",
		Action: func(c *cli.Context) {},
	},
	{
		Name:   "delete",
		Usage:  "delete idea",
		Action: func(c *cli.Context) {},
	},
	{
		Name:   "done",
		Usage:  "done idea",
		Action: func(c *cli.Context) {},
	},
	{
		Name:   "start",
		Usage:  "start work idea",
		Action: func(c *cli.Context) {},
	},
	{
		Name:  "topic",
		Usage: "manage your topics",
		Subcommands: []cli.Command{
			{
				Name:   "add",
				Usage:  "add topic",
				Action: func(c *cli.Context) {},
			},
			{
				Name:   "delete",
				Usage:  "delete topic",
				Action: func(c *cli.Context) {},
			},
			{
				Name:   "list",
				Usage:  "list topics",
				Action: func(c *cli.Context) {},
			},
			{
				Name:   "set",
				Usage:  "set a default topic",
				Action: func(c *cli.Context) {},
			},
		},
	},
}
