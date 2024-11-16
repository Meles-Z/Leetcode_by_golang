package main

import "fmt"

func main() {
	nums := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			nums <- i
		}
		close(nums)
	}()
	go func() {
		for num := range nums {
			squares <- num * num
		}
		close(squares)
	}()
	for square := range squares {
		fmt.Println(square)
	}
}
