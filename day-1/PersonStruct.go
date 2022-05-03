package main

import "fmt"

type person struct {
	firstName         string
	lastName          string
	favouriteFlavours []string
}

func PersonStruct() {
	personOne := person{
		firstName:         "Jordan",
		lastName:          "Shaw",
		favouriteFlavours: []string{"Vanilla", "Salted Caramel"},
	}

	personTwo := person{
		firstName:         "Winston",
		lastName:          "Churchill",
		favouriteFlavours: []string{"Tobacco", "Whiskey"},
	}

	fmt.Println("Print values of Person One")
	fmt.Println(personOne)

	fmt.Println("\nPrint values of Person Two")
	fmt.Println(personTwo)

	fmt.Println("\nRange over Favourite Flavours of Person One")
	for _, v := range personOne.favouriteFlavours {
		fmt.Println(v)
	}

	fmt.Println("\nRange over Favourite Flavours of Person Two")
	for _, v := range personTwo.favouriteFlavours {
		fmt.Println(v)
	}

}