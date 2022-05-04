package day2

import "fmt"

func closure() {
	fmt.Println("\033[33m", "\nClosure Exercise: ", "\033[0m")

	a := doubler(1)
	b := doubler(1)

	fmt.Println("First call of a: ", a())
	fmt.Println("Second call of a: ", a())
	fmt.Println("Third call of a: ", a())

	fmt.Println("First call of b: ", b())
}

func doubler(x int) func() int {
	total := x
	return func() int {
		total = total * 2
		return total
	}
}
