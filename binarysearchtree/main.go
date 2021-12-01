package main

import "fmt"

type Tree struct {
	Root *Node
}

type Node struct {
	key   byte
	left  *Node
	right *Node
}

//Creating a tree
func (t *Tree) Insert(data byte) {
	if t.Root == nil {
		t.Root = &Node{
			key: data,
		}
	} else {
		t.Root.Insert(data)
	}
}

// Insert inserts data into the Node
func (n *Node) Insert(data byte) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.Insert(data)
		}
	}
}

func main() {
	var t Tree

	t.Insert('F')
	t.Insert('B')
	t.Insert('A')
	t.Insert('D')
	t.Insert('C')
	t.Insert('E')
	t.Insert('G')
	t.Insert('I')
	t.Insert('H')

	fmt.Printf("PreOrder: ")
	printPreOrderTraversal(t.Root)
	fmt.Println()

	fmt.Printf("PostOrder: ")
	printPostOrderTraversal(t.Root)
	fmt.Println()

	fmt.Printf("InOrder: ")
	printInOrderTraversal(t.Root)
	fmt.Println()

	// Find ...
	err := t.Find('I')
	if err != nil {
		fmt.Printf("Value %c Not Found\n", 'I')
	} else {
		fmt.Printf("Value %c Found\n", 'I')
	}

	err = t.Find('L')
	if err != nil {
		fmt.Printf("Value %c Not Found\n", 'L')
	} else {
		fmt.Printf("Value %c Found\n", 'L')
	}
}

// printPreOrderTraversal Performs the PreOrder traversal in the tree
func printPreOrderTraversal(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%c ", n.key)
		printPreOrderTraversal(n.left)
		printPreOrderTraversal(n.right)
	}
}

// printPostOrderTraversal Performs the PostOrder traversal in the tree
func printPostOrderTraversal(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrderTraversal(n.left)
		printPostOrderTraversal(n.right)
		fmt.Printf("%c ", n.key)
	}
}

// printInOrderTraversal Performs the InOrder traversal in the tree
func printInOrderTraversal(n *Node) {
	if n == nil {
		return
	} else {
		printInOrderTraversal(n.left)
		fmt.Printf("%c ", n.key)
		printInOrderTraversal(n.right)
	}
}

//Find looks for a particular node value in the tree
func (b *Tree) Find(data byte) error {
	node := b.findRec(b.Root, data)
	if node == nil {
		return fmt.Errorf("value: %d not found in tree", data)
	}
	return nil
}

func (b *Tree) findRec(node *Node, data byte) *Node {
	if node == nil {
		return nil
	}
	if node.key == data {
		return b.Root
	}
	if data < node.key {
		return b.findRec(node.left, data)
	}
	return b.findRec(node.right, data)
}
