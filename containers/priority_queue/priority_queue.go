// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package priority_queue provides ...
package priority_queue

import (
	"constraints"
	"github.com/modern-dev/gtl/utility"
)

type (
	// A PriorityQueue is a container adaptor that provides constant time lookup of the largest (by default) element,
	// at the expense of logarithmic insertion and extraction.
	// A user-provided comparator can be supplied to change the ordering, e.g. using utility.Greater[T]
	// would cause the smallest element to appear as the Top().
	PriorityQueue[T any] struct {
		heapList []T
		size     int
		cmpInst  utility.Compare[T]
	}
)

// NewPriorityQueue constructs the PriorityQueue.
func NewPriorityQueue[T constraints.Ordered]() *PriorityQueue[T] {
	var emptyEl T

	return &PriorityQueue[T]{
		heapList: []T{emptyEl},
		cmpInst:  &utility.Less[T]{},
	}
}

// NewPriorityQueueWithComparator constructs the PriorityQueue.
// A utility.Compare type providing a strict weak ordering.
func NewPriorityQueueWithComparator[T any](comparator utility.Compare[T]) *PriorityQueue[T] {
	var emptyEl T

	return &PriorityQueue[T]{
		heapList: []T{emptyEl},
		cmpInst:  comparator,
	}
}

// Size returns the number of elements in the PriorityQueue.
// Complexity - constant.
func (h *PriorityQueue[T]) Size() int {
	return h.size
}

// Empty checks if the PriorityQueue has no elements
// Complexity - constant.
func (h *PriorityQueue[T]) Empty() bool {
	return h.Size() == 0
}

// Push pushes the given element value to the PriorityQueue.
// Complexity - logarithmic number of comparisons.
func (h *PriorityQueue[T]) Push(value T) {
	h.heapList = append(h.heapList, value)
	h.size++

	h.siftUp(h.size)
}

// Pop removes the top element from the PriorityQueue
// Complexity - logarithmic number of comparisons.
func (h *PriorityQueue[T]) Pop() T {
	root := h.heapList[1]

	h.heapList[1] = h.heapList[h.size]
	h.heapList = h.heapList[:len(h.heapList)-1]
	h.size--

	h.siftDown(1)

	return root
}

// Top returns reference to the top element in the PriorityQueue.
// Complexity - constant.
func (h *PriorityQueue[T]) Top() T {
	return h.heapList[1]
}

func (h *PriorityQueue[T]) siftUp(i int) {
	for i/2 > 0 {
		if !h.cmpInst.Cmp(h.heapList[i], h.heapList[i/2]) {
			h.heapList[i], h.heapList[i/2] = h.heapList[i/2], h.heapList[i]
		}

		i /= 2
	}
}

func (h *PriorityQueue[T]) minChild(i int) int {
	if (i*2)+1 > h.size {
		return i * 2
	}

	if !h.cmpInst.Cmp(h.heapList[i*2], h.heapList[i*2+1]) {
		return i * 2
	}

	return i*2 + 1
}

func (h *PriorityQueue[T]) siftDown(i int) {
	for (i * 2) <= h.size {
		mc := h.minChild(i)

		if h.cmpInst.Cmp(h.heapList[i], h.heapList[mc]) {
			h.heapList[i], h.heapList[mc] = h.heapList[mc], h.heapList[i]
		}

		i = mc
	}
}
