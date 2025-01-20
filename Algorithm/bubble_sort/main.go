package main

import "fmt"

func bubbleSort(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
	return nums
}
func main() {
	nums := []int{1, 6, 2, 5}
	fmt.Println(bubbleSort(nums))

}
