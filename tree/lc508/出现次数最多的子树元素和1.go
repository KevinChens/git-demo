package lc508

/*
分析：DFS
对于每棵子树，其子树元素和等于子树根结点的元素值，加上左子树的元素和，以及右子树的元素和。
用哈希表统计每棵子树的元素和的出现次数，计算出现次数的最大值maxCnt，最后将出现次数等于maxCnt 的所有元素和返回。
时间复杂度：O(n),n是二叉树的结点个数
空间复杂度：O(n), hash表和递归的栈空间均为O(n)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) (ans []int) {
	// 统计每棵子树的元素和的出现次数
	cnt := map[int]int{}
	maxCnt := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + dfs(node.Left) + dfs(node.Right)
		cnt[sum]++
		if cnt[sum] > maxCnt {
			maxCnt = cnt[sum]
		}
		return sum
	}
	dfs(root)
	// maxCnt的key进行返回
	for s, c := range cnt {
		if c == maxCnt {
			ans = append(ans, s)
		}
	}
	return
}
