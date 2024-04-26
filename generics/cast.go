package generics

import (
	"fmt"

	"github.com/mandelsoft/goutils/errors"
)

// TryCast is like Cast, but reports
// whether the assertion is possible or not.
func TryCast[T, O any](o O) (T, bool) {
	var i any = o
	t, ok := i.(T)
	return t, ok
}

// TryCastE casts one type parameter to another type parameter,
// which have a sub type relation.
// This cannot be described by type parameter constraints in Go, because
// constraints may not be type parameters again.
func TryCastE[T any, O any](o O) (T, error) {
	var _nil T
	var s any = o
	if t, ok := s.(T); ok {
		return t, nil
	}
	return _nil, errors.ErrInvalid("type", fmt.Sprintf("%T", o))
}

// Cast asserts a type given by a type parameter for a value
// This is not directly suppoerted by Go.
//
//	func [O any](...) {
//	   x := i.(O)
//	}
func Cast[T, O any](o O) T {
	var i any = o
	t := i.(T)
	return t
}

// CastPointer maps a pointer P to an interface type I
// avoiding typed nil pointers. Nil pointers will be mapped
// to nil interfaces.
func CastPointer[I any, E any, P PointerType[E]](e P) I {
	var _nil I
	if e == nil {
		return _nil
	}
	var i any = e
	return i.(I)
}
