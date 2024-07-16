package main

import (
	"fmt"

	sBst "github.com/patrickn2/Golang-Binary-Search-Tree/bst"
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
	min, _ := bst.FindMin()
	fmt.Println("Find Min", min)
	max, _ := bst.FindMax()
	fmt.Println("Find Max", max)

	bst.Delete(5)
	fmt.Println("Depth Pre Order After Delete 5", bst.DepthPreOrder())

}
