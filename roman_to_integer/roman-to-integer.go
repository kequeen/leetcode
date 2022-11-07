package leetcode

//https://leetcode.cn/problems/roman-to-integer/description/
//罗马数字转整数
func romanToInt(s string) int {
	res := 0
	lenS := len(s)
	if lenS == 0 {
		return res
	}

	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < lenS-1; i++ {
		//判定是否是那种特殊情况

		//其实结题思路里面做的更简单，只需要判断左边的数是否小于右边的数就好了，小于就是减法
		//https://leetcode.cn/problems/roman-to-integer/solutions/774992/luo-ma-shu-zi-zhuan-zheng-shu-by-leetcod-w55p/
		if s[i] == 'I' && (s[i+1] == 'V' || s[i+1] == 'X') {
			res += -1
		} else if s[i] == 'X' && (s[i+1] == 'L' || s[i+1] == 'C') {
			res += -10
		} else if s[i] == 'C' && (s[i+1] == 'D' || s[i+1] == 'M') {
			res += -100
		} else {
			value, ok := romanMap[s[i]]
			if !ok {
				return 0
			}
			res += value

		}
	}
	value, ok := romanMap[s[lenS-1]]
	if !ok {
		return 0
	}
	res += value
	return res
}
