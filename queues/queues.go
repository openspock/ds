// Package queues offers various queue implementations.
package queues

import (
	"fmt"

	"github.com/openspock/ds/lists"
)

// ConcurrentEnqueueFifo is a first in first out implementation of the queue data structure.
// The queue is concurrent for Enqueue operation only.
type ConcurrentEnqueueFifo struct {

	// Internally, this queue stores it's data in a linked list.
	// Concurrency is handled outside this store.
	store *lists.LinkedList
	len   uint64
	qchan chan interface{}
}

// MakeFifo makes a FIFO queue of a given length.
func MakeFifo(length uint64) *ConcurrentEnqueueFifo {
	list := lists.MakeLinkedList()
	q := &ConcurrentEnqueueFifo{
		store: list,
		len:   length,
		qchan: make(chan interface{}),
	}
	go func() {
		for {
			select {
			case elem := <-q.qchan:
				q.store.AddAt(elem, uint64(0))
			}
		}
	}()
	return q
}

// Count returns the count of elements in the queue.
func (q *ConcurrentEnqueueFifo) Count() uint64 {
	if q.store == nil {
		return uint64(0)
	}
	return q.store.Count()
}

// Enqueue adds an element to the tail of the FIFO queue
// returning an error if the queue is full.
func (q *ConcurrentEnqueueFifo) Enqueue(i interface{}) error {
	//return q.store.AddAt(i, uint64(0))
	if q.Count()+1 > q.len {
		return fmt.Errorf("len: should be %d, queue is full", q.len)
	}
	q.qchan <- i
	return nil
}

// Dequeue removes an element from the head of the FIFO queue
// returning it, and, optionally an error if the queue is empty.
func (q *ConcurrentEnqueueFifo) Dequeue() (interface{}, error) {
	return q.store.RemoveTail()
}
