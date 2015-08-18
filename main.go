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
	// Create default data directory.
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		createDataPath()
	}

	app.Run(os.Args)
}
