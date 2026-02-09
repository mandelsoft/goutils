package generics

import (
	"github.com/mandelsoft/goutils/general"
	"reflect"
)

type PointerType[P any] interface {
	*P
}

// Deprecated: use PointerTo.
func Pointer[T any](t T) *T {
	return &t
}

// PointerValue provides the value of a pointer it is not nil,
// otherwise a default value or the zero value is returned.
func PointerValue[T any](p *T, def ...T) T {
	if p != nil {
		return *p
	}
	return general.OptionalNonZero(def...)
}

// PointerTo returns a pointer to a given value.
func PointerTo[T any](t T) *T {
	return &t
}

func TypeOf[T any]() reflect.Type {
	var t T
	return reflect.TypeOf(&t).Elem()
}

// CanAssign checks whether an I is assignable to a variable of type O.
func CanAssign[I, O any]() bool {
	return TypeOf[I]().AssignableTo(TypeOf[O]())
}

// Implements checks whether an I implements an interface O.
func Implements[I, O any]() bool {
	return TypeOf[I]().Implements(TypeOf[O]())
}

// ObjectFor provides an object of type T.
// If T is a pointer type a pointer to an appropriate object is returned.
// For a non-pointer type the object is returned as value.
func ObjectFor[T any]() T {
	t := TypeOf[T]()
	if t.Kind() == reflect.Ptr {
		return reflect.New(t.Elem()).Interface().(T)
	}
	return reflect.New(t).Elem().Interface().(T)
}
