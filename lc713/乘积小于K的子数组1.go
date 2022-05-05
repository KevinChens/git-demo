package lc713

/*
分析：
用双指针来维护一个滑动窗口，窗口不断向右滑动，窗口右边界(r)为固定轴，左边界(l)则是一个变动轴。
此窗口代表的意义为：
以窗口右边界为结束点的区间，其满足乘积小于k所能维持的最大窗口。
最后答案是是窗口在每个位置的最大长度的累计和。
因为针对上一位置的窗口，移动一次后相对增加出来的个数便是r-l+1
时间复杂度：O(n)
空间复杂度；O(1)
 */

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	left, right, ans := 0, 0, 0
	multi := 1
	
	for ; right < len(nums); right++ {
		multi *= nums[right]
		// 判断
		for multi >= k {
			multi /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}


