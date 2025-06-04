package main

import "fmt"

// defer

func logging() {
	fmt.Println("finished calling function") // this function will be called after runApp function is finished
}

func runApp() {
	defer logging()
	fmt.Println("Running application")
}