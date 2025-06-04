package main

import "fmt"

func sayHello() {
	fmt.Println("Hello, World!")
}

func paramater(name string) {
	fmt.Println("Hello", name)
}

// Function with return value
func getHello(name string) string {
	return "Hello" + name
}

func result() {
	name := "Mustofa"
	result := getHello(name)

	fmt.Println(result)
}

