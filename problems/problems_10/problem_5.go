package problems_10

/* Problem Info
Asked by: Jane Street

cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and last element of that pair. For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4.

*/

//pair struct represents the pair
type pair struct {
	first int
	last  int
}

func cons(a int, b int) *pair {
	return &pair{first: a, last: b}
}

func car(pair *pair) int {
	return pair.first
}

func cdr(pair *pair) int {
	return pair.last
}
