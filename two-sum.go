package main

import "fmt"

func TwoSums(nums []int, target int) []int {
	// var numberMap map[int]int = make(map[int]int)

	// for i := 0; i < len(nums); i++ {
	// 	numberNeeded := target - nums[i]
	// 	fmt.Println("numberNeeded", numberNeeded)

	// 	if indexNeeded, found := numberMap[numberNeeded]; found {
	// 		return []int{indexNeeded, i}
	// 	}

	// 	numberMap[nums[i]] = i
	// 	fmt.Println(numberMap)
	// }
	// return nil


	var numberMap map[int]int = make(map[int]int)

	for index, num := range nums {
		numberNeeded := target - num // misal: target = 9, num = 2, numberNeeded = 7
		fmt.Println("numberNeeded", numberNeeded)

		indexNeeded, found := numberMap[numberNeeded]; // check apakan numberNeeded ada di map, jika ada maka found = true
		fmt.Println("indexNeeded", indexNeeded)
		if found {
			return []int{indexNeeded, index} // index dari 2 yang di numberMap, index dari 7 yang di nums
		}

		numberMap[num] = index // iterasi pertama index = 0, num = 2, map[2] = 0
	}
	return nil
}