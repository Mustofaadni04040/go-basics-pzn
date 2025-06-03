package main

import "fmt"

func BooleanOperations() {
	var a bool = true
	var b bool = false

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!a)
	fmt.Println(!b)

	var nilaiAkhir int = 80
	var absensi int = 70

	var lulus bool = nilaiAkhir >= 80 && absensi >= 70
	fmt.Println("Apakah lulus?", lulus)

}