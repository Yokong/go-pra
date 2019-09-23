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

func intMin(x, y int) int {
	if x < y {
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

func fib(n int) int {
	a, b := 1, 2
	if n == 0 {
		return 0
	}
	if n == 1 {
		return a
	}
	if n == 2 {
		return b
	}
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
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
	maxRes := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = intMax(dp[i-1]+nums[i], nums[i])
		if dp[i] > maxRes {
			maxRes = dp[i]
		}
	}
	return maxRes
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

// 5. 三角形
func minimumTotal(triangle [][]int) int {
	minNums := make([]int, len(triangle[len(triangle)-1]))
	for i := 0; i < len(triangle[len(triangle)-1]); i++ {
		minNums[i] = triangle[len(triangle)-1][i]
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			minNums[j] = intMin(minNums[j], minNums[j+1]) + triangle[i][j]
		}
	}
	return minNums[0]
}

// 6. 最长上升子序列
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	lis := 1
	for i := 1; i < len(dp); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if lis < dp[i] {
			lis = dp[i]
		}
	}
	return lis
}

// 最小路径和
func minPathSum(grid [][]int) int {
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] += intMin(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {
	a := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	b := minPathSum(a)
	fmt.Println(b)
}
