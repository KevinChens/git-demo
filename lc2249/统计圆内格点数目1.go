package lc2249

import "sort"

/*
分析：遍历坐标系中的所有点，根据圆的方程过滤出落在圆上面的点
满足圆心为(a,b) 半径为r的圆的方程为(x-a)²+(y-b)²<=r²说明点(x,y)在圆内
先按半径从大到小排序，这样可以更早地遇到包含当前枚举的点的圆
时间复杂度：O(n^2)
空间复杂度：O(1)
 */

func countLatticePoints(circles [][]int) int {
	sort.Slice(circles, func(i, j int) bool {
		// 按半径从大到小排序
		return circles[i][2] > circles[j][2]
	})
	var res int
	// 遍历坐标系所有点
	for x := 0; x <= 200; x++ {
		for y := 0; y <= 200; y++ {
			for _, c := range circles {
				// 判断(x, y)是否在圆内
				if (x-c[0])*(x-c[0])+(y-c[1])*(y-c[1]) <= c[2]*c[2] {
					res++
					break
				}
			}
		}
	}
	return res
}