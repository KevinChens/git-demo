package lc28

import (
	"fmt"
	"testing"
)

func TestStrStr1(t *testing.T) {
	// 第二层循环没有懂，调试一下，看一下过程
	// haystack[i+j] needle[j] j是needle的位置，i是haystack的位置，相加进行偏移
	haystack := "hello"
	needle := "ll"
	res := StrStr1(haystack, needle)
	fmt.Printf("res:%d\n", res)
	if res != 2 {
		t.Errorf("error, res:%d", res)
	}
}
