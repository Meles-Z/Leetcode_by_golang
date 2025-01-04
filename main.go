package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func makeArray(length int) []int {

	array := []int{}
	for i := 0; i < length; i++ {
		array = append(array, rand.Intn(10))
	}
	return array
}


var wg sync.WaitGroup

func main() {
	ch := make(chan []int, 10)
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			makeArray := makeArray(10)
			ch<- makeArray
		}()
	}
	wg.Wait()
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
