package lc876

/*
分析：遍历一次链表，将节点存储在切片，返回中间节点
时间复杂度；O(n)
空间复杂度：O(n)
 */

func middleNode2(head *ListNode) *ListNode {
	curr := head
	nodes := make([]*ListNode, 0)
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}
	return nodes[len(nodes)/2]
}
