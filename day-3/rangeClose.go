package day3

import "fmt"

func RangeClose() {
	fmt.Println("\033[33m", "\nRange Exemplar Output: ", "\033[0m")

	c := make(chan int)

	// send
	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}
		// need to close c other will continue to wait after loop is done and result in deadlock
		close(c)
	}()

	var testingSlice = []int{}

	// receive
	// lack of a go here means that IT WILL BLOCK UNTIL IT HAS COMPLETED IT's JOB
	// means you don't need a wait group or anything
	for v := range c {
		testingSlice = append(testingSlice, v)
	}

	fmt.Println("Resulting slice: ", testingSlice)
}
