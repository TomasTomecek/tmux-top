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

var c *conf.ConfigurationManager = conf.LoadConf()

func print_mem(*cli.Context) {
	used, total := mem.GetMemStats()
	mem_intervals := c.GetMemIntervals()
	separator := c.GetMemSeparator()
	separator_bg := c.GetMemSeparatorBg()
	separator_fg := c.GetMemSeparatorFg()
	total_bg := c.GetMemTotalBg()
	total_fg := c.GetMemTotalFg()

	fmt.Printf("%s%s%s",
		display.DisplayFloat64(used, 2, mem_intervals, true, "B", total),
		display.DisplayString(separator, separator_bg, separator_fg),
		display.PrintFloat64(total, 2, total_bg, total_fg, true, "B"))
}

func print_load(*cli.Context) {
	one, five, fifteen := load.GetCPULoad()
	load_intervals := c.GetLoadIntervals()

	fmt.Printf("%s %s %s",
		display.DisplayFloat64(one, 2, load_intervals, false, "", 0.0),
		display.DisplayFloat64(five, 2, load_intervals, false, "", 0.0),
		display.DisplayFloat64(fifteen, 2, load_intervals, false, "", 0.0))
}

func print_net(*cli.Context) {
	net_stats := net.GetNetStats(c)
	conf_interfaces := c.GetNetInterfaces()
	net_intervals := c.GetNetIntervals()
	for _, net_stat := range net_stats {
		label := conf_interfaces[net_stat.Name].Alias
		if label == "" {
			label = net_stat.Name
		}
		fmt.Printf("%s %s ",
			display.DisplayString(label, conf_interfaces[net_stat.Name].LabelColorBg,
				conf_interfaces[net_stat.Name].LabelColorFg),
			display.DisplayString(net_stat.Address, conf_interfaces[net_stat.Name].AddressColorBg,
				conf_interfaces[net_stat.Name].AddressColorFg))
		rendered_up := display.DisplayFloat64(net_stat.Tx, 1, net_intervals, true, "B", 0.0)
		if rendered_up != "" {
			fmt.Printf("%s %s ",
				display.DisplayString(c.GetNetUpLabel(), c.GetNetUpLabelBg(), c.GetNetUpLabelFg()),
				rendered_up)
		}
		rendered_down := display.DisplayFloat64(net_stat.Rx, 1, net_intervals, true, "B", 0.0)
		if rendered_down != "" {
			fmt.Printf("%s %s ",
				display.DisplayString(c.GetNetDownLabel(), c.GetNetDownLabelBg(), c.GetNetDownLabelFg()),
				rendered_down)
		}
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
