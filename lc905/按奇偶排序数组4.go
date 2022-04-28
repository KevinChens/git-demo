package lc905

/*
分析：原地交换
先从nums左侧开始遍历，如果遇到的是偶数，就表示这个元素已经排好序了，继续从左往右遍历，直到遇到一个奇数。
然后从nums右侧开始遍历，如果遇到的是奇数，就表示这个元素已经排好序了，继续从右往左遍历，直到遇到一个偶数。
交换这个奇数和偶数的位置，并且重复两边的遍历，直到在中间相遇，nums排序完毕
时间复杂度：O(n)
空间复杂度：O(1)
 */

func sortArrayByParity4(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left] % 2 == 0 {
			left++
		}
		for left < right && nums[right] % 2 == 1 {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}
