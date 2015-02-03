package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"tmux-top/conf"
	"tmux-top/display"
	"tmux-top/load"
)

var configuration conf.Configuration = conf.LoadConf()

func print_load(*cli.Context) {
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
