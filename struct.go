package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func Struct() {
	customer1 := Customer{
		Name:    "Mustofa",
		Address: "Jl. Raya No. 1",
		Age:     30,
	}

	fmt.Println("Customer 1:", customer1)
}