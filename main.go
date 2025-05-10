package main

import "fmt"

// import (
// 	"intro-to-go/utils"
// )

const a = 1

var b = 2

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}

	add(2, 3)
	add(a, b)
}

func main() {
	// utils.Intro()
	// utils.Function()
	// utils.Scope()
	// utils.TypesOfFunction()
	a := 2

	fmt.Println(a)

	call()
}

// func init() {
// 	fmt.Println("init function called", a)
// 	a := 2
// 	fmt.Println(a)
// }
