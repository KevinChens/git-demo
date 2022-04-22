package lc206

/*
分析：迭代反转链表
时间复杂度O(n)，需要遍历一次链表
空间复杂度O(1)，不需要额外的辅助空间
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	// prev-->curr-->next
	// 要保存prev,next
	// 将curr的next指向prev; prev，curr同步后移，迭代
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		// 开始反转
		curr.Next = prev
		// prev,curr同步后移
		prev = curr
		curr = next
	}
	// 返回新头节点prev
	return prev
}
