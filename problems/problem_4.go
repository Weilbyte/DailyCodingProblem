package problems

import (
	"fmt"
	"os"
	"sort"
)

//Problem4 prints a brief info about the problem and runs the solution
func Problem4() {
	problemInfo("Stripe", "Given an array of integers, find the first missing positive integer in linear time and constant space. In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.")
	arrayraw, err := InputNumbersBySpace("Provide list of numbers for array, seperated by space")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Printf("Smallest positive integer is %v", problem4solution(arrayraw))
}

func problem4solution(arrayraw []int) int {
	array := []int{}
	for _, x := range arrayraw {
		if containsNum(array, x) || (x < 0) {
			continue
		}
		array = append(array, x)
	}
	sort.Ints(array)
	return findSmallestMissing(array)
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
