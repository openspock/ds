// Package lists contain various implementations of the list data structures.
package lists

import (
	"fmt"
)

type node struct {
	elem     interface{}
	next     *node
	previous *node
}

// LinkedList implementation of a list.
type LinkedList struct {
	head *node
	tail *node
}

// MakeLinkedList makes and returns an empty linked list.
func MakeLinkedList() *LinkedList {
	return &LinkedList{nil, nil}
}

// Count returns the count of elements in this linked list.
func (list *LinkedList) Count() uint64 {
	i := uint64(0)
	for node := list.head; node != nil; node = node.next {
		i = i + 1
	}
	return i
}

// AddTail adds an element to the tail of this linked list.
func (list *LinkedList) AddTail(elem interface{}) {
	if list.tail == nil {
		list.head = &node{elem, nil, nil}
		list.tail = list.head
		return
	}
	newNode := &node{elem, nil, list.tail}
	if list.head.next == nil {
		list.head.next = newNode
	}
	tail := list.tail
	tail.next = newNode
	list.tail = newNode
}

// AddAt adds an element at a specific index.
func (list *LinkedList) AddAt(elem interface{}, index uint64) error {
	n, err := list.nodeAt(index)
	if err != nil {
		return err
	}
	if n == nil { // this means, we're adding to the tail of the list
		list.AddTail(elem)
		return nil
	}
	newNode := node{elem, n, n.previous}
	n.previous.next = &newNode
	n.previous = &newNode

	return nil
}

// RemoveTail removes the element at the tail, update tail to tail - 1.
// Return the tail element.
func (list *LinkedList) RemoveTail() (interface{}, error) {
	count := list.Count()
	if count == 0 {
		return nil, fmt.Errorf("count: cannot remove tail as count is 0")
	} else if count == 1 {
		elem := list.head.elem
		list.head = nil
		list.tail = nil
		return elem, nil
	}
	elem := list.tail.elem
	newTail := list.tail.previous
	list.tail = newTail
	list.tail.next = nil
	return elem, nil
}

// RemoveAt removes an element at an index and returns it.
func (list *LinkedList) RemoveAt(index uint64) (interface{}, error) {
	n, err := list.nodeAt(index)
	if err != nil {
		return nil, err
	}
	prev := n.previous
	next := n.next
	prev.next = next
	next.previous = prev
	return n.elem, nil
}

// Iterate will execute a lambda function over each element of the
// linked list.
func (list *LinkedList) Iterate(f func(i interface{}) error) error {
	for node := list.head; node != nil; node = node.next {
		f(node.elem)
	}
	return nil
}

func (list *LinkedList) nodeAt(index uint64) (*node, error) {
	err := list.validateCount(index)
	if err != nil {
		return nil, err
	}
	node := list.head
	for i := uint64(0); i <= index; i++ {
		node = node.next
	}
	return node, nil
}

func (list *LinkedList) validateCount(index uint64) error {
	count := list.Count()
	if count+1 < index {
		return fmt.Errorf("count: size of the list is %d", count)
	}
	return nil
}
