package main

import (
	"math"
	"strconv"
)

//Write a program that prints the numbers from 1 to 100.
//But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”. F
//or numbers which are multiples of both three and five print “FizzBuzz “.

//FizzBuzz is excellent
func FizzBuzz() [100]string {

	var i int
	var result [100]string

	for i = 0; i < 100; i++ {
		var f = float64(i)
		if math.Mod(f, 3) == 0 && math.Mod(f, 5) == 0 {
			result[i] = "fizzbuzz"
		} else if math.Mod(f, 3) == 0 {
			result[i] = "fizz"
		} else if math.Mod(f, 5) == 0 {
			result[i] = "fizzbuzz"
		} else {
			result[i] = strconv.FormatFloat(f, 'f', -1, 64)
		}
	}

	return result
}
