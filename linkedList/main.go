package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtBegin(data int) {
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

// [3|next][8|next][9|next]
func (list *LinkedList) insertBeforeValue(data, targetValue int) {
	newNode := Node{data: data, next: nil}

	if list.head.data == targetValue {
		newNode.next = list.head
		list.head = &newNode
	}
	current := list.head
	for current.next != nil {
		if current.next.data == targetValue {
			newNode.next = current.next
			current.next = &newNode
			return
		}
		current = current.next
	}

}

// [3|next][8|next][9|next]
func (list *LinkedList) insertAfterValue(data int, targetValue int) {
	newNode := Node{data: data, next: nil}
	if list.head == nil {
		list.head = &newNode
	}
	current := list.head
	for current != nil {
		if current.data == targetValue {
			newNode.next = current.next
			current.next = &newNode
			return

		}
		current = current.next
	}

}

func (list *LinkedList) deleteAtFirst() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}
	list.head = list.head.next
}

// [3|next][8|next][9|next]
func (list *LinkedList) deleteAtEnd() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}
	if list.head.next == nil {
		list.head = nil
		return
	}
	current := list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

// [3|next][8|next][9|next]
func (list *LinkedList) deleteAfterValue(target int) {
	if list.head == nil {
		fmt.Println("List is empty")
	}
	for list.head.next == nil {
		list.head = nil
		return
	}
	current := list.head
	for current != nil {
		if current.data == target {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

// [3|next][8|next][9|next]
func (list *LinkedList) deleteBeforeValue(target int) {
	if list.head == nil || list.head.next == nil {
		fmt.Println("Not enough to perfom operations")
		return
	}
	// if the second node is targeted
	if list.head.next.data == target {
		list.head = list.head.next
		return
	}

	// traverse on list
	prev := list.head
	current := list.head.next
	for current.next != nil {
		if current.next.data == target {
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}
}
func (list *LinkedList) display() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
func main() {
	list := LinkedList{}
	list.insertAtBegin(3)
	list.insertAtBegin(2)
	list.insertAtEnd(7)
	list.insertAtEnd(4)
	list.insertBeforeValue(9, 7)
	list.insertBeforeValue(10, 3)
	list.insertAfterValue(5, 4)
	list.insertAfterValue(8, 5)
	list.deleteAtFirst()
	list.deleteAtEnd()
	list.deleteAfterValue(9)
	list.deleteBeforeValue(4)
	list.display()
}