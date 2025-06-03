package main

import "fmt"

type NoKTP string
type Married bool

func TypeDeclarations() {
	var ktpMustofa NoKTP = "12987123123"
	var married Married = false
	fmt.Println("No KTP Mustofa:", ktpMustofa)
	fmt.Println("Married status?", married)
}