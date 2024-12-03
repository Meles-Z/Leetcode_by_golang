package main

import "fmt"

func reverse(x int) int {
	num := 0
	strInt := fmt.Sprintf("%d", x)
	for i := len(strInt) - 1; i >= 0; i-- {
		num = int(strInt)
		
	}
	return num
}
func main() {
	fmt.Println(reverse(123))

}
