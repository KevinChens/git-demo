package lc138

/*
分析：迭代+拆分，复制节点跟在原节点后面，再处理random，最后拆分
时间复杂度：O(n)，遍历链表3次
空间复杂度：O(1)
 */

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 复制节点跟在原节点后面
	// 1->2->3  ==>  1->1'->2->2'->3->3'
	for curr := head; curr != nil; curr = curr.Next.Next {
		clone := &Node{Val: curr.Val, Next: curr.Next}
		curr.Next = clone
	}

	// 处理random指针
	for curr := head; curr != nil; curr = curr.Next.Next {
		if curr.Random != nil {
			// 为什么是curr.Random.Next? 不应该是curr.Random吗
			// 连接节点，curr.Random的值是curr.Random.Next
			curr.Next.Random = curr.Random.Next
		}
	}
	// 分离两个链表
	cloneHead := head.Next
	for curr := head; curr != nil && curr.Next != nil; curr = curr.Next {
		// 将next分离出来
		node := curr.Next
		curr.Next = curr.Next.Next
		if node.Next != nil {
			node.Next = node.Next.Next
		}
	}
	// 原始链表头：head: 1->2->3
	// 克隆链表头：cloneHead: 1'->2'->3'
	return cloneHead
}


