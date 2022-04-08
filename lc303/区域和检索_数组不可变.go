package lc303

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */


/*
时间复杂度：O(n), 每次调用sumRange时，都遍历left到right之间的元素，进行累加
如果sumRange方法被反复调用，每次都是O(n)，查询的代价比较大
空间复杂度：O(n), 需要创建一个长度为n的数组NumArray
 */

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

// SumRange 未优化，逐次遍历
func (this *NumArray) SumRange(left, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.nums[i]
	}
	return sum
}

/*
时间复杂度：O(1), 使⽤前缀和技巧
初始化O(n)，需要遍历数组nums计算前缀和, 其中n是数组nums的长度;
每次检索O(1)，只需要得到两个下标处的前缀和，然后计算差值。
空间复杂度：O(n), 需要创建一个长度为n+1的前缀和数组preSums
*/

type NumArray1 struct {
	preSums []int
}

func Constructor1(nums []int) NumArray1 {
	preSums := make([]int, len(nums)+1)
	//构造前缀和
	for i, v := range nums {
		preSums[i+1] = preSums[i] + v
	}
	return NumArray1{preSums: preSums}
}

func (this *NumArray1) SumRange1(left, right int) int {
	return this.preSums[right+1] - this.preSums[left]
}