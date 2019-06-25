package sak

import (
	"reflect"
)

func IsZero(i interface{}) bool {
	v := reflect.ValueOf(i)
	k := v.Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	case reflect.Invalid:
		return true
	default:
		return i == reflect.Zero(v.Type()).Interface()
	}
}
