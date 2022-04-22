package lc234

/*
分析：快慢指针，中间断开，反转比较，判断是否值相等
时间复杂度：O(n)
空间复杂度：O(1)
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	// fast初始化为head.Next, 中点为slow.Next
	// fast初始化为head, 中点为slow
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 找中点，断开两个链表（需要用到中点前一个节点）
	tail := reverse(slow.Next)
	slow.Next = nil

	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		// 反转
		curr.Next = prev
		// 后移
		prev = curr
		curr = next
	}
	return prev
}
