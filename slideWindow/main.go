package main

import "fmt"

func maximumSubArray(nums []int, k int) int {
	//[2,3,4,4,5]
	var windowStart, windowEnd, maxSum, windowSum int
	windowStart = 0

	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	maxSum = windowSum
	for windowEnd = k; windowEnd < len(nums); windowEnd++ {
		windowSum = windowSum + nums[windowEnd] - nums[windowStart]
		if windowSum > maxSum {
			maxSum = windowSum
		}
		windowStart++
	}
	return maxSum
}

func main() {
	nums := []int{1, 5, 4, 8, 7, 1, 9}
    k := 3
    fmt.Println(maximumSubArray(nums, k))
}
