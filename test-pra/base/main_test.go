package main

import (
	"testing"
)

func BenchmarkReverseStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseStr("Yoko")
	}
}

func BenchmarkReverseStrV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseStrV2("Yoko")
	}
}
