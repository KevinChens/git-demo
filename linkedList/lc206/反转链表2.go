package lc206

/*
分析：递归反转
时间复杂度O(n)，需要反转n个节点；
空间复杂度O(n)，需要递归调用n层
*/

func reverseList2(head *ListNode) *ListNode {
	// 0和1个节点不用反转
	if head == nil || head.Next == nil {
		return head
	}
	// 2个及以上递归反转
	newHead := reverseList2(head.Next)
	// A->B ;  B.Next = A, A.Next = nil; B->A
	head.Next.Next = head
	head.Next = nil
	// 返回newHead
	return newHead
}
