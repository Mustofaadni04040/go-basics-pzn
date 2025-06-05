package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married1() {
	man.Name = "Mr. " + man.Name
}

func ExamplewithoutPointer() {
	eko := Man{"Eko"}
	eko.Married1()
	fmt.Println(eko.Name) // {"Eko"}
}

func (man *Man) Married2() {
	man.Name = "Mr. " + man.Name // untuk method selalu direkomendasikan menggunakan pointer
}

func ExamplewithPointer() {
	eko := Man{"Eko"}
	eko.Married2()
	fmt.Println(eko.Name) // {"Mr. Eko"}
}