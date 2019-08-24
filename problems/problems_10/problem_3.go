package problems10

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

func traversetree(root *node, s *serial) {
	if root != nil {
		pfx := "#"
		if len(s.Value) == 0 {
			pfx = ""
		}
		s.Value += fmt.Sprintf("%v%s", pfx, strconv.Itoa(root.Value))
		traversetree(root.Left, s)
		traversetree(root.Right, s)
	}
}

func serializetree(root *node) string {
	serialized := serial{Value: ""}
	traversetree(root, &serialized)
	return serialized.Value
}

func deserializetree(serialized string) *node {
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
