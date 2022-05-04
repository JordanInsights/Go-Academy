package day2

import "fmt"

func summingFunction() {
	fooSum := foo(1, 2, 3)
	fmt.Println("Return value of foo: ", fooSum)

	x := []int{1, 2, 3}
	barSum := bar(x)
	fmt.Println("Return value of bar: ", barSum)
}

func foo(xi ...int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}

	return sum
}

func bar(x []int) int {
	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum
}
