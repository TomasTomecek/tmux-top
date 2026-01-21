package main

import (
	"fmt"
	"github.com/TomasTomecek/tmux-top/conf"
	"github.com/TomasTomecek/tmux-top/disk"
	display "github.com/TomasTomecek/tmux-top/display"
	"github.com/TomasTomecek/tmux-top/io"
	"github.com/TomasTomecek/tmux-top/journal"
	"github.com/TomasTomecek/tmux-top/load"
	"github.com/TomasTomecek/tmux-top/mem"
	"github.com/TomasTomecek/tmux-top/net"
	"github.com/TomasTomecek/tmux-top/sens"
	cli "github.com/urfave/cli/v2"
	"os"
)

var c *conf.ConfigurationManager = conf.LoadConf()

func print_mem(ctx *cli.Context) error {
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
	return nil
}

func print_load(ctx *cli.Context) error {
	one, five, fifteen := load.GetCPULoad()
	load_intervals := c.GetLoadIntervals()

	fmt.Printf("%s %s %s",
		display.DisplayFloat64(one, 2, load_intervals, false, "", 0.0),
		display.DisplayFloat64(five, 2, load_intervals, false, "", 0.0),
		display.DisplayFloat64(fifteen, 2, load_intervals, false, "", 0.0))
	return nil
}

func print_net(ctx *cli.Context) error {
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
		rendered_up := display.DisplayFloat64(net_stat.TxDiff, 1, net_intervals, true, "B", 0.0)
		if rendered_up != "" {
			fmt.Printf("%s %s ",
				display.DisplayString(c.GetNetUpLabel(), c.GetNetUpLabelBg(), c.GetNetUpLabelFg()),
				rendered_up)
		}
		rendered_down := display.DisplayFloat64(net_stat.RxDiff, 1, net_intervals, true, "B", 0.0)
		if rendered_down != "" {
			fmt.Printf("%s %s ",
				display.DisplayString(c.GetNetDownLabel(), c.GetNetDownLabelBg(), c.GetNetDownLabelFg()),
				rendered_down)
		}
	}
	return nil
}

func print_io(ctx *cli.Context) error {
	io_stats := io.GetIOStats(c)
	conf_devices := c.GetIODevices()
	io_intervals := c.GetIOIntervals()
	for _, stat := range io_stats {
		rendered_read := display.DisplayFloat64(stat.BytesRead, 1, io_intervals, true, "B", 0.0)
		rendered_write := display.DisplayFloat64(stat.BytesWritten, 1, io_intervals, true, "B", 0.0)
		if rendered_write == "" && rendered_read == "" {
			continue
		}
		label := conf_devices[stat.Name].Alias
		if label == "" {
			label = stat.Name
		}
		fmt.Printf("%s ",
			display.DisplayString(label, conf_devices[stat.Name].LabelColorBg,
				conf_devices[stat.Name].LabelColorFg))
		if rendered_read != "" {
			fmt.Printf("%s %s ",
				display.DisplayString(c.GetIOReadLabel(), c.GetIOReadLabelBg(), c.GetIOReadLabelFg()),
				rendered_read)
		}
		if rendered_write != "" {
			fmt.Printf("%s %s ",
				display.DisplayString(c.GetIOWriteLabel(), c.GetIOWriteLabelBg(), c.GetIOWriteLabelFg()),
				rendered_write)
		}
	}
	return nil
}

func print_sens(ctx *cli.Context) error {
	template := c.GetSensorsTemplate(ctx.String("format"))
	sens.PrintSensorStats(template)
	return nil
}

func print_disk(ctx *cli.Context) error {
	template := c.GetDiskTemplate(ctx.String("format"))
	disk.PrintDiskStats(template, c)
	return nil
}

func print_journal(ctx *cli.Context) error {
	timeframe := ctx.String("timeframe")
	format := ctx.String("format")

	if format == "" {
		// Simple display mode with intervals
		errorCount, err := journal.GetJournalErrorCount(timeframe)
		if err != nil {
			fmt.Printf("Journal unavailable")
			return nil
		}

		journal_intervals := c.GetJournalIntervals()
		fmt.Printf("%s %s",
			display.DisplayString("J", c.GetJournalLabelBg(), c.GetJournalLabelFg()),
			display.DisplayFloat64(float64(errorCount), 0, journal_intervals, false, "", 0.0))
	} else {
		// Template mode
		template := c.GetJournalTemplate(format)
		journal.PrintJournalStats(template, timeframe)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Version = "0.1.1"
	app.Name = "tmux-top"
	app.Usage = "monitoring information for your tmux status line"
	app.Commands = []*cli.Command{
		{
			Name:    "net",
			Aliases: []string{"n"},
			Usage:   "show net stats",
			Action:  print_net,
		},
		{
			Name:    "mem",
			Aliases: []string{"m"},
			Usage:   "show memory stats",
			Action:  print_mem,
		},
		{
			Name:    "load",
			Aliases: []string{"l"},
			Usage:   "show load",
			Action:  print_load,
		},
		{
			Name:    "io",
			Aliases: []string{"i"},
			Usage:   "show I/O stats",
			Action:  print_io,
		},
		{
			Name:    "sensors",
			Aliases: []string{"s"},
			Usage:   "show sensor stats (temperature)",
			Action:  print_sens,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "format",
					Aliases: []string{"f"},
					Usage:   "Format the output using the given Go template",
				},
			},
		},
		{
			Name:    "disk",
			Aliases: []string{"d"},
			Usage:   "show disk space stats",
			Action:  print_disk,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "format",
					Aliases: []string{"f"},
					Usage:   "Format the output using the given Go template",
				},
			},
		},
		{
			Name:    "journal",
			Aliases: []string{"j"},
			Usage:   "show journald error counts",
			Action:  print_journal,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "timeframe",
					Aliases: []string{"t"},
					Value:   "1h",
					Usage:   "Time window for error aggregation (1m, 5m, 1h, 24h)",
				},
				&cli.StringFlag{
					Name:    "format",
					Aliases: []string{"f"},
					Usage:   "Format the output using the given Go template",
				},
			},
		},
		{
			Name:  "generate-man",
			Usage: "Generate man page",
			Action: func(ctx *cli.Context) error {
				man, err := app.ToMan()
				if err != nil {
					return err
				}
				fmt.Print(man)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
