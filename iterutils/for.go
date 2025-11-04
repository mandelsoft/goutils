package iterutils

import (
	"github.com/mandelsoft/goutils/general"
	"iter"
)

// For provides an iterator for a sequence of elements.
func For[T any](v ...T) iter.Seq[T] {
	return func(yield func(v T) bool) {
		for _, e := range v {
			if !yield(e) {
				return
			}
		}
	}
}

// ForMapped provides an iterator providing elements of type O for a given
// sequence of elements of type T mapped to a target type O.
func ForMapped[O, T any](mapper general.MapperFunc[T, O], v ...T) iter.Seq[O] {
	return func(yield func(v O) bool) {
		for _, e := range v {
			if !yield(mapper(e)) {
				return
			}
		}
	}
}
