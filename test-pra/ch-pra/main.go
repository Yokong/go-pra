package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func fib2(x int) int {
	if x < 2 {
		return x
	}
	a, b := 1, 2
	for i := 3; i < x; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	start := time.Now()
	fibN := fib(n)
	end := time.Since(start)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	fmt.Println(end)
}
