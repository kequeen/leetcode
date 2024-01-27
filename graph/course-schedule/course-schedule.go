package leetcode

// https://leetcode.cn/problems/course-schedule/description/
// 课程表
// 标准的拓扑排序
// 重新回看，意识到这个算法不是最优的，因为有可能已经计算出了有环，但是还是会继续算完
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

func canFinishV2(numCourses int, prerequisites [][]int) bool {
	//将数据转换为邻接链表
	link := make(map[int][]int)
	for _, item := range prerequisites {
		link[item[0]] = append(link[item[0]], item[1])
	}
	visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		//加一个条件，时间就能用优化很多，说明很多时候，函数调用的成本还是比较大的
		if visited[i] == 0 && hasCycle(link, i, visited) {
			return false
		}
	}
	return true
}

// 用邻接链表的方式进行存储
// 0 是未访问，1是正在访问，2是已访问
func hasCycle(link map[int][]int, visit int, visited []int) bool {
	if visited[visit] == 1 {
		return true
	}
	if visited[visit] == 2 {
		return false
	}
	//正在访问中
	visited[visit] = 1
	currentVisitArr := link[visit]
	for _, i := range currentVisitArr {
		if hasCycle(link, i, visited) {
			return true
		}
	}
	visited[visit] = 2
	return false
}
