package main

import (
	"fmt"
)

var a = 1

func main() {
	// utils.Intro()
	// utils.Function()
	// utils.Scope()

	// age := 20

	// if age >= 18 {
	// 	a := 3
	// 	fmt.Println(a)
	// }

	fmt.Println("Main Function called", a)

	func() {
		fmt.Println("Anonymous Function called", a)
	}()

	(func(num int) {
		fmt.Println("IIFE Function called", num)
	})(100)
}

func init() {
	fmt.Println("init function called", a)
	a := 2
	fmt.Println(a)
}
