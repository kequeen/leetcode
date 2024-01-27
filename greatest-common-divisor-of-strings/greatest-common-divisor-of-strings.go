package leetcode

// 两个字符串的最大公约数
// https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/
// 其实就是将问题转换为求两个数的最大公约数的过程
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

// 求最大公约数
func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
