package main

import "fmt"

func sayHello() {
	fmt.Println("Hello, World!")
}

func paramater(name string) {
	fmt.Println("Hello", name)
}

// Function with return value
func getHello(name string) string {
	return "Hello" + name
}

func result() {
	name := "Mustofa"
	result := getHello(name)

	fmt.Println(result)
}

// mutiple return values
func returnMultipleValues() (string, string) {
	return "Mustofa", "Adny"
}

func getFullName() {
	firstName, _ := returnMultipleValues() // : _ ignores the return value
	fmt.Println("Full Name:", firstName)
}

