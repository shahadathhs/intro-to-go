package utils

import "fmt"

func Scope() {
	add(a, b)

	// c := 3
	// d := 4

	// add(a, b)
	// add(c, d)

	// add(a, c)
	// add(b, d)

	// age := 20
	// if age >= 18 {
	// 	fmt.Println("adult")
	// } else if age >= 10 {
	// 	fmt.Println("teenager")
	// } else {
	// 	fmt.Println("child")
	// }

	// switch age {
	// case 10:
	// 	fmt.Println("10")
	// case 20:
	// 	fmt.Println("20")
	// case 30:
	// 	fmt.Println("30")
	// default:
	// 	fmt.Println("default")
	// }
}

var (
	a = 1
	b = 2
)

func printNum(num int) {
	fmt.Println(num)
}

func add(x int, y int) {
	z := x + y
	printNum(z)
}
