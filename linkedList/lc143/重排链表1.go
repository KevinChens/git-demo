package lc143

/*
分析：找到中点断开，翻转后面部分，然后合并前后两个链表
时间复杂度：O(n)
空间复杂度：O(1)
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	mid := findMiddle(head)
	tail := reverseList(mid.Next)
	mid.Next = nil

	head = mergeTwoLists(head, tail)
}

func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	toggle := true

	for list1 != nil && list2 != nil {
		// 节点切换, list1, list2轮询添加
		if toggle {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		toggle = !toggle
		head = head.Next
	}

	if list1 == nil {
		head.Next = list2
	} else {
		head.Next = list1
	}
	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		// 保存next，防止被覆盖
		next := head.Next
		// 反转
		head.Next = prev
		// 后移
		prev = head
		head = next
	}
	return prev
}


