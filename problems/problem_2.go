package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Problem2() {
	problemInfo("Uber", "Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.")
	problem2solution()
}

func problem2solution() {
	numArrayOriginal := []int{}
	numArrayResult := []int{}
	fmt.Print("Provide a list of numbers to add to array, seperated by a space: ")
	numListInput := bufio.NewScanner(os.Stdin)
	numListInput.Scan()
	numListInputArray := strings.Fields(numListInput.Text())
	for _, x := range numListInputArray {
		res, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		numArrayOriginal = append(numArrayOriginal, res)
	}
	for xi, _ := range numArrayOriginal {
		sum := 0
		for yi, y := range numArrayOriginal {
			if yi == xi {
				continue
			}
			if sum == 0 {
				sum = y
				continue
			}
			sum = sum * y
		}
		numArrayResult = append(numArrayResult, sum)
	}
	fmt.Println(numArrayResult)
	fmt.Println("Done")
}
