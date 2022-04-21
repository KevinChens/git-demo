package lc141

/*
分析：使用哈希表来存储所有已经访问过的节点，遍历所有节点，每次遍历到一个节点时，判断该节点此前是否被访问过
时间复杂度：O(n)
空间复杂度：O(n)
 */

func hasCycle2(head *ListNode) bool {
	visited := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := visited[head]; ok {
			return true
		}
		visited[head] = struct{}{}
		head = head.Next
	}
	return false
}

