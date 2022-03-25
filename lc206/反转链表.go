package lc206

// ListNode 定义链表
type ListNode struct {
	Val int
	Next *ListNode
}
/*
时间复杂度O(n)，需要反转n个节点；
空间复杂度O(n)，需要递归调用n层
*/
func reverseList(head *ListNode) *ListNode {
	//0和1个节点不用反转
	if head == nil || head.Next == nil { return head }
	//2个及以上需要递归反转
	newHead := reverseList(head.Next)
	//2个节点之间如何反转
	//第2个节点的next指向第1个节点，第1个节点的next指向nil
	head.Next.Next = head
	head.Next = nil
	//返回newHead
	return newHead
}

/*
时间复杂度O(n)，需要遍历一次链表
空间复杂度O(1)，不需要额外的辅助空间
*/
func reverseList2(head *ListNode) *ListNode {
	//迭代反转链表
	//prev-->curr-->next
	//要保存prev,next
	//将curr的next指向prev
	//prev，curr同步后移，迭代
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		//开始反转
		curr.Next = prev
		//prev,curr同步后移
		prev = curr
		curr = next
	}
	//返回新头节点prev
	return prev
}
