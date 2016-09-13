package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	cmds := []cli.Command{
		NewRunCommand(),
		NewInstallCommand(),
	}

	app := cli.NewApp()
	app.Commands = cmds
	app.Version = "1.0.0"
	app.Usage = "Docker Run Tool is a CLI wrapper for running utility containers"
	app.UsageText = "drt help\n\tdrt run {PATH/TO/MANIFEST.yml} {ARGS}"

	e := app.Run(os.Args)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
