package lc74

/*
分析：可以将二维转为一维，进行二分查找
时间复杂度：O(logn)
空间复杂度：O(1)
 */

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	start, end := 0, row*col-1
	for start+1 < end {
		mid := (start+end) / 2
		val := matrix[mid/col][mid%col]
		if val > target {
			end = mid
		} else if val < target {
			start = mid
		} else {
			return true
		}
	}
	if matrix[start/col][start%col] == target || matrix[end/col][end%col] == target {
		return true
	}
	return false
}
