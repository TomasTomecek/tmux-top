package io

import (
	//"fmt"
	//"net"
	"github.com/TomasTomecek/tmux-top/conf"
	"io/ioutil"
	"testing"
)

func TestReadEntries(t *testing.T) {
	entries, err := ioutil.ReadDir(BASE_PATH)
	if len(entries) <= 0 {
		t.Error("no devices in /sys/block")
	}
	if err != nil {
		t.Error("error during reading /sys/block")
	}
	c := &conf.ConfigurationManager{
		Default: &conf.Configuration{
			IO: &conf.IOConfiguration{
				Devices: &map[string]conf.IODeviceConfiguration{
					entries[0].Name(): conf.IODeviceConfiguration{
						Alias: "Banana",
					}}}}}

	stats := readStats(c, entries)
	if len(stats) <= 0 {
		t.Error("no stats for devices")
	}
}

func TestGetIOStats(t *testing.T) {
	entries, err := ioutil.ReadDir(BASE_PATH)
	if len(entries) <= 0 {
		t.Error("no devices in /sys/block")
	}
	if err != nil {
		t.Error("error during reading /sys/block")
	}
	c := &conf.ConfigurationManager{
		Default: &conf.Configuration{
			IO: &conf.IOConfiguration{
				Devices: &map[string]conf.IODeviceConfiguration{
					entries[0].Name(): conf.IODeviceConfiguration{
						Alias: "Banana",
					}}}}}

	stats := GetIOStats(c)
	if len(stats) <= 0 {
		t.Error("no stats for devices")
	}
}
