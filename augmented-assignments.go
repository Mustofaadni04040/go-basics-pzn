package main

import "fmt"

func AugmentedAssignments() {
	var a int = 10
	a += 10
	fmt.Println(a)
	a -= 10
	fmt.Println(a)
	a *= 10
	fmt.Println(a)
	a /= 10
	fmt.Println(a)
}