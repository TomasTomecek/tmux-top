package sens

import (
	"github.com/TomasTomecek/tmux-top/conf"
	"testing"
)

func TestReadStats(t *testing.T) {
	template := conf.Init_template()
	te, err := template.Parse("{{.|printf \"%#v\"}}")
	if err != nil {
		panic(err)
	}
	PrintSensorStats(*te)
}

func TestReadStatsDefaultTempl(t *testing.T) {
	c := &conf.ConfigurationManager{
		Default: conf.LoadConfFromBytes([]byte(conf.GetDefaultConf())),
	}
	PrintSensorStats(c.GetSensorsTemplate(""))
}

func TestMissingHwmon(t *testing.T) {
	// Test that getSensorsStats handles missing hwmon directory gracefully
	// This simulates s390x and other environments where hwmon doesn't exist
	oldPath := BASE_PATH
	BASE_PATH = "/nonexistent/hwmon/path/"
	defer func() { BASE_PATH = oldPath }()

	stats := getSensorsStats()
	if len(stats.Devices) != 0 {
		t.Errorf("Expected empty devices when hwmon is missing, got %d devices", len(stats.Devices))
	}
}
