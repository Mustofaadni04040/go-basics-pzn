package main

import "fmt"

type Address3 struct {
	City, Province, Country string
}

func ExampleNew1() {
	var alamat1 *Address3 = new(Address3) // = &Address3{}
	var alamat2 *Address3 = alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1) // {"", "", "Indonesia"}
	fmt.Println(alamat2) // {"", "", "Indonesia"}
}