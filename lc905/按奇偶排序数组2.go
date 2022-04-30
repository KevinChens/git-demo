package lc905

/*
分析：遍历2次，第一次把偶数添加到res切片，第二次把奇数添加到res切片，最后返回res切片
时间复杂度：O(n)
空间复杂度：O(1)
 */

func sortArrayByParity2(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		if num % 2 == 0 {
			res = append(res, num)
		}
	}
	for _, num := range nums {
		if num % 2 != 0 {
			res = append(res, num)
		}
	}
	return res
}