package links

type LinkNode struct {
	data int
	next *LinkNode
}

// hasCycle1 哈希表发
func hasCycle1(head *LinkNode) bool {
	if head == nil || head.next == nil {
		return false
	}

	nodeMap := make(map[LinkNode]struct{})
	curr := head
	for curr.next != nil {
		curr = curr.next
		if _, ok := nodeMap[*curr]; ok {
			return true
		}

		nodeMap[*curr] = struct{}{}
	}

	return false
}

// hasCycle2 快慢指针
// 思路： 从头节点开始，快指针一次走两步，慢指针一次走一步，若链表有环，则快慢指针一定会相遇。若快慢指针不相遇，说明没环。
func hasCycle2(head *LinkNode) bool {
	if head == nil || head.next == nil || head.next.next == nil {
		return false
	}

	lowPtr := head.next
	fastPtr := head.next.next
	for lowPtr != fastPtr {
		if fastPtr.next == nil || fastPtr.next.next == nil {
			return false
		}

		lowPtr = lowPtr.next
		fastPtr = fastPtr.next.next
	}

	return true
}
