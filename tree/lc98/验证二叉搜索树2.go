package lc98

/*
分析：分治法，	判断左 MAX < 根 < 右 MIN
时间复杂度：O(n)
空间复杂度：O(n)
 */

type ResultType struct {
	IsValid bool
	// 记录左右两边最大最小值，和根节点进行比较
	Max *TreeNode
	Min *TreeNode
}
func isValidBST2(root *TreeNode) bool {
	result := helper(root)
	return result.IsValid
}

func helper(root *TreeNode) ResultType {
	result := ResultType{}
	// check
	if root == nil {
		result.IsValid = true
		return result
	}

	left := helper(root.Left)
	right := helper(root.Right)

	if !left.IsValid || !right.IsValid {
		result.IsValid = false
		return result
	}
	// 判断
	if left.Max != nil && left.Max.Val >= root.Val {
		result.IsValid = false
		return result
	}
	if right.Min != nil && right.Min.Val <= root.Val {
		result.IsValid = false
		return result
	}

	result.IsValid = true
	// 如果左边还有更小的，用更小的节点
	result.Min = root
	if left.Min != nil {
		result.Min = left.Min
	}
	result.Max = root
	if right.Max != nil {
		result.Max = right.Max
	}
	return result
}
