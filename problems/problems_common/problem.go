//Package problems_common contains common functions shared between all the problems
package problems_common

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func problemInfo(askedBy string, problem string) {
	fmt.Printf("\nAsked by: %v\nProblem: %s\n\n", askedBy, problem)
}

// InputNumbersBySpace input, which should be numbers, and splits them into an array by whitespace.
// Returns the result array.
func InputNumbersBySpace(inputText string) (array []int, err error) {
	result := []int{}
	fmt.Print(inputText + ": ")
	stringInput := bufio.NewScanner(os.Stdin)
	stringInput.Scan()
	stringSpace := strings.Fields(stringInput.Text())
	for _, n := range stringSpace {
		for _, r := range n {
			if unicode.IsLetter(r) {
				return []int{}, errors.New("Input accepts numbers only")
			}
		}
		nInteger, err := strconv.Atoi(n)
		result = append(result, nInteger)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}
	return result, nil
}
