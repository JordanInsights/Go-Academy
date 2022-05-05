package day3

import "fmt"

func BrokenChannelsExample() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}

func ChannelsExample() {
	c := make(chan string)

	go func() {
		c <- "Example One"
	}()

	fmt.Println(<-c)
}

func BufferChannelsExample() {
	c := make(chan string, 1)
	c <- "Successful Buffer Example"
	fmt.Println(<-c)
}

func UnsuccessfulBufferExample() {
	c := make(chan string, 1)
	c <- "This part would work"
	c <- "But this will break it as there's only a buffer of 1, no where to store the second value"
	fmt.Println(<-c)
}
