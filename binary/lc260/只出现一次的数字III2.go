package lc260

/*
分析：hash table；
我们可以使用一个哈希映射统计数组中每一个元素出现的次数。
在统计完成后，我们对哈希映射进行遍历，将所有只出现了一次的数放入答案中；
时间复杂度：O(n)
空间复杂度：O(n)
 */

func singleNumber2(nums []int) []int{
	freq := map[int]int{}
	ans := make([]int, 0)
	for _, num := range nums {
		freq[num]++
	}
	for num, occ := range freq {
		if occ == 1 {
			ans = append(ans, num)
		}
	}
	return ans
}
