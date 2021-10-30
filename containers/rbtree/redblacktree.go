// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package rbtree

import (
	"constraints"
	"github.com/modern-dev/gtl/utility"
)

const (
	red color = iota
	black
)

type (
	RBTree[T comparable] struct {
		root    *nodeHandle[T]
		nilNode *nodeHandle[T]
		cmpInst utility.Compare[T]
		size    int
		dupl    bool
	}

	nodeHandle[T any] struct {
		col    color
		left   *nodeHandle[T]
		right  *nodeHandle[T]
		parent *nodeHandle[T]
		value  T
	}

	color uint8
)

func NewRBTree[T constraints.Ordered](allowDuplicates bool) *RBTree[T] {
	nilNode := &nodeHandle[T]{col: black}

	return &RBTree[T]{
		root:    nilNode,
		nilNode: nilNode,
		cmpInst: &utility.Less[T]{},
		dupl:    allowDuplicates,
	}
}

// NewRBTreeWithComparator creates an empty tree with provided comparator for items.
func NewRBTreeWithComparator[T comparable](comparator utility.Compare[T], allowDuplicates bool) *RBTree[T] {
	nilNode := &nodeHandle[T]{col: black}

	return &RBTree[T]{
		root:    nilNode,
		nilNode: nilNode,
		cmpInst: comparator,
		dupl:    allowDuplicates,
	}
}

// Insert adds value into the tree.
// Complexity O(log n), where n is the number of elements in the tree.
func (rbt *RBTree[T]) Insert(value T) {
	newNode := &nodeHandle[T]{
		col:    red,
		left:   rbt.nilNode,
		right:  rbt.nilNode,
		parent: rbt.nilNode,
		value:  value,
	}

	parentNode := rbt.findInsertNode(newNode)
	newNode.parent = parentNode

	rbt.addChild(parentNode, newNode)

	rbt.size++

	rbt.insertFixup(newNode)
}

// Find tries to find the value in the tree.
// Returns 2 values.
// First value is an item if it was found, otherwise zero value for type parameter.
// Second value is bool indicating whether an item was found.
// Complexity O(log n), where n is the number of elements in the tree.
func (rbt *RBTree[T]) Find(value T) (T, bool) {
	node, found := rbt.searchFromNode(rbt.root, value)

	return node.value, found
}

// Max returns max item in the tree according to the comparator.
// Complexity O(log n), where n is the number of elements in tree.
func (rbt *RBTree[T]) Max() T {
	node := rbt.maximum(rbt.root)

	return node.value
}

// Min returns min item in the tree according to the comparator
// Complexity O(log n), where n is the number of elements in the tree.
func (rbt *RBTree[T]) Min() T {
	node := rbt.minimum(rbt.root)

	return node.value
}

// Erase deletes the item from the tree. Has no effect if the item was not in the tree.
// Complexity O(log n), where n is the number of elements in the tree.
func (rbt *RBTree[T]) Erase(value T) {
	node, found := rbt.searchFromNode(rbt.root, value)

	if !found {
		return
	}

	if color, nodeToFix := rbt.deleteNode(node); color == black {
		rbt.deleteFixup(nodeToFix)
	}

	rbt.size--
}

// Size return the number of elements in the tree.
// Complexity O(1).
func (rbt *RBTree[T]) Size() int {
	return rbt.size
}

func (rbt *RBTree[T]) Empty() bool {
	return rbt.Size() == 0
}

func (rbt *RBTree[T]) searchFromNode(node *nodeHandle[T], value T) (*nodeHandle[T], bool) {
	it := node

	for it != rbt.nilNode {
		cmpRes := rbt.cmpInst.Cmp(value, it.value)

		// discard duplicates
		if value == it.value && !rbt.dupl {
			return it, true
		}

		if cmpRes {
			it = it.left
		} else {
			it = it.right
		}
	}

	return rbt.nilNode, false
}

func (rbt *RBTree[T]) findInsertNode(newNode *nodeHandle[T]) *nodeHandle[T] {
	y, x := rbt.nilNode, rbt.root

	for x != rbt.nilNode {
		y = x
		if rbt.cmpInst.Cmp(newNode.value, x.value) {
			x = x.left
		} else {
			x = x.right
		}
	}

	return y
}

