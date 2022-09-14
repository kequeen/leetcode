package leetcode

//https://leetcode.cn/problems/island-perimeter/
//求岛屿的周长，在我看来也是遍历的问题
//这里采用DFS遍历
func islandPerimeter(grid [][]int) int {
	permeter := 0
	if grid == nil {
		return permeter
	}
	row := len(grid)
	column := len(grid[0])
	//因为不能改变已经访问过的节点，所以用个二维数组存储下访问过的节点
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, column)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if visited[i][j] || grid[i][j] == 0 {
			return
		}
		//标记访问过
		visited[i][j] = true
		//初始化长度为4
		initLen := 4
		//上下左右是陆地就-1
		if i-1 >= 0 {
			if grid[i-1][j] == 1 {
				initLen--
				if visited[i-1][j] {
					dfs(i-1, j)
				}
			}
		}

		if i+1 < row {
			if grid[i+1][j] == 1 {
				initLen--
				if visited[i+1][j] {
					dfs(i+1, j)
				}
			}
		}

		if j-1 >= 0 {
			if grid[i][j-1] == 1 {
				initLen--
				if visited[i][j-1] {
					dfs(i, j-1)
				}
			}
		}
		if j+1 < column {
			if grid[i][j+1] == 1 {
				initLen--
				if visited[i][j+1] {
					dfs(i, j+1)
				}
			}
		}

		permeter = permeter + initLen
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				dfs(i, j)
			}
		}
	}

	return permeter
}
