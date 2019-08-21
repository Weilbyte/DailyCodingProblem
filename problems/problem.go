package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func problemInfo(askedBy string, problem string) {
	fmt.Printf("\nAsked by: %v\nProblem: %s\n\n", askedBy, problem)
}

func InputNumbersBySpace(inputText string) []int {
	result := []int{}
	fmt.Print(inputText + ": ")
	stringInput := bufio.NewScanner(os.Stdin)
	stringInput.Scan()
	stringSpace := strings.Fields(stringInput.Text())
	for _, n := range stringSpace {
		nInteger, err := strconv.Atoi(n)
		result = append(result, nInteger)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}
	return result
}
