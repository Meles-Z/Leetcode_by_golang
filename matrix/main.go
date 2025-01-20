package main

import "fmt"

func Identity(order int) [][]float64 {
	matrix := make([][]float64, order)
	for i := 0; i < order; i++ {
		row := make([]float64, order)
		matrix[i] = row
		row[i] = 1
	}
	return matrix
}

func AddMatrix(mat1 [2][2]int, mat2 [2][2]int) [2][2]int {
	var sum [2][2]int
	for l := 0; l < 2; l++ {
		for j := 0; j < 2; j++ {
			sum[l][j] = mat1[l][j] + mat2[l][j]
		}
	}
	return sum
}

func SubstructMatrix(mat1 [2][2]int, mat2 [2][2]int) [2][2]int {
	var sum [2][2]int
	for l := 0; l < 2; l++ {
		for j := 0; j < 2; j++ {
			sum[l][j] = mat1[l][j] - mat2[l][j]
		}
	}
	return sum
}

func transpose(matrix1 [2][2]int) [2][2]int {
	var m int
	var l int
	var transMatrix [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			transMatrix[l][m] = matrix1[m][l]
		}
	}
	return transMatrix
}

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
	fmt.Println(Identity(4))
	fmt.Println(AddMatrix([2][2]int{{1, 2}, {1, 2}}, [2][2]int{{3, 4}, {5, 6}}))
	fmt.Println(SubstructMatrix([2][2]int{{6, 1}, {1, 2}}, [2][2]int{{3, 4}, {5, 6}}))
	fmt.Println(transpose([2][2]int{{6, 1}, {1, 2}}))
}
