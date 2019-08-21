package problems

import (
	"fmt"
)

//Problem5 prints a brief info about the problem and runs the solution
func Problem5() {
	problemInfo("Jane Street", "cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and last element of that pair. For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4.")
	problem5solution()
}

func problem5solution() {
	fmt.Println("car(cons(3, 4)) = ", car(cons(3, 4)))
	fmt.Println("cdr(cons(3, 4)) = ", cdr(cons(3, 4)))
}

//Pair struct represents the pair
type Pair struct {
	first int
	last  int
}

func cons(a int, b int) *Pair {
	return &Pair{first: a, last: b}
}

func car(pair *Pair) int {
	return pair.first
}

func cdr(pair *Pair) int {
	return pair.last
}
