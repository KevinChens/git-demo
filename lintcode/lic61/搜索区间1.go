package lic61

/*
分析：核心点就是找第一个target的索引，和最后一个target的索引，所以用两次二分搜索分别找第一次和最后一次的位置
时间复杂度：O(logn)
空间复杂度：o(1)
 */

func SearchRange(a []int, target int) []int {
	if len(a) == 0 {
		return []int{-1, -1}
	}
	res := make([]int, 0)
	start, end := 0, len(a)-1

	for start+1 < end {
		mid := (start+end) / 2
		if a[mid] < target {
			start = mid
		} else if a[mid] > target {
			end = mid
		} else if a[mid] == target {
			// 继续向左边找，找到第一个target的位置
			end = mid
		}
	}
	// 搜索左边的索引
	if a[start] == target {
		res = append(res, start)
	} else if a[end] == target {
		res = append(res, end)
	} else {
		res = append(res, -1, -1)
		return res
	}

	start, end = 0, len(a)-1
	for start+1 < end {
		mid := (start+end) / 2
		if a[mid] < target {
			start = mid
		} else if a[mid] > target {
			end = mid
		} else if a[mid] == target {
			// 继续向右边找，找到最后一个target的位置
			start = mid
		}
	}
	// 搜索右边的索引
	if a[end] == target {
		res = append(res, end)
	} else if a[start] == target {
		res = append(res, start)
	} else {
		res = append(res, -1, -1)
		return res
	}
	return res
}
