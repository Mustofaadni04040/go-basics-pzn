package main

import "fmt"

func IfExpressions() {
	name := "Mustofa Adny"

	if name == "Mustofa Adny" {
		fmt.Println("Hello, ", name)
	} else if name == "Ucok" {
		fmt.Println("Hello, ucok")
	} else {
		fmt.Println("Hello, unknown")
	}

	// Short statement
	if length := len(name); length > 10 {
		fmt.Println("Name is too long, length:", length)
	} else {
		fmt.Println("Name is acceptable, length:", length)
	}
}