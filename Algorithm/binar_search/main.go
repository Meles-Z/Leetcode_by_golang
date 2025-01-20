package main

import (
	"fmt"
	"sort"
)

func binarySearch(num []int, target int) int {
	start := 0
	end := len(num) - 1

	for start <= end {
		middle := (start + end) / 2
		if num[middle] == target {
			return middle
		} else if num[middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}
func main() {
	data := []int{4, 5, 1, 9, 6, 10}
	sort.Ints(data)
	fmt.Println(binarySearch(data, 5))
}
