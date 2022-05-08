package lc203

/*
分析：遇见相等的，直接删除，注意head可能被删除，设置一个dummy
时间复杂度：O(n), 只遍历一次链表
空间复杂度：O(1)
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{-1, head}
	curr := dummy

	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}