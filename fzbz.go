package fzbz

import (
	"fmt"
)

// FizzBuzz parse int to FizzBuzz string:
//
//	FizzBuzz(1)  = "1"
//	FizzBuzz(3)  = "Fizz"
//	FizzBuzz(5)  = "Buzz"
//	FizzBuzz(15) = "FizzBuzz"
func FizzBuzz(num int) string {
	if num%15 == 0 {
		return "FizzBuzz"
	}

	if num%3 == 0 {
		return "Fizz"
	}

	if num%5 == 0 {
		return "Buzz"
	}

	return fmt.Sprint(num)
}
