package main

import "fmt"

func permutaionOfString(s string) []string {
	if len(s) == 0 {
		return []string{""}
	}
	var result []string
	for i, char := range s {
		remaining := s[:i] + s[i+1:]
		for _, perm := range permutaionOfString(remaining) {
			result = append(result, string(char)+perm)
		}
	}
	return result
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	var result [][]int
	for i, num := range nums {
		rest := append([]int{}, nums[:i]...)
		rest = append(rest, nums[i+1:]...)
		for _, perm := range permute(rest) {
			result = append(result, append([]int{num}, perm...))
		}
	}

	return result
}

// remove reptitive
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
func permut(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	var result [][]int
	for i, num := range nums {
		res := append([]int{}, nums[:i]...)
		res = append(res, nums[i+1:]...)
		for _, perm := range permut(res) {
			newPerm := append([]int{num}, perm...)
			if !exist(result, newPerm) {
				result = append(result, newPerm)
			}
		}
	}

	return result
}

func main() {
	fmt.Println(permutaionOfString("CAT"))
}
