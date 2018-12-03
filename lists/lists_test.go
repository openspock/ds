package lists

import (
	"fmt"
	"testing"
)

func TestAddElemToEmptyLList(t *testing.T) {
	l := MakeLinkedList()
	l.AddTail(0)
	if l.Count() != uint64(1) {
		t.Error("count: should be 0")
	}
}

func TestRemoveTailFromLinkedList(t *testing.T) {
	l := MakeLinkedList()
	l.AddTail(0)
	val, _ := l.RemoveTail()
	if val != 0 {
		t.Error("elem: incorrect, should be 0")
	}
	if l.Count() != 0 {
		t.Error("count: list should be empty")
	}
}

func TestAddElemAtRandomIndexToLinkedList(t *testing.T) {
	l := makeIntLinkedList(10000)
	l.AddAt(10001, 335)
	if l.Count() != 10001 {
		t.Errorf("count: should be 10001, is %d", l.Count())
	}

	val, _ := l.RemoveAt(335)
	if val != 10001 {
		t.Error("elem: incorrect element returned")
	}
}

func TestListRemoveTail(t *testing.T) {
	l := makeIntLinkedList(10)
	for {
		if l.Count() == 0 {
			break
		}
		_, err := l.RemoveTail()
		if err != nil {
			t.Error(err)
		}
	}
}

func TestLinkedListIterator(t *testing.T) {
	l := makeIntLinkedList(10)
	if err := l.Iterate(func(i interface{}) {
		fmt.Println(i)
	}); err != nil {
		t.Error(err)
	}
}

func makeIntLinkedList(count int) *LinkedList {
	l := MakeLinkedList()
	for i := 0; i < count; i++ {
		l.AddTail(i)
	}
	return l
}
