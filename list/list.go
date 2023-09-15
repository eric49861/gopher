package list

import "fmt"

var (
	ErrListIsNil = fmt.Errorf("nil list\n")
)

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

func (n *Node[T]) Value() T {
	return n.value
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

// NewList create a new linked list
func NewList[T any]() *List[T] {
	dummy := &Node[T]{
		next: nil,
		prev: nil,
	}
	return &List[T]{
		head: dummy,
		tail: dummy,
		len:  0,
	}
}

// Insert nodes to list tail
func (l *List[T]) Insert(values ...T) error {
	if l == nil {
		return ErrListIsNil
	}
	for _, v := range values {
		l.tail.next = &Node[T]{
			value: v,
			next:  nil,
			prev:  l.tail,
		}
		l.tail = l.tail.next
	}
	l.len += len(values)
	return nil
}

func (l *List[T]) InsertHead(value T) error {
	if l == nil {
		return ErrListIsNil
	}
	n := &Node[T]{
		value: value,
		next:  l.head.next,
		prev:  l.head,
	}
	l.head.next = n
	if l.Empty() {
		l.tail = n
	} else {
		n.next.prev = n
	}
	l.len++
	return nil
}

// Remove delete node of list head
// if list is nil, panic
func (l *List[T]) Remove() error {
	if l == nil || l.Empty() {
		return ErrListIsNil
	}
	l.head.next = l.head.next.next
	if l.Len() == 1 {
		l.tail = l.head
	} else {
		l.head.next.prev = l.head
	}
	l.len--
	return nil
}

// RemoveTail delete node of list tail
func (l *List[T]) RemoveTail() error {
	if l == nil || l.Empty() {
		return ErrListIsNil
	}
	l.tail = l.tail.prev
	l.tail.next = nil
	l.len--
	return nil
}

func (l *List[T]) Head() *Node[T] {
	if l == nil || l.Empty() {
		return nil
	}
	return l.head.next
}

func (l *List[T]) Tail() *Node[T] {
	if l == nil || l.Empty() {
		return nil
	}
	return l.tail
}

// Len get number of elements of list
func (l *List[T]) Len() int {
	if l == nil {
		return 0
	}
	return l.len
}

// Empty if the list is empty, it will return true, or return false
func (l *List[T]) Empty() bool {
	return l.Len() == 0
}

// Clear delete all nodes of list
func (l *List[T]) Clear() {
	if l == nil {
		return
	}
	l.head.next = nil
	l.tail = l.head
	l.len = 0
}

func (l *List[T]) ForEach(callback func(node *Node[T])) error {
	if l == nil {
		return ErrListIsNil
	}
	temp := l.head.next
	for temp != nil {
		callback(temp)
		temp = temp.next
	}
	return nil
}
