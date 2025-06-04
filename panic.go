package main

import "fmt"

func endApp() {
	fmt.Println("End Application")
}

func runApplication(error bool) {
	defer endApp() // defer will call endApp function after runApplication is finished, even if there is a panic

	if error {
		panic("Application Error") // panic will stop the application and call defer functions
	}
}