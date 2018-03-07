package binarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBinarySearchTree(t *testing.T) {
	n := NewBinarySearchTree(5)
	assert.Equal(t, n.Value, 5)
	assert.Nil(t, n.Left)
	assert.Nil(t, n.Right)
}

func TestInsert(t *testing.T) {
	n := NewBinarySearchTree(5)

	n.Insert(4)
	assert.Equal(t, n.Left.Value, 4)

	n.Insert(3)
	assert.Equal(t, n.Left.Left.Value, 3)

	n.Insert(6)
	assert.Equal(t, n.Right.Value, 6)

	assert.Error(t, n.Insert(5))
}

func TestSearch(t *testing.T) {
	n := NewBinarySearchTree(5)

	n.Insert(7)
	n.Insert(9)
	n.Insert(4)

	assert.Equal(t, n.Search(4).Value, 4)
	assert.Nil(t, n.Search(1))
}

func TestFindMinNode(t *testing.T) {
	n := NewBinarySearchTree(5)
	n.Insert(4)
	n.Insert(1)
	n.Insert(9)

	assert.Equal(t, n.FindMinNode().Value, 1)

	n = NewBinarySearchTree(5)
	assert.Equal(t, n.FindMinNode(), n)
}

func TestFindMaxNode(t *testing.T) {
	n := NewBinarySearchTree(5)
	n.Insert(4)
	n.Insert(1)
	n.Insert(9)

	assert.Equal(t, n.FindMaxNode().Value, 9)

	n = NewBinarySearchTree(5)
	assert.Equal(t, n.FindMaxNode(), n)
}

func TestDelete(t *testing.T) {
	n := NewBinarySearchTree(5)
	assert.Equal(t, n, n.Delete(1))

	assert.Nil(t, n.Delete(5))

	// left and right child are nil
	n.Insert(6)
	assert.NotNil(t, n.Right.Value)
	assert.Nil(t, n.Delete(6).Right)

	// left child is nil and right child is not nil
	n = NewBinarySearchTree(5)
	n.Insert(8)
	n.Insert(12)
	n.Insert(9)
	assert.NotNil(t, n.Right.Right.Left)
	n = n.Delete(8)
	assert.Equal(t, n.Right.Value, 12)
	assert.Equal(t, n.Right.Left.Value, 9)

	// left child is not nil and right child is nil
	n = NewBinarySearchTree(8)
	n.Insert(7)
	n.Insert(6)
	assert.NotNil(t, n.Left.Left)
	n = n.Delete(7)
	assert.Equal(t, n.Left.Value, 6)
	assert.Nil(t, n.Left.Left)

	// left child and right are not nil
	n = NewBinarySearchTree(5)
	n.Insert(8)
	n.Insert(12)
	n.Insert(9)
	n.Insert(7)
	assert.NotNil(t, n.Right.Right.Left)
	n = n.Delete(8)
	assert.Nil(t, n.Right.Right.Left)
	assert.Equal(t, n.Right.Value, 9)
	assert.Equal(t, n.Right.Left.Value, 7)
	assert.Equal(t, n.Right.Right.Value, 12)
}

func TestPreOrderTraverse(t *testing.T) {
	n := NewBinarySearchTree(5)
	n.Insert(5)
	n.Insert(3)
	n.Insert(1)
	n.Insert(4)
	n.Insert(8)
	n.Insert(7)
	n.Insert(10)

	nodes := n.PreOrderTraverse()
	assert.Equal(t, nodes[0].Value, 5)
	assert.Equal(t, nodes[1].Value, 3)
	assert.Equal(t, nodes[2].Value, 1)
	assert.Equal(t, nodes[3].Value, 4)
	assert.Equal(t, nodes[4].Value, 8)
	assert.Equal(t, nodes[5].Value, 7)
	assert.Equal(t, nodes[6].Value, 10)
}

func TestPostOrderTraverse(t *testing.T) {
	n := NewBinarySearchTree(5)
	n.Insert(5)
	n.Insert(3)
	n.Insert(1)
	n.Insert(4)
	n.Insert(8)
	n.Insert(7)
	n.Insert(10)

	nodes := n.PostOrderTraverse()
	assert.Equal(t, nodes[0].Value, 1)
	assert.Equal(t, nodes[1].Value, 3)
	assert.Equal(t, nodes[2].Value, 4)
	assert.Equal(t, nodes[3].Value, 5)
	assert.Equal(t, nodes[4].Value, 7)
	assert.Equal(t, nodes[5].Value, 8)
	assert.Equal(t, nodes[6].Value, 10)
}

func TestInOrderTraverse(t *testing.T) {
	n := NewBinarySearchTree(5)
	n.Insert(5)
	n.Insert(3)
	n.Insert(1)
	n.Insert(4)
	n.Insert(8)
	n.Insert(7)
	n.Insert(10)

	nodes := n.InOrderTraverse()
	assert.Equal(t, nodes[0].Value, 1)
	assert.Equal(t, nodes[1].Value, 4)
	assert.Equal(t, nodes[2].Value, 3)
	assert.Equal(t, nodes[3].Value, 7)
	assert.Equal(t, nodes[4].Value, 10)
	assert.Equal(t, nodes[5].Value, 8)
	assert.Equal(t, nodes[6].Value, 5)
}
