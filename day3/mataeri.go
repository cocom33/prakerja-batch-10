package day3

import (
	"fmt"

	"slices"
)

func materi3() {
	// array
	// var score [5]int = [5]int{1,2}
	// score[0] = 10
	// fmt.Println(score)

	// slice
	baju := []int{1,2,3,4}
	// fmt.Println(baju)
	baju = slices.Delete(baju, 1, 2)
	fmt.Println(baju)
	for _, v := range baju {
		fmt.Print(v, " ")
	}

	// map
	// var pendapatan map[string]int
	pendapatan := map[string]int{}
	pendapatan["jan"] = 900 
	
	result, isFound := pendapatan["maret"]
	
	fmt.Println("\n", result, isFound)
}