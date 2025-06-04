package main

import "fmt"

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func testStructMethod() {
	customer := Customer{
		Name:    "Mustofa",
		Address: "Jl. Raya No. 1",
		Age:     30,
	}

	customer.sayHello("Ucok") // calling method sayHello on customer2
}