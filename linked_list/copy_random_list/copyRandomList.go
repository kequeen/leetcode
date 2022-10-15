package leetcode

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//https://leetcode.cn/leetbook/read/top-interview-questions/xam1wr/
//复制带随机指针的链表
//一开始自己的想法太复杂了，复制完整个链表之后，再看随机指针的位置，然后在新的链表上再按照这个位置去指向

var nodeMap map[*Node]*Node

func copyRandomList(head *Node) *Node {
	nodeMap = map[*Node]*Node{}
	return deepCopy(head)
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}

	//如果已经创建过，那就返回
	n, ok := nodeMap[node]
	if ok {
		return n
	}

	//否则就进行节点的创建
	newNode := &Node{
		Val: node.Val,
	}
	nodeMap[node] = newNode
	//将问题不断递归
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}
