package lc148

/*
分析：归并排序，找中点，合并
时间复杂度：O(n*logn)
空间复杂度：O(logn), 空间复杂度主要取决于递归调用的栈空间
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

// 快慢指针找中间值

func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	// 快指针先为nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	if list1 == nil {
		head.Next = list2
	} else {
		head.Next = list1
	}
	return dummy.Next
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 找中间节点
	middle := findMiddle(head)
	// 断开中间节点
	tail := middle.Next
	middle.Next = nil
	// 递归
	left := mergeSort(head)
	right := mergeSort(tail)
	// 合并
	result := mergeTwoLists(left, right)
	return result
}
