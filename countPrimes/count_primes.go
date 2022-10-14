package leetcode

// 计算质数
//https://leetcode.cn/leetbook/read/top-interview-questions/x20yuc/

//以前是不是想过筛法，筛法真的是很有意思的方法,数学证明上比较有意思，但实际写下来貌似没那么容易写下来

//还是用了比较常规的方法，但是这个方法超时了
func countPrimes(n int) int {
	//数组存储下当前的质数
	primeArr := []int{}
	for i := 2; i < n; i++ {
		if len(primeArr) == 0 {
			primeArr = append(primeArr, i)
		}
		primeFlag := true
		for _, value := range primeArr {
			if i%value == 0 {
				primeFlag = false
				break
			}
		}
		if primeFlag {
			primeArr = append(primeArr, i)
		}
	}
	return len(primeArr)
}

//尝试使用筛法
func countPrimesV2(n int) int {
	//用数组存储下是否为质数
	isPrime := make([]bool, n)
	for i, _ := range isPrime {
		isPrime[i] = true
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
			//这个筛法确实比较精妙
			for j := 2 * i; j < n; j = j + i {
				isPrime[j] = false
			}
		}

	}
	return count
}
