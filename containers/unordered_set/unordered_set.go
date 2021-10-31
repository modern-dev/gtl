// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package unordered_set

import (
	"github.com/modern-dev/gtl/containers/linked_hash_map"
)

type (
	// UnorderedSet is unordered set based on linked_hash_map.LinkedHashMap.
	// It can contain comparable elements only.
	UnorderedSet[T comparable] struct {
		lhm *linked_hash_map.LinkedHashMap[T, interface{}]
	}
)

// NewUnorderedSet TODO
func NewUnorderedSet[T comparable]() *UnorderedSet[T] {
	return &UnorderedSet[T]{
		linked_hash_map.NewLinkedHashMap[T, interface{}](false),
	}
}

/* Capacity members */

// Empty checks if there are elements in UnorderedSet.
// Complexity — constant.
// Returns true if the UnorderedSet is empty, false otherwise.
func (us *UnorderedSet[T]) Empty() bool {
	return us.lhm.Empty()
}

// Size returns the number of elements in the container.
// Complexity — constant.
func (us *UnorderedSet[T]) Size() int {
	return us.lhm.Size()
}

/* Modifiers members */

// Clear erases all elements from the container. After this call, Size() returns zero.
// Complexity — linear in the size of the container.
func (us *UnorderedSet[T]) Clear() {
	us.lhm.Clear()
}

// Insert inserts element into UnorderedSet.
// Complexity — average case - O(1), worst case O(Size()).
func (us *UnorderedSet[T]) Insert(item T) {
	us.lhm.Insert(item, nil)
}

// Erase deletes the element from unordered_set if it contains an element
// does nothing otherwise.
// Complexity — constant.
func (us *UnorderedSet[T]) Erase(item T) {
	us.lhm.Erase(item)
}

/* Lookup members */

// Contains checks if UnorderedSet contains given element.
// Complexity — constant.
// Returns true if UnorderedSet includes the element, false otherwise.
func (us *UnorderedSet[T]) Contains(item T) bool {
	return us.lhm.Contains(item)
}
