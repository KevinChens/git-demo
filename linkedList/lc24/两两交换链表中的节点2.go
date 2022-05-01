package lc24

/*
分析：迭代
时间复杂度：O(n)
空间复杂度：O(1)
 */

func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{-1, head}
	temp := dummy

	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		// 迭代
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		temp = node1
	}
	return dummy.Next
}
