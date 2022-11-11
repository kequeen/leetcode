package leetcode

//https://leetcode.cn/problems/palindrome-number/
//判断回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	//其实最简单的办法应该是判断反转后的数是否与原来的相等
	//当然放到数组中再前后遍历也行
	newNum := 0
	num := x
	for num != 0 {
		temp := num % 10
		newNum = newNum*10 + temp
		num = num / 10
	}
	if newNum == x {
		return true
	}
	return false
}

//不过题解里面只用反转一半其实效率更高
