package journal

import (
	"github.com/TomasTomecek/tmux-top/conf"
	"testing"
)

func TestGetJournalStats(t *testing.T) {
	stats, err := GetJournalStats("1h")
	if err != nil {
		t.Logf("Warning: journalctl failed (may not be available): %v", err)
		return
	}

	// Validate structure
	if stats.TimeFrame != "1h" {
		t.Errorf("Expected timeframe '1h', got '%s'", stats.TimeFrame)
	}

	// Counts should be non-negative
	if stats.ErrorCount < 0 || stats.WarningCount < 0 || stats.CriticalCount < 0 {
		t.Errorf("Negative counts found: error=%d, warning=%d, critical=%d",
			stats.ErrorCount, stats.WarningCount, stats.CriticalCount)
	}

	// TotalCount should equal sum
	expected := stats.ErrorCount + stats.WarningCount + stats.CriticalCount
	if stats.TotalCount != expected {
		t.Errorf("TotalCount mismatch: got %d, expected %d", stats.TotalCount, expected)
	}
}

func TestPrintJournalStats(t *testing.T) {
	c := &conf.ConfigurationManager{
		Default: conf.LoadConfFromBytes([]byte(conf.GetDefaultConf())),
	}

	tmpl := c.GetJournalTemplate("")

	// Should not panic
	PrintJournalStats(tmpl, "1h")
}

func TestConvertTimeframe(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1h", "-1h"},
		{"5m", "-5m"},
		{"24h", "-24h"},
		{"-1h", "-1h"}, // Already negative, should not change
		{"-5m", "-5m"},
	}

	for _, tt := range tests {
		result := convertTimeframe(tt.input)
		if result != tt.expected {
			t.Errorf("convertTimeframe(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}
