package sak

import (
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	isPos := func(i interface{}) bool {
		return i.(int) > 0
	}
	tests := []struct {
		desc     string
		iterable interface{}
		want     bool
	}{
		{"true case", []int{1, 2, 3}, true},
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
		desc     string
		iterable interface{}
		want     bool
	}{
		{"true case", []int{-1, 2, -3}, true},
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

func TestFilter(t *testing.T) {
	isPos := func(i interface{}) bool {
		return i.(int) > 0
	}
	tests := []struct {
		desc     string
		iterable interface{}
		want     []interface{}
	}{
		{"true case", []int{-1, 2, -3}, []interface{}{2}},
		{"filter all case", []int{-1, -3}, nil},
		{"filter empty case", []int{}, nil},
	}

	for _, tt := range tests {
		got := Filter(isPos, tt.iterable)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: got=%t want=%t", tt.desc, got, tt.want)
		}
	}
}

func TestIn(t *testing.T) {
	x := 3

	tests := []struct {
		desc     string
		iterable interface{}
		want     bool
	}{
		{"found case", []int{1, 2, 3}, true},
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

func TestIndex(t *testing.T) {
	x := 3

	tests := []struct {
		desc     string
		iterable interface{}
		want     int
	}{
		{"found case", []int{1, 2, 3}, 2},
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

func TestOr(t *testing.T) {
	isPos := func(i interface{}) bool {
		return i.(int) > 0
	}
	tests := []struct {
		desc     string
		iterable interface{}
		want     interface{}
	}{
		{"first found case", []int{-1, 2, -3}, 2},
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

func TestIntIn(t *testing.T) {
	tests := []struct {
		desc string
		i    int
		is   []int
		want bool
	}{
		{"normal", 1, []int{1, 2}, true},
		{"missing", 1, []int{2, 3}, false},
		{"empty list", 1, []int{}, false},
		{"nil list", 1, nil, false},
	}

	for _, tt := range tests {
		got := IntIn(tt.i, tt.is)
		if got != tt.want {
			t.Errorf("%s case: got=%t, want=%t", tt.desc, got, tt.want)
		}
	}
}

func TestIntOr(t *testing.T) {
	tests := []struct {
		desc   string
		inputs []int
		want   int
	}{
		{"zero", []int{0}, 0},
		{"first", []int{1, 2}, 1},
		{"second", []int{0, 2}, 2},
		{"all zero", []int{0, 0}, 0},
	}

	for _, tt := range tests {
		got := IntOr(tt.inputs[0], tt.inputs[1:]...)
		if got != tt.want {
			t.Errorf("%s case: got=%d, want=%d", tt.desc, got, tt.want)
		}
	}
}

func TestStrIn(t *testing.T) {
	tests := []struct {
		desc string
		s    string
		ss   []string
		want bool
	}{
		{"normal", "world", []string{"hey", "world"}, true},
		{"missing", "hello", []string{"hey", "world"}, false},
		{"empty list", "hello", []string{}, false},
		{"nil list", "hello", nil, false},
	}

	for _, tt := range tests {
		got := StrIn(tt.s, tt.ss)
		if got != tt.want {
			t.Errorf("%s case: got=%t, want=%t", tt.desc, got, tt.want)
		}
	}
}

func TestStrOr(t *testing.T) {
	tests := []struct {
		desc   string
		inputs []string
		want   string
	}{
		{"empty", []string{""}, ""},
		{"first", []string{"foo", "bar"}, "foo"},
		{"second", []string{"", "foo"}, "foo"},
		{"all empty", []string{"", ""}, ""},
	}

	for _, tt := range tests {
		got := StrOr(tt.inputs[0], tt.inputs[1:]...)
		if got != tt.want {
			t.Errorf("%s case: got=%s, want=%s", tt.desc, got, tt.want)
		}
	}
}

func TestUintIn(t *testing.T) {
	tests := []struct {
		desc string
		i    uint
		is   []uint
		want bool
	}{
		{"normal", 1, []uint{1, 2}, true},
		{"missing", 1, []uint{2, 3}, false},
		{"empty list", 1, []uint{}, false},
		{"nil list", 1, nil, false},
	}

	for _, tt := range tests {
		got := UintIn(tt.i, tt.is)
		if got != tt.want {
			t.Errorf("%s case: got=%t, want=%t", tt.desc, got, tt.want)
		}
	}
}

func TestUintOr(t *testing.T) {
	tests := []struct {
		desc   string
		inputs []uint
		want   uint
	}{
		{"zero", []uint{0}, 0},
		{"first", []uint{1, 2}, 1},
		{"second", []uint{0, 2}, 2},
		{"all zero", []uint{0, 0}, 0},
	}

	for _, tt := range tests {
		got := UintOr(tt.inputs[0], tt.inputs[1:]...)
		if got != tt.want {
			t.Errorf("%s case: got=%d, want=%d", tt.desc, got, tt.want)
		}
	}
}
