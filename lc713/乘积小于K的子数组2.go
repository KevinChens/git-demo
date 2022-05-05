package lc713

import (
	"math"
	"sort"
)

/*
分析：
我们可以使用二分查找解决这道题目，即对于固定的i，二分查找出最大的j满足nums[i]到nums[j]的乘积小于k
不太懂，暂略
时间复杂度：O(nlogn), 其中n是nums数组的长度。二分查找的时间复杂度为O(logn)
空间复杂度：O(n), 用来存储前缀和数组prefix
 */

func numSubarrayProductLessThanK2(nums []int, k int) (ans int) {
	if k == 0 {
		return 0
	}
	n := len(nums)
	logPrefix := make([]float64, n+1)

	for i, num := range nums {
		logPrefix[i+1] = logPrefix[i] + math.Log(float64(num))
	}

	logK := math.Log(float64(k))
	for j := 1; j <= n; j++ {
		l := sort.SearchFloat64s(logPrefix[:j], logPrefix[j]-logK+1e-10)
		ans += j - l
	}
	return ans
}
