package priority_queue

import (
	"fmt"
)

var (
	ErrPriorityQueueIsNil = fmt.Errorf("queue is nil\n")
)

type PriorityQueue[T any] struct {
	values []T
	cmp    func(v1 T, v2 T) int
}

func New[T any](cmp func(v1 T, v2 T) int) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		values: make([]T, 0),
		cmp:    cmp,
	}
}

func (p *PriorityQueue[T]) Empty() bool {
	return p.Len() == 0
}

func (p *PriorityQueue[T]) Len() int {
	return len(p.values)
}

func (p *PriorityQueue[T]) Push(v T) {
	p.values = append(p.values, v)
	p.up()
}

func (p *PriorityQueue[T]) Pop() {
	if p.Empty() {
		panic(ErrPriorityQueueIsNil)
	}
	p.swap(0, p.Len()-1)
	p.values = p.values[:p.Len()-1]
	p.down()
}

func (p *PriorityQueue[T]) Top() T {
	if p.Empty() {
		panic(ErrPriorityQueueIsNil)
	}
	return p.values[0]
}

// down adjust big root heap from top to bottom
func (p *PriorityQueue[T]) down() {
	i := 0
	for i < p.Len() && 2*i+1 < p.Len() && 2*i+2 < p.Len() {
		if p.cmp(p.values[i], p.values[2*i+1]) < 0 {
			if p.cmp(p.values[2*i+1], p.values[2*i+2]) >= 0 {
				// swap p.values[i] and its left child
				p.swap(i, 2*i+1)
				i = 2*i + 1
			} else {
				// swap p.values[i] ans its right child
				p.swap(i, 2*i+2)
				i = 2*i + 2
			}
		} else if p.cmp(p.values[i], p.values[2*i+2]) < 0 {
			// swap p.values[i] and its left child
			p.swap(i, 2*i+2)
			i = 2*i + 2
		} else {
			return
		}
	}
}

// up adjust big root heap from bottom to top
// because the first index of heap is 0
// so if a child's index is i, the index of it's parent is (i+1)/2 -1
func (p *PriorityQueue[T]) up() {
	i := p.Len() - 1
	parent := (i+1)/2 - 1
	for i >= 0 && parent >= 0 {
		if p.cmp(p.values[i], p.values[parent]) > 0 {
			p.swap(i, parent)
		} else {
			return
		}
		i = parent
		parent = (i+1)/2 - 1
	}
}

func (p *PriorityQueue[T]) swap(i, j int) {
	p.values[i], p.values[j] = p.values[j], p.values[i]
}
