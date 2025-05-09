package utils

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func getNumbers (a int, b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}

func printSomething() {
	fmt.Println("Hello World")
}

func Intro() {
	// fmt.Println("Hello World")

	// a := 1
	// b := 2
	// fmt.Println(a + b)

	// int
	// float
	// string
	// bool
	// var c = 1
	// var c int = 1
	// c := 2
	// // var d = 1.1
	// // var d float64 = 1.1
	// d := 1.1
	// // var e = "test"
	// // var e string = "test"
	// e := "test"
	// // var f = true
	// // var f bool = true
	// f := true

	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Println(f)

	// * reassign
	// c = 3
	// fmt.Println(c)
	// d = 3.3
	// fmt.Println(d)
	// e = "test2"
	// fmt.Println(e)
	// f = false
	// fmt.Println(f)

	// age := 20
	// if age >= 18 {
	// 	fmt.Println("adult")
	// } else if age >= 10 {
	// 	fmt.Println("teenager")
	// } else {
	// 	fmt.Println("child")
	// }

	// <, <=, >, >=, ==, !=
	// and &&
	// or ||
	// not !

	// age := 19
	// if age >= 18 && age < 30 {
	// 	fmt.Println("adult")
	// }

	// if age >= 18 || age < 30 {
	// 	fmt.Println("adult")
	// }

	// if !(age == 20) {
	// 	fmt.Println("You are not 20")
	// }

	// age := 21
	// // switch case
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

	a := 1
	b := 2
	c := 3
	d := 4

	fmt.Println(sum(a, b))

	sum2 := sum(c, d)
	fmt.Println(sum2)

	sum, mul := getNumbers(a, b)
	fmt.Println(sum, mul)
	printSomething()

}