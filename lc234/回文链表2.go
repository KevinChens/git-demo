package lc234

/*
分析：将值复制到数组中后用双指针法
时间复杂度：O(n)
空间复杂度：O(n)
 */

func isPalindrome2(head *ListNode) bool {
	var vals []int

	for curr := head; curr != nil; curr = curr.Next {
		vals = append(vals, curr.Val)
	}

	l := len(vals)
	for i, v := range vals[:l/2] {
		if v != vals[l-1-i] {
			return false
		}
	}
	return true
}


