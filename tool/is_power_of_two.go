package leetcode

// 判断一个数是否2的幂次方
func IsPowerOfTwo(n int) bool {
	// 这个写法写的还是比较巧妙的
	return n > 0 && (n&(n-1)) == 0
}
