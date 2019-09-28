package main

import (
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(5)
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		climbStairs(5)
	}
}

func BenchmarkLengthOfLIS(b *testing.B) {
	a := []int{1, 2, 3, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLIS(a)
	}
}

func BenchmarkDungeon(b *testing.B) {
	a := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateMinimumHP(a)
	}
}
