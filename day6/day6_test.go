package day6

import "testing"

func TestRunDay6(t *testing.T) {

	result := Run()
	expectedResult := 24568703 //351282
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, expectedResult)
	}
}
