package links

import (
	"fmt"
	"testing"
)

func TestLinkNode(T *testing.T) {
	node1 := &LinkNode{}
	node2 := &LinkNode{}
	node3 := &LinkNode{}
	node4 := &LinkNode{}
	// node5 := &LinkNode{}
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node2

	fmt.Println(hasCycle2(node1))
}
