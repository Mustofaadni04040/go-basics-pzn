package main

import "fmt"

func Variable() {
	var name string
	name = "Mustofa Adny"

	age := 30 // short variable declaration without var and type

	fmt.Println("Hello", name)
	fmt.Println("Umur", age)

	name = "Ucok Udin"
	fmt.Println("Hello", name)

	// Multiple variable declaration
	var (
		firsName = "Mustofa"; 
		lastName = "Adny"
	)
	fmt.Println("First name", firsName)
	fmt.Println("Last name", lastName)
}