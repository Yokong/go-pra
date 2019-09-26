package main

import "fmt"

const (
	NotFound = -1
)

func main() {
	//a := []int{-1, -2, 3, 4, -1, 1, -3, -4, 5}
	//max := maxSubSequenceSum(a)
	//fmt.Println(max)
	fmt.Println(highPow(2, 40))
}

// 最大子序列
func maxSubSequenceSum(a []int) int {
	thisSum, maxSum := 0, 0
	for j := 0; j < len(a); j++ {
		thisSum += a[j]
		if thisSum > maxSum {
			maxSum = thisSum
		} else if thisSum < 0 {
			thisSum = 0
		}
	}
	return maxSum
}

// 通过二分查找元素下标
func binarySearch(a []int, target int) int {
	low, mid, high := 0, 0, len(a)-1
	for low <= high {
		mid = (low + high) / 2
		if a[mid] < target {
			low = mid + 1
		} else if a[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	return NotFound
}

func highPow(x int64, n uint) int64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n%2 == 0 {
		return highPow(x*x, n/2)
	} else {
		return highPow(x*x, n/2) * x
	}
}
