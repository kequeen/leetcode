package leetcode

//https://leetcode.cn/problems/word-search/description/?favorite=2cktkvj
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

// 单次搜索，也是很经典的题目
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	wordLen := len(word) - 1
	m := len(board)
	n := len(board[0])
	//上下左右四个方向
	direction := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	//还需要记录每个节点的访问状态
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}

	//check函数
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		//剪枝
		if board[i][j] != word[k] {
			return false
		}
		//如果已经遍历完
		if k == wordLen {
			return true
		}
		visit[i][j] = true
		//defer函数还是很方便的
		defer func() {
			//回溯还原
			visit[i][j] = false
		}()
		//未遍历完，则继续循环,向上下左右四个方向
		for _, item := range direction {
			if i+item[0] >= 0 && i+item[0] < m && j+item[1] >= 0 && j+item[1] < n && !visit[i+item[0]][j+item[1]] {
				//继续遍历
				// 如果这种写法，会导致有不合适的就会被错误剪枝
				// return check(i+item[0], j+item[1], k+1)

				if check(i+item[0], j+item[1], k+1) {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//在check函数里面做剪枝
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
