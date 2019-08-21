package problems

import (
	"testing"
)

func TestProblem1Solution(t *testing.T) {
	arr := []int{10, 15, 3, 7}
	ik := 17
	test := problem1solution(arr, ik)
	if test != true {
		t.Errorf("problem1solution: Expected true, got %v", test)
	}

}
