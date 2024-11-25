package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InsertAtBegin(data int) {
	newNode := Node{data: data, next: nil}
	if list.head == nil {
		list.head = &newNode
		return
	}
	newNode.next = list.head
	list.head = &newNode
}
func (list *LinkedList) insertAtEnd(data int) {
	newNode := Node{data: data, next: nil}
	if list.head == nil {
		list.head = &newNode
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = &newNode
}
func (list *LinkedList) insertBeforeValue(data, target int) {
	newNode := Node{data: data, next: nil}
	if list.head == nil {
		fmt.Println("Target value is empty")
		return
	}
	if list.head.data == target {
		newNode.next = list.head
		list.head = &newNode
	}
	current := list.head
	for current.next != nil {
		if current.next.data == target {
			newNode.next = current.next
			current.next = &newNode
			return
		}
		current = current.next
	}

}
func (list *LinkedList) display() {
	if list.head != nil {
		current := list.head
		for current != nil {
			fmt.Println(current.data)
			current = current.next

		}
	}
}
func main() {
	list := LinkedList{}
	list.InsertAtBegin(4)
	list.InsertAtBegin(5)
	list.insertAtEnd(6)
	list.insertBeforeValue(10, 6)
	list.display()
}
