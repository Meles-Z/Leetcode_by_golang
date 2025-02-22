package main

import "fmt"

type CircularArray struct {
	data  []int
	head  int
	tail  int
	size  int
	cap   int
}

func NewCircularArray(capacity int) *CircularArray {
	return &CircularArray{
		data: make([]int, capacity),
		cap:  capacity,
	}
}

func (ca *CircularArray) Enqueue(value int) error {
	if ca.size == ca.cap {
		return fmt.Errorf("queue is full")
	}
	ca.data[ca.tail] = value
	ca.tail = (ca.tail + 1) % ca.cap
	ca.size++
	return nil
}


func (ca *CircularArray) Dequeue() (int, error) {
	if ca.size == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	value := ca.data[ca.head]
	ca.head = (ca.head + 1) % ca.cap
	ca.size--
	return value, nil
}

func main() {
	ca := NewCircularArray(5)
	ca.Enqueue(10)
	ca.Enqueue(20)
	fmt.Println(ca.Dequeue()) // Outputs: 10
	fmt.Println(ca.Dequeue()) // Outputs: 20
}


package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type CircularLinkedList struct {
	tail *Node
	size int
}

func (cl *CircularLinkedList) Add(value int) {
	newNode := &Node{value: value}
	if cl.tail == nil {
		cl.tail = newNode
		cl.tail.next = newNode // Point to itself
	} else {
		newNode.next = cl.tail.next
		cl.tail.next = newNode
		cl.tail = newNode
	}
	cl.size++
}

func (cl *CircularLinkedList) Display() {
	if cl.tail == nil {
		fmt.Println("List is empty")
		return
	}
	current := cl.tail.next
	for i := 0; i < cl.size; i++ {
		fmt.Print(current.value, " -> ")
		current = current.next
	}
	fmt.Println("(circular)")
}

func main() {
	cl := &CircularLinkedList{}
	cl.Add(1)
	cl.Add(2)
	cl.Add(3)
	cl.Display() // Outputs: 1 -> 2 -> 3 -> (circular)
}

package main

import "fmt"

type QueueNode struct {
	value int
	next  *QueueNode
}

type Queue struct {
	front *QueueNode
	rear  *QueueNode
	size  int
}

func (q *Queue) Enqueue(value int) {
	newNode := &QueueNode{value: value}
	if q.rear == nil {
		q.front, q.rear = newNode, newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.size++
}

func (q *Queue) Dequeue() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	value := q.front.value
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	q.size--
	return value, nil
}

func lengthOfLongestSubstring(s string) int {
    charMap := make(map[rune]int) // A map to store the last seen index of each character
    left := 0                      // Left pointer of the sliding window
    maxLength := 0                 // Variable to store the maximum length of substring without repeating characters

    // Iterate through the string with the right pointer (i.e., the right side of the sliding window)
    for right, char := range s {
        // If the character has been seen and its index is within the current window, move the left pointer
        if lastIndex, found := charMap[char]; found && lastIndex >= left {
            left = lastIndex + 1 // Move the left pointer to the position right after the last occurrence
        }

        // Update the map with the current character's latest index
        charMap[char] = right

        // Calculate the length of the current window and update maxLength if it's larger
        currentLength := right - left + 1
        if currentLength > maxLength {
            maxLength = currentLength
        }
    }

    return maxLength // Return the maximum length found
}


func main() {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Dequeue()) // Outputs: 1
	fmt.Println(q.Dequeue()) // Outputs: 2
}
