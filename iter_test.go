package sak

import (
	"reflect"
	"testing"
)

func TestIter(t *testing.T) {
	want := []int{0, 1, 2}
	i := Iter(want)
	got := []int{i.Next().(int), i.Next().(int), i.Next().(int)}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Iter: %v != %v", got, want)
	}

	i = Iter(Iter([]int{0, 1, 2}))
	got = []int{i.Next().(int), i.Next().(int), i.Next().(int)}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("IterIter: %v != %v", got, want)
	}

	// Test non-iterable panics.
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		Iter(31)
	}()
	if !panicked {
		t.Error("Iter(31) should have panicked")
	}
}

func TestIter_HasNext(t *testing.T) {
	i := new(arrayOrSliceIterator)
	i.data = reflect.ValueOf([]int{0})
	i.dataLen = i.data.Len()
	if !i.HasNext() {
		t.Error("HasNext() should be true")
	}
	i.Next()
	if i.HasNext() {
		t.Error("HasNext() should be false")
	}
}

func TestIter_Next(t *testing.T) {
	i := &arrayOrSliceIterator{data: reflect.ValueOf([]int{0})}
	if x := i.Next().(int); x != 0 {
		t.Errorf("%d != 0", x)
	}
	if i.idx != 1 {
		t.Error("idx should have incremented")
	}

	want := []int{0, 1, 2}
	i = &arrayOrSliceIterator{data: reflect.ValueOf(want)}
	got := []int{i.Next().(int), i.Next().(int), i.Next().(int)}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%v != %v", got, want)
	}

	// Test too many Next()s panics.
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		i.Next()
	}()
	if !panicked {
		t.Error("Next() should have panicked")
	}
}
