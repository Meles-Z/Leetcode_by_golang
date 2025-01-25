package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func squareOfNums(nums []int, result chan<- int) {
	defer wg.Done()
	for _, num := range nums {
		result <- num * num
	}
}

func main() {
	val := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(val))
	wg.Add(1)
	go squareOfNums(val, ch)
	wg.Wait()
	close(ch)
	var square []int
	for res := range ch {
		square = append(square, res)
	}
	fmt.Println("Squares:", square)
}
