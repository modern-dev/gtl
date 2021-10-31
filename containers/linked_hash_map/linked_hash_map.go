// Copyright 2021. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package linked_hash_map

import (
	"github.com/modern-dev/gtl"
)

type (
	// LinkedHashMap TODO
	LinkedHashMap[TKey comparable, TVal any] struct {
		store      map[TKey][]*node[TVal]
		head, tail *node[TVal]
		dupl       bool
		size       int
	}

	node[TVal any] struct {
		val        TVal
		prev, next *node[TVal]
	}
)

// NewLinkedHashMap TODO
func NewLinkedHashMap[TKey comparable, TVal any](allowDuplicates bool) *LinkedHashMap[TKey, TVal] {
	head, tail := &node[TVal]{}, &node[TVal]{}
	head.next = tail
	tail.prev = head

	return &LinkedHashMap[TKey, TVal]{
		make(map[TKey][]*node[TVal]),
		head,
		tail,
		allowDuplicates,
		0,
	}
}

/* Capacity members */

// Empty checks if the container has no elements.
// Complexity — constant.
func (lhm *LinkedHashMap[TKey, TVal]) Empty() bool {
	return lhm.Size() == 0
}

// Size returns the number of elements in the container.
// Complexity — constant.
func (lhm *LinkedHashMap[TKey, TVal]) Size() int {
	return lhm.size
}

/* Modifiers members */

// Clear erases all elements from the container. After this call, Size() returns zero.
// Complexity — constant.
func (lhm *LinkedHashMap[TKey, TVal]) Clear() {
	lhm.head.next = lhm.tail
	lhm.tail.prev = lhm.head
	lhm.store = make(map[TKey][]*node[TVal])
	lhm.size = 0
}

// Insert inserts element into the container.
// Complexity — average case - O(1), worst case - O(size()).
func (lhm *LinkedHashMap[TKey, TVal]) Insert(key TKey, val TVal) bool {
	if _, found := lhm.store[key]; found && !lhm.dupl {
		return false
	}

	newNode := &node[TVal]{
		val,
		nil,
		nil,
	}

	lhm.store[key] = append(lhm.store[key], newNode)
	lhm.size++

	lhm.appendNode(newNode)

	return true
}

// Assign assigns to the current element if the key already exists.
// Throws gtl.OutOfRange if the container does not have an element with the specified key.
// Complexity — constant.
func (lhm *LinkedHashMap[TKey, TVal]) Assign(key TKey, val TVal) {
	lhm.getNode(key, 0).val = val
}

// InsertOrAssign inserts an element or assigns to the current element if the key already exists.
// Returns true if the insertion took place and false if the assignment took place.
// Complexity — average case - O(1), worst case - O(Size()).
func (lhm *LinkedHashMap[TKey, TVal]) InsertOrAssign(key TKey, val TVal) bool {
	if lhm.Contains(key) {
		lhm.Assign(key, val)

		return false
	}

	return lhm.Insert(key, val)
}

// Erase removes all elements with the key equivalent to key.
// Complexity — average case - O(1), worst case - O(Size()).
func (lhm *LinkedHashMap[TKey, TVal]) Erase(key TKey) {
	if _, found := lhm.store[key]; !found {
		return
	}

	for _, remNode := range lhm.store[key] {
		lhm.removeNode(remNode)
	}

	lhm.size -= len(lhm.store[key])

	delete(lhm.store, key)
}

/* Lookup members */

// At returns a reference to the mapped value of the element with key equivalent to key.
// If no such element exists, a panic of type gtl.OutOfRange is thrown.
// Complexity — constant.
func (lhm *LinkedHashMap[TKey, TVal]) At(key TKey) TVal {
	return lhm.getNode(key, 0).val
}

// Count returns the number of elements with key that compares equal to the specified argument key.
// Complexity — linear in the number of elements with key on average, worst case linear in the size of the container.
func (lhm *LinkedHashMap[TKey, TVal]) Count(key TKey) int {
	items, found := lhm.store[key]

	if !found {
		return 0
	}

	return len(items)
}

// Contains checks if there is an element with key equivalent to key in the container.
// Complexity — constant.
func (lhm *LinkedHashMap[TKey, TVal]) Contains(key TKey) bool {
	_, found := lhm.store[key]

	return found
}

/* Private members */

func (lhm *LinkedHashMap[TKey, TVal]) appendNode(newNode *node[TVal]) {
	next := lhm.head.next

	next.prev = newNode
	newNode.next = next
	lhm.head.next = newNode
	newNode.prev = lhm.head
}

func (lhm *LinkedHashMap[TKey, TVal]) removeNode(remNode *node[TVal]) {
	prev, next := remNode.prev, remNode.next

	prev.next = next
	next.prev = prev
}

func (lhm *LinkedHashMap[TKey, TVal]) getNode(key TKey, idx int) *node[TVal] {
	if items, found := lhm.store[key]; lhm.Empty() || !found || len(items) == 0 {
		panic(gtl.OutOfRange)
	}

	return lhm.store[key][idx]
}
