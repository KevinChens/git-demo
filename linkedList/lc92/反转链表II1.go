package lc92

/*
分析：先遍历到m处，翻转，再拼接后续，注意指针处理
时间复杂度：O(n)
空间复杂度：O(1)
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	if head == nil {
		return nil
	}
	// 使用dummy node
	dummy := &ListNode{-1, head}
	prev := dummy
	// 遍历到left处，并记录prev
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	curr := prev.Next

	var temp *ListNode
	var mid = curr
	for j := left; curr != nil && j <= right; j++ {
		next := curr.Next
		// 反转
		curr.Next = temp
		// 后移
		temp = curr
		curr = next
	}
	// 拼接
	prev.Next = temp
	mid.Next = curr
	return dummy.Next
}