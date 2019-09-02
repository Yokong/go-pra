package main

import "fmt"

// 1. 爬楼梯
func climbStairs(n int) int {
	dp := make([]int, n+3)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	r := climbStairs(3)
	fmt.Println(r)
}
