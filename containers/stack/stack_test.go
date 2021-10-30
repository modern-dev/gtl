// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package stack

import "testing"

func TestNewStack(t *testing.T) {
	var (
		structStack = NewStack[struct{}]()
		intStack    = NewStack[int]()
		sliceStack  = NewStack[[]float64]()
	)

	checkStackSize(structStack, 0, t)
	checkStackSize(intStack, 0, t)
	checkStackSize(sliceStack, 0, t)
}

func checkStackSize[T any](s *Stack[T], expected int, t *testing.T) {
	if s.Size() != expected {
		t.Errorf("stack should have size %d but got %d", expected, s.Size())
	}
}
