package main

import "fmt"

func main() {
	fmt.Println("Hello there!")
	a := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			fmt.Println(a[i][j])
		}
	}
}
