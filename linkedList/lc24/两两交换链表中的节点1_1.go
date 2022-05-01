package lc24

/*
分析：官方的递归
时间复杂度：
空间复杂度：
 */

func swapPairs1_1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	// 递归
	head.Next = swapPairs(newHead.Next)
	// 反转
	newHead.Next = head
	return newHead
}

