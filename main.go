package main

import "fmt"

func missingLessNumber(nums []int) int {
	greater := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			greater = nums[i]
		}
	}
	return greater
}

func main() {
	gr := []int{3, 5, 13, 5}
	fmt.Println(missingLessNumber(gr))
}
