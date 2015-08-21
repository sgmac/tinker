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
			var idea string
			reqArgs := 2
			topic := getDefaultTopic()
			if topic != "" {
				reqArgs = 1
			}
			if len(c.Args()) < reqArgs {
				cli.ShowSubcommandHelp(c)
				return
			}
			// If there is no default topic, or if there is, but the first
			// arg is a topic, add the idea to that topic. Preference arg topic.

			if topic == "" {
				topic = c.Args()[0]
			}

			if len(c.Args()) > 2 {
				if isValidTopic(c.Args()[0]) {
					topic = c.Args()[0]
				}
			}
			idea = strings.Join(c.Args()[1:], " ")
			addIdea(topic, idea)
		},
	},
	{
		Name:  "list",
		Usage: "list ideas",
		Action: func(c *cli.Context) {
			reqArgs := 1
			topic := getDefaultTopic()
			if topic != "" {
				reqArgs = 0
			}
			if len(c.Args()) < reqArgs {
				cli.ShowSubcommandHelp(c)
				return
			}
			if len(c.Args()) == 1 {
				topic = c.Args()[0]
			}
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
				Name:  "get",
				Usage: "get default topic",
				Action: func(c *cli.Context) {
					readDefaultTopic()
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
