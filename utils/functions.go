package utils

import "fmt"

func total(x int, y int) {
	fmt.Println(x + y)
}

func higherOrderFunction(x int, y int, op func(int, int)) {
	op(x, y)
}

func callFunction() {
	func() {
		fmt.Println("Anonymous Function called")
	}()
}

// func add(x int, y int) {
// 	z := x + y
// 	fmt.Println(z)
// }

func TypesOfFunction() {
	a := 1

	// func() {
	// 	fmt.Println("Anonymous Function called", a)
	// }()

	// (func(num int) {
	// 	fmt.Println("IIFE Function called", num)
	// })(100)

	// f := func() {
	// 	fmt.Println("Anonymous Function called using function expression", a)
	// }

	// f()

	// total(a, 2)

	higherOrderFunction(a, 2, total)

	callFunction()

	// add(2, 3)
	// add(a, 2)
}
