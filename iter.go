package sak

import "reflect"

type Iterable interface {
	// Iter returns an Iterator for the backing data structure.
	Iter() Iterator
}

type Iterator interface {
	// HasNext returns true if there is more data in this Iterator.
	HasNext() bool
	// Next returns the next datum in this Iterator.
	// Next will panic
	Next() interface{}
}

// arrayOrSliceIterator is a wrapper for Arrays and Slices passed to Iter.
type arrayOrSliceIterator struct {
	data    reflect.Value
	dataLen int
	idx     int
}

func (i *arrayOrSliceIterator) HasNext() bool {
	return i.idx < i.dataLen
}

func (i *arrayOrSliceIterator) Next() interface{} {
	defer func() { i.idx++ }()
	return i.data.Index(i.idx).Interface()
}

// Iter returns an Iterator around an iterable, specifically:
//   an array, a slice, an Iterable, or another Iterator.
// Calling Iter on an Iterator is idempotent and just returns the Iterator.
// If iterable is a slice or array, concurrent modifications to the
// iterable could cause unexpected results.
func Iter(iterable interface{}) Iterator {
	if iterator, ok := iterable.(Iterator); ok {
		return iterator
	}
	if iterable, ok := iterable.(Iterable); ok {
		return iterable.Iter()
	}

	i := &arrayOrSliceIterator{idx: 0}
	switch reflect.TypeOf(iterable).Kind() {
	case reflect.Slice, reflect.Array:
		i.data = reflect.ValueOf(iterable)
		i.dataLen = i.data.Len()
	default:
		panic("iterable must be slice, array, or Iterator")
	}
	return i
}
