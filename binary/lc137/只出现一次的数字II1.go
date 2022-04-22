package lc137

/*
分析：
由于数组中的元素都在int（即32位整数）范围内，因此我们可以依次计算答案的每一个二进制位是0还是1
对于数组中的每一个元素x，我们使用位运算(x >> i)&1得到x的第i个二进制位，
并将它们相加再对3取余，得到的结果一定为0或1，即为答案的第i个二进制位


时间复杂度：O(n), 遍历一次数组
空间复杂度：O(1)
 */

func singleNumber(nums []int) int {
	// 统计每位1的个数
	var res int
	for i := 0; i < 64; i++ {
		sum := 0
		for _, num := range nums {
			// 统计1的个数
			sum += (num >> i) & 1
		}
		// 还原位00^10=10, 或者用|也可以
		res ^= (sum % 3) << i
		// result |= (sum % 3) << i
	}
	return res
}
