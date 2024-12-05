package main

import (
	"fmt"
)

// Define a ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to add two linked lists
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}
	return head.Next
}

// Helper function to create a linked list from a slice
func createLinkedList(nums []int) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummyHead.Next
}

// Helper function to print a linked list
func printLinkedList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print(" -> ")
		}
		l = l.Next
	}
	fmt.Println()
}

func main() {
	// Example usage
	l1 := createLinkedList([]int{2, 4, 3}) // 342
	l2 := createLinkedList([]int{5, 6, 4}) // 465

	fmt.Println("First List:")
	printLinkedList(l1)
	fmt.Println("Second List:")
	printLinkedList(l2)

	result := addTwoNumbers(l1, l2)
	fmt.Println("Resultant List:")
	printLinkedList(result) // Output: 7 -> 0 -> 8 (807)
}
