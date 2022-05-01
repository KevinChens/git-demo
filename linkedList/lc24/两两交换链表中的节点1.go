package lc24

/*
分析：将链表反转化为一个子问题，递归解决
先反转两个，后面的节点也以同样的方式反转，最后将反转的节点连接起来
时间复杂度：O(n)
空间复杂度：O(n)
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	return helper(head)
}

func helper(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 保存下一阶段的头指针
	nextHead := head.Next.Next

	next := head.Next
	// 反转当前指针
	next.Next = head
	// 递归
	head.Next = helper(nextHead)
	return next
}