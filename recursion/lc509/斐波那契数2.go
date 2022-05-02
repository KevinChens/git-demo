package lc509

/*
分析：递归+备忘录，减少重复的递归
时间复杂度：O(n)
空间复杂度：O(n)
 */

func fib2(n int) int {
	return dfs(n)
}

// 缓存备忘录
var m = make(map[int]int)

func dfs(n int) int {
	if n < 2 {
		return n
	}
	// 读取缓存
	if m[n] != 0 {
		return m[n]
	}
	ans := dfs(n-2)+dfs(n-1)
	// 缓存已经计算过的值
	m[n] = ans
	return ans
}