package main

import "fmt"

func Ups() interface{} {
	return "Ups"
}

func Kosong() {
	kosong := Ups()

	fmt.Println(kosong)
}