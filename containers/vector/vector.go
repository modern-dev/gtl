// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package vector

type Vector[T any] struct {
	ar []T
}

func NewVector[T any]() *Vector[T] {
	return &Vector[T]{
		[]T{},
	}
}

// At returns a reference to the element at specified location pos, with bounds checking.
// If pos is not within the range of the container, a panic is thrown.
// Complexity - O(1).
func (v *Vector[T]) At(pos int) T {
	return v.ar[pos]
}

// PushBack appends the given element value to the end of the container.
// Complexity - amortized constant e.g. O(1).
func (v *Vector[T]) PushBack(item T) {
	v.ar = append(v.ar, item)
}

// Size returns the number of elements in the container.
// Complexity - O(1).
func (v *Vector[T]) Size() int {
	return len(v.ar)
}

// Empty checks if the container has no elements.
// Complexity - O(1).
func (v *Vector[T]) Empty() bool {
	return v.Size() == 0
}

// Capacity returns the number of elements that the container has currently allocated space for.
// Complexity - O(1).
func (v *Vector[T]) Capacity() int {
	return cap(v.ar)
}

// Front returns a reference to the first element in the container.
// Calling Front on an empty container is undefined.
// Complexity - O(1).
func (v *Vector[T]) Front() T {
	return v.ar[0]
}

// Back returns a reference to the last element in the container.
// Calling Back on an empty container causes undefined behavior.
// Complexity - O(1).
func (v *Vector[T]) Back() T {
	return v.ar[v.Size()-1]
}

func (v *Vector[T]) PopBack() T {
	res := v.Back()

	v.ar = v.ar[:v.Size()-1]

	return res
}

// Pop(pos int) T
// Erase(pos int)
// Insert(pos int, el T)
// Clear()
// Reverse()
