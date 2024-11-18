package main

import "fmt"

func isMonotonic(num []int) bool {
	if len(num) <= 1 {
		return true
	}
	isDeacreasing := true
	isIncreasing := true
	for i := 0; i < len(num)-1; i++ {
		if num[i] > num[i+1] {
			isIncreasing = false
			fmt.Println("deacrease")
		} else if num[i] < num[i+1] {
			fmt.Println("increase")
			isDeacreasing = false
		} else {
			fmt.Println("Neither increase nor deacrease")
		}
	}

	return isDeacreasing || isIncreasing
}
func main() {
	num := []int{4, 4, 3, 6}
	fmt.Println(isMonotonic(num))
}
