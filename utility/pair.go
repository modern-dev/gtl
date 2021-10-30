// Copyright 2021. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utility

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

func MakePair[T1 any, T2 any](first T1, second T2) *Pair[T1, T2] {
	return &Pair[T1, T2]{first, second}
}
