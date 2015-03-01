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

type NetStatDiff struct {
	Name    string
	Address string
	RxDiff  float64
	TxDiff  float64
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

func read_all_net_stats(c *conf.ConfigurationManager) map[string]NetStat {
	response := make(map[string]NetStat)
	conf_interfaces := c.GetNetInterfaces()
	ifs, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return response
	}
	for _, intf := range ifs {
		_, ok := conf_interfaces[intf.Name]
		if !ok {
			continue
		}
		addrs, err := intf.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if len(addrs) <= 0 {
			continue
		}

		intf_stats_path_rx := fmt.Sprintf("/sys/class/net/%s/statistics/rx_bytes", intf.Name)
		intf_stats_path_tx := fmt.Sprintf("/sys/class/net/%s/statistics/tx_bytes", intf.Name)
		rx_float, tx_float := read_net_stats(intf_stats_path_rx, intf_stats_path_tx)

		for _, address := range addrs {
			address_s := address.String()
			if !strings.Contains(address_s, ":") { // TODO: add support for IPv6
				s := NetStat{
					Name:    intf.Name,
					Address: address_s,
					Rx:      rx_float,
					Tx:      tx_float,
				}
				response[intf.Name] = s
			}
		}
	}
	return response
}

func GetNetStats(c *conf.ConfigurationManager) []NetStatDiff {
	response := make([]NetStatDiff, 0)

	old_stats := read_all_net_stats(c)

	time.Sleep(1000 * time.Millisecond)

	new_stats := read_all_net_stats(c)

	for key, value := range new_stats {
		if old_value, exists := old_stats[key]; exists {
			d := NetStatDiff{
				Name:    value.Name,
				Address: value.Address,
				RxDiff:  value.Rx - old_value.Rx,
				TxDiff:  value.Tx - old_value.Tx,
			}
			response = append(response, d)
		}
	}

	return response
}
