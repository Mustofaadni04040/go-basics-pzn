package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func Example1() {
	// pass by value
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1 // copy value from address1

	address2.City = "Bandung"

	fmt.Println(address1) // address 1 is not changed
	fmt.Println(address2) // address 2 has own value

	// pass by reference : pointer
	address3 := &address1 // reference to address1: address3 doesn't copy value from address1 but reference to address1. has same value

	address3.City = "Pamulang"
	fmt.Println(address1)
	fmt.Println(address2) 
	fmt.Println(address3)
}