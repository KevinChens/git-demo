package lc1

/*
时间复杂度O(n^2)，两层遍历
空间复杂度O(1)，不需要额外空间
*/
func twoSum(nums []int, target int) []int {
	//穷举遍历
	for i, v := range nums {
		//记录第一个数的index
		index := i
		//计算第2个数
		num := target - v
		//查找第二个数的位置
		for j := i+1; j < len(nums); j++ {
			//找到返回下标
			if num == nums[j] {
				return []int{index, j}
			}
		}
	}
	//没找到返回nil
	return nil
}

/*
时间复杂度O(n)，遍历1次数组存储进map
空间复杂度O(n)，需要一个map
*/
func twoSum2(nums []int, target int) []int {
	//map存储index和value
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	//查找target
	for j, num := range nums {
		temp := target - num
		if index, ok := m[temp]; ok {
			//排除找到自身，target=2*num的情况
			if index == j {
				continue
			}else {
				return []int{index, j}
			}
		}
	}
	//没找到，返回nil
	return nil
}
//边存map，边找
func twoSum3(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		//map存v和i
		temp := target - v
		//判断第2个数是否存在
		if index, ok := m[temp]; ok {
			return []int{index, i}
		}else {
			//不存在插入当前数的index和value
			m[v] = i
		}
	}
	//没找到返回nil
	return nil
}
