package display

import (
	"fmt"
	"testing"
)

func TestDisplayString(t *testing.T) {
	s := "banana"
	expected_response := fmt.Sprintf("#[bg=black,fg=white]%s#[bg=default,fg=default]", s)
	response := DisplayString(s, "black", "white")
	if response != expected_response {
		t.Error(fmt.Sprintf("Strings don't match: '%s' '%s'", response, expected_response))
	}
}

func TestDisplayFloat(t *testing.T) {
	f := 3.14159
	expected_response := fmt.Sprintf("#[bg=black,fg=white]%.2f#[bg=default,fg=default]", f)
	response := PrintFloat64(f, 2, "black", "white", false, "")
	if response != expected_response {
		t.Error(fmt.Sprintf("Strings don't match: '%s' '%s'", response, expected_response))
	}
}
