package day4

import "testing"

func TestRun(t *testing.T) {

	result := Run()

	if result != 455 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 455)
	}
}
