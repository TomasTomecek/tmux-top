// code for displaying stuff in tmux-aware format

package tmux_display

import (
	"conf"
	"fmt"
	"math"
	"strconv"
)

func DisplayString(value string, bg_color, fg_color string) (response string) {
	response = fmt.Sprintf("[bg=%s,fg=%s]%s[bg=default,fg=default]", bg_color, fg_color, value)
	return
}

func PrintFloat64(value float64, precision int, bg_color, fg_color string) (response string) {
	float_str := strconv.FormatFloat(value, 'f', precision, 64)
	if bg_color == "" && fg_color == "" {
		response = float_str
	} else if bg_color == "" {
		response = fmt.Sprintf("[fg=%s]%s[fg=default]", fg_color, float_str)
	} else if fg_color == "" {
		response = fmt.Sprintf("[bg=%s]%s[bg=default]", bg_color, float_str)
	} else {
		response = fmt.Sprintf("[bg=%s,fg=%s]%s[bg=default,fg=default]", bg_color, fg_color, float_str)
	}
	return
}

func DisplayFloat64(value float64, precision int, intervals []conf.IntervalDisplay) string {
	for i, v := range intervals {
		if math.IsNaN(v.to) && math.IsNaN(v.from) {
			return PrintFloat64(value, precision, v.bg_color, v.fg_color)
		} else if math.IsNaN(v.from) && value < v.to {
			return PrintFloat64(value, precision, v.bg_color, v.fg_color)
		} else if math.IsNaN(v.to) && v.from <= value {
			return PrintFloat64(value, precision, v.bg_color, v.fg_color)
		} else if v.from <= value && value < v.to {
			return PrintFloat64(value, precision, v.bg_color, v.fg_color)
		}
	}
	return PrintFloat64(value, precision, "", "")
}
