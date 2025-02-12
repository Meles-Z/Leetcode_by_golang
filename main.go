// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func runAllFunc(wg *sync.WaitGroup, fun ...func()) {
// 	wg.Add(len(fun))
// 	for _, fn := range fun {
// 		go func ()  {
// 			defer wg.Done()
// 			 fn()

// 		}()
// 	}
// }
// func main() {
// 	var wg sync.WaitGroup

// 	test1 := func() {
// 		fmt.Println("Hi there one")
// 	}

// 	test2 := func() {
// 		fmt.Println("Hi there 2")
// 	}

// 	func3 := func() {
// 		fmt.Println("Hi there3")
// 	}

// 	runAllFunc(&wg, test1, test2, func3)
// 	wg.Wait()

// 	now := time.Now()
// 	elipse := time.Since(now)
// 	fmt.Println(elipse)
// }

// package main

// import "fmt"

// func main() {
// we must have two go routines one for send and one for receive
// both of them are gourutines
// if it's one goroutines it will hapen deadlock
// main function is default goroutines

// ch := make(chan int)
// // let me dive into unbuffered channel

// go func() {
// 	ch <- 5
// }()
// val := <-ch
// fmt.Println(val)

// 	ch:=make(chan int, 3)
// 	ch<-5
// 	ch<-6
// 	ch<-7
// 	val1:=<-ch
// 	val2:=<-ch
// 	val3:=<-ch
// 	go func ()  {
// 		ch<-8
// 	}()
// 	val4:=<-ch
// 	fmt.Println(val1,val2, val3, val4)
// }

package main

import (
	"fmt"
	"sync"
)

func evenNumber() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Even:", i)
	}
}

func oddNumber() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Odd:", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		evenNumber()
	}()

	go func() {
		defer wg.Done()
		oddNumber()
	}()
	wg.Wait()
}
