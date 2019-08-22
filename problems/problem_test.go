package problems

import (
	"reflect"
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

func TestProblem2Solution(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{120, 60, 40, 30, 24}

	test := problem2solution(arr)
	if !reflect.DeepEqual(test, expected) {
		t.Errorf("problem2solution: Expected %v, got %v", expected, test)
	}
}

func TestProblem3Solution(t *testing.T) {
	test1 := problem3solution([]int{5, 2, -4, 3, 21, 19, 25})
	test2 := problem3solution([]int{10, 6, 18, 4, 8, 15, 21, 22})

	if test1 != "Passed" {
		t.Errorf("problem3solution/test1: Expected 'Passed', got '%v'", test1)
	}
	if test2 != "Passed" {
		t.Errorf("problem3solution/test2: Expected 'Passed', got '%v'", test2)
	}
}

func TestProblem4Solution(t *testing.T) {
	test1 := problem4solution([]int{3, 4, -1, 1})
	test2 := problem4solution([]int{1, 1, 2, 0})

	if test1 != 2 {
		t.Errorf("problem4solution/test1: Expected 2, got '%v'", test1)
	}
	if test2 != 3 {
		t.Errorf("problem4solution/test2: Expected 3, got '%v'", test2)
	}
}

func TestProblem5Solution(t *testing.T) {
	conspair := Pair{first: 3, last: 4}
	constest := cons(3, 4)
	cartest := car(cons(3, 4))
	cdrtest := cdr(cons(3, 4))
	if (conspair.first != constest.first) || (conspair.last != constest.last) {
		t.Errorf("problem5/constest: Expected %v, got '%v'", conspair, constest)
	}
	if cartest != 3 {
		t.Errorf("problem5/cartest: Expected 3, got '%v'", cartest)
	}
	if cdrtest != 4 {
		t.Errorf("problem5/cdrtest: Expected 4, got '%v'", cdrtest)
	}

}
