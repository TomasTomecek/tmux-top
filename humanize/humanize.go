package humanize

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Dehumanize(num float64, unit string) (float64, error) {
	suffixes := map[string]int{"K": 1, "M": 2, "G": 3, "T": 4, "P": 5, "E": 6, "Z": 7}
	suffix := "B"
	var prefix string

	if strings.HasSuffix(unit, suffix) {
		prefix = unit[:len(unit)-1]
	} else {
		prefix = unit
	}
	prefix = strings.ToUpper(prefix)

	if s, ok := suffixes[prefix]; ok {
		value := num * math.Pow(1024.0, float64(s))
		return value, nil
	} else {
		return 0.0, fmt.Errorf("Unknown unit: '%s'", prefix)
	}
}

func DehumanizeString(value string) (float64, error) {
	r, _ := regexp.Compile(`([0-9.]+)\s*([a-zA-Z]+)`)
	r_simple, _ := regexp.Compile("([0-9.]+)")
	m := r.FindStringSubmatch(value)
	m_simple := r_simple.FindStringSubmatch(value)

	if len(m_simple) == 0 {
		return 0.0, fmt.Errorf("Invalid value: '%s'", value)
	}
	f, _ := strconv.ParseFloat(m_simple[0], 64)
	if len(m) == 3 {
		return Dehumanize(f, m[2])
	}
	return f, nil
}

func Absolutize(value string) (response float64, e error) {
	var v string
	suffix := "%"

	if strings.HasSuffix(value, suffix) {
		v = value[:len(value)-1]
	} else {
		e = fmt.Errorf("Missing percent sign in '%s'", value)
		return response, e
	}
	return strconv.ParseFloat(v, 64)
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
