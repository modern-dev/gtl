// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package unordered_set

import (
	"testing"
)

const addsCount = 3000

func TestNewUnorderedSet(t *testing.T) {
	s := NewUnorderedSet[int]()

	checkSize(s, 0, t)
	checkEmpty(s, true, t)
}

func TestInsert(t *testing.T) {
	s := NewUnorderedSet[int]()

	checkSet(s, 0, true, t)

	for i := 0; i < addsCount; i++ {
		s.Insert(i)
		checkSet(s, i+1, false, t)
	}
}

func TestEmpty(t *testing.T) {
	s := NewUnorderedSet[int]()

	checkEmpty(s, true, t)

	s.Insert(1)
	s.Insert(42)

	checkEmpty(s, false, t)
}

func TestContains(t *testing.T) {
	s := NewUnorderedSet[int]()

	checkSet(s, 0, true, t)

	for i := 0; i < addsCount/2; i = i + 2 {
		s.Insert(i)

		if s.Contains(i) != true {
			t.Errorf("Expected unordered_set to contain %d", i)
		}

		if s.Contains(i+1) == true {
			t.Errorf("Expected unordered_set to not contain %d", i+1)
		}
	}
}

func TestErase(t *testing.T) {
	s := NewUnorderedSet[int]()

	for i := 0; i < addsCount; i++ {
		s.Insert(i)
	}

	checkSet(s, addsCount, false, t)

	for i := 0; i < addsCount; i++ {
		checkSet(s, addsCount-i, false, t)
		s.Erase(i)
	}

	checkSet(s, 0, true, t)

	for i := 0; i < addsCount; i++ {
		checkSet(s, 0, true, t)
		s.Erase(i)
	}
}

func TestClear(t *testing.T) {
	s := NewUnorderedSet[int]()

	for i := 0; i < addsCount; i++ {
		s.Insert(i)
	}

	checkSet(s, addsCount, false, t)

	s.Clear()

	checkEmpty(s, true, t)
}

func checkSet[T comparable](s *UnorderedSet[T], size int, isEmpty bool, t *testing.T) {
	checkSize(s, size, t)
	checkEmpty(s, isEmpty, t)
}

func checkSize[T comparable](s *UnorderedSet[T], size int, t *testing.T) {
	if size != s.Size() {
		t.Errorf("Expected UnorderedSet size %d, got %d", size, s.Size())
	}
}

func checkEmpty[T comparable](s *UnorderedSet[T], isEmpty bool, t *testing.T) {
	if isEmpty != s.Empty() {
		t.Errorf("Expected IsEmpty to be %v, got %v", isEmpty, s.Empty())
	}
}
