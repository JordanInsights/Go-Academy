package day1

import (
	"fmt"
)

func ShortDeclarationOperator() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("Single Print Statement:", x, y, z)

	fmt.Println("\nMultiple Print Statements: ")
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)
}

var (
	X int
	Y string
	Z bool
)
