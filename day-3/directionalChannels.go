package day3

import "fmt"

func BrokenDirectionalChannels() {
	c := make(chan<- int, 2)

	c <- 42
	c <- 43

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T", c)
}

// Channel types read from left to right

func DirectionChannels() {
	fmt.Println("\033[33m", "\nDirectional Channels Exemplar code: ", "\033[0m")

	c := make(chan int)    // bi-directional
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // sned

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	//specific to general doesn't assign
	// i.e. trying to assign a recieve/send channel to a bi-directional doesn't work
	// c = cr
	// c = cs

	// specific to specific doesn't assign
	// i.e. trying to assign a recieve to a send doesn't work, and vice-versa

	// However, general to specific assigns
	// i.e. you can assign a bidirectional channel to a receive/send
	// cr = c
	// cs = c

	// general to specific converts
	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))
	fmt.Println("-----")
}

func DirectionChannelsUseCase() {
	c := make(chan int)

	// send
	go foo(c)

	// receive
	// lack of a go here means that IT WILL BLOCK UNTIL IT HAS COMPLETED IT's JOB
	// means you don't need a wait group or anything
	bar(c)

	fmt.Println("about to exit")
}

//send
func foo(c chan<- int) {
	c <- 42
}

//send
func bar(c <-chan int) {
	fmt.Println(<-c)
}
