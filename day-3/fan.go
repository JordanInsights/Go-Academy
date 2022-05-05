package day3

// Fan In / Out Pattern
// This is a concurrent design pattern

// If you have a chunk of work of indeterminate size
// you can fan it out to as many goroutines as possible.
// They will all work concurrently

// Then, when the work gets results the results are
// fanned in to a single channel. So you end with all
// the results in one place

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func FanInFunc() {
	fmt.Println("\033[33m", "\nFan In Exemplar Output: ", "\033[0m")

	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go FanInSend(even, odd)
	go FanInReceive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

}

// send channel
func FanInSend(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(even)
	close(odd)
}

// receive channel
func FanInReceive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

// FAN OUT
func FanOutFunc() {
	fmt.Println("\033[33m", "\nFan Out Exemplar Output: ", "\033[0m")

	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("Exiting...")
}

func FanOutFuncThrottled() {
	fmt.Println("\033[33m", "\nThrottled Fan Out Exemplar Output: ", "\033[0m")

	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go throttledFanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("Exiting...")
}

func populate(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func throttledFanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWork(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
