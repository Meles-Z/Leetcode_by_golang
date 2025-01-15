package main

import "fmt"

func returnTwoDimenstion(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	return matrix
}

func threeDimensions(rows, cols int) [][][]int {
	matrix := make([][][]int, rows)
	for i:=range matrix{
		matrix[i]=make([][]int, rows)
		for j:=range matrix[i]{
			matrix[i][j]=make([]int, cols)
		}
	}
	return matrix
}
func main() {
	row := 3
	cols := 4
	matrix := returnTwoDimenstion(row, cols)
	for i := 0; i < row; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = i*cols + j + 1
		}
	}
	for i := 0; i < len(matrix); i++ {

		fmt.Println(matrix[i])
	}
  matr:=threeDimensions(row, cols)
  fmt.Println("three Dimesnsiton", matr)
}
