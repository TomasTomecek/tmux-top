package load

import (
	"testing"
)

func TestGetCPULoad(t *testing.T) {
	one, _, _ := GetCPULoad()
	if one <= 0.0 {
		t.Error("Suspicous, one minute load _can't_ be 0.00")
	}
}
