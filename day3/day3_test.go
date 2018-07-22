package day3

import "testing"

type TestExpecation struct {
	input          int
	expectedOutput int
}

func TestRun(t *testing.T) {

	runs := []TestExpecation{
		{1, 0},
		{2, 1},
		{4, 1},
		{5, 2},
	}

	for _, run := range runs {
		result := Run(run.input)

		if result != run.expectedOutput {
			t.Errorf("Result was incorrect, for %d, got: %d, want: %d.", run.input, result, run.expectedOutput)
		}
	}

}
