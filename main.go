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
}

func init() {
	fmt.Println("init function called", a)
	a := 2
	fmt.Println(a)
}
