package queues

import (
	"testing"
	"time"
)

func TestFifoCount(t *testing.T) {
	fifo := makeFifoAndPopulateRandomInts(10)
	if fifo == nil {
		t.Error("nil: expected to be non nil")
	}
	if fifo.Count() != 10 {
		t.Errorf("count: expected 10, is %d", fifo.Count())
	}
}

func TestFifoEnQDeQSingleInt(t *testing.T) {
	fifo := MakeFifo(1)
	if fifo.Count() != 0 {
		t.Errorf("count: expected 0, is %d", fifo.Count())
	}
	x := 513
	fifo.Enqueue(x)
	if fifo.Count() != 1 {
		t.Errorf("count: expected 1, is %d", fifo.Count())
	}
	time.Sleep(1 * time.Second)
	if y, _ := fifo.Dequeue(); x != y {
		t.Errorf("val: should be %d, is %d", x, y)
	}
}

func makeFifoAndPopulateRandomInts(len uint64) *ConcurrentEnqueueFifo {
	q := MakeFifo(len)
	for i := uint64(0); i < len; i++ {
		_ = q.Enqueue(i)
	}
	time.Sleep(3 * time.Second)
	return q
}
