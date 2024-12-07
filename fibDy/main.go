package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(otherFibonacciTech(50))
	elipse := time.Since(start)
	fmt.Println(elipse)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibDynamic(n int) int {
	a := []int{1}
	if n == 1 {
		return 1
	}
	a = append(a, 1)
	if n == 2 {
		return 2
	}
	for i := 2; i < n; i++ {
		a = append(a, a[i-1]+a[i-2])
	}
	return a[len(a)-1]
}

// Memoization
var memo = make(map[int]int)

func memoization(n int) int {
	_, ok := memo[n]

	if ok {
		return memo[n]

	}
	if n <= 2 {
		return 1
	}
	memo[n] = memoization(n-1) + memoization(n-2)
	return memo[n]
}

func tabulation(n int) int {
	var result []int = make([]int, n+2)

	result[0] = 0
	result[1] = 1

	for i := 0; i <= n-1; i++ {

		result[i+1] += result[i]

		result[i+2] += result[i]

	}

	return result[n]
}

func otherFibonacciTech(n int) int {
	if n < 2 {
		return n
	}
	prev1, prev2 := 1, 0
	for i := 2; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}
	return prev1
}
