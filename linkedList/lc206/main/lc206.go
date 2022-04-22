package main

import "fmt"

// ListNode 定义链表
type ListNode struct {
	Val int
	Next *ListNode
}

// NewHead 生成头节点
func NewHead() *ListNode {
	return &ListNode{}
}

// ShowList 打印链表的值
func (head *ListNode) ShowList() {
	curr := head
	//head为虚拟头节点
	//curr := head.Next
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

// Append 尾插法（顺序）插入节点
func (tail *ListNode) Append(value int) {
	curr := tail
	node := &ListNode{Val: value}
	//链表为nil，直接插入第1个位置
	if curr == nil {
		curr = node
	}else {
		//链表不空，尾插法插入节点
		for curr.Next != nil {
			curr = curr.Next
		}
		//找到尾巴插入节点
		curr.Next = node
	}
}

//反转链表
/*
时间复杂度O(n)，需要反转n个节点；
空间复杂度O(n)，需要递归调用n层
*/
func reverseList(head *ListNode) *ListNode {
	//0和1个节点不用反转
	if head == nil || head.Next == nil { return head}
	//2个及以上递归反转
	newHead := reverseList(head.Next)
	//2个节点之间如何反转，第2个节点next指向第1个节点，第1个节点next指向nil
	head.Next.Next = head
	head.Next = nil
	//返回newHead
	return newHead
}
/*
时间复杂度O(n)，需要遍历一次链表
空间复杂度O(1)，不需要额外的辅助空间
*/
func reverseList2(head *ListNode) *ListNode {
	//迭代反转链表
	//prev-->curr-->next
	//要保存prev(用于curr指向)和next(用于curr移动)，将curr的next指向prev，prev和curr同步后移迭代
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		//开始反转
		curr.Next = prev
		//prev，curr同步后移迭代
		prev = curr
		curr = next
	}
	//返回新头节点prev
	return prev
}

func main() {
	head := NewHead()
	head.Append(1)
	head.Append(2)
	head.Append(3)
	head.Append(4)
	head.Append(5)

	fmt.Println("head")
	head.ShowList()
	//resHead := reverseList(head)
	resHead := reverseList2(head)
	fmt.Println("reverse head")
	resHead.ShowList()
}
