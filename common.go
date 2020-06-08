package sak

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
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

func JSONMustMarshal(v interface{}) []byte {
	bs, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Sprintf("%+v", errors.WithStack(err)))
	}
	return bs
}

func JSONMustUnmarshal(data []byte, v interface{}) {
	err := errors.WithStack(json.Unmarshal(data, v))
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
}
