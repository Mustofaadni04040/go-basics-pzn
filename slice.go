package main

import "fmt"

func Slice() {
	// A slice is a dynamically-sized, flexible view into the elements of an array.
	// It is a descriptor of an array segment and provides a way to work with arrays without needing to know their size at compile time.

	names := [...]string{"Mustofa", "Adny", "Ucok", "Udin"}
	slice := names[1:3] // index : length array | var slice []string = names[1:3]

	fmt.Println(slice[0]) // Output: Adny
	fmt.Println(slice[1]) // Output: Ucok
	fmt.Println(slice) // Output: [Adny Ucok]
}