// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package rbtree

import (
	"constraints"
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := NewRBTree[int](false)

	assertTreeSize(tree, 0, t)
}

func TestTreeLen(t *testing.T) {
	tree := NewRBTree[int](false)

	assertTreeSize(tree, 0, t)

	items := []int{6, 2, 5, 1, 4, 21, 42}

	for i, item := range items {
		tree.Insert(item)
		assertTreeSize(tree, i+1, t)
	}

	assertTreeSize(tree, len(items), t)
}

func TestTreeSearch(t *testing.T) {
	runTreeSearch[int]([]int{5, 3, 1, 2, 4, 12, 10, 42, 13}, []int{0, 6, 11, 17, 18}, t)

	runTreeSearch[float64]([]float64{4, 12, 10, 32, 51, 25, 14, 4, 3}, []float64{0, 1, 5, 7, 13, 33}, t)
}

func TestTreeInorder(t *testing.T) {
	elements1 := []int{5, 3, 1, 2, 4, 12, 10, 42, 13}
	tree1 := treeFromSlice[int](elements1)

	elements2 := []float64{4, 12, 10, 32, 51, 25, 14, 4, 3}
	tree2 := treeFromSlice[float64](elements2)

	assertTreeSize(tree1, len(elements1), t)
	assertTreeSize(tree2, len(elements2), t)
}

func TestTreeMax(t *testing.T) {
	maxInt := 412
	intTree := treeFromSlice[int]([]int{25, 12, 35, 14, 52, 15, 235, 51, 2, 124, 5, 235, maxInt, 55})

	maxFloat := 321.512
	floatTree := treeFromSlice[float64]([]float64{10.5, 25, 12, 24.52, maxFloat, 33.42, 52.55})

	if intTree.Max() != maxInt {
		t.Errorf("Max() should return max value in tree. Got %v, expected %v", intTree.Max(), maxInt)
	}

	if floatTree.Max() != maxFloat {
		t.Errorf("Max() should return max value in tree. Got %v, expected %v", floatTree.Max(), maxFloat)
	}
}

func TestTreeMin(t *testing.T) {
	minInt := 2
	intTree := treeFromSlice[int]([]int{25, 12, 35, 14, 52, 15, 235, 51, minInt, 124, 5, 235, 55})

	minFloat := 1.111
	floatTree := treeFromSlice[float64]([]float64{10.5, 25, 12, 24.52, minFloat, 33.42, 52.55})

	if intTree.Min() != minInt {
		t.Errorf("Min() should return min value in tree. Got %v, expected %v", intTree.Min(), minInt)
	}

	if floatTree.Min() != minFloat {
		t.Errorf("Min() should return min value in tree. Got %v, expected %v", floatTree.Min(), minFloat)
	}
}

func TestTreeDelete(t *testing.T) {
	items := []int{3, 2, 5, 7, 4, 1, 5, 6, 2}
	tree := treeFromSlice[int](items)

	cases := []struct {
		item                   int
		shouldExistAfterDelete bool
		size                   int
	}{
		{5, true, len(items) - 1},
		{5, false, len(items) - 2},
		{3, false, len(items) - 3},
		{3, false, len(items) - 3},
	}

	runTreeDelete[int](tree, cases, t)

	floatItems := []float64{0.25, 52.3, 52, 12.22, 31, 16, 22.7, 11.5, 12.22, 12.11}
	floatTree := treeFromSlice[float64](floatItems)

	floatCases := []struct {
		item                   float64
		shouldExistAfterDelete bool
		size                   int
	}{
		{0.25, false, len(floatItems) - 1},
		{16.0, false, len(floatItems) - 2},
		{12.22, true, len(floatItems) - 3},
		{12.22, false, len(floatItems) - 4},
	}

	runTreeDelete[float64](floatTree, floatCases, t)
}

func runTreeSearch[T constraints.Ordered](existingElements, notExistingElements []T, t *testing.T) {
	tree := treeFromSlice[T](existingElements)

	for _, el := range existingElements {
		assertTreeValueSearch[T](tree, el, true, t)
	}

	for _, el := range notExistingElements {
		assertTreeValueSearch[T](tree, el, false, t)
	}
}

func runTreeDelete[T comparable](tree *RBTree[T], cases []struct {
	item                   T
	shouldExistAfterDelete bool
	size                   int
}, t *testing.T) {
	for _, c := range cases {
		tree.Erase(c.item)
		assertTreeValueSearch[T](tree, c.item, c.shouldExistAfterDelete, t)
		assertTreeSize(tree, c.size, t)
	}
}

func treeFromSlice[T constraints.Ordered](elements []T) *RBTree[T] {
	tree := NewRBTree[T](false)

	for _, el := range elements {
		tree.Insert(el)
	}

	return tree
}

func assertTreeSize[T comparable](tree *RBTree[T], expected int, t *testing.T) {
	if expected != tree.Size() {
		t.Errorf("Expected tree Len() to be %d, got %d", expected, tree.Size())
	}
}

func assertTreeValueSearch[T comparable](tree *RBTree[T], value T, expected bool, t *testing.T) {
	if _, exist := tree.Find(value); exist != expected {
		t.Errorf("Search test failed for element %v. Got %v, expected %v", value, exist, expected)
	}
}
