package lc509

/*
分析：dp+滚动数组, 神奇的dp
state:每一项的和等于前两项的和
function:f(n)=f(n-1)+f(n-2)
initialization：f(0), f(1)
answer:f(n-1)
时间复杂度：O(n)
空间复杂度：(1), 滚动数组让空间复杂度从O(n)到O(1)
 */

func fib3(n int) int {
	if n < 2 {
		return n
	}
	// ?为什么是这个值
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}


