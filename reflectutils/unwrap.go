package reflectutils

import (
	"reflect"

	"github.com/mandelsoft/goutils/generics"
	"github.com/modern-go/reflect2"
)

type Unwrappable[T any] interface {
	Unwrap() T
}

func Unwrap[T any](o any) T {
	return generics.Cast[T](UnwrapAny(o))
}

func UnwrapAny(o any) any {
	if reflect2.IsNil(o) {
		return o
	}

	v := reflect.ValueOf(o)

	m, ok := v.Type().MethodByName("Unwrap")
	if !ok || m.Type.NumIn() != 1 || m.Type.NumOut() != 1 {
		return nil
	}

	return m.Func.Call([]reflect.Value{v})[0].Interface()
}
