package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
)

type IntervalDisplay struct {
	From    *float64 `json:"from"`
	To      *float64 `json:"to"`
	BgColor string   `json:"bg_color"`
	FgColor string   `json:"fg_color"`
}

func (i *IntervalDisplay) GetFrom() float64 {
	if i.From == nil {
		return math.NaN()
	}
	return *i.From
}

func (i *IntervalDisplay) GetTo() float64 {
	if i.To == nil {
		return math.NaN()
	}
	return *i.To
}

type NetIFConfiguration struct {
	Alias          string `json:"alias"`
	LabelColorFg   string `json:"label_color_fg"`
	LabelColorBg   string `json:"label_color_bg"`
	AddressColorFg string `json:"address_color_fg"`
	AddressColorBg string `json:"address_color_bg"`
}

type NetConfiguration struct {
	Interfaces  *map[string]NetIFConfiguration `json:"interfaces"`
	Threshold   *float64                       `json:"threshold"`
	UpLabel     *string                        `json:"upload_label"`
	UpLabelBg   *string                        `json:"upload_label_bg"`
	UpLabelFg   *string                        `json:"upload_label_fg"`
	DownLabel   *string                        `json:"download_label"`
	DownLabelBg *string                        `json:"download_label_bg"`
	DownLabelFg *string                        `json:"download_label_fg"`
	Intervals   *[]IntervalDisplay             `json:"intervals"`
}

type MemConfiguration struct {
	Intervals   *[]IntervalDisplay `json:"intervals"`
	Separator   *string            `json:"separator"`
	SeparatorBg *string            `json:"separator_bg"`
	SeparatorFg *string            `json:"separator_fg"`
	TotalBg     *string            `json:"total_bg"`
	TotalFg     *string            `json:"total_fg"`
}

type LoadConfiguration struct {
	Intervals *[]IntervalDisplay `json:"intervals"`
}

type Configuration struct {
	Load *LoadConfiguration `json:"load"`
	Net  *NetConfiguration  `json:"net"`
	Mem  *MemConfiguration  `json:"mem"`
}

func loadConfFromFile(path string) []byte {
	response, _ := ioutil.ReadFile(path)
	/*if err != nil {
		fmt.Println("error:", err)
	}*/
	return response
}

func loadConfFromBytes(json_input []byte) *Configuration {
	var conf *Configuration = new(Configuration)
	err := json.Unmarshal(json_input, &conf)

	if err != nil {
		fmt.Println("error:", err)
	}

	return conf
}

func LoadConf() *ConfigurationManager {
	var c *ConfigurationManager = new(ConfigurationManager)
	bytes := loadConfFromFile("~/.tmux-top")
	if len(bytes) > 0 {
		c.User = loadConfFromBytes(bytes)
	}
	c.Default = loadConfFromBytes([]byte(default_conf))
	return c
}
