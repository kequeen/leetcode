package leetcode

//矩阵的最短路径和
//https://leetcode.cn/problems/minimum-path-sum/?favorite=2cktkvj
// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
func minPathSum(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	if row == 0 || column == 0 {
		return 0
	}
	dp := make([][]int, row)
	for k := 0; k < row; k++ {
		dp[k] = make([]int, column)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if i == 0 && j == 0 {
				continue
			}
			//判断边界问题
			if i-1 < 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j-1 < 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[row-1][column-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
