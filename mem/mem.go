package mem

import (
	"fmt"
	"github.com/TomasTomecek/tmux-top/humanize"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func GetMemStats() (used, total float64) {
	var free, buffers, cached, available float64
	contents, err := os.ReadFile("/proc/meminfo")
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
			total, _ = humanize.Dehumanize(total, line_chunks[2])
		} else if line_chunks[0] == "MemAvailable" {
			available, _ = strconv.ParseFloat(line_chunks[1], 64)
			available, _ = humanize.Dehumanize(available, line_chunks[2])
		} else if line_chunks[0] == "MemFree" {
			free, _ = strconv.ParseFloat(line_chunks[1], 64)
			free, _ = humanize.Dehumanize(free, line_chunks[2])
		} else if line_chunks[0] == "Buffers" {
			buffers, _ = strconv.ParseFloat(line_chunks[1], 64)
			buffers, _ = humanize.Dehumanize(buffers, line_chunks[2])
		} else if line_chunks[0] == "Cached" {
			cached, _ = strconv.ParseFloat(line_chunks[1], 64)
			cached, _ = humanize.Dehumanize(cached, line_chunks[2])
		}
	}
	used = total - available
	return
}

func main() {
	used, total := GetMemStats()
	fmt.Printf("U: %f, T: %f\n", used, total)
}
