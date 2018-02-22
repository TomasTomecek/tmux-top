// code for displaying stuff in tmux-aware format

package display

import (
	"fmt"
	"github.com/TomasTomecek/tmux-top/conf"
	huma "github.com/TomasTomecek/tmux-top/humanize"
	"math"
	"strconv"
)

func DisplayString(value string, bg_color, fg_color string) (response string) {
	response = fmt.Sprintf("#[bg=%s,fg=%s]%s", bg_color, fg_color, value)
	return
}

func PrintFloat64(value float64, precision int, bg_color, fg_color string, humanize bool, suffix string) (response string) {
	var float_str string
	if humanize {
		float_str = huma.Humanize(value, precision, suffix)
	} else {
		float_str = strconv.FormatFloat(value, 'f', precision, 64)
	}
	if bg_color == "" && fg_color == "" {
		response = float_str
	} else if bg_color == "" {
		response = fmt.Sprintf("#[fg=%s]%s", fg_color, float_str)
	} else if fg_color == "" {
		response = fmt.Sprintf("#[bg=%s]%s", bg_color, float_str)
	} else {
		response = fmt.Sprintf("#[bg=%s,fg=%s]%s", bg_color, fg_color, float_str)
	}
	return
}

func DisplayFloat64(value float64, precision int, intervals []conf.IntervalDisplay, humanize bool, suffix string, max_value float64) string {
	for _, v := range intervals {
		from, is_from_relative := v.GetFrom()
		if is_from_relative {
			from = max_value * from
		}
		to, is_to_relative := v.GetTo()
		if is_to_relative {
			to = max_value * to
		}

		if math.IsNaN(to) && math.IsNaN(from) {
			return PrintFloat64(value, precision, v.BgColor, v.FgColor, humanize, suffix)
		} else if math.IsNaN(from) && value < to {
			return PrintFloat64(value, precision, v.BgColor, v.FgColor, humanize, suffix)
		} else if math.IsNaN(to) && from <= value {
			return PrintFloat64(value, precision, v.BgColor, v.FgColor, humanize, suffix)
		} else if from <= value && value < to {
			return PrintFloat64(value, precision, v.BgColor, v.FgColor, humanize, suffix)
		}
	}
	return ""
}
