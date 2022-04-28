package lc905

/*
分析：遍历一次数组，将偶数添加到even切片，将奇数添加到odd切片，最后将even和odd拼接返回
时间复杂度：O(n)，遍历一次数组
空间复杂度：O(1), 结果不计空间复杂度？
 */

func sortArrayByParity(nums []int) []int {
	even := make([]int, 0)
	odd := make([]int, 0)
	for _, num := range nums {
		if num % 2 != 0 {
			odd = append(odd, num)
		} else {
			even = append(even, num)
		}
	}
	even = append(even, odd...)
	return even
}
