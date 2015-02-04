package net

import (
	"fmt"
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
	return
}

func GetNetStats() []NetStat {
	response := make([]NetStat, 0)
	ifs, err := net.Interfaces()
	if err != nil {
		fmt.Println(nil)
	}
	for _, intf := range ifs {
		if intf.Name == "lo" {
			continue
		}
		addrs, err := intf.Addrs()
		if len(addrs) <= 0 {
			continue
		}

		intf_stats_path_rx := fmt.Sprintf("/sys/class/net/%s/statistics/rx_bytes", intf.Name)
		intf_stats_path_tx := fmt.Sprintf("/sys/class/net/%s/statistics/tx_bytes", intf.Name)
		rx, err := ioutil.ReadFile(intf_stats_path_rx)
		if err != nil {
			fmt.Println(err)
		}
		rx_float, err := strconv.ParseFloat(strings.Trim(string(rx), "\n"), 64)
		if err != nil {
			fmt.Println(err)
		}
		tx, err := ioutil.ReadFile(intf_stats_path_tx)
		if err != nil {
			fmt.Println(err)
		}
		tx_float, err := strconv.ParseFloat(strings.Trim(string(tx), "\n"), 64)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(1000 * time.Millisecond)

		rx_after, err := ioutil.ReadFile(intf_stats_path_rx)
		if err != nil {
			fmt.Println(err)
		}
		rx_float_after, err := strconv.ParseFloat(strings.Trim(string(rx_after), "\n"), 64)
		if err != nil {
			fmt.Println(err)
		}
		tx_after, err := ioutil.ReadFile(intf_stats_path_tx)
		if err != nil {
			fmt.Println(err)
		}
		tx_float_after, err := strconv.ParseFloat(strings.Trim(string(tx_after), "\n"), 64)
		if err != nil {
			fmt.Println(err)
		}

		rx_diff := rx_float_after - rx_float
		tx_diff := tx_float_after - tx_float

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
	return response
}

func main() {
	fmt.Println(GetNetStats())
}
