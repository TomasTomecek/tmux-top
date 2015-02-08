package mem

import (
	"fmt"
	display "github.com/TomasTomecek/tmux-top/display"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func GetMemStats() (used, total float64) {
	var free float64
	contents, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return
	}
	lines := strings.Split(strings.TrimSpace(string(contents)), "\n")
	for _, line := range lines {
		line_chunks := strings.FieldsFunc(line,
			func(r rune) bool {
				return r == ':' || unicode.IsSpace(r)
			})
		if line_chunks[0] == "MemTotal" {
			total, _ = strconv.ParseFloat(line_chunks[1], 64)
			// no idea why /proc/meminfo outputs stuff in kilobytes
			// what is this, 1990?
			total = display.Dehumanize(total, line_chunks[2])
		}
		if line_chunks[0] == "MemFree" {
			free, _ = strconv.ParseFloat(line_chunks[1], 64)
			free = display.Dehumanize(free, line_chunks[2])
		}
	}
	used = total - free
	return
}

func main() {
	used, total := GetMemStats()
	fmt.Printf("U: %f, T: %f\n", used, total)
}
