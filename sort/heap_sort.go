package sort

// down adjust from top to bottom
func down[T any](values []T, start, end int, cmp func(v1 T, v2 T) int) {
	lchild := 2*start + 1
	rchild := 2*start + 2
	parent := start
	for parent < end && lchild < end && rchild < end {
		child := 0
		if cmp(values[lchild], values[rchild]) >= 0 {
			child = lchild
		} else {
			child = rchild
		}
		if cmp(values[parent], values[child]) >= 0 {
			return
		}
		swap(values, parent, child)
		parent = child
		lchild = 2*parent + 1
		rchild = 2*parent + 2
	}
}

// swap exchange the position of two elements of an array
func swap[T any](values []T, i, j int) {
	values[i], values[j] = values[j], values[i]
}

func HeapSort[T any](values []T, cmp func(v1 T, v2 T) int) {
	// construct big root heap
	for i := len(values)/2 - 1; i >= 0; i-- {
		down(values, i, len(values), cmp)
	}
	// heap sort
	for i := 0; i < len(values); i++ {
		swap(values, 0, len(values)-i-1)
		down(values, 0, len(values)-i-1, cmp)
	}
}
