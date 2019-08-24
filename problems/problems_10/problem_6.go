package problems_10

/* Problem Info
Asked by: Google

An XOR linked list is a more memory efficient doubly linked list.
Instead of each node holding next and prev fields, it holds a field named both, which is an XOR of the next node and the previous node.
Implement an XOR linked list; it has an add(element) which adds the element to the end, and a get(index) which returns the node at index.

*/

import (
	"unsafe"
)

//xorNode struct represents a node in the XOR linked list
type xorNode struct {
	value int
	both  uintptr
}

func addxor(xnode *xorNode, value int) {
	var previous uintptr
	current := unsafe.Pointer(xnode)
	xor := previous ^ xnode.both
	for xor != 0 {
		previous = uintptr(current)
		current = unsafe.Pointer(xor)
		xnode = (*xorNode)(current)
		xor = previous ^ xnode.both
	}
	xnode.both = previous ^ uintptr(unsafe.Pointer(&xorNode{value: value, both: uintptr(current)}))
}

func getxor(xnode *xorNode, index int) *xorNode {
	var previous uintptr
	current := unsafe.Pointer(xnode)
	xor := previous ^ xnode.both
	for i := 0; i < index; i++ {
		previous = uintptr(current)
		current = unsafe.Pointer(xor)
		xnode = (*xorNode)(current)
		xor = previous ^ xnode.both
	}
	return xnode
}
