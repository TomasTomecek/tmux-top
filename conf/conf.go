package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type IntervalDisplay struct {
	From     float64 `json:"from"`
	To       float64 `json:"to"`
	Bg_color string  `json:"bg_color"`
	Fg_color string  `json:"fg_color"`
}

type NetIFConfiguration struct {
	Alias          string `json:alias`
	LabelColorFg   string `json:label_color_fg`
	LabelColorBg   string `json:label_color_bg`
	AddressColorFg string `json:address_color_fg`
	AddressColorBg string `json:address_color_bg`
}

type NetConfiguration struct {
	Interfaces map[string]NetIFConfiguration `json:interfaces`
	Threshold  float64                       `json:threshold`
	Intervals  []IntervalDisplay             `json:"intervals"`
}

type LoadConfiguration struct {
	Intervals []IntervalDisplay `json:"intervals"`
}

type Configuration struct {
	Load LoadConfiguration `json:"load"`
	Net  NetConfiguration  `json:"net"`
}

var default_conf string = `
{
	"load": {
		"intervals": [{
			"to": 1.0,
			"bg_color": "blue",
			"fg_color": "black"
		}]
	},
	"net": {
		"interfaces": {
			"wlp3s0": {
				"alias": "E",
				"label_color_fg": "white",
				"label_color_bg": "default",
				"address_color_fg": "blue",
				"address_color_bg": "default"
			}
		},
		"threshold": 0.0,
		"intervals": [{
			"to": 1.0,
			"bg_color": "blue",
			"fg_color": "black"
		}]
	}

}`

func (c *Configuration) GetConf(conf_module string) LoadConfiguration {
	return c.Load
}

func loadConfFromFile(path string) []byte {
	response, _ := ioutil.ReadFile(path)
	/*if err != nil {
		fmt.Println("error:", err)
	}*/
	return response
}

func loadConfFromBytes(json_input []byte) Configuration {
	var conf Configuration
	err := json.Unmarshal(json_input, &conf)

	if err != nil {
		fmt.Println("error:", err)
	}

	return conf
}

func LoadConf() Configuration {
	bytes := loadConfFromFile("~/.tmux-top")
	if len(bytes) <= 0 {
		return loadConfFromBytes([]byte(default_conf))
	} else {
		return loadConfFromBytes(bytes)
	}
}
