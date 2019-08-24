package problems10

/* Problem Info
Asked by: Uber

Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

*/

func arrallnumproduct(numArrayOriginal []int) []int {
	numArrayResult := []int{}
	for xi := range numArrayOriginal {
		sum := 1
		for yi, y := range numArrayOriginal {
			if yi == xi {
				continue
			}
			sum *= y
		}
		numArrayResult = append(numArrayResult, sum)
	}
	return numArrayResult
}
