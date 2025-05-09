package utils

import "fmt"

func Function() {
	// Welcome Message
	fmt.Println("Welcome to function in utils package")

	// Get User Input
	var name string
	fmt.Println("Please enter your name")
	fmt.Scanln(&name)
	fmt.Println("Hello", name)

	// Get User Input
	var age int
	fmt.Println("Please enter your age")
	fmt.Scanln(&age)
	fmt.Println("You are", age, "years old")
}
