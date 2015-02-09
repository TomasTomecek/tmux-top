// code for displaying stuff in tmux-aware format

package tmux_display

import (
	"fmt"
	"github.com/TomasTomecek/tmux-top/conf"
	"math"
	"strconv"
	"strings"
)

// FIXME: error checking is non-existent
func Dehumanize(num float64, unit string) float64 {
	suffixes := map[string]int{"K": 1, "M": 2, "G": 3, "T": 4, "P": 5, "E": 6, "Z": 7}
	suffix := "B"
	var prefix string

	if strings.HasSuffix(unit, suffix) {
		prefix = unit[:len(unit)-1]
	} else {
		prefix = unit
	}
	prefix = strings.ToUpper(prefix)

	return num * math.Pow(1024.0, float64(suffixes[prefix]))
}

func Humanize(num float64, precision int, suffix string) string {
	suffixes := []string{"", "K", "M", "G", "T", "P", "E", "Z"}
	if suffix == "" {
		suffix = "B"
	}
	format_str := fmt.Sprintf("%%3.%df%%s%%s", precision)
	for _, unit := range suffixes {
		if math.Abs(num) < 1024.0 {
			return fmt.Sprintf(format_str, num, unit, suffix)
		}
		num = num / 1024.0
	}
	return fmt.Sprintf("%.1f%s%s", num, "Yi", suffix)
}

func DisplayString(value string, bg_color, fg_color string) (response string) {
	response = fmt.Sprintf("#[bg=%s,fg=%s]%s#[bg=default,fg=default]", bg_color, fg_color, value)
	return
}

func PrintFloat64(value float64, precision int, bg_color, fg_color string, humanize bool, suffix string) (response string) {
	var float_str string
	if humanize {
		float_str = Humanize(value, precision, suffix)
	} else {
		float_str = strconv.FormatFloat(value, 'f', precision, 64)
	}
	if bg_color == "" && fg_color == "" {
		response = float_str
	} else if bg_color == "" {
		response = fmt.Sprintf("#[fg=%s]%s#[fg=default]", fg_color, float_str)
	} else if fg_color == "" {
		response = fmt.Sprintf("#[bg=%s]%s#[bg=default]", bg_color, float_str)
	} else {
		response = fmt.Sprintf("#[bg=%s,fg=%s]%s#[bg=default,fg=default]", bg_color, fg_color, float_str)
	}
	return
}

func DisplayFloat64(value float64, precision int, intervals []conf.IntervalDisplay, humanize bool, suffix string) string {
	for _, v := range intervals {
		if math.IsNaN(v.GetTo()) && math.IsNaN(v.GetFrom()) {
			return PrintFloat64(value, precision, v.BgColor, v.FgColor, humanize, suffix)
		} else if math.IsNaN(v.GetFrom()) && value < v.GetTo() {
			return PrintFloat64(value, precision, v.BgColor, v.FgColor, humanize, suffix)
		} else if math.IsNaN(v.GetTo()) && v.GetFrom() <= value {
			return PrintFloat64(value, precision, v.BgColor, v.FgColor, humanize, suffix)
		} else if v.GetFrom() <= value && value < v.GetTo() {
			return PrintFloat64(value, precision, v.BgColor, v.FgColor, humanize, suffix)
		}
	}
	return PrintFloat64(value, precision, "default", "default", humanize, suffix)
}