func (rbt *RBTree[T]) addChild(node, child *nodeHandle[T]) {
	if node == rbt.nilNode {
		rbt.root = child

		return
	}

	if rbt.cmpInst.Cmp(child.value, node.value) {
		node.left = child

		return
	}

	node.right = child
}

func (rbt *RBTree[T]) insertFixup(z *nodeHandle[T]) {
	for z.parent.col == red {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right

			if y.col == red {
				z.parent.col = black
				y.col = black
				z.parent.parent.col = red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent

					rbt.leftRotate(z)
				}

				z.parent.col = black
				z.parent.parent.col = red

				rbt.rightRotate(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left

			if y.col == red {
				z.parent.col = black
				y.col = black
				z.parent.parent.col = red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent

					rbt.rightRotate(z)
				}

				z.parent.col = black
				z.parent.parent.col = red

				rbt.leftRotate(z.parent.parent)
			}
		}
	}

	rbt.root.col = black
}

func (rbt *RBTree[T]) deleteNode(z *nodeHandle[T]) (color, *nodeHandle[T]) {
	var x *nodeHandle[T]
	y := z
	originalColor := y.col

	if z.left == rbt.nilNode {
		x = z.right

		rbt.transplant(z, z.right)

		return originalColor, x
	}

	if z.right == rbt.nilNode {
		x = z.left

		rbt.transplant(z, z.left)

		return originalColor, x
	}

	y = rbt.minimum(z.right)

	originalColor = y.col
	x = y.right

	if y.parent == z {
		x.parent = y
	} else {
		rbt.transplant(y, y.right)

		y.right = z.right
		y.right.parent = y
	}

	rbt.transplant(z, y)

	y.left = z.left
	y.left.parent = y
	y.col = z.col

	return originalColor, x
}

func (rbt *RBTree[T]) deleteFixup(x *nodeHandle[T]) {
	for x != rbt.root && x.col == black {
		if x == x.parent.left {
			w := x.parent.right

			if w.col == red {
				w.col = black
				x.parent.col = red

				rbt.leftRotate(x.parent)

				w = x.parent.right
			}

			if w.left.col == black && w.right.col == black {
				w.col = red
				x = x.parent
			} else {
				if w.right.col == black {
					w.left.col = black
					w.col = red

					rbt.rightRotate(w)

					w = x.parent.right
				}

				w.col = x.parent.col
				x.parent.col = black
				w.right.col = black

				rbt.leftRotate(x.parent)

				x = rbt.root
			}
		} else {
			w := x.parent.left

			if w.col == red {
				w.col = black
				x.parent.col = red

				rbt.rightRotate(x.parent)

				w = x.parent.left
			}

			if w.left.col == black && w.right.col == black {
				w.col = red
				x = x.parent
			} else {
				if w.left.col == black {
					w.right.col = black
					w.col = red

					rbt.leftRotate(w)

					w = x.parent.left
				}

				w.col = x.parent.col
				x.parent.col = black
				w.left.col = black

				rbt.rightRotate(x.parent)

				x = rbt.root
			}
		}
	}

	x.col = black
}

func (rbt *RBTree[T]) transplant(a, b *nodeHandle[T]) {
	if a.parent == rbt.nilNode {
		rbt.root = b
	} else if a == a.parent.left {
		a.parent.left = b
	} else {
		a.parent.right = b
	}

	b.parent = a.parent
}

func (rbt *RBTree[T]) minimum(node *nodeHandle[T]) *nodeHandle[T] {
	for node.left != rbt.nilNode {
		node = node.left
	}

	return node
}

func (rbt *RBTree[T]) maximum(node *nodeHandle[T]) *nodeHandle[T] {
	for node.right != rbt.nilNode {
		node = node.right
	}

	return node
}

func (rbt *RBTree[T]) leftRotate(x *nodeHandle[T]) {
	y := x.right
	x.right = y.left

	if y.left != rbt.nilNode {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == rbt.nilNode {
		rbt.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func (rbt *RBTree[T]) rightRotate(x *nodeHandle[T]) {
	y := x.left
	x.left = y.right

	if y.right != rbt.nilNode {
		y.right.parent = x
	}

	y.parent = x.parent

	if x.parent == rbt.nilNode {
		rbt.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}

	y.right = x
	x.parent = y
}
