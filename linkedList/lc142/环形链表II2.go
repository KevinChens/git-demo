package lc142

/*
分析：我们遍历链表中的每个节点，并将它记录下来；一旦遇到了此前遍历过的节点，就可以判定链表中存在环。
时间复杂度：O(n)
空间复杂度：O(n)
 */

func detectCycle2(head *ListNode) *ListNode {
	visited := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := visited[head]; ok {
			return head
		}
		visited[head] = struct{}{}
		head = head.Next
	}
	return nil
}
