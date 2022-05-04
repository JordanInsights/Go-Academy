package day1

import (
	"fmt"
)

func Day1() {
	ShortDeclarationOperator()

	fmt.Println("\n'Zero Value' variables with package level scope, for integers, strings, and booleans:")
	fmt.Println("X (int): ", X)
	fmt.Println("Y (string): ", Y)
	fmt.Println("Z (bool): ", Z)

	fmt.Println("\033[33m", "\nFizzBuzz Exercise:", "\033[0m")
	FizzBuzz()

	fmt.Println("\033[33m", "\nAppendToSlice Exercise:", "\033[0m")
	AppendToSlice()

	fmt.Println("\033[33m", "\nInitialise two VALUES of TYPE person Exercise:", "\033[0m")
	PersonStruct()

	fmt.Println("\033[33m", "\nAnonymous Struct Exercise:", "\033[0m")
	AnonymousStruct()
}
