package main

import "fmt"

func findMissingNumberInArray(nums []int) []int {
	missed := []int{}
	if len(nums) == 0 {
		return []int{}
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := nums[i] + 1; j < nums[i+1]; j++ {
			missed = append(missed, j)
		}
	}
	return missed
}

func main() {
	fmt.Println(findMissingNumberInArray([]int{1, 3, 4, 6}))
}
