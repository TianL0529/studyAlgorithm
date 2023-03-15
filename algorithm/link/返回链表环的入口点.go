package links

// 思考： 若需要返回入环的第一个节点，该怎么做。
// 思路： 从头节点开始，快指针一次走两步，慢指针一次走一步，若链表有环，则快慢指针一定会在环上的某个节点相遇，之后，快指针回到开头变成一次走一步，慢指针在原地开始继续一次走一步，再次相遇的节点一定在入环的第一个节点处。
func solution(head *LinkNode) *LinkNode {
	if head == nil || head.next == nil || head.next.next == nil {
		return nil
	}

	lowPtr := head.next
	fastPtr := head.next.next
	for lowPtr != fastPtr {
		if fastPtr.next == nil || fastPtr.next.next == nil {
			return nil
		}

		lowPtr = lowPtr.next
		fastPtr = fastPtr.next.next
	}

	fastPtr = head
	for lowPtr != fastPtr {
		lowPtr = lowPtr.next
		fastPtr = fastPtr.next
	}

	return lowPtr
}
