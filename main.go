package main

import "fmt"

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func exist(nums [][]int, target []int) bool {
	for _, val := range nums {
		if compare(val, target) {
			return true
		}
	}
	return false
}
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	var result [][]int
	for i, num := range nums {
		res := append([]int{}, nums[:i]...)
		res = append(res, nums[i+1:]...)
		for _, perm := range permute(res) {
			newPerm := append([]int{num}, perm...)
			if !exist(result, newPerm) {
				result = append(result, newPerm)
			}
		}
	}

	return result
}

func main() {
	fmt.Println(permute([]int{1, 1, 2}))
}
