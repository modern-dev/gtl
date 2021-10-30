// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package stack

import (
	. "github.com/modern-dev/gtl/containers/deque"
)

// Stack is a container adapter that gives the programmer the functionality of a stack
// - specifically, a LIFO (last-in, first-out) data structure.
type Stack[T any] struct {
	dq *Deque[T]
}

// NewStack TODO
func NewStack[T any]() *Stack[T] {
	dq := NewDeque[T]()

	return &Stack[T]{dq}
}

// Size returns the number of elements in the underlying container.
// Complexity - constant e.g. O(1).
// Returns the number of elements in the container.
func (s *Stack[T]) Size() int {
	return s.dq.Size()
}

// Empty checks if the underlying container has no elements.
// Complexity - constant e.g. O(1).
// Returns true if the underlying container is empty, false otherwise.
func (s *Stack[T]) Empty() bool {
	return s.dq.Empty()
}

// Push pushes the given element value to the top of the Stack.
// Complexity - constant e.g. O(1).
func (s *Stack[T]) Push(item T) {
	s.dq.PushBack(item)
}

// Top returns reference to the top element in the Stack.
// This is the most recently pushed element. This element will be removed on a call to Pop.
// Complexity - constant e.g. O(1).
func (s *Stack[T]) Top() T {
	return s.dq.Back()
}

// Pop removes the top element from the Stack.
// Returns the object at the top of this Stack.
func (s *Stack[T]) Pop() T {
	return s.dq.PopBack()
}
