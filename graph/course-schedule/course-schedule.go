package leetcode

//https://leetcode.cn/problems/course-schedule/description/
//课程表
//标准的拓扑排序

func canFinish(numCourses int, prerequisites [][]int) bool {
	//采用dfs的方式
	valid := true
	//存储每个节点是否访问过，0表示没有访问过，1表示访问过，2表示已经访问过
	visited := make([]int, numCourses)
	//采用邻接链表的方式存储
	edges := make([][]int, numCourses)
	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
	}
	//构造邻接链表
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}
