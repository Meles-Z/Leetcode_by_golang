package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return n
	}
	return n * factorial(n-1)

}

func ffactorial(n int) int {
	var value int = 1
	for i := 1; i <= n; i++ {
		value *= i
	}
	return value
}

func NCR(n int, r int) int {
	res := factorial(n) / (factorial(n-r) * factorial(r))
	return res
}
func main() {
	fmt.Println(factorial(5))

}
