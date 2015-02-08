package net

import (
	"fmt"
	"github.com/TomasTomecek/tmux-top/conf"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
	"time"
)

type NetStat struct {
	Name    string
	Address string
	Rx      float64
	Tx      float64
}

func (s *NetStat) String() (response string) {
	return fmt.Sprintf("%s %s: U %.1f D %.1f", s.Name, s.Address,
		s.Tx, s.Rx)
}

func read_net_stats(rx_path, tx_path string) (float64, float64) {
	var rx_float, tx_float float64 = -1.0, -1.0
	var err error

	rx, err := ioutil.ReadFile(rx_path)
	if err != nil {
		fmt.Println(err)
	}
	rx_float, err = strconv.ParseFloat(strings.Trim(string(rx), "\n"), 64)
	if err != nil {
		rx_float = -1.0
		fmt.Println(err)
	}
	tx, err := ioutil.ReadFile(tx_path)
	if err != nil {
		fmt.Println(err)
	}
	tx_float, err = strconv.ParseFloat(strings.Trim(string(tx), "\n"), 64)
	if err != nil {
		tx_float = -1.0
		fmt.Println(err)
	}
	return rx_float, tx_float
}

func GetNetStats(iface_conf conf.NetConfiguration) []NetStat {
	response := make([]NetStat, 0)
	ifs, err := net.Interfaces()
	if err != nil {
		fmt.Println(nil)
	}
	for _, intf := range ifs {
		_, ok := iface_conf.Interfaces[intf.Name]
		if !ok {
			fmt.Printf("%s not requested\n", intf.Name)
			continue
		}
		addrs, err := intf.Addrs()
		if err != nil {
			fmt.Println(err)
		}
		if len(addrs) <= 0 {
			continue
		}

		intf_stats_path_rx := fmt.Sprintf("/sys/class/net/%s/statistics/rx_bytes", intf.Name)
		intf_stats_path_tx := fmt.Sprintf("/sys/class/net/%s/statistics/tx_bytes", intf.Name)
		rx_float, tx_float := read_net_stats(intf_stats_path_rx, intf_stats_path_tx)

		time.Sleep(1000 * time.Millisecond)

		rx_float_after, tx_float_after := read_net_stats(intf_stats_path_rx, intf_stats_path_tx)

		rx_diff := rx_float_after - rx_float
		tx_diff := tx_float_after - tx_float

		if rx_diff < iface_conf.Threshold && tx_diff < iface_conf.Threshold {
			continue
		}

		for _, address := range addrs {
			address_s := address.String()
			if !strings.Contains(address_s, ":") { // ipv6
				s := NetStat{
					Name:    intf.Name,
					Address: address_s,
					Rx:      rx_diff,
					Tx:      tx_diff,
				}
				fmt.Println(s)
				response = append(response, s)
			}
		}
	}
	fmt.Println(response)
	return response
}
