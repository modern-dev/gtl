// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package priority_queue

import "testing"

func TestPriorityQueue(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		pq := NewPriorityQueue[int]()

		pq.Push(5)
		pq.Push(6)
		pq.Push(7)
		pq.Push(9)
		pq.Push(14)
		pq.Push(11)
		pq.Push(10)

		want := 14

		if got := pq.Pop(); got != want {
			t.Errorf("Pop() = %v, want %v", got, want)
		}

		if pq.Size() != 6 {
			t.Errorf("Expected PriorityQueue size to be {%d}, got {%d} instead", 6, pq.Size())
		}
	})
}
