package main

import (
	"strings"

	"github.com/codegangsta/cli"
)

var commands = []cli.Command{
	{
		Name:  "add",
		Usage: "add idea topic ",
		Action: func(c *cli.Context) {
			if len(c.Args()) < 2 {
				cli.ShowSubcommandHelp(c)
				return
			}
			topic := c.Args()[0]
			idea := strings.Join(c.Args()[1:], " ")
			addIdea(topic, idea)
		},
	},
	{
		Name:  "list",
		Usage: "list ideas",
		Action: func(c *cli.Context) {
			if len(c.Args()) < 1 {
				cli.ShowSubcommandHelp(c)
				return
			}
			topic := c.Args()[0]
			listIdeas(topic)
		},
	},
	{
		Name:  "delete",
		Usage: "delete idea",
		Action: func(c *cli.Context) {
			if len(c.Args()) < 2 {
				cli.ShowSubcommandHelp(c)
				return
			}
			topic := c.Args()[0]
			id := c.Args()[1]
			deleteIdea(topic, id)
		},
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
				Name:  "add",
				Usage: "add topic",
				Action: func(c *cli.Context) {
					if len(c.Args()) < 1 {
						cli.ShowSubcommandHelp(c)
						return
					}
					topic := strings.Join(c.Args(), " ")
					addNewTopic(topic)
				},
			},
			{
				Name:  "delete",
				Usage: "delete topic",
				Action: func(c *cli.Context) {
					topic := strings.Join(c.Args(), " ")
					deleteTopic(topic)
				},
			},
			{
				Name:  "list",
				Usage: "list topics",
				Action: func(c *cli.Context) {
					listTopics()
				},
			},
			{
				Name:  "set",
				Usage: "set a default topic",
				Action: func(c *cli.Context) {
					topic := strings.Join(c.Args(), " ")
					if len(c.Args()) < 1 {
						cli.ShowSubcommandHelp(c)
						return
					}
					setDefaultTopic(topic)
				},
			},
		},
	},
}
