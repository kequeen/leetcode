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

type land struct {
	row    int
	column int
}
