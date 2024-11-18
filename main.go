package main

import "fmt"

func main() {
	num := [4]int{2, 3, 4, 5}
	greater := num[0]
	for i := 0; i < len(num); i++ {
		if num[i] > greater {
			greater = num[i]
		}
	}
	fmt.Println(greater)
}
