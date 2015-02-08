// code for displaying stuff in tmux-aware format

package tmux_display

import (
	"fmt"
	"github.com/TomasTomecek/tmux-top/conf"
	"math"
	"strconv"
)

func FormatNicely(num float64) string {
	suffixes := []string{"", "K", "M", "G", "T", "P", "E", "Z"}
	suffix := "B"
	for _, unit := range suffixes {
		if math.Abs(num) < 1024.0 {
			return fmt.Sprintf("%3.1f%s%s", num, unit, suffix)
		}
		num = num / 1024.0
	}
	return fmt.Sprintf("%.1f%s%s", num, "Yi", suffix)
}

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
	for _, v := range intervals {
		if math.IsNaN(v.To) && math.IsNaN(v.From) {
			return PrintFloat64(value, precision, v.Bg_color, v.Fg_color)
		} else if math.IsNaN(v.From) && value < v.To {
			return PrintFloat64(value, precision, v.Bg_color, v.Fg_color)
		} else if math.IsNaN(v.To) && v.From <= value {
			return PrintFloat64(value, precision, v.Bg_color, v.Fg_color)
		} else if v.From <= value && value < v.To {
			return PrintFloat64(value, precision, v.Bg_color, v.Fg_color)
		}
	}
	return PrintFloat64(value, precision, "", "")
}
