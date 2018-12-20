package sak

import "reflect"

type Iterable interface {
	// Iter returns an Iterator for the backing data structure.
	Iter() Iterator
}

type Iterator interface {
	// HasNext returns true if there is more data in this Iterator.
	HasNext() bool
	// Len returns the number of remaining data in this Iterator.
	Len() int
	// Next returns the next datum in this Iterator.
	// Next will panic
	Next() interface{}
}

// sakIterator is a wrapper for Arrays and Slices passed to Iter.
type sakIterator struct {
	data reflect.Value
	pos int
}

func (i *sakIterator) HasNext() bool {
	return i.Len() > 0
}

func (i *sakIterator) Len() int {
	return i.data.Len() - i.pos
}

func (i *sakIterator) Next() interface{} {
	defer func() { i.pos++ }()
	return i.data.Index(i.pos).Interface()
}

// Iter returns an Iterator around an iterable, specifically:
//   an array, a slice, an Iterable, or another Iterator.
// Calling Iter on an Iterator is idempotent and just returns the Iterator.
func Iter(iterable interface{}) Iterator {
	if iterator, ok := iterable.(Iterator); ok {
		return iterator
	}
	if iterable, ok := iterable.(Iterable); ok {
		return iterable.Iter()
	}

	i := &sakIterator{pos: 0}
	switch reflect.TypeOf(iterable).Kind() {
	case reflect.Slice, reflect.Array:
		i.data = reflect.ValueOf(iterable)
	default:
		panic("iterable must be slice, array, or Iterator")
	}
	return i
}

// Any returns true if any element in iter satisfies fn.
func Any(fn func(elem interface{}) bool, iterable interface{}) bool {
	return Or(fn, iterable) != nil
}

// All returns true if all elements in iter satisfy fn.
func All(fn func(elem interface{}) bool, iterable interface{}) bool {
	iter := Iter(iterable)
	for iter.HasNext() {
		if t := fn(iter.Next()); !t {
			return false
		}
	}
	return true
}

// Or returns the first element from iter that satisfies fn.
// Returns nil if no element satifies fn.
func Or(fn func(elem interface{}) bool, iterable interface{}) interface{} {
	iter := Iter(iterable)
	for iter.HasNext() {
		elem := iter.Next()
		if t := fn(elem); t {
			return elem
		}
	}
	return nil
}

// Index returns the index of elem into the data structure behind iterable.
// This only makes sense if the backing data structure is ordered and indexable.
// Returns -1 if elem is not found in iterable.
// elem and elements in iterable must be comparable.
func Index(elem interface{}, iterable interface{}) int {
	iter := Iter(iterable)
	for idx := 0; iter.HasNext(); idx++ {
		elemFromIterable := iter.Next()
		if elem == elemFromIterable {
			return idx
		}
	}
	return -1
}

// In returns true if elem is found in iterable.
// elem and elements in iterable must be comparable.
func In(elem interface{}, iterable interface{}) bool {
	iter := Iter(iterable)
	for idx := 0; iter.HasNext(); idx++ {
		elemFromIterable := iter.Next()
		if elemFromIterable == elem {
			return true
		}
	}
	return false
}

