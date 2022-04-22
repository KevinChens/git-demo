package lc138

/*
分析：回溯+hash table;
利用回溯的方式，让每个节点的拷贝操作相互独立；
对于curr，我们要对Next和Random进行拷贝，
然后进行curr的后继节点和curr的随机指针指向的节点拷贝，拷贝完成后将创建的cloneHead的指针返回；
用hash table记录每一个节点对应新节点的创建情况
时间复杂度：O(n), 每个节点最多被访问2次
空间复杂度：O(n)
 */

var cacheNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	// 避免重复拷贝，直接返回指针即可
	if n, has := cacheNode[node]; has {
		return n
	}
	// 节点拷贝
	cloneHead := &Node{Val: node.Val}
	cacheNode[node] = cloneHead
	// Next和Random指针拷贝
	cloneHead.Next = deepCopy(node.Next)
	cloneHead.Random = deepCopy(node.Random)
	return cloneHead
}

func copyRandomList2(head *Node) *Node {
	cacheNode = map[*Node]*Node{}
	return deepCopy(head)
}