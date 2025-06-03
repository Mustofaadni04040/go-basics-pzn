package main

import "fmt"

func Boolean() {
	fmt.Println("True = ", true)
	fmt.Println("False = ", false)
	fmt.Println("True AND False = ", true && false)
	fmt.Println("True OR False = ", true || false)
	fmt.Println("NOT True = ", !true)
	fmt.Println("NOT False = ", !false)
	fmt.Println("True XOR False = ", true != false)
	fmt.Println("False XOR True = ", false != true)
}