package main

import (
	"belajar-golang-dasar/helper"
	_ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	Slice()
	SliceFunctions()

	result := helper.SayHello("Mustofa Adny")
	fmt.Println(result)
}