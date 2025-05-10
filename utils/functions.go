package utils

import "fmt"

func TypesOfFunction() {
	a := 1

	// func() {
	// 	fmt.Println("Anonymous Function called", a)
	// }()

	// (func(num int) {
	// 	fmt.Println("IIFE Function called", num)
	// })(100)

	f := func() {
		fmt.Println("Anonymous Function called using function expression", a)
	}

	f()
}
