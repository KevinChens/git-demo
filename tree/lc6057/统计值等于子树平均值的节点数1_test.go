package lc6057

import (
	"fmt"
	"testing"
)

func TestAverageOfSubtree(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 0,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 1,
				Left: nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: nil,
			Right: &TreeNode{
				Val: 6,
				Left: nil,
				Right: nil,
			},
		},
	}
	ans := AverageOfSubtree(root)
	fmt.Println(ans)
}
