package conf

import (
	"encoding/json"
	"fmt"
	"github.com/TomasTomecek/tmux-top/humanize"
	"io/ioutil"
	"math"
	"os"
	"path"
)

type IntervalDisplay struct {
	From    *string `json:"from"`
	To      *string `json:"to"`
	BgColor string  `json:"bg_color"`
	FgColor string  `json:"fg_color"`
}

func loadIntervalValue(s string) (float64, bool) {
	f, err := humanize.DehumanizeString(s)
	if err == nil {
		return f, false
	}
	f, err = humanize.Absolutize(s)
	if err == nil {
		return f, true
	}
	return 0.0, false
}

func (i *IntervalDisplay) GetFrom() (float64, bool) {
	if i.From == nil {
		return math.NaN(), false
	}
	return loadIntervalValue(*i.From)
}

func (i *IntervalDisplay) GetTo() (float64, bool) {
	if i.To == nil {
		return math.NaN(), false
	}
	return loadIntervalValue(*i.To)
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
	UpLabel     *string                        `json:"upload_label"`
	UpLabelBg   *string                        `json:"upload_label_bg"`
	UpLabelFg   *string                        `json:"upload_label_fg"`
	DownLabel   *string                        `json:"download_label"`
	DownLabelBg *string                        `json:"download_label_bg"`
	DownLabelFg *string                        `json:"download_label_fg"`
	Intervals   *[]IntervalDisplay             `json:"intervals"`
}

type IODeviceConfiguration struct {
	Alias        string `json:"alias"`
	LabelColorFg string `json:"label_color_fg"`
	LabelColorBg string `json:"label_color_bg"`
}

type IOConfiguration struct {
	Devices      *map[string]IODeviceConfiguration `json:"devices"`
	ReadLabel    *string                           `json:"read_label"`
	ReadLabelBg  *string                           `json:"read_label_bg"`
	ReadLabelFg  *string                           `json:"read_label_fg"`
	WriteLabel   *string                           `json:"write_label"`
	WriteLabelBg *string                           `json:"write_label_bg"`
	WriteLabelFg *string                           `json:"write_label_fg"`
	Intervals    *[]IntervalDisplay                `json:"intervals"`
}

type MemConfiguration struct {
	Intervals   *[]IntervalDisplay `json:"intervals"`
	Separator   *string            `json:"separator"`
	SeparatorBg *string            `json:"separator_bg"`
	SeparatorFg *string            `json:"separator_fg"`
	TotalBg     *string            `json:"total_bg"`
	TotalFg     *string            `json:"total_fg"`
}

// configuration in a config file
type SensorsConfiguration struct {
	Template *string `json:"template"`
}

type LoadConfiguration struct {
	Intervals *[]IntervalDisplay `json:"intervals"`
}

type Configuration struct {
	Load    *LoadConfiguration    `json:"load"`
	Net     *NetConfiguration     `json:"net"`
	Mem     *MemConfiguration     `json:"mem"`
	IO      *IOConfiguration      `json:"io"`
	Sensors *SensorsConfiguration `json:"sensors"`
}

func loadConfFromFile(path string) []byte {
	response, _ := ioutil.ReadFile(path)
	/*if err != nil {
		fmt.Println("error:", err)
	}*/
	return response
}

func LoadConfFromBytes(json_input []byte) *Configuration {
	var conf *Configuration = new(Configuration)
	err := json.Unmarshal(json_input, &conf)

	if err != nil {
		fmt.Println("error:", err)
	}

	return conf
}

func LoadConf() *ConfigurationManager {
	var c *ConfigurationManager = new(ConfigurationManager)
	home_dir := os.Getenv("HOME")
	bytes := loadConfFromFile(path.Join(home_dir, ".tmux-top"))
	if len(bytes) > 0 {
		c.User = LoadConfFromBytes(bytes)
	}
	c.Default = LoadConfFromBytes([]byte(default_conf))
	return c
}
