package lc278

/*
分析：二分搜索
时间复杂度：O(logn)
空间复杂度：O(1)
 */

func firstBadVersion(n int) int {
	start, end := 0, n
	for start+1 < end {
		mid := (start+end) / 2
		if isBadVersion(mid) {
			end = mid
		} else if isBadVersion(mid) == false {
			start = mid
		}
	}
	if isBadVersion(start) {
		return start
	}
	return end
}

func isBadVersion(x int) bool {
	// 略
	return true
}