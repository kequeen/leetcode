package leetcode

// https://leetcode.cn/problems/pascals-triangle/description/?envType=study-plan-v2&envId=top-100-liked
// 杨辉三角
// 不过这个写法并不算很优雅，因为本身对于golang的一些初始化的问题，这里描述得不算特别清晰
func generate(numRows int) [][]int {
	ans := [][]int{}
	last := []int{1}
	ans = append(ans, last)
	for i := 2; i <= numRows; i++ {
		last := ans[i-2]
		temp := []int{}
		for j := 0; j < i; j++ {
			//常规的情况，但左右需要特殊处理
			if j == 0 {
				temp = append(temp, 1)
			} else if j == i-1 {
				temp = append(temp, 1)
			} else {
				temp = append(temp, last[j-1]+last[j])
			}
		}
		ans = append(ans, temp)
	}
	return ans
}

// 更加优雅一些的写法
func generateV2(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i, _ := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}
