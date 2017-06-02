package net

import (
	"fmt"
	"github.com/TomasTomecek/tmux-top/conf"
	"net"
	"testing"
)

func TestReadNetStats(t *testing.T) {
	ifs, err := net.Interfaces()
	if err != nil {
		t.Error("Can't get network interfaces", err)
	}
	first_if_name := ifs[0].Name
	intf_stats_path_rx := fmt.Sprintf("/sys/class/net/%s/statistics/rx_bytes", first_if_name)
	intf_stats_path_tx := fmt.Sprintf("/sys/class/net/%s/statistics/tx_bytes", first_if_name)
	rx, tx := read_net_stats(intf_stats_path_rx, intf_stats_path_tx)
	if rx < 0.0 {
		t.Error("Rx is less than zero")
	}
	if tx < 0.0 {
		t.Error("Tx is less than zero")
	}
}

func TestReadEntries(t *testing.T) {
	ifs, err := net.Interfaces()
	if len(ifs) <= 0 {
		t.Error("no network interfaces")
	}
	if err != nil {
		t.Error("error during getting network interfaces")
	}
	c := &conf.ConfigurationManager{
		Default: &conf.Configuration{
			Net: &conf.NetConfiguration{
				Interfaces: &map[string]conf.NetIFConfiguration{
					ifs[0].Name: conf.NetIFConfiguration{
						Alias: "Banana",
					}}}}}

	stats := read_all_net_stats(c)
	if len(stats) <= 0 {
		t.Error("no stats for network interfaces")
	}
}

func TestGetIOStats(t *testing.T) {
	ifs, err := net.Interfaces()
	if len(ifs) <= 0 {
		t.Error("no network interfaces")
	}
	if err != nil {
		t.Error("error during getting network interfaces")
	}
	c := &conf.ConfigurationManager{
		Default: &conf.Configuration{
			Net: &conf.NetConfiguration{
				Interfaces: &map[string]conf.NetIFConfiguration{
					ifs[0].Name: conf.NetIFConfiguration{
						Alias: "Banana",
					}}}}}

	stats := GetNetStats(c)
	if len(stats) <= 0 {
		t.Error("no stats for network interfaces")
	}

	new_stats := GetNetStats(c)

	if stats[0].Name != new_stats[0].Name {
		t.Error("order of net stats broke")
	}
}
