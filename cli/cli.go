package main

import (
	"fmt"
	"github.com/TomasTomecek/tmux-top/conf"
	"github.com/TomasTomecek/tmux-top/display"
	"github.com/TomasTomecek/tmux-top/load"
	"github.com/TomasTomecek/tmux-top/net"
	"github.com/codegangsta/cli"
	"os"
)

var configuration conf.Configuration = conf.LoadConf()

func print_load(*cli.Context) {
	one, five, fifteen := load.GetCPULoad()
	fmt.Printf("%s %s %s",
		tmux_display.DisplayFloat64(one, 2, configuration.Load.Intervals),
		tmux_display.DisplayFloat64(five, 2, configuration.Load.Intervals),
		tmux_display.DisplayFloat64(fifteen, 2, configuration.Load.Intervals))
}

func print_net(*cli.Context) {
	net_stats := net.GetNetStats(configuration.Net)
	for _, net_stat := range net_stats {
		label := configuration.Net.Interfaces[net_stat.Name].Alias
		if label == "" {
			label = net_stat.Name
		}
		fmt.Printf("%s %s U %s D %s",
			tmux_display.DisplayString(label, configuration.Net.Interfaces[net_stat.Name].LabelColorBg,
				configuration.Net.Interfaces[net_stat.Name].LabelColorFg),
			tmux_display.DisplayString(net_stat.Address, configuration.Net.Interfaces[net_stat.Name].AddressColorBg,
				configuration.Net.Interfaces[net_stat.Name].AddressColorFg),
			tmux_display.DisplayFloat64(net_stat.Tx, 1, configuration.Net.Intervals),
			tmux_display.DisplayFloat64(net_stat.Rx, 1, configuration.Net.Intervals),
		)
	}
}

func main() {

	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Commands = []cli.Command{
		{
			Name:      "net",
			ShortName: "n",
			Usage:     "show net stats ",
			Action:    print_net,
		},
		{
			Name:      "load",
			ShortName: "l",
			Usage:     "show ",
			Action:    print_load,
		}}

	app.Run(os.Args)
}
