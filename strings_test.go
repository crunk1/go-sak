package sak

import "testing"

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
