/*
docs: `strace sensors`
*/

package sens

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

var BASE_PATH string = "/sys/class/hwmon/"

type TemperatureStat struct {
	Label       string
	CurrentTemp float64
	// MaxTemp       float64
	// CritTemp      float64
	// CritAlarmTemp float64
}

// a list of stats for this device
type DeviceStat struct {
	Name      string            `json:"Name"`
	LowValue  float64           `json:"LowValue"`
	HighValue float64           `json:"HighValue"`
	Stats     []TemperatureStat `json:"Stats"`
}

type SensorsStats struct {
	Devices []DeviceStat
}

func getFileContentToInt(p string) (float64, error) {
	fd, err := ioutil.ReadFile(p)
	if err != nil {
		return 0.0, err
	}
	i, err := strconv.ParseInt(strings.TrimSpace(string(fd)), 10, 64)
	if err != nil {
		return 0.0, err
	}
	return float64(i) / 1000, nil
}

func getDeviceStats(f os.FileInfo) DeviceStat {
	sens_stats := make([]TemperatureStat, 0)
	d := DeviceStat{
		Stats:     sens_stats,
		Name:      "",
		LowValue:  1000000.0,
		HighValue: -100.0,
	}
	device_path, err := filepath.EvalSymlinks(path.Join(BASE_PATH, f.Name()))
	if err != nil {
		return d
	}

	name_f, err := ioutil.ReadFile(path.Join(device_path, "name"))
	name := ""
	if err == nil {
		name = strings.TrimSpace(string(name_f))
	}
	d.Name = name

	count := 1
	for {
		cur_p := path.Join(
			device_path,
			fmt.Sprintf("temp%s_input", strconv.Itoa(count)),
		)
		cur_f, err := getFileContentToInt(cur_p)
		if err != nil {
			break
		}

		label_f, err := ioutil.ReadFile(path.Join(
			device_path,
			fmt.Sprintf("temp%s_label", strconv.Itoa(count)),
		))
		label := ""
		if err == nil {
			label = strings.TrimSpace(string(label_f))
		}

		if d.LowValue > cur_f {
			d.LowValue = cur_f
		}
		if d.HighValue < cur_f {
			d.HighValue = cur_f
		}

		s := TemperatureStat{
			CurrentTemp: cur_f,
			Label:       label,
		}
		sens_stats = append(sens_stats, s)

		count++
	}
	d.Stats = sens_stats
	return d
}

func getSensorsStats() SensorsStats {
	device_stats := make([]DeviceStat, 0)
	s := SensorsStats{
		Devices: device_stats,
	}
	entries, err := ioutil.ReadDir(BASE_PATH)
	if err != nil {
		// If the hwmon directory doesn't exist (e.g., on s390x or in containers),
		// return empty stats instead of crashing
		if os.IsNotExist(err) {
			return s
		}
		panic(err)
	}
	// fd, os.FileInfo
	for _, e := range entries {
		d := getDeviceStats(e)
		device_stats = append(device_stats, d)
	}
	s.Devices = device_stats
	return s
}

func PrintSensorStats(t template.Template) {
	d := getSensorsStats()
	err := t.Execute(os.Stdout, d)
	if err != nil {
		panic(err)
	}
}
