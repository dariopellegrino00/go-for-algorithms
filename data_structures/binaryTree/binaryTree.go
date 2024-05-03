package binarytree

import (
	"fmt"
	"strings"
)

type treeNode[T comparable] struct {
	value T
	left  *treeNode[T]
	right *treeNode[T]
}

type BinaryTree[T comparable] struct {
	root *treeNode[T]
}

func NewTree[T comparable](root *treeNode[T]) *BinaryTree[T] {
	return &BinaryTree[T]{root}
}

// terminal node with no left or right
func NewTermNode[T comparable](value T) *treeNode[T] {
	return &treeNode[T]{value, nil, nil}
}

func NewNode[T comparable](value T, left *treeNode[T], right *treeNode[T]) *treeNode[T] {
	return &treeNode[T]{value, left, right}
}

// i created a praticular in order printing
func (tree BinaryTree[T]) String() string {
	return tree.root.preorderString("")
}

// change how the offset works ??
func (node *treeNode[T]) preorderString(offset string) string {
	if node == nil {
		return offset + "*\n"
	}
	sprinted := fmt.Sprint(node.value)
	res := offset + "*" + sprinted + "\n"
	if node.left == nil && node.right == nil {
		return res
	}

	newOffset := offset + strings.Repeat(" ", len(sprinted))
	return res + node.left.preorderString(newOffset) + node.right.preorderString(newOffset)
}

func (tree *BinaryTree[T]) InOrderVisit() []T {
	return *tree.root.inOrderVisit(&([]T{}))
}

func (node *treeNode[T]) inOrderVisit(nodes *[]T) *[]T {
	if node != nil {
		node.left.inOrderVisit(nodes)
		*nodes = append(*nodes, node.value)
		node.right.inOrderVisit(nodes)
	}
	return nodes
}

func (tree *BinaryTree[T]) PreOrderVisit() []T {
	return *tree.root.preOrderVisit(&[]T{})
}

func (node *treeNode[T]) preOrderVisit(nodes *[]T) *[]T {
	if node != nil {
		*nodes = append(*nodes, node.value)
		node.left.preOrderVisit(nodes)
		node.right.preOrderVisit(nodes)
	}
	return nodes
}

func (tree *BinaryTree[T]) PostOrderVisit() []T {
	return *tree.root.postOrderVisit(&[]T{})
}

func (node *treeNode[T]) postOrderVisit(nodes *[]T) *[]T {
	if node != nil {
		node.left.postOrderVisit(nodes)
		node.right.postOrderVisit(nodes)
		*nodes = append(*nodes, node.value)
	}
	return nodes
}
