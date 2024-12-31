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

func main() {
	fmt.Println(permutaionOfString("CAT"))
}
