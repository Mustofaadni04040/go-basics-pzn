package main

import "fmt"

func ForLoops() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Counter =", counter)
	}

	fmt.Println("Finish")

	// for range loop
	names := []string{"Mustofa", "Adny", "Siti", "Budi"}

	for i := 0; i < len(names); i++ {
		fmt.Println("Name at index", i, "is", names[i])
	}

	for index, name := range names { // if index is not needed, can removed by using _
		fmt.Println("Name at index", index, "is", name)
	}
}