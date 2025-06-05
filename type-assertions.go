package main

import "fmt"

func Random() interface{} {
	return "OK"
}

func Assertion() {
	result := Random()
	resultString := result.(string) // conversion type from interface to string
	fmt.Println(resultString)

	resultInt := result.(int) // wrong : panic 
	fmt.Println(resultInt)
}

func AssertionSwitch() {
	result := Random()

	switch value := result.(type) { // ambil type return result
		case string:
			fmt.Println("String", value)
		case int:
			fmt.Println("Integer", value)
		default:
			fmt.Println("Unknown")
	}
}