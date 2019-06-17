package main

import "fmt"

func main() {
	a := []int{-1, -2, 3, 4, -1, 1, -3, -4, 5}
	max := maxSubSequenceSum(a)
	fmt.Println(max)
}

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
