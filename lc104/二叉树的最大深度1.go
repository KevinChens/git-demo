package lc104

/*
分析：遍历⼆叉树计算答案的思路，回溯思想，BFS；
遍历一次⼆叉树，⽤⼀个外部变量记录每个节点所在的深度，取最⼤值就得到最⼤深度。
时间复杂度：O(n), 只需要遍历一次二叉树，每个节点只会被访问一次，其中n为二叉树的节点个数。
空间复杂度：O(height), 递归遍历需要栈空间，大小为二叉树的深度height
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//记录最大深度
var res int
//记录遍历到的当前节点的深度
var depth int

func MaxDepth1(root *TreeNode) int {
	//全局变量初始化
	res = 0
	depth = 0
	//遍历二叉树计算树的最大深度
	traverse(root)
	return res
}

// 二叉树的递归遍历，回溯思想
func traverse(root *TreeNode) {
	if root == nil {
		//到达叶子节点，更新最大深度
		res = max(res, depth)
		/*匿名函数返回最大值
		res = func(x, y int) int {
			if x < y {
				return y
			}
			return x
		}(res, depth)
		 */
		//返回
		return
	}
	//前序位置，进入二叉树节点前执行，当前节点depth+1
	depth++
	traverse(root.Left)
	traverse(root.Right)
	//后序位置，离开二叉树节点后执行，当前节点depth-1
	depth--
}

// 返回两个数中的最大值
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}



