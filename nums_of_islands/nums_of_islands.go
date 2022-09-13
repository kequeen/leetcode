package leetcode

import (
	"container/list"
)

//https://leetcode.cn/problems/number-of-islands/
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。

//其实我理解就是遍历，看看需要遍历几次才能够遍历完
//目前这个解法会超时，还是需要再多想想
func numIslands(grid [][]byte) int {
	num := 0
	if grid == nil {
		return num
	}
	//开始遍历，每次都是DFS遍历
	//数组存储是否遍历过
	row := len(grid)
	column := len(grid[0])
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, column)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			//如果已经遍历过
			if visited[i][j] {
				continue
			}

			//如果是水
			if grid[i][j] == '0' {
				visited[i][j] = true
				continue
			}
			l := list.New()
			l.PushBack(land{i, j})
			num++
			for l.Len() != 0 {
				node := l.Front()
				//上下左右都需要去判断
				currnetRow := node.Value.(land).row
				currnetColumn := node.Value.(land).column
				//标记为遍历过
				visited[currnetRow][currnetColumn] = true

				//上面
				if currnetRow-1 >= 0 {
					if !visited[currnetRow-1][currnetColumn] && grid[currnetRow-1][currnetColumn] == '1' {
						l.PushBack(land{currnetRow - 1, currnetColumn})
					}
				}
				//下面
				if currnetRow+1 < row {
					if !visited[currnetRow+1][currnetColumn] && grid[currnetRow+1][currnetColumn] == '1' {
						l.PushBack(land{currnetRow + 1, currnetColumn})
					}
				}
				//左边
				if currnetColumn-1 >= 0 {
					if !visited[currnetRow][currnetColumn-1] && grid[currnetRow][currnetColumn-1] == '1' {
						l.PushBack(land{currnetRow, currnetColumn - 1})
					}
				}
				//右边
				if currnetColumn+1 < column {
					if !visited[currnetRow][currnetColumn+1] && grid[currnetRow][currnetColumn+1] == '1' {
						l.PushBack(land{currnetRow, currnetColumn + 1})
					}
				}
				l.Remove(node)
			}
		}
	}

	return num
}

//官方的解法跟我的解法差不多，只不过没有用map存储已经访问过的陆地节点，而是直接标记为水，试试看这种解法是否会超时
//感觉不是因为这个原因超时的，可能是因为go里面用的list来模拟队列这种实在太重了
func numIslandsV2(grid [][]byte) int {
	num := 0
	if grid == nil {
		return num
	}
	//开始遍历，每次都是BFS遍历
	row := len(grid)
	column := len(grid[0])

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {

			//如果是水
			if grid[i][j] == '0' {
				continue
			}
			l := list.New()
			l.PushBack(land{i, j})
			num++
			for l.Len() != 0 {
				node := l.Front()
				//上下左右都需要去判断
				currnetRow := node.Value.(land).row
				currnetColumn := node.Value.(land).column
				//标记为遍历过
				grid[currnetRow][currnetColumn] = '0'

				//上面
				if currnetRow-1 >= 0 {
					if grid[currnetRow-1][currnetColumn] == '1' {
						l.PushBack(land{currnetRow - 1, currnetColumn})
					}
				}
				//下面
				if currnetRow+1 < row {
					if grid[currnetRow+1][currnetColumn] == '1' {
						l.PushBack(land{currnetRow + 1, currnetColumn})
					}
				}
				//左边
				if currnetColumn-1 >= 0 {
					if grid[currnetRow][currnetColumn-1] == '1' {
						l.PushBack(land{currnetRow, currnetColumn - 1})
					}
				}
				//右边
				if currnetColumn+1 < column {
					if grid[currnetRow][currnetColumn+1] == '1' {
						l.PushBack(land{currnetRow, currnetColumn + 1})
					}
				}
				l.Remove(node)
			}
		}
	}

	return num
}

//尝试采用DFS遍历
//DFS这种方式确实不会超时，不过我觉得可能是调用go的list创建代价太高了
func numIslandsV3(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	nums := 0
	//开始遍历，每次都是BFS遍历
	row := len(grid)
	column := len(grid[0])
	var dfs func(i int, j int)
	dfs = func(i, j int) {
		if grid[i][j] == '0' {
			//如果已经是水，则结束
			return
		}
		grid[i][j] = '0'
		//遍历上下左右
		if i-1 >= 0 && grid[i-1][j] == '1' {
			dfs(i-1, j)
		}
		if i+1 < row && grid[i+1][j] == '1' {
			dfs(i+1, j)
		}
		if j-1 >= 0 && grid[i][j-1] == '1' {
			dfs(i, j-1)
		}
		if j+1 < column && grid[i][j+1] == '1' {
			dfs(i, j+1)
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == '1' {
				nums++
				dfs(i, j)
			}
		}
	}
	return nums
}

type land struct {
	row    int
	column int
}
