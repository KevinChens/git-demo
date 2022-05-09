package lc442

/*
分析：对时空复杂度有限制，感觉要用math相关知识
将元素交换到对应的位置, 巧妙的条件
由于给定的n个数都在[1,n]的范围内，如果有数字出现了两次，就意味着[1,n]中有数字没有出现过。
我们对数组进行一次遍历，进行放置。
当遍历到位置i时，我们知道nums[i]应该被放在位置nums[i]−1。
因此我们交换num[i]和nums[nums[i]−1]即可，直到待交换的两个元素相等为止。
时间复杂度：O(n)
空间复杂度：O(1)
 */

func findDuplicates(nums []int) (ans []int) {
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if num-1 != i {
			ans = append(ans, num)
		}
	}
	return ans
}
