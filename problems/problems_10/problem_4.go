package problems_10

/* Problem Info
Asked by: Stripe

Given an array of integers, find the first missing positive integer in linear time and constant space.
In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.

*/

import (
	"sort"
)

func containsNum(array []int, number int) bool {
	for _, x := range array {
		if x == number {
			return true
		}
	}
	return false
}

func findSmallestMissing(arrayinput []int) int {
	array := []int{}
	for _, x := range arrayinput {
		if containsNum(array, x) || (x < 0) {
			continue
		}
		array = append(array, x)
	}
	sort.Ints(array)
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
