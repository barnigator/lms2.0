package main

import (
	"errors"
	"time"
)

func Fib(n int) int {
	dp := make([]int, n+2)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func countFib(n int, out chan<- int) {
	num := Fib(n)
	out <- num
	close(out)
}

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	chanFib := make(chan int)
	if n < 0 {
		return 0, errors.New("n must be non-negative")
	}
	go countFib(n, chanFib)

	// return <-chanFib, nil

	select {
	case num := <-chanFib:
		return num, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}
