package sak

import "testing"

func TestIntIn(t *testing.T) {
	tests := []struct {
		desc string
		i int
		is []int
		want bool
	} {
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
		desc string
		inputs []int
		want int
	} {
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
		s string
		ss []string
		want bool
	} {
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
		desc string
		inputs []string
		want string
	} {
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
		i uint
		is []uint
		want bool
	} {
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
		desc string
		inputs []uint
		want uint
	} {
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
