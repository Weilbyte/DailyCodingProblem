package main

import (
	"bufio"
	"fmt"
	"github.com/weilbyte/DailyCodingProblem/problems"
	"os"
	"strconv"
)

func main() {
	curSolvedList := []int{1, 2, 3, 4, 5, 6} //Add new number for each solved problem
	var curSolvedString string
	var curSolvedCount int8
	for _, number := range curSolvedList {
		pfx := ", "
		if len(curSolvedString) == 0 {
			pfx = ""
		}
		curSolvedString += fmt.Sprintf("%v#%s", pfx, strconv.Itoa(number))
		curSolvedCount++
	}
	fmt.Printf("I have currently solved %v daily problems: %s\n", curSolvedCount, curSolvedString)
	fmt.Print("Which one do you want to run?\n#")
	selProblem := bufio.NewScanner(os.Stdin)
	selProblem.Scan()
	switch selProblem.Text() {
	case "1":
		problems.Problem1()
	case "2":
		problems.Problem2()
	case "3":
		problems.Problem3()
	case "4":
		problems.Problem4()
	case "5":
		problems.Problem5()
	case "6":
		problems.Problem6()
	default:
		fmt.Println("Uh oh, that didnt work :c")
	}
}
