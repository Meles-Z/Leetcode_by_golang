package main

import (
	"fmt"
	"sync"
	"time"
)

func runAllFunc(wg *sync.WaitGroup, fun ...func()) {
	wg.Add(len(fun))
	for _, fn := range fun {
		go func ()  {
			defer wg.Done()
			 fn()
			
		}()
	}
}
func main() {
	var wg sync.WaitGroup

	test1 := func() {
		fmt.Println("Hi there one")
	}

	test2 := func() {
		fmt.Println("Hi there 2")
	}

	func3 := func() {
		fmt.Println("Hi there3")
	}

	runAllFunc(&wg, test1, test2, func3)
	wg.Wait()

	now := time.Now()
	elipse := time.Since(now)
	fmt.Println(elipse)
}