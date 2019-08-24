package problems_10

/* Problem Info
Asked by: Google

Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

*/

import (
	"fmt"
	"strconv"
	"strings"
)

//node epresents a binary tree node
type node struct {
	Left  *node
	Right *node
	Value int
}

//serial holds the serialized string
type serial struct {
	Value string
}

//move to unit test
func problem3solution(input []int) string {

	Tree := node{Right: nil, Left: nil, Value: input[0]}
	for xi, x := range input {
		if xi == 0 {
			continue
		}
		addNode(&Tree, x)
	}
	if Tree.Left.Left.Value == deserialize(serialize(&Tree)).Left.Left.Value {
		return "Passed"
	}
	return "Failed"
}

func addNode(tree *node, val int) {
	insertee := &node{Value: val}
	if insertee.Value > tree.Value {
		if tree.Right == nil {
			tree.Right = &node{Value: insertee.Value}
		}
		addNode(tree.Right, val)
	}
	if insertee.Value < tree.Value {
		if tree.Left == nil {
			tree.Left = &node{Value: insertee.Value}
		}
		addNode(tree.Left, val)
	}
}

func traverse(root *node, s *serial) {
	if root != nil {
		pfx := "#"
		if len(s.Value) == 0 {
			pfx = ""
		}
		s.Value += fmt.Sprintf("%v%s", pfx, strconv.Itoa(root.Value))
		traverse(root.Left, s)
		traverse(root.Right, s)
	}
}

func serialize(root *node) string {
	serialized := serial{Value: ""}
	traverse(root, &serialized)
	return serialized.Value
}

func deserialize(serialized string) *node {
	sArray := strings.Split(serialized, "#")
	i, e := strconv.Atoi(sArray[0])
	if e != nil {
		fmt.Println("Error has occureddd!", e)
	}
	Result := node{Right: nil, Left: nil, Value: i}
	for _, x := range sArray {
		ii, e := strconv.Atoi(x)
		if e != nil {
			fmt.Println("Error has occured!", e)
		}
		addNode(&Result, ii)
	}
	return &Result
}
