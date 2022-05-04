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
	p1 := person{
		firstName: "Jordan",
		lastName:  "Shaw",
	}

	fmt.Println(p1)

	changeMe(&p1)
	fmt.Println(p1)
}
