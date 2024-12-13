package main

import (
	"fmt"
	"sort"
)

func missedArray(nums []int) []int {
	sort.Ints(nums)
	missed := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := nums[i] + 1; j < nums[i+1]; j++ {
			missed = append(missed, j)
		}
	}
	return missed
}

func main() {
	fmt.Println(missedArray([]int{9, 6, 2, 3, 5, 7, 0}))

}
