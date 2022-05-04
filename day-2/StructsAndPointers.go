package day2

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func changeMe(p *person) {
	p.firstName = "New"
	p.lastName = "Name"
}

func structsAndPointers() {
	fmt.Println("\033[33m", "\nStructs and Pointers Exercise: ", "\033[0m")

	p1 := person{
		firstName: "Jordan",
		lastName:  "Shaw",
	}

	fmt.Println(p1)

	changeMe(&p1)
	fmt.Println(p1)
}
