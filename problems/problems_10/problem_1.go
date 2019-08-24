//Package problems10 contains the problems from 0-10
package problems10

/* Problem Info
Asked by: Google

Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

*/

func adduptok(numList []int, K int) bool {
	match := false
	for _, x := range numList {
		for _, y := range numList {
			if x+y == K {
				match = true
			}
		}
	}
	return match
}
