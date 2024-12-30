package main

import "fmt"

func isHappy(n int) bool {
	seen := make(map[int]bool)
	for n != 1 {
		if seen[n] {
			return false
		}
		seen[n] = true
		n=somofSquareNumber(n)
	}
	return true
}
func somofSquareNumber(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
func main() {
	fmt.Println(isHappy(2))
}
