package main

import "fmt"

// import (
// 	"intro-to-go/utils"
// )

var a = 1

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	// utils.Intro()
	// utils.Function()
	// utils.Scope()
	// utils.TypesOfFunction()

	add(2, 3)
	add(a, 2)
}

// func init() {
// 	fmt.Println("init function called", a)
// 	a := 2
// 	fmt.Println(a)
// }
