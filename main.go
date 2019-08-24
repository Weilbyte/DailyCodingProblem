package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	curSolvedList := []int{1, 2, 3, 4, 5, 6, 7} //Add new number for each solved problem
	var curSolvedString string
	for _, number := range curSolvedList {
		if len(curSolvedString) == 0 {
			curSolvedString += fmt.Sprintf("%v#%s", ", ", strconv.Itoa(number))
		}
		curSolvedString += fmt.Sprintf("%v#%s", "", strconv.Itoa(number))
	}
	fmt.Printf("I have currently solved %v daily problems: %s\n", len(curSolvedList), curSolvedString)
	bufio.NewScanner(os.Stdin).Scan()
}
