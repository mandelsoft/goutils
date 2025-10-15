package iterutils

import "iter"

func SingleValue[T any](v T) iter.Seq[T] {
	return func(yield func(v T) bool) {
		yield(v)
	}
}
