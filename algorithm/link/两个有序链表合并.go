package links

// 合并两个排序的链表: 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
// 递归
func mergeTwoLists1(l1 *LinkNode, l2 *LinkNode) *LinkNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var curr *LinkNode
	if l1.data < l2.data {
		curr = l1
		curr.next = mergeTwoLists1(l1.next, l2)
	} else {
		curr = l2
		curr.next = mergeTwoLists1(l1, l2.next)
	}

	return curr
}
