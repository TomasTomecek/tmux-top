package main

import (
	"fmt"
	"github.com/TomasTomecek/tmux-top/conf"
	display "github.com/TomasTomecek/tmux-top/display"
	"github.com/TomasTomecek/tmux-top/load"
	"github.com/TomasTomecek/tmux-top/mem"
	"github.com/TomasTomecek/tmux-top/net"
	"github.com/codegangsta/cli"
	"os"
)

var configuration conf.Configuration = conf.LoadConf()

func print_mem(*cli.Context) {
	used, total := mem.GetMemStats()
	separator := display.DisplayString(configuration.Mem.Separator, configuration.Mem.SeparatorBg,
		configuration.Mem.SeparatorFg)
	fmt.Printf("%s%s%s",
		display.DisplayFloat64(used, 2, configuration.Mem.Intervals, true, "B"),
		separator,
		display.PrintFloat64(total, 2, configuration.Mem.TotalBg, configuration.Mem.TotalFg, true, "B"))
}

func print_load(*cli.Context) {
	one, five, fifteen := load.GetCPULoad()
	fmt.Printf("%s %s %s",
		display.DisplayFloat64(one, 2, configuration.Load.Intervals, false, ""),
		display.DisplayFloat64(five, 2, configuration.Load.Intervals, false, ""),
		display.DisplayFloat64(fifteen, 2, configuration.Load.Intervals, false, ""))
}

func print_net(*cli.Context) {
	net_stats := net.GetNetStats(configuration.Net)
	for _, net_stat := range net_stats {
		label := configuration.Net.Interfaces[net_stat.Name].Alias
		if label == "" {
			label = net_stat.Name
		}
		fmt.Printf("%s %s %s %s %s %s",
			display.DisplayString(label, configuration.Net.Interfaces[net_stat.Name].LabelColorBg,
				configuration.Net.Interfaces[net_stat.Name].LabelColorFg),
			display.DisplayString(net_stat.Address, configuration.Net.Interfaces[net_stat.Name].AddressColorBg,
				configuration.Net.Interfaces[net_stat.Name].AddressColorFg),
			display.DisplayString("U", "default", "white"),
			display.DisplayFloat64(net_stat.Tx, 1, configuration.Net.Intervals, true, "B"),
			display.DisplayString("D", "default", "white"),
			display.DisplayFloat64(net_stat.Rx, 1, configuration.Net.Intervals, true, "B"),
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
			Name:      "mem",
			ShortName: "m",
			Usage:     "show memory stats ",
			Action:    print_mem,
		},
		{
			Name:      "load",
			ShortName: "l",
			Usage:     "show load",
			Action:    print_load,
		}}

	app.Run(os.Args)
}
