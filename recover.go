package main

import "fmt"

func endApps() {
	fmt.Println("End Application")
	message := recover() // recover for goodRecover function
	fmt.Println("Something happened:", message)
}

// wrong way to recover
func wrongRecover(error bool) {
	defer endApp()

	if error {
		panic("Application Error")
	}

	message := recover() // recover will catch the panic and return the message
	fmt.Println("Something happened:", message)
}

// good way to recover
func goodRecover(error bool) {
	defer endApps()

	if error {
		panic("Application Error") // panic will stop the application and call defer functions
	}
}