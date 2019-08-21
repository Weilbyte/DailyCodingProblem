package problems

import (
	"fmt"
	"sort"
)

//Daily Coding Problem #4
func Problem4() {
	problemInfo("Stripe", "Given an array of integers, find the first missing positive integer in linear time and constant space. In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.")
	problem4solution()
}

func problem4solution() {
	arrayraw := InputNumbersBySpace("Provide list of numbers for array, seperated by space")
	array := []int{}
	for _, x := range arrayraw {
		if containsNum(array, x) || (x < 0) {
			continue
		}
		array = append(array, x)
	}
	sort.Ints(array)
	fmt.Println("\nSmallest positive integer is ", findSmallestMissing(array))
}

func containsNum(array []int, number int) bool {
	for _, x := range array {
		if x == number {
			return true
		}
	}
	return false
}

func findSmallestMissing(array []int) int {
	num := 1
	for _, x := range array {
		if x == 0 {
			continue
		}
		if x == num {
			num++
			continue
		}
		break
	}
	return num
}
