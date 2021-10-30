// Copyright 2021. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utility

import "testing"

func TestLess(t *testing.T) {
	l := Less[int]{}

	if l.Cmp(1, 2) != true {
		t.Errorf("expected 1 to be less than 2, got false instead")
	}
}

func TestGreater(t *testing.T) {
	g := Greater[int]{}

	if g.Cmp(2, 1) != true {
		t.Errorf("expected 2 to be greater than 1, got false instead")
	}
}
