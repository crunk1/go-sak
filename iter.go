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

type iter struct {
	data reflect.Value
	pos int
}

func (i *iter) HasNext() bool {
	return i.Len() > 0
}

func (i *iter) Len() int {
	return i.data.Len() - i.pos
}

func (i *iter) Next() interface{} {
	defer func() { i.pos++ }()
	return i.data.Index(i.pos).Interface()
}

// Iter returns an Iterator around an iterable, specifically: an array, a slice, or another Iterator.
// Calling Iter on an Iterator is idempotent and just returns the Iterator.
func Iter(iterable interface{}) Iterator {
	i := &iter{pos: 0}
	if iterator, ok := iterable.(Iterator); ok {
		return iterator
	}

	switch reflect.TypeOf(iterable).Kind() {
	case reflect.Slice, reflect.Array:
		i.data = reflect.ValueOf(iterable)
	default:
		panic("iterable must be slice, array, or Iterator")
	}
	return i
}

// Any returns true if any element in iter satisfies fn.
func Any(fn func(i interface{}) bool, iter Iterator) bool {
	return First(fn, iter) != nil
}

// All returns true if all elements in iter satisfy fn.
func All(fn func(i interface{}) bool, iter Iterator) bool {
	for iter.HasNext() {
		if t := fn(iter.Next()); !t {
			return false
		}
	}
	return true
}

// First returns the first element from iter that satisfies fn.
func First(fn func(i interface{}) bool, iter Iterator) interface{} {
	for iter.HasNext() {
		i := iter.Next()
		if t := fn(i); t {
			return i
		}
	}
	return nil
}

// Index returns the index of the element into the data structure behind iter.
// This only makes sense if the backing data structure is ordered and indexable.
// This uses shallow equivalence.
// Returns -1 if element is not found in iter.
func Index(element interface{}, iter Iterator) int {
	for idx := 0; iter.HasNext(); idx++ {
		i := iter.Next()
		if i == element {
			return idx
		}
	}
	return -1
}