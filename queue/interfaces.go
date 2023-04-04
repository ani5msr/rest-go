package queue

type QueueRest[T comparable] struct {
	items   chan T
	counter uint64
}
