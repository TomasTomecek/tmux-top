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
