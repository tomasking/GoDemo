package day5

import "testing"

func TestRunDay5(t *testing.T) {

	result := Run()
	expectedResult := 24568703 //351282
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, expectedResult)
	}
}
