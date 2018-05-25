package sak


func StrIn(s string, ss []string) bool {
	fn := func(i interface{}) bool {
		return s == i
	}
	return Any(fn, Iter(ss))
}


func StrOr(s string, ss ...string) string {
	if s != "" || len(ss) == 0 {
		return s
	}
	fn := func(i interface{}) bool {
		return i != ""
	}
	result := First(fn, Iter(ss))
	if result == nil {
		return ""
	}
	return result.(string)
}