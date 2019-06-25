package sak

import (
	"testing"
)

type Test_IsZero_I interface {
	Do()
}
type Test_isZero_IImplementer struct {}
func (t *Test_isZero_IImplementer) Do() {}

func Test_IsZero(t *testing.T) {
	// Some test inputs.
	var i1, i2 Test_IsZero_I
	i2 = &Test_isZero_IImplementer{}
	var ch1, ch2 chan int
	ch2 = make(chan int)
	var fn1, fn2 func()
	fn2 = func() {}
	var m1, m2 map[int]int
	m2 = make(map[int]int)
	var p1, p2 *int
	someInt := 0
	p2 = &someInt
	var ss1, ss2 []string
	ss2 = make([]string, 0)
	type someStruct struct {
		X int
		Y *int
		Z string
	}

	cases := []struct {
		input interface{}
		want  bool
	}{
		{nil, true},  // 0
		{i1, true},
		{i2, false},
		{false, true},
		{true, false},
		{int(0), true},  // 5
		{int8(0), true},
		{int16(0), true},
		{int32(0), true},
		{int64(0), true},
		{uint(0), true},  // 10
		{uint8(0), true},
		{uint16(0), true},
		{uint32(0), true},
		{uint64(0), true},
		{int(1), false},  // 15
		{int8(1), false},
		{int16(1), false},
		{int32(1), false},
		{int64(1), false},
		{uint(1), false},  // 20
		{uint8(1), false},
		{uint16(1), false},
		{uint32(1), false},
		{uint64(1), false},
		{float32(0), true},  // 25
		{float64(0), true},
		{float32(1), false},
		{float64(1), false},
		{complex(float32(0), float32(0)), true},
		{complex(float32(0), float32(1)), false},  // 30
		{complex(float32(1), float32(0)), false},
		{complex(float32(1), float32(1)), false},
		{complex(float64(0), float64(0)), true},
		{complex(float64(0), float64(1)), false},
		{complex(float64(1), float64(0)), false},  // 35
		{complex(float64(1), float64(1)), false},
		{[2]int{}, true},
		{[2]int{1}, false},
		{ch1, true},
		{ch2, false},  // 40
		{fn1, true},
		{fn2, false},
		{m1, true},
		{m2, false},
		{p1, true},  // 45
		{p2, false},
		{ss1, true},
		{ss2, false},
		{"", true},
		{"foo", false},  // 50
		{someStruct{}, true},
		{someStruct{X:1}, false},
		{someStruct{Y:&someInt}, false},
		{someStruct{Z:"foo"}, false},
	}

	for i, c := range cases {
		got := IsZero(c.input)
		if got != c.want {
			t.Errorf("case %d: got=%t want=%t", i, got, c.want)
		}
	}
}
