package main

import (
	"cliif/tmux_display"
	"conf"
	"fmt"
	"github.com/codegangsta/cli"
	"load"
	"os"
)

var configuration conf.Configuration = conf.LoadConfiguration()

func print_load() {
	one, five, fifteen := load.GetCPULoad()
	fmt.Printf("%s %s %s",
		tmux_display.DisplayFloat64(one, 2, configuration.Load.Intervals),
		tmux_display.DisplayFloat64(five, 2, configuration.Load.Intervals),
		tmux_display.DisplayFloat64(fifteen, 2, configuration.Load.Intervals))
}

func main() {

	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Commands = []cli.Command{
		{
			Name:      "load",
			ShortName: "l",
			Usage:     "show ",
			Action:    print_load,
		}}

	app.Run(os.Args)
}
