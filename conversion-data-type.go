package main

import "fmt"

func Conversion() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println("Nilai 32 = ", nilai32) // Output: 32768
	fmt.Println("Nilai 64 = ", nilai64) // Output: 32768
	fmt.Println("Nilai 16 = ", nilai16) // Output: -32768 (overflow occurs when converting from int32 to int16)

	var name = "Mustofa Adny"
	var e uint8= name[0]
	var string = string(e)

	fmt.Println("Huruf pertama:", e) // Output: 77 (ASCII value of 'M')
	fmt.Println("Huruf pertama:", string) // Output: M
}