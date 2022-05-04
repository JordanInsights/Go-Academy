package day1

import (
	"fmt"
)

func AppendToSlice() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("\nAppend single value")
	x = append(x, 52)
	fmt.Println(x)

	fmt.Println("\nAppend multiple values")
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}

	// for i := 0; i < len(y); i++ {
	// 	x = append(x, y[i])
	// }

	// for _, num := range y {
	// 	x = append(x, num)
	// }

	// Most concise
	x = append(x, y...)

	fmt.Println("\nAppend a slice of values to a slice")
	fmt.Println(x)
}
