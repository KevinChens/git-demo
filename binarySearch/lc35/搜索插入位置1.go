package lc35

/*
分析：使用二分搜索模板3, 找到第一个>=target的元素位置
时间复杂度：O(logn)
空间复杂度：O(1)
 */

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := (start+end) / 2
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else if nums[mid] == target {
			start = mid
		}
	}
	if nums[start] >= target {
		return start
	} else if nums[end] >= target {
		return end
	} else if nums[end] < target {
		// target比所有值都大
		return end+1
	}
	return 0
}
