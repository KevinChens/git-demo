package lc191

/*
分析：
时间复杂度：O(logn), 循环次数等于n的二进制位中1的个数
空间复杂度：O(1)
 */

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num &= num-1
		res++
	}
	return res
}
