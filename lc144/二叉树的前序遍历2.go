package lc144

/*
分析：迭代，用迭代的方式实现解法1的递归函数，两种方式是等价的，
区别在于递归的时候隐式地维护了一个栈，而我们在迭代的时候需要显式地将这个栈模拟出来
时间复杂度：O(n), 其中n是二叉树的节点数。每个节点只被访问一次。
空间复杂度：O(n), 迭代过程中显式栈的开销
 */

func preorderTraversal2(root *TreeNode) (res []int) {
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			//前序：根-左
			node = node.Left
		}
		//出栈,根-右
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}