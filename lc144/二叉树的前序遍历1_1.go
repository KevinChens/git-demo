package lc144

//记录前序顺序
var res []int

func PreorderTraversal1_1(root *TreeNode) []int {
	//没有使用匿名函数，怎么就不对了
	//还是全局变量的问题，初始化一下
	res = nil
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	//前序位置
	res = append(res, root.Val)
	traverse(root.Left)
	traverse(root.Right)
}
