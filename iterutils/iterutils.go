package iterutils

import (
	"github.com/mandelsoft/goutils/general"
	"iter"
)

// Convert transforms an iterator of type I into an iterator of type O by
// applying a type assertion during iteration.
func Convert[O, I any](in iter.Seq[I]) iter.Seq[O] {
	return func(yield func(O) bool) {
		in(func(v I) bool {
			return yield(any(v).(O))
		})
	}
}

// ConvertFunc transforms an input iterator using a provided mapper function
// and returns a new iterator of mapped outputs.
func ConvertFunc[O, I any](in iter.Seq[I], mapper general.MapperFunc[I, O]) iter.Seq[O] {
	return func(yield func(O) bool) {
		in(func(v I) bool {
			return yield(mapper(v))
		})
	}
}

func Reverse[T any](in iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		elems := Get(in)
		for i := range elems {
			if !yield(elems[len(elems)-i-1]) {
				return
			}
		}
	}
}
