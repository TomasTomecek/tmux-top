package conf

import (
	"fmt"
	"strings"
	"text/template"
)

/*
This ugly boilerplate code enables user to override default settings

it could be done by via dynamic retrieval of values from struct, but that
would feel like dynamic language:

http://stackoverflow.com/a/18931036/909579
*/

type ConfigurationManager struct {
	User    *Configuration
	Default *Configuration
}

func (c *ConfigurationManager) GetLoadIntervals() []IntervalDisplay {
	if c.User != nil {
		if c.User.Load != nil {
			if c.User.Load.Intervals != nil {
				return *c.User.Load.Intervals
			}
		}
	}
	return *c.Default.Load.Intervals
}

func (c *ConfigurationManager) GetMemIntervals() []IntervalDisplay {
	if c.User != nil {
		if c.User.Mem != nil {
			if c.User.Mem.Intervals != nil {
				return *c.User.Mem.Intervals
			}
		}
	}
	return *c.Default.Mem.Intervals
}

func (c *ConfigurationManager) GetMemSeparator() string {
	if c.User != nil {
		if c.User.Mem != nil {
			if c.User.Mem.Separator != nil {
				return *c.User.Mem.Separator
			}
		}
	}
	return *c.Default.Mem.Separator
}

func (c *ConfigurationManager) GetMemSeparatorBg() string {
	if c.User != nil {
		if c.User.Mem != nil {
			if c.User.Mem.SeparatorBg != nil {
				return *c.User.Mem.SeparatorBg
			}
		}
	}
	return *c.Default.Mem.SeparatorBg
}

func (c *ConfigurationManager) GetMemSeparatorFg() string {
	if c.User != nil {
		if c.User.Mem != nil {
			if c.User.Mem.SeparatorFg != nil {
				return *c.User.Mem.SeparatorFg
			}
		}
	}
	return *c.Default.Mem.SeparatorFg
}

func (c *ConfigurationManager) GetMemTotalBg() string {
	if c.User != nil {
		if c.User.Mem != nil {
			if c.User.Mem.TotalBg != nil {
				return *c.User.Mem.TotalBg
			}
		}
	}
	return *c.Default.Mem.TotalBg
}

func (c *ConfigurationManager) GetMemTotalFg() string {
	if c.User != nil {
		if c.User.Mem != nil {
			if c.User.Mem.TotalFg != nil {
				return *c.User.Mem.TotalFg
			}
		}
	}
	return *c.Default.Mem.TotalFg
}

func (c *ConfigurationManager) GetNetInterfaces() map[string]NetIFConfiguration {
	if c.User != nil {
		if c.User.Net != nil {
			if c.User.Net.Interfaces != nil {
				return *c.User.Net.Interfaces
			}
		}
	}
	return *c.Default.Net.Interfaces
}

func (c *ConfigurationManager) GetNetIntervals() []IntervalDisplay {
	if c.User != nil {
		if c.User.Net != nil {
			if c.User.Net.Intervals != nil {
				return *c.User.Net.Intervals
			}
		}
	}
	return *c.Default.Net.Intervals
}

func (c *ConfigurationManager) GetNetUpLabel() string {
	if c.User != nil {
		if c.User.Net != nil {
			if c.User.Net.UpLabel != nil {
				return *c.User.Net.UpLabel
			}
		}
	}
	return *c.Default.Net.UpLabel
}

func (c *ConfigurationManager) GetNetUpLabelBg() string {
	if c.User != nil {
		if c.User.Net != nil {
			if c.User.Net.UpLabelBg != nil {
				return *c.User.Net.UpLabelBg
			}
		}
	}
	return *c.Default.Net.UpLabelBg
}

func (c *ConfigurationManager) GetNetUpLabelFg() string {
	if c.User != nil {
		if c.User.Net != nil {
			if c.User.Net.UpLabelFg != nil {
				return *c.User.Net.UpLabelFg
			}
		}
	}
	return *c.Default.Net.UpLabelFg
}

func (c *ConfigurationManager) GetNetDownLabel() string {
	if c.User != nil {
		if c.User.Net != nil {
			if c.User.Net.DownLabel != nil {
				return *c.User.Net.DownLabel
			}
		}
	}
	return *c.Default.Net.DownLabel
}

func (c *ConfigurationManager) GetNetDownLabelBg() string {
	if c.User != nil {
		if c.User.Net != nil {
			if c.User.Net.DownLabelBg != nil {
				return *c.User.Net.DownLabelBg
			}
		}
	}
	return *c.Default.Net.DownLabelBg
}

