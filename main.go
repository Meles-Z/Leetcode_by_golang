package main

import (
	"fmt"
	"time"
)

func OtherMethod() {
	start := time.Now() // Record the start time

	for i := 0; i <= 10; i++ {
		square := i * i
		fmt.Println(square)
	}

	elapsed := time.Since(start) 
	fmt.Println("Execution time for OtherMethod:", elapsed)
}

func main() {
	OtherMethod()
	start:=time.Now()
	nums := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
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
	elspsed:=time.Since(start)
	fmt.Println("Time to execution", elspsed)
}
