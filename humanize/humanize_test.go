package humanize

import (
	"fmt"
	"testing"
)

// func Dehumanize(num float64, unit string) (float64, error) {

func TestDehumanize1(t *testing.T) {
	d, _ := Dehumanize(1.0, "KB")
	if d != 1024.0 {
		t.Error(fmt.Sprintf("%f != 1024.0", d))
	}
}

func TestDehumanize2(t *testing.T) {
	d, _ := Dehumanize(1.0, "K")
	if d != 1024.0 {
		t.Error(fmt.Sprintf("%f != 1024.0", d))
	}
}

func TestDehumanize3(t *testing.T) {
	d, _ := Dehumanize(1.0, "MB")
	if d != 1048576.0 {
		t.Error(fmt.Sprintf("%f != 1048576.0", d))
	}
}

// func DehumanizeString(value string) (float64, error) {

func TestDehumanizeString1(t *testing.T) {
	d, _ := DehumanizeString("1.0MB")
	if d != 1048576.0 {
		t.Error(fmt.Sprintf("%f != 1048576.0", d))
	}
}

func TestDehumanizeString2(t *testing.T) {
	d, _ := DehumanizeString("1.0 M")
	if d != 1048576.0 {
		t.Error(fmt.Sprintf("%f != 1048576.0", d))
	}
}

func TestDehumanizeString3(t *testing.T) {
	d, _ := DehumanizeString("1.0")
	if d != 1.0 {
		t.Error(fmt.Sprintf("%f != 1.0", d))
	}
}

func TestDehumanizeString4(t *testing.T) {
	d, _ := DehumanizeString("512kB")
	if d != 524288 {
		t.Error(fmt.Sprintf("%f != 524288", d))
	}
}

// func Absolutize(value string) (response float64, e error) {

func TestAbsolutize1(t *testing.T) {
	d, _ := Absolutize("0.1%")
	if d != 0.1 {
		t.Error(fmt.Sprintf("%f != 0.1", d))
	}
}

func TestAbsolutizeE2(t *testing.T) {
	_, e := Absolutize("0.1")
	if e == nil {
		t.Error("'0.1' should raise an error")
	}
}
