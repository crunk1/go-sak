package sak

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