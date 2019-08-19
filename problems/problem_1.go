package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Problem1() {
	problemInfo("Google", "Given a list of numbers and a number k, return whether any two numbers from the list add up to k.")
	problem1solution()
}

func problem1solution() {
	numList := InputNumbersBySpace("Provide list of numbers seperated by a space")
	match := false
	fmt.Print("k=")
	kStr := bufio.NewScanner(os.Stdin)
	kStr.Scan()
	k, err := strconv.Atoi(kStr.Text())
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	for _, x := range numList {
		for _, y := range numList {
			if x+y == k {
				match = true
			}
		}
	}
	fmt.Println(match)
	fmt.Println("\nDone.")
}
