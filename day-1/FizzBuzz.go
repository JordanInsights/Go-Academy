package main

import (
	"fmt"
)

func FizzBuzz() {
	fizzBuzzSlice := []string{}
	for i := 1; i <= 100; i++ {
		divisibleByThree := isDivisble(i, 3)
		divisibleByFive := isDivisble(i, 5)

		returnString := ""

		if divisibleByThree {
			returnString += "Fizz"
		}

		if divisibleByFive {
			returnString += "Buzz"
		}

		if len(returnString) == 0 {
			returnString = fmt.Sprint(i)
		}

		fizzBuzzSlice = append(fizzBuzzSlice, returnString)
	}

	fmt.Println(fizzBuzzSlice)
}

func isDivisble(numerator, denominator int) bool {
	zeroRemainder := numerator%denominator == 0
	return zeroRemainder
}
