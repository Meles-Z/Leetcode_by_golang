package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InserAtEnd(data int) {
	newNode := Node{data: data, next: nil}
	if list.head == nil {
		list.head = &newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = &newNode
}
func (list *LinkedList) InsertAtFirst(data int) {
	newNode := Node{data: data, next: list.head}
	if list.head == nil {
		list.head = &newNode
		return
	}
	list.head = &newNode
}

func (list *LinkedList) InsertAfterValue(data, val int) {
	newNode := Node{data: data, next: nil}
	if list.head == nil {
		list.head = &newNode
	}
	current := list.head
	for current != nil {
		if current.data == val {
			newNode.next = current.next
			current.next = &newNode
		}
		current = current.next

	}

}

func (list *LinkedList) InsertBeforeValue(data, val int) {
	newNode := &Node{data: data}

	// Case 1: If the list is empty
	if list.head == nil {
		fmt.Println("The list is empty. Cannot insert before the value.")
		return
	}

	// Case 2: If the value is in the head node
	if list.head.data == val {
		newNode.next = list.head
		list.head = newNode
		return
	}

	// Case 3: Traverse the list to find the value
	current := list.head
	for current.next != nil {
		if current.next.data == val {
			newNode.next = current.next // Link the new node to the target node
			current.next = newNode      // Link the previous node to the new node
			return                      // Exit after insertion
		}
		current = current.next
	}

	// If the value is not found
	fmt.Println("Value not found in the list.")
}

func (list *LinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func main() {
	list := LinkedList{}
	list.InserAtEnd(5)
	list.InserAtEnd(6)
	list.InsertAtFirst(4)
	list.InsertAfterValue(9, 5)
	list.InsertBeforeValue(4,9)
	list.Display()
}
