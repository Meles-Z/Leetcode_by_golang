package main

import "fmt"

func reverse(n int) int {
	rev := 0
	for n != 0 {
		a := n % 10
		rev = rev*10 + a
		n = n / 10
	}
	return rev
}
func main() {
	fmt.Println(reverse(456))

}