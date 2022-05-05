package day3

import (
	"fmt"
	"time"
)

type operation struct {
	action    string
	parameter string
	response  chan string
}

var requests chan operation = make(chan operation)

func OrderCoffee(coffeeType string) {
	op := operation{
		action:    "order",
		parameter: coffeeType,
	}

	requests <- op
}

func Start() {
	go monitorRequests()
}

func monitorRequests() {
	// our protected data
	var lastCoffeeMade string

	for op := range requests {
		switch op.action {
		case "order":
			requestedCoffee := op.parameter
			makeCoffee(requestedCoffee)
			lastCoffeeMade = requestedCoffee
		case "lastmade":
			op.response <- lastCoffeeMade
		}
		// To Be Implemented
		fmt.Println(op)
	}

	close(done)
}

func makeCoffee(coffee string) {
	fmt.Println("Brewing " + coffee)
	time.Sleep(2 * time.Second)
	fmt.Println(coffee + " now ready")
}

func GetLastCofeeMade() string {
	answer := make(chan string)

	op := operation{
		action:    "lastmade",
		parameter: "",
		response:  answer,
	}

	requests <- op
	return <-answer
}

var done chan struct{} = make(chan struct{})

func Stop() {
	close(requests)
	<-done
}

func ActorModel() {
	Start()
	defer Stop()

	go OrderCoffee("Mocha")
	go OrderCoffee("Choca")
	go fmt.Println("Last made answer: " + GetLastCofeeMade())
	go OrderCoffee("Tea, milk, no sugar!")
	go fmt.Println("Last made answer: " + GetLastCofeeMade())
	go OrderCoffee("Banana Shake")
	go OrderCoffee("Strawberry Shake")
	go fmt.Println("Last made answer: " + GetLastCofeeMade())

	time.Sleep(100 * time.Millisecond)
}
