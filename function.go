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

// function as a value
func getGoodBye(name string) string {
	return "Good Bye " + name
}

func returnGoodBye() {
	goodBye := getGoodBye

	fmt.Println(goodBye("Ucok"))
}

type FilterType func(string) string // type declaration

// function as a parameter
func sayHelloWithFilter(name string, filter FilterType) {
	filtredName := filter(name)

	fmt.Println("Hello", filtredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "*****"
	} else {
		return name
	}
}

func sayHelloWithFilterExample() {
	sayHelloWithFilter("Ucok", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter) 

	filter := spamFilter
	sayHelloWithFilter("Budi", filter) 
}

// anonymous function
type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blacklisted:", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func user() {
	blacklist := func(name string) bool {
		return name == "Anjing" || name == "Babi"
	}

	registerUser("Ucok", blacklist)
	registerUser("Anjing", blacklist)
	registerUser("Babi", blacklist)
}
