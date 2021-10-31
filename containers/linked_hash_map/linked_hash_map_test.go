// Copyright 2021. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package linked_hash_map

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestLinkedHashMapCapacityMembers(t *testing.T) {
	lhm := NewLinkedHashMap[int, string](false)
	lhmm := NewLinkedHashMap[int, string](true)

	if !lhm.Empty() {
		t.Errorf("expected lhm not to be empty")
	}

	if !lhmm.Empty() {
		t.Errorf("expected lhmm not to be empty")
	}

	for _, num := range []int{1, 2, 3, 4, 5, 5} {
		lhm.Insert(num, strconv.Itoa(num))
		lhmm.Insert(num, strconv.Itoa(num))
	}

	if lhm.Size() != 5 {
		t.Errorf("expected lhm to be the size of 5")
	}

	if lhmm.Size() != 6 {
		t.Errorf("expected lhmm to be the size of 6")
	}

	lhm.Clear()
	lhmm.Clear()

	if !lhm.Empty() {
		t.Errorf("expected lhm not to be empty")
	}

	if !lhmm.Empty() {
		t.Errorf("expected lhmm not to be empty")
	}
}

func TestLinkedHashMapModifierAndLookupMembers(t *testing.T) {
	cases := []struct {
		name     string
		insert   []int
		expected []int
	}{
		{"Example 1", []int{1, 2, 3, 4, 5}, []int{1, 3, 5}},
		{"Example 2", []int{-10, -50, 10, 50}, []int{-10, 10}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			lhm := NewLinkedHashMap[int, string](false)

			for _, num := range tt.insert {
				lhm.Insert(num, strconv.Itoa(num))
			}

			for _, num := range tt.expected {
				if !lhm.Contains(num) {
					t.Errorf("expected container to have value {%d}, but got false", num)
				}

				if lhm.At(num) != strconv.Itoa(num) {
					t.Errorf("expected key {%d} to have value {%s}", num, strconv.Itoa(num))
				}
			}
		})
	}
}

func TestLinkedHashMapModifierAndLookupMembers2(t *testing.T) {
	cases := []struct {
		name    string
		insert  []int
		removed []int
	}{
		{"Example 1", []int{1, 2, 3, 4, 5}, []int{1, 3, 5}},
		{"Example 2", []int{-10, -50, 10, 50}, []int{-10, 10}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			lhm := NewLinkedHashMap[int, string](false)

			for _, num := range tt.insert {
				lhm.Insert(num, strconv.Itoa(num))
			}

			for _, num := range tt.removed {
				lhm.Erase(num)
			}

			for _, num := range tt.removed {
				if lhm.Contains(num) {
					t.Errorf("expected container not to have value {%d}", num)
				}
			}
		})
	}
}

func TestLinkedHashMapModifierAndLookupMembers3(t *testing.T) {
	cases := []struct {
		name     string
		insert   []int
		modify   [][]int
		expected [][]int
	}{
		{
			"Example 1",
			[]int{1, 2, 3, 4, 5},
			[][]int{{1, 10}, {3, 30}, {5, 50}, {7, 70}},
			[][]int{{3, 30}, {7, 70}},
		}, {
			"Example 2",
			[]int{-10, -50, 10, 50},
			[][]int{{-10, -20}, {10, 20}, {100, 200}},
			[][]int{{-10, -20}, {100, 200}}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			lhm := NewLinkedHashMap[int, int](false)

			for _, num := range tt.insert {
				lhm.Insert(num, num)
			}

			for _, mod := range tt.modify {
				lhm.InsertOrAssign(mod[0], mod[1])
			}

			for _, exp := range tt.expected {
				if !lhm.Contains(exp[0]) {
					t.Errorf("expected container to have the key of {%d}", exp[0])
				}

				if lhm.At(exp[0]) != exp[1] {
					t.Errorf("expected container to have the key of {%d} be associated with {%d}, got {%d} instead",
						exp[0], exp[1], lhm.At(exp[0]))
				}
			}
		})
	}
}

func TestLinkedHashMultiMapModifierAndLookupMembers(t *testing.T) {
	cases := []struct {
		name     string
		insert   []int
		expected [][]int
	}{
		{"Example 1", []int{1, 2, 3, 4, 5, 3, 5, 5}, [][]int{{1, 1}, {3, 2}, {5, 3}, {7, 0}}},
		{"Example 2", []int{-10, -50, 10, 50, 10, 10, 1}, [][]int{{10, 3}, {1, 1}, {-20, 0}}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			lhm, rndInt := NewLinkedHashMap[int, string](true), rand.Intn(100)

			for _, num := range tt.insert {
				lhm.Insert(num, strconv.Itoa(num))
			}

			for _, exp := range tt.expected {
				if lhm.Count(exp[0]) != exp[1] {
					t.Errorf("expected container to have {%d} key {%d} times, got {%d} instead",
						exp[0], exp[1], lhm.Count(exp[0]))
				}
			}

			lhm.Clear()

			if !lhm.Empty() {
				t.Errorf("expected container to be empty")
			}

			defer func() { recover() }()

			lhm.Erase(rndInt)
			lhm.At(rndInt)

			t.Errorf("expected to panic")
		})
	}
}
