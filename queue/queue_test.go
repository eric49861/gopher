package queue

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 100; i++ {
		q.Enqueue(r.Intn(100))
	}
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", q.Front())
		q.Dequeue()
	}
}
