package lc104

/*
分析：官方的答案对BFS有些修改，使用queue存放当前层的所有节点；
每次拓展下一层时，不同于BFS的每次只从队列里拿出一个节点，
还需要将队列里的所有节点都拿出来进行拓展，这样能保证每次拓展完的时候队列里存放的是当前层的所有节点，
一层一层地进行拓展，最后用一个变量ans来维护拓展的次数，也就是二叉树的最大深度。
时间复杂度：O(n), 每个节点只访问一次
空间复杂度：O(n), 使用了queue保存了当前层的所有节点
*/

func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//queue保存当前层所有节点
	var queue []*TreeNode
	queue = append(queue, root)
	//记录最大深度
	ans := 0
	//层序遍历，并一层一层扩展
	for len(queue) > 0 {
		//当前层节点个数
		sz := len(queue)
		for sz > 0 {
			//出队
			node := queue[0]
			queue = queue[1:]
			//下一层节点存放
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			//出队后，节点个数减一
			sz--
		}
		//每层拓展，深度+1
		ans++
	}
	return ans
}


