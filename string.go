package main

import "fmt"

func String(name string) {
	fmt.Println("Hello", name)
	fmt.Println("Panjang string:", len(name))
	fmt.Println("Huruf pertama:", name[0]) // byte type
}