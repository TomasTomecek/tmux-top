package conf

import (
	"fmt"
	"testing"
)

// func loadIntervalValue(s string) (float64, bool) {

func TestLoadIntervalValue1(t *testing.T) {
	d, r := loadIntervalValue("1.0")
	if d != 1.0 {
		t.Error(fmt.Sprintf("%f != 1.0", d))
	}
	if r {
		t.Error("Value should be absolute!")
	}
}

func TestLoadIntervalValue3(t *testing.T) {
	d, r := loadIntervalValue("1B")
	if d != 1.0 {
		t.Error(fmt.Sprintf("%f != 1.0", d))
	}
	if r {
		t.Error("Value should be absolute!")
	}
}

func TestLoadIntervalValue2(t *testing.T) {
	d, r := loadIntervalValue("10%")
	if d != 0.1 {
		t.Error(fmt.Sprintf("%f != 0.1", d))
	}
	if !r {
		t.Error(fmt.Errorf("Value '%f' should be relative!", d))
	}
}
