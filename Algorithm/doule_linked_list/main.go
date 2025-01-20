package main

import (
	"fmt"
)

type Node struct {
	prev *Node
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtBegining(data int) {
	newNode := Node{prev: nil, data: data, next: nil}
	if list.head == nil {
		list.head = &newNode
		return
	}
	newNode.next = list.head
	list.head = &newNode
}

// [prev| 5| next] [prev| 6| next] [prev| 8| next]
func (list *LinkedList) insertAtEnd(data int) {
	newNode := Node{prev: nil, data: data, next: nil}
	if list.head == nil {
		list.head = &newNode
		return
	}
	curr := list.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &newNode
	newNode.prev = curr
}

// [prev| 5| next] [prev| 6| next] [prev| 8| next]
func (list *LinkedList) insertBeforeValue(data, target int) {
	newNode := Node{prev: nil, data: data, next: nil}
	if list.head == nil {
		fmt.Println("List empty. cannot insert before target")
		return
	}
	if list.head.data == target {
		newNode.next = list.head
		list.head.prev = &newNode
		list.head = &newNode
	}
	curr := list.head
	for curr != nil {
		if curr.next != nil && curr.next.data == target {
			newNode.next = curr.next
			newNode.prev = curr
			curr.next.prev = &newNode
			curr.next = &newNode
			return
		}
		curr = curr.next
	}
}

// [prev| 5| next] [prev| 6| next] [prev| 8| next]

func (list *LinkedList) insertAfterValue(data, target int) {
	newNode := Node{prev: nil, data: data, next: nil}
	if list.head == nil {
		fmt.Println("List is empty. cannot insert before target value")
		return
	}
	if list.head.data == target {
		list.head.next = &newNode
		newNode.prev = list.head
		return
	}
	curr := list.head
	for curr != nil {
		if curr.data == target {
			newNode.next = curr.next
			curr.next = &newNode
			curr.next.prev = &newNode
			newNode.prev = curr
			return
		}
		curr = curr.next
	}
}

func (list *LinkedList) deleteAtFirst() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}
	// for only one node
	if list.head.next == nil {
		list.head = nil
		return
	}
	list.head = list.head.next
	list.head.next.prev = nil
}

func (list *LinkedList) deleteAtEnd() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}
	if list.head.next == nil {
		list.head = nil
		return
	}
	curr := list.head
	for curr.next.next != nil {
		curr = curr.next
	}
	curr.next = nil
}

// [prev| 5| next] [prev| 6| next] [prev| 8| next]
func (list *LinkedList) deleteBeforeValue(value int) {
	if list.head == nil || list.head.next == nil {
		fmt.Println("List is empy")
		return
	}
	if list.head.next.data == value {
		list.head = list.head.next
		list.head.prev = nil
		return
	}
	curr := list.head.next
	for curr != nil {
		if curr.data == value {
			// Delete the node before `curr`
			if curr.prev != nil && curr.prev.prev != nil {
				curr.prev.prev.next = curr
				curr.prev = curr.prev.prev
			}
			return
		}
		curr = curr.next
	}
}

func (list *LinkedList) display() {
	curr := list.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}
func main() {
	list := LinkedList{}
	list.insertAtBegining(5)
	list.insertAtBegining(4)
	list.insertAtEnd(6)
	list.insertAtEnd(2)
	list.insertBeforeValue(10, 6)
	list.insertAfterValue(8, 2)
	list.deleteAtFirst()
	list.deleteAtEnd()
	list.deleteBeforeValue(2)
	list.display()
}
