package lc21

/*
分析：递归遍历
时间复杂度：O(n+m)
空间复杂度：O(n+m)
 */

func mergeTwoLists2(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		// 递归
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}else {
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}
}
