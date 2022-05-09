package lc937

import (
	"sort"
	"strings"
	"unicode"
)

/*
分析：自定义排序，按照题目要求排序即可
对于go的sort库不熟悉
时间复杂度：O(n*logn), 其中n是logs的字符数，是平均情况下内置排序的时间复杂度
空间复杂度：O(n), 需要新建数组保存log和下标，需要将每条log进行拆分
 */

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		s, t := logs[i], logs[j]
		s1 := strings.SplitN(s, " ", 2)[1]
		s2 := strings.SplitN(t, " ", 2)[1]
		isDigit1 := unicode.IsDigit(rune(s1[0]))
		isDigit2 := unicode.IsDigit(rune(s2[0]))
		if isDigit1 && isDigit2 {
			return false
		}
		if !isDigit1 && !isDigit2 {
			return s1 < s2 || s1 == s2 && s < t
		}
		return !isDigit1
	})
	return logs
}