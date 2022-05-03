package main

import (
	"fmt"
)

func main() {
	ShortDeclarationOperator()

	fmt.Println("\n'Zero Value' variables with package level scope, for integers, strings, and booleans:")
	fmt.Println("X (int): ", X)
	fmt.Println("Y (string): ", Y)
	fmt.Println("Z (bool): ", Z)

	fmt.Println("\nOutput of the FizzBuzz task: ")
	FizzBuzz()

	fmt.Println("\nOutput of AppendToSlice task: ")
	AppendToSlice()

	fmt.Println("\nInitialise two VALUES of TYPE person, print the values")
	PersonStruct()
}
