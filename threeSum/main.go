package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				// Check if the triplet sums to 0
				if nums[i]+nums[j]+nums[k] == 0 {
					triplet := []int{nums[i], nums[j], nums[k]}

					// Avoid duplicates: check if triplet already exists in result
					if !contains(result, triplet) {
						result = append(result, triplet)
					}
				}
			}
		}
	}

	return result
}

// Helper function to check for duplicate triplets
func contains(result [][]int, triplet []int) bool {
	for _, res := range result {
		matchCount := 0
		for _, num := range res {
			for _, t := range triplet {
				if num == t {
					matchCount++
				}
			}
		}
		if matchCount == 3 {
			return true
		}
	}
	return false
}


func threeSumm(nums []int) [][]int {
    var result [][]int

    // Step 1: Sort the array
    sort.Ints(nums)

    // Step 2: Iterate over the array
    for i := 0; i < len(nums)-2; i++ {
        // Skip duplicates for the current index `i`
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        // Step 3: Use two-pointer technique
        left, right := i+1, len(nums)-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                // Found a triplet
                result = append(result, []int{nums[i], nums[left], nums[right]})
                
                // Skip duplicates for `left` and `right`
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                for left < right && nums[right] == nums[right-1] {
                    right--
                }

                left++
                right--
            } else if sum < 0 {
                left++ // Move `left` pointer to increase the sum
            } else {
                right-- // Move `right` pointer to decrease the sum
            }
        }
    }

    return result
}


func main() {
	t := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(t))
}
