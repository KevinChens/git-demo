package lc237

/*
分析：因为题目传入的是要被删除的节点，同时不是末尾节点，
所以可以copy下一个节点的值，再删除下一个节点，整体效果就是删除了本节点
时间复杂度：O(1)
空间复杂度：O(1)
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
