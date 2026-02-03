package journal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

// convertTimeframe converts shorthand (1h, 5m) to journalctl standard format
// Converts "1h" -> "-1h", "5m" -> "-5m" (negative relative time)
// This is the standard systemd/journalctl format for relative time
func convertTimeframe(timeframe string) string {
	if !strings.HasPrefix(timeframe, "-") {
		return "-" + timeframe
	}
	return timeframe
}

type JournalStats struct {
	TimeFrame     string `json:"timeframe"`
	ErrorCount    int64  `json:"error_count"`
	WarningCount  int64  `json:"warning_count"`
	CriticalCount int64  `json:"critical_count"`
	TotalCount    int64  `json:"total_count"`
}

func GetJournalStats(timeframe string) (JournalStats, error) {
	stats := JournalStats{
		TimeFrame: timeframe,
	}

	// Get critical count (0-2: emerg, alert, crit)
	criticalCount, err := getJournalCountByPriorityRange(timeframe, "0", "2")
	if err != nil {
		return stats, err
	}
	stats.CriticalCount = criticalCount

	// Get error count (3 only: err)
	errorCount, err := getJournalCountByPriorityRange(timeframe, "3", "3")
	if err != nil {
		return stats, err
	}
	stats.ErrorCount = errorCount

	// Get warning count (4 only: warning)
	warningCount, err := getJournalCountByPriorityRange(timeframe, "4", "4")
	if err != nil {
		return stats, err
	}
	stats.WarningCount = warningCount

	stats.TotalCount = stats.ErrorCount + stats.WarningCount + stats.CriticalCount

	return stats, nil
}

func getJournalCountByPriorityRange(timeframe, fromPriority, toPriority string) (int64, error) {
	cmd := exec.Command("journalctl",
		"--since="+convertTimeframe(timeframe),
		"-p", fromPriority+".."+toPriority,
		"--no-pager",
		"-q",
		"--output=cat")

	output, err := cmd.Output()
	if err != nil {
		// If journalctl fails, return 0 instead of error for graceful degradation
		return 0, nil
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return 0, nil
	}

	return int64(len(lines)), nil
}

func GetJournalErrorCount(timeframe string) (int64, error) {
	stats, err := GetJournalStats(timeframe)
	if err != nil {
		return 0, err
	}
	return stats.ErrorCount, nil
}

func PrintJournalStats(tmpl template.Template, timeframe string) {
	stats, err := GetJournalStats(timeframe)
	if err != nil {
		fmt.Printf("Journal unavailable")
		return
	}

	err = tmpl.Execute(os.Stdout, stats)
	if err != nil {
		panic(err)
	}
}
