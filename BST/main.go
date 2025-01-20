package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

// insert the node

func (n *Node) insert(k int) {
	if n.data < k {
		// move to right
		if n.right == nil {
			n.right = &Node{data: k}
		} else {
			n.right.insert(k)
		}
	} else if n.data > k {
		// move to left
		if n.left == nil {
			n.left = &Node{data: k}
		} else {
			n.left.insert(k)
		}
	}
}

// search the node
func (node *Node) Search(k int) bool {
	if node == nil {
		return false
	}
	if node.data < k {
		//move to right
		return node.right.Search(k)
	} else if node.data > k {
		// move to left
		return node.left.Search(k)
	}
	return true
}
func (node *Node) print() {
	if node == nil {
		return
	}
	node.left.print()
	fmt.Println(node.data)
	node.right.print()
}
func main() {
	node := Node{
		data: 50,
	}
	node.insert(90)
	node.insert(100)
	node.insert(200)
	node.insert(30)
	node.insert(45)
	fmt.Println(node)
	fmt.Println(node.Search(100))
	node.print()
}
