package lc136

/*
分析：两个相同的数异或=0, 0和任何异或=任何
时间复杂度：O(n)
空间复杂度：O(1)
 */

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single = single ^ num
	}
	return single
}
