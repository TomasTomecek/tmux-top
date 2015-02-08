package net

import (
	"fmt"
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
