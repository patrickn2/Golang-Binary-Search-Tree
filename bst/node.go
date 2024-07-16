package bst

type Node struct {
	value int
	left  *Node
	right *Node
}

func newNode(value int) *Node {
	return &Node{
		value: value,
		left:  nil,
		right: nil,
	}
}

func (n *Node) insertNode(newNode *Node) {
	if newNode.value > n.value {
		if n.right != nil {
			n.right.insertNode(newNode)
			return
		}
		n.right = newNode
		return
	}
	if newNode.value < n.value {
		if n.left != nil {
			n.left.insertNode(newNode)
			return
		}
		n.left = newNode
		return
	}
}

func (n *Node) searchNode(value int) bool {
	if value > n.value {
		if n.right != nil {
			return n.right.searchNode(value)
		}
		return false
	}

	if value < n.value {
		if n.left != nil {
			return n.left.searchNode(value)
		}
		return false
	}
	return true
}

func (n *Node) depthPreOrder() []int {
	result := []int{n.value}
	if n.left != nil {
		result = append(result, n.left.depthPreOrder()...)
	}
	if n.right != nil {
		result = append(result, n.right.depthPreOrder()...)
	}
	return result
}

func (n *Node) depthInOrder() []int {
	var result []int
	if n.left != nil {
		result = append(result, n.left.depthInOrder()...)
	}
	result = append(result, n.value)
	if n.right != nil {
		result = append(result, n.right.depthInOrder()...)
	}
	return result
}

func (n *Node) depthPostOrder() []int {
	var result []int
	if n.left != nil {
		result = append(result, n.left.depthPostOrder()...)
	}
	if n.right != nil {
		result = append(result, n.right.depthPostOrder()...)
	}
	return append(result, n.value)
}
