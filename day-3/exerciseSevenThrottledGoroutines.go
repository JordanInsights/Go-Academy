package day3

import (
	"fmt"
)

func ExerciseSeven() {
	const goroutines int = 10

	fmt.Println("\033[33m", "\nExercise Seven Output: ", "\033[0m")
	c := make(chan int)

	for i := 0; i < goroutines; i++ {
		go func(c2 chan int) {
			for i2 := 0; i2 < 2; i2++ {
				// v := rand.Intn(500)
				v := i2
				c2 <- v
			}
			// cancel()
		}(c)
	}

	for i3 := range c {
		fmt.Println(i3)
	}

	close(c)
}
