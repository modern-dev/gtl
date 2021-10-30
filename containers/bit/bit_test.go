// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package bit

import (
	"testing"
)

const MOD = 1e9 + 7

func TestNewBIT(t *testing.T) {
	expected := 10
	bit := NewBIT(expected)

	if bit.Len() != expected {
		t.Errorf("Expected to get %d, got %d", expected, bit.Len())
	}
}

func TestLeetcode1649(t *testing.T) {
	tests := []struct {
		name         string
		instructions []int
		expected     int
	}{
		{"Example 1", []int{1, 5, 6, 2}, 1},
		{"Example 2", []int{1, 2, 3, 6, 5, 4}, 3},
		{"Example 3", []int{1, 3, 3, 3, 2, 4, 2, 1, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := createSortedArray(tt.instructions)

			if result != tt.expected {
				t.Errorf("Expected to get %d, got %d", tt.expected, result)
			}
		})
	}
}

func createSortedArray(instructions []int) int {
	/*
	* 1649. Create Sorted Array through Instructions
	* https://leetcode.com/problems/create-sorted-array-through-instructions/
	*
	* Given an integer array instructions, you are asked to create a sorted array from the elements in instructions.
	* You start with an empty container nums. For each element from left to right in instructions, insert it into nums.
	* The cost of each insertion is the minimum of the following:
	*
	* The number of elements currently in nums that are strictly less than instructions[i].
	* The number of elements currently in nums that are strictly greater than instructions[i].
	* For example, if inserting element 3 into nums = [1,2,3,5], the cost of insertion is min(2, 1) (elements 1 and 2 are less than 3,
	* element 5 is greater than 3) and nums will become [1,2,3,3,5].
	*
	* Return the total cost to insert all elements from instructions into nums. Since the answer may be large, return it modulo 1e9 + 7
	 */
	ans, n, bit := 0, len(instructions), NewBIT(100002)

	for i := 0; i < n; i++ {
		lsum, rsum := bit.Query(instructions[i]-1), i-bit.Query(instructions[i])
		ans += min(lsum, rsum)
		ans %= MOD
		bit.Update(instructions[i], 1)
	}

	return ans % MOD
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
