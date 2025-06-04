package main

import "fmt"

func Closure() {
	counter := 0

	increment := func() {
		fmt.Println("Incrementing counter")
		counter++
	}

	increment()
	increment()

	fmt.Println("Counter = ", counter) // counter variable is changed by the closure function increment
}