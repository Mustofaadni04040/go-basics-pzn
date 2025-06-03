package main

import "fmt"

func Map() {
	person := map[string]string{
		"name":    "Mustofa Adny",
		"address": "Jakarta",
	}

	fmt.Println("Name:", person["name"])
	fmt.Println("Address:", person["address"])
}