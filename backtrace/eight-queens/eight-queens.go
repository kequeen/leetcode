package leetdoce

// https://leetcode.cn/problems/eight-queens-lcci/description/
// ```
// 设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。这里的“对角线”指的是所有的对角线，不只是平分整个棋盘的那两条对角线。

// 注意：本题相对原题做了扩展

// 示例:

//  输入：4
//  输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//  解释: 4 皇后问题存在如下两个不同的解法。
// [
//  [".Q..",  // 解法 1
//   "...Q",
//   "Q...",
//   "..Q."],

//	["..Q.",  // 解法 2
//	 "Q...",
//	 "...Q",
//	 ".Q.."]
//
// ]
// ```
func solveNQueens(n int) [][]string {
	var ans [][]string
	//棋盘的初始化
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		board[i] = row
	}

	//我们有几个内容需要全局记录,1、每一列是否被使用过；2、主对角线；3、副对角线
	//这里用了一种比较取巧的方式去判断这个对角线是否有数据，
	//同一个对角线上 i+j 的值是相等的
	colUsed := make([]bool, n)
	dig1Used := make([]bool, 2*n)
	dig2Used := make([]bool, 2*n)
	backtrace(0, n, colUsed, dig1Used, dig2Used, board, &ans)
	return ans
}

func backtrace(row int, n int, colUsed []bool, dig1Used []bool, dig2Used []bool, board [][]byte, ans *[][]string) {
	if row == n {
		//最终的结果
		var solution []string
		for _, v := range board {
			solution = append(solution, string(v))
		}
		*ans = append(*ans, solution)
		return
	}

	//判断是否可以放下
	for col := 0; col < n; col++ {
		if !colUsed[col] && !dig1Used[row+col] && !dig2Used[row-col+n-1] {
			board[row][col] = 'Q'
			colUsed[col] = true
			dig1Used[row+col] = true
			dig2Used[row-col+n-1] = true
			backtrace(row+1, n, colUsed, dig1Used, dig2Used, board, ans)
			//复原
			board[row][col] = '.'
			colUsed[col] = false
			dig1Used[row+col] = false
			dig2Used[row-col+n-1] = false
		}
	}
}
