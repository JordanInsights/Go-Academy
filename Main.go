package main

import (
	"fmt"
	day1 "go-academy/day-1"
	day2 "go-academy/day-2"
)

func main() {
	fmt.Println("Which day's code would you like to run?")
	var x int
	fmt.Scanf("%d", &x)

	switch x {
	case 1:
		fmt.Println("\nRunning the Day 1 code...")
		day1.Day1()
	case 2:
		fmt.Println("\nRunning the Day 2 code...")
		day2.Day2()
	default:
		fmt.Println("\nApologies, there's no code for the day you entered :(")
	}
}
