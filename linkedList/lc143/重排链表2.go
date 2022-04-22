package lc143

/*
分析：利用线性表
时间复杂度：O(n)
空间复杂度：O(n)
 */

func reorderList2(head *ListNode) {
	if head == nil {
		return
	}

	var nodes []*ListNode
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
