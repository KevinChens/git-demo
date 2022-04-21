package lc21

/*
分析：通过 dummy node链表，连接各个元素, 迭代
时间复杂度：O(n+m), 遍历两个链表
空间复杂度：O(1)
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	head := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}

	if list1 == nil {
		// 连接list2未处理完的节点
		head.Next = list2
	} else {
		// 连接list1未处理完的节点
		head.Next = list1
	}
	return dummy.Next
}