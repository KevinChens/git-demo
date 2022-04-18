package lc28

// 官方暴力匹配

func strStr1_1(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	outer:
		for i := 0; i+m <= n; i++ {
			for j := range needle {
				if haystack[i+j] != needle[j] {
					continue outer
				}
			}
			return i
		}
	// 不存在
	return -1
}
