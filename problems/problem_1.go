package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Problem1 prints a brief info about the problem, gathers user input and runs the solution
func Problem1() {
	problemInfo("Google", "Given a list of numbers and a number k, return whether any two numbers from the list add up to k.")
	numList, err := InputNumbersBySpace("Provide list of numbers seperated by a space")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Print("k=")
	KStr := bufio.NewScanner(os.Stdin)
	KStr.Scan()
	K, err := strconv.Atoi(KStr.Text())
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	problem1solution(numList, K)
}

func problem1solution(numList []int, K int) bool {
	match := false
	for _, x := range numList {
		for _, y := range numList {
			if x+y == K {
				match = true
			}
		}
	}
	return match
}
