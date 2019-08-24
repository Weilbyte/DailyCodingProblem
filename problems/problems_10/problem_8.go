package problems10

/* Problem Info
Asked by: Google

A unival tree (which stands for "universal value") is a tree where all nodes under it have the same value.
Given the root to a binary tree, count the number of unival subtrees.

Notes:
node struct borrowed from problem_3.go
*/

func countUnivalHelper(node *node) int {
	result := 0
	countUnival(node, &result)
	return result
}

func countUnival(node *node, counter *int) {
	if node != nil {
		if (node.Left == nil) && (node.Right == nil) {
			*counter++
		} else {
			if node.Left.Value == node.Right.Value {
				*counter++
			}
		}
		countUnival(node.Left, counter)
		countUnival(node.Right, counter)
	}
}
