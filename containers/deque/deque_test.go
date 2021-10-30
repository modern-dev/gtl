// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package deque

import (
	"testing"
)

const enqueuesCount = 200

func TestNewDeque(t *testing.T) {
	var (
		structDeque = &Deque[struct{}]{}
		intDeque    = &Deque[int]{}
		sliceDeque  = &Deque[[]float64]{}
	)

	checkDequeSize(structDeque, 0, t)
	checkDequeSize(intDeque, 0, t)
	checkDequeSize(sliceDeque, 0, t)
}

func TestPushFront(t *testing.T) {
	d := &Deque[int]{}

	for i := 0; i < enqueuesCount; i++ {
		d.PushFront(i)
		checkDequeSize(d, i+1, t)
	}
}

func TestPushBack(t *testing.T) {
	d := &Deque[int]{}

	for i := 0; i < enqueuesCount; i++ {
		d.PushBack(i)
		checkDequeSize(d, i+1, t)
	}
}

func TestPushPopFront(t *testing.T) {
	d := &Deque[int]{}

	d.PushFront(1)
	el := d.PopFront()

	checkDequeSize(d, 0, t)
	if el != 1 {
		t.Errorf("Expected to get %d, got %d", 1, el)
	}
}

func TestPushFrontPopBack(t *testing.T) {
	d := &Deque[int]{}

	d.PushFront(1)
	el := d.PopBack()

	checkDequeSize(d, 0, t)
	if el != 1 {
		t.Errorf("Expected to get %d, got %d", 1, el)
	}
}

func TestPopFront(t *testing.T) {
	d := &Deque[string]{}

	d.PushBack("string")

	value := d.Front()
	if value != "string" {
		t.Errorf("Expected to get %s, got %s", "string", value)
	}

	d.PushBack("value")
	d.PushFront("front")

	value = d.Front()
	if value != "front" {
		t.Errorf("Expected to get %s, got %s", "front", value)
	}
}

func TestPopBack(t *testing.T) {
	d := &Deque[string]{}

	d.PushBack("string")

	value := d.Back()
	if value != "string" {
		t.Errorf("Expected to get %s, got %s", "string", value)
	}

	d.PushBack("value")
	d.PushFront("front")

	value = d.Back()
	if value != "value" {
		t.Errorf("Expected to get %s, got %s", "value", value)
	}
}

func TestDeque(t *testing.T) {
	d := &Deque[int]{}

	for i := 0; i < enqueuesCount/2; i++ {
		d.PushFront(i)
	}

	checkDequeSize(d, enqueuesCount/2, t)

	for i := enqueuesCount / 2; i < enqueuesCount; i++ {
		d.PushBack(i)
	}

	checkDequeSize(d, enqueuesCount, t)

	if d.Empty() {
		t.Errorf("deque should not be empty")
	}

	for i := enqueuesCount; i > 0; i-- {
		d.PopFront()
	}

	if d.Empty() != true {
		t.Errorf("deque should be empty")
	}
}

func checkDequeSize[T any](Deque *Deque[T], expected int, t *testing.T) {
	if Deque.Size() != expected {
		t.Errorf("deque should have size %d but got %d", expected, Deque.Size())
	}
}
