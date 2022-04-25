package lc155

/*
分析：用两个栈实现，一个最小栈始终保证最小值在顶部
时间复杂度：O(1), 所有操作都是O(1)
空间复杂度：O(n)
 */

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type MinStack struct {
	min []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min: make([]int, 0),
		stack: make([]int, 0),
	}
}

func (ms *MinStack) Push(val int) {
	min := ms.GetMin()
	if val < min {
		ms.min = append(ms.min, val)
	} else {
		// ? 感觉不需要，除非是为了维持两个栈同样大小？
		// 因为pop时，min，stack都要pop，需要同等大小，min作为填充
		// 或者每次pop比较是不是min，是就也要pop min
		ms.min = append(ms.min, min)
	}
	ms.stack = append(ms.stack, val)
}

func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.min = ms.min[:len(ms.min)-1]
}

func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return 0
	}
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	if len(ms.min) == 0 {
		return 1 << 31
	}
	return ms.min[len(ms.min)-1]
}


