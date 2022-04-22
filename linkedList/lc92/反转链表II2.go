package lc92

/*
分析：一次遍历+头插法
时间复杂度：O(n)
空间复杂度：O(1)
 */

func reverseBetween2(head *ListNode, left, right int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{-1, head}
	prev := dummy
	// 找到反转前的位置
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	curr := prev.Next

	for i := 0; i < right-left; i++ {
		next := curr.Next
		// 头插法反转
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dummy.Next
}