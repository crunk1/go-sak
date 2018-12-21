package sak

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

// Filter creates a slice of elements that satisfy fn.
func Filter(fn func(elem interface{}) bool, iterable interface{}) []interface{} {
	iter := Iter(iterable)
	var results []interface{}
	for iter.HasNext() {
		elem := iter.Next()
		if fn(elem) {
			results = append(results, elem)
		}
	}
	return results
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

func IntIn(i int, is []int) bool {
	fn := func(x interface{}) bool {
		return i == x
	}
	return Any(fn, is)
}

func IntOr(i int, is ...int) int {
	if i != 0 || len(is) == 0 {
		return i
	}
	fn := func(x interface{}) bool {
		return x != 0
	}
	result := Or(fn, is)
	if result == nil {
		return 0
	}
	return result.(int)
}

func StrIn(s string, ss []string) bool {
	fn := func(i interface{}) bool {
		return s == i
	}
	return Any(fn, ss)
}

func StrOr(s string, ss ...string) string {
	if s != "" || len(ss) == 0 {
		return s
	}
	fn := func(i interface{}) bool {
		return i != ""
	}
	result := Or(fn, ss)
	if result == nil {
		return ""
	}
	return result.(string)
}

func UintIn(i uint, is []uint) bool {
	fn := func(x interface{}) bool {
		return i == x
	}
	return Any(fn, is)
}

func UintOr(i uint, is ...uint) uint {
	if i != 0 || len(is) == 0 {
		return i
	}
	fn := func(x interface{}) bool {
		return x != 0
	}
	result := Or(fn, is)
	if result == nil {
		return 0
	}
	return result.(uint)
}
