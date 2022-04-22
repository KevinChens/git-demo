package lc191

/*
分析：循环检查二进制位，直接循环检查给定整数n的二进制位的每一位是否为1
时间复杂度：O(k), k是int型二进制的位数
空间复杂度：O(1)
 */

func hammingWeight2(num uint32) int {
	res := 0
	for i := 0; i < 32; i++ {
		if 1 << i & num > 0 {
			res++
		}
	}
	return res
}
