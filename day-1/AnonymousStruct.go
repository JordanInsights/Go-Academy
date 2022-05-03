package main

import (
	"fmt"
)

func AnonymousStruct() {
	p1 := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Miss",
		lastName:  "MoneyPenny",
		age:       27,
	}

	fmt.Println(p1)
}
