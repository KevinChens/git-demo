package lc153

/*
分析：二分搜索，最后一个值作为target，然后往左移动，最后比较start、end的值
时间复杂度：O(logn)
空间复杂度：O(1)
 */

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1

	for start+1 < end {
		mid := (start+end) / 2
		// 最后一个数作为target
		target := nums[end]
		if nums[mid] <= target {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] > nums[end] {
		return nums[end]
	}
	return nums[start]
}
