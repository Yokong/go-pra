package main

import (
	"fmt"
)

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

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

// 2. 打家劫舍
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = intMax(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = intMax(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

// 3. 最大子段和
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max_res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = intMax(dp[i-1]+nums[i], nums[i])
		if dp[i] > max_res {
			max_res = dp[i]
		}
	}
	return max_res
}

// 4. 找零钱
func coinChange(coins []int, amount int) int {
	var dp []int
	for i := 0; i <= amount; i++ {
		dp = append(dp, -1)
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] != -1 {
				if dp[i] == -1 || dp[i] > dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}

	return dp[amount]
}

func main() {
	a := []int{1, 2, 5, 7, 10}
	fmt.Println(coinChange(a, 19))
}
