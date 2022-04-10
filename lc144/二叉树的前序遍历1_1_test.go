package lc144

import (
	"fmt"
	"testing"
)

func TestPreorderTraversal1_1(t *testing.T) {
	//var root = &TreeNode{}, 并不是空树，默认val=0
	if res := PreorderTraversal1_1(nil); res != nil {
		fmt.Println(res)
		t.Errorf("error, %d", res)
	}

}
