package lc704

/*
分析：二分模板
时间复杂度：O(logn)
空间复杂度：O(1)
 */

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start+end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		}
	}
	return -1
}