func (c *ConfigurationManager) GetNetDownLabelFg() string {
	if c.User != nil {
		if c.User.Net != nil {
			if c.User.Net.DownLabelFg != nil {
				return *c.User.Net.DownLabelFg
			}
		}
	}
	return *c.Default.Net.DownLabelFg
}

func (c *ConfigurationManager) GetIODevices() map[string]IODeviceConfiguration {
	if c.User != nil {
		if c.User.IO != nil {
			if c.User.IO.Devices != nil {
				return *c.User.IO.Devices
			}
		}
	}
	return *c.Default.IO.Devices
}

func (c *ConfigurationManager) GetIOIntervals() []IntervalDisplay {
	if c.User != nil {
		if c.User.IO != nil {
			if c.User.IO.Intervals != nil {
				return *c.User.IO.Intervals
			}
		}
	}
	return *c.Default.IO.Intervals
}

func (c *ConfigurationManager) GetIOReadLabel() string {
	if c.User != nil {
		if c.User.IO != nil {
			if c.User.IO.ReadLabel != nil {
				return *c.User.IO.ReadLabel
			}
		}
	}
	return *c.Default.IO.ReadLabel
}

func (c *ConfigurationManager) GetIOReadLabelBg() string {
	if c.User != nil {
		if c.User.IO != nil {
			if c.User.IO.ReadLabelBg != nil {
				return *c.User.IO.ReadLabelBg
			}
		}
	}
	return *c.Default.IO.ReadLabelBg
}

func (c *ConfigurationManager) GetIOReadLabelFg() string {
	if c.User != nil {
		if c.User.IO != nil {
			if c.User.IO.ReadLabelFg != nil {
				return *c.User.IO.ReadLabelFg
			}
		}
	}
	return *c.Default.IO.ReadLabelFg
}

func (c *ConfigurationManager) GetIOWriteLabel() string {
	if c.User != nil {
		if c.User.IO != nil {
			if c.User.IO.WriteLabel != nil {
				return *c.User.IO.WriteLabel
			}
		}
	}
	return *c.Default.IO.WriteLabel
}

func (c *ConfigurationManager) GetIOWriteLabelBg() string {
	if c.User != nil {
		if c.User.IO != nil {
			if c.User.IO.WriteLabelBg != nil {
				return *c.User.IO.WriteLabelBg
			}
		}
	}
	return *c.Default.IO.WriteLabelBg
}

func (c *ConfigurationManager) GetIOWriteLabelFg() string {
	if c.User != nil {
		if c.User.IO != nil {
			if c.User.IO.WriteLabelFg != nil {
				return *c.User.IO.WriteLabelFg
			}
		}
	}
	return *c.Default.IO.WriteLabelFg
}

func (c *ConfigurationManager) GetDiskMounts() ([]string, error) {
	if c.User != nil {
		if c.User.Disk != nil {
			if c.User.Disk.Mounts != nil {
				return *c.User.Disk.Mounts, nil
			}
		}
	}
	return *c.Default.Disk.Mounts, nil
}

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func tmux_display(bg, fg string, value interface{}) (response string) {
	response = fmt.Sprintf("#[bg=%s,fg=%s]%v#[bg=default,fg=default]", bg, fg, value)
	return
}

func Init_template() template.Template {
	tmpl := template.New("")
	funcMap := template.FuncMap{
		"replace":      replace,
		"tmux_display": tmux_display,
	}
	tmpl.Funcs(funcMap)
	return *tmpl
}

// format -- CLI option
func (c *ConfigurationManager) GetSensorsTemplate(format string) template.Template {
	template_s := format
	if format == "" {
		template_s = *c.Default.Sensors.Template
		if c.User != nil {
			if c.User.Sensors != nil {
				if c.User.Sensors.Template != nil {
					template_s = *c.User.Sensors.Template
				}
			}
		}
	}
	template := Init_template()
	t, err := template.Parse(template_s)
	if err != nil {
		panic(err)
	}
	return *t
}

// format -- CLI option
func (c *ConfigurationManager) GetDiskTemplate(format string) template.Template {
	template_s := format
	if format == "" {
		template_s = *c.Default.Disk.Template
		if c.User != nil {
			if c.User.Disk != nil {
				if c.User.Disk.Template != nil {
					template_s = *c.User.Disk.Template
				}
			}
		}
	}
	template := Init_template()
	t, err := template.Parse(template_s)
	if err != nil {
		panic(err)
	}
	return *t
}
