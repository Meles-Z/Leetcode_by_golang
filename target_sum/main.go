package main

import "fmt"

func twoSum(num []int, target int) []int {
	for i := 0; i < len(num)-1; i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i]+num[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
func main() {
	num := []int{1, 3, 4, 5}
	fmt.Println(twoSum(num, 8))

}
