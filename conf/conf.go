package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type IntervalDisplay struct {
	From    float64 `json:"from"`
	To      float64 `json:"to"`
	BgColor string  `json:"bg_color"`
	FgColor string  `json:"fg_color"`
}

type NetIFConfiguration struct {
	Alias          string `json:"alias"`
	LabelColorFg   string `json:"label_color_fg"`
	LabelColorBg   string `json:"label_color_bg"`
	AddressColorFg string `json:"address_color_fg"`
	AddressColorBg string `json:"address_color_bg"`
}

type NetConfiguration struct {
	Interfaces map[string]NetIFConfiguration `json:"interfaces"`
	Threshold  float64                       `json:"threshold"`
	Intervals  []IntervalDisplay             `json:"intervals"`
}

type MemConfiguration struct {
	Intervals   []IntervalDisplay `json:"intervals"`
	Separator   string            `json:"separator"`
	SeparatorBg string            `json:"separator_bg"`
	SeparatorFg string            `json:"separator_fg"`
	TotalBg     string            `json:"total_bg"`
	TotalFg     string            `json:"total_fg"`
}

type LoadConfiguration struct {
	Intervals []IntervalDisplay `json:"intervals"`
}

type Configuration struct {
	Load LoadConfiguration `json:"load"`
	Net  NetConfiguration  `json:"net"`
	Mem  MemConfiguration  `json:"mem"`
}

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
