package lc76

import "math"

/*
分析：
我们可以用一个哈希表表示t中所有的字符以及它们的个数，
用一个哈希表动态维护窗口中所有的字符以及它们的个数，
如果这个动态表中包含t的哈希表中的所有字符，并且对应的个数都不小于t的哈希表中各个字符的个数，
那么当前的窗口是可行的。
难题，不太懂
时间复杂度：O(C⋅∣s∣+∣t∣)，最坏情况下左右指针对s的每个元素各遍历一遍，
哈希表中对s中的每个元素各插入、删除一次，对t中的元素各插入一次。
每次检查是否可行会遍历整个t的哈希表，哈希表的大小与字符集的大小有关，设字符集大小为C
则渐进时间复杂度为O(C⋅∣s∣+∣t∣)
空间复杂度：O(C)，这里用了两张哈希表作为辅助空间，每张哈希表最多不会存放超过字符集大小的键值对，
我们设字符集大小为C ，则渐进空间复杂度为O(C)
 */

func miniWindow(s, t string) string {
	// 保存滑动窗口字符集
	win := make(map[byte]int)
	// 保存需要的字符集
	need := make(map[byte]int)
	for _, c := range t {
		need[byte(c)]++
	}
	// 窗口
	left, right := 0, 0
	// match记录匹配次数
	match := 0
	start, end := 0, 0
	min := math.MaxInt64
	var c byte
	for right < len(s) {
		c = s[right]
		right++
		// 在需要的字符集里面，添加到窗口字符集里面
		if need[c] != 0 {
			win[c]++
			// 如果当前字符的数量匹配需要的字符的数量，则match+1
			if win[c] == need[c] {
				match++
			}
		}
		// 当所有字符数量都匹配之后，开始缩紧窗口
		for match == len(need) {
			// 获取结果
			if right-left < min {
				min = right-left
				start = left
				end = right
			}
			c = s[left]
			left++
			// 左指针指向不再需要字符集则直接跳过
			if need[c] != 0 {
				// 左指针指向字符数量和需要的字符相等时，右移之后match值就不匹配则减一
				// 因为win里面的字符数可能比较多，如有10个A，但需要的字符数量可能为3
				// 所以在压死骆驼的最后一根稻草时，match才减一，这时候才跳出循环
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
		}
	}
	if min == math.MaxInt64 {
		return ""
	}
	return s[start:end]
}
