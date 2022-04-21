package lc141

/*
分析：快慢指针，快慢指针相同则有环，证明：如果有环每走一步快慢指针距离会减 1
时间复杂度：O(n)
空间复杂度：O(1)
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		// 比较指针是否相等，不要使用val比较
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
