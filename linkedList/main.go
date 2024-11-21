package main

import "fmt"

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
	list.InsertAtFirst(7)
	list.Display()
}