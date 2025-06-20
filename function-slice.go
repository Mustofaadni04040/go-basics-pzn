package main

import "fmt"

func SliceFunctions() {
	var iniArray  = [...]string{"Satu", "Dua", "Tiga", "Empat", "Lima"}
	var iniSlice = []string {"Satu", "Dua", "Tiga", "Empat", "Lima"}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days) // [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	daySlice2 := append(daySlice1, "Libur Baru") // menambahkan elemen baru ke dalam array days, yang mana array di go tidak bisa diubah ukurannya, maka akan membuat array baru di internal
	daySlice2[0] = "Ups" 
	fmt.Println(daySlice2) // [Ups Minggu Baru Libur Baru]
	fmt.Println(days) // [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	// make slice
	var newSlice []string = make([]string, 2, 5) // membuat slice dengan panjang default terisi 2 dan kapasitas 5
	newSlice[0] = "Mustofa"
	newSlice[1] = "Adny"
	fmt.Println(newSlice) // [Mustofa Adny]

	newSlice2 := append(newSlice, "Ucok", "Udin") // menambahkan elemen baru ke dalam slice (memakai kapasitas slice newSlice karena kapasitasnya 5)
	fmt.Println(newSlice2) // [Mustofa Adny Ucok Udin]

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice) // menyalin isi dari fromSlice ke toSlice
	fmt.Println(toSlice) // [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]
}