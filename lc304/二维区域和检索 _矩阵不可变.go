package lc304

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

/*
分析：初始化时对矩阵的每一行计算前缀和，检索时对二维区域中的每一行计算子数组和，然后对每一行的子数组和计算总和。
具体实现方面，创建m行n+1列的二维数组preSums，其中m和n分别是矩阵matrix的行数和列数，preSums[i]为matrix[i]的前缀和数组。
将preSums的列数设为n+1的目的是为了方便计算每一行的子数组和，不需要对col1=0的情况特殊处理。

时间复杂度：检索是O(m)
初始化O(mn)，需要遍历矩阵matrix计算二维前缀和；
每次检索O(m)，其中m和n分别是矩阵matrix的行数和列数。
检索需要对二维区域中的每一行计算子数组和，二维区域的行数不超过m，计算每一行的子数组和的时间复杂度是O(1)。
空间复杂度：O(mn)，其中m和n分别是矩阵matrix的行数和列数。
需要创建一个m行n+1列的前缀和数组preSums
 */

type NumMatrix struct {
	//记录NumMatrix中矩阵[0,0,i-1,j-1]的元素和
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	preSum := make([][]int, len(matrix))
	for i, row := range matrix {
		preSum[i] = make([]int, len(row)+1)
		//初始化，计算矩阵每一行前缀和
		for j, v := range row {
			preSum[i][j+1] = preSum[i][j] + v
		}
	}
	return NumMatrix{preSum: preSum}
}

func (nm *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	sum := 0
	//检索，计算每一行子数组和，并计算总和
	for i := row1; i <= row2; i++ {
		sum += nm.preSum[i][col2+1] - nm.preSum[i][col1]
	}
	return sum
}

/*
分析：二维前缀和，进一步优化检索过程
具体实现方面，创建m+1行n+1列的二维数组sums。
时间复杂度：
初始化O(mn)，需要遍历矩阵matrix计算二维前缀和；
每次检索O(1)，其中m和n分别是矩阵matrix的行数和列数。
空间复杂度：O(mn)，其中m和n分别是矩阵matrix的行数和列数。
需要创建一个m+1行n+1列的二维前缀和数组sums。
 */

type NumMatrix1 struct {
	preSums [][]int
}

func Constructor1(matrix [][]int) NumMatrix1 {
	ml := len(matrix)
	if ml == 0 {
		return NumMatrix1{}
	}

	nl := len(matrix[0])
	preSums := make([][]int, ml+1)
	preSums[0] = make([]int, nl+1)
	//计算二维数组前缀和
	for i, row := range matrix {
		preSums[i+1] = make([] int, nl+1)
		for j, v := range row {
			preSums[i+1][j+1] = preSums[i+1][j] + preSums[i][j+1] - preSums[i][j] + v
		}
	}
	return NumMatrix1{preSums: preSums}
}

func (nm *NumMatrix1) SumRegion1(row1, col1, row2, col2 int) int {
	return nm.preSums[row2+1][col2+1] - nm.preSums[row1][col2+1] - nm.preSums[row2+1][col1] + nm.preSums[row1][col1]
}

