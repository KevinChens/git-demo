package lc28

/*
分析：遍历给定字符串字符haystack，判断以当前字符开头字符串是否等于目标字符串needle
暴力匹配
时间复杂度：O(n*m)，两层for循环遍历; 其中n是字符串haystack的长度，m是字符串needle的长度。
空间复杂度：O(1), 只需要常数的空间保存若干变量
 */

func StrStr1(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	// i不需要到len-1
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			//?
			if haystack[i+j] != needle[j] {
				break
			}
		}
		// 判断字符串长度是否相等
		//? 感觉多此一举，不用判断，需要，保证了将needle遍历完
		if len(needle) == j {
			return i
		}
	}
	// 不存在
	return -1
}
