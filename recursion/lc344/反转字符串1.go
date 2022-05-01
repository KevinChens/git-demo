package lc344

/*
分析：双指针两头交换
时间复杂度：O(n)
空间复杂度：O(1)
 */

func reverseString(s []byte) {
	l := len(s)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
