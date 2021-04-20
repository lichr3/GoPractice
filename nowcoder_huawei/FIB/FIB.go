package main

import "fmt"

func helper(memo map[int]int ,N int) int {
	if N == 1 || N == 2 { return 1 }
	if memo[N] != 0 { return memo[N] }
	memo[N] = helper(memo, N-1) + helper(memo, N-2)
	return memo[N]
}

func fib(N int) int{
	if N < 1 { return 0 }
	memo := make(map[int]int)
	return helper(memo, N)
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(fib(N))
}