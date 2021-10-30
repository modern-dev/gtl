// Copyright 2021. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utility

import "constraints"

type (
	Compare[T any] interface {
		Cmp(T, T) bool
	}
	Less[T constraints.Ordered]    struct{}
	Greater[T constraints.Ordered] struct{}
)

func (l *Less[T]) Cmp(lhs T, rhs T) bool {
	return lhs < rhs
}

func (g *Greater[T]) Cmp(lhs T, rhs T) bool {
	return lhs > rhs
}
