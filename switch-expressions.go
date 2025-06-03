package main

import "fmt"

func Switch() {
	name := "Mustofa Adny"

	switch name {
	case "Mustofa Adny":
		fmt.Println("Hello, ", name)
	case "Ucok":
		fmt.Println("Hello, ucok")
	default:
		fmt.Println("Hello, unknown")
	}

	// Short statement
	switch length := len(name); length > 10 {
		case true:
			fmt.Println("Name is too long, length:", length)
		default:
			fmt.Println("Name is acceptable, length:", length)
	}

	length := len(name)
	// switch without condition
	switch {
		case length > 10:
			fmt.Println("Name is too long, length:", length)
		default:
			fmt.Println("Name is acceptable, length:", length)
	}
}