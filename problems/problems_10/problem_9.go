package problems10

/* Problem Info
Asked by: Airbnb

Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.

For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5] should return 10, since we pick 5 and 5.

*/

import "math"

func naSum(array []int) int {
	incl := 0
	excl := 0
	for _, i := range array {
		new := 0
		if excl > incl {
			new = excl
		} else {
			new = incl
		}
		incl = excl + i
		excl = new
	}
	return int(math.Max(float64(excl), float64(incl)))
}
