package main

import "fmt"

func ComparisonOperations() {
	var a int = 10
	var b int = 20

	var result1 bool = a < b

	var name1 string = "Mustofa"
	var name2 string = "Adny"

	var result2 bool = name1 == name2

	fmt.Println("Is a less than b?", result1)
	fmt.Println("Is name1 equal to name2?", result2)
}