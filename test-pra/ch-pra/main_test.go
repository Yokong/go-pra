package main

import (
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(45)
	}
}

func BenchmarkFib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib2(45)
	}
}
