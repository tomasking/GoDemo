package main

import "testing"

func TestFizzBuzz(t *testing.T) {

	result := FizzBuzz()

	if result[2] != "Fizz" {
		t.Error("Expected Fizz, got ", result[2])
	}
}

// func TestSum(t *testing.T) {
// 	total := Sum(5, 5)
// 	if total != 10 {
// 		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
// 	}
// }
