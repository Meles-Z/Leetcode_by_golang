package main

import "fmt"

func returnMatrix(rows, col int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, col)
	}
	return matrix
}
func main() {

	matrix := returnMatrix(3, 4)
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			matrix[i][j] = i*3 + j + 1
		}
	}
	for _, row:=range matrix{
		fmt.Println( row)
	}
}
