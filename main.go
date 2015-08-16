package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tinker"
	app.Usage = "Keep your ideas"
	app.Flags = []cli.Flag{}
	app.Commands = commands
	app.Run(os.Args)
}
