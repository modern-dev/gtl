// Copyright 2021. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utility

import "testing"

func TestMakePair(t *testing.T) {
	p1 := MakePair[int, string](1488, "Deus vult")
	expectedFirst, expectedSecond := 1488, "Deus vult"

	if p1.First != expectedFirst || p1.Second != expectedSecond {
		t.Errorf("expected {%d, %s}, got {%d, %s}", expectedFirst, expectedSecond, p1.First, p1.Second)
	}
}
