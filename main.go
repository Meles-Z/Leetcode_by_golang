package main

import "fmt"

func slideWindow(nums []int, target int) []int {
	if len(nums) < 3 {
		return []int{}
	}

	sum := nums[0] + nums[1] + nums[2]
	for i := 3; i < len(nums); i++ {
		if sum == target {
			return []int{nums[i-3], nums[i-2], nums[i-1]}
		}
		// start slide window
		sum = sum - nums[i-3] + nums[i]
	}
	if sum == target {
		return []int{nums[len(nums)-3], nums[len(nums)-2], nums[len(nums)-1]}
	}
	return []int{}
}

func targetSum(nums []int, target int) []int {
	k := 3 // Size of the sliding window
	windowSum := 0
	// Initialize the first window sum
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	for i := k; i < len(nums); i++ {
		// If the window sum matches the target, return the window
		if windowSum == target {
			return nums[i-k : i] // Return the window as a slice
		}

		// Slide the window:
		// Remove the first number of the previous window
		windowSum -= nums[i-k]

		// Add the next number to the window
		windowSum += nums[i]
	}

	// Final check for the last window
	if windowSum == target {
		return nums[len(nums)-k:]
	}

	return nil // Return nil if no such window is found
}

func main() {
	fmt.Println(targetSum([]int{1, 2, 3, 4, 5}, 12))
}
