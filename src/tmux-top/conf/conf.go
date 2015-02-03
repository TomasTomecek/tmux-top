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

type LoadConfiguration struct {
	Intervals []IntervalDisplay `json:"intervals"`
}

type Configuration struct {
	Load LoadConfiguration `json:"load"`
}

/*
{
	"load": {
		"intervals": [{...}]
	},
	"net": {
		label: {...}
		"colors": [{...}]
	}
}
*/
var default_conf string = `
{
	"load": {
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
