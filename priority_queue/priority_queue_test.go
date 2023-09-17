package priority_queue

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Struct struct {
	v int
}

func TestPriorityQueue_Push(t *testing.T) {
	p := New[Struct](func(v1 Struct, v2 Struct) int {
		if v1.v < v2.v {
			return -1
		} else if v1.v == v2.v {
			return 0
		}
		return 1
	})
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	ss := make([]Struct, 0)
	for i := 0; i < 100; i++ {
		ss = append(ss, Struct{
			v: r.Int() % 1000,
		})
	}
	for i := 0; i < 100; i++ {
		p.Push(ss[i])
	}
	//
	for !p.Empty() {
		fmt.Printf("%d ", p.Top())
		p.Pop()
	}
}
