package lc905

/*
分析：双指针+一次遍历，遍历一次数组，偶数从左边加入res，奇数从右边加入res，最后返回res
时间复杂度：O(n)
空间复杂度：O(1)
 */

func sortArrayByParity3(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	left, right := 0, l-1
	for _, num := range nums {
		if num % 2 == 0 {
			res[left] = num
			left++
		} else {
			res[right] = num
			right--
		}
	}
	return res
}