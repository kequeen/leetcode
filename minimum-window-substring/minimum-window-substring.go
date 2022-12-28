package leetcode

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
// https://leetcode.cn/problems/minimum-window-substring/?favorite=2cktkvj
// 最简单的可以用暴力方式，m * n的时间复杂度

// 感觉这个方法一直打补丁也是有些问题
// 最终还是会超时，某些测试用例有问题
func minWindow(s string, t string) string {
	lenS := len(s)
	lenT := len(t)
	beg := 0
	end := 0
	minW := lenS
	if lenT == 0 || lenS < lenT {
		return ""
	}
	//存储所有数据
	sMap := make(map[byte]int)
	count := 0
	flag := false
	for i := 0; i < lenT; i++ {
		sMap[t[i]]++
		count++
	}
	copyCount := count
	copySmap := make(map[byte]int)
	copyMap(sMap, copySmap)
	//从左到右遍历，不断遍历剪枝，假如遍历到count已清零，则直接返回
	for i := 0; i <= lenS-lenT; i++ {
		if copySmap[s[i]] == 0 {
			continue
		}
		copySmap[s[i]]--
		copyCount--
		if copyCount == 0 {
			beg = i
			end = i
			flag = true
		}
		for j := i + 1; j < lenS && copyCount != 0; j++ {

			if copySmap[s[j]] != 0 {
				copyCount--
				copySmap[s[j]]--
			}
			if copyCount == 0 {
				if j-i < minW {
					minW = j - i
					beg = i
					end = j
					flag = true
				}
				break
			}
		}

		//重置初始状态
		copyMap(sMap, copySmap)
		copyCount = count
	}
	if !flag {
		return ""
	}
	return s[beg : end+1]
}

func copyMap(map1, map2 map[byte]int) {
	for k, v := range map1 {
		map2[k] = v
	}
}

// 没想到最终还是滑动窗口，其实字符串比较，或者是子串的问题，基本都是滑动窗口，应该主动往这方面去想
func minWindowV2(s string, t string) string {
	//其实就是不停滑动窗口
	lenS := len(s)
	lenT := len(t)
	if lenT == 0 || lenS < lenT {
		return ""
	}
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)
	for i := 0; i < lenT; i++ {
		tMap[t[i]]++
	}

	//校验内容是否相同
	check := func() bool {
		for k, v := range tMap {
			if sMap[k] < v {
				return false
			}
		}
		return true
	}
	//开始从左到右遍历，满足条件之后，就可以左边开始收缩
	ansL := -1
	ansR := -1
	minLen := lenS + 1
	for l, r := 0, 0; r < lenS; r++ {
		if l <= r && tMap[s[r]] > 0 {
			sMap[s[r]]++
		}
		//如果满足条件，就开始收缩左边，寻求当前满足条件的最短字符串
		for check() && l <= r {
			if r-l+1 < minLen {
				minLen = r - l + 1
				ansL = l
				ansR = l + minLen
			}
			//如果在sMap中，则做减法
			if _, ok := sMap[s[l]]; ok {
				sMap[s[l]]--
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}

	return s[ansL:ansR]
}
