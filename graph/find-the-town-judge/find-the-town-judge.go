package leetcode

//https://leetcode.cn/problems/find-the-town-judge/description/
//找到小镇的法官

func findJudge(n int, trust [][]int) int {
	//其实只需要考虑入度和出度，不需要考虑邻接链表或者矩阵的存储
	//最终满足条件的应该是入度为N-1，出度为0
	indegree := make(map[int]int)
	outdegree := make(map[int]int)
	for _, item := range trust {
		indegree[item[1]]++
		outdegree[item[0]]++
	}
	for i := 1; i <= n; i++ {
		if indegree[i] == n-1 && outdegree[i] == 0 {
			return i
		}
	}
	return -1
}
