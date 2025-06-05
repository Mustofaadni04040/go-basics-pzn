package main

import "fmt"

type Address2 struct {
	City, Province, Country string
}

func Example() {
	address1 := Address2{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1  // pointer to address1
	address2.City = "Pamulang" // address1 and address2 city are changed to "Pamulang"
	fmt.Println(address1) // {"Pamulang", "Jawa Barat", "Indonesia"}
	fmt.Println(address2) // &{"Pamulang", "Jawa Barat", "Indonesia"}

	address2 = &Address2{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // {"Pamulang", "Jawa Barat", "Indonesia"}
	fmt.Println(address2) // {"Jakarta", "DKI Jakarta", "Indonesia"}
}

func ExampleAsterisk(){ 
	address1 := Address2{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1  // pointer to address1
	address2.City = "Pamulang" // address1 and address2 city are changed to "Pamulang"
	fmt.Println(address1) // {"Pamulang", "Jawa Barat", "Indonesia"}
	fmt.Println(address2) // &{"Pamulang", "Jawa Barat", "Indonesia"}

	*address2 = Address2{"Jakarta", "DKI Jakarta", "Indonesia"} // semua yang mengacu pada address1 akan berubah ke address yang diberi asterisk operator
	fmt.Println(address1) // {"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address2) // {"Jakarta", "DKI Jakarta", "Indonesia"}
}