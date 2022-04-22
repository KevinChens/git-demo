package lc142

/*
分析：快慢指针，快慢相遇之后，慢指针回到头，快慢指针步调一致一起移动，相遇点即为入环点
时间复杂度：O(n)
空间复杂度：O(1)
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if fast == slow {
			// slow从头开始移动，fast从第一次相交点开始一个一个移动
			slow = head
			fast = fast.Next
			// 再次相遇，返回slow
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}
