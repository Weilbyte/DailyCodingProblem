package problems

import (
	"fmt"
	"os"
)

//Problem2 prints a brief info about the problem and runs the solution
func Problem2() {
	problemInfo("Uber", "Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.")
	problem2solution()
}

func problem2solution() {
	numArrayOriginal, err := InputNumbersBySpace("Provide a list of numbers to add to array, seperated by a space")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	numArrayResult := []int{}
	for xi := range numArrayOriginal {
		sum := 1
		for yi, y := range numArrayOriginal {
			if yi == xi {
				continue
			}
			sum *= y
		}
		numArrayResult = append(numArrayResult, sum)
	}
	fmt.Println(numArrayResult)
	fmt.Println("Done")
}
