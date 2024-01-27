package leetcode

// https://leetcode.cn/problems/combination-sum/?favorite=2cktkvj
// 组合总数
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	comb := []int{}
	//其实每次搜索都是一个分叉，要不选择这个分叉，要不不选择另外一个分叉，就像人生一样，只不过人生很难回溯
	var dfs func(int, int)
	dfs = func(n, idx int) {
		//判定遍历是否终止
		if idx == len(candidates) {
			return
		}
		if n == 0 {
			//这种方式会导致结果不符合预期，因为golang本身的数组切片传的是引用,这种方式会导致后续数组切片的内容发生了变更，这个结果也会被变更
			//ans = append(ans, comb)
			//要采用这种直接复制数组的方式
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		//分叉,不选这个数
		dfs(n, idx+1)
		//选择这个数
		if n-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(n-candidates[idx], idx)
			//恢复原装
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}
