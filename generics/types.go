package generics

import (
	"reflect"
)

type PointerType[P any] interface {
	*P
}

// Deprecated: use PointerTo.
func Pointer[T any](t T) *T {
	return &t
}

// PointerTo returns a pointer to a given value.
func PointerTo[T any](t T) *T {
	return &t
}

func TypeOf[T any]() reflect.Type {
	var t T
	return reflect.TypeOf(&t).Elem()
}
