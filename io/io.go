/*
docs https://www.kernel.org/doc/Documentation/block/stat.txt
*/

package io

import (
	//"fmt"
	"github.com/TomasTomecek/tmux-top/conf"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var BASE_PATH string = "/sys/block/"

type IOStat struct {
	Name           string
	SectorsRead    float64
	SectorsWritten float64
}

type IOStatDiff struct {
	Name         string
	BytesRead    float64
	BytesWritten float64
}

func readStats(c *conf.ConfigurationManager, devs []os.FileInfo) map[string]IOStat {
	response := make(map[string]IOStat)
	conf_devices := c.GetIODevices()

	for _, d := range devs {
		if _, ok := conf_devices[d.Name()]; !ok {
			continue
		}
		dev_path := path.Join(BASE_PATH, d.Name(), "stat")
		contents, err := ioutil.ReadFile(dev_path)
		if err != nil {
			continue
		}
		line_chunks := strings.FieldsFunc(string(contents),
			func(r rune) bool {
				return unicode.IsSpace(r)
			})
		read, err := strconv.ParseFloat(line_chunks[2], 64)
		if err != nil {
			continue
		}
		write, err := strconv.ParseFloat(line_chunks[6], 64)
		if err != nil {
			continue
		}
		i := IOStat{
			Name:           d.Name(),
			SectorsRead:    read,
			SectorsWritten: write,
		}
		response[d.Name()] = i
	}
	return response
}

func GetIOStats(c *conf.ConfigurationManager) []IOStatDiff {
	response := make([]IOStatDiff, 0)
	entries, err := ioutil.ReadDir(BASE_PATH)
	if err != nil {
		return response
	}
	old_stats := readStats(c, entries)
	time.Sleep(1000 * time.Millisecond)
	new_stats := readStats(c, entries)
	for key, value := range new_stats {
		if old_value, exists := old_stats[key]; exists {
			d := IOStatDiff{
				Name:         value.Name,
				BytesRead:    (value.SectorsRead - old_value.SectorsRead) * 512.0,
				BytesWritten: (value.SectorsWritten - old_value.SectorsWritten) * 512.0,
			}
			response = append(response, d)
		}
	}
	return response
}
