package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Problem1() {
	problemInfo("Google", "Given a list of numbers and a number k, return whether any two numbers from the list add up to k.")
	problem1solution()
}

func problem1solution() {
	numList := []int{}
	fmt.Print("Provide list of numbers seperated by a space: ")
	numListStringInput := bufio.NewScanner(os.Stdin)
	numListStringInput.Scan()
	numListString := strings.Fields(numListStringInput.Text())
	for _, n := range numListString {
		nInteger, err := strconv.Atoi(n)
		numList = append(numList, nInteger)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}
	fmt.Print("k=")
	kStr := bufio.NewScanner(os.Stdin)
	kStr.Scan()
	k, err := strconv.Atoi(kStr.Text())
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	for _, x := range numList {
		for _, y := range numList {
			xysum := x + y
			if xysum == k {
				fmt.Printf("The sum of %v and %s make k\n", strconv.Itoa(x), strconv.Itoa(y))
			}
		}
	}
	fmt.Println("\n Done.")
}
