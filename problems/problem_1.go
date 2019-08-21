package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Problem1 prints a brief info about the problem and runs the solution
func Problem1() {
	problemInfo("Google", "Given a list of numbers and a number k, return whether any two numbers from the list add up to k.")
	problem1solution()
}

func problem1solution() {
	numList := InputNumbersBySpace("Provide list of numbers seperated by a space")
	match := false
	fmt.Print("k=")
	KStr := bufio.NewScanner(os.Stdin)
	KStr.Scan()
	K, err := strconv.Atoi(KStr.Text())
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	for _, x := range numList {
		for _, y := range numList {
			if x+y == K {
				match = true
			}
		}
	}
	fmt.Println(match)
	fmt.Println("\nDone.")
}
