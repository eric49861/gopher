package queue

import "fmt"

var (
	ErrQueueIsNil = fmt.Errorf("queue is nil")
)

type Queue[T any] struct {
	values []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		values: make([]T, 0),
	}
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}

func (q *Queue[T]) Len() int {
	return len(q.values)
}

func (q *Queue[T]) Enqueue(v T) {
	q.values = append(q.values, v)
}

func (q *Queue[T]) Dequeue() {
	if q.Empty() {
		panic(ErrQueueIsNil)
	}
	q.values = q.values[1:]
}

func (q *Queue[T]) Front() T {
	if q.Empty() {
		panic(ErrQueueIsNil)
	}
	return q.values[0]
}
