package codec

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

//要不断去保持每天的代码量或者写作
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		//用递归的方式，其实就不需要栈或者队列去保存了
		if root == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(root.Val))
		sb.WriteByte(',')
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	//二叉树解析,假设输入正常
	//其实就相当于一个局部的全局变量，可以当一个队列来用，先进先出遍历
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		val := sp[0]
		if val == "null" {
			sp = sp[1:]
			return nil
		}
		value, _ := strconv.Atoi(val)
		sp = sp[1:]
		return &TreeNode{value, build(), build()}
	}

	return build()

}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
