package main

import "fmt"

func addTwoSlices(nums1 []int, nums2 []int) []int {
	if len(nums1) != len(nums2) {
		fmt.Println("Error: slice have the same length")
		return nil
	}

	sum := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		sum[i] = nums1[i] + nums2[i]
	}
	return sum

}

func main() {
	nm1 := []int{3, 4, 5}
	nm2 := []int{5, 6, 7}
	fmt.Println(addTwoSlices(nm1, nm2))

}
