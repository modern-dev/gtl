// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package queue

import (
	. "github.com/modern-dev/gtl/containers/deque"
)

// Queue is a container adapter that gives the programmer the functionality of a queue
// - specifically, a FIFO (first-in, first-out) data structure.
type Queue[T any] struct {
	dq Deque[T]
}

// Size returns number of elements in queue
// Complexity O(1)
func (q *Queue[T]) Size() int {
	return q.dq.Size()
}

// Empty checks if Queue has no element
// Complexity O(1)
func (q *Queue[T]) Empty() bool {
	return q.dq.Empty()
}

// Front returns value of the first element in Queue
// Complexity O(1)
func (q *Queue[T]) Front() T {
	return q.dq.Front()
}

// Back returns value of the last element in Queue
// Complexity O(1)
func (q *Queue[T]) Back() T {
	return q.dq.Back()
}

// Push inserts element at the end of Queue
// Complexity O(1)
func (q *Queue[T]) Push(element T) {
	q.dq.PushBack(element)
}

// Pop removes and returns first element of Queue
// Complexity O(1)
func (q *Queue[T]) Pop() T {
	return q.dq.PopFront()
}
