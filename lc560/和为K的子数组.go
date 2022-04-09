package lc560

/*
分析：把所有⼦数组都穷举出来，算它们的和，看看谁的和等于k，借助前缀和技巧很容易得到
时间复杂度：O(n^2), 需要穷举所有子数组
空间复杂度：O(n), 需要n+1的前缀和数组
 */

func subarraySum(nums []int, k int) int {
	nl := len(nums)
	//构造前缀和
	preSum := make([]int, nl+1)
	for i := 0; i < nl; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	res := 0
	//穷举所有子数组
	for i := 1; i <= nl; i++ {
		for j := 0; j < i; j++ {
			if preSum[j] - preSum[i] == k {
				res++
			}
		}
	}
	return res
}

/*
分析：⽤哈希表，在记录前缀和的同时记录该前缀和出现的次数
时间复杂度：O(n), 使用hash表减少了内层循环遍历
空间复杂度：O(n), 需要一个n大小的hash表
 */

func subarraySum1(nums []int, k int) int {
	nl := len(nums)
	//preSum是一个map：key：前缀和；value：出现次数
	preSum := make(map[int]int)
	preSum[0] = 1

	res := 0
	sumI := 0
	for i := 0; i < nl; i++ {
		sumI += nums[i]
		//前缀和：nums[0...j]
		sumJ := sumI - k
		//若有前缀和，则更新答案
		if v, ok := preSum[sumJ]; ok {
			res += v
		}else {
			//把前缀和nums[0...i]加入map并记录出现次数
			preSum[sumI] += 1
		}
	}
	return res
}