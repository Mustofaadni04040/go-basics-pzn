package main

import "fmt"

func ArrayFunctions() {
	var values = [...]int{
		90,
		80,
		70,
	}

	fmt.Println(values)
	fmt.Println("Length:", len(values))
	values[0] = 100
	fmt.Println("Updated values:", values)
}