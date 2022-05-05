package day3

import "fmt"

// This code uses a select statement to pull values
// off multiple channels, it will pull off whatever
// value is ready to be pulled off

func Select() {
	fmt.Println("\033[33m", "\nSelect Exemplar Output: ", "\033[0m")

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	// receive
	selectReceive(even, odd, quit)

	fmt.Println("Exiting")
}

func selectReceive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("From the even channel: ", v)
		case v := <-odd:
			fmt.Println("From the odd channel: ", v)
		case v := <-quit:
			fmt.Println("From the odd channel: ", v)
			return
		}
	}
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- 0
}
