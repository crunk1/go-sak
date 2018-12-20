package sak

import(
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	isPos := func(i interface{}) bool {
		return i.(int) > 0
	}
	tests := []struct {
		desc string
		iterable interface{}
		want bool
	} {
		{"true case", []int{1,2,3}, true},
		{"true empty case", []int{}, true},
		{"false case", []int{-1}, false},
	}

	for _, tt := range tests {
		got := All(isPos, tt.iterable)
		if got != tt.want {
			t.Errorf("%s: got=%t want=%t", tt.desc, got, tt.want)
		}
	}
}

func TestAny(t *testing.T) {
	isPos := func(i interface{}) bool {
		return i.(int) > 0
	}
	tests := []struct {
		desc string
		iterable interface{}
		want bool
	} {
		{"true case", []int{-1,2,-3}, true},
		{"false case", []int{-1, -3}, false},
		{"false empty case", []int{}, false},
	}

	for _, tt := range tests {
		got := Any(isPos, tt.iterable)
		if got != tt.want {
			t.Errorf("%s: got=%t want=%t", tt.desc, got, tt.want)
		}
	}
}

func TestOr(t *testing.T) {
	isPos := func(i interface{}) bool {
		return i.(int) > 0
	}
	tests := []struct {
		desc string
		iterable interface{}
		want interface{}
	} {
		{"first found case", []int{-1,2,-3}, 2},
		{"none found case", []int{-1, -3}, nil},
		{"none found empty case", []int{}, nil},
	}

	for _, tt := range tests {
		got := Or(isPos, tt.iterable)
		if got != tt.want {
			t.Errorf("%s: got=%v want=%v", tt.desc, got, tt.want)
		}
	}
}

func TestIndex(t *testing.T) {
	x := 3

	tests := []struct {
		desc string
		iterable interface{}
		want int
	} {
		{"found case", []int{1,2,3}, 2},
		{"not found case", []int{-1, -3}, -1},
		{"not found empty case", []int{}, -1},
	}
	for _, tt := range tests {
		got := Index(x, tt.iterable)
		if got != tt.want {
			t.Errorf("%s: got=%d want=%d", tt.desc, got, tt.want)
		}
	}
}

func TestIn(t *testing.T) {
	x := 3

	tests := []struct {
		desc string
		iterable interface{}
		want bool
	} {
		{"found case", []int{1,2,3}, true},
		{"not found case", []int{-1, -3}, false},
		{"not found empty case", []int{}, false},
	}
	for _, tt := range tests {
		got := In(x, tt.iterable)
		if got != tt.want {
			t.Errorf("%s: got=%t want=%t", tt.desc, got, tt.want)
		}
	}
}

func TestIter(t *testing.T) {
	want := []int{0,1,2}
	i := Iter(want)
	got := []int{i.Next().(int), i.Next().(int), i.Next().(int)}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Iter: %v != %v", got, want)
	}

	i = Iter(Iter([]int{0,1,2}))
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
	i := &sakIterator{data: reflect.ValueOf([]int{0})}
	if !i.HasNext() {
		t.Error("HasNext() should be true")
	}
	i.Next()
	if i.HasNext() {
		t.Error("HasNext() should be false")
	}
}

func TestIter_Len(t *testing.T) {
	tests := []struct {
		desc string
		data []int
		want int
	} {
		{"positive case", []int{-1,2,-3}, 3},
		{"empty case", []int{}, 0},
	}

	for _, tt := range tests {
		i := &sakIterator{data: reflect.ValueOf(tt.data)}
		got := i.Len()
		if got != tt.want {
			t.Errorf("%s: got=%d want=%d", tt.desc, got, tt.want)
		}
	}

	i := &sakIterator{data: reflect.ValueOf([]int{0,1,2})}
	i.Next()
	if i.Len() != 2 {
		t.Error("Len() should have returned 2 (one less than 3 after a Next())")
	}
}

func TestIter_Next(t *testing.T) {
	i := &sakIterator{data: reflect.ValueOf([]int{0})}
	if x := i.Next().(int); x != 0 {
		t.Errorf("%d != 0", x)
	}
	if i.pos != 1 {
		t.Error("pos should have incremented")
	}

	want := []int{0,1,2}
	i = &sakIterator{data: reflect.ValueOf(want)}
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