package lc203

/*
分析：递归, 对于给定的链表，首先对除了头节点head以外的节点进行删除操作，
然后判断head的节点值是否等于给定的val, 等于head就被删除
时间复杂度：O(n)
空间复杂度：O(n)
 */

func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}