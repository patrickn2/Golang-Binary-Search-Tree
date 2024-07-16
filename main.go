package main

import (
	"fmt"

	sBst "github.com/patrickn2/golangbinarysearchtree/bst"
)

func main() {
	bst := sBst.NewBst()

	fmt.Println("Tree is empty:", bst.IsEmpty())

	bst.Add(10)
	bst.Add(5)
	bst.Add(15)
	bst.Add(3)
	bst.Add(7)

	fmt.Println("Depth Pre Order", bst.DepthPreOrder())
	fmt.Println("Depth In Order", bst.DepthInOrder())
	fmt.Println("Depth Post Order", bst.DepthPostOrder())
	fmt.Println("Breadth Traversal Order", bst.LevelOrder())
	fmt.Println("Find Min", bst.FindMin())
	fmt.Println("Find Max", bst.FindMax())

}
