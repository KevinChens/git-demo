package lc82

/*
分析：链表头结点可能被删除，所以用 dummy node辅助删除
时间复杂度：O(n)
空间复杂度：O(1)
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	curr := dummy

	var rmVal int
	// 两个节点
	for curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			// 记录已经删除的值，用于后序节点判断
			rmVal = curr.Next.Val
			for curr.Next != nil && curr.Next.Val == rmVal {
				curr.Next = curr.Next.Next
			}
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}


