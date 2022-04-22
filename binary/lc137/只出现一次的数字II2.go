package lc137

/*
分析：hash table保存出现次数，查找出现一次的即可
时间复杂度：O(n)
空间复杂度：O(n)
 */

func singleNumber2(nums []int) int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}
	for num, occ := range freq {
		if occ == 1 {
			return num
		}
	}
	// 只是作为一个返回值，实际上不会发生，一定会在hash table找到结果
	return 0
}
