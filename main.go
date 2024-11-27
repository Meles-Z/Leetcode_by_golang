package main

import "fmt"

type CircleQueue struct {
	data  []int
	size  int
	front int
	rare  int
}

func NewCircularLinkedList(size int) *CircleQueue {
	return &CircleQueue{
		data:  make([]int, size),
		size:  size,
		front: -1,
		rare:  -1,
	}
}
func (q *CircleQueue) isFull() bool {
	return (q.rare+1)%q.size == q.front
}

func main() {
	fmt.Print("Hello circluar linked list")
}
