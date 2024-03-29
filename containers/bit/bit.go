// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// See the implementation details https://www.hackerearth.com/practice/notes/binary-indexed-tree-or-fenwick-tree/

package bit

type BIT struct {
	tree []int
	size int
}

// NewBIT returns an object representing Binary Index Tree or Fenwick Tree of given size.
func NewBIT(size int) *BIT {
	return &BIT{
		make([]int, size),
		size,
	}
}

// Update updates a node in Binary Index Tree (BIT) at given index in BIT.
// The given value `val` is added to BIT[i] and all of its ancestors in tree.
// Complexity - O(LogN).
func (this *BIT) Update(index, val int) {
	for ; index <= this.size; index += index & -index {
		this.tree[index] += val
	}
}

// Query returns a sum of first [0...index] elements of Binary Index Tree.
// Complexity - O(LogN).
func (this *BIT) Query(index int) int {
	sum := 0

	for ; index > 0; index -= index & -index {
		sum += this.tree[index]
	}

	return sum
}

// Len returns the size of the BIT.
// Complexity - O(1).
func (this *BIT) Len() int {
	return this.size
}
