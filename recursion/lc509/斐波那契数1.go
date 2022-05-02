package lc509

/*
分析：纯递归
时间复杂度：O(2^n)
空间复杂度：O(n)
 */

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1)+fib(n-2)
}