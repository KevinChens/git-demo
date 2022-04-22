package lc136

/*
分析：hash table记录频次，找到出现次数为1的返回即可
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
	return 0
}