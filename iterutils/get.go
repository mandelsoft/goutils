package iterutils

import (
	"github.com/mandelsoft/goutils/generics"
	"iter"
)

// Get provides a value slice from iterating ofer an iterator.
func Get[T any](seq iter.Seq[T]) []T {
	list := []T{}
	for e := range seq {
		list = append(list, e)
	}
	return list
}

// ToSliceOf is like Get but allows to generalize the base type
// of the slice. I or all elements must be assignable to O.
func ToSliceOf[O, I any](seq iter.Seq[I]) []O {
	list := []O{}
	for e := range seq {
		list = append(list, generics.Cast[O](e))
	}
	return list
}
