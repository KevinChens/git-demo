package lc104

import (
	"fmt"
	"testing"
)

func TestMaxDepth1(t *testing.T) {
	//[1, nil, 2]
	//测试能得到正确答案，但leetcode就是不通过
	/*原因：使用了全局变量，力扣判题机制会对每个测试用例，都会初始化一次，测试用例共享全局变量和类变量
	  解决：避免申明类内静态变量以及全局变量；在函数入口重新初始化全局变量
	 */
	var root = &TreeNode{
		Val: 1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: nil,
			Right: nil,
		},
	}
	res := MaxDepth1(root)
	fmt.Println(res)
	if res := MaxDepth1(root); res != 2 {
		t.Errorf("maxDepth is expected 2, but %d got", res)
	}

}
