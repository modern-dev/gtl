// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package deque

type (
	// Deque basic generic deque (double-ended queue) implementation based on double-linked list
	Deque[T any] struct {
		head   *node[T]
		tail   *node[T]
		length int
	}

	// node is a generic double-linked list node use for Deque implementation
	node[T any] struct {
		Value T
		Next  *node[T]
		Prev  *node[T]
	}
)

// NewDeque TODO
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		head: &node[T]{},
		tail: &node[T]{},
	}
}

// Empty checks if Deque has no element
func (d *Deque[T]) Empty() bool {
	return d.length == 0
}

// Size returns the number of elements in Deque
// Complexity - O(1).
func (d *Deque[T]) Size() int {
	return d.length
}

// PushBack adds element to the end of Deque
// Complexity - O(1).
func (d *Deque[T]) PushBack(element T) {
	node := &node[T]{Value: element, Next: nil, Prev: nil}

	if d.Empty() {
		d.insertIntoEmpty(node)
		return
	}

	d.tail.Next = node
	node.Prev = d.tail
	d.tail = node
	d.length++
}

// PushFront adds element before first element of Deque
// Complexity - O(1).
func (d *Deque[T]) PushFront(element T) {
	node := &node[T]{Value: element, Next: nil, Prev: nil}

	if d.Empty() {
		d.insertIntoEmpty(node)
		return
	}

	d.head.Prev = node
	node.Next = d.head
	d.head = node
	d.length++
	return
}

// PopBack returns and removes the last element from Deque.
// Calling PopBack() on an empty Deque results in undefined behavior.
// Complexity - O(1).
func (d *Deque[T]) PopBack() T {
	if d.tail == nil {
		value := d.head.Value
		d.reset()
		return value
	}

	value := d.tail.Value
	d.tail = d.tail.Prev
	d.length--
	return value
}

// PopFront returns and removes the first element from Deque.
// Calling PopFront on an empty Deque results in undefined behavior.
// Complexity - O(1).
func (d *Deque[T]) PopFront() T {
	if d.head == nil {
		value := d.tail.Value
		d.reset()
		return value
	}

	value := d.head.Value
	d.head = d.head.Next
	d.length--
	return value
}

// Front returns values of the first element in Deque.
// Calling Front on an empty Deque results in undefined behavior.
// Complexity - O(1).
func (d *Deque[T]) Front() T {
	if d.head == nil {
		return d.tail.Value
	}

	return d.head.Value
}

// Back returns values of the last element in Deque.
// Calling Back on an empty Deque results in undefined behavior.
// Complexity - O(1).
func (d *Deque[T]) Back() T {
	if d.tail == nil {
		return d.head.Value
	}

	return d.tail.Value
}

func (d *Deque[T]) insertIntoEmpty(node *node[T]) {
	d.tail = node
	d.head = node
	d.length++
}

func (d *Deque[T]) reset() {
	d.tail = nil
	d.head = nil
	d.length = 0
}
