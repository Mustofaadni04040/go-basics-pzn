package main

import "fmt"

type Address4 struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia1(address Address4) {
	address.Country = "Indonesia"
}

func ExampleWithoutPointer() {
	address := Address4{"Subang", "Jawa Barat", ""}
	ChangeAddressToIndonesia1(address) // {"Subang", "Jawa Barat", ""}

	fmt.Println(address) // not changed
}

func ChangeAddressToIndonesia2(address *Address4) {
	address.Country = "Indonesia"
}

func ExampleWithPointer() {
	address := Address4{"Subang", "Jawa Barat", ""}
	ChangeAddressToIndonesia2(&address) // {"Subang", "Jawa Barat", "Indonesia"}

	fmt.Println(address) // changed
}