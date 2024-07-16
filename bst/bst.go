package bst

import "errors"

type BST struct {
	root *Node
}

func NewBst() *BST {
	return &BST{
		root: nil,
	}
}

func (b *BST) IsEmpty() bool {
	return b.root == nil
}

func (b *BST) Add(value int) {
	nNode := newNode(value)
	if b.IsEmpty() {
		b.root = nNode
		return
	}
	b.root.insertNode(nNode)
}

func (b *BST) Search(value int) bool {
	if b.IsEmpty() {
		return false
	}
	return b.root.searchNode(value)
}

func (b *BST) DepthPreOrder() []int {
	if b.IsEmpty() {
		return []int{}
	}
	return b.root.depthPreOrder()
}

func (b *BST) DepthInOrder() []int {
	if b.IsEmpty() {
		return []int{}
	}
	return b.root.depthInOrder()
}

func (b *BST) DepthPostOrder() []int {
	if b.IsEmpty() {
		return []int{}
	}
	return b.root.depthPostOrder()
}

func (b *BST) LevelOrder() []int {
	if b.IsEmpty() {
		return []int{}
	}
	var result []int
	var queue []*Node
	queue = append(queue, b.root)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		result = append(result, curr.value)
		if curr.left != nil {
			queue = append(queue, curr.left)
		}
		if curr.right != nil {
			queue = append(queue, curr.right)
		}
	}
	return result
}

func (b *BST) FindMin() (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("Binary Search Tree is Empty")
	}
	curr := b.root
	for {
		if curr.left != nil {
			curr = curr.left
			continue
		}
		return curr.value, nil
	}
}

func (b *BST) FindMax() (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("Binary Search Tree is Empty")
	}
	curr := b.root
	for {
		if curr.right != nil {
			curr = curr.right
			continue
		}
		return curr.value, nil
	}
}
