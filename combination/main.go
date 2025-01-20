package main

import "fmt"

// Function to find all unique combinations
func combinationSum(candidates []int, target int) [][]int {
    var result [][]int // Stores all valid combinations
    var current []int  // Stores the current combination being explored

    // Helper function for backtracking
    var backtrack func(index int, target int)
    backtrack = func(index int, target int) {
        // Base case: when the target becomes 0, a valid combination is found
        if target == 0 {
            combination := append([]int{}, current...) // Copy the current combination
            result = append(result, combination)
            return
        }

        // Explore all candidates starting from the current index
        for i := index; i < len(candidates); i++ {
            if candidates[i] <= target { // Only proceed if the candidate can fit into the target
                current = append(current, candidates[i]) // Add candidate to the current combination
                backtrack(i, target-candidates[i])       // Recur with reduced target
                current = current[:len(current)-1]       // Backtrack by removing the last candidate
            }
        }
    }

    // Start the backtracking process
    backtrack(0, target)
    return result
}

func main() {
    fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // Output: [[2, 2, 3], [7]]
    fmt.Println(combinationSum([]int{2, 3, 5}, 8))    // Output: [[2, 2, 2, 2], [2, 3, 3], [3, 5]]
    fmt.Println(combinationSum([]int{2}, 1))          // Output: []
}
