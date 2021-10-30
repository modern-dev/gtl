// Copyright 2020. The GTL Authors. All rights reserved.
// https://github.com/modern-dev/gtl
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package queue

import (
	"testing"
)

const queueEnqueuesCount = 300

func TestNewQueue(t *testing.T) {
	var (
		structQueue = &Queue[struct{}]{}
		intQueue    = &Queue[int]{}
		sliceQueue  = &Queue[[]float64]{}
	)

	checkQueueSize(structQueue, 0, t)
	checkQueueSize(intQueue, 0, t)
	checkQueueSize(sliceQueue, 0, t)
}

func TestEnqueue(t *testing.T) {
	queue := &Queue[int]{}

	for i := 0; i < queueEnqueuesCount; i++ {
		queue.Push(i)
		checkQueueSize(queue, i+1, t)
	}
}

func TestDequeue(t *testing.T) {
	queue := &Queue[int]{}

	queue.Push(1)
	el := queue.Pop()

	checkQueueSize(queue, 0, t)
	if el != 1 {
		t.Errorf("Expected to get %d, got %d", 1, el)
	}
}

func TestPeek(t *testing.T) {
	queue := &Queue[int]{}

	queue.Push(1)
	el := queue.Front()

	checkQueueSize(queue, 1, t)
	if el != 1 {
		t.Errorf("Expected to get %d, got %d", 1, el)
	}
}

func TestQueue(t *testing.T) {
	queue := &Queue[float64]{}

	for i := 0; i < queueEnqueuesCount; i++ {
		queue.Push(float64(i))
	}

	checkQueueSize(queue, queueEnqueuesCount, t)

	if queue.Empty() {
		t.Errorf("Queue should not be empty")
	}

	for i := queueEnqueuesCount; i > 0; i-- {
		queue.Pop()
	}

	if !queue.Empty() {
		t.Errorf("Deque should be empty")
	}
}

func checkQueueSize[T any](queue *Queue[T], expected int, t *testing.T) {
	if queue.Size() != expected {
		t.Errorf("Queue should have size %d but got %d", expected, queue.Size())
	}
}
