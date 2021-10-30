// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package union_find

type DisjointSet struct {
	size         int
	rank, parent []int
}

// NewDisjointSet returns an object representing a unordered_set containing n element.
func NewDisjointSet(n int) *DisjointSet {
	inst := &DisjointSet{
		n,
		make([]int, n),
		make([]int, n),
	}

	for i := 0; i < n; i++ {
		inst.parent[i] = i
	}

	return inst
}

// Len returns the size of the disjoint unordered_set.
// Complexity - O(1).
func (this *DisjointSet) Len() int {
	return this.size
}

// Find returns the root of the tree that contains x.
// Complexity - O(α(n)), where α is the inverse Ackermann function.
func (this *DisjointSet) Find(x int) int {
	if this.parent[x] != x {
		this.parent[x] = this.Find(this.parent[x])
	}

	return this.parent[x]
}

// Union merges the sets represented by x node and the y node into a single unordered_set.
// Returns whether or not the nodes were disjoint before the union operation (i.e. if the operation had an effect).
// Complexity - O(α(n)), where α is the inverse Ackermann function.
func (this *DisjointSet) Union(x, y int) bool {
	xSet, ySet := this.Find(x), this.Find(y)

	if xSet == ySet {
		return false
	}

	if this.rank[xSet] < this.rank[ySet] {
		this.parent[xSet] = ySet
	} else if this.rank[xSet] > this.rank[ySet] {
		this.parent[ySet] = xSet
	} else {
		this.parent[ySet] = xSet
		this.rank[xSet] = this.rank[xSet] + 1
	}

	return true
}
