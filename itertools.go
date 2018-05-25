package sak

import "reflect"

type Iterable interface {
	Iter() Iterator
}

type Iterator interface {
	HasNext() bool
	Len() int
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

func Any(fn func(i interface{}) bool, iter Iterator) bool {
	return First(fn, iter) != nil
}

func All(fn func(i interface{}) bool, iter Iterator) bool {
	for iter.HasNext() {
		if t := fn(iter.Next()); !t {
			return false
		}
	}
	return true
}

func First(fn func(i interface{}) bool, iter Iterator) interface{} {
	for iter.HasNext() {
		i := iter.Next()
		if t := fn(i); t {
			return i
		}
	}
	return nil
}
