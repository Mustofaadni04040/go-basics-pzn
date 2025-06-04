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

// Named return values
func getCompleteName() (firstName string, lastName string) {
	firstName = "Mustofa"
	lastName = "Adny"
	
	return firstName, lastName
}

func getFullNameWithNamedReturn() {
	firstName, lastName := getCompleteName()

	fmt.Println("Full Name with Named Return:", firstName, lastName)
}

// Function with variadic parameters
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func sumAllNumbers() {
	total := sumAll(10, 20, 30, 40, 50)

	fmt.Println("Total sum is:", total)
}

// slice parameters
func sumAllSlice() {
	numbers := []int{10, 20, 30, 40, 50}

	total := sumAll(numbers...) // using ... to pass slice as variadic parameters
	fmt.Println("Total sum from slice: ", total)
}
