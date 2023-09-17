package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestHeapSort(t *testing.T) {
	arr := make([]int, 1000)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 1000; i++ {
		arr[i] = r.Intn(10000)
	}
	HeapSort[int](arr, func(v1 int, v2 int) int {
		return v1 - v2
	})
	fmt.Printf("%v\n", arr)
}
