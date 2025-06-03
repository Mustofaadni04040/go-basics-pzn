package main

import "fmt"

func Constant() {
	const firstName string = "Mustofa"
	const lastName = "Adny"

	fmt.Println("First name:", firstName)
	fmt.Println("Last name:", lastName)

	const (
		title = "Learning Go"
		description = "Learning Go programming language basics"
	)

	fmt.Println("Title:", title)
	fmt.Println("Description:", description)
}