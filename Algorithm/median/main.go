package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merge := []int{}
	merge1 := []int{}
	merge2 := []int{}
	var median float64
	for i := 0; i < len(nums1); i++ {
		merge1 = append(merge1, nums1[i])
	}
	for i := 0; i < len(nums2); i++ {
		merge2 = append(merge2, nums2[i])
	}
	merge = append(merge1, merge2...)
	for i := 0; i < len(merge)-1; i++ {
		for j := 0; j < len(merge)-i-1; j++ {
			if merge[j] > merge[j+1] {
				merge[j], merge[j+1] = merge[j+1], merge[j]
			}
		}
	}
	fmt.Println(merge)
	middle := len(merge) / 2
	if len(merge)%2 == 0 {
		median = float64(merge[middle]+merge[middle-1]) / 2.0
		return median
	} else {
		median = float64(merge[middle])
	}
	return median
}
func main() {
	n1 := []int{1,3}
	n2 := []int{2}
	fmt.Println(findMedianSortedArrays(n1, n2))
}
