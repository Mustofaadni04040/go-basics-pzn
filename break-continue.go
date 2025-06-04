package main

import "fmt"

func BreakContinue() {
	for i := range 10 {
		if i == 5 {
			break // totally break the iteration
		}

		fmt.Println("Counter = ", i)
	}

	for i := range 10 {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Odd Counter = ", i) // continue will skip the iteration if the condition is met: will skip even numbers
	}
}