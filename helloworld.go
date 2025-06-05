package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	Slice()
	SliceFunctions()

	result := helper.SayHello("Mustofa Adny")
	fmt.Println(result)
}