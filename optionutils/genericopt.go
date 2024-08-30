package optionutils

import (
	"fmt"

	"github.com/mandelsoft/goutils/generics"
	"github.com/mandelsoft/goutils/reflectutils"
)

type genricoption[S any, B any, T any] struct {
	value T
}

func (o *genricoption[S, B, T]) ApplyTo(opts B) {
	reflectutils.CallMethodByInterfaceVA[S](opts, o.value)
}

// WithGenericOption probides a generic option implementation for Option[T]
// intended for options based on an option setter interface S implemented
// by the option set B implementing S for the value type T. Hereby, B must
// implement S, which cannot be expressed by Go generics.
func WithGenericOption[S, B any, T any](v T) Option[B] {
	var b B

	if _, ok := generics.TryCast[S](b); !ok {
		panic(fmt.Sprintf("%T must be %s", b, generics.TypeOf[S]()))
	}
	return &genricoption[S, B, T]{v}
}