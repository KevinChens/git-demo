package lc74

import "sort"

/*
分析：两次二分查找，根据特性
对矩阵的第一列的元素二分查找，找到最后一个不大于目标值的元素，
然后在该元素所在行中二分查找目标值是否存在。
可以暂时忽略，没太get
时间复杂度：O(logmn)
空间复杂度：O(1)
 */

func searchMatrix2(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	}) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}
