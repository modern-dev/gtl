// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package union_find

import (
	"sort"
	"testing"
)

func TestNewDisjointSet(t *testing.T) {
	expected := 5
	ds := NewDisjointSet(expected)

	if ds.Len() != expected {
		t.Errorf("Expected to get %d, got %d", expected, ds.Len())
	}
}

func TestUnionFind(t *testing.T) {
	ds := NewDisjointSet(5)
	ds.Union(0, 2)
	ds.Union(4, 2)
	ds.Union(3, 1)

	given1, given2 := ds.Find(4) == ds.Find(0), ds.Find(1) == ds.Find(0)

	if !given1 {
		t.Errorf("Expected to get %t, got %t", true, given1)
	}

	if given2 {
		t.Errorf("Expected to get %t, got %t", false, given2)
	}
}

func TestLeetcode1101(t *testing.T) {
	logs, N := [][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}}, 6
	expected, given := 20190301, earliestAcq(logs, N)

	if given != expected {
		t.Errorf("Expected to get %d, got %d", given, expected)
	}
}

func earliestAcq(logs [][]int, N int) int {
	/*
	* 1101. The Earliest Moment When Everyone Become Friends
	* https://leetcode.com/problems/the-earliest-moment-when-everyone-become-friendsz
	*
	* In a social group, there are N people, with unique integer ids from 0 to N-1.
	* We have a list of logs, where each logs[i] = [timestamp, id_A, id_B] contains a non-negative integer timestamp, and the ids of two different people.
	* Each log represents the time in which two different people became friends.  Friendship is symmetric: if A is friends with B, then B is friends with A.
	* Let's say that person A is acquainted with person B if A is friends with B, or A is a friend of someone acquainted with B.
	* Return the earliest time for which every person became acquainted with every other person. Return -1 if there is no such earliest time.
	 */

	ds := NewDisjointSet(N)

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	for i := 0; i < len(logs); i++ {
		if ds.Union(logs[i][1], logs[i][2]) {
			N--
		}

		if N == 1 {
			return logs[i][0]
		}
	}

	return -1
}
