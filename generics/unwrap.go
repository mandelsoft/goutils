package generics

import (
	"reflect"

	"github.com/mandelsoft/goutils/general"
	"github.com/modern-go/reflect2"
)

type Unwrappable[I any] interface {
	Unwrap() I
}

// UnwrapWith unwraps at most one level with
// unwrapper of given unwrap type.
func UnwrapWith[I any](o any) I {
	var _nil I
	if !reflect2.IsNil(o) {
		if u, ok := o.(Unwrappable[I]); ok {
			return u.Unwrap()
		}
	}
	return _nil
}

// UnwrapAllWith tries to unwrap an object using
// a given unwrap type until it reaches
// an object not implementing the unwrapper
// or unwrapping to nil.
// The last non-nil object implementing I is returned.
// Otherwise, the zero value of type I
// is returned.
func UnwrapAllWith[I any](o any) I {
	var _nil I
	if reflect2.IsNil(o) {
		return _nil
	}
	for {
		if u, ok := o.(Unwrappable[I]); ok {
			n := u.Unwrap()
			if reflect2.IsNil(n) {
				break
			}
			o = n
		}
	}
	if i, ok := o.(I); ok {
		return i
	}
	return _nil
}

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

// UnwrapAny unwraps an object as log as there is
// any Unwrap method.
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

func MapUnwrapped[R any, I any](o interface{}, f general.MapperFunc[I, R]) R {
	var _nil R

	if u, ok := UnwrapUntil[I](o); ok {
		return f(u)
	}
	return _nil
}
