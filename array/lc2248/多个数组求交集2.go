package lc2248

import "sort"

/*
分析：递归，分解为子数组进行判断，每次判断两个数组的交集，所以先写出求解两个数组交集的函数：
时间复杂度：O(kn*logk), k个数组，数组每个长度为n
空间复杂度：O(kn), k个数组
 */

func intersection2(nums [][]int) []int {
	res := recursion(nums, 0, len(nums)-1)
	sort.Ints(res)
	return res
}

func intersectionOfTwoArrays(nums1, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	// 将nums1加入set
	set := make(map[int]struct{})
	for _, v := range nums1 {
		set[v] = struct{}{}
	}
	// 遍历nums2，将set中相同的值，加入newSet, newSet是两个数组的交集
	newSet := make(map[int]struct{})
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			newSet[v] = struct{}{}
		}
	}
	//将newSet的值保存到res，并返回
	res := make([]int, 0)
	for v, _ := range newSet {
		res = append(res, v)
	}
	sort.Ints(res)
	return res
}

func recursion(nums [][]int, start, end int) []int {
	// 递归截至条件
	if start == end {
		return nums[start]
	}
	// 递归中间处理逻辑
	mid := (start+end) / 2
	left := recursion(nums, start, mid)
	right := recursion(nums, mid+1, end)
	return intersectionOfTwoArrays(left, right)
}
