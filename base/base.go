// Package base offers some basic data structures for golang.
package base

// Pair is a plain key value pair.
type Pair struct {
	Key, Value interface{}
}

// Adder is an interface that permits it's implementors to add elements
// a data structure, like a list.
//
// AddHead can be easily executed by calling AddTail(elem, 0)
type Adder interface {

	// Add an elemen to the tail end
	AddTail(Elem interface{})

	// Add an element, at a specific index.
	AddAt(Elem interface{}, index uint64) error
}

// Remover is an interface that permits it's implementors to add elements
// a data structure, like a list.
//
// RemoveHead can be easily executed by calling AddTail(elem, 0)
type Remover interface {

	// Remove and return the elemtent at index.
	RemoveAt(index uint64) (interface{}, error)

	// Remove and return the tail element.
	RemoveTail() (interface{}, error)
}

// ListAdderRemover is an interface that allows implementations to
// add/ remove elements.
type ListAdderRemover interface {
	Adder
	Remover
}

// SizeCounter allows it's implementations to return a count of elements
// contained in a data structure.
type SizeCounter interface {

	// Count returns the count of elements.
	Count() uint64
}

// Iterator offers it's implementors to execute a lambda function on each
// element in the data structure.
type Iterator interface {

	// Iterate will execute the lambda function f on each element of the
	// data structure.
	//
	// It'll throw an error on the first occurence.
	Iterate(f func(i interface{}) error) error
}
