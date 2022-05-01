package lc2248

import "sort"

/*
分析：map，记录每个数出现次数，次数与数组个数相等的数，就是每个数组的交集，sort.Ints排序返回res
时间复杂度：O(n)
空间复杂度：O(n)
 */

func intersection(nums [][]int) []int {
	freq := make(map[int]int, 0)
	for _, num := range nums {
		for _, v := range num {
			freq[v]++
		}
	}
	l := len(nums)
	res := make([]int, 0)
	for val, count := range freq {
		if count == l {
			res = append(res, val)
		}
	}
	sort.Ints(res)
	return res
}
