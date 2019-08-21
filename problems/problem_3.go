package problems

import (
	"fmt"
	"strconv"
	"strings"
)

//Daily Coding Problem #3
func Problem3() {
	problemInfo("Google", "Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.")
	problem3solution()
}

//Represents a binary tree node
type Node struct {
	Left  *Node
	Right *Node
	Value int
}

//Holds serialized string
type Serial struct {
	Value string
}

func problem3solution() {
	fmt.Println("Attempting with BST of 5 2 -4 3 21 19 25")
	input := [7]int{5, 2, -4, 3, 21, 19, 25}
	Tree := Node{Right: nil, Left: nil, Value: input[0]}
	for xi, x := range input {
		if xi == 0 {
			continue
		}
		addNode(&Tree, x)
	}
	if Tree.Left.Left.Value == deserialize(serialize(&Tree)).Left.Left.Value {
		fmt.Println("Passed.")
	}
}

func addNode(tree *Node, val int) {
	insertee := &Node{Value: val}
	if insertee.Value > tree.Value {
		if tree.Right == nil {
			tree.Right = &Node{Value: insertee.Value}
		}
		addNode(tree.Right, val)
	}
	if insertee.Value < tree.Value {
		if tree.Left == nil {
			tree.Left = &Node{Value: insertee.Value}
		}
		addNode(tree.Left, val)
	}
}

func traverse(root *Node, s *Serial) {
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

func serialize(root *Node) string {
	serialized := Serial{Value: ""}
	traverse(root, &serialized)
	return serialized.Value
}

func deserialize(serialized string) *Node {
	sArray := strings.Split(serialized, "#")
	i, e := strconv.Atoi(sArray[0])
	if e != nil {
		fmt.Println("Error has occureddd!", e)
	}
	Result := Node{Right: nil, Left: nil, Value: i}
	for _, x := range sArray {
		ii, e := strconv.Atoi(x)
		if e != nil {
			fmt.Println("Error has occured!", e)
		}
		addNode(&Result, ii)
	}
	return &Result
}
