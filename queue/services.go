package queue

import (
	"sync/atomic"
)

func NewQueueRest[T comparable]() *QueueRest[T] {
	return &QueueRest[T]{
		items:   make(chan T, 1),
		counter: 0,
	}
}
func (q *QueueRest[T]) Enqueue(item T) {
	// counter variable atomically incremented
	atomic.AddUint64(&q.counter, 1)
	// put item to channel
	q.items <- item
}

func (q *QueueRest[T]) Dequeue() T {
	// read item from channel
	item := <-q.items
	// counter variable decremented atomically.
	atomic.AddUint64(&q.counter, ^uint64(0))
	return item
}
