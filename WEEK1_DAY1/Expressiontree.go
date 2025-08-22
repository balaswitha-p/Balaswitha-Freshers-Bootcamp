package main

import "fmt"

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func Preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%s ", node.Value)
	Preorder(node.Left)
	Preorder(node.Right)
}

func Postorder(node *Node) {
	if node == nil {
		return
	}
	Postorder(node.Left)
	Postorder(node.Right)
	fmt.Printf("%s ", node.Value)
}

func main() {
	nodeA := &Node{Value: "a"}
	nodeB := &Node{Value: "b"}
	nodeC := &Node{Value: "c"}

	nodeMinus := &Node{Value: "-", Left: nodeB, Right: nodeC}

	root := &Node{Value: "+", Left: nodeA, Right: nodeMinus}

	fmt.Println("Expression Tree Traversal:")

	fmt.Print("Preorder Traversal: ")
	Preorder(root)
	fmt.Println()

	fmt.Print("Postorder Traversal: ")
	Postorder(root)
	fmt.Println()
}
