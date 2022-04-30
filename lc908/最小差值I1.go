package lc908

/*
分析：数学知识，
更改后的整数数组nums的最低分数为maxNum−minNum−2k
时间复杂度：O(n)
空间复杂度：O(1)
 */

func smallestRangeI(nums []int, k int) int {
	minNum, maxNum := nums[0], nums[0]
	for _, num := range nums[1:] {
		if num < minNum {
			minNum = num
		} else if num > maxNum {
			maxNum = num
		}
	}
	res := maxNum - minNum - 2*k
	if res > 0 {
		return res
	}
	return 0
}
