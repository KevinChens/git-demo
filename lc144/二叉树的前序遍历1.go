package lc144


/*
分析：递归，根-左-右
定义preorder(root)表示当前遍历到root节点的答案。
按照定义，我们只要首先将root节点的值加入答案，
然后递归调用preorder(root.left)来遍历root节点的左子树，
最后递归调用preorder(root.right)来遍历root节点的右子树即可，
递归终止的条件为碰到空节点。
时间复杂度：O(n), 其中n是二叉树的节点数，每一个节点只被访问一次。
空间复杂度：O(n), 递归过程中栈空间的开销
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal1(root *TreeNode) []int {
	//当前遍历到节点的值
	var res []int
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		//前序位置
		res = append(res, root.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	//递归整个二叉树
	preorder(root)
	return res
}
