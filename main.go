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
    for l:=0; l<2; l++{
        for j:=0; j<2; j++{
            sum[l][j]=mat1[l][j]+mat2[l][j]
        }
    }
    return sum
}
func main() {
	fmt.Println(Identity(4))
    fmt.Println(AddMatrix([2][2]int{{1,2},{1,2}}, [2][2]int{{3,4},{5,6}}))

}
