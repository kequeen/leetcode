package leetcode

import (
	"container/list"
)

//https://leetcode.cn/problems/valid-parentheses/?favorite=2cktkvj
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

//其实我理解就是一个栈，将数据放入栈中，如果下一个数据跟栈顶的数据是一对，则弹出
//否则就入栈
func isValid(s string) bool {
	lenS := len(s)
	if lenS == 0 {
		return true
	}
	pMap := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	//突然间感觉用这个栈的结构，不如直接用切片来得简单
	stack := list.New()
	stack.PushBack(s[0])
	for i := 1; i < lenS; i++ {
		top := stack.Back()
		_, ok := pMap[s[i]]
		if ok {
			//如果在里面
			stack.PushBack(s[i])
		} else {
			if top == nil {
				return false
			}
			//如果不在的话，就要看是否匹配了
			key := top.Value.(byte)
			if pMap[key] != s[i] {
				return false
			} else {
				stack.Remove(top)
			}

		}
	}
	if stack.Len() == 0 {
		return true
	}

	return false
}
