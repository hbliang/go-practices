package binarysearchtree

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewBinarySearchTree(value int) *Node {
	return &Node{Value: value}
}

func (n *Node) Insert(value int) error {
	if value == n.Value {
		return fmt.Errorf("You can't insert duplicate value %d", value)
	}

	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else if value > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}

	return nil
}

func (n *Node) Search(value int) *Node {
	if n == nil {
		return nil
	}

	if value == n.Value {
		return n
	} else if value < n.Value {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}

func (n *Node) Delete(value int) *Node {
	if n == nil {
		return nil
	}

	if value < n.Value {
		n.Left = n.Left.Delete(value)
		return n
	} else if value > n.Value {
		n.Right = n.Right.Delete(value)
		return n
	}

	if n.Left == nil && n.Right == nil {
		return nil
	} else if n.Left != nil && n.Right == nil {
		return n.Left
	} else if n.Left == nil && n.Right != nil {
		return n.Right
	} else {
		minNode := n.Right.FindMinNode()
		n.Value = minNode.Value
		n.Right = n.Right.Delete(minNode.Value)
		return n
	}
}

func (n *Node) FindMinNode() *Node {
	if n.Left == nil {
		return n
	}

	minNode := n.Left

	for minNode.Left != nil {
		minNode = minNode.Left
	}

	return minNode
}

func (n *Node) FindMaxNode() *Node {
	if n.Right == nil {
		return n
	}

	maxNode := n.Right

	for maxNode.Right != nil {
		maxNode = maxNode.Right
	}

	return maxNode
}

func (n *Node) PreOrderTraverse() []*Node {
	var nodes []*Node

	nodes = append(nodes, n)

	if n.Left != nil {
		nodes = append(nodes, n.Left.PreOrderTraverse()...)
	}

	if n.Right != nil {
		nodes = append(nodes, n.Right.PreOrderTraverse()...)
	}

	return nodes
}

func (n *Node) PostOrderTraverse() []*Node {
	var nodes []*Node

	if n.Left != nil {
		nodes = append(nodes, n.Left.PostOrderTraverse()...)
	}

	nodes = append(nodes, n)

	if n.Right != nil {
		nodes = append(nodes, n.Right.PostOrderTraverse()...)
	}

	return nodes
}

func (n *Node) InOrderTraverse() []*Node {
	var nodes []*Node

	if n.Left != nil {
		nodes = append(nodes, n.Left.InOrderTraverse()...)
	}

	if n.Right != nil {
		nodes = append(nodes, n.Right.InOrderTraverse()...)
	}

	nodes = append(nodes, n)

	return nodes
}
