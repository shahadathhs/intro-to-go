package utils

import "fmt"

func welcomeMessage() {
	fmt.Println("Welcome to function in utils package")
}

func getUserFullName() string {
	var name string
	var firstName string
	var lastName string
	fmt.Println("Please enter your first name")
	fmt.Scanln(&firstName)
	fmt.Println("Please enter your last name")
	fmt.Scanln(&lastName)
	name = firstName + " " + lastName
	return name
}

func getUserAge () int {
	var age int
	fmt.Println("Please enter your age")
	fmt.Scanln(&age)
	return age
}

func getUserAddress () string {
	var address string
	fmt.Println("Please enter your address")
	fmt.Scanln(&address)
	return address
}

func printUserData(name string, age int, address string) {
	fmt.Println("Here is your provided data")
	fmt.Println("Your full name is", name)
	fmt.Println("Your age is", age)
	fmt.Println("Your address is", address)
}

func goodbyeMessage() {
	fmt.Println("Thank you for using our service, Goodbye")
}

func Function() {
	welcomeMessage()
	
	name := getUserFullName()
	fmt.Println("Your full name is", name)

	age := getUserAge()
	fmt.Println("Your age is", age)

	address := getUserAddress()
	fmt.Println("Your address is", address)
	
	printUserData(name, age, address)

	goodbyeMessage()
}
