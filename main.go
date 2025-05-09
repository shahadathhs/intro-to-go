package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	a := 1
	b := 2
	fmt.Println(a + b)

	// int
	// float
	// string
	// bool
	// var c = 1
	// var c int = 1
	c := 2
	// var d = 1.1
	// var d float64 = 1.1
	d := 1.1
	// var e = "test"
	// var e string = "test"
	e := "test"
	// var f = true
	// var f bool = true
	f := true

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	// * reassign
	c = 3
	fmt.Println(c)
	d = 3.3
	fmt.Println(d)
	e = "test2"
	fmt.Println(e)
	f = false
	fmt.Println(f)
}