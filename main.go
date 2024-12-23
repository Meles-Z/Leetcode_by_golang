package main

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = original[i*n : (i+1)*n]
	}
	return result
}
func main() {
	fmt.Println(construct2DArray([]int{1, 2, 3, 4}, 2, 2))
}
