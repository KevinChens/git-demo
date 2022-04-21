package lc86


/*
分析：模拟，small，large拼接
时间复杂度：O(n)
空间复杂度：O(1)
 */

func partition2(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large

	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	// 拼接
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}