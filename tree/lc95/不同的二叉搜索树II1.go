package lc95

/*
分析:
出口为当start>end 的时候，当前二叉搜索树为空，返回空节点即可
不太懂
时间复杂度：O(n*Gn), Gn为卡特兰数，整个算法的时间复杂度取决于可行二叉搜索树的个数，
对于n个点生成的二叉搜索树数量等价于数学上第n个卡特兰数，用Gn表示
空间复杂度：O(n*Gn), n个点生成的二叉搜索树有Gn棵，每棵有n个节点，因此存储的空间需要O(n*Gn)
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
		// return nil不行吗
		// 不行
	}
	ans := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		// 递归生成所有左右子树
		lefts := generate(start, i-1)
		rights := generate(i+1, end)
		// 拼接左右子树后返回
		for j := 0; j < len(lefts); j++ {
			for k := 0; k < len(rights); k++ {
				root := &TreeNode{
					Val: i,
					Left: nil,
					Right: nil,
				}
				root.Left = lefts[j]
				root.Right = rights[k]
				ans = append(ans, root)
			}
		}
	}
	return ans
}
