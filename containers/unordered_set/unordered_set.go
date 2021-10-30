// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package unordered_set

// UnorderedSet is unordered unordered_set based on standard map
// It can contain comparable elements only
type UnorderedSet[T comparable] struct {
	table map[T]bool
}

func NewUnorderedSet[T comparable]() *UnorderedSet[T] {
	return &UnorderedSet[T]{
		make(map[T]bool),
	}
}

// Size returns the number of elements in the container.
// Complexity - O(1).
func (s *UnorderedSet[T]) Size() int {
	return len(s.table)
}

// Empty checks if there are elements in Set.
// Complexity - O(1).
// Returns true if the unordered_set is empty, false otherwise.
func (s *UnorderedSet[T]) Empty() bool {
	return s.Size() == 0
}

// Insert inserts element into Set.
// Complexity - O(1).
func (s *UnorderedSet[T]) Insert(item T) {
	s.table[item] = true
}

// Contains checks if Set contains given element.
// Complexity - O(1).
// returns true if Set includes the element, false otherwise.
func (s *UnorderedSet[T]) Contains(item T) bool {
	_, exists := s.table[item]

	return exists
}

// Erase deletes the element from unordered_set if it contains an element
// does nothing otherwise.
// Complexity - O(1).
func (s *UnorderedSet[T]) Erase(item T) {
	if s.Contains(item) {
		delete(s.table, item)
	}
}
