package lc876

import "fmt"

/*
分析：快慢指针，fast = head.Next slow = head; fast到达终点，slow.Next即为中点
时间复杂度：O(n)
空间复杂度：O(1)
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast == nil {
		fmt.Println(slow.Next)
		return slow
	}
	fmt.Println(slow.Next.Next)
	return slow.Next
}