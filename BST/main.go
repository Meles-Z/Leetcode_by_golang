package main

import "fmt"

type Node struct {
	key   byte
	left  *Node
	right *Node
}
type Tree struct {
    root *Node
}

func (t *Tree) insert(data byte){
    newNode:=&Node{key: data}
    if t.root==nil{
        t.root=newNode
    }else{
        t.root.insert(data)
    }
}

func (n *Node)insert(data byte){
    if data<=n.key{
        if n.left==nil{
            n.left=&Node{key: data}
        }else{
            n.left.insert(data)
        }
    }else{
        if n.right==nil{
            n.right=&Node{key: data}
        }else{
            n.right.insert(data)
        }
    }
}
func (t *Tree) PrintTree() {
	if t.root == nil {
		fmt.Println("The tree is empty")
	} else {
		t.root.print("", true, "")
	}
}

func (n *Node) print(prefix string, isTail bool, indent string) {
	if n.right != nil {
		n.right.print(prefix+indent, false, "│   ")
	}
	fmt.Printf("%s%s%c\n", prefix, indent, n.key)
	if n.left != nil {
		n.left.print(prefix+indent, true, "    ")
	}
}

func main() {
    var t Tree
	t.insert('F')
	t.insert('G')
	t.insert('S')
	t.insert('R')
	t.insert('A')
	t.insert('C')

	t.PrintTree()

}
