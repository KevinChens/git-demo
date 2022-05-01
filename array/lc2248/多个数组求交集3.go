package lc2248

import "sort"

/*
分析：set是一个集合，set里的元素不能重复；
golang没有set，利用map key唯一，实现set；map的value不需要，为了节省内存空间，value可为空结构；
在Go中，空结构通常不使用任何内存；set := make(map[int]struct{})

首先用一个set把第一行数字存进去, 然后从第二行开始遍历，层级遍历;
把每层和上一层一样的数字存进一个新的set里，然后更新set用于下一行查找，
最后set里面的数字数目就是全部的交集
时间复杂度：O(n), 遍历一次二维数组
空间复杂度：O(n)，需要额外空间set
 */

func intersection3(nums [][]int) []int {
	// 将第一行数字加入set
	set := make(map[int]struct{})
	for _, v := range nums[0] {
		set[v] = struct{}{}
	}

	for i := 1; i < len(nums); i++ {
		newSet := make(map[int]struct{})
		// 遍历后面每行数字，判断是否有与set相同的值
		// 将set中相同的值，加入newSet，不断更新set
		for _, v := range nums[i] {
			if _, ok := set[v]; ok {
				newSet[v] = struct{}{}
			}
		}
		set = newSet
	}
	// 将set的值保存到res，并排序返回
	res := make([]int, 0)
	for v, _ := range set {
		res = append(res, v)
	}
	sort.Ints(res)
	return res
}
