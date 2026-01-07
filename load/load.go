package load

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetCPULoad() (one, five, fifteen float64) {
	contents, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return
	}
	fields := strings.Fields(string(contents))
	one, _ = strconv.ParseFloat(fields[0], 64)
	five, _ = strconv.ParseFloat(fields[1], 64)
	fifteen, _ = strconv.ParseFloat(fields[2], 64)
	return
}

func main() {
	one, five, fifteen := GetCPULoad()
	fmt.Printf("CPU load is %f %f %f\n", one, five, fifteen)
}
