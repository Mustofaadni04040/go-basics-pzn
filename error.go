package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	} else {
		return nilai / pembagi, nil
	}
}

func ExampleMain2() {
	hasil, err := Pembagian(100, 10)

	// error checking
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err.Error())
	}
}