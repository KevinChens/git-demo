package lc86

/*
分析：将大于 x 的节点，放到另外一个链表，最后连接这两个链表
时间复杂度：O(n)
空间复杂度：O(1)
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	// check
	if head == nil {
		return nil
	}
	headDummy := &ListNode{-1, head}
	tailDummy := &ListNode{-1, nil}
	tail := tailDummy
	head = headDummy

	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			// 移除大于x的节点
			node := head.Next
			head.Next = head.Next.Next
			// 放到另外一个链表
			tail.Next = node
			tail = tail.Next
		}
	}
	// 拼接两个链表
	tail.Next = nil
	head.Next = tailDummy.Next
	return headDummy.Next
}
