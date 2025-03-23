package generics

import (
	"reflect"

	"github.com/modern-go/reflect2"
)

// UnwrapUntil calls *Unwrap() any* until
// the given type/interface is reached.
// If this is not possible false it returned.
func UnwrapUntil[T any](e any) (T, bool) {
	var _nil T

	for {
		if u, ok := e.(T); ok {
			return u, true
		}
		e = UnwrapAny(e)
		if reflect2.IsNil(e) {
			return _nil, false
		}
	}
}

func UnwrapAny(e any) any {
	if e == nil {
		return nil
	}
	v := reflect.ValueOf(e)
	m := v.MethodByName("Unwrap")
	if !m.IsValid() {
		return nil
	}

	t := m.Type()
	if t.NumIn() != 0 || t.NumOut() != 1 {
		return nil
	}
	r := m.Call(nil)
	return r[0].Interface()
}
