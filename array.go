package main

import "fmt"

func Array() {
	var names [2]string
	names[0] = "Mustofa"
	names[1] = "Adny"

	fmt.Println("Aray of first name:", names[0])
	fmt.Println("Array of last name:", names[1])

	var values = [5]int {1,2,3,4,5}
	fmt.Println("Array of values:", values)
}