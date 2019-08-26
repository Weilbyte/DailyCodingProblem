package problems10

/*

Unit tests for problems 0-10

*/

import (
	"reflect"
	"testing"
)

func Test_Problem1(t *testing.T) {
	test := adduptok([]int{10, 15, 3, 7}, 17)
	if test != true {
		t.Errorf("problem_1/test01: Expected true, got %v", test)
	}
}

func Test_Problem2(t *testing.T) {
	expected := []int{120, 60, 40, 30, 24}
	test := arrallnumproduct([]int{1, 2, 3, 4, 5})
	if !reflect.DeepEqual(test, expected) {
		t.Errorf("problem_2/test01: Expected %v, got %v", expected, test)
	}
}

func Test_Problem3(t *testing.T) {
	arr := []int{5, 2, -4, 3, 21, 19, 25}
	tree := node{Right: nil, Left: nil, Value: arr[0]}
	for xi, x := range arr {
		if xi == 0 {
			continue
		}
		addNode(&tree, x)
	}
	test := deserializetree(serializetree(&tree)).Left.Left.Value
	if tree.Left.Left.Value != test {
		t.Errorf("problem_3/test01: Expected -4, got '%v'", test)
	}
}

func Test_Problem4(t *testing.T) {
	test := findSmallestMissing([]int{3, 4, -1, 1})
	if test != 2 {
		t.Errorf("problem_4/test01: Expected 2, got '%v'", test)
	}
	test = findSmallestMissing([]int{1, 2, 0})
	if test != 3 {
		t.Errorf("problem_4/test02: Expected 3, got '%v'", test)
	}
}

func Test_Problem5(t *testing.T) {
	expected := pair{first: 3, last: 4}
	testpair := cons(3, 4)
	if (testpair.first != expected.first) || (testpair.last != expected.last) {
		t.Errorf("problem_4/test01: Expected %v, got '%v'", expected, testpair)
	}
	test := car(cons(3, 4))
	if test != 3 {
		t.Errorf("problem_5/test02: Expected 3, got '%v'", test)
	}
	test = cdr(cons(3, 4))
	if test != 4 {
		t.Errorf("problem_5/test03: Expected 4, got '%v'", test)
	}
}

func Test_Problem6(t *testing.T) {
	testnode := &xorNode{value: 00, both: uintptr(0)}
	addxor(testnode, 10)
	addxor(testnode, 20)
	test := getxor(testnode, 0).value
	if test != 00 {
		t.Errorf("problem_6/test01: Expected 00, got '%v'", test)
	}
	test = getxor(testnode, 1).value
	if test != 10 {
		t.Errorf("problem_6/test02: Expected 10, got '%v'", test)
	}
	test = getxor(testnode, 2).value
	if test != 20 {
		t.Errorf("problem_6/test03: Expected 20, got '%v'", test)
	}
}

func Test_Problem7(t *testing.T) {
	test := decode(encode("aaa"))
	if test != 3 {
		t.Errorf("problem_7/test01: Expected 3, got '%v'", test)
	}
	test = decode(encode("fff"))
	if test != 1 {
		t.Errorf("problem_7/test02: Expected 1, got '%v'", test)
	}
	test = decode(encode("zf"))
	if test != 2 {
		t.Errorf("problem_7/test03: Expected 2, got '%v'", test)
	}
}

func Test_Problem8(t *testing.T) {
	testtree := &node{Value: 0, Left: &node{Value: 1}, Right: &node{Value: 0, Right: &node{Value: 0}, Left: &node{Value: 1, Right: &node{Value: 1}, Left: &node{Value: 1}}}}
	test := countUnivalHelper(testtree)
	if test != 5 {
		t.Errorf("problem_8/test01: Expected 5, got '%v'", test)
	}
}

func Test_Problem9(t *testing.T) {
	test := naSum([]int{2, 4, 6, 2, 5})
	if test != 13 {
		t.Errorf("problem_9/test01: Expected 13, got '%v'", test)
	}
	test = naSum([]int{5, 1, 1, 5})
	if test != 10 {
		t.Errorf("problem_9/test02: Expected 10, got '%v'", test)
	}
}

func Test_Problem10(t *testing.T) {
	test := runAfter(runAfterExample, 200)
	if test != true {
		t.Errorf("problem_10/test01: Expected true, got '%v'", test)
	}
}
