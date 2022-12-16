package leetcode

//有一条河流长度为N，其中某些位置上有可以停留的石子格，假设人上一次跨过的距离为X，则当次可以选择的步长为X-1，X，X+1，第一次人的步长默认为1
//给你一个格子的列表，第一个数默认为0，最后一个数代表对岸的种点，判断能否过河

func canCrossRiver(arr []int) bool {
	if len(arr) == 0 || len(arr) == 1 {
		return true
	}
	//采用贪心算法，每一步，都应该去走它所能到达的下一步的最大值。当前位置 -> 下一步可能的最大值之间，如果有其他的步数，都可以忽略

	//记录当前位置
	current := arr[0]
	//记录当前最大步数
	maxStep := 1
	//记录上次的值
	lastValue := 0
	for i := 1; i < len(arr); {
		//如果下一步不满足条件
		if arr[i]-current > maxStep+1 {
			return false
		}
		//小于的maxStep - 1 的话，直接往前走
		if arr[i]-current < maxStep-1 {
			i++
		}

		//判断是否需要更新current 和 maxStep
		for i < len(arr) && arr[i]-current <= maxStep+1 {
			lastValue = arr[i]
			i++
		}
		maxStep = lastValue - current
		current = lastValue
	}
	return true
}
