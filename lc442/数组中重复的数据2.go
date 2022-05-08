package lc442

/*
分析：使用正负号作为标记
我们也可以给nums[i]加上负号表示数i+1已经出现过一次。
具体地，我们首先对数组进行一次遍历。
当遍历到位置i时，我们考虑nums[nums[i]−1]的正负性
由于nums[i]本身可能已经为负数，因此在将nums[i]作为下标或者放入答案时，需要取绝对值
有点没懂
时间复杂度：O(n)
空间复杂度：O(1)
 */

func findDuplicates2(nums []int) (ans []int) {
	for _, x := range nums {
		if x < 0 {
			x = -x
		}
		if nums[x-1] > 0 {
			nums[x-1] = -nums[x-1]
		} else {
			ans = append(ans, x)
		}
	}
	return ans
}