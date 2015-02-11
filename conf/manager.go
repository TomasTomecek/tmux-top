package conf

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
