package links

func reverseLink(head *LinkNode) *LinkNode {
	if head == nil || head.next == nil {
		return head
	}

	var (
		prev *LinkNode
		curr = head
	)
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
}
