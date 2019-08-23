package problems

import (
	"fmt"
	"unsafe"
)

//Problem6 prints a brief info about the problem and runs the solution
func Problem6() {
	problemInfo("Google", "An XOR linked list is a more memory efficient doubly linked list. Instead of each node holding next and prev fields, it holds a field named both, which is an XOR of the next node and the previous node. Implement an XOR linked list; it has an add(element) which adds the element to the end, and a get(index) which returns the node at index.")
	problem6solution()
}

func problem6solution() {
	xnode := &XORNode{value: 0, both: uintptr(0)}
	addxor(xnode, 10)
	addxor(xnode, 20)
	addxor(xnode, 30)
	fmt.Println("Added XOR nodes with values 10, 20 and 30.")
	fmt.Printf("Values of XOR nodes at indexes 0, 1, 2 and 3 are: %v %v %v and %v", getxor(xnode, 0).value, getxor(xnode, 1).value, getxor(xnode, 2).value, getxor(xnode, 3).value)
}

//XORNode struct represents a node in the XOR linked list
type XORNode struct {
	value int
	both  uintptr
}

func addxor(xnode *XORNode, value int) {
	var previous uintptr
	current := unsafe.Pointer(xnode)
	xor := previous ^ xnode.both
	for xor != 0 {
		previous = uintptr(current)
		current = unsafe.Pointer(xor)
		xnode = (*XORNode)(current)
		xor = previous ^ xnode.both
	}
	xnode.both = previous ^ uintptr(unsafe.Pointer(&XORNode{value: value, both: uintptr(current)}))
}

func getxor(xnode *XORNode, index int) *XORNode {
	var previous uintptr
	current := unsafe.Pointer(xnode)
	xor := previous ^ xnode.both
	for i := 0; i < index; i++ {
		previous = uintptr(current)
		current = unsafe.Pointer(xor)
		xnode = (*XORNode)(current)
		xor = previous ^ xnode.both
	}
	return xnode
}
