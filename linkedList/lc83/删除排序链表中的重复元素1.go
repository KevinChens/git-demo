package lc83

/*
分析：先遍历，全部删除后，再移动下一个元素
时间复杂度：O(n), 遍历一次链表
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
	curr := head
	for curr != nil {
		// 全部删除完再移动到下一个元素
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		}
		curr = curr.Next
	}
	return head
}
