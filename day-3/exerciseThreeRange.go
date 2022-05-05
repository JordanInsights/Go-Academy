package day3

import (
	"fmt"
)

func ExerciseThree() {
	fmt.Println("\033[33m", "\nOutput from exercise 3; Range over a Channel: ", "\033[0m")
	c := gen()
	var values []int = receive(c)

	fmt.Println("Values from channel: ", values)
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) []int {
	channelSlice := []int{}
	// RANGE BLOCKS UNTIL THE CHANNEL IS CLOSED
	for v := range c {
		channelSlice = append(channelSlice, v)
	}
	return channelSlice
}
