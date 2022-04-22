package lc260

/*
分析：位运算
对于任意一个在数组nums中出现两次的元素，该元素的两次出现会被包含在同一类中；
对于任意一个在数组nums中只出现了一次的元素，即x1和x2 ，它们会被包含在不同类中。
因此，如果我们将每一类的元素全部异或起来，那么其中一类会得到x1，另一类会得到x2。
这样我们就找出了这两个只出现一次的元素。
时间复杂度：O(n)
空间复杂度：O(1)
 */

func singleNumber(nums []int) []int {
	// xorSum = a^b, 接下来将a b分开即可
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	// 去掉末尾的1后异或diff就得到最后一个1的位置
	//diff = ( diff & (diff-1) ) ^ diff

	// 也可以使用位运算x & -x 取出x的二进制表示中最低位那个1
	lsb := xorSum ^ -xorSum
	a, b := 0, 0

	for _, num := range nums {
		// 分两类，分别全体异或，得到a，b
		if lsb & num == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
