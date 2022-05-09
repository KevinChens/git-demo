package lc942

/*
分析：贪心
考虑perm[0]问题，其他都是相同的不同规模的问题
s[0]='I', perm[0]=0;
s[0]='D', perm[0]=n
时间复杂度：O(n)
空间复杂度：O(1)
 */

func diStringMatch(s string) []int {
	n := len(s)
	perm := make([]int, n+1)
	min, max := 0, n

	for i, ch := range s {
		if ch == 'I' {
			perm[i] = min
			min++
		} else {
			perm[i] = max
			max--
		}
	}
	// 最后只剩一个数min=max
	perm[n] = min
	return perm
}
